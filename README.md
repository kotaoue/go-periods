# go-periods

[![Go](https://github.com/kotaoue/go-periods/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/kotaoue/go-periods/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/kotaoue/go-periods/branch/main/graph/badge.svg)](https://codecov.io/gh/kotaoue/go-periods)
[![Go Report Card](https://goreportcard.com/badge/github.com/kotaoue/go-periods)](https://goreportcard.com/report/github.com/kotaoue/go-periods)
[![License](https://img.shields.io/github/license/kotaoue/go-periods)](https://github.com/kotaoue/go-periods/blob/main/LICENSE)

Package for date period handling in Go.

## Installation

```bash
go get github.com/kotaoue/go-periods/periods
```

## Import

```go
import "github.com/kotaoue/go-periods/periods"
```

## Usage

| Function | Description |
| ---- | ---- |
| `periods.Split(period string) ([]string, error)` | Split a date period string into a slice of date strings |

### Supported formats

| Format | Example | Description |
| ---- | ---- | ---- |
| `YYYYMMDD` | `20220101` | Single day |
| `YYYYMMDD-YYYYMMDD` | `20220101-20220105` | Day range |
| `YYYYMM` | `202201` | Single month |
| `YYYYMM-YYYYMM` | `202201-202203` | Month range |

## Examples

### Single day

```go
package main

import (
	"fmt"

	"github.com/kotaoue/go-periods/periods"
)

func main() {
	period, err := periods.Split("20220101")
	if err != nil {
		panic(err)
	}
	fmt.Println(period)
}
```

Output:

```
[20220101]
```

### Day range

```go
package main

import (
	"fmt"

	"github.com/kotaoue/go-periods/periods"
)

func main() {
	days, err := periods.Split("20220101-20220105")
	if err != nil {
		panic(err)
	}
	fmt.Println(days)
}
```

Output:

```
[20220101 20220102 20220103 20220104 20220105]
```

### Month range

```go
package main

import (
	"fmt"

	"github.com/kotaoue/go-periods/periods"
)

func main() {
	months, err := periods.Split("202201-202203")
	if err != nil {
		panic(err)
	}
	fmt.Println(months)
}
```

Output:

```
[202201 202202 202203]
```