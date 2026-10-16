# Sortinator

Sortinator is a Go CLI that reads a YAML file with an `ips:` key, sorts the IP addresses numerically (IPv4 and IPv6),
and prints the result. It was built to test local inference with llama.cpp.

## Goal

The main goal of creating the Sortinator tool is to learn how to set up local inference with llama.cpp and gain skills
in agent coding using the Junie CLI.

## Build

```bash
go build -o ./bin/sortinator .
# or
go install github.com/mishagin/sortinator@latest
```

## Usage

```bash
./bin/sortinator test.yaml
```

**Input** (`test.yaml`):

```yaml
ips:
  - 10.0.0.1
  - 8.8.8.8
  - 192.168.0.1
```

**Output**:

```yaml
ips:
  - 8.8.8.8
  - 10.0.0.1
  - 192.168.0.1
```

Invalid entries are skipped and logged as warnings.

## Flags

| Flag        | Short | Description                    |
|-------------|-------|--------------------------------|
| `--file`    | `-f`  | Write output to specified file |
| `--reverse` | `-r`  | Sort in descending order       |
| `--unique`  | `-u`  | Remove duplicate IPs           |
| `--verbose` | `-v`  | Show statistics summary        |
| `--version` |       | Print version                  |

## Examples

```bash
# Reverse sort, deduplicate, verbose
./bin/sortinator test.yaml -r -u -v

# Write output to a file
./bin/sortinator test.yaml -f output.yaml
```
