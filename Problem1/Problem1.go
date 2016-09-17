package main

import "fmt"

func main()  {
    
    var sum int;

    for i := 0;i  <= 100 ; i++ {
        if i % 3 ==0 {
            sum = sum + i;
        }
        if i % 5 == 0{
            sum = sum +i;
        }
    }

//
    fmt.Println(sum);

}