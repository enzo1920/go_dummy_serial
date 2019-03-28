package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tarm/serial"
)

const AppVersion = "1.0.4 beta"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func reverse(array_byte []byte) []byte {
	for i := 0; i < len(array_byte)/2; i++ {
		j := len(array_byte) - i - 1
		array_byte[i], array_byte[j] = array_byte[j], array_byte[i]
	}
	return array_byte
}

func main() {

	fmt.Println(AppVersion)
	file, err := os.OpenFile("dummy_serial.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		log.Print("\r\n")
	}
	defer file.Close()
	log.SetOutput(file)
	log.Printf("\n Start logging!")
	log.Print("\r\n")

	//create result file
	f, err := os.Create("results.txt")
	check(err)
	defer f.Close()

	//parse flags
	var port int
	var speed int
	var timeout int
	flag.IntVar(&port, "port", 6, "please number of  com-port, int")
	flag.IntVar(&timeout, "timeout", 5, "please timeout in secs int")
	flag.IntVar(&speed, "speed", 9600, "please enter speed of port int")

	flag.Parse()

	comport := "COM" + strconv.Itoa(port)

	byte_arr := []byte("= ")
	/*
		testrr := []byte("=021000 ")
		fmt.Printf("\n testarr_str %v", testrr)
		testrrr1 := reverse(testrr)
		fmt.Printf("\n testrrr1 %s", testrrr1)
		_, err1 := f.Write(testrrr1)
		check(err1)*/

	log.Printf("port:%s, speed %d, timeout: %d \r\n", comport, speed, timeout)
	c := &serial.Config{Name: comport, Baud: speed, ReadTimeout: time.Duration(timeout) * time.Second}
	stream, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
		log.Print("\r\n")
		return
	}
	buf := make([]byte, 8)

	for {
		n, errr := stream.Read(buf)
		if errr != nil {
			log.Fatal(errr)
			log.Print("\r\n")
		}
		fmt.Println("i read: ", n)
		fmt.Printf("\nbuffer:%v", buf)

		if n == 8 {
			if buf[0] == byte_arr[0] && buf[7] == byte_arr[1] {
				fmt.Println("Found! exit!")
				log.Printf("\n 0x3d and 0x20 found in place [0] and [6] exit!")
				fmt.Printf("Found! %s\n", buf)
				fmt.Printf("Reversed buff and write in results! %s\n", reverse(buf))
				n2, err := f.Write(reverse(buf))
				check(err)
				fmt.Printf("wrote %d bytes\n", n2)

				log.Printf("\n Stop logging!")
				log.Print("\r\n")

				os.Exit(0)
			} else {
				fmt.Printf("\nNot found in: %s", buf)
			}

		}

	}

}
