---
sessionId: session-261016-133752-1f0m
---

# Requirements

### Overview & Goals

Change the CLI output format from plain text (one IP per line) to a YAML list, matching the input structure. This allows
users to pipe the output back into YAML pipelines or use it programmatically.

### Scope

- **In Scope**: Output formatting in `cmd/root.go` — replace the plain-text loop with YAML list serialization using the
  existing `gopkg.in/yaml.v3` dependency.
- **Out of Scope**: Adding new flags, modifying input parsing, changing internal sorter/validator logic, adding tests
  (unless explicitly requested).

### User Stories

- As a user, I want the sorted IPs printed as a YAML list so that I can use the output directly in other YAML documents.
- As a user, I want file output (`-f`) to also produce valid YAML when writing to a file.

### Functional Requirements

1. When no `-f` flag is given, stdout prints a YAML list of sorted IPs (e.g., `\n- 8.8.8.8\n- 10.0.0.1`).
2. When `-f output.yaml` is used, the file contains the same YAML list format.
3. Empty input (no valid IPs) prints an empty YAML list (`[]` or just nothing — consistent with YAML).
4. Verbose logging (`-v`) and all other flags remain unchanged.

# Technical Design

### Current Implementation

- `cmd/root.go` lines 101-105 build output as plain text by appending `ip + "\n"` for each IP.
- Output is printed via `fmt.Print()` (stdout) or written to a file via `os.WriteFile()`.
- The `IPConfig` struct (line 53-55) already uses `yaml:"ips"` from `gopkg.in/yaml.v3`, proving YAML serialization works
  in this project.

### Proposed Changes

**File**: `cmd/root.go`

Replace the plain-text output builder (lines 101-105):

```go
// Before
var output strings.Builder
for _, ip := range sortedIPs {
output.WriteString(ip + "\n")
}
```

With YAML list serialization:

```go
// After
type IPOutput struct {
IPs []string `yaml:"ips"`
}
yamlOutput, err := yaml.Marshal(IPOutput{IPs: sortedIPs})
if err != nil {
log.Fatal().Err(err).Msg("failed to marshal output")
return err
}
```

Then pass `string(yamlOutput)` to the existing file/stdout output logic (lines 117-125).

### Key Decisions

1. **Reuse `yaml.v3`** — already a dependency, no new packages needed.
2. **Wrap in `ips:` key** — output mirrors input structure (`IPConfig`), making it round-trippable.
3. **No new flags** — the `-f` flag behavior is unchanged, only content format changes.

### Risks

- **Backward compatibility**: Users expecting plain text will see a different format. This is a breaking change but
  aligns with the task request.
- **Empty output**: `yaml.Marshal` on an empty slice produces `ips: []`, which is valid YAML and acceptable.

# Testing

### Validation Approach

Build the binary and run it against `test.yaml` to verify YAML output format.

### Key Scenarios

1. `./bin/sortinator test.yaml` — verifies stdout produces a valid YAML list with `ips:` key.
2. `./bin/sortinator test.yaml -r -u` — verifies flags work correctly with YAML output.
3. `./bin/sortinator test.yaml -f /tmp/out.yaml && cat /tmp/out.yaml` — verifies file output is also YAML.

### Edge Cases

- Empty IP list → should produce `ips: []`.
- All entries invalid → same as empty (no valid IPs).
- Single IP → should produce a single-item YAML list.

# Delivery Steps

### ✓ Step 1: Replace plain-text output with YAML list serialization

Modify cmd/root.go to serialize the sorted IPs as a YAML list instead of plain text.

- Replace the plain-text `strings.Builder` loop (lines 101-105) with `yaml.Marshal()` wrapping IPs in an
  `IPOutput{IPs: sortedIPs}` struct.
- Pass the marshaled YAML string to the existing output logic (stdout or file).
- Keep all other logic (verbose stats, flags, error handling) unchanged.

### ✓ Step 2: Build and verify YAML output

Build the binary and run it against `test.yaml` to confirm correct YAML list output.

- Run `go build -o ./bin/sortinator .`
- Run `./bin/sortinator test.yaml` and verify stdout is valid YAML with `ips:` key.
- Run `./bin/sortinator test.yaml -r -u` and `-f /tmp/out.yaml` to verify flags work with the new format.