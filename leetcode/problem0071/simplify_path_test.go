package problem0071

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(path string) string) {
	testCases := []struct {
		path string
		want string
	}{
		{
			path: "/home/",
			want: "/home",
		},
		{
			path: "/home//foo/",
			want: "/home/foo",
		},
		{
			path: "/home/user/Documents/../Pictures",
			want: "/home/user/Pictures",
		},
		{
			path: "/../",
			want: "/",
		},
		{
			path: "/.../a/../b/c/../d/./",
			want: "/.../b/d",
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.path)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestSimplifyPath(t *testing.T) {
	test(t, simplifyPath)
}
