# Contributing

Thanks for investing time in improving this module! 

## 1. Before You Start

- Review the [Code of Conduct](CODE_OF_CONDUCT.md). Participating implies acceptance of its terms.
- Search existing [issues](https://github.com/bytemare/hash/issues) and pull requests to avoid duplicating work. For substantial or breaking changes, open an issue first so we can agree on scope.
- Familiarise yourself with the architecture and testing expectations.

## 2. Development Environment

- One among the three latest Go versions (CI runs the current stable toolchain and the two previous releases).
- `go`, `git`, `make`, and a POSIX-compatible shell.
- Optional tools used by CI (installed automatically in workflows but helpful locally): `golangci-lint`, `govulncheck`, and other tooling you can install with `make update-linters`.

## 3. Workflow and Branching

1. Fork the repository and create topic branches from `main` (for example `feat/blake3-support`, `docs/security-refresh`).
2. Keep changes focused. Separate refactors, dependency bumps, and feature work into distinct pull requests.
3. Reference related issues in your branch description or pull request.

## 4. Commit Standards

- Follow [Conventional Commits](https://www.conventionalcommits.org/) for clear history (`feat:`, `fix:`, `docs:`, `test:`, `chore:` and so on).
- Every commit must include a `Signed-off-by` trailer to satisfy the [Developer Certificate of Origin](https://developercertificate.org/). Use `git commit -s` to add it automatically.
- Every commit must be cryptographically signed.
- Commit only what you have built and tested locally. Avoid large unrelated changes in a single commit.

## 5. Quality Checks

1. Run fast local checks before pushing:
   ```bash
   make -C .github test
   ```
2. Mirror the CI matrix when possible:
   ```bash
   make -C .github lint vulncheck test cover fuzz
   ```
   These targets run `golangci-lint`, `govulncheck`, unit tests, race coverage, and fuzzing seeds.
3. Ensure `go mod tidy` produces no diff and that coverage does not regress meaningfully. If coverage drops, explain why in the pull request.
4. Update documentation when behaviour or APIs change. Architecture or security changes should be reflected in the relevant `docs/` files.

## 6. Opening a Pull Request

1. Push your branch and open a PR against `main`.
2. Fill out the pull request template, including the commands you ran and any follow-up tasks.
3. Keep the description focused on why the change is necessary and what risks were considered. Link issues or discussions for additional context.
4. If the change affects documentation, link to the updated files in the PR body so reviewers can verify rendering quickly.

## 7. Review Expectations

- Expect at least one maintainer review. Response times are usually within a few business days. Comment if you need a quicker turnaround.
- Be responsive to feedback. If discussion stalls, summarise remaining concerns so the maintainer can make a decision.
- Maintainers may close stale PRs after reasonable attempts to coordinate.

## 8. Issue Guidance

- Include reproduction steps, expected versus actual behaviour, and environment details (Go version, OS/arch).
- For vulnerabilities, use the private GitHub Security Advisory form rather than public issues.
- Feature requests should describe the use case and, when possible, sketch the desired API.

## 9. Releases and Post-Merge Tasks

- Maintainers follow the process in [docs/releasing.md](../docs/releasing.md). Contributors assisting with release notes should provide changelog entries and highlight migration steps.
- After a change lands, watch for CI status and respond quickly if regressions are reported by downstream users.

## 10. Further Reading

- Governance model: [docs/governance.md](../docs/governance.md)
- Roadmap and open initiatives: [docs/roadmap.md](../docs/roadmap.md)

Thank you for helping keep the project reliable and secure!
