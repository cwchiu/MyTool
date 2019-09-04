package server

import (
    "net"
    "fmt"
    "log"
)

func factory(port int, f func(conn net.Conn)) {

    ss, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        panic(err)
    }
    log.Println(fmt.Sprintf("Listen: %d", port))
    for {
        conn, err := ss.Accept()
        if err != nil {
            log.Println("accept error:", err)
            break
        }
        go f(conn)
    }
}