# set
🤲 Effective and minimal set implementation for Go

Now support:

- String set 
- ... 

## Usage

```go

package main

import (
    "fmt"

    "github.com/danyanya/set"
)

func main() {
    s := set.String{}
    s.Store("aa")

    if s.Load("aa") {
        fmt.Println("aa is in set")
    } else {
        fmt.Println("aa is not in set")
    }
}


```


## Copyright

Made by Sliusar Danya in 2020. 
 
