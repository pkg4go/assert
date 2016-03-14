package assert

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
	a.NotEqual("1", "2")
	a.NotEqual([]int{1, 2, 3}, []int{1, 2})
}

func TestReport(t *testing.T) {
	// a := A{t}

	// a.Nil(true)
	// a.NotNil(nil)
	// a.Equal(1, 2)
	// a.NotEqual(1, 1)
}
