package main

import(
    "fmt"
)
const num = 5
func Repeated(character string)string{
    var repeated string

    for i:=0;i<num;i++{
        repeated+=character
    }
    return repeated
}
func Hello(name string)string{
    engConst:= "Hello, "
    if name == ""{
        name = "World"
    }
    return engConst + name
}
func main(){
    fmt.Println(Hello("Chris"))
    fmt.Println(Repeated("a"))
}
