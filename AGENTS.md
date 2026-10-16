# AGENTS.md

Sortinator is a Go CLI that reads a YAML file with an `ips:` key, sorts the IPs numerically (IPv4 and IPv6) and prints
the result.

## Stack

Go, cobra (CLI), yaml (parsing), zerolog (logging).

## Layout

- `main.go`: entry point, calls `cmd.Execute()`
- `cmd/`: root command, flags and sort logic
- `internal/validator/`: IP validation and filtering
- `internal/sorter/`: numeric sorting, reverse, deduplication
- `bin/`: built binaries

## Commands

- Build: `go build -o ./bin/sortinator .`
- Test: `go test ./...`
- Format and vet: `gofmt -l .` and `go vet ./...`
- Try it: `./bin/sortinator test.yaml`

## Behavior rules

- Input is a positional YAML file path. Only the `ips:` key is supported.
- Invalid entries are skipped and logged as zerolog warnings. They do not stop execution.
- Flags: `-f/--file`, `-r/--reverse`, `-u/--unique`, `-v/--verbose`, `--version`.
- Output goes to stdout by default, or to a file with `-f`.
- Exit code is 0 on success and non-zero on fatal errors (missing file, YAML parse error).
- IPv4-mapped IPv6 addresses (`::ffff:a.b.c.d`) are distinct from their IPv4 form.
- Sort by binary value using `net.ParseIP` and `To16()`.

## Out of scope

DNS or network calls, config files, environment variables, other YAML keys, in-place editing.

## Code rules

- Keep business logic in `internal/`, CLI wiring in `cmd/`.
- Use zerolog for all logging. Do not use `fmt.Println` for diagnostics.
- Add no dependencies beyond cobra, yaml.v3 and zerolog without asking.
- Do not touch source code when the task is documentation only.

## Git policy

- Commit after each small, logical change. Do not batch unrelated changes.
- Examples of one commit each: a new package, adding logging, adding the root command, adding a flag, adding tests.
- The code must build before each commit.
- Commit messages follow this rule: <insert your commit message rule here>
- Never push, force-push, amend or rewrite history unless asked.
