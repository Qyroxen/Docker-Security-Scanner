# Docker Security Scanner

![CI](https://github.com/Qyroxen/Docker-Security-Scanner/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Docker-Security-Scanner/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Docker-Security-Scanner?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Docker-Security-Scanner)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Docker-Security-Scanner)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Docker-Security-Scanner?style=social)](https://github.com/Qyroxen/Docker-Security-Scanner/stargazers)

## What is it?

Docker Security Scanner is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Docker-Security-Scanner.git
cd Docker-Security-Scanner
go build -o dockersecurityscanner .

# Run
./dockersecurityscanner --help
```

## CLI Usage

```bash
# Basic usage
./dockersecurityscanner

# With flags
./dockersecurityscanner --verbose --output json

# Get help
./dockersecurityscanner --help
```

## Examples

```bash
# Example 1
./dockersecurityscanner example1

# Example 2
./dockersecurityscanner example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o dockersecurityscanner .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Docker-Security-Scanner/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Docker-Security-Scanner?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Docker-Security-Scanner/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Docker-Security-Scanner?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Docker-Security-Scanner/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Docker-Security-Scanner" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Docker-Security-Scanner/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Docker-Security-Scanner" alt="Pull Requests">
  </a>
</p>
