# Releasing

This project publishes Go modules following Semantic Versioning.
Releases are cut from tags on the `main` branch and automated workflows.

## Release Checklist

- Next SemVer tag is determined (`vMAJOR.MINOR.PATCH`).
- A tag is released on main
  - ```git tag -s vX.Y.Z```
  - ```git push origin vX.Y.Z```

[- Automation publishes artifacts]: #
[- Pushing the tag triggers `.github/workflows/wf-release.yaml`.]: #
[- SLSA: The workflow builds a source archive, generates a CycloneDX SBOM, records checksums, uploads an SBOM attestation and provenance `.intoto.jsonl` assets.]: #
[- If the automated release does not include human-readable notes, edit the GitHub release, paste the changelog notes, and save.]: #

## Emergency Releases

For high-severity security issues, coordinate privately via the process in [.github/SECURITY.md](../.github/SECURITY.md). Patch branches should include only the minimal changes required to resolve the issue.
