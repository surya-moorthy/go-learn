package main

import (
	"fmt"
	"math"
)

//anonymous function  a function is declared 

func deliver(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

//methods

type Vertex struct {
	X , Y float64
}

func (v Vertex) Abs() float64 {
	fmt.Println( "x1 : ",v.X)
   return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex ) Scale( f float64) {
	fmt.Println("x2 : ",v.X)
	v.X = v.X * f
	fmt.Println("x3 : ",v.X)
	v.Y = v.Y * f

}
func Scalefunc(v *Vertex , f float64) {
	fmt.Println("x2 : ",v.X)
	v.X = v.X * f
	fmt.Println("x3 : ",v.X)
	v.Y = v.Y * f

}



func main() {
	hypo := func(x float64, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypo(5,12))
	fmt.Println(deliver(hypo))

	v := Vertex{5 ,6}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())

	// functions vs methods in case of pointer receiver
	
}


// value receiver and pointer receiver when to use this and that?
/*
    mutating -> this case can use pointer receiver 
	for longer structs ->  if we use it in sender receiver , we need to allocate to all the all values including copy , in that case we use pointer receiver     
*/

/*
    
*/