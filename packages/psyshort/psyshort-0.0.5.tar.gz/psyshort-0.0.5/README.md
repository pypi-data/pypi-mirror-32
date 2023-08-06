Making psycopg2 a little simpler to use.

Usage:
```python
from psyshort import Psyshort
psy = Psyshort(
    hostname="db.example.com",
    dbname="example_database",
    username="postgres_user",
    password="pa$$w0rd"
    )
```