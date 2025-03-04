package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
)

type Person struct {
	Name string		`json:"name"`
	Age int 		`json:"age"`
	Address string	`json:"address"`
	SoyDev bool		`json:"soydev"`
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years old) lives in %v", p.Name, p.Age, p.Address)
}

func (p Person) Jump(){
	fmt.Println("Jumping")
}

func (p Person) PickUpBeer() (bool,error){
	if p.Age < 18 {
		return false, errors.New("you are not allowed to pick up beer")
	}
	return true, nil
}

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

	// struct 

	// Creating an instance of Person using struct literal
	person1 := Person{
		Name:    "Alice",
		Age:     25,
		Address: "New York",
		SoyDev:  true,
	}

	// Calling the String method
	fmt.Println(person1.String()) // Output: Alice (25 years old) lives in New York

	// Calling the Jump method
	person1.Jump() // Output: Jumping

	// Calling the PickUpBeer method
	canPickUp, err := person1.PickUpBeer()
	if err != nil {
		fmt.Println(err) // If under 18, it will print: "you are not allowed to pick up beer"
	} else {
		fmt.Println("Can pick up beer:", canPickUp)
	}

	// Creating another instance of Person with age below 18 to check beer restriction
	person2 := Person{
		Name:    "Bob",
		Age:     16,
		Address: "Los Angeles",
		SoyDev:  false,
	}

	// Calling the PickUpBeer method for person2
	canPickUp, err = person2.PickUpBeer()
	if err != nil {
		fmt.Println(err) // Output: "you are not allowed to pick up beer"
	} else {
		fmt.Println("Can pick up beer:", canPickUp)
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

	 fmt.Println("This is the end")

}