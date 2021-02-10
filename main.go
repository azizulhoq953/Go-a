package main
import "fmt"
//import "reflect"

func main(){

//primitive data type
//rune,byte=alias

//composite Data type
//array
//slice
//map
//struct

//slice

//more then array


//var students [10]string
//var students [10]string

//students[0]="Azizul"
//students[1]="Galib"
//students[3]="Safin"
//students[4]="Sakib"
//students[5]="Akib"

//fruit := students[0:6]
//fmt.Println(fruit)

//x:=make([]string,5)
//fmt.Println(x,len(x))

var Mango []string

Mango=append(Mango, "Apple","\n banana","\nWater Milon")

fmt.Println(Mango,len(Mango) )
fmt.Printf("%T",Mango)

//fmt.Printf("%T",Mango)

//a := reflect.TypeOf(Mango).Kind().String()

//fmt.Printf("\n")

//fmt.Println(a)

}