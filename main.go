/*
	runlib imports the given package and runs the given function within from the package.
	The function that should be run must not receive or return any parameters.
*/
package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"flag"
	"fmt"
)

var pkg = flag.String("pkg", "", "the package that should be imported, if not provided, the package in the current directory is chosen")
var fun = flag.String("func", "Main", "the function that should be run")
var args = flag.String("args", "", "the arguments for the program")
var gopathForTempfile = ""

var template = `
package main

import lib "%s"

func main() {
	lib.%s()
}
`

func ok(err error) {
	if err != nil {
		panic(err)
	}
}

func install() {
	cmd := exec.Command("go", "install", *pkg)
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: Could not install %q: %s\n", *pkg, out)
		os.Exit(1)
	}
}

func run() {
	dir, err := ioutil.TempDir(gopathForTempfile, "go-runlib_")
	ok(err)
	defer os.RemoveAll(dir)
	mainFile := filepath.Join(dir, "main.go")
	err = ioutil.WriteFile(
		mainFile,
		[]byte(fmt.Sprintf(template, *pkg, *fun)),
		0666)
	ok(err)

	args_ := []string{
		"run", mainFile,
	}

	if args != nil && *args != "" {
		args_ = append(args_, *args)
	}
	//args_ = append(args_, os.Args...)
	cmd := exec.Command("go", args_...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}

func findPackageForDir(dir string, wd string) string {
	if !strings.HasPrefix(wd, dir) {
		return ""
	}
	path, err := filepath.Rel(dir, wd)
	if err != nil {
		return ""
	}
	return path
}

func findPackagePath(gopath string, wd string) string {
	gopaths := filepath.SplitList(gopath)

	for _, gp := range gopaths {
		path := findPackageForDir(filepath.Join(gp, "src"), wd)
		if path != "" {
			return path
		}
	}
	return ""
}

func setGopathForTempfile() {

	gopaths := os.Getenv("GOPATH")
	if gopaths == "" {
		panic("GOPATH is not set")
	}

	gopathForTempfile = filepath.SplitList(gopaths)[0]
}

func setDefaultPkg() {
	if pkg == nil || *pkg == "" {
		wd, err := os.Getwd()
		ok(err)
		wd, err = filepath.Abs(wd)
		ok(err)
		path := findPackagePath(os.Getenv("GOPATH"), wd)
		if path == "" {
			fmt.Printf("Error: can't find package path for %q\n", wd)
			os.Exit(1)
		}
		*pkg = path
	}
}

func main() {
	flag.Parse()

	setGopathForTempfile()
	setDefaultPkg()
	install()
	run()
}
