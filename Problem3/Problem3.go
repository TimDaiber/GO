








package main    

import ("fmt")

func main () {



fmt.Println("Enter a String")

var bla string

    fmt.Scan(&bla)

    //fmt.Println(bla)

    fmt.Println(Rev(bla))
}

func Rev (s string) string {
    runes := []rune(s)
    for i,j :=0, len(runes)-1; i<j;i,j=i+1,j-1{
        runes[i],runes[j]= runes[j],runes[i]
    }
    return string (runes)
}