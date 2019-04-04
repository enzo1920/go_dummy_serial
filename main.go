package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/enzo1920/go_dummy_serial/version"
	"github.com/tarm/serial"
)

const AppVersion = "1.0.10 beta"

func check(e error) {
	if e != nil {
		log.Fatal(e)
		log.Print("\r\n")
	}
}

func reverse(array_byte []byte) []byte {
	for i := 0; i < len(array_byte)/2; i++ {
		j := len(array_byte) - i - 1
		array_byte[i], array_byte[j] = array_byte[j], array_byte[i]
	}
	return array_byte
}

func Comparer(need_to_comp []byte, fname string) bool {
	byte_arr := []byte("= -")
	var comp bool
	if need_to_comp[0] == byte_arr[0] && (need_to_comp[7] == byte_arr[1] || need_to_comp[7] == byte_arr[2]) {
		fmt.Println("Found! exit!")
		log.Printf("\n 0x3d and 0x20 found in place [0] and [7] exit!")
		fmt.Printf("Found! %s\n", need_to_comp)
		fmt.Printf("Reversed buff and write in results! %s\n", reverse(need_to_comp[1:7]))

		log.Printf("\n Stop logging!")
		log.Print("\r\n")
		comp = true

		//create result file
		f, err := os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		check(err)
		defer f.Close()
		//если - , то записываем число с минусом
		if need_to_comp[7] == byte_arr[2] {
			//chislo:=reverse(need_to_comp[1:7])
			need_to_comp[0] = need_to_comp[7]
			n2, err := f.Write(need_to_comp[:7])
			check(err)
			fmt.Printf("wrote %d bytes\n", n2)

		} else {

			n2, err := f.Write(need_to_comp[1:7])
			check(err)
			fmt.Printf("wrote %d bytes\n", n2)
		}

		//os.Exit(0)
	} else {
		fmt.Printf("\nNot found in: %v", need_to_comp)
		log.Printf("\nNot found in: %v", need_to_comp)
		comp = false
		//buf = buf[:0]
		//tmp_buf = tmp_buf[:0]
		//n=0
	}
	return comp
}

func Exists(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {

	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	fmt.Println(AppVersion)

	fileresults := "./results.txt"

	if Exists(fileresults) {
		os.Remove(fileresults)
	}

	file, err := os.OpenFile("dummy_serial.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
	var timeout int
	flag.IntVar(&port, "port", 6, "please number of  com-port, int")
	flag.IntVar(&timeout, "timeout", 5, "please timeout in secs int")
	flag.IntVar(&speed, "speed", 9600, "please enter speed of port int")

	flag.Parse()

	comport := "COM" + strconv.Itoa(port)

	//"=021000 "

	log.Printf("port:%s, speed %d, timeout: %d \r\n", comport, speed, timeout)
	c := &serial.Config{Name: comport, Baud: speed, ReadTimeout: time.Duration(timeout) * time.Second}
	stream, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
		log.Print("\r\n")
		return
	}
	buf := make([]byte, 8)
	tmp_buf := make([]byte, 8)

	for {
		n, errr := stream.Read(buf)
		check(errr)
		if n < len(buf) {
			for {

				k, errr := stream.Read(tmp_buf)
				check(errr)
				if k > 0 {
					for i := 0; i < k; i++ {
						if n >= len(buf) {
							//fmt.Printf("\n stop need compare!")
							break
						} else {

							buf[n+i] = tmp_buf[i]

						}
					}
					n = n + k
				} else {
					break
				}
				//clear tmp_buf
				tmp_buf = tmp_buf[:0]
				k = 0

			}
		}
		//fmt.Printf("\n run comparer")
		if Comparer(buf, fileresults) {
			os.Exit(0)
		}
		buf = buf[:0]
		n = 0

	}
}
