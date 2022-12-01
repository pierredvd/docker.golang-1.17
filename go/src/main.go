package main

import (
    "strconv"
    "fmt"
    "net"
)
	
func main() {
    fmt.Println("hello world")

    err := listen(80)
    fmt.Println(err)

}

func listen(port int) (err error){
    strpos := strconv.Itoa(port)
    listener, err := net.Listen("tcp", fmt.Sprintf(":%s", strpos))
    if err != nil{
        return err
    }
    defer listener.Close()

    fmt.Println(fmt.Sprintf("Server is now listening port %s", strpos))

    var ch chan error = make(chan error, 0)
    defer close(ch)
    go func(){
        for{
            conn, err := listener.Accept()
            if err != nil{
                ch <- err
            } else {
                go accept(conn)
            }
        }
    }()
    return <-ch
}

func accept(sock net.Conn){
    
    fmt.Println("entering connexion from", sock.RemoteAddr().String())

    var ch chan []byte  = make(chan []byte, 1020)
    go func(){
        var recvData        = make([]byte, 1024)
        for {
            numBytes, err := sock.Read(recvData)
            if err != nil{
                return
            }
            ch <- recvData[:numBytes]
        }
    }()

    data := string(<- ch)
    fmt.Println("receive data ", data)
    sock.Write([]byte("HTTP1/1 200 OK\r\nContent-Type: text/html\r\nContent-Length: 58\r\n\r\n<!doctype><html><body><h1>Hello world !</h1></body></html>"))
    sock.Close()

}