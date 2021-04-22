package gotest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// Assert
//
// The assert package provides some helpful methods that allow you to write better test code in Go.
// Prints friendly, easy to read failure descriptions
// Allows for very readable code
// Optionally annotate each assertion with a message
//
// Notes:
//
// - Using the assert package, we can get read of the deep equal call and
//   use a helper to "assert" that testing is equal.
// - Assert has many helper functions inclduing
// 		- assert.NotNil assert.Nil (asserts that an error is nil or not nil)
// 		- assert.JSONEq (good for testing http responses)
// 		- assert.True assert.False (asserts that a variable is true or false)
func TestSplit_Assert(t *testing.T) {
	got := Split("a/b/c", "/")
	want := []string{"a", "b", "c"}
	assert.Equal(t, want, got)
}

// Suite
//
// The suite package provides functionality that you might be used to from more common object
// oriented languages. With it, you can build a testing suite as a struct, build
// setup/teardown methods and testing methods on your struct, and run them with
// 'go test' as per normal.
//
// Notes:
//
// - Suite is for bigger packages that require a lot of setup and teardown code.
// - You can assigned paths and variables to the suite.
// - Instead of using assert, you can simply do t.Equal(want, got) without the
// 	 need of passing the test variable.
// - We make the test function a member of the struct so testing can be carried
//   out.
type StringsTestSuite struct {
	suite.Suite
}

// Assert testing has begun.
func TestStrings(t *testing.T) {
	suite.Run(t, new(StringsTestSuite))
}

// Assigns test base.
func (t *StringsTestSuite) SetupSuite() {
	fmt.Println("Im setting something up")
}

// The same principle applies, but we are using a suite.
func (t *StringsTestSuite) TestSplit() {
	got := Split("a/b/c", "/")
	want := []string{"a", "b", "c"}
	t.Equal(want, got)
}
