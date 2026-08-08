---
mode: agent
---
You are working under YOLO (auto-approve) mode, so try to continue, only stopping for questions if your own investigation falls short.

1. Read PROJECT.md. Find the first unchecked item under "Roadmap".
2. Assess feasibility as a single PR-sized unit of work. If it's too big
   for one focused session, break it into numbered sub-steps and rewrite
   that roadmap item as nested checkboxes instead of proceeding — then stop
   and report the breakdown for approval.
3. If it's appropriately sized: restate the task, list files to touch,
   flag any new dependency per the dependency policy in
   copilot-instructions.md. Stop and wait for "go".
4. Once approved: create branch `feature/step-N-<short-name>` from latest
   main. Implement. Run `make check`; fix and re-run per the failure
   handling policy (max 2 attempts) until green.
5. Commit, push the branch, open a PR via `gh pr create` with a summary of
   what changed and confirmation that checks pass. Do NOT merge.
6. Mark the roadmap item `[x]` in PROJECT.md on `main` only after you tell
   me it's merged — do not check it off from the feature branch.