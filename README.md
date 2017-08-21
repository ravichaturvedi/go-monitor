# GoMonitor

[![Build Status](https://travis-ci.org/ravichaturvedi/go-monitor.svg?branch=master)](https://travis-ci.org/ravichaturvedi/go-monitor)

Golang framework to monitor custom components via a HTTP Server.
It can be embedded in an application or can run as a standalone binary. 



## Getting Started

Download prebuilt `gomonitor-mac` (for MAC OSX platform) binary from the [Release](https://github.com/ravichaturvedi/go-monitor/releases) page.

For other platforms, checkout the sources and build `go-monitor/cmd/gomonitor/main.go` locally with `go 1.8 or higher`

Run the binary as follows.

```
./gomonitor-mac

...
Starting server: http://0.0.0.0:1234
...
```

## Feature

- Embeddable within an application to monitor its components.
- Standalone `gomonitor` to run as `sidecars`.
- Remote monitor via http server listening on port `1234`
- Scheduling plugin to run after every time interval
- Internal caching for scheduled plugin.
- Extensible for writing custom plugin.  


## Background

It's always better to have some monitoring enabled for the application so that we can monitor the application remotely.
However, there doesn't seem to be a framework available that can enable custom component monitoring by custom code execution.

So built `gomonitor` to support the same use case where developer can embed it within its own application defining their custom component monitor as pluggable components. 


## Usages

Following snippet provides good overview to write custom plugin and start `gomonitor`.

```go
func main() {
    p := map[string]plugin.Plugin{"hello": plugin.New(func() (interface{}, error) {
            return "Hello World!", nil
        })}
        
    d := map[string]time.Duration{"hello": time.Second}
    
	monitor.New(p, d).Start(context.Background())
}
```