# Log

DEPRECATED: use github.com/yadisnel/go-ms/v2/logger interface

This is the global logger for all go-ms based libraries.

## Set Logger

Set the logger for go-ms libraries

```go
// import go-ms/util/log
import "github.com/yadisnel/go-ms/util/log"

// SetLogger expects github.com/yadisnel/go-ms/debug/log.Log interface
log.SetLogger(mylogger)
```
