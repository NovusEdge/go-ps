# go-ps
`go-ps` is a simple implementation of a port-scanner in golang.

## Usage
To obtain the module, simply execute in shell:

```zsh
$ go get github.com/NovusEdge/go-ps
```

<br>

#### Example use-cases:

*] Importing the module:

```go
import gops "github.com/NovusEdge/go-ps"
```



*] Declaring an a variable of type `PortScanner`:

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


* `Scan()` iteartes over all ports in range [startPort, endPort] and reports which ports are open by printing to stdout.


```go
/* Scan() parameters:
	startPort [int] Port from which to start scanning (inclusive)
	endPort   [int] Port on which to stop the scanning. (inclusive)
	timeout   [time.Duration] timeout for each port being scanned.
*/
ps.Scan(1, 1024, 500*time.Millisecond)
```


