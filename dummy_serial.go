package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tarm/serial"
)

func main() {

	file, err := os.OpenFile("com_port_reader.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		log.Print("\r\n")
	}
	defer file.Close()
	log.SetOutput(file)
	log.Printf("\n Start logging!")
	log.Print("\r\n")

	//parse flags
	var port int
	var speed int
	var cmd string
	flag.IntVar(&port, "port", 6, "please number of  com-port, int")
	flag.IntVar(&speed, "speed", 9600, "please enter speed of port int")
	flag.StringVar(&cmd, "command", "ATI", "please enter a command , string ")
	flag.Parse()

	comport := "COM" + strconv.Itoa(port)

	log.Printf("port:%s, speed %d, cmd %s\r\n", comport, speed, cmd)

	c := &serial.Config{Name: comport, Baud: speed}
	stream, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
		log.Print("\r\n")
		return
	}
	cmd = cmd + "\r\n"
	n, err := stream.Write([]byte(cmd))
	if err != nil {
		log.Fatal(err)
		log.Print("\r\n")
	}
	fmt.Println("write", n)

	buf := make([]byte, 1024)
	for {
		n, errr := stream.Read(buf)
		if errr != nil {
			log.Fatal(errr)
			log.Print("\r\n")
		}
		s := string(buf[:n])
		fmt.Println(s)
	}

}
