package lib

import "fmt"

func Log(remote_addr string, method string, path string) {
    fmt.Println(remote_addr, method, path)
}
