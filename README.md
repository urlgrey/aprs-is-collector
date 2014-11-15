APRS-IS Collector [![Circle CI](https://circleci.com/gh/urlgrey/aprs-is-collector.png?style=badge)](https://circleci.com/gh/urlgrey/aprs-is-collector)
==============

APRS-IS Collector that sends APRS-IS packets to an [APRS-Dashboard service](https://github.com/urlgrey/aprs-dashboard) endpoint.

Installation
------------

### Build
**Note**: You must have a Go compiler installed in order to build this project.

```shell
go build -o aprs-is-collector
```

### Running
Set the `APRS_DASHBOARD_HOST` environment variable to indicate the APRS-Dashboard installation location:
```shell
export APRS_DASHBOARD_HOST="127.0.0.1:3000"
```

Run the binary:
```shell
aprs-is-collector
```
