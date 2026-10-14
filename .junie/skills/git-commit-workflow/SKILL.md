---
name: git-commit-workflow
description: Standardized Git workflow for making changes, linting, testing, and committing with a strict commit message format
---

# Git Commit Workflow

Rules:

1. Always run `git status` first to understand the current state before taking any action.
2. Based on the changelog, prepare a commit message according to the skill.
3. Never push changes, interactively rebase, reset the --hard option, or delete branches unless explicitly stated.
4. Never commit secrets, credentials, or `.env` files-mark them instead.
5. If your working tree contains unrelated changes, ask permission before committing everything with `git add .`;
   committing specific files is preferable.

## Commit message format

```
<type>(<scope>): <short summary>

<body: what changed and why, 1-2 lines>
```

Rules:

- Types: feat, fix, refactor, test, chore, docs.
- Scope = affected package/module (e.g. validator, sorter).
- Short summary must be no more than 50 characters.
- Body is 1-2 short sentences.
- Avoid overly verbose descriptions or unnecessary details.
- Don't use colon and semicolon in the description.
