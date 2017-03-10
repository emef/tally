### tally

Metric recording service.

-----

install

```
PROTOC=/usr/local/bin/protoc make clean all install`
```

-----

todo (semi-prioritized order)

* client package(s)
* testing app for data collection
* frontend for charts / dashboards
* process to clean-up/merge intra-hourly files (compactor)
* convert to docker swarm
* all-in-one cmd
* load testing / benchmarks
* query aggregations / expressions
* > 1m granularity via aggregation in tree walking
* (alternatively) different granularity trees
* config file
* health / load factor

un-prioritized

* proper logging
* code documentation
* system documentation