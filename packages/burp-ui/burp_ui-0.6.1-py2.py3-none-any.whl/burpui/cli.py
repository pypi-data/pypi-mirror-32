# -*- coding: utf8 -*-
"""
Burp-UI is a web-ui for burp backup written in python with Flask and
jQuery/Bootstrap

.. module:: burpui.cli
    :platform: Unix
    :synopsis: Burp-UI CLI module.

.. moduleauthor:: Ziirish <hi+burpui@ziirish.me>
"""
import os
import sys
import click

if os.getenv('BUI_MODE') in ['server', 'ws'] or 'websocket' in sys.argv:
    try:
        from gevent import monkey
        monkey.patch_socket()
    except ImportError:
        pass

from .app import create_app  # noqa
from .exceptions import BUIserverException  # noqa
from six import iteritems  # noqa

try:
    from flask_socketio import SocketIO  # noqa
    WS_AVAILABLE = True
except ImportError:
    WS_AVAILABLE = False

ROOT = os.path.dirname(os.path.realpath(__file__))
DEBUG = os.getenv('BUI_DEBUG') or os.getenv('FLASK_DEBUG') or False
DEBUG = DEBUG and DEBUG.lower() not in ['false', 'no', '0']

VERBOSE = os.getenv('BUI_VERBOSE') or 0
if VERBOSE:
    try:
        VERBOSE = int(VERBOSE)
    except ValueError:
        VERBOSE = 0

# UNITTEST is used to skip the burp-2 requirements for modes != server
UNITTEST = os.getenv('BUI_MODE') not in ['server', 'manage', 'celery', 'legacy', 'ws']
CLI = os.getenv('BUI_MODE') not in ['server', 'legacy']

try:
    app = create_app(
        conf=os.environ.get('BUI_CONFIG'),
        verbose=VERBOSE,
        logfile=os.environ.get('BUI_LOGFILE'),
        debug=DEBUG,
        gunicorn=False,
        unittest=UNITTEST,
        cli=CLI,
        websocket_server=(os.getenv('BUI_MODE') == 'ws' or 'websocket' in sys.argv)
    )
except:
    import traceback
    traceback.print_exc()
    sys.exit(1)

try:
    from .app import create_db
    from .ext.sql import db
    from flask_migrate import Migrate

    # This may have been reseted by create_app
    if isinstance(app.database, bool):
        app.config['WITH_SQL'] = app.database
    else:
        app.config['WITH_SQL'] = app.database and \
            app.database.lower() != 'none'

    if app.config['WITH_SQL']:
        create_db(app, True)

        mig_dir = os.getenv('BUI_MIGRATIONS')
        if mig_dir:
            migrate = Migrate(app, db, mig_dir)
        else:
            migrate = Migrate(app, db)
except ImportError:
    pass


def _die(error, appli=None):
    appli = " '{}'".format(appli) if appli else ''
    click.echo(
        click.style(
            'Unable to initialize the application{}: {}'.format(appli, error),
            fg='red'
        ),
        err=True
    )
    sys.exit(2)


@app.cli.command()
def legacy():
    """Legacy server for backward compatibility."""
    click.echo(
        click.style(
            'If you want to pass options, you should run \'python -m burpui '
            '-m legacy [...]\' instead',
            fg='yellow'
        )
    )
    app.manual_run()


@app.cli.command()
@click.option('-b', '--bind', default='127.0.0.1',
              help='Which address to bind to for the websocket server')
@click.option('-p', '--port', default=5001,
              help='Which port to listen on for the websocket server')
@click.option('-d', '--debug', default=False, is_flag=True,
              help='Whether to start the websocket server in debug mode')
def websocket(bind, port, debug):
    """Start a new websocket server."""
    try:
        from .ext.ws import socketio
    except ImportError:
        _die('Missing requirement, did you ran \'pip install'
             ' "burp-ui[websocket]"\'?', 'websocket')
    socketio.run(app, host=bind, port=port, debug=debug)


@app.cli.command()
@click.option('-b', '--backend', default='BASIC',
              help='User Backend (default is BASIC).')
@click.option('-p', '--password', help='Password to assign to user.',
              default=None)
@click.option('-a', '--ask', default=False, is_flag=True,
              help='If no password is provided and this flag is enabled, '
                   'you\'ll be prompted for one, else a random one will be '
                   'generated.')
