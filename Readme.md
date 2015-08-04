[![Build status][travis-img]][travis-url]
[![License][license-img]][license-url]
[![GoDoc][doc-img]][doc-url]

### assert

* assert only for testing

```go
import . "github.com/pkg4go/assert"
import "testing"

func TestBasic(t *testing.T) {
  a := A{t}

  a.Nil(nil)

  a.NotNil(0)
  a.NotNil("")
  a.NotNil(123)

  a.Equal(1, 1)
  a.Equal("1", "1")
  a.Equal([]int{1, 2, 3}, []int{1, 2, 3})

  a.NotEqual(1, 2)
  a.NotEqual(1, "1")
  a.NotEqual([]int{1, 2, 3}, []int{1, 2})
}
```

### License
MIT

[travis-img]: https://img.shields.io/travis/pkg4go/assert.svg?style=flat-square
[travis-url]: https://travis-ci.org/pkg4go/assert
[license-img]: https://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[doc-img]: https://img.shields.io/badge/GoDoc-reference-blue.svg?style=flat-square
[doc-url]: http://godoc.org/github.com/pkg4go/assert
