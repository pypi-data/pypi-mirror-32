### redlock-dist
Implementation of distributed locking with [Redis](https://redis.io)

[Distributed locks with Redis](https://redis.io/topics/distlock)

### Usage
Install:

```
pip install redlock-dist
```

example:

```

from redis import StrictRedis
from redlock_dist import RedLock

rds = StrictRedis(host='localhost', port=6379, db=0)

"""
:param expire: lock key ttl(ms)
:param retry_times: retry times to get lock
:param retry_interval: retry interval(ms)
"""
with RedLock(rds, 'key-name', expires=30000, retry_times=3, retry_interval=50) as red_lock:

    if red_lock:
        print('get lock')
    else:
        print('do not get lock')
```

### Principle and notes
Because Redis is single-threaded and you can use `SET key value NX PX expires`
