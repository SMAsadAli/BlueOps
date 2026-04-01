# Contributing

## Scope

BlueOps is being built in small, guide-aligned milestones. Keep changes focused, easy to review, and limited to the phase or step being worked on.

## Local Setup

- Use Go 1.22 or newer.
- Run `make build` to confirm the CLI compiles.
- Run `make test` before opening a pull request.
- Run `gofmt` on changed Go files.

## Change Guidelines

- Prefer small PR-sized changes over bundled feature drops.
- Keep command wiring, business logic, and output concerns separated.
- Default to safe behavior for any operational command.
- Add or update tests when behavior changes.
- Update documentation when the command surface or configuration model changes.

## Pull Requests

- Write a clear title that matches the change scope.
- Describe what changed, why it changed, and how it was verified.
- Call out follow-up work explicitly instead of bundling it into the same PR.

## Issues

If you find a bug or design gap, open an issue with:

- a short description of the problem
- reproduction steps
- expected behavior
- actual behavior
- relevant environment details
