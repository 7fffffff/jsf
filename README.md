# jsf

This package provides an implementation of the 64-bit, 3-rotation
variant of Bob Jenkins' unnamed "small noncryptographic PRNG"

For more information and the original implementation in C:  
<https://burtleburtle.net/bob/rand/smallprng.html>

## Example

jsf.JSF64 implements the math/rand.Source64 interface

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
	
	"github.com/7fffffff/jsf"
)

func main() {
	// Always seed with something, or you'll get nothing but zeroes
	src := jsf.New(time.Now().UnixNano())
	rng := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(rng.Int63())
	}
}
```

## 32-bit

See the jsf32 subdirectory for an implementation of the 32-bit version.
On some CPUs it might be a bit faster for generating 32-bit values, if
called directly

```
cpu: AMD FX(tm)-6120 Six-Core Processor
pkg: github.com/7fffffff/jsf
BenchmarkUint32-6       184545068                6.793 ns/op
BenchmarkUint64-6       180301332                6.394 ns/op
pkg: github.com/7fffffff/jsf/jsf32
BenchmarkUint32-6       258255520                4.663 ns/op
BenchmarkUint64-6       100000000                10.71 ns/op
```

```
cpu: AMD Ryzen 5 2600X Six-Core Processor
pkg: github.com/7fffffff/jsf
BenchmarkUint32-12      410623514                2.946 ns/op
BenchmarkUint64-12      412275786                2.885 ns/op
pkg: github.com/7fffffff/jsf/jsf32
BenchmarkUint32-12      424884342                2.795 ns/op
BenchmarkUint64-12      291792106                4.098 ns/op
```

The math/rand Source and Source64 interfaces only use Int63() and
Uint64(), so the 64-bit version seems preferable for that use

## API Documentation

<https://pkg.go.dev/github.com/7fffffff/jsf>

## License

This software is licensed under the Blue Oak Model License 1.0.0:  
<https://blueoakcouncil.org/license/1.0.0>