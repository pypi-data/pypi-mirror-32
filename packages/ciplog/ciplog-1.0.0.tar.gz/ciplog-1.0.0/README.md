# CipLog

CipLog is an easy Python package that used for you write your logs.

CipLog has three types of log, **info**, **warning** and **error**

CipLog supports python 3+.

# Examples 

## Install

> Under your virtualenv do:  

```
$ pip install ciplog
```

## Getting Started

You can set a service name and a log path or use default.
````python
from ciplog import  CipLog

log = CipLog(service_name='news-api', log_path='/your/log/path')

```` 

### Info
```python
from ciplog import CipLog

log = CipLog()

log.info(200, 'news', 'news that successfully registered.')

```

### Warning
```python
from ciplog import CipLog

log = CipLog()

log.warning('AB456*', 'news', 'The unicode is not defined.')
```

### Error
```python
from ciplog import CipLog

log = CipLog()

log.error('UX3023', 500, 'news', 'Service is timeout.')

```