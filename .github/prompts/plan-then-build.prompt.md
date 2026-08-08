---
mode: agent
---
You are working under YOLO (auto-approve) mode, so treat the plan step
below as the only checkpoint before you start making changes.

Task: ${input:task}

Do the following, in order:

1. **Restate** the task in your own words in 1-3 sentences.
2. **File impact list** — every file you expect to create, modify, or delete.
3. **Dependency check** — state explicitly: "No new dependencies needed"
   OR name the dependency and justify why stdlib/x/sys/windows/custom code
   is insufficient.
4. **Approach** — a short paragraph or bullet list of the implementation
   plan, including any Windows API surfaces touched.
5. Stop here and output exactly: `Ready to proceed — reply "go" to continue.`

Do not call any tools (read-only exploration aside) until the user replies.
Once approved, implement the plan, then run `make check` and report the
result before declaring the task complete.