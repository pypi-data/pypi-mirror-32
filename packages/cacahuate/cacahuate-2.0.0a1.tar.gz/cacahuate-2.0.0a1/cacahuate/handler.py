from coralillo.errors import ModelNotFoundError
from datetime import datetime
from importlib import import_module
from pymongo import MongoClient
import simplejson as json
import pika
import pymongo

from cacahuate.errors import CannotMove, ElementNotFound, InconsistentState, \
    MisconfiguredProvider
from cacahuate.logger import log
from cacahuate.models import Execution, Pointer, User
from cacahuate.xml import Xml
from cacahuate.node import make_node, Exit, Action, Validation
from cacahuate.grammar import Condition


class Handler:
    ''' The actual process machine, it is in charge of moving the pointers
    among the graph of nodes '''

    def __init__(self, config):
        self.config = config
        self.mongo = None

    def __call__(self, channel, method, properties, body: bytes):
        ''' the main callback of cacahuate '''
        message = json.loads(body)

        if message['command'] == 'cancel':
            self.cancel_execution(message)
        elif message['command'] == 'step':
            try:
                self.call(message, channel)
            except (ModelNotFoundError, CannotMove, ElementNotFound,
                    MisconfiguredProvider, InconsistentState
                    ) as e:
                log.error(str(e))
        else:
            log.warning('Unrecognized command {}'.format(message['command']))

        if not self.config['RABBIT_NO_ACK']:
            channel.basic_ack(delivery_tag=method.delivery_tag)

    def call(self, message: dict, channel):
        pointer, user, input = self.recover_step(message)
        execution = pointer.proxy.execution.get()

        xml = Xml.load(self.config, execution.process_name, direct=True)
        xmliter = iter(xml)

        node = make_node(xmliter.find(
            lambda e: e.getAttribute('id') == pointer.node_id
        ))

        to_queue = []  # pointers to be sent to the queue

        # node's lifetime ends here
        self.teardown(node, pointer, user, input)

        # retrieve the state so we can use it in wakeup and to find next node
        collection = self.get_mongo()[
            self.config['MONGO_EXECUTION_COLLECTION']
        ]
        state = next(collection.find({'id': execution.id}))

        next_nodes = self.next(xml, node, state, input)

        for node in next_nodes:
            # node's begining of life
            pointer = self.wakeup(node, execution, channel, state)

            # async nodes don't return theirs pointers so they are not queued
            if pointer:
                to_queue.append(pointer)

        if execution.proxy.pointers.count() == 0:
            self.finish_execution(execution)

        channel.queue_declare(
            queue=self.config['RABBIT_QUEUE'],
            durable=True
        )

        for pointer in to_queue:
            channel.basic_publish(
                exchange='',
                routing_key=self.config['RABBIT_QUEUE'],
                body=json.dumps({
                    'command': 'step',
                    'pointer_id': pointer.id,
                }),
                properties=pika.BasicProperties(
                    delivery_mode=2,
                ),
            )

    def next(self, xml, node, state, input):
        ''' Given a position in the script, return the next position '''
        if isinstance(node, Exit):
            return []

        elif isinstance(node, Validation) and \
                input[0]['inputs']['items']['response']['value'] == 'reject':
            # find the data backwards
            first_node_found = False
            first_invalid_node = None
            invalidated = set(
                i['ref']
                for i in input[0]['inputs']['items']['inputs']['value']
            )

            for element in iter(xml):
                node = make_node(element)

                more_fields = node.get_invalidated_fields(invalidated, state)

                invalidated.update(more_fields)

                if more_fields and not first_node_found:
                    first_node_found = True
                    first_invalid_node = node

            comment = input[0]['inputs']['items']['comment']['value']

            def get_update_keys(invalidated):
                ikeys = set()
                fkeys = set()
                akeys = set()
                nkeys = set()
                ckeys = set()

                for key in invalidated:
                    node, actor, form, input = key.split('.')
                    index, ref = form.split(':')

                    ikeys.add(('state.items.{node}.actors.items.{actor}.'
                               'forms.{index}.inputs.items.{input}.'
                               'state'.format(
                                    node=node,
                                    actor=actor,
                                    index=index,
                                    input=input,
                                ), 'invalid'))

                    fkeys.add(('state.items.{node}.actors.items.{actor}.'
                               'forms.{index}.state'.format(
                                    node=node,
                                    actor=actor,
                                    index=index,
                                ), 'invalid'))

                    akeys.add(('state.items.{node}.actors.items.{actor}.'
                               'state'.format(
                                    node=node,
                                    actor=actor,
                                ), 'invalid'))

                    nkeys.add(('state.items.{node}.state'.format(
                        node=node,
                    ), 'invalid'))

                for key, _ in nkeys:
                    key = '.'.join(key.split('.')[:-1]) + '.comment'
                    ckeys.add((key, comment))

                return fkeys | akeys | nkeys | ikeys | ckeys

            updates = dict(get_update_keys(invalidated))

            # update state
            collection = self.get_mongo()[
                self.config['MONGO_EXECUTION_COLLECTION']
            ]
            collection.update_one({
                'id': state['id'],
            }, {
                '$set': updates,
            })

            return [first_invalid_node]

        # Return next node by simple adjacency, works for actions and accepted
        # validations

        try:
            # Return next node by simple adjacency
            xmliter = iter(xml)
            xmliter.find(lambda e: e.getAttribute('id') == node.id)

            element = next(xmliter)

            context = None
            while element.tagName == 'if':
                if context is None:
                    context = Condition(state['state'])

                condition = xmliter.get_next_condition()

                if not context.parse(condition):
                    xmliter.expand(element)

                element = next(xmliter)

            return [make_node(element)]
        except StopIteration:
            # End of process
            return []

    def wakeup(self, node, execution, channel, state):
        ''' Waking up a node often means to notify someone or something about
        the execution, this is the first step in node's lifecycle '''

        # create a pointer in this node
        pointer = self.create_pointer(node, execution)
        log.debug('Created pointer p:{} n:{} e:{}'.format(
            pointer.id,
            node.id,
            execution.id,
        ))

        # notify someone
        notified_users = self.notify_users(node, pointer, channel, state)

        # update registry about this pointer
        collection = self.get_mongo()[self.config['MONGO_HISTORY_COLLECTION']]
        collection.insert_one({
            'started_at': datetime.now(),
            'finished_at': None,
            'execution': {
                'id': execution.id,
                'name': execution.name,
                'description': execution.description,
            },
            'node': node.to_json(),
            'notified_users': notified_users,
            'actors': {
                '_type': ':map',
                'items': {},
            },
            'process_id': execution.process_name
        })

        # nodes with forms are not queued
        if not node.is_async():
            return pointer

    def teardown(self, node, pointer, user, input):
        ''' finishes the node's lifecycle '''
        execution = pointer.proxy.execution.get()
        actor_json = {
            '_type': 'actor',
            'state': 'valid',
            'user': user.to_json(include=[
                '_type',
                'fullname',
                'identifier',
            ]),
            'forms': input,
        }

        # update history
        collection = self.get_mongo()[self.config['MONGO_HISTORY_COLLECTION']]
        collection.update_one({
            'execution.id': execution.id,
            'node.id': pointer.node_id,
        }, {
            '$set': {
                'finished_at': datetime.now(),
                'actors.items.{identifier}'.format(
                    identifier=user.identifier,
                ): actor_json,
            },
        })

        # update state
        collection = self.get_mongo()[
            self.config['MONGO_EXECUTION_COLLECTION']
        ]
        collection.update_one({
            'id': execution.id,
        }, {
            '$set': {
                'state.items.{node}.actors.items.{identifier}'.format(
                    node=node.id,
                    identifier=user.identifier,
                ): actor_json,
                'state.items.{node}.state'.format(node=node.id): 'valid',
            },
        })

        log.debug('Deleted pointer p:{} n:{} e:{}'.format(
            pointer.id,
            pointer.node_id,
            execution.id,
        ))

        pointer.delete()

    def finish_execution(self, execution):
        """ shuts down this execution and every related object """
        mongo = self.get_mongo()
        collection = mongo[self.config['MONGO_EXECUTION_COLLECTION']]
        collection.update_one({
            'id': execution.id
        }, {
            '$set': {
                'status': 'finished',
                'finished_at': datetime.now()
            }
        })

        log.debug('Finished e:{}'.format(execution.id))

        execution.delete()

    def notify_users(self, node, pointer, channel, state):
        users = node.get_actors(self.config, state)

        if type(users) != list:
            raise MisconfiguredProvider('Provider returned non list')

        channel.exchange_declare(
            exchange=self.config['RABBIT_NOTIFY_EXCHANGE'],
            exchange_type='direct'
        )

        notified_users = []

        for user in users:
            if not isinstance(user, User):
                raise MisconfiguredProvider(
                    'User returned by hierarchy provider is not User, '
                    'but {}'.format(type(user))
                )

            notified_users.append(user.to_json())

            user.proxy.tasks.add(pointer)

            mediums = self.get_contact_channels(user)

            for medium, params in mediums:
                log.debug('Notified user {} via {} about n:{} e:{}'.format(
                    user.identifier,
                    medium,
                    node.id,
                    pointer.proxy.execution.get().id,
                ))
                channel.basic_publish(
                    exchange=self.config['RABBIT_NOTIFY_EXCHANGE'],
                    routing_key=medium,
                    body=json.dumps({**{
                        'pointer': pointer.to_json(include=['*', 'execution']),
                    }, **params}),
                    properties=pika.BasicProperties(
                        delivery_mode=2,
                    ),
                )

        return notified_users

    def get_mongo(self):
        if self.mongo is None:
            client = MongoClient(self.config['MONGO_URI'])
            db = client[self.config['MONGO_DBNAME']]

            self.mongo = db

        return self.mongo

    def get_contact_channels(self, user: User):
        return [('email', {'email': user.get_x_info('email')})]

    def create_pointer(self, node, execution: Execution):
        ''' Given a node, its process, and a specific execution of the former
        create a persistent pointer to the current execution state '''
        pointer = Pointer(
            node_id=node.id,
            name=node.name,
            description=node.description,
        ).save()

        pointer.proxy.execution.set(execution)

        return pointer

    def recover_step(self, message: dict):
        ''' given an execution id and a pointer from the persistent storage,
        return the asociated process node to continue its execution '''
        pointer = Pointer.get_or_exception(message['pointer_id'])

        return (
            pointer,
            User.get_by('identifier', message.get('user_identifier')),
            message['input'],
        )

    def cancel_execution(self, message):
        execution = Execution.get_or_exception(message['execution_id'])

        for pointer in execution.proxy.pointers.get():
            pointer.delete()

        collection = self.get_mongo()[
            self.config['MONGO_EXECUTION_COLLECTION']
        ]

        collection.update_one({
            'id': execution.id
        }, {
            '$set': {
                'status': 'cancelled',
                'finished_at': datetime.now()
            }
        })

        execution.delete()
