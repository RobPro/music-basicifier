---
mode: agent
---
You are working under YOLO (auto-approve) mode, so try to continue, only stopping for questions if your own investigation falls short.

Important workflow rules:
- Always create a new branch for every roadmap increment. Do not reuse an existing feature branch or a branch that already has a PR, even if the branch name looks related to the current task.
- Never finish a roadmap increment without an open GitHub PR for the new branch. After pushing the branch, always attempt to create a PR via `gh pr create` (or the equivalent GitHub workflow if `gh` is unavailable). If PR creation is blocked by auth or tooling, stop and report the block instead of silently skipping PR creation.
- Never mark a roadmap item complete in PROJECT.md from a feature branch. Do not change the checkbox to `[x]` until the PR has been merged into main. After merge, update the corresponding `[x]` item in PROJECT.md on main and push that change as a follow-up commit.

1. Read PROJECT.md. Find the first unchecked item under "Roadmap".
2. Assess feasibility as a single PR-sized unit of work. If it's too big
   for one focused session, break it into numbered sub-steps and rewrite
   that roadmap item as nested checkboxes instead of proceeding — then stop
   and report the breakdown for approval.
3. If it's appropriately sized: restate the task, list files to touch,
   flag any new dependency per the dependency policy in
   copilot-instructions.md. Stop and wait for "go".
4. Once approved: treat this as a fresh change and create a new, unique
   branch from latest main. The default branch name should be
   `feature/step-N-<short-name>`, but before using it, check whether that
   branch already exists locally or on origin, and whether GitHub already has
   a PR for it (open or merged). If the name is already taken or already
   used for a PR, append a short suffix such as `-2`, `-3`, or
   `-<YYYYMMDD>` to make the branch name unique. Do not reuse a previously
   merged or already-used branch for a new increment.
5. Implement. Run `make check`; fix and re-run per the failure handling
   policy (max 2 attempts) until green. Commit, push the new branch, and
   create a PR via `gh pr create` with a summary of what changed and
   confirmation that checks pass. If `gh pr create` reports that a PR already
   exists for that branch, reuse it rather than creating a duplicate. Do NOT
   merge.
6. After the PR is merged, update the roadmap item `[x]` in PROJECT.md on
   `main` and push that change. Do not check it off from the feature branch.