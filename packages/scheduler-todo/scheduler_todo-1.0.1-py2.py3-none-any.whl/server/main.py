from flask import Flask, jsonify, request
from pathlib import Path
import sqlite3, sys
from . import __version__
home_dir = str(Path.home())
app = Flask(__name__)

def scheduler_server():
    if len(sys.argv) == 2 and sys.argv[1] in ['version', 'ver', '--v']:
        print('scheduler-server version:' + __version__)
    else:
        conn = sqlite3.connect(home_dir + '/server.db')
        cur = conn.cursor()
        create_db = 'create table if not exists server(user text not null, year integer not null, category text not null, month integer not null, day integer not null, what text not null, done integer)'
        cur.execute(create_db)
        conn.close()
        app.run(host='0.0.0.0', port=8865)


@app.route('/pull/<user_id>')
def pull_user(user_id):
    conn = sqlite3.connect(home_dir + '/server.db')
    cur = conn.cursor()
    select_data = 'select * from server where user=?'
    cur.execute(select_data, (user_id,))
    result = cur.fetchall()
    print(result)
    data = {'account': user_id, 'result': result}
    print("calender from", user_id, "is pulled from server")
    return jsonify(data)


@app.route('/push/<user_id>', methods=['POST'])
def push_userr(user_id):
    conn = sqlite3.connect(home_dir + '/server.db')
    cur = conn.cursor()
    delete_db = 'delete from server where user=? and what=?'
    insert_db = 'insert into server (user, year, category, month, day, what, done) values (?,?,?,?,?,?,?)'
    received_json = request.get_json()
    content = received_json['result']
    for x in content:
        cal = [user_id] + x
        print(cal)
        cur.execute(delete_db, (cal[1], cal[4],))
        cur.execute(insert_db, cal)
    conn.commit()
    conn.close()
    print("calender from", user_id, "is pushed to server")
    return "You pushed your data", 200
