# Kun Foundation Library - Project Rules

Reply in Chinese; use Conventional Commits with Chinese or English descriptions (e.g., `feat(log): ...`, `fix(maps): ...`).

## Overview

Kun is the foundational utility and common infrastructure library for the Yao application engine ecosystem. It provides unified exception handling, zero-allocation type casting, structured contextual logging, dynamic map/record manipulation, string utilities, and gRPC communication helpers.

- **Ecosystem Role**: Root foundation library imported by `gou`, `xun`, and `yao`. Zero external dependencies on sister repos.

## Tech Stack

- **Core Runtime**: Go 1.24+ (Toolchain 1.25.5)
- **Logging Backend**: Logrus (`github.com/sirupsen/logrus`) + ColorJSON formatting
- **RPC & Plugins**: HashiCorp `go-plugin`, `google.golang.org/grpc`, `google.golang.org/protobuf`
- **Testing**: `github.com/stretchr/testify` (assert, require)

## Common Commands & Fast Feedback Loop

```bash
# Fast Feedback Inner Loop (<2s)
go test -v -run TestAny ./any/...       # Test safe type casting
go test -v -run TestMaps ./maps/...     # Test record/map access
go test -v -run TestException ./exception/... # Test exception throw/catch
go test -race ./any/...                 # Verify race safety on package
make vet                                # Static analysis (go vet)

# Full Verification
make test                               # Run full test suite across active packages
make fmt                                # Format Go files with gofmt -s
make lint                               # Run golint check
```

## Navigation & Key Entrypoints

| Subsystem / Utility | Primary Entry File / Directory |
| :--- | :--- |
| **Exception Throw & Catch Semantics** | `exception/exception.go` |
| **Contextual Structured Logging** | `log/log.go` |
| **Zero-Panic Type Casting (`any.Of`)** | `any/any.go`, `any/conv.go` |
| **Record & Map Dot-Notation Access** | `maps/maps.go`, `maps/record.go` |
| **String Helpers & Token Generators** | `str/str.go` |
| **High-Precision Numeric Conversion** | `num/num.go` |
| **Datetime & ISO8601 Helpers** | `day/day.go` |
| **gRPC Client/Server Plugin Wrapper** | `grpc/` |

## Directory Structure

```
kun/
├── exception/         # Exception throwing, catching, error code & stack unwinding
├── log/               # Structured contextual logging (log.With, log.Info, log.Error)
├── any/               # High-performance zero-panic type casting (Of, ToInt, ToString, etc.)
├── maps/              # Dynamic Map & Record utilities, dot-notation access (Get/Set/Has)
├── str/               # String helpers, case conversions (Snake, Camel, Kebab) & token generators
├── num/               # High-precision numeric parsing, rounding & safety conversions
├── day/               # Datetime formatting, calculation and ISO8601 parsing
├── grpc/              # gRPC client/server wrapper for HashiCorp go-plugin integration
├── interfaces/        # Shared core interface definitions
├── share/             # Common constants, data structures & environment detection
└── utils/             # General utility functions and reflection helpers
```

## System Invariants (DO NOT BREAK)

1. **Zero-Panic Guarantee**: Methods in `any/` and `maps/` must NEVER allow standard Go `panic()` to escape. Always return safe default zero-values on type mismatch.
2. **API Backward Compatibility**: As the root dependency of `gou`, `xun`, and `yao`, public API signatures in `kun` MUST remain backward-compatible at all costs.
3. **Structured Stack Unwinding**: `exception.Throw()` uses structured panics that MUST be intercepted cleanly via `defer exception.Catch(...)`.

## Footguns & Anti-Patterns (DO NOT)

- **DO NOT** use standard Go `panic(err)`. Always use `exception.New(...).Throw()` or return typed errors.
- **DO NOT** use reflection (`reflect.TypeOf` / `reflect.ValueOf`) on hot paths in `any/` and `maps/`. Always prefer explicit type switches (`switch v := val.(type)`).
- **DO NOT** introduce dependencies from `gou`, `xun`, or `yao` into `kun` (strictly circular-dependency illegal).
