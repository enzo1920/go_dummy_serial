# Simple serial port reader



## Getting Started
This app for read som COM-port and write results to results.txt
### Prerequisites for build
```
go get github.com/tarm/serial
go get github.com/enzo1920/go_dummy_serial/version
```

###Run programm 
```
make windows
or
make linux 
```
command will create a binary files
run on win32:
```
./bin/goserialreadlin32  -port=<COM port int> // 1 or 5 or other int value
```
after find a some value app will create results.txt with value and exit
###Run programm 
log in current dir with *.log



## Authors

***Sergei Larichkin** - - https://github.com/enzo1920/

