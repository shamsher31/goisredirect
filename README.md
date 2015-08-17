# goisredirect
Check if a number is a redirect HTTP status code

### How to install
```go
go get github.com/shamsher31/goisredirect
```

### How to use
```go
package main

import (
	"fmt"
	"github.com/shamsher31/goisredirect"
)

func main() {
  fmt.Println(redirect.Is(300))
  // true
}
```
###Aliasing Imports
If you already have package name ```redirect``` you can use following.
```go
package main

import (
	"fmt"
	pckRedirect "github.com/shamsher31/goisredirect"
)

func main() {
  fmt.Println(pckRedirect.Is(401))
  // false
}
```

### Related
[goisvdo](https://github.com/shamsher31/goisvdo)<br>
[goistext](https://github.com/shamsher31/goistext)<br>
[goisimg](https://github.com/shamsher31/goisimg)<br>

### Why
This package is inspired by [is-redirect](https://www.npmjs.com/package/is-redirect) npm module.

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
