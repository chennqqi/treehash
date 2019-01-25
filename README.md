# treehash

Implements the SHA256 tree hash algorithm (as used by [aws glacier](https://docs.aws.amazon.com/amazonglacier/latest/dev/checksum-calculations.html))

## Installation

Download and install :

```
$ go get github.com/downeast/treehash
```

Add it to your code :

```go
import "github.com/downeast/treehash"
```

## Use

```go
file, _ := os.Open("filename")
th := treehash.New()
io.Copy(th, file)
checksum := fmt.Sprintf("%x", th.Sum(nil))
```