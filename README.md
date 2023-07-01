# Logzo

logzo builds some useful stuff on top of slog

## Tater

tater implements a simple, single rotating file based os.Writer.

```go
	ttr := tater.New(filename)
	ttr.Write([]byte("spam"))  // writes "spam" to filename
    ttr.Rotate()               // moves filename to filename+".1"
    ttr.Write([]byte("eggs))   // writes "eggs" to filename
```

## Multipass

multipass multiplexes logging to multiple slog Handlers.

```go
	h0 := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
    h1 := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
    h := multipass.New(h0, h1)

	slog.SetDefault(slog.New(h))
    slog.Debug("spam")           // logs "spam" to stderr/text
    slog.Info("eggs")            // logs "eggs" to stderr/text and stdout/json
```

## Logging

logging combines slog, tater, and multipass to log messages as text to stderr and as json to a file.

```go
	logging.Init("./log.txt")
	slog.Info("spam")         // log "spam" to stderr/text and ./log.txt/json
    logging.Rotate()          // rotate ./log.txt to ./log.txt.1
```
