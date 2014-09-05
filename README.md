runlib
======

a go runner for libraries

Usage
-----

use 

```
runlib -pkg=github.com/user/package -func=X -args='-x=y'
```

to run the following code with the arguments `-x=y`

```go
package main

import lib "github.com/user/package"

func main() {
    lib.X()
}
```

or if you are in the directory `$GOPATH/src/github.com/user/package` and want to run the function `Main` and passing no arguments, run

```
runlib
```

StdIn, StdErr and StdOut are bound to the ones of the main process and available inside the package.