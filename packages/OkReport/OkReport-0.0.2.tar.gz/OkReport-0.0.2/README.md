# OkReport

[![](https://img.shields.io/badge/pip-v0.0.2%20okcat-yellow.svg)](https://pypi.python.org/pypi/OkReport)

Send line-chart report of okcheck

## Install

```
pip install okreport
```

## Help

```
okreport --help
```

![](https://git.llsapp.com/client-infra/okreport/raw/master/arts/help.png)

## Config

Create file:

```
[MAILGUN]
DomainName = <domain name>
ApiKey = <api key>
From = <name><email address>
To = <email address>
Cc = <email address 1>,<email address 2>
Subject = <subject>
```

And provide the config file path to okreport when run the command as: `okreport -c=/path/to/config/file`
