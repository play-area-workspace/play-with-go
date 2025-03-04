package main

import (
	"fmt"
	"runtime"
	"log"
	"os"
)



func main() {
	fmt.Println("This is the begining")
	anotherFile()
	a := "Hey a new variable"
	fmt.Println(a)
	var c int
	c = 8
	fmt.Println(c)
	primes := [6]int{2,3,6,8,4,5}

	var slice []int = primes[1:4]

	fmt.Println(slice)

	m:= make(map[string]int)

	m["ONE"] = 1

	sum:=0
	// for loop
	for i:= 0; i < 10; i++{
		sum+=1
		fmt.Println(sum)
	} 

	// imitating while loop with for
	for sum < 20{
		sum+=1
		fmt.Println("while ",sum)
	}

	// for range loop
	fmt.Println("for range loop")
	for i, v := range primes{
		fmt.Println(i,v)
	}

	if sum > 10 {
		fmt.Println("sum is greater than 10")
	}else{
		fmt.Println("sum is less than 10")
	}

	switch os:= runtime.GOOS;os {
		case "darwin":
			fmt.Println("OS X")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Printf("Other OS %s \n",os)
	}

	// error handling
	f,err := os.Open("filename.ext")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// read file
	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes: %s\n", n, string(buf))
	
	/* 
	defer f.Close()
    buffer := make([]byte, 1024)
    for {
        n, err := f.Read(buffer)
        if err != nil {
            if err == os.EOF {
                break
            }
            log.Fatal(err)
        }
        fmt.Print(string(buffer[:n]))
    } 
	 */	

}