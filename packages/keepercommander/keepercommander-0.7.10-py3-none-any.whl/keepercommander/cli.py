#  _  __  
# | |/ /___ ___ _ __  ___ _ _ ®
# | ' </ -_) -_) '_ \/ -_) '_|
# |_|\_\___\___| .__/\___|_|
#              |_|            
#
# Keeper Commander 
# Copyright 2017 Keeper Security Inc.
# Contact: ops@keepersecurity.com
#
import sys
import getpass
import json
import click
import datetime
import time

from keepercommander.record import Record
from keepercommander import display, api, imp_exp
from keepercommander.params import KeeperParams
from keepercommander.error import AuthenticationError, CommunicationError

####### shell
@click.command(help = 'Use Keeper interactive shell')
@click.pass_obj
def shell(params):
    loop(params)

####### list
@click.command(help = 'Display all record UID/titles')
@click.pass_obj
def list(params):
    try:
        prompt_for_credentials(params)
        api.sync_down(params)
        if (len(params.record_cache) > 0):
            results = api.search_records(params, '')
            display.formatted_records(results)

        if (len(params.shared_folder_cache) > 0):
            results = api.search_shared_folders(params, '')
            display.formatted_shared_folders(results)

        if (len(params.team_cache) > 0):
            results = api.search_teams(params, '')
            display.formatted_teams(results)

    except Exception as e:
        raise click.ClickException(e)

####### list_sf
@click.command(help = 'Display all Shared Folder UID/titles')
@click.pass_obj
def list_sf(params):
    try:
        prompt_for_credentials(params)
        api.sync_down(params)

        if (len(params.shared_folder_cache) > 0):
            results = api.search_shared_folders(params, '')
            display.formatted_shared_folders(results)

    except Exception as e:
        raise click.ClickException(e)

####### list_teams
@click.command(help = 'Display all Teams')
@click.pass_obj
def list_teams(params):
    try:
        prompt_for_credentials(params)
        api.sync_down(params)

        if (len(params.team_cache) > 0):
            results = api.search_teams(params, '')
            display.formatted_teams(results)

    except Exception as e:
        raise click.ClickException(e)

####### get
@click.command('get', help = 'Display specified Keeper record')
@click.pass_obj
@click.argument('uid')
def get(params, uid):
    try:
        prompt_for_credentials(params)
        api.sync_down(params)
        if uid:
            api.get_record(params,uid).display()
    except Exception as e:
        raise click.ClickException(e)

####### search
@click.command(help = 'Search vault with a regular expression')
@click.argument('regex')
@click.pass_obj
def search(params, regex):
    try:
        prompt_for_credentials(params)
        api.sync_down(params)
        if (len(params.record_cache) == 0): 
            print('No records')
            return
        results = api.search_records(params, regex) 
        display.formatted_records(results)
    except Exception as e:
        raise click.ClickException(e)

####### rotate
@click.command(help = 'Rotate Keeper record')
@click.pass_obj
@click.argument('uid')
@click.option('--match', help='regular expression to select records for password rotation')
@click.option('--print', flag_value=True, help='display the record content after rotation')
def rotate(params, uid, match, print):
    try:
        prompt_for_credentials(params)
        api.sync_down(params)
        if uid:
            api.rotate_password(params, uid)
            if print:
                display.print_record(params, uid)
        else:
            if filter:
                results = api.search_records(params, match)
                for r in results:
                    api.rotate_password(params, r.record_uid)
                    if print:
                        display.print_record(params, uid)
    except Exception as e:
        raise click.ClickException(e)

####### import
@click.command('import', help='Import password records from local file')
@click.pass_obj
@click.option('--format', type=click.Choice(['tab-separated', 'json']))
@click.argument('filename')
def _import(params, format, filename):
    try:
        prompt_for_credentials(params)
        imp_exp._import(params, format, filename)
    except Exception as e:
        raise click.ClickException(e)

