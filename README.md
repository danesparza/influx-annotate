# influx-annotate
Command line tool to add annotations to an InfluxDB database.  Suitable for batch files, deployment scripts, unicorn logging

## Quick start 
Get the [latest release](https://github.com/danesparza/influx-annotate/releases/latest) for your platform -- it's just a single executable.  

Make sure you've already got [InfluxDB](https://docs.influxdata.com/influxdb/v1.2/introduction/installation/) up and running.  In InfluxDb, [create the database](https://docs.influxdata.com/influxdb/v1.2/guides/writing_data/#creating-a-database-using-the-http-api) that you want to log to.

Log your annotation data:
```
influx-annotate -server="http://your-influxdb-server:8086" -tags="application=Super unicorn app" -value="Build_20170414.2"
```



