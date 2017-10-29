# Go Packet Server

TCP server implementation that cuts incoming data into packets. 
Most TCP server example codes read data up to `byte 0` which terminates message and treats it as string which is very limited. 
Example: [repo](https://github.com/firstrow/tcp_server).  
  
In some cases you want your app to send and receive `integer`, `byte`, `float` or raw data, then this package may come in handy.  

## Installation
```
go get -u github.com/lafreak/go-packet-server
```

## Client implementation
- [C# (NuGet package)](https://github.com/lafreak/PacketClient-cs)
  
## Packet class
Structure of example packet:
```
0x07 0x00 0x08 0xC2 0x00 0x00 0x00 ...
|_______| |__| |_____________________|
    |      |      |
    |      |      `----- Data
    |      |
    |      `------------ Type
    |
    `------------------- Size
```

`Size` - first 2 bytes are reserved for packet size. This way server is sure how many bytes should be read from stream.  
`Type` - packet identifier. Each packet can be treated like an event, you can subscribe to each incoming packet of certain type.  
`Data` - custom information. This is space for all data that comes with packet. 
It can be interpreted on many different ways for example:  
- `1 integer value` - interpreted as number `194` (`C2 00 00 00`)  
- `2 byte values` - interpreted as byte `194` (`C2 00`) and `0` (`00 00`)  
  
Notice: all data is interpreted in little endian order. Strings terminate with `0`.

## Read from packet
Example packet:  `0x13 0x00 0x01 0x2D 0xFF 0xFF 0xFF 0xFF 0xFF 0xFF 0xFF 0xFF 0x47 0x6F 0x4C 0x61 0x6E 0x67 0x00`
``` go
game.On(1 /* Packet Type */, func(s *server.Session, p *server.Packet) {
  var smallNumber byte
  var bigNumber int64
  var message string
  
  p.Read(&smallNumber, &bigNumber, &message)
  
  // smallNumber == 45
  // bigNumber == -1
  // message == "GoLang"
}
```

## Write to packet
``` go
p := server.NewPacket(9 /* Type */)
p.Write(3 /*int*/, "Go", byte(2))
```
Packet `p` will represent: `0x0B 0x00 0x09 0x03 0x00 0x00 0x00 0x47 0x6F 0x00 0x02`  
Or just use `session.Send` method that creates and sends packet for you:
``` go
game.On(200 /* When server receives packet of type 200 */, func(s *server.Session, p *server.Packet) {
  s.Send(9 /* Type */, 3, "Go", byte(2))
})
```

## Sample usage
``` go
package main

import (
  "fmt"
  "github.com/lafreak/go-packet-server"
)

func main() {
  game := server.New("localhost:3000")

  // Subscribe to connection event.
  game.OnConnected(func(s *server.Session) {
    fmt.Println("Client connected.")
  })

  // Subscribe to disconnection event.
  game.OnDisconnected(func(s *server.Session) {
    fmt.Println("Client disconnected.")
  })

  // This event fires when undefined packet was received.
  // It means client send packet of type you did not subscribed to.
  game.OnUnknownPacket(func(s *server.Session, p *server.Packet) {
    fmt.Println("Unknown packet:", p.Type())
  })

  // Subscribe to incoming packets of type 25.
  // Read integer from packet stream.
  game.On(25, func(s *server.Session, p *server.Packet) {
    var n int
    p.Read(&n)
    fmt.Println("N was received:", n)
  })

  game.Start()
}
```

## Todo
- packet encrypt/decrypt
- read/write raw data and float
- code client for various languages
