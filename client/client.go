package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
)

var text string

func getInput(c net.Conn) {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print(">> ")
        text, _ = reader.ReadString('\n')
        fmt.Fprintf(c, text+"\n")
}

func main () {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide port.")
                return
        }

        CONNECT := arguments[1]
        c, err := net.Dial("tcp", CONNECT)
        if err != nil {
                fmt.Println(err)
                return
        }

        for {
                // reader := bufio.NewReader(os.Stdin)
                // fmt.Print(">> ")
                // text, _ = reader.ReadString('\n')
                // fmt.Fprintf(c, text+"\n")
                go getInput(c)

                message, _ := bufio.NewReader(c).ReadString('\n')
                if strings.TrimSpace(strings.ToUpper(string(text))) == "EXIT" {
                        fmt.Println("TCP client exiting...")
                        return
                }
                if message != c.LocalAddr().String() + "\n" {
                        fmt.Print(message)
                } else {        
                        continue
                }
        }
}