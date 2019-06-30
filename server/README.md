# tailf
This codebase was created to create client/server websocket application to implement a log watching solution 
(similar to the tail -f command in UNIX). However, in this case, the log file is hosted on a remote machine.


# How it works
```
.
├── main.go
├── handler.go
├── tail
│   └── tail.go
├── conn
│   ├── conn.go      // ws connection read/write
│   └── conn_mgr.go  // connection manager
├── file
│   └── file.go      // read, write to channel and file monitoring
└── mtime
    ├── mtime.go         // genric wrapper for mtime for the file
    ├── mtime_darwin.go  // mtime syscall specific to darwin
    ├── mtime_windows.go // mtime syscall specific to darwin
    └── mtime_linux.go   // mtime syscall specific to darwin
```

# Getting started

## Install the Golang
https://golang.org/doc/install
```
## Build

* `go build`

## Run

* `./tailf -port :8080 -file /var/log/system.log`
