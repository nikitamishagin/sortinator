# Sortinator

Sortinator is a Go CLI that reads a YAML file with an `ips:` key, sorts the IP addresses numerically (IPv4 and IPv6), and prints the result. It was built to test local inference with llama.cpp.

## Build

```bash
go build -o sortinator .
# or
go install github.com/mishagin/sortinator@latest
```

## Usage

```bash
./sortinator test.yaml
```

**Input** (`test.yaml`):

```yaml
ips:
  - 10.0.0.1
  - 8.8.8.8
  - 192.168.0.1
```

**Output**:

```
8.8.8.8
10.0.0.1
192.168.0.1
```

Invalid entries are skipped and logged as warnings.

## Flags

| Flag            | Short | Description                  |
|-----------------|-------|------------------------------|
| `--file`        | `-f`  | Write output to specified file |
| `--reverse`     | `-r`  | Sort in descending order     |
| `--unique`      | `-u`  | Remove duplicate IPs         |
| `--verbose`     | `-v`  | Show statistics summary      |
| `--version`     |       | Print version                |

## Examples

```bash
# Reverse sort, deduplicate, verbose
./sortinator test.yaml -r -u -v

# Write output to a file
./sortinator test.yaml -f output.txt
```