@click.option('-v', '--verbose', default=False, is_flag=True,
              help='Add extra debug messages.')
@click.argument('name')
def create_user(backend, password, ask, verbose, name):
    """Create a new user."""
    try:
        msg = app.load_modules(True)
    except Exception as e:
        msg = str(e)

    if msg:
        _die(msg, 'create_user')

    click.echo(click.style('[*] Adding \'{}\' user...'.format(name), fg='blue'))
    try:
        handler = getattr(app, 'uhandler')
    except AttributeError:
        handler = None

    if not handler or len(handler.backends) == 0 or \
            backend not in handler.backends:
        click.echo(click.style('[!] No authentication backend found', fg='red'))
        sys.exit(1)

    back = handler.backends[backend]

    if back.add_user is False:
        click.echo(click.style("[!] The '{}' backend does not support user "
                               "creation".format(backend), fg='red'))
        sys.exit(2)

    if not password:
        if ask:
            import getpass
            password = getpass.getpass()
            confirm = getpass.getpass('Confirm: ')
            if password != confirm:
                click.echo(click.style("[!] Passwords mismatch", fg='red'))
                sys.exit(3)
        else:
            import random

            alphabet = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLM" \
                       "NOPQRSTUVWXYZ"
            pw_length = 8
            mypw = ""

            for i in range(pw_length):
                next_index = random.randrange(len(alphabet))
                mypw += alphabet[next_index]
            password = mypw
            click.echo(
                click.style(
                    '[+] Generated password: {}'.format(password),
                    fg='blue'
                )
            )

    success, message, _ = back.add_user(name, password)
    click.echo(click.style(
        '[+] Success: {}{}'.format(
            success, ' -> {}'.format(message) if verbose and message else ''
        ),
        fg='green' if success else 'red')
    )


@app.cli.command()
@click.option('-p', '--password', help='Password to assign to user.',
              default=None)
@click.option('-u', '--username', help='Provide the username to get the full '
              'configuration line.',
              default=None)
@click.option('-b', '--batch', default=False, is_flag=True,
              help='Don\'t be extra verbose so that you can use the output '
              'directly in your scripts. Requires both -u and -p.')
def hash_password(password, username, batch):
    """Hash a given password to fill the configuration file."""
    from werkzeug.security import generate_password_hash

    if batch and (not username or not password):
        click.echo(click.style(
            'You need to provide both a username and a password using the '
            '-u and -p flags!',
            fg='red')
        )
        sys.exit(1)

    askpass = False
    if not password:
        askpass = True
        import getpass
        password = getpass.getpass()

    hashed = generate_password_hash(password)
    if not batch:
        click.echo("'{}' hashed into: {}".format(
            password if not askpass else '*' * 8,
            hashed)
        )
    if username:
        if not batch:
            click.echo(click.style(
                '#8<{}'.format('-' * 77),
                fg='blue')
            )
        click.echo('{} = {}'.format(username, hashed))
        if not batch:
            click.echo(click.style(
                '#8<{}'.format('-' * 77),
                fg='blue')
            )


@app.cli.command()
@click.argument('language')
def init_translation(language):
    """Initialize a new translation for the given language."""
    try:
        import babel  # noqa
    except ImportError:
        click.echo(
            click.style('Missing i18n requirements, giving up', fg='yellow')
        )
        return
    os.chdir(os.path.join(ROOT, '..'))
    os.system('pybabel extract -F babel.cfg -k __ -k lazy_gettext -o messages.pot burpui')
    os.system('pybabel init -i messages.pot -d burpui/translations -l {}'.format(language))
    os.unlink('messages.pot')


@app.cli.command()
def update_translation():
    """Update translation files."""
    try:
        import babel  # noqa
    except ImportError:
        click.echo(
            click.style('Missing i18n requirements, giving up', fg='yellow')
        )
        return
    os.chdir(os.path.join(ROOT, '..'))
    os.system('pybabel extract -F babel.cfg -k __ -k lazy_gettext -o messages.pot burpui')
    os.system('pybabel update -i messages.pot -d burpui/translations')
    os.unlink('messages.pot')


@app.cli.command()
def compile_translation():
    """Compile translations."""
    try:
        import babel  # noqa
    except ImportError:
        click.echo(
            click.style('Missing i18n requirements, giving up', fg='yellow')
        )
        return
    os.chdir(os.path.join(ROOT, '..'))
    os.system('pybabel compile -f -d burpui/translations')


