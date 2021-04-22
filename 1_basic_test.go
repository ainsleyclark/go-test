package gotest

import (
	"reflect"
	"testing"
)

// Simple
//
// Notes:
//
// - The name of the test function must start with Test.
// - The test function must take one argument of type *testing.T. A *testing.T
// 	 is a type injected by the testing package itself, to provide ways to print,
//	 skip, and fail the test.
func TestSplit(t *testing.T) {
	got := Split("a/b/c", "/")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

// Coverage:
//
// Notes:
// Checks for all possible code branches.
// go test -coverprofile=c.out
//
// PASS
// coverage: 100.0% of statements
// ok      github.com/ainsleyclark/go-test/1-basic 0.056s

// Going beyond 100% coverage.
//
// We need to test boundary conditions, for example, what happens if we try to
// split it on a comma?
// This is where table driven tests come into play.
func TestSplitWrongSep(t *testing.T) {
	got := Split("a/b/c", ",")
	want := []string{"a/b/c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
