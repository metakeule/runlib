package main

import "testing"

func TestFindPackagePath(t *testing.T) {
	type test struct {
		gopath      string
		wd          string
		packagePath string
	}

	corpus := []test{
		{"/a/b/c", "/a/b/c/src/e/f/g", "e/f/g"},
		{"/x/z/y", "/a/b/c/src/e/f/g", ""},
		{"/x/z/y:/a/b/c", "/a/b/c/src/e/f/g", "e/f/g"},
		{"/x/z/y:/p/t/k", "/a/b/c/src/e/f/g", ""},
	}

	for _, testcase := range corpus {
		gotPath := findPackagePath(testcase.gopath, testcase.wd)

		if gotPath != testcase.packagePath {
			t.Errorf("expecting path %q for GOPATH %q and wd %q, but got %q", testcase.packagePath, testcase.gopath, testcase.wd, gotPath)
		}
	}

}
