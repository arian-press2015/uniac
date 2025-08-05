# UnIaC

UnIaC is a unified `Infrastructure as Code` tool that converts YAML configurations into existing `IaC` tools like `Terraform HCL`, with a focus on simplicity and provider-agnostic infrastructure management.

The goal is to have a unified language for infrastructure, to prevent confusions, simplify transitions, and setting a common language.

## Getting Started

### Prerequisites
- Go 1.24 or later

### Installation

1. Clone the repository:

```bash
git clone https://github.com/arian-press2015/uniac.git
cd uniac
```

2. Build the CLI:

```bash
make build
```

### Usage

Run the uniac CLI to get detailed help:

```bash
./uniac
```

### Development

- Structure: uses a go repo with `cmd/`, `pkg/` and `internal/` directories.
- Core package: `pkg/core/world.go` defines the world model.
- CLI: `cmd/uniac/main.go` handles execution.

### Roadmap

- Add HCL transpilation for providers
- Implement local working environment
- Add cost optimization (`uniac-cost`) and mappings (`uniac-mappings`).

### Contributing

Contributions are welcomed. open issues or PRs on Github

### License

UnIaC is licensed under the GNU General Public License version 2 (GPLv2). See the LICENSE file for details.

