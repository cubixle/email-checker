# email-checker
Email checker is package with a collection of different methods of checking email addresses.

```go
package main

import "github.com/lukerodham/email-checker"

func main() {
    is, err := checker.IsGeneric("example@gmail.com")
    if err != nil {
        panic(err)
    }

    if is {
        // do something?
    }
}
```
