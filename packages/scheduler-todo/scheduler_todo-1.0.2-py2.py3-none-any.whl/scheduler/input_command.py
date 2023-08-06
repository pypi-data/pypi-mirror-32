from .valid_date import valid_date
from .make_string import make_string
from termcolor import colored
from pathlib import Path
from colorama import init
import sqlite3, requests, json
from . import __version__

def print_calendar(year, month):
    import calendar
    c = calendar.TextCalendar(calendar.SUNDAY)
    string = c.formatmonth(year, month)
    print(string)


def done_percent(all, done):
    number_of_done = int(done * 100 / all)
    print(colored('>' * (number_of_done // 2), 'green') + colored('>' * ((100 - number_of_done)//2), 'red'))
    print("{}% Completed".format(number_of_done))


def input_command(command):
    init()
    home_dir = str(Path.home())
    if type(command) is str:
        command = command.split(' ')
    conn = sqlite3.connect(home_dir + "/scheduler.db")
    cur = conn.cursor()
    # 전체 도움말
    title_string = "%-35s|%-36s|%-40s\n"%("function", "command", "example")
    add_help_string = ("-" * 35 + '+') + ("-" * 36 + '+') + ("-" * 30 + '+') + '\n'
    add_help_string += "%-35s|%-45s|%-40s\n"%("add schedule with category", colored("add {due} {content} in {category}", 'yellow'), colored("add 2018/3/2 go school in school", 'cyan'))
    add_help_string += "%-35s|%-45s|%-40s\n"%("add schedule without category", colored("add {due} {content}", 'yellow'), colored("add 2018/3/2 go school", 'cyan'))
    delete_help_string = ("-" * 35 + '+') + ("-" * 36 + '+') + ("-" * 30 + '+') + '\n'
    delete_help_string += "%-35s|%-45s|%-40s\n"%("delete all schedule", colored("delete all", 'yellow'), colored("delete all", 'cyan'))
    delete_help_string += "%-35s|%-45s|%-40s\n"%("delete schedule with category", colored("delete in {category}", 'yellow'), colored("delete in hoesung", 'cyan'))
    delete_help_string += "%-35s|%-45s|%-40s\n"%("delete schedule with content", colored("delete {content}", 'yellow'), colored("delete hit hoesung", 'cyan'))
    update_help_string = ("-" * 35 + '+') + ("-" * 36 + '+') + ("-" * 30 + '+') + '\n'
    update_help_string += "%-35s|%-45s|%-40s\n"%("update state with content", colored("update {content} {done/undone}", 'yellow'), colored("update hit hoesung done", 'cyan'))
    update_help_string += "%-35s|%-45s|%-40s\n"%("update due date with content", colored("update {content} at {due}", 'yellow'), colored("update hit hoesung at 2018/7/1", 'cyan'))
    update_help_string += ("-" * 35 + '+') + ("-" * 36 + '+') + ("-" * 30 + '+') + '\n'
    update_help_string += "%-35s|%-45s|%-40s\n"%("update state with category", colored("update in {category} {done/undone}", 'yellow'), colored("update in school done", 'cyan'))
    update_help_string += "%-35s|%-45s|%-40s\n"%("update due date with category", colored("update in {category} at {due}", 'yellow'), colored("update in school at 2018/7/1", 'cyan'))
    show_help_string = ("-" * 35 + '+') + ("-" * 36 + '+') + ("-" * 30 + '+') + '\n'
    show_help_string += "%-35s|%-45s|%-40s\n"%("get all schedule", colored("show all", 'yellow'), colored("show all", 'cyan'))
    show_help_string += "%-35s|%-45s|%-40s\n"%("get schedule with content", colored("show {content}", 'yellow'), colored("show hit hoesung", 'cyan'))
    show_help_string += "%-35s|%-45s|%-40s\n"%("get all schedule in category", colored("show in {category}", 'yellow'), colored("show in school", 'cyan'))
    show_help_string += "%-35s|%-45s|%-40s\n"%("get all schedule at month", colored("show in {month}", 'yellow'), colored("show at july", 'cyan'))
    show_help_string += "%-35s|%-45s|%-40s\n"%("get all calender at specific month", colored("show cal {year/month}", 'yellow'), colored("show cal 2018/03", 'cyan'))
    full_help_string = title_string + add_help_string + delete_help_string + update_help_string + show_help_string
    add_help_string = title_string + add_help_string
    show_help_string = title_string + show_help_string
    update_help_string = title_string + update_help_string
    delete_help_string = title_string + delete_help_string
    # 새 스케쥴 만드는 sql 구문
    insert_data = 'insert into todo (category, year, month, day, what, done) values (?,?,?, ?, ?, ?)'
    # 모든 스케쥴 삭제 sql 구문
    delete_all_data = 'delete from todo'
    # 내용으로 스케쥴 삭제 sql 구문
    delete_data = 'delete from todo where what = ?'
    # 분류로 스케쥴 삭제 sql 구문
    delete_data_cat = 'delete from todo where category = ?'
    # 모든 스케쥴 선택 sql 구문
    select_data_all = 'select * from todo order by year, month, day'
    select_data_all_done = 'select * from todo where done = 1 order by year,month, day'
    select_data_all_undone = 'select * from todo where done = 0 order by year,month, day'
    # id로 스케쥴 선택 sql 구문
    select_data = 'select * from todo where what=? order by year, month, day'
    select_data_done = 'select * from todo where what=? and  done = 1 order by year,month, day'
    # 카테고리로 스케쥴 선택 sql 구문
    select_data_cat = 'select * from todo where category=?'
    select_data_cat_done = 'select * from todo where category=? and done = 1'
    # 월별로 스케쥴 선택 sql 구문
    select_data_mon = 'select * from todo where month=?'
    select_data_mon_done = 'select * from todo where month=? and done = 1'
    month_dic = {"January": 1, "january": 1, "Jan": 1, "jan": 1, "February": 2, "Feb": 2, "february": 2, "feb": 2, "March": 3, "march": 3,
                 "Mar": 3, "mar": 3, "April": 4, "april": 4, "Apr": 4, "apr": 4, "May": 5, "may": 5, "June": 6, "june": 6, "July": 7, "july": 7,
                 "August": 8, "august": 8, "Aug": 8, "aug": 8, "September": 9, "Sep": 9, "september": 9, "sep": 9, "October": 10, "Oct": 10,
                 "october": 10, "oct": 10, "November": 11, "Nov": 11, "november": 11, "nov": 11, "December": 12, "Dec": 12, "december": 12, "dec": 12}
    # id로 일정 업데이트 sql 구문
    update_data_by_id = 'update todo set done = ? where what = ?'
    update_data_by_id_due = 'update todo set year = ?, month = ?, day = ? where what = ?'
    update_data_by_cat = 'update todo set done = ? where category = ?'
    update_data_by_cat_due = 'update todo set year = ?, month = ?, day = ? where category = ?'
    if command[0] == 'add':
        if len(command) > 2:
            # 명령의 두번째 단어에 /가 없으면 날짜가 없는 것으로 간주하고, 추가하지 않는다.
            if '/' in command[1] and len(command[1].split('/')) == 3:
                year, month, day = command[1].split('/')
                if int(year) > 9999:
                    print("Please input year under 10000")
                    return 1
                if not valid_date(int(year), int(month), int(day)):
                    print("Date is not valid")
                    return 1
                # in 키워드를 통해 어느 category에 넣을지 정할 수 있다.
                if 'in' in command:
                    category_split = command.index('in')
                    category_list = command[category_split + 1:]
                    content_list = command[2:category_split]
                # in 키워드가 없으면 no category 처리한다.
                else:
                    category = 'No category'
                    content_list = command[2:]
                # content, category를 띄어쓰기로 묶기
                content = ''
                for x in content_list:
                    content += x + ' '
                if len(content) > 22:
                    print("plz enter content less than 20 letters")
                    return 1
                category = ''
                if 'in' in command:
                    for x in category_list:
                        category += x + ' '
                else:
                    category = "No category"
                category = category.strip()
                content = content.strip()
                cur.execute(insert_data, (category, int(year), int(month), int(day), content, 0))
                print('schedule ' + content + ' in ' + category + ' at ' + command[1])
                conn.commit()
            # 명령어에 날짜가 없는 경우 다시 입력 받기
            else:
                print('There is no date in command')
        # 날짜도 없으면 스케쥴 추가 문자열 출력
        else:
            print(add_help_string.strip())
    elif command[0] == 'delete':
        # 모두 스케쥴 삭제
        if command[1] == 'all':
            cur.execute(delete_all_data)
            conn.commit()
        # 카테고리로 스케쥴 찾아서 삭제:
        elif command[1] == 'in':
            string = ''
            command_list = command[2:]
            for x in command_list:
                string += x + ' '
            string = string.strip()
            try:
                cur.execute(delete_data_cat, (string,))
                conn.commit()
                print(string, "분류의 일정이 제거되었습니다.")
            except:
                print(delete_help_string.strip())
        # id로 스케쥴 찾아서 삭제
        else:
            try:
                string = ''
                command_list = command[1:]
                for x in command_list:
                    string += x + ' '
                string = string.strip()
                cur.execute(delete_data, (string,))
                conn.commit()
                print(string, " 일정이 제거되었습니다.")
            # id로 된 일정이 없으면 예외처리
            except:
                print(delete_help_string.strip())
            cur.execute(select_data_all)
            result = cur.fetchall()
            print(make_string(result))
    elif command[0] == 'show':
        all_length = 0
        done_length = 0
        # 스케쥴 모두 보기
        if command[1] == 'all':
            cur.execute(select_data_all)
            result = cur.fetchall()
            all_length = len(result)
            print(make_string(result))
            cur.execute(select_data_all_done)
            done_length = len(cur.fetchall())
        # 카테고리로 스케쥴 검색
        elif command[1] == 'in':
            category = ''
            for x in command[2:]:
                category += x + ' '
            cur.execute(select_data_cat, (category.strip(),))
            result = cur.fetchall()
            all_length = len(result)
            print(make_string(result))
            conn.commit()
            cur.execute(select_data_cat_done, (category.strip(), ))
            done_length = len(cur.fetchall())
        # 날짜의 월로 검색
        elif command[1] == 'at':
            if command[2] in month_dic.keys():
                cur.execute(select_data_mon, (month_dic[command[2]],))
                result = cur.fetchall()
                all_length = len(result)
                print(make_string(result))
                cur.execute(select_data_mon_done, (month_dic[command[2]],))
                done_length = len(cur.fetchall())
            else:
                print("invalid month")
        # 끝난 일정 검색
        elif command[1] == "done":
            cur.execute(select_data_all_done)
            result = cur.fetchall()
            print(make_string(result))
        # 끝나지 못한 일정 검색
        elif command[1] == "undone":
            cur.execute(select_data_all_undone)
            result = cur.fetchall()
            print(make_string(result))
        # 달력 보기 기능
        elif (command[1] == 'calender' or command[1] == 'cal') and len(command) == 3:
            if '/' in command[2] and command[2].split("/")[0].isdigit() and command[2].split("/")[1].isdigit():
                if 70 < int(command[2].split("/")[0]) < 100:
                    year = int(command[2].split("/")[0]) + 1900
                elif 0 < int(command[2].split("/")[0]) <= 70:
                    year = int(command[2].split("/")[0]) + 2000
                else:
                    year = int(command[2].split("/")[0])
                print_calendar(year, int(command[2].split("/")[1]))
            else:
                print(show_help_string.strip())
        # 제목으로 검색 보기
        else:
            content_list = command[1:]
            content = ''
            for x in content_list:
                content += x + ' '
            content = content.strip()
            try:
                cur.execute(select_data, (content,))
                result = cur.fetchall()
                all_length = len(result)
                print(make_string(result))
                cur.execute(select_data_done, (content, ))
                done_length = len(cur.fetchall())
            except:
                print(show_help_string)
        # 완료한 비율 검색
        if all_length != 0:
            done_percent(all_length, done_length)
    # 일정을 끝났는지 안끝났는지 명령어
    elif command[0] == 'update':
        # 두번째 키워드가 category이면,
        if command[1] == 'in' and len(command) > 3:
            category = ''
            if command[-1] == 'done':
                for x in category[2:-1]:
                    category += x + ' '
                if category == '':
                    category = command[-2]
                cur.execute(select_data_cat, (category.strip(),))
                result = cur.fetchall()
                if not result:
                    print('no schedule found')
                    return 1                
                print('schedule in :', category, '\'s state is changed to done')
                cur.execute(update_data_by_cat, (1, category,))
                conn.commit()
            elif command[-1] == 'undone':
                for x in category[2:-1]:
                    category += x + ' '
                if category == '':
                    category = command[-2]                    
                cur.execute(select_data_cat, (category.strip(),))
                result = cur.fetchall()
                if not result:
                    print('no schedule found')
                    return 1 
                print('schedule in :', category, '\'s state is changed to undone')
                cur.execute(update_data_by_id, (0, category,))
                conn.commit()
            elif command[-2] == 'at' and len(command[-1].split('/')) == 3:
                for x in category[2:-2]:
                    category += x + ' '
                if category == '':
                    category = command[-3]                
                cur.execute(select_data_cat, (category.strip(),))
                result = cur.fetchall()
                if not result:
                    print('no schedule found')
                    return 1
                year, month, day = command[-1].split('/')
                if not valid_date(int(year), int(month), int(day)):
                    print("Date is not valid")
                    return 1
                print('schedule in :', category, '\'s due is changed to ' + command[4])
                cur.execute(update_data_by_cat_due, (int(year), int(month), int(day), category,))
                conn.commit()
            else:
                print(update_help_string)
                return 1
        elif command[1] == 'help':
            print(update_help_string)
        # 이외 사항이면 제목으로 검색
        else:
            if command[-1] in ['done', 'undone']:
                content = ''
                position = command[1:-1]
                for x in position:
                    content += x + ' '
                position = content.strip()
                try:
                    cur.execute(select_data, (position,))
                    result = cur.fetchall()
                except:
                    print('no schedule found')
                    return 1
            elif command[-2] == 'at':
                content = ''
                position = command[1:-2]
                for x in position:
                    content += x + ' '
                position = content.strip()
                try:
                    cur.execute(select_data, (position,))
                    result = cur.fetchall()
                except:
                    print('no schedule found')
                    return 1                
            if command[-1] == 'done':
                print('content:', position, '\'s state is changed to done')
                cur.execute(update_data_by_id, (1, position,))
                conn.commit()
            elif command[-1] == 'undone':
                print('content:', position, '\'s state is changed to undone')
                cur.execute(update_data_by_id, (0, position,))
                conn.commit()
            elif command[-2] == 'at' and len(command[-1].split('/')) == 3:
                try:
                    year, month, day = command[-1].split('/')
                except:
                    print("Date is not valid")
                    return 1
                if not valid_date(int(year), int(month), int(day)):
                    print("Date is not valid")
                    return 1
                print('content:', position, '\'s due is changed to ' + command[-1])
                cur.execute(update_data_by_id_due, (int(year), int(month), int(day), position,))
                conn.commit()
            conn.commit()
        conn.commit()
    elif command[0] == 'pull':
        address = input("please input ip address of server: ")
        account = input("please input your account: ")
        link = 'http://' + address + ':8865/pull/' + account
        try:
            response = requests.get(link)
            received_json = json.loads(response.text)
            print("Server Connected")
            print("Loading")
        except:
            print("server error!")
            return 1
        result = received_json['result']
        insert_data = 'insert into todo (year, category, month, day, what, done) values (?,?,?, ?, ?, ?)'
        for x in result:
            delete_data = 'delete from todo where year=? and category=? and month=? and day=? and what=?'
            cur.execute(delete_data, x[1:-1])
            cur.execute(insert_data, x[1:])
        conn.commit()
        conn.close()
    elif command[0] == 'push':
        address = input("please input ip address of server: ")
        account = input("please input your account: ")
        link = 'http://' + address + ':8865/push/' + account
        cur.execute(select_data_all)
        result = cur.fetchall()
        send_json = {'result': result}
        try:
            requests.post(link, json=send_json)
            print("Server Connected")
            print("Loading")
        except:
            print("server error!")
            return 1
        conn.close()
    elif command[0] == 'sync':
        address = input("please input ip address of server: ")
        account = input("please input your account: ")
        link = 'http://' + address + ':8865/push/' + account
        cur.execute(select_data_all)
        result = cur.fetchall()
        send_json = {'result': result}
        delete_db = 'delete from todo where what=? and month=? and day=?'
        try:
            requests.post(link, json=send_json)
            print("Server Connected")
            print("Loading")
        except:
            print("server error!")
            return 1       
        link = 'http://' + address + ':8865/pull/' + account
        try:
            response = requests.get(link)
            received_json = json.loads(response.text)
            print("Server Connected")
            print("Loading")
        except:
            print("server error!")
            return 1
        result = received_json['result']
        for x in result:
            cur.execute(delete_db, (x[4], x[2], x[3]))
            cur.execute(insert_data, x[1:])
        print("Successfully sync from server.")
        conn.commit()
        conn.close()
    elif command[0] == 'exit':
        return 1
    elif command[0] == 'help':
        print(full_help_string.strip())
    elif command[0] == 'version':
        print(__version__)
    else:
        print('To get more info, plz input command \'help\'')
    conn.close()
    return 0
