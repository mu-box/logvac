[![logvac logo](http://microbox.rocks/assets/readme-headers/logvac.png)](http://microbox.cloud/open-source#logvac)
[![Build Status](https://github.com/mu-box/logvac/actions/workflows/ci.yaml/badge.svg)](https://github.com/mu-box/logvac/actions)

# Logvac

Simple, lightweight, api-driven log aggregation service with realtime push capabilities and historical persistence.

## Routes:

| Route | Description | Payload | Output |
| --- | --- | --- | --- |
| **Get** /remove-token | Remove a log read/write token | *'X-USER-TOKEN' and 'X-AUTH-TOKEN' headers  | success message string |
| **Get** /add-token | Add a log read/write token | *'X-USER-TOKEN' and 'X-AUTH-TOKEN' headers  | success message string |
| **Post** / | Post a log | *'X-USER-TOKEN' header and json Log object | success message string |
| **Get** / | List all services | *'X-USER-TOKEN' header | json array of Log objects |
Note: * = only if 'auth-address' configured

### Query Parameters:
| Parameter | Description |
| --- | --- |
| **auth** | Replacement for 'X-USER-TOKEN' |
| **id** | Filter by id |
| **tag** | Filter by tag |
| **type** | Filter by type |
| **start** | Start time (unix epoch(nanoseconds)) at which to view logs older than (defaults to now) |
| **end** | End time (unix epoch(nanoseconds)) at which to view logs newer than (defaults to 0) |
| **limit** | Number of logs to read (defaults to 100) |
| **level** | Severity of logs to view (defaults to 'trace') |
`?id=my-app&tag=apache%5Berror%5D&type=deploy&start=0&limit=5`

## Data types:
### Log:
```json
{
  "id": "my-app",
  "tag": "build-1234",
  "type": "deploy",
  "priority": "4",
  "message": "$ mv microbox/.htaccess .htaccess\n[✓] SUCCESS"
}
```
| Field | Description |
| --- | --- |
| **time** | Timestamp of log (`time.Now()` on post) |
| **id** | Id or hostname of sender |
| **tag** | Tag for log |
| **type** | Log type (commonly 'app' or 'deploy'. default value configured via `log-type`) |
| **priority** | Severity of log (0(trace)-5(fatal)) |
| **message*** | Log data |
Note: * = required on submit


## Usage

add auth key - attempt
```
$ curl -ik https://localhost:6360/add-key -H 'X-USER-TOKEN: user'
HTTP/1.1 401 Unauthorized
```

add auth key - success
```
$ curl -ik https://localhost:6360/add-key -H 'X-USER-TOKEN: user' -H 'X-AUTH-TOKEN: secret'
HTTP/1.1 200 OK
```

publish log - attempt
```
$ curl -ik https://localhost:6360 -d '{"id":"my-app","type":"deploy","message":"$ mv microbox/.htaccess .htaccess\n[✓] SUCCESS"}'
HTTP/1.1 401 Unauthorized
```

publish log - success
```
$ curl -ik https://localhost:6360 -H 'X-USER-TOKEN: user' -d '{"id":"my-app","type":"deploy","message":"$ mv microbox/.htaccess .htaccess\n[✓] SUCCESS"}'
sucess!
HTTP/1.1 200 OK
```

get deploy logs
```
$ curl -k https://localhost:6360?kind=deploy -H 'X-USER-TOKEN: user'
[{"time":"2016-03-07T15:48:57.668893791-07:00","id":"my-app","tag":"","type":"deploy","priority":0,"message":"$ mv microbox/.htaccess .htaccess\n[✓] SUCCESS"}]
```

get app logs
```
$ curl -k https://localhost:6360 -H 'X-USER-TOKEN: user'
[]
```

### Contributing

Contributions to the logvac project are welcome and encouraged. Logvac is a [Microbox](https://microbox.cloud) project and contributions should follow the [Microbox Contribution Process & Guidelines](https://docs.microbox.cloud/contributing/).

### Licence

This project is released under [The MIT License](http://opensource.org/licenses/MIT).

[![open source](http://microbox.rocks/assets/open-src/microbox-open-src.png)](http://microbox.cloud/open-source)
