package main

import "fmt"

//import "fmt"

func Sum(numbers[] int)int{
    sum:=0
    for _,number:=range numbers{
        sum+=number

    }
    return sum
}

func SumAll(numbersToSum ...[]int)[]int{
    var sum []int
    for _,numbers:=range numbersToSum{
        sum = append(sum,Sum(numbers))
    }
    return sum
}

func SumAllTail(numbersToSum ...[]int)[]int{
    var sum []int
    for _,numbers:=range numbersToSum{
        if len(numbers) == 0{
            sum = append(sum, 0)
        }else{
            tails:=numbers[1:]
            fmt.Println(tails)
            sum = append(sum,Sum(tails))
        }
    }
    return sum
}
func main(){
    var num1 = []int{1,2}
    var me = []int{0,9}

    fmt.Println(SumAllTail(num1,me))
}
