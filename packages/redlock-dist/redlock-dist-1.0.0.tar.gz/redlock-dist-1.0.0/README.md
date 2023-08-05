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

with RedLock(rds, 'key-name') as red_lock:
    if red_lock:
        print('get lock')
    else:
        print('do not get lock')
```

### Principle and notes
Because Redis is single-threaded and you can use `SET key value NX PX expires`
