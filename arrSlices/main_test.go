package main

import (
	"reflect"
	"testing"
)
func TestSum(t *testing.T){
   t.Run("slices",func(t *testing.T){
        num:=[]int{1,2,3,4,5}
        got:=Sum(num)
        want:=15

        if got!=want{
            t.Errorf("got %q want %v, %d",got,want,num)
        }
        
        
    })
}
func TestSumAll(t *testing.T){
    got := SumAll([]int{1, 2}, []int{0, 9})
    want:=[]int{3,9}

    if !reflect.DeepEqual(got,want){
        t.Errorf("got %v want %v",got,want)
    }
}
func TestSumAllTail(t *testing.T){
    got:=SumAllTail([]int{1,2},[]int{0,9})
    want:=[]int{2,9}

    if !reflect.DeepEqual(got,want){
        t.Errorf("got %v want %v",got , want)
    }
}



















