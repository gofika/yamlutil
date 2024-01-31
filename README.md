[![codecov](https://codecov.io/gh/gofika/yamlutil/branch/main/graph/badge.svg)](https://codecov.io/gh/gofika/yamlutil)
[![Build Status](https://github.com/gofika/yamlutil/workflows/build/badge.svg)](https://github.com/gofika/yamlutil)
[![go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/gofika/yamlutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofika/yamlutil)](https://goreportcard.com/report/github.com/gofika/yamlutil)
[![Licenses](https://img.shields.io/github/license/gofika/yamlutil)](LICENSE)

# yamlutil

golang yaml utils for common use


## Basic Usage

### Installation

To get the package, execute:

```bash
go get github.com/gofika/yamlutil
```

### Example

```go
package main

import (
	"fmt"

	"github.com/gofika/yamlutil"
)

type Foo struct {
	Name  string `yaml:"name"`
	Value int    `yaml:"value"`
}

func main() {
	foo := &Foo{
		Name:  "Jason",
		Value: 100,
	}
	name := "foo.yaml"
	// write struct to file
	err := yamlutil.WriteFile(name, foo)
	if err != nil {
		fmt.Printf("WriteFile failed. err: %s\n", err.Error())
		return
	}
	// read struct from file
	bar, err = yamlutil.ReadFile[Foo](name)
	if err != nil {
		fmt.Printf("ReadFile failed. err: %s\n", err.Error())
		return
	}
	fmt.Printf("bar.Name: %s\n", bar.Name)
	fmt.Printf("bar.Value: %d\n", bar.Value)
}
```