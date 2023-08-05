# -*- coding: utf-8 -*-

"""Console script for xml_delta."""
import click

import settings

import logging
logger = logging.getLogger(__name__)

@click.command()
@click.option('-u', 'url', required=True, help='Checks if endpoint is responding.')
@click.option('-t', 'timeout', default=10, help='Timeout for request.')
def check_network(url, timeout):
    from .network_manager import NetworkManager
    if NetworkManager.ping(url=url, timeout=timeout):
        click.echo('Ping success.')
    else:
        click.echo('Ping failed.')

@click.command()
@click.option('-cs', 'connection_string', required=False, help='Checks if db is responding.')
def check_db(connection_string):
    from .storage_manager import StorageManager
    try:
        if not connection_string:
            connection_string = settings.STORAGE_DATABASE_URL
        storage_manager = StorageManager(connection_string)
        storage_manager.session.execute('SELECT 1')
        click.echo("Connection @ '%s' successful." % connection_string)
    except Exception:
        click.echo("Connection @ '%s' failed." % connection_string)

@click.command()
@click.option('-cs', 'connection_string', required=False, help='Checks if db is responding.')
def drop_tables(connection_string):
    from .storage_manager import StorageManager
    try:
        if not connection_string:
            connection_string = settings.STORAGE_DATABASE_URL
        storage_manager = StorageManager(connection_string)
        storage_manager.drop_all()
        click.echo('DropTables finished successfully.')
    except Exception as ex:
        click.echo('DropTables failed with error "%r."', ex.message)

@click.command()
@click.option('-type', 'type', required=True, help='Type to count rows for')
@click.option('-cs', 'connection_string', required=False, help='Connection string for db.')
def count_type(type, connection_string):
    from .storage_manager import StorageManager
    try:
        if not connection_string:
            connection_string = settings.STORAGE_DATABASE_URL
        storage_manager = StorageManager(connection_string)
        click.echo('%d records' % storage_manager.count_type(type))
    except Exception as ex:
        click.echo('Count Type failed with error "%r."', ex.message)

@click.command()
@click.option('-type', 'type', required=True, help='Type of file parsed.')
@click.option('-tag', 'tag', required=True, help='Tag to extract from xml.')
@click.option('-sk', 'sk', required=True, help='Storage key tag to use for identification of records in db level (inserts, updates, deletes). '
                                               'Can be comma separated list of values.')
@click.option('-nk', 'nk', required=False, help='Network key tag to use for network operations.')
@click.option('-file', 'file', type=click.Path(), required=True, help='File to parse.')
@click.option('-endpoint', 'endpoint', default=None, help='Endpoint to post data to.')
def xml_delta(file, type, tag, sk, nk, endpoint):
    if not nk:
        nk = sk
    from delta_task import DeltaTask
    delta_task = DeltaTask(data_type=type, file_path=file, tag=tag, storage_key=sk, network_key=nk, endpoint=endpoint)
    delta_task.run()

@click.command()
def check_file_log():
    logger.info('Check 1.2.')
