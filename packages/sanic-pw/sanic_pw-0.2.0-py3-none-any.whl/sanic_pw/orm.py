import logging
import os
from importlib import import_module

from peewee import Model as PeeweeModel, Proxy
from cached_property import cached_property
from peewee_migrate.router import Router
from playhouse.db_url import connect

from .models import Model, BaseSignalModel


LOGGER = logging.getLogger(__name__)


class Peewee(object):

    def __init__(self, app=None):
        """
        Initialize the plugin.
        """
        self.app = app
        self.database = Proxy()

        if app is not None:
            self.init_app(app)

    def init_app(self, app, database=None):
        """
        Initialize an application.
        """
        if not app:
            raise RuntimeError('Invalid application.')

        self.app = app
        if not hasattr(app, 'extensions'):
            app.extensions = {}
        app.extensions['peewee'] = self

        app.config.setdefault('PEEWEE_CONNECTION_PARAMS', {})
        app.config.setdefault('PEEWEE_DATABASE_URI', 'sqlite:///peewee.sqlite')
        app.config.setdefault('PEEWEE_MANUAL', False)
        app.config.setdefault('PEEWEE_MIGRATE_DIR', 'migrations')
        app.config.setdefault('PEEWEE_MIGRATE_TABLE', 'migratehistory')
        app.config.setdefault('PEEWEE_MODELS_CLASS', Model)
        app.config.setdefault('PEEWEE_MODELS_IGNORE', [])
        app.config.setdefault('PEEWEE_MODELS_MODULE', '')
        app.config.setdefault('PEEWEE_READ_SLAVES', '')
        app.config.setdefault('PEEWEE_USE_READ_SLAVES', True)

        # Initialize database
        params = app.config['PEEWEE_CONNECTION_PARAMS']
        database = database or app.config.get('PEEWEE_DATABASE_URI')
        if not database:
            raise RuntimeError('Invalid database.')
        database = get_database(database, **params)

        # Configure read slaves
        slaves = app.config['PEEWEE_READ_SLAVES']
        if isinstance(slaves, str):
            slaves = slaves.split(',')
        self.slaves = [get_database(slave, **params) for slave in slaves if slave]

        self.database.initialize(database)
        if self.database.database == ':memory:':
            app.config['PEEWEE_MANUAL'] = True

        # TODO: Replace this code, when a solution will be found
        #if not app.config['PEEWEE_MANUAL']:
        #    app.before_request(self.connect)
        #    app.teardown_request(self.close)

    def connect(self):
        """
        Initialize connection to database.
        """
        LOGGER.info('Connecting [%s]', os.getpid())
        return self.database.connect()

    def close(self, response):
        """
        Close connection to database.
        """
        LOGGER.info('Closing [%s]', os.getpid())
        if not self.database.is_closed():
            self.database.close()
        return response

    @cached_property
    def Model(self):
        """
        Bind model to self database.
        """
        Model_ = self.app.config['PEEWEE_MODELS_CLASS']
        meta_params = {'database': self.database}
        if self.slaves and self.app.config['PEEWEE_USE_READ_SLAVES']:
            meta_params['read_slaves'] = self.slaves

        Meta = type('Meta', (), meta_params)
        return type('Model', (Model_,), {'Meta': Meta})

    @property
    def models(self):
        """
        Return self.application models.py.
        """
        Model_ = self.app.config['PEEWEE_MODELS_CLASS']
        ignore = self.app.config['PEEWEE_MODELS_IGNORE']

        models = []
        if Model_ is not Model:
            try:
                mod = import_module(self.app.config['PEEWEE_MODELS_MODULE'])
                for model in dir(mod):
                    models = getattr(mod, model)
                    if not isinstance(model, PeeweeModel):
                        continue
                    models.append(models)
            except ImportError:
                return models
        elif isinstance(Model_, BaseSignalModel):
            models = BaseSignalModel.models

        return [m for m in models if m._meta.name not in ignore]

    def cmd_create(self, name, auto=False):
        """
        Create a new migration.
        """
        LOGGER.setLevel('INFO')
        LOGGER.propagate = 0

        router = Router(self.database,
                        migrate_dir=self.app.config['PEEWEE_MIGRATE_DIR'],
                        migrate_table=self.app.config['PEEWEE_MIGRATE_TABLE'])

        if auto:
            auto = self.models

        router.create(name, auto=auto)

    def cmd_migrate(self, name=None, fake=False):
        """
        Run migrations.
        """
        LOGGER.setLevel('INFO')
        LOGGER.propagate = 0

        router = Router(self.database,
                        migrate_dir=self.app.config['PEEWEE_MIGRATE_DIR'],
                        migrate_table=self.app.config['PEEWEE_MIGRATE_TABLE'])

        migrations = router.run(name, fake=fake)
        if migrations:
            LOGGER.warn('Migrations are completed: %s' % ', '.join(migrations))

    def cmd_rollback(self, name):
        """
        Rollback migrations.
        """
        LOGGER.setLevel('INFO')
        LOGGER.propagate = 0

        router = Router(self.database,
                        migrate_dir=self.app.config['PEEWEE_MIGRATE_DIR'],
                        migrate_table=self.app.config['PEEWEE_MIGRATE_TABLE'])

        router.rollback(name)

    def cmd_list(self):
        """
        List migrations.
        """
        LOGGER.setLevel('DEBUG')
        LOGGER.propagate = 0

        router = Router(self.database,
                        migrate_dir=self.app.config['PEEWEE_MIGRATE_DIR'],
                        migrate_table=self.app.config['PEEWEE_MIGRATE_TABLE'])

        LOGGER.info('Migrations are done:')
        LOGGER.info('\n'.join(router.done))
        LOGGER.info('')
        LOGGER.info('Migrations are undone:')
        LOGGER.info('\n'.join(router.diff))

    def cmd_merge(self):
        """
        Merge migrations.
        """
        LOGGER.setLevel('DEBUG')
        LOGGER.propagate = 0

        router = Router(self.database,
                        migrate_dir=self.app.config['PEEWEE_MIGRATE_DIR'],
                        migrate_table=self.app.config['PEEWEE_MIGRATE_TABLE'])

        router.merge()

    @cached_property
    def manager(self):
        """
        Integration with the Sanic-Script package.
        """
        from sanic_script import Manager, Command

        manager = Manager(usage="Migrate database.")
        manager.add_command('create', Command(self.cmd_create))
        manager.add_command('migrate', Command(self.cmd_migrate))
        manager.add_command('rollback', Command(self.cmd_rollback))
        manager.add_command('list', Command(self.cmd_list))
        manager.add_command('merge', Command(self.cmd_merge))

        return manager

    @cached_property
    def cli(self):
        import click

        @click.group()
        def cli():
            pass

        @cli.command()
        @click.argument('name')
        @click.option('--auto', is_flag=True)
        def create(name, auto=False):
            """
            Create a new migration.
            """
            return self.cmd_create(name, auto)

        @cli.command()
        @click.argument('name', default=None, required=False)
        @click.option('--fake', is_flag=True)
        def migrate(name, fake=False):
            """
            Run migrations.
            """
            return self.cmd_migrate(name, fake)

        @cli.command()
        @click.argument('name')
        def rollback(name):
            """
            Rollback migrations.
            """
            return self.cmd_rollback(name)

        @cli.command()
        def list():
            """
            List migrations.
            """
            return self.cmd_list()

        return cli


def get_database(obj, **params):
    """
    Get database from given URI/Object.
    """
    if isinstance(obj, str):
        return connect(obj, **params)
    return obj
