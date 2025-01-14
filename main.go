package main

import (
	"fmt"
	"math"
)

var python bool

func add( a , b int) int {
	return a + b
}

func Pow(x,n,lim float64) float64 {
	if v := math.Pow(x,n) ; v < lim {
         return v
	}
    //can't use  v here
	return lim
}

func div(a,b int) int {
	if( b == 0) {
		fmt.Print("cannot divide with Zero") 
	}
	div := a/b
    return div
	
}

func main() {
	// a,b := 32 , 45
    // fmt.Println("Addition of numbers:",add(a,b),python)

	// var i int
	// var s string
	// var b bool
	// var f float64
     
	// num := 47.3

	// var float float64 = float64(num)

	// fmt.Printf("%v %q %v %v",i,s,b,f)
	// fmt.Println(float)

	// sum := 0
	// for i := 0 ; i < 10 ; i++ {
	// 	sum += i
	// }
     
    // sum := 1
	// // for sum < 10 {
	// // 	sum += sum
	// // }

	// for {
	// 	if (sum > 100) {
	// 		break
	// 	}
	// 	sum += sum
	// 	fmt.Println(sum)
	// }
    
     
	
}

/*
   variables can be declared in two ways , 
   ** var name type  , variable name and its type are given 
   or for multiple variable declarations 
   ** var (
        name1 type1 
		name2 type2 
   )

   this type of declaration can be used when the initial value is unknown 
   this can be used at the global and local scope of declaration 

Another way will be  name := value
    this type of declaration is used when the value is known 
	this can be used inly at the local scope 
*/
/*
   varaibles and their types 

   1)basic :
       number  ->> int  int8 int16 int32 int128  we can declare the size level like int8 ,int16 based on the level needed 
               ->> unit same as the level in int provides only positive numbers
               ->> byte
			   ->> rune 
			   ->> float
	   , string , bool
	2)Aggregate : -> they have collection of similar or different basic data types
	    array : A collection of similar data types  		
		struct : A collection of heterogenous data types
	3)Reference : -> thay don't have specific types but they refer to the data types
	    pointers , slice , function , channel , maps  
    4)Interfaces

	type inference -  go make us to change type of variable explicitly 
    type reference - this means when assign a type explicitly or implicitly with := , we can't use that variable with another type o value
*/

/*
   const   
      -> it allows to have basic data type values 
	  -> it is Constant
	  -> when to use?  this is used when you to need a hard coded values like url , long digit number , something like that
*/

/*
   loops 
    -> Go has only for loops 
*/