####### create_sf
@click.command('create_sf', help='Create shared folders from JSON input file')
@click.pass_obj
@click.argument('filename')
def create_sf(params, filename):
    try:
        prompt_for_credentials(params)
        imp_exp.create_sf(params, filename)
    except Exception as e:
        raise click.ClickException(e)

####### export
@click.command(help='Export password records from Keeper')
@click.pass_obj
@click.option('--format', type=click.Choice(['tab-separated', 'json']))
@click.argument('filename')
def export(params, format, filename):
    try:
        prompt_for_credentials(params)
        imp_exp.export(params, format, filename)
    except Exception as e:
        raise click.ClickException(e)

####### delete-all
@click.command('delete-all', help='Delete all Keeper records on server')
@click.confirmation_option(prompt='Are you sure you want to delete all Keeper records on the server?')
@click.pass_obj
def delete_all(params):
    try:
        prompt_for_credentials(params)
        imp_exp.delete_all(params)
    except Exception as e:
        raise click.ClickException(e)

stack = []

def goodbye():
    print('\nGoodbye.\n')
    sys.exit()

def get_params_from_config(config_filename):
    params = KeeperParams()
    params.config_filename = 'config.json'
    if config_filename:
        params.config_filename = config_filename

    try:
        with open(params.config_filename) as config_file:

            try:
                params.config = json.load(config_file)

                if 'user' in params.config:
                    params.user = params.config['user']

                if 'server' in params.config:
                    params.server = params.config['server']

                if 'password' in params.config:
                    params.password = params.config['password']

                if 'challenge' in params.config:
                    try:
                        import keepercommander.yubikey.yubikey
                        challenge = params.config['challenge']
                        params.password = keepercommander.yubikey.yubikey.get_response(challenge)
                    except Exception as e:
                        print(e)
                        sys.exit(1)

                if 'timedelay' in params.config:
                    params.timedelay = params.config['timedelay']

                if 'mfa_token' in params.config:
                    params.mfa_token = params.config['mfa_token']

                if 'mfa_type' in params.config:
                    params.mfa_type = params.config['mfa_type']

                if 'commands' in params.config:
                    params.commands = params.config['commands']

                if 'plugins' in params.config:
                    params.plugins = params.config['plugins']

                if 'debug' in params.config:
                    params.debug = params.config['debug']

            except:
                print('Error: Unable to parse JSON file ' + params.config_filename)
                raise

    except IOError:
        if config_filename:
            print('Error: Unable to open config file ' + config_filename)
        pass

    if not params.server:
        params.server = 'https://keepersecurity.com/api/v2/'

    return params


