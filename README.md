# usb-barcode-scanner

Small usb hid barcode scanner written in go.
This project is basically a go port of the node [usb-barcode-scanner](https://github.com/YaroslavSl/usb-barcode-scanner) library.

## Usage

```go
package main

import (
    "fmt"
    "github.com/sdorra/usb-barcode-sanner"
)

func onScan(barcode string) {
    fmt.Println(barcode)
}

func onError(err error) {
    panic(err)
}

func main() {
    scanner.Start(0xac90, 0x3002, onScan, onError)
}
```