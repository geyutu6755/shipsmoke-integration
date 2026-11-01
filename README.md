# shipsmoke-integration

Independent end-to-end integration test for the
[shipsmoke](https://github.com/geyutu6755/shipsmoke) GitHub Action.

This repository is deliberately a *separate project*: it builds its own tiny
CLI (`teapot`), packages it like a release, and verifies the package with the
shipsmoke action consumed exactly as an outside user would:

```yaml
- uses: geyutu6755/shipsmoke@v1
  with:
    config: .shipsmoke.yml
    report: shipsmoke-report.json
```

The workflow runs on ubuntu-24.04, macos-15 (arm64) and windows-2025 and
asserts that the action's `exit-code` output is `0` and that the JSON report
is produced with `overall == PASS`. A second job verifies the refusal path:
a tampered artifact must fail the action.
