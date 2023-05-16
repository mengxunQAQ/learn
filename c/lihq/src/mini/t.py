import socket               
 
s = socket.socket()         
host = socket.gethostname() 
host = "10.120.6.250"
port = 12345                
s.bind((host, port)) 
s.listen(5)
import time
time.sleep(300)
