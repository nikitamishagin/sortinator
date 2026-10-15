---
sessionId: session-261016-005122-1g12
---

# Requirements

### Overview & Goals

Enhance the zerolog ConsoleWriter with custom color mappings per log level and a cleaner timestamp format for improved
readability.

### Scope

- Modify `cmd/root.go` to configure custom colors and timestamp format
- No changes to `internal/` packages — logging calls remain unchanged
- No new dependencies beyond what's already in `go.mod`

### User Stories

- As a user running sortinator in a terminal, I want color-coded log levels that are easy to distinguish at a glance.
- As a developer, I want a clean, consistent timestamp format in the logs.

### Functional Requirements

- Error logs: red
- Warning logs: yellow
- Info logs: cyan (changed from default green)
- Timestamps: clean `HH:MM:SS` format
- Verbose output (with `-v`): statistics logged with appropriate colors
- Invalid IP entries: warning logs with yellow color

### Non-Functional Requirements

- Use only zerolog's built-in colorization (no new dependencies)
- Maintain backward compatibility — output is still valid console text
- Colors respect `NoColor` flag if terminal doesn't support it

# Technical Design

### Current Implementation

- `cmd/root.go` line 105-106 sets up zerolog with `ConsoleWriter`:
  ```go
  zerolog.SetGlobalLevel(zerolog.InfoLevel)
  log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
  ```
- Default zerolog colors: Trace=blue, Debug=none, Info=green, Warn=yellow, Error=Fatal=Panic=red
- Timestamp format: ISO8601 (default)

### Key Decisions

1. **Use zerolog's `FormatLevel` and `FormatTimestamp` custom formatters** — zerolog supports custom `Formatter`
   functions that receive the raw value and return a formatted string. This is the cleanest approach without forking
   zerolog.
2. **Use ANSI color codes directly** — zerolog already has `colorize()` helper, but it's unexported. We'll use raw ANSI
   escape sequences which zerolog's `colorize` uses internally.
3. **Cyan for Info level** — more distinctive from green terminal backgrounds than the default.

### Proposed Changes

**File: `cmd/root.go`**

Replace the `ConsoleWriter` initialization (lines 105-106) with:

```go
zerolog.SetGlobalLevel(zerolog.InfoLevel)

writer := zerolog.ConsoleWriter{
Out:        os.Stderr,
TimeFormat: "15:04:05",
FormatLevel: func (i interface{}) string {
return colorizeLevel(i)
},
}
log.Output(writer)
```

Where `colorizeLevel` uses ANSI codes:

```go
const (
ansiReset = "\033[0m"
ansiRed = "\033[31m"
ansiYellow = "\033[33m"
ansiCyan = "\033[36m"
ansiBold = "\033[1m"
)

func colorizeLevel(level interface{}) string {
l, ok := level.(string)
if !ok {
return ""
}
var color string
switch l {
case "ERR", "FAT", "PAN":
color = ansiRed
case "WRN":
color = ansiYellow
case "INF":
color = ansiCyan
default:
return l
}
return ansiBold + color + strings.ToUpper(l) + ansiReset
}
```

### Data Models / Contracts

- No new types or interfaces added
- Existing logging calls in `validator.go` and `root.go` remain unchanged

### Risks

- ANSI codes may not render correctly in non-ANSI terminals — zerolog's `go-colorable` handles Windows, but we should
  test
- The `colorizeLevel` function must handle input types correctly (zerolog passes the level string)

# Delivery Steps

###  * Step 1: Add custom color formatters to zerolog ConsoleWriter

Implement enhanced zerolog logging with custom color mappings and a cleaner timestamp format.

- Add ANSI color constants (red, yellow, cyan, bold, reset) to `cmd/root.go`
- Add a `colorizeLevel()` helper function that returns color-coded level strings (ERR=FAT=red, WRN=yellow, INF=cyan)
- Replace the existing `ConsoleWriter` initialization with a configured instance that includes:
    - `TimeFormat: "15:04:05"` for clean HH:MM:SS timestamps
    - `FormatLevel` set to `colorizeLevel` for custom colored level tags
- Ensure the `os` import is present for `os.Stderr`

### Step 2: Verify build and test with sample input

Build the project and verify the enhanced logging works correctly.

- Run `go build -o sortinator .` to confirm the build succeeds
- Run `./sortinator test.yaml -v` to verify colored output in the console
- Run `go vet ./...` to check for static analysis issues