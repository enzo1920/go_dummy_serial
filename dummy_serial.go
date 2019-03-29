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
	
	//"=021000 "
	//testrr := []byte("xyz")
	/*fmt.Printf("\n testarr_str %v", testrr)
	testrrr1 := []byte("87659")
	fmt.Printf("\n testrrr1 %s", testrrr1)
	_, err1 := f.Write(testrrr1)
	check(err1)
	thirdarr:= append(testrr ,testrrr1...)
	fmt.Printf("\n thirdarr %s", thirdarr)*/
/*
	b := make([]byte, 8)
	b = []byte("87665966")
	b = append(b, testrr...)
	b = append(b, byte_arr...)
	fmt.Printf("\n b %s", b)
	b = b[:0]
	fmt.Printf("\n len(b) %s cap(b)%s", len(b), cap(b))
	*/

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
	
	//buf1 :=buf[:n]
		for {
			n, errr := stream.Read(buf)
	        check(errr)
	        fmt.Println("i read: ", n)
	        fmt.Printf("\nbuffer:%v", buf)
			k, errr := stream.Read(tmp_buf)
			check(errr)
			fmt.Printf("\ni read k-bytes: %v", k)
			fmt.Printf("\ntmp_buf:%v", tmp_buf)
			time.Sleep(time.Duration(timeout) * time.Second)
	        if n > 0 || k > 0{
				for i := 0;  i<=k; i++ {
					if n+i == len(buf) { 
						break 
					}else{
						buf[n+i] = tmp_buf[i]
					}
	
				}
				fmt.Printf("\n add to buf? now  buffer contains :%v", buf)
				//clear tmp_buf
				tmp_buf = tmp_buf[:0]
				n = n+k
				if n >= len(buf) {
					if buf[0] == byte_arr[0] && buf[7] == byte_arr[1] {
						fmt.Println("Found! exit!")
						log.Printf("\n 0x3d and 0x20 found in place [0] and [7] exit!")
						fmt.Printf("Found! %s\n", buf)
						fmt.Printf("Reversed buff and write in results! %s\n", reverse(buf))
						n2, err := f.Write(reverse(buf))
						check(err)
						fmt.Printf("wrote %d bytes\n", n2)
			
						log.Printf("\n Stop logging!")
						log.Print("\r\n")
			
						os.Exit(0)
					} else {
							fmt.Printf("\nNot found in: %v", buf)
							buf = buf[:0]
							tmp_buf = tmp_buf[:0]
						}
					 //break 
					}
			} 
			
        }


	



}