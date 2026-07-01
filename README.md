# condition-gitlab-ci

Allows releases only when semrel is running inside GitLab CI.

This plugin is distributed as the standalone Go binary `semrel-plugin-condition-gitlab-ci`. Semrel executes the binary as a subprocess, provides plugin configuration through `SEMREL_PLUGIN_*` environment variables, provides release context through `SEMREL_*` environment variables, reads standard output, and treats exit code `0` as success and any non-zero exit code as failure. Install the binary in `~/.semrel/plugins/` or anywhere on your `$PATH`.

## Installation

### Binary

```bash
go install github.com/SemRels/condition-gitlab-ci/cmd/plugin@latest
```

### Docker

Pre-built, multi-platform images (linux/amd64, linux/arm64) are published to the GitHub Container Registry on every release:

```bash
docker pull ghcr.io/semrels/condition-gitlab-ci:latest
```

Images are signed with [cosign](https://github.com/sigstore/cosign) and include a full SBOM attestation. Verify the signature:

```bash
cosign verify ghcr.io/semrels/condition-gitlab-ci:latest \
  --certificate-identity-regexp 'https://github.com/SemRels/condition-gitlab-ci/.github/workflows/release.yml.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```


## Configuration

```yaml
plugins:
  - name: condition-gitlab-ci
    path: ~/.semrel/plugins/semrel-plugin-condition-gitlab-ci
    env:
      {}
```

## `SEMREL_PLUGIN_*` variables

| Name | Required | Description | Default |
| --- | --- | --- | --- |
| _None_ | - | This plugin does not require any `SEMREL_PLUGIN_*` variables. It relies on CI-provided environment state. | - |

## `SEMREL_*` release context used

This plugin does not consume any `SEMREL_*` release context variables directly.

## Example behavior

The plugin checks the CI environment and succeeds when `GITLAB_CI=true`. Outside GitLab CI it exits non-zero to stop the release.

## GitLab CI component

In addition to the condition plugin binary above, this project ships a ready-to-use
[GitLab CI/CD component](templates/release.yml) that wraps semrel in a single `release`
job:

```yaml
include:
  - remote: 'https://raw.githubusercontent.com/SemRels/condition-gitlab-ci/main/templates/release.yml'
```

Set a masked/protected `SEMREL_TOKEN` CI/CD variable, then push to your default branch.
Supported inputs (`dry_run`, `config`, `stage`, `semrel_version`) are documented at the
top of [`templates/release.yml`](templates/release.yml). Note: formal listing in the
[GitLab CI/CD Catalog](https://docs.gitlab.com/ee/ci/components/) requires this project
to also exist as a GitLab-hosted project with the Catalog feature enabled — the `remote:`
include above works without that and needs no GitLab account on your side.

## License

Apache-2.0
