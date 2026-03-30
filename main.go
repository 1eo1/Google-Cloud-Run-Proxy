package main

import (
    "io"
    "net"
    "os"
)

func handleClient(clientConn net.Conn, targetAddr string) {
    defer clientConn.Close()

    remoteConn, _ := net.Dial("tcp", targetAddr)
    defer remoteConn.Close()

    go func() {
        io.Copy(remoteConn, clientConn)
    }()

    io.Copy(clientConn, remoteConn)
}

func main() {
    listenAddr := ":" + os.Getenv("PORT")
    targetPort := os.Getenv("V2RAY_SERVER_PORT")
    if targetPort == "" {
        targetPort == ":80"
        }
    targetAddr := os.Getenv("V2RAY_SERVER_IP") + targetPort
    listener, _ := net.Listen("tcp", listenAddr)
    defer listener.Close()

    for {
        clientConn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleClient(clientConn, targetAddr)
    }
}
