package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	eve := make(chan int)
	odd:= make(chan int)
	quit:= make(chan int)

	//Send the Data trought the chanel
	go send(eve,odd,quit)
	//Recieve the data trough the data
	recieve(eve,odd,quit)

	fmt.Println("The program is about to exit")
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


//////////////////////////////FIXME: CHANNELS IN GO FIXME:////////////////////////
//FIXME: Chanels in GO 
func ej4(){
	//make the channel
	c:= make(chan int)
	//send 
	go func(){
		for i := 0; i<5; i++{
			c <-i
		}
		close(c)
	}()
	//recieve
	for v:= range c{
		fmt.Println(v)
	}
	fmt.Println("The program is ending`")
}
////////////////////FIXME: SELECT IN CHANELS/////////////
//FIXME: Select
func recieve(e,o,q <-chan int){
	for{
		select{
		case v:= <-e:
			fmt.Println("From the eve channel:",v)
		case v:= <-o:
			fmt.Println("From the eve channel:",v)
		case v:= <-q:
			fmt.Println("From the eve channel:",v)
			return 
		}
	}
}
func send(e,o,q chan<-int){
	for i:= 0;i <100;i++{
		if i%2==0{
			e<-i
			}else{
				o<-i
			}
	}
	q<-0	
}
