[![logvac logo](http://assets.microbox.rocks/readme-headers/logvac.png)](http://microbox.cloud/open-source#logvac)  
[![Build Status](https://travis-ci.org/mu-box/logvac.svg)](https://travis-ci.org/mu-box/logvac)

# Logvac

Simple, lightweight, api-driven log aggregation service with realtime push capabilities and historical persistence.

## Usage

Logvac can receive logs from rsyslog
>/etc/rsyslog.d/01-logvac-example.conf
>```
# rsyslog.conf style - more info look at rsyslog.conf(5)
# Single '@' sends to UDP
*.* @127.0.0.1:514
# Double '@' sends to TCP
*.* @@127.0.0.1:6361
```
> `sudo service rsyslog restart` with the preceding config file should start dumping logs to logvac

See http examples [here](../api/README.md)  

### Contributing

Contributions to the logvac project are welcome and encouraged. Logvac is a [Microbox](https://microbox.cloud) project and contributions should follow the [Microbox Contribution Process & Guidelines](https://docs.microbox.cloud/contributing/).

### Licence

Mozilla Public License Version 2.0

[![open source](http://assets.microbox.rocks/open-src/microbox-open-src.png)](http://microbox.cloud/open-source)
