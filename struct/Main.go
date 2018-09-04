package main

import "fmt"

type Employee struct {
    name string
    id int
    manager string
    dependents []string
}

func main() {
    fmt.Println("Hello, Go! Structs...\n");
    amit :=  Employee{
        "Amit Verma",
        7630,
        "Sendilraj Selvaraju",
        []string{"Rashi Dhawan",
         "Avantika Verma",
         "Aditya Verma",
         "Arun Kumar Verma",
         "Santosh Verma"} }

    fmt.Println(amit)
}
