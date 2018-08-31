import thriftpy
pingpong_thrift = thriftpy.load("pingpong.thrift", module_name="pingpong_thrift")

from thriftpy.rpc import make_server

class Dispatcher(object):
    def ping(self):
        result = [1,2,3]
        print(result)
        return result  # error

server = make_server(pingpong_thrift.PingPong, Dispatcher(), '127.0.0.1', 6000)
print("servering......")
server.serve()
