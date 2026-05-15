# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

Go module `go.openviz.dev/installer` — Helm charts and supporting tooling for installing the OpenViz stack. It also exposes a Go API package describing chart values so that other OpenViz components can consume strongly typed installation parameters.

Charts shipped:
- `charts/grafana-operator` — the OpenViz Grafana operator (this is the `grafana-tools` binary).
- `charts/monitoring-operator` — Prometheus/Alertmanager stack helpers.
- `charts/trickster` — Trickster proxy install.
- `charts/kube-grafana-dashboards` — pre-built dashboard bundle.

## Architecture

- `charts/` — one subdirectory per Helm chart. Each has `Chart.yaml`, `values.yaml`, `templates/`, plus generated artifacts `doc.yaml`, `README.md`, and (for charts with a Go-typed schema) `values.openapiv3_schema.yaml`.
- `apis/installer/v1alpha1/` — Go types backing the chart values (single API group `installer:v1alpha1`). Used for OpenAPI/values-schema generation and as a typed surface for downstream programs.
  - `register.go`, `install/`, `fuzzer/` — standard k8s scheme registration and round-trip fuzz helpers.
- `crds/` — top-level CRD manifests; chart-specific CRDs live under each chart's `crds/`.
- `catalog/` — image catalog driven by `kmodules.xyz/image-packer`. `imagelist.yaml` is the source of truth; the script outputs (`copy-images.sh`, etc.) and `catalog/README.md` (CVE report) are auto-generated. Do not hand-edit anything in `catalog/` except `imagelist.yaml`.
- `hack/scripts/` — codegen/release helpers: `update-catalog.sh`, `update-chart-dependencies.sh`, `import-crds.sh`, `ct.sh`, `open-pr.sh`, `update-release-tracker.sh`.
- `tests/` — chart-testing (`ct`) configuration.
- `lintconf.yaml` — YAML lint config used by `make ct`.
- `vendor/` — vendored Go deps.

## Common commands

All Make targets run inside the `ghcr.io/appscode/golang-dev` Docker image — Docker must be running.

- `make gen` — regenerate everything: `codegen manifests` (clientset, OpenAPI, CRDs, values schemas, chart docs).
- `make codegen` — clientset only.
- `make manifests` — `gen-crds gen-values-schema gen-chart-doc`.
- `make gen-values-schema` — regenerate `values.openapiv3_schema.yaml` from `apis/installer/v1alpha1`.
- `make gen-chart-doc` — regenerate per-chart `README.md`.
- `make update-charts` — refresh chart-level metadata (one target per chart subdir).
- `make fmt` — gofmt + goimports.
- `make lint` — golangci-lint.
- `make unit-tests` / `make test` — Go unit tests.
- `make ct` — `chart-testing` lint+test against the charts.
- `make verify` — `verify-modules` (re-run `gen fmt`, ensure `go mod tidy && go mod vendor` is clean).
- `make add-license` / `make check-license` — manage license headers.

Auxiliary helpers (invoked outside Make):

- `./hack/scripts/update-catalog.sh` — regenerate `catalog/` from `imagelist.yaml` via image-packer.
- `./hack/scripts/import-crds.sh` — pull CRDs from dependent repos into the chart `crds/` dirs.
- `./hack/scripts/update-chart-dependencies.sh` — refresh `Chart.lock` / subchart pins.

Run a single Go test (requires a local Go toolchain):

```
go test ./apis/installer/v1alpha1/... -run TestName -v
```

## Conventions

- Module path is `go.openviz.dev/installer` (vanity URL); imports must use that, not the GitHub URL (`open-viz/installer`).
- Edit `apis/installer/v1alpha1/*_types.go` to change a chart's values surface, then run `make gen` so the generated clientset, `values.openapiv3_schema.yaml`, and per-chart `README.md` stay in sync. Do not hand-edit `zz_generated.*.go`, generated chart `README.md` files, `values.openapiv3_schema.yaml`, or anything under `catalog/` except `imagelist.yaml`.
- License: Apache-2.0; use `make add-license` to apply headers to new files.
- Sign off commits (`git commit -s`); contributions follow the DCO (`DCO` file).
- Vendor directory is checked in — `go mod tidy && go mod vendor` must leave the tree clean (enforced by `verify-modules`).
- When adding a new chart: create `charts/<name>/`, add a values-schema target by following the existing chart layout, and add it to `imagelist.yaml` if it ships images.
