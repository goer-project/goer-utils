# Logger
A logger based on [Zap](<https://pkg.go.dev/go.uber.org/zap>)

## Usage

```go
l := logger.NewChannel(logger.Channel{
    Path:    "/your-path/log.log",
    Level:   "debug",
    Days:    14,
    Console: true,
    Format: "json",
})

l.Info("This is a log.")
```