import redis


r = redis.StrictRedis(host='localhost', port=6379, db=0)

class Drift:
    def __init__(self):
        self.count_hash = "count"
        self.user_hash = "user"
        
    def throw(self, user, msg):
        
        key = user + ":throw"
    
        if r.hget(self.count_hash, key) != 10:
            r.hincrby(self.count_hash, key)
            count = r.hget(self.count_hash, key)
            msg_hash = user + ":{count}".format(count=count)
            r.sadd(self.user_hash, msg_hash)
            r.hmset("msg", {msg_hash:msg})
    
    
    def pick(self, user):
        key = user + ":pick"
    
        if r.hget(self.count_hash, key) != 10:
            r.hincrby(self.count_hash, key)
            msg_hash = r.srandmember(self.user_hash, 1)
            result = r.hget("msg", msg_hash[0])
        else:
            result = "Today's opportunity is gone!"
            return result


if __name__ == '__main__':
    drift = Drift()
    drift.throw("Alen", "good night!")
    result = drift.pick("Bob")
    print(result)


