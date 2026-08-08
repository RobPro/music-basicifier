# AI Agent Rules for this project

This is a Windows desktop application written in Go. These rules apply to
every agent interaction, YOLO mode or not.

## Planning requirement (read this first)
Before editing or creating any file, for any non-trivial task:
1. Restate the task in one or two sentences.
2. List every file you expect to create or modify.
3. If the task might require a new dependency, name it and justify why
   the standard library, `golang.org/x/sys/windows`, or a small amount of
   custom Go code cannot do the job instead.
4. Stop and wait for explicit confirmation ("go" / "approved") before
   making any tool calls that write files or run commands.

Only skip this for genuinely trivial edits (typo fixes, renaming a
variable, adding a comment).

## Dependency policy — be conservative
- The ONLY approved third-party dependency for GUI work is `fyne.io/fyne/v2`,
  already pinned in go.mod. Do not add a second GUI framework.
- Do NOT run `go get -u`, `go get <pkg>@latest`, or `go mod tidy` in a way
  that changes existing version pins, unless the task explicitly says to
  upgrade a specific package (e.g. "bump fyne to fix CVE-XXXX").
- Do NOT add any new third-party module without first stating, in the plan
  step above, why stdlib or `x/sys/windows` can't do it.
- Prefer writing small amounts of custom Go over importing a library for
  something under ~200 lines of straightforward logic (e.g. simple string
  parsing, basic retry logic, small data structures).
- Any change to go.mod or go.sum must be called out explicitly in your
  summary of what you did — never a silent side effect of another task.
- Run with `GOFLAGS=-mod=readonly` assumed; if a command fails because of
  this, stop and report it rather than removing the flag.

## Windows-specific
- Use `golang.org/x/sys/windows` for raw Win32 API access when needed
  (already an indirect dependency via Fyne, safe to use directly).
- Do not shell out to PowerShell/cmd for things Go can do natively
  (file I/O, registry access via x/sys/windows/registry, etc.).

## Quality gate — run before declaring any task done
Always run `make check` (or `.\check.ps1` on native Windows shells) after
making changes. If it fails, fix the issue and re-run before reporting
completion. Do not report a task as done with a failing build, vet, or
test.

## Failure handling
If `make check` fails, attempt at most 2 fixes. If still failing after 2
attempts, STOP and report the failure with full output rather than trying
further changes. Do not proceed to additional features while checks are red.

## Style
- Standard Go formatting (`gofmt`/`goimports`) is mandatory.
- Prefer explicit error handling over panics; no swallowed errors.
- Package-level comments and exported-symbol doc comments required.

## Project context
Read `PROJECT.md` at the repo root before planning any feature. It defines
scope, architecture decisions already made, and out-of-scope items — do
not revisit those decisions without being asked.