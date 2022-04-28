#!/usr/bin/env python
import queue, threading
from urllib.request import urlopen

hosts = ["http://yahoo.com", "http://google.com", "http://amazon.com",
"http://ibm.com", "http://apple.com"]          

import threading, queue

q = queue.Queue()

def worker():
    while True:
        host = q.get()
        url = urlopen(host)
        print("\n" + host + "\n")
        print(url.read(1024))
        q.task_done()

# Turn-on the worker thread.
threading.Thread(target=worker, daemon=True).start()

# Send thirty task requests to the worker.
for host in hosts:
    q.put(host)


# Block until all tasks are done.
q.join()
print('All work completed')