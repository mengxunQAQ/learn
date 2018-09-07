import logging
import tornado.escape
import tornado.ioloop
import tornado.options
import tornado.web
import tornado.websocket
import os.path
import uuid
import json

from tornado.options import define, options

define("port", default=8888, help="run on the given port", type=int)

class Application(tornado.web.Application):
    def __init__(self):
        handlers = [
            (r"/", MainHandler),
            (r"/chatsocket/", ChatSocketHandler),
        ]
        settings = dict(
            cookie_secret="__TODO:_GENERATE_YOUR_OWN_RANDOM_VALUE_HERE__",
            template_path=os.path.join(os.path.dirname(__file__), "templates"),
            static_path=os.path.join(os.path.dirname(__file__), "static"),
            xsrf_cookies=True,
        )
        super(Application, self).__init__(handlers, **settings)


class MainHandler(tornado.web.RequestHandler):
    def get(self):
        room = self.get_argument("room", None)
        if room:
            ChatSocketHandler.create_new_room(room)
        self.finish(json.dumps(list(ChatSocketHandler.room.keys())))


class ChatSocketHandler(tornado.websocket.WebSocketHandler):
    cache_size = 200
    room = dict()

    def check_origin(self, origin):
        return True

    def get_compression_options(self):
        # Non-None enables compression with default options.
        return {}

    def open(self):
        self.room_code = self.get_argument("room", None)
        ChatSocketHandler.room[self.room_code]["waiters"].append(self)

    def on_close(self):
        ChatSocketHandler.room[self.room_code]["waiters"].append(self)

    @classmethod
    def update_cache(cls, chat):
        cls.room[chat["room"]]["cache"].append(chat)
        if len(cls.room[chat["room"]]["cache"]) > cls.cache_size:
            cls.room[chat["room"]]["cache"] = cls.room[chat["room"]]["cache"][-cls.cache_size:]

    @classmethod
    def send_updates(cls, chat):
        logging.info("sending message to %d waiters", len(cls.room[chat["room"]]["waiters"]))
        for waiter in cls.room[chat["room"]]["waiters"]:
            try:
                waiter.write_message(chat)
            except:
                logging.error("Error sending message", exc_info=True)

    @classmethod
    def create_new_room(self, room):
        if room not in self.room:
            self.room[room] = {'cache': [],
                               'waiters': []}
        return True

    def on_message(self, message):
        logging.info("got message %r", message)
        parsed = tornado.escape.json_decode(message)
        chat = {
            "room": self.room_code,
            "id": str(uuid.uuid4()),
            "body": parsed["body"],
        }
        # chat["html"] = tornado.escape.to_basestring(
        #     self.render_string("message.html", message=chat))

        ChatSocketHandler.update_cache(chat)
        ChatSocketHandler.send_updates(chat)


def main():
    tornado.options.parse_command_line()
    app = Application()
    app.listen(options.port, address='0.0.0.0')
    tornado.ioloop.IOLoop.current().start()


if __name__ == "__main__":
    main()


