[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/gosv)](https://pkg.go.dev/github.com/s0rg/gosv)
[![Build](https://github.com/s0rg/gosv/workflows/ci/badge.svg)](https://github.com/s0rg/gosv/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/gosv)](https://goreportcard.com/report/github.com/s0rg/gosv)
[![Maintainability](https://api.codeclimate.com/v1/badges/e1c002df2b4571e01537/maintainability)](https://codeclimate.com/github/s0rg/gosv/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/e1c002df2b4571e01537/test_coverage)](https://codeclimate.com/github/s0rg/gosv/test_coverage)
[![License](https://img.shields.io/badge/license-MIT%20License-blue.svg)](https://github.com/s0rg/gosv/blob/main/LICENSE)

# gosv

golang rule-based string validator

# usage
```go
import "github.com/s0rg/gosv"

var MyRules = []gosv.Rule{
    gosv.MinLen(8),
    gosv.MaxLen(64),
    gosv.MinLowers(1),
    gosv.MinUppers(1),
    gosv.MinNumbers(1),
    gosv.MaxDuplicates(0.2),  // limit duplicates to 20%
    gosv.MaxSequencies(0.1),  // limit sequencies to 10% (i.e. 'abc', '123')
    gosv.MaxEntropyDiff(0.3), // limit entropy diff between ideal and current entropy values
}

func MyPasswordValidator(input string) error {
    return gosv.Validate(input, MyRules...)
}
```

# license

MIT
