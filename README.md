# msgpack

msgpack implementation avoiding allocations.

## Usage

```go
package main

import (
  "github.com/dgrr/msgpack"
)

func main() {
  b := msgpack.AppendString(nil, "Hello world")
  
  // setup connection or whatever you want
  
  conn.Write(b)
}
```
