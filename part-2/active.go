package main

import (
	"fmt"
	// "runtime"
	// "time"
)

type Group struct {
	Admin string 
    Member string
}

func main() {
	//switch case

    //3 ways we can abke to use switch case 
    // 1.variable declared inside the switch 
	// 2.variable declared outside the switch and used  
	// 3.with conditionals like else if


	// fmt.Println("System OS?")
    // switch os := runtime.GOOS ; os {
	// case "darwin" : 
	//     fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default : 
	//    fmt.Printf("%s \n",os)
	// }


    // today := time.Now().Weekday();
	// fmt.Println("When is Tuesday")
	// switch time.Tuesday {
	// case today + 0 : 
	//     fmt.Println("Today")
	// case today +  1:
	// 	fmt.Println("Tomorrow.")
	// default :
	//    fmt.Println("We passed it")
	// }


    // t := time.Now()
	// switch {
	// case t.Hour() < 12 :
	// 	  fmt.Println("Good Morning")
	// case t.Hour() < 17 :
	// 	  fmt.Println("Good Morning")
	// default : 
	//    fmt.Println("Good Night")
	// }

     //defer

	 /*  defer is used to push a function or a set of code is pushed into a stack and at the end of the program the functions are popped sequentially */
         
	// fmt.Println("counting")
	// for i := 0 ; i < 10 ; i++ {
	// 	defer fmt.Println(i)
	// }
    // fmt.Println("counting done")

	//pointer it is a data type that contains memory address of a variable , we can access that address to change or print the value stored in that address

	// i, j :=  23 , 345
    
	// p := &i
	// fmt.Println(*p)
	// *p = 22
	// fmt.Println(*p)
	// fmt.Println(&i)
	// fmt.Println(&j)

	// //struct 
	// Group1 :=  Group{Admin : "Admin1",Member : "Member1"}
    // p := &Group1
	// p.Admin = "Admin2"
	// fmt.Println(Group1)

	// //capitalizing struct members is important here , when struct is exported from one package to another 

	// var (
    //    g1 = Group{Admin:"Admin2"}
    //    g2 = Group{Member:"Member2"}
    //    g3 = Group{}
    //    g4 = &Group{Admin:"Admin2"}
	// )

	// fmt.Println(g1,g2,g3,g4)

	// //array
	// // go is statically typed and size is a part of the type
	// //so array sizes are fixed  
	// // collection of homogenous data type

	// var a [2]string
	// a[0] , a[1] = "Hello" , "World"
	// primes := [4]int{1,2,3,4}
	// fmt.Println(a , primes)
	
    //  // slice is gnenerally a method to create an array whether from the parent array or into s seperate slice(array)
	// //slice  this is like an array but is can be used to store multiple arrays , we can create our own slice or we can create an array and slice from it 
    // primes := [4]int{1,2,3,4}

	// num := primes[1:3] 
	// fmt.Println(num)
    
    // s := []int{1,2,3,4,5,6}
    
	// fmt.Printf("len=%d cap=%d",len(s) , cap(s))

	// var a []int
	// fmt.Print("\n",a ,len(a), cap(a))
	// if a == nil {
	// 	fmt.Println("nil")
	// }
    
	// //make()  is a pre built function to initialize reference types like alice , map , channel
   
	// m := make([]int,5)
	// fmt.Println(m)
    //the zero value of all the reference types will be nil 
	// var s []int
	// fmt.Println(s) // Output: []
	// fmt.Println(s == nil)
	
	// s = append(s,1)
	// fmt.Println(s) 

	//appending a slice , whenever there is no memory , at that when we try to append a new element the capacity of gets double , when Length == Capacity , at the next iteration capacity gets doubled 
	//How?  first when the slice is full , there are no memory location to store the next element , so when appending slice finds a new memory location and backing array and then increases the capacity of the array

	// var s []int
	
	// for i := 0 ; i < 20 ; i++ {
	// 	s = append(s, i)
	// 	fmt.Printf("iteration:%d len=%d cap=%d \n",i,len(s),cap(s))
	// }

	// what if we know the length of the slice then, we use

	a := make([]int , 0 , 5)
    
	for i := 0 ; i <  5; i++ {
		a = append(a , i)
		fmt.Printf("iteration:%d len=%d cap=%d \n",i,len(a),cap(a))
	}
     
    // maps  ->  these contains key value pairs like object in js 
    

}

