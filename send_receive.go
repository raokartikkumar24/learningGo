
package main

import(
  "bytes"
  "flag"
  "net"
  "fmt"
  "time"
  "io"
  "os"
)

/*
* Each execution of this program does just one thing:
*   it either sends or receives (and descrads) a "repeates" number of predefined
*   messages in either TCP (connected) or UDP (connectionless) fashion.
*
* Its command-line parameters are:
*/
var n     = flag.Int("n", 1, "Number or repeates for the send/receive action (< 0 means infinitely).")
var proto = flag.String("p", "tcp4", "Protocol to use in the send/receive action.")
var to    = flag.String("to", "", "If present, the IPv4 address of the send action. Determines send mode.")
var port  = flag.Int("port", 8800, "Protocol to use in the send/receive action.")
var msg   = flag.String("msg", "Msg", "If sending, message to be sent.")

func main() {
  // Parse arguments
  flag.Parse()

  if *to != "" { // Send
    // repeate
    for i := 0; *n < 0 || i < *n; i++ {
      var addr string = fmt.Sprintf("%s:%d", *to, *port)
      //Connect TCP
      conn, err := net.Dial(*proto, addr)
      if err != nil {
        panic(err)
      }
      //simple write
      fmt.Printf("Send %d of %d %s protocol packets to address %s.\n", i+1, *n, *proto, addr)
      conn.Write([]byte(*msg + "\n"))
      conn.Close()
      time.Sleep(1 * time.Second)
    }
  }else{        // Receive?
    var addr string = fmt.Sprintf(":%d", *port)
    fmt.Printf("Receive %s protocol packets on port %d.\n", *proto, *port)
    if *proto == "tcp4" || *proto == "tcp6" {
      // listen to incoming tcp connections
      l, err := net.Listen(*proto, addr)
      if err != nil {
        panic(err)
      }
      defer l.Close()
      // A common pattern is to start a loop to continously accept connections
      for {
          //accept connections using Listener.Accept()
        c, err := l.Accept()
        if err!= nil {
          panic(err)
        }
        //It's common to handle accepted connection on different goroutines
        go io.Copy(os.Stdout, c)
      }
    } else if *proto == "udp4" || *proto == "udp6" {
      // listen to incoming udp packets
      for {
        pc, err := net.ListenPacket(*proto, addr)
        if err != nil {
        	panic(err)
        }

        //simple read
        buffer := make([]byte, 1024)
        pc.ReadFrom(buffer)
        io.Copy(os.Stdout, bytes.NewReader(buffer))

        pc.Close()
      }
    } else {
      panic("No valid protocol specified. Should be either 'tcp4', 'tcp6', 'udp4' or 'udp6'.")
    }
  }
}

