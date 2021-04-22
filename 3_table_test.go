package gotest

// Table
//
// We declare a structure to hold our test inputs and expected outputs. This is our table.
// The tests structure is usually a local declaration because we want to reuse this
// name for other tests in this package.
//
// Notes:
// 	- By using a map[string]struct can can declare what test is running, so if it fails
// 	  we have a detailed message on what test case failed.
// 	- We range over the test cases, split them and see if they are equal.
//  - By using a map, we can harness the order of the test cases. Map iteration is
//    undefined, this means each time we run go test, our tests are going to be
//    potentially run in a different order.
//  - You can see that the last test fails, we need to account for that.
func (t *StringsTestSuite) TestSplit_Table() {
	tt := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		"Simple": {
			input: "a/b/c",
			sep:   "/",
			want:  []string{"a", "b", "c"},
		},
		"Wrong Sep": {
			input: "a/b/c",
			sep:   ",",
			want:  []string{"a/b/c"},
		},
		"No Sep": {
			input: "abc",
			sep:   "/",
			want:  []string{"abc"},
		},
		// TODO: Uncomment here for demo
		//"Trailing Sep": {
		//	input: "a/b/c/",
		//	sep:   "/",
		//	want:  []string{"a", "b", "c"},
		//},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := Split(test.input, test.sep)
			t.Equal(test.want, got)
		})
	}
}
