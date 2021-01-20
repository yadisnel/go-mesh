# Log

DEPRECATED: use github.com/yadisnel/go-ms/v2/logger interface

This is the global logger for all ms based libraries.

## Set Logger

Set the logger for ms libraries

```go
// import go-ms/util/log
import "github.com/yadisnel/go-ms/v2util/log"

// SetLogger expects github.com/yadisnel/go-ms/v2debug/log.Log interface
log.SetLogger(mylogger)
```
