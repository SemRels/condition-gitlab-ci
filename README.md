# condition-gitlab-ci

GitLab CI condition plugin for Semantic Release.

Validates GitLab CI runtime conditions before a Semantic Release is executed.

## Documentation

- Docs (coming soon): <https://github.com/SemRels/semrel/tree/main/docs/plugins/condition-gitlab-ci>
- Template source: <https://github.com/SemRels/plugin-template>

## Repository Layout

`	ext
cmd/plugin/              Plugin entry point
internal/plugin/         Business logic scaffold
internal/grpc/           gRPC transport scaffold
proto/v1                 Symlink to the SemRel protobuf contract
.github/workflows/       CI, release, and security automation
`

## Development

`ash
go build ./cmd/plugin
go test ./...
`

## Configuration Example

`yaml
plugins:
  - name: condition-gitlab-ci
    type: condition
    config:
      require_branch: main
      require_pipeline_source: push
      require_protected_ref: true
`

## Status

This repository is bootstrapped from SemRels/plugin-template and is ready for implementation.