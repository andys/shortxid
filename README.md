# shortxid

Generates short, base62, unique identifiers that are sortable by timestamp.

Base62 has the advantage of being easy to read, sort, and regexp, and is cut-and-paste with auto text selection.

Each ID is built using:
* Optional global- and locally-prepended strings
* 5 bytes worth of milliseconds since 2023 (should work  up to 2057)
* 1 byte worth of  "instance ID" (passed in as an integer during initialization, modulo 256)
* 2 bytes worth of atomic incrementing counter

This means, out of the box, it isn't globally unique like a UUID, so you are required to *EITHER*:
* Init with an incrementing server or deployment ID, *OR*
* Use the global string prepend to uniquely identify the running code (eg. IP address or MAC address or long random number).

With a bit of up-front work to get an incrementing machine or deploy ID, your application can enjoy very short IDs,
that generate fast across a distributed cluster, as compared with other UID generators that end up absorbing a full
timestamp, MAC address, IP address/ports, etc.


```go

package main

import (
	"fmt"
	"github.com/andys/shortxid"
)

func main() {
  myInstanceID := 123
  generator := shortxid.NewGenerator(myInstanceID, "MyApp-")
  ID := generator.NewID("XID-")

  fmt.Println("result: ", ID)
}
```

Result:
```
result:  MyApp-XID-DgqfsyMdmIz
```
