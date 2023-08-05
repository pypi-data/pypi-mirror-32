# nameko-redis

Pony dependency for nameko services

## Installation
```
pip install nameko-pony
```

## Usage
app.py
```python
from nameko.rpc import rpc
from nameko_pony import Database, Required, PonySession


db = Database()

class Person(db.Entity):
    name = Required(str)


class MyService(object):
    name = "my_service"

    db_session = PonySession(db)

    @rpc
    def hello(self, name):
        with self.db_session:
            Person(name='hello')
```

config.yml
```yml
AMQP_URI: 'pyamqp://guest:guest@localhost'
DATABASE_URI: 'sqlite:///:memory:'
# DATABASE_URI: 'mysql://usr:pwd@localhost/db_name'
```