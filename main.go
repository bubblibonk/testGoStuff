package main

import(
    "fmt"
)
func Hello(name string)string{
    engConst:= "Hello, "
    if name == ""{
        name = "World"
    }
    return engConst + name
}
func main(){
    fmt.Println(Hello("Chris"))
}