def do_command(params):
    if (params.command == 'q'):
        return False

    elif (params.command == 'l'):
        if (len(params.record_cache) == 0):
            print('No records')
        else:
            results = api.search_records(params, '')
            display.formatted_records(results)

    elif (params.command == 'lsf'):
        if (len(params.shared_folder_cache) == 0): 
            print('No shared folders')
        else:
            results = api.search_shared_folders(params, '') 
            display.formatted_shared_folders(results)

    elif (params.command == 'lt'):
        if (len(params.team_cache) == 0): 
            print('No teams')
        else:
            results = api.search_teams(params, '') 
            display.formatted_teams(results)

    elif (params.command[:2] == 'g '):
        if (api.is_shared_folder(params, params.command[2:])):
            sf = api.get_shared_folder(params, params.command[2:])
            if sf:
                sf.display()
        elif (api.is_team(params, params.command[2:])):
            team = api.get_team(params, params.command[2:])
            if team:
                team.display()
        else:
            r = api.get_record(params, params.command[2:])
            if r:
                r.display()

    elif (params.command[:2] == 'r '):
        api.rotate_password(params, params.command[2:])

    elif (params.command[:2] == 'd '):
        api.delete_record(params, params.command[2:])

    elif (params.command == 'c'):
        print(chr(27) + "[2J")

    elif (params.command[:2] == 's '):
        results = api.search_records(params, params.command[2:])
        display.formatted_records(results)

    elif (params.command[:2] == 'b '):
        results = api.search_records(params, params.command[2:])
        for r in results:
            api.rotate_password(params, r.record_uid)

    elif (params.command[:3] == 'an '):
        api.append_notes(params, params.command[3:])

    elif (params.command == 'd'):
        api.sync_down(params)

    elif (params.command == 'a'):
        record = Record()
        while not record.title:
            record.title = input("... Title (req'd): ")
        record.folder = input("... Folder: ")
        record.login = input("... Login: ")
        record.password = input("... Password: ")
        record.login_url = input("... Login URL: ")
        record.notes = input("... Notes: ")
        while True:
            custom_dict = {} 
            custom_dict['name'] = input("... Custom Field Name : ") 
            if not custom_dict['name']:
                break

            custom_dict['value'] = input("... Custom Field Value : ") 
            custom_dict['type'] = 'text' 
            record.custom_fields.append(custom_dict)

        api.add_record(params, record)

    elif (params.command == 'h'):
        display.formatted_history(stack)

    elif (params.command == 'debug'):
        if params.debug:
            params.debug = False
            print('Debug OFF')
        else:
            params.debug = True
            print('Debug ON')

    elif params.command == '':
        pass

    else:
        print('\n\nShell Commands:\n')
        print('  d         ... download & decrypt data')
        print('  l         ... list folders and record titles')
        print('  lsf       ... list shared folders')
        print('  lt        ... list teams')
        print('  s <regex> ... search with regular expression')
        print('  g <uid>   ... get record or shared folder details for uid')
        print('  r <uid>   ... rotate password for uid')
        print('  b <regex> ... rotate password for matches of regular expression')
        print('  a         ... add a new record interactively')
        print('  an <uid>  ... append some notes to the specified record')
        print('  c         ... clear the screen')
        print('  h         ... show command history')
        print('  q         ... quit')
        print('')

    if params.command:
        if params.command != 'h':
            stack.append(params.command)
            stack.reverse()

    return True

def runcommands(params):
    keep_running = True
    timedelay = params.timedelay

    while keep_running:
        for c in params.commands:
            params.command = c
            print('Executing [' + params.command + ']...')
            try:
                if not do_command(params):
                    print('Command ' + params.command + ' failed.')
            except CommunicationError as e:
                print("Communication Error:" + str(e.message))
            except AuthenticationError as e:
                print("AuthenticationError Error: " + str(e.message))
            except:
                print('An unexpected error occurred: ' + str(sys.exc_info()[0]))

            params.command = ''

        if (timedelay == 0):
            keep_running = False
        else:
            print(datetime.datetime.now().strftime('%Y/%m/%d %H:%M:%S') + \
                ' Waiting for ' + str(timedelay) + ' seconds')
            time.sleep(timedelay)


def prompt_for_credentials(params):
        while not params.user:
            params.user = getpass.getpass(prompt='User(Email): ', stream=None)

        while not params.password:
            params.password = getpass.getpass(prompt='Password: ', stream=None)


def loop(params):

    display.welcome()

    try:

        prompt_for_credentials(params)

        # if commands are provided, execute those then exit
        if params.commands:
            runcommands(params)
            goodbye()

        if params.debug: print('Params: ' + str(params))

        # start with a sync download
        if not params.command:
            params.command = 'd'

        # go into interactive mode
        while True:
            if not params.command:
                try:
                    params.command = input("Keeper > ")
                except KeyboardInterrupt:
                    print('')
                except EOFError:
                    raise KeyboardInterrupt

            try:
                if not do_command(params):
                    raise KeyboardInterrupt
            except CommunicationError as e:
                print("Communication Error:" + str(e.message))
            except AuthenticationError as e:
                print("AuthenticationError Error: " + str(e.message))
            except KeyboardInterrupt as e:
                raise
            except:
                print('An unexpected error occurred: ' + str(sys.exc_info()[0]))
                raise

            params.command = ''

    except KeyboardInterrupt:
        goodbye()
