# ISO8583
Go for iso8583

This is a go lang project that assist in processing ISO 8583 message standards.
``` go get github.com/CliffordN/ISO8583
```
 # Example

 ```go
 package main

 import (
  	"github.com/CliffordN/ISO8583"
 )

  func main() {

    msg := isoFormatter.Message{}
    
    msg.MTI = "0200"
    msg.AddData(2, 12345678)
    msg.AddData(3, "310000")
    msg.AddData(4, isoFormatter.PadAmount("20.00", 12, "0"))
    msg.AddData(7, 303103707)
    msg.AddData(11, 133337)
    msg.AddData(13, 2015)
    msg.AddData(37, "010215040512")
    msg.AddData(49, "404")
    msg.AddData(102, "123456789")
    
    isomsg := msg.Build()
    fmt.Println(isomsg)

   }
 ```

  # output
 ```
 0200f228000008008000000000000400000008123456783100000000000002000303103707133337201501021504051240409123456789
 ```

 Additional example you can see in tester.go