@app.cli.command()
@click.option('-b', '--burp-conf-cli', 'bconfcli', default=None,
              help='Burp client configuration file')
@click.option('-s', '--burp-conf-serv', 'bconfsrv', default=None,
              help='Burp server configuration file')
@click.option('-c', '--client', default='bui',
              help='Name of the burp client that will be used by Burp-UI '
                   '(defaults to "bui")')
@click.option('-h', '--host', default='::1',
              help='Address of the status server (defaults to "::1")')
@click.option('-r', '--redis', default=None,
              help='Redis URL to connect to')
@click.option('-d', '--database', default=None,
              help='Database to connect to for persistent storage')
@click.option('-p', '--plugins', default=None,
              help='Plugins location')
@click.option('-n', '--dry', is_flag=True,
              help='Dry mode. Do not edit the files but display changes')
def setup_burp(bconfcli, bconfsrv, client, host, redis, database, plugins, dry):
    """Setup burp client for burp-ui."""
    if app.vers != 2:
        click.echo(
            click.style(
                'Sorry, you can only setup the Burp 2 client',
                fg='red'
            ),
            err=True
        )
        sys.exit(1)

    if not app.standalone:
        click.echo(
            click.style(
                'Sorry, only the standalone mode is supported',
                fg='red'
            ),
            err=True
        )
        sys.exit(1)

    try:
        msg = app.load_modules(True)
    except Exception as e:
        msg = str(e)

    if msg:
        _die(msg, 'setup_burp')

    from .misc.parser.utils import Config
    from .app import get_redis_server
    import difflib
    import tempfile

    parser = app.client.get_parser()
    orig = source = None
    conf_orig = []
    if dry:
        try:
            with open(app.conf.options.filename) as fil:
                conf_orig = fil.readlines()
        except:
            pass

        orig = source = app.conf.options.filename
        (_, temp) = tempfile.mkstemp()
        app.conf.options.filename = temp

    # handle migration of old config files
    if app.conf.section_exists('Burp2'):
        if app.conf.rename_section('Burp2', 'Burp', source):
            click.echo(
                click.style(
                    'Renaming old [Burp2] section',
                    fg='blue'
                )
            )
            app.conf._refresh(True)

    refresh = False
    if not app.conf.lookup_section('Burp', source):
        refresh = True
    if not app.conf.lookup_section('Global', source):
        refresh = True
    if (database or redis) and not app.conf.lookup_section('Production', source):
        refresh = True

    if refresh:
        app.conf._refresh(True)

    def _edit_conf(key, val, attr, section='Burp', obj=app.client):
        if val and (((key not in app.conf.options[section]) or
                    (key in app.conf.options[section] and
                    val != app.conf.options[section][key])) and
                    getattr(obj, attr) != val):
            app.conf.options[section][key] = val
            app.conf.options.write()
            click.echo(
                click.style(
                    'Adding new option: "{}={}" to section [{}]'.format(
                        key,
                        val,
                        section
                    ),
                    fg='blue'
                )
            )
            return True
        return False

    def _color_diff(line):
        if line.startswith('+'):
            return click.style(line, fg='green')
        elif line.startswith('-'):
            return click.style(line, fg='red')
        elif line.startswith('^'):
            return click.style(line, fg='blue')
        return line

    refresh = False
    refresh |= _edit_conf('bconfcli', bconfcli, 'burpconfcli')
    refresh |= _edit_conf('bconfsrv', bconfsrv, 'burpconfsrv')
    refresh |= _edit_conf('plugins', plugins, 'plugins', 'Global', app)

    if refresh:
        app.conf._refresh(True)

    if redis:
        try:
            # detect missing modules
            import redis as redis_client  # noqa
            import celery  # noqa
            import socket
            if ('redis' not in app.conf.options['Production'] or
                'redis' in app.conf.options['Production'] and
                app.conf.options['Production']['redis'] != redis) and \
                    app.redis != redis:
                app.conf.options['Production']['redis'] = redis

            rhost, rport, _ = get_redis_server(app)
            ret = -1
            for res in socket.getaddrinfo(rhost, rport, socket.AF_UNSPEC, socket.SOCK_STREAM):
                if ret == 0:
                    break
                af, socktype, proto, _, sa = res
                try:
                    s = socket.socket(af, socktype, proto)
                except socket.error:
                    continue
                try:
                    ret = s.connect_ex(sa)
                except:
                    continue

            if ret == 0:
                app.conf.options['Production']['celery'] = 'true'

                app.conf.options['Production']['storage'] = 'redis'

                app.conf.options['Production']['cache'] = 'redis'
            else:
                click.echo(
                    click.style(
                        'Unable to contact the redis server, disabling it',
                        fg='yellow'
                    )
                )
                app.conf.options['Production']['storage'] = 'default'
                app.conf.options['Production']['cache'] = 'default'
                if app.use_celery:
                    app.conf.options['Production']['celery'] = 'false'

            app.conf.options.write()
            app.conf._refresh(True)
        except ImportError:
            click.echo(
                click.style(
                    'Unable to activate redis & celery. Did you ran the '
                    '\'pip install burp-ui[celery]\' and '
                    '\'pip install burp-ui[gunicorn-extra]\' commands first?',
                    fg='yellow'
                )
            )

    if database:
        try:
            from .ext.sql import db  # noqa
            if ('database' not in app.conf.options['Production'] or
                'database' in app.conf.options['Production'] and
                app.conf.options['Production']['database'] != database) and \
                    app.database != database:
                app.conf.options['Production']['database'] = database
                app.conf.options.write()
                app.conf._refresh(True)
        except ImportError:
            click.echo(
                click.style(
                    'It looks like some dependencies are missing. Did you ran '
                    'the \'pip install "burp-ui[sql]"\' command first?',
                    fg='yellow'
                )
            )

    if dry:
        temp = app.conf.options.filename
        app.conf.options.filename = orig
        after = []
        try:
            if not os.path.exists(temp) or os.path.getsize(temp) == 0:
                after = conf_orig
            else:
                with open(temp) as fil:
                    after = fil.readlines()
                os.unlink(temp)
        except:
            pass
        diff = difflib.unified_diff(conf_orig, after, fromfile=orig, tofile='{}.new'.format(orig))
        out = ''
        for line in diff:
            out += _color_diff(line)
        if out:
            click.echo_via_pager(out)

    bconfcli = bconfcli or app.conf.options['Burp'].get('bconfcli') or \
        getattr(app.client, 'burpconfcli')
    bconfsrv = bconfsrv or app.conf.options['Burp'].get('bconfsrv') or \
        getattr(app.client, 'burpconfsrv')
    dest_bconfcli = bconfcli

    if not os.path.exists(bconfsrv):
        click.echo(
            click.style(
                'Unable to locate burp-server configuration, aborting!',
                fg='red'
            ),
            err=True
        )
        sys.exit(1)

    confsrv = Config(bconfsrv, parser, 'srv')
    confsrv.set_default(bconfsrv)
    confsrv.parse()

    if host not in ['::1', '127.0.0.1']:
        bind = confsrv.get('status_address')
        if (bind and bind not in [host, '::', '0.0.0.0']) or not bind:
            click.echo(
                click.style(
                    'It looks like your burp server is not exposing it\'s '
                    'status port in a way that is reachable by Burp-UI!',
                    fg='yellow'
                )
            )
            click.echo(
                click.style(
                    'You may want to set the \'status_address\' setting with '
                    'either \'{}\', \'::\' or \'0.0.0.0\' in the {} file '
                    'in order to make Burp-UI work'.format(host, bconfsrv),
                    fg='blue'
                )
            )

    status_port = confsrv.get('status_port', [4972])
    if 'max_status_children' not in confsrv:
        click.echo(
            click.style(
                'We need to set the number of \'max_status_children\'. '
                'Setting it to 15.',
                fg='blue'
            )
        )
        confsrv['max_status_children'] = 15
        status_port = status_port[0]
    else:
        max_status_children = confsrv.get('max_status_children')
        found = False
        for idx, value in enumerate(max_status_children):
            if value >= 15:
                found = True
                if idx >= len(status_port):
                    status_port = status_port[-1]
                else:
                    status_port = status_port[idx]
                break
        if not found:
            click.echo(
                click.style(
                    'We need to raise the number of \'max_status_children\'. '
                    'Raising it to 15 instead of {}.'.format(max_status_children),
                    fg='yellow'
                )
            )
            confsrv['max_status_children'][-1] = 15
            status_port = status_port[-1]

    if 'restore_client' not in confsrv:
        confsrv['restore_client'] = client
    else:
        restore = confsrv.getlist('restore_client')
        if client not in restore:
            confsrv['restore_client'].append(client)

    confsrv['monitor_browse_cache'] = True

    ca_client_dir = confsrv.get('ca_csr_dir')
    if ca_client_dir and not os.path.exists(ca_client_dir):
        try:
            os.makedirs(ca_client_dir)
        except IOError as exp:
            click.echo(
                click.style(
                    'Unable to create "{}" dir: {}'.format(ca_client_dir, exp),
                    fg='yellow'
                ),
                err=True
            )

    if confsrv.dirty:
        if dry:
            (_, dstfile) = tempfile.mkstemp()
        else:
            dstfile = bconfsrv

        confsrv.store(conf=bconfsrv, dest=dstfile, insecure=True)
        if dry:
            before = []
            after = []
            try:
                with open(bconfsrv) as fil:
                    before = fil.readlines()
            except:
                pass
            try:
                with open(dstfile) as fil:
                    after = fil.readlines()
                os.unlink(dstfile)
            except:
                pass
            diff = difflib.unified_diff(before, after, fromfile=bconfsrv, tofile='{}.new'.format(bconfsrv))
            out = ''
            for line in diff:
                out += _color_diff(line)
            if out:
                click.echo_via_pager(out)

    if confsrv.get('clientconfdir'):
        bconfagent = os.path.join(confsrv.get('clientconfdir'), client)
    else:
        click.echo(
            click.style(
                'Unable to find "clientconfdir" option, you will have to '
                'setup the agent by your own',
                fg='yellow'
            )
        )
        bconfagent = os.devnull

    if not os.path.exists(bconfcli):
        clitpl = """
mode = client
port = 4971
status_port = 4972
server = ::1
password = abcdefgh
cname = {0}
protocol = 1
pidfile = /tmp/burp.client.pid
syslog = 0
stdout = 1
progress_counter = 1
network_timeout = 72000
server_can_restore = 0
cross_all_filesystems=0
ca_burp_ca = /usr/sbin/burp_ca
ca_csr_dir = /etc/burp/CA-client
ssl_cert_ca = /etc/burp/ssl_cert_ca-client-{0}.pem
ssl_cert = /etc/burp/ssl_cert-bui-client.pem
ssl_key = /etc/burp/ssl_cert-bui-client.key
ssl_key_password = password
ssl_peer_cn = burpserver
include = /home
exclude_fs = sysfs
exclude_fs = tmpfs
nobackup = .nobackup
exclude_comp=bz2
exclude_comp=gz
""".format(client)

        if dry:
            (_, dest_bconfcli) = tempfile.mkstemp()
        with open(dest_bconfcli, 'w') as confcli:
            confcli.write(clitpl)

    parser = app.client.get_parser()

    confcli = Config(dest_bconfcli, parser, 'srv')
    confcli.set_default(dest_bconfcli)
    confcli.parse()

    if confcli.get('cname') != client:
        confcli['cname'] = client
    if confcli.get('server') != host:
        confcli['server'] = host
    if confcli.get('status_port')[0] != status_port:
        c_status_port = confcli.get_raw('status_port')
        c_status_port[0] = status_port

    if confcli.dirty:
        if dry:
            (_, dstfile) = tempfile.mkstemp()
        else:
            dstfile = bconfcli

        confcli.store(conf=bconfcli, dest=dstfile, insecure=True)
        if dry:
            before = []
            after = []
            try:
                with open(bconfcli) as fil:
                    before = fil.readlines()
            except:
                pass
            try:
                with open(dstfile) as fil:
                    after = fil.readlines()
                os.unlink(dstfile)
            except:
                pass

            if dest_bconfcli != bconfcli:
                # the file did not exist
                os.unlink(dest_bconfcli)
                before = []

            diff = difflib.unified_diff(before, after, fromfile=bconfcli, tofile='{}.new'.format(bconfcli))
            out = ''
            for line in diff:
                out += _color_diff(line)
            if out:
                click.echo_via_pager(out)

    if not os.path.exists(bconfagent):

        agenttpl = """
password = abcdefgh
"""

        if not dry:
            with open(bconfagent, 'w') as confagent:
                confagent.write(agenttpl)
        else:
            before = []
            after = ['{}\n'.format(x) for x in agenttpl.splitlines()]
            diff = difflib.unified_diff(before, after, fromfile='None', tofile=bconfagent)
            out = ''
            for line in diff:
                out += _color_diff(line)
            if out:
                click.echo_via_pager(out)

    else:
        confagent = Config(bconfagent, parser, 'cli')
        confagent.set_default(bconfagent)
        confagent.parse()

        if confagent.get('password') != confcli.get('password'):
            click.echo(
                click.style(
                    'It looks like the passwords in the {} and the {} files '
                    'mismatch. Burp-UI will not work properly until you fix '
                    'this'.format(bconfcli, bconfagent),
                    fg='yellow'
                )
            )


@app.cli.command()
@click.option('-c', '--client', default='bui',
              help='Name of the burp client that will be used by Burp-UI '
                   '(defaults to "bui")')
@click.option('-h', '--host', default='::1',
              help='Address of the status server (defaults to "::1")')
@click.option('-t', '--tips', is_flag=True,
              help='Show you some tips')
def diag(client, host, tips):
    """Check Burp-UI is correctly setup."""
    if app.vers != 2:
        click.echo(
            click.style(
                'Sorry, you can only setup the Burp 2 client',
                fg='red'
            ),
            err=True
        )
        sys.exit(1)

    if not app.standalone:
        click.echo(
            click.style(
                'Sorry, only the standalone mode is supported',
                fg='red'
            ),
            err=True
        )
        sys.exit(1)

    try:
        msg = app.load_modules(True)
    except Exception as e:
        msg = str(e)

    if msg:
        _die(msg, 'diag')

    from .misc.parser.utils import Config
    from .app import get_redis_server

    if 'Production' in app.conf.options and \
            'redis' in app.conf.options['Production']:
        try:
            # detect missing modules
            import redis as redis_client  # noqa
            import celery  # noqa
            import socket

            rhost, rport, _ = get_redis_server(app)
            ret = -1
            for res in socket.getaddrinfo(rhost, rport, socket.AF_UNSPEC, socket.SOCK_STREAM):
                if ret == 0:
                    break
                af, socktype, proto, _, sa = res
                try:
                    s = socket.socket(af, socktype, proto)
                except socket.error:
                    continue
                try:
                    ret = s.connect_ex(sa)
                except:
                    continue

            if ret != 0:
                click.echo(
                    click.style(
                        'Unable to contact the redis server, disabling it',
                        fg='yellow'
                    )
                )

        except ImportError:
            click.echo(
                click.style(
                    'Unable to activate redis & celery. Did you ran the '
                    '\'pip install "burp-ui[celery]"\' and '
                    '\'pip install "burp-ui[gunicorn-extra]"\' commands first?',
                    fg='yellow'
                )
            )

    if 'Production' in app.conf.options and \
            'database' in app.conf.options['Production']:
        try:
            from .ext.sql import db  # noqa
        except ImportError:
            click.echo(
                click.style(
                    'It looks like some dependencies are missing. Did you ran '
                    'the \'pip install "burp-ui[sql]"\' command first?',
                    fg='yellow'
                )
            )

    section = 'Burp'
    if not app.conf.section_exists(section):
        click.echo(
            click.style(
                'Section [Burp] not found, looking for the old [Burp2] section '
                'instead.',
                fg='yellow'
            )
        )
        section = 'Burp2'
        if not app.conf.section_exists(section):
            click.echo(
                click.style(
                    'No [Burp*] section found at all!',
                    fg='red'
                )
            )
            section = 'Burp'

    bconfcli = app.conf.options.get(section, {}).get('bconfcli') or \
        getattr(app.client, 'burpconfcli')
    bconfsrv = app.conf.options.get(section, {}).get('bconfsrv') or \
        getattr(app.client, 'burpconfsrv')

    try:
        app.client.status()
    except Exception as e:
        if 'Unable to spawn burp process' in str(e):
            try:
                app.client._spawn_burp(verbose=True)
            except Exception as e:
                msg = str(e)
        else:
            msg = str(e)

    if msg:
        click.echo(click.style(msg, fg='red'))
        if 'could not connect' in msg:
            click.echo(
                click.style(
                    'It looks like your burp-client can not reach your '
                    'burp-server. Please check both your \'server\' setting in '
                    'your \'{}\' file and \'status_address\' in your \'{}\' '
                    'file.'.format(bconfcli, bconfsrv),
                    fg='yellow'
                )
            )

    errors = False
    if os.path.exists(bconfcli):
        parser = app.client.get_parser()

        confcli = Config(bconfcli, parser, 'srv')
        confcli.set_default(bconfcli)
        confcli.parse()

        if confcli.get('cname') != client:
            click.echo(
                click.style(
                    'The cname of your burp client does not match: '
                    '{} != {}'.format(confcli.get('cname'), client),
                    fg='yellow'
                )
            )
            errors = True
        if confcli.get('server') != host:
            click.echo(
                click.style(
                    'The burp server address does not match: '
                    '{} != {}'.format(confcli.get('server'), host),
                    fg='yellow'
                )
            )
            errors = True

    else:
        click.echo(
            click.style(
                'No client conf file found: {} does not exist'.format(bconfcli),
                fg='red'
            ),
            err=True
        )
        errors = True

    if os.path.exists(bconfsrv):
        parser = app.client.get_parser()

        confsrv = Config(bconfsrv, parser, 'srv')
        confsrv.set_default(bconfsrv)
        confsrv.parse()

        if host not in ['::1', '127.0.0.1']:
            bind = confsrv.get('status_address')
            if (bind and bind not in [host, '::', '0.0.0.0']) or not bind:
                click.echo(
                    click.style(
                        'It looks like your burp server is not exposing it\'s '
                        'status port in a way that is reachable by Burp-UI!',
                        fg='yellow'
                    )
                )
                click.echo(
                    click.style(
                        'You may want to set the \'status_address\' setting with '
                        'either \'{}\', \'::\' or \'0.0.0.0\' in the {} file '
                        'in order to make Burp-UI work'.format(host, bconfsrv),
                        fg='blue'
                    )
                )

        max_status_children = confsrv.get('max_status_children', [-1])
        if all([x < 15 for x in max_status_children]):
            click.echo(
                click.style(
                    '\'max_status_children\' is to low, you need to set it to '
                    '15 or more. Please edit your {} file.'.format(bconfsrv),
                    fg='blue'
                )
            )
            errors = True

        restore = []
        if 'restore_client' in confsrv:
            restore = confsrv.getlist('restore_client')

        if client not in restore:
            click.echo(
                click.style(
                    'Your burp client is not listed as a \'restore_client\'. '
                    'You won\'t be able to view other clients stats!',
                    fg='yellow'
                )
            )
            errors = True

        if 'monitor_browse_cache' not in confsrv or not \
                confsrv.get('monitor_browse_cache'):
            click.echo(
                click.style(
                    'For performance reasons, it is recommended to enable the '
                    '\'monitor_browse_cache\'.',
                    fg='yellow'
                )
            )
            errors = True

        ca_client_dir = confsrv.get('ca_csr_dir')
        if ca_client_dir and not os.path.exists(ca_client_dir):
            try:
                os.makedirs(ca_client_dir)
            except IOError as exp:
                click.echo(
                    click.style(
                        'Unable to create "{}" dir: {}'.format(ca_client_dir, exp),
                        fg='yellow'
                    ),
                    err=True
                )

        if confsrv.get('clientconfdir'):
            bconfagent = os.path.join(confsrv.get('clientconfdir'), client)
        else:
            click.echo(
                click.style(
                    'Unable to find "clientconfdir" option. Something is wrong '
                    'with your setup.',
                    fg='yellow'
                )
            )
            bconfagent = 'ihopethisfiledoesnotexistbecauseitisrelatedtoburpui'

        if not os.path.exists(bconfagent) and bconfagent.startswith('/'):
            click.echo(
                click.style(
                    'Unable to find the {} file.'.format(bconfagent),
                    fg='yellow'
                )
            )
            errors = True
        else:
            confagent = Config(bconfagent, parser, 'cli')
            confagent.set_default(bconfagent)
            confagent.parse()

            if confagent.get('password') != confcli.get('password'):
                click.echo(
                    click.style(
                        'It looks like the passwords in the {} and the {} files '
                        'mismatch. Burp-UI will not work properly until you fix '
                        'this.'.format(bconfcli, bconfagent),
                        fg='yellow'
                    )
                )
    else:
        click.echo(
            click.style(
                'Unable to locate burp-server configuration: {} does not '
                'exist.'.format(bconfsrv),
                fg='red'
            ),
            err=True
        )
        errors = True

    if errors:
        if not tips:
            click.echo(
                click.style(
                    'Some errors have been found in your configuration. '
                    'Please make sure you ran this command with the right flags! '
                    '(see --help for details).'.format(sys.argv[0], sys.argv[1]),
                    fg='red'
                ),
                err=True
            )
        else:
            click.echo(
                click.style(
                    '\n'
                    'Well, if you are sure about your settings, you can run the '
                    'following command to help you setup your Burp-UI agent. '
                    '(Note, the \'--dry\' flag is here to show you the '
                    'modifications that will be applied. Once you are OK with '
                    'those, you can re-run the command without the \'--dry\' flag):',
                    fg='blue'
                )
            )
            click.echo('    > bui-manage setup_burp --host="{}" --client="{}" --dry'.format(host, client))
    else:
        click.echo(
            click.style(
                'Congratulations! It seems everything is alright. Burp-UI '
                'should run without any issue now.',
                fg='green'
            )
        )


@app.cli.command()
@click.option('-v', '--verbose', is_flag=True,
              help='Dump parts of the config (Please double check no sensitive'
              ' data leaked)')
@click.option('-l', '--load', is_flag=True,
              help='Load all configured modules for full summary')
def sysinfo(verbose, load):
    """Returns a couple of system informations to help debugging."""
    from .desc import __release__, __version__
    import platform

    msg = None
    if load:
        try:
            msg = app.load_modules(True)
        except Exception as e:
            msg = str(e)

    backend_version = app.vers
    if not app.standalone:
        backend_version = 'multi'

    colors = {
        'True': 'green',
        'False': 'red',
    }
    embedded_ws = str(app.websocket)
    available_ws = str(WS_AVAILABLE)

    click.echo('Python version:      {}.{}.{}'.format(sys.version_info[0], sys.version_info[1], sys.version_info[2]))
    click.echo('Burp-UI version:     {} ({})'.format(__version__, __release__))
    click.echo('OS:                  {}:{} ({})'.format(platform.system(), platform.release(), os.name))
    if platform.system() == 'Linux':
        click.echo('Distribution:        {} {} {}'.format(*platform.dist()))
    click.echo('Single mode:         {}'.format(app.standalone))
    click.echo('Backend version:     {}'.format(backend_version))
    click.echo('WebSocket embedded:  {}'.format(click.style(embedded_ws, fg=colors[embedded_ws])))
    click.echo('WebSocket available: {}'.format(click.style(available_ws, colors[available_ws])))
    click.echo('Config file:         {}'.format(app.config.conffile))
    if load:
        if not app.standalone and not msg:
            click.echo('Agents:')
            for agent, obj in iteritems(app.client.servers):
                client_version = server_version = 'unknown'
                try:
                    app.client.status(agent=agent)
                    client_version = app.client.get_client_version(agent=agent)
                    server_version = app.client.get_server_version(agent=agent)
                except BUIserverException:
                    pass
                alive = obj.ping()
                if alive:
                    status = click.style('ALIVE', fg='green')
                else:
                    status = click.style('DISCONNECTED', fg='red')
                click.echo(' - {} ({})'.format(agent, status))
                click.echo('   * client version: {}'.format(client_version))
                click.echo('   * server version: {}'.format(server_version))
        elif not msg:
            server_version = 'unknown'
            try:
                app.client.status()
                server_version = app.client.get_server_version()
            except BUIserverException:
                pass
            click.echo('Burp client version: {}'.format(app.client.client_version))
            click.echo('Burp server version: {}'.format(server_version))
    if verbose:
        click.echo('>>>>> Extra verbose informations:')
        click.echo(click.style(
            '!!! PLEASE MAKE SURE NO SENSITIVE DATA GET EXPOSED !!!',
            fg='red'
        ))
        sections = [
            'WebSocket',
            'Burp',
            'Production',
            'Global',
        ]
        sections.reverse()
        for section in sections:
            if section in app.config.options:
                click.echo()
                click.echo('    8<{}BEGIN[{}]'.format('-' * (67 - len(section)), section))
                for key, val in iteritems(app.config.options.get(section, {})):
                    click.echo('    {} = {}'.format(key, val))
                click.echo('    8<{}END[{}]'.format('-' * (69 - len(section)), section))

    if load and msg:
        _die(msg, 'sysinfo')
