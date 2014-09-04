runlib
======

a go runner for libraries

Usage
-----

use 

```
runlib -pkg=github.com/user/package -func=X
```

to run the following code

```go
package main

import lib "github.com/user/package"

func main() {
    lib.X()
}
```

or if you are in the directory `$GOPATH/src/github.com/user/package` and want to run the function `Main`, run

```
runlib
```

StdIn, StdErr and StdOut are bound to the ones of the main process and available inside the package, as well as the environment and `os.Args`. 
The latter includes the arguments needed by runlib: `pkg` and `func`.