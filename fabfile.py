# pip3 install fabric

from fabric import Connection, task, connection

@task
def status(c):
    """
    systemctl status
    """
    # sudo コマンドで 'systemctl status' を叩く
    c.sudo('systemctl status')

@task
def get_pull(c):
    """
    git pull
    """
    with c.cd('/home/isucon/isucari'):
        c.run('git pull', warn=True)

@task
def put_nginx(c):
    """
    put nginx.conf
    """
    c.put('./nginx.conf', '/etc/nginx/nginx.conf')
    c.put('./isucari.conf', '/etc/nginx/sites-available/isucari.conf')

@task
def get_nginx(c):
    """
    get nginx.conf
    """
    c.get('/etc/nginx/nginx.conf', './nginx.conf')
    c.get('/etc/nginx/sites-available/isucari.conf', './isucari.conf')
