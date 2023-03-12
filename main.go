package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){

}
///////////////TODO: Ej1 TODO:///////////////////
type person struct{
	first string 
}
func (p *person) speak(){
	fmt.Println("Hello")
}
type human interface {
	speak()
}
func saySomething(h human){
	h.speak()
}
///////////////TODO: Eje2 TODO://///////////////
//Created an incrementer program 
func ej2(){
	var wg sync.WaitGroup

	increment:=0
	gs:=100
	wg.Add(gs)

	for i :=0; i<gs; i++{
		go func(){
			v := increment
			runtime.Gosched()
			v++
			increment=v
			wg.Done() 
		}()
	}
	wg.Wait()
	fmt.Println(increment)

}