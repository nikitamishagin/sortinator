---
sessionId: session-261014-221934-1rwl
---

# Requirements

### Overview & Goals

Create a `README.md` file for the `sortinator` Go CLI project. The README should be minimal and focused, serving as the
entry point for users and demonstrating skills in using local AI agents for development.

### Scope

- **In Scope**: A single `README.md` file at the project root.
- **Out of Scope**: No changes to source code, no new files besides the README.

### User Stories

- As a visitor to the repository, I can quickly understand what Sortinator does.
- As a developer, I can build the project from scratch using the provided instructions.
- As a user, I can see concrete examples of how to run the CLI and understand all available flags.

### Functional Requirements

- Brief project description mentioning its purpose (testing local inference with llama.cpp).
- Quick build instructions (`go build` or `go install`).
- Usage examples showing the CLI command, flags, and expected output.
- A flag reference table (`-f`, `-r`, `-u`, `-v`).
- Minimal prose — no fluff, no architecture deep-dives.

# Delivery Steps

### Step 1: Create README.md

Create a minimal README.md at the project root with the following sections:

- **Title & short description**: "Sortinator — sorts IP addresses from YAML files", mentioning its use for testing local
  inference (llama.cpp).
- **Build**: `go build` and/or `go install` commands.
- **Usage**: A quick example showing `sortinator test.yaml` with sample input/output.
- **Flags**: A table documenting `-f`/`--file`, `-r`/`--reverse`, `-u`/`--unique`, `-v`/`--verbose`.
- Brief, clean Markdown formatting — no badges, no extra sections.