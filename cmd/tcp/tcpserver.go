package tcp

import (
  "fmt"
  "net"
  "os"
)

const (
  CONN_HOST = "localhost"
  CONN_PORT = "3333"
  CONN_TYPE = "tcp"
  MAX_WORKERS = 10
)

func handleRequest(conn net.Conn) {
  defer conn.Close()

  buf := make([]byte, 1024)
  reqLen, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading", err.Error())
    return
  }
  fmt.Print(buf)

  fmt.Println("Received:", string(buf[:reqLen]))
}

func worker(id int, jobs <-chan net.Conn) {
  for conn := range jobs {
    fmt.Printf("Worker %d handling connection\n", id)
    handleRequest(conn)
  }
}

func TcpServer() {
  l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
  if err != nil {
    fmt.Println("Error listening:", err.Error())
    os.Exit(1)
  }
  defer l.Close()

  fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

  jobs := make(chan net.Conn, MAX_WORKERS)

  for i := 1; i <= MAX_WORKERS; i++ {
    go worker(i, jobs)
  }

  for {
    conn, err := l.Accept()
    if err != nil {
      fmt.Println("Error accepting connection:", err.Error())
      continue
    }

    jobs <- conn
  }
}
