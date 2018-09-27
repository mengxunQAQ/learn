import redis


r = redis.StrictRedis(host='localhost', port=6379, db=0)


def throw(user, msg):
    if r.scard(user) < 10:
        r.sadd(user, msg)
        r.sadd("members", user)
        print("瓶子已经扔向大海...")
    else:
        print("今天扔瓶子机会已经用完~~~明天再来吧")


def pick(user):
    key = user+":count"
    r.incr(key)
    pick_count = int(r.get(key))
    if  pick_count < 10:
        member = r.srandmember("members", number=1)[0]
        msg = str(r.spop(member))
        print(msg)
    else:
        print("今天捡瓶子机会已经用完~~~明天再来吧")


if __name__ == '__main__':
        choice = input("捡瓶子还是瓶子？1-扔瓶子 2-捡瓶子：")
        if choice == "1":
            user = input("输入您的用户名：")
            while True:
                msg = input("说点想说的：")
                throw(user, msg)
                option = input("是否接着扔（y/n）：")
                if option != "y":
                    break

        elif choice == "2":
            user = input("输入您的用户名：")
            while True:
                pick(user)
                option = input("是否接着捡（y/n）：")
                if option != "y":
                    break

