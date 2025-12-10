# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Terraform provider for Solace Software Event Broker, built using the Terraform Plugin Framework. The provider enables infrastructure-as-code management of Solace brokers via the SEMP (Solace Element Management Protocol) API.

**Minimum supported broker version**: 10.4

**Important limitation**: Generally available for production services contained within a Message VPN. Resources outside Message VPNs are not supported in production.

## Essential Commands

### Building and Installing
- `make install` - Build and install provider to `${GOBIN}` (for local development)
- `make build` - Build binary to current directory
- `make dep` - Tidy Go dependencies
- `make clean` - Remove build artifacts

### Testing
- `make test` - Run unit tests
- `make test-coverage` - Run tests with HTML coverage report (generates `reports/cover.html`)
- `make testacc` - Run acceptance tests (requires `TF_ACC=1`, takes ~2 hours)
- Acceptance tests automatically start a Docker test broker using testcontainers-go

### Code Quality
- `make fmt` - Format code with gofmt
- `make vet` - Run go vet

### Documentation
- `make generate-docs` - Generate Terraform provider documentation

### Development Broker
- `make newbroker` - Start a local Solace broker in Docker (default: `tag=10.25.0.24`)
- Example: `make newbroker tag=10.26.0.25`
- Runs on ports: 8080 (SEMP), 55555, 8008, 1883, etc.
- Default credentials: admin/admin

## Code Architecture

### Generated Code
**Critical**: Most provider code is auto-generated from SEMP API specs. Generated files live in `internal/broker/generated/`. Each Solace object type (msg_vpn, queue, etc.) has its own generated file.

To regenerate code:
```bash
make generate-code
```
This clones `broker-terraform-code-generator`, generates code from SEMP spec in `ci/swagger_spec/`, and updates `internal/broker/generated/`.

### Core Package Structure
- `internal/broker/` - Core provider logic
  - `provider.go` - Provider implementation
  - `resource.go` - Generic resource CRUD operations
  - `datasource.go` - Data source implementations
  - `schema.go` - Schema utilities
  - `generated/` - Auto-generated resource/datasource definitions
  - `testacc/` - Acceptance test framework
- `internal/semp/` - SEMP API client wrapper
- `cmd/` - CLI commands (generate, version, completion)
  - `generator/` - Config generation from existing broker
  - `client/` - SEMP client for CLI

### Provider Binary Dual Purpose
The provider binary serves two purposes:
1. **Standard Terraform provider** - Started by Terraform CLI
2. **Standalone CLI tool** - Run `terraform-provider-solacebroker generate` to export existing broker configs to HCL

### Config Generator Feature
The provider can reverse-engineer Terraform configs from a running broker:
```bash
terraform-provider-solacebroker generate \
  --url=http://localhost:8080 \
  solacebroker_msg_vpn.myvpn test vpn-config.tf
```
This generates HCL for the `test` message VPN and all child objects.

### Resource Implementation Pattern
Resources use a generic framework in `resource.go` that:
1. Uses generated entity definitions from `generated/`
2. Handles CRUD via SEMP API with automatic retry logic
3. Converts between Terraform state and SEMP API objects
4. Manages resource hierarchies (parent/child relationships)

### Testing Strategy
- **Unit tests**: Standard Go tests (`_test.go` files)
- **Acceptance tests**: In `internal/broker/testacc/`
  - Require `TF_ACC=1` environment variable
  - Automatically spin up Solace broker in Docker via testcontainers
  - Test real Terraform operations against broker
  - Each generated resource has acceptance tests

## Development Workflow

### Using Local Provider Build
After `make install`, configure Terraform to use local build:

Create/update `~/.terraformrc`:
```hcl
provider_installation {
  dev_overrides {
    "registry.terraform.io/solaceproducts/solacebroker" = "${GOBIN}"
  }
  direct {}
}
```

Or use `TF_CLI_CONFIG_FILE` environment variable.

### Environment Variables for Provider
- `SOLACEBROKER_URL` - Broker URL (e.g., `http://localhost:8080`)
- `SOLACEBROKER_USERNAME` - Admin username
- `SOLACEBROKER_PASSWORD` - Admin password
- `SOLACEBROKER_BEARER_TOKEN` - Alternative to username/password
- `SOLACEBROKER_INSECURE_SKIP_VERIFY` - Skip SSL verification
- `SOLACEBROKER_SKIP_API_CHECK` - Skip platform validation
- `SOLACEBROKER_REGISTRY_OVERRIDE` - Override registry address

### Test Environment Variables
- `TF_ACC=1` - Enable acceptance tests

## Important Architectural Details

### SEMP API Version Handling
The provider is built for a specific SEMP version and platform (software vs appliance). Version checking occurs in:
- `main.go` - Validates platform matches expected
- `version.go` - Defines provider version
- `internal/broker/sempversion.go` - SEMP version utilities

### Platform Check
Provider validates it's talking to a software broker (not appliance) unless `SOLACEBROKER_SKIP_API_CHECK=true`.

### Hierarchical Resources
Resources have parent/child relationships (e.g., queue subscriptions belong to queues). The framework in `resource.go` handles:
- Building import identifiers with parent IDs
- Cascading creates/deletes
- Proper Terraform state management

### Write-Only Attributes
Some broker attributes are write-only (e.g., passwords). The generator handles these specially:
- Coupled with non-write-only attributes when possible
- Generated as variable references in config generator
- May require manual specification in generated configs

## Common Development Scenarios

### Adding Support for New Broker Feature
1. Update SEMP spec in `ci/swagger_spec/`
2. Run `make generate-code`
3. Run `make generate-docs`
4. Test with `make testacc`

### Debugging Provider
Set `-debug` flag when running provider directly:
```bash
terraform-provider-solacebroker -debug
```
Or set `SOLACEBROKER_DEBUG_RUN` environment variable.

### Running Single Test
```bash
go test -run TestAccResourceQueue ./internal/broker/testacc/
```

### Testing Config Generator
```bash
# Build provider
make install

# Start test broker
make newbroker

# Generate config
~/go/bin/terraform-provider-solacebroker generate \
  --url=http://localhost:8080 \
  --username=admin --password=admin \
  solacebroker_msg_vpn.default default default.tf
```

## CI/CD Pipeline
GitHub Actions workflows in `.github/workflows/`:
- `provider-test-pipeline.yml` - Main test pipeline
- `provider-acceptance-test.yml` - Acceptance tests
- `core-pipeline-main-branch-only.yml` - Main branch validation
- `daily-sanity-main.yml` - Daily sanity checks
- `provider-release.yml` - Release automation

## Key Dependencies
- Terraform Plugin Framework v1.9.0
- Terraform Plugin Testing v1.7.0
- testcontainers-go v0.30.0 (for acceptance tests)
- hashicorp/go-retryablehttp v0.7.7
- spf13/cobra v1.8.1 (for CLI)
