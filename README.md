Filesync
===
fork from  github.com/elgs/filesync/gsyncd     

fix some bugs


Filesync is an utility written in Golang which helps you to keep the files on the client up to date with the files on the server. Only the changed parts of files on the server are downloaded. Therefore it's great to synchronize your huge, and frequently changing files.

Server
===
Installation
---

`go get -u github.com/garfcat/filesync/gsyncd`

Run
---
`gsyncd gsyncd.json`

Configuration
---
gsyncd.json
```json
{
    "ip": "0.0.0.0",
    "port": 6776,
    "monitors": {
        "home_garfcat_desktop_a": "/home/garfcat/Desktop/a",
        "home_garfcat_desktop_b": "/home/garfcat/Desktop/b"
    }
}
```


Client
===
Installtion
---

`go get github.com/garfcat/filesync/gsync`

Run
---
`gsync gsync.json`

Configuration
---
gsync.json
```json
{
    "ip": "127.0.0.1",
    "port": 6776,
    "monitors": {
        "home_garfcat_desktop_a": "/home/garfcat/Desktop/c",
        "home_garfcat_desktop_b": "/home/garfcat/Desktop/d"
    }
}
```
