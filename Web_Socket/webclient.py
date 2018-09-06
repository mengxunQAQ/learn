from websocket import create_connection
import json
ws = create_connection("ws://192.168.0.243:8888/chatsocket")
ws.send(json.dumps({"body":123}))
result =  ws.recv()
print(result)
ws.close()