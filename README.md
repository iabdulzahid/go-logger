# go-logger

## **Overview**
The Go Logger package provides a customizable logging solution with support for JSON and plain-text formats, log rotation, and dynamic log level configuration. This package is built using Uber's Zap library for high performance and flexibility.

---

## **Features**
- Supports different log levels: `debug`, `info`, `warn`, `error`, `fatal`.
- Outputs logs to multiple destinations (`stdout`, `file`, etc.).
- Supports JSON and plain-text formats.
- Includes caller information (file, function, and line number).
- Customizable timestamp format.
- File rotation with options for size, backup count, and compression.

---

## **Installation**
Add the logger package to your project:

```bash
go get github.com/iabdulzahid/go-logger
```
