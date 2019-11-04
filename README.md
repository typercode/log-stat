# logStat
logStat for text file

## Install

`go get github.com/qiaodaimadelaowang/log-stat/...`

## Usage

```bash
➜  Downloads ./logStat -h
Usage of ./logStat:
  -f string
    	set the log file name(eg: Log.log.2019-11-03)
  -fp string
    	set the log file prefix,Ignore if f is set. (eg: MSSM-Auth.log.) (default "MSSM-Auth.log.")
  -ma string
    	set the mail address (eg: xxx@xxx.com)
  -p string
    	set the log file path(eg: /path/to/file/logs/app-name)
    	
    	
➜  ./logStat -p "/Users/tinyhuiwang/temp/a" -fp MSSM-Auth.log.

2019-11-03
filePath:/Users/tinyhuiwang/temp/a/MSSM-Auth.log.2019-11-03
mailContent:
appId: 1
api-version-id: 2,3

appId: 2
api-version-id: ,

appId: 3
api-version-id: 2,

appId: 1
api-version-id: 2,

appId: 2
api-version-id: 3,

appId: 1
api-version-id: 17,

appId: 2
api-version-id: ,

appId: 2
api-version-id: 1,

appId: 3
api-version-id: 17,47,

appId:
api-version-id: 1,

appId: id1
api-version-id: ,

appId: id2
api-version-id: 1,2,
```