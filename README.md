# go-ps
`go-ps` is a simple implementation of a port-scanner in golang.

## Usage
To obtain the module, simply execute in shell:

```zsh
$ go get github.com/NovusEdge/go-ps
```

<br>

#### Example use-case:

* Importing the module:

```go
import gops "github.com/NovusEdge/go-ps"
```

<br>
<br>


* Declaring an a variable of type `PortScanner`:
 
```go
// Decalring a PortScanner object.

/*
struct definition:
type PortScanner struct {
	Domain   string
	Protocol string
}
*/

ps := gops.PortScanner{
	Domain: "scanme.nmap.org",
	Protocol: "tcp",
}

```

<br>
<br>

* `Scan()` iterates over all ports in range [startPort, endPort] and reports which ports are open by _printing to stdout_.


```go
/* Scan() parameters:
	startPort [int] Port from which to start scanning (inclusive)
	endPort   [int] Port on which to stop the scanning. (inclusive)
	timeout   [time.Duration] timeout for each port being scanned.
*/
ps.Scan(1, 1024, 500*time.Millisecond)
```

Output: 
```
[*] Port 80 open
...
...
```

<br>
<br>

* `Scans()` iterates over all ports in range [startPort, endPort] and _returns_ a list of open ports.


```go
/* Scans() parameters:
	startPort [int] Port from which to start scanning (inclusive)
	endPort   [int] Port on which to stop the scanning. (inclusive)
	timeout   [time.Duration] timeout for each port being scanned.

Returns: chan int
*/

ports := ps.Scans(1, 1024, 500*time.Millisecond)

fmt.Println("Open Ports: ", ports)
```

Output:
```
Open Ports: [ 80 448 ... ]
```

***

### Sample program:
```go
package main

import (
    "fmt"
    "time"

    gops "github.com/NovusEdge/go-ps"
)

func main() {
    ps := gops.PortScanner{
	Domain: "scanme.nmap.org",
	Protocol: "tcp",
    }
    
    // Scan and report open ports in stdout:
    ps.Scan(1, 1024, 500*time.Millisecond)

    // Scan ports and get a list of open ports:
    openPorts := ps.Scans()

    // Use in whatever way you like :)
    // For this sample, we'll just print it to stdout:

    fmt.Println(openPorts)
}
```


