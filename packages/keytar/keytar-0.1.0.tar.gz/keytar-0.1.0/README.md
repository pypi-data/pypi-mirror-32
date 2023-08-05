# Key Value Database Wrapper

_Prevent vendor lock-in with shared functionality._

This is a pip package aimed at easing switching between multiple NoSql databases.

## Use Case

The package's main goal is a common wrapper around Key Value databases. Making the code
portable across multiple providers.

This is in no means a package that takes good advantage of any of the various databases'
strengths. Making it ill-suited at the moment for resource intensive or optimization
focuesed projects. The only goal is to provide a common abstraction amongst simple key
value commands.

## Example usage

In this example, a table was created on AWS DynamoDb. This connects and modifies the resource.

```python
from keytar import Database, Key, db_types
import boto3

TABLE = "keytar_test"

# Initialization does require an already established connection
resource = boto3.resource("dynamodb", region_name="us-east-1")
db = Database(resource, db_types.AWS_DYNAMO)

db.set(TABLE, Key("greg", key_name="name"), {'age': 10})
res = db.get(TABLE, Key("greg", key_name="name"))

print(res)
```

## Suported Types

+ AWS Dyanamo Db
  - Get & Set operations
+ GCP Cloud Datastore
  - Get & Set operations
  - Custom queries
  - Random keys
+ Redis
  - Get & Set operations
