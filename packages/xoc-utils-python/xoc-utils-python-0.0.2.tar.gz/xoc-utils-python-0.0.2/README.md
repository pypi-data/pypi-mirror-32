# xoc-utils-python

An XO core project in python that provides utils functionalities, including data processing and Kafka client.

----

## Kafka

### Producer

```
producer = KProducer('kafka_server_origin', 'service_name')
message = json.dumps({
    'hello': 'world'
})
producer.produce('topic_name', message)
```
Example [here](examples/kafka_producer.py)

### Consumer

Consumer without retry
```
consumer = KConsumer('kafka_server_origin', 'group_id', 'service_name', poll_timeout=1, retries=0)

# Add handler for a topic
def printMsg(msg):
    print('consuming message: ', msg)
    raise Exception('I am Exception.')
consumer.add_handler('topic_name', printMsg)
consumer.consume()
```
Example [here](examples/kafka_consumer.py).  
If you want to use retry feature, you can just set retries to the times you want it to retry, like 3.  
Example [here](examples/kafka_consumer_retry.py)


## Others
 
### Singleton
This provide you a way to create a singleton instance.

```
from xoc_utils import Singleton

class YourClass(metaclass=Singleton):
	pass
``` 


## Development
You get start from [here](https://packaging.python.org/tutorials/distributing-packages/#packages).

To debug your package, you don't have to release it, you can just run
```
python setup.py install
```
then it will be installed on your machine.