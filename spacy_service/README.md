* simple spaCy parser micro-service wrapper

* production run without gevent (sync) can be run multi-threaded with w
  but takes x times the amount of RAM, better to use --threads

* CAREFUL: this app runs Tensorflow, --threads > 1 ruins the graph and the program will fail.

```
gunicorn --bind 0.0.0.0:8000 --timeout 120 --threads 1 server:app
```

* gevent is stricly single threaded only, --threads has no effects in gevent

```
gunicorn -k gevent --bind 0.0.0.0:8000 --timeout 120 server:app
```
