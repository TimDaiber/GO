





<<<<<<< HEAD
=======


>>>>>>> 7cd7c6711947773cb514b4a09dc621e2424dbfef
package main

import "fmt"

<<<<<<< HEAD
func main(){

    var  typemycounter int64;

    var primeNumbers[] int64;


}
=======
func main() {
	// Declare an array of size 11000 type int64
	var testNumberArray[100001] int64
    var counter int64 = 1
    //var primearray[] int
    var ki int64 = 1
	for ki = 1; ki < 50009; ki++ {
        
		if ki % 2 == 0 || ki % 3 ==0 || ki % 5 ==0 || ki % 7 == 0 {
            
        
        }else{
            testNumberArray[counter] = ki 
            counter++
        }
    }
	
    fmt.Println(testNumberArray[10001])
    /*
    for  k := 1 ; k < len(testNumberArray) ; k++ {
        
        if testNumberArray[k]==0{
            fmt.Println(k)
            k = 10000
        }
    }
    */
	

}
>>>>>>> 7cd7c6711947773cb514b4a09dc621e2424dbfef
