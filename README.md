# gotest is a Go library for test

[![Golang](https://img.shields.io/badge/language-Go-brightgreen.svg?style=flat)](https://golang.org)
[![Build Status](https://travis-ci.org/philchia/gotest.svg?branch=master)](https://travis-ci.org/philchia/gotest)
[![Coverage Status](https://coveralls.io/repos/github/philchia/gotest/badge.svg?branch=master)](https://coveralls.io/github/philchia/gotest?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/philchia/gotest)](https://goreportcard.com/report/github.com/philchia/gotest)
[![GoDoc](https://godoc.org/github.com/philchia/gotest?status.svg)](https://godoc.org/github.com/philchia/gotest)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://opensource.org/licenses/MIT)

## Usage

### assert

#### Assert equal

```go
import "github.com/philchia/gotest/assert"

func TestEqual(t *testing.T) {
    assert.Equal(t, 1, 1)
}

```

#### Assert not equal

```go
import "github.com/philchia/gotest/assert"

func TestNotEqual(t *testing.T) {
    assert.NotEqual(t, 1, 2)
}

```

#### Assert nil

```go
import "github.com/philchia/gotest/assert"

func TestNil(t *testing.T) {
    assert.Nil(t, nil)
}

```

#### Assert not nil

```go
import "github.com/philchia/gotest/assert"

func TestNotNil(t *testing.T) {
    assert.NotNil(t, 2)
}

```

#### Assert less than

```go
import "github.com/philchia/gotest/assert"

func TestLessThan(t *testing.T) {
    assert.LessThan(t, 2, 3)
}

```

#### Assert less than or equal

```go
import "github.com/philchia/gotest/assert"

func TestLessThanOrEqual(t *testing.T) {
    assert.LessThanOrEqual(t, 2, 2)
}

```

#### Assert greater than

```go
import "github.com/philchia/gotest/assert"

func TestGreaterThan(t *testing.T) {
    assert.GreaterThan(t, 2, 1)
}

```

#### Assert greater than or equal

```go
import "github.com/philchia/gotest/assert"

func TestGreaterThanOrEqual(t *testing.T) {
    assert.GreaterThanOrEqual(t, 2, 2)
}

```

#### Assert true

```go
import "github.com/philchia/gotest/assert"

func TestTrue(t *testing.T) {
    assert.True(t, 2 > 1)
}

```

#### Assert false

```go
import "github.com/philchia/gotest/assert"

func TestFalse(t *testing.T) {
    assert.False(t, 2 < 1)
}

```

### bench

```go
import "github.com/philchia/gotest/bench"

func BenchFunc(b *testing.B) {
    bench.Benchmark(b, func(){
        log.Println("benchmark")
    })
}

```

### mock

## Todo

## License

gotest code published under MIT license