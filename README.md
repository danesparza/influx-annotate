# influx-annotate [![CircleCI](https://circleci.com/gh/danesparza/influx-annotate.svg?style=svg)](https://circleci.com/gh/danesparza/influx-annotate)
Command line tool to add annotations to an InfluxDB database.  

Suitable for batch files, deployment scripts, unicorn logging

## Quick start 

* Make sure you've already got [InfluxDB](https://docs.influxdata.com/influxdb/v1.2/introduction/installation/) up and running.  In InfluxDb, [create the database](https://docs.influxdata.com/influxdb/v1.2/guides/writing_data/#creating-a-database-using-the-http-api) that you want to log to.
* Get the [latest release](https://github.com/danesparza/influx-annotate/releases/latest) for your platform (it's just a single executable) 
* Log your annotation data:
```
influx-annotate -server="http://your-influxdb-server:8086" -tags="application=Super unicorn app" -value="Build_20170414.2"
```

If you need help, just run `influx-annotate -h`.

There are a few command line parameters available:

Parameter       | Description
----------      | -----------
server          | The [InfluxDB server](https://docs.influxdata.com/influxdb/v1.2/concepts/glossary/#server) to use.  Defaults to 'http://localhost:8086'
database        | The [InfluxDB database](https://docs.influxdata.com/influxdb/v1.2/concepts/glossary/#database) to use. Defaults to 'events'
name            | The [measurement name](https://docs.influxdata.com/influxdb/v1.2/concepts/glossary/#measurement).  Defaults to 'build'
value           | The [measurement value](https://docs.influxdata.com/influxdb/v1.2/concepts/glossary/#field-value)
tags            | Comma separated list of [tags](https://docs.influxdata.com/influxdb/v1.2/concepts/glossary/#tag-set).  Example: `application=My app name,env=dev,machine=server01`

