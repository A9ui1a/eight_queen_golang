package main

import "fmt"

var c int =0

func main() {
    board:=make([]int ,8)
    putQueen(board,0)
    fmt.Println("答案數量=>",c)
}

func putQueen(b []int,col int){
        if col>=8{
            fmt.Printf("%v\n",b)
            c++
            return
        }
        for i:=0;i<8;i++{
              b[col]=i+1
              if SafePut(b,i+1,col){
                putQueen(b,col+1)

              }

        }
}

func SafePut(a []int,b int,k int) bool{
    for i:=0;i<k;i++{
        if a[i]==b{
            return false
        }
    }
    for i:=0;i<8;i++{
        if k>i{
            if  a[k-i-1]==b-i-1|| a[k-i-1]==b+i+1{
                return false
            }
        }
    }


    return true
}