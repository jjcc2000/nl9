package main

import (
	"fmt"
	"runtime"
	"sync"
	
)

func main(){
	e:=make(chan int)
	o := make(chan int)
	f := make(chan int)
	go send_F(e,o)
	
	go recieve_F(e,o,f)

	for v:= range f{
		fmt.Println("This is the Value",v)
	}
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
			fmt.Println("From the odd channel:",v)
		case v:= <-q:
			fmt.Println("From the quit channel:",v)
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
//FIXME: The comma statement
//Is to check if the channel is still open//
func checOk(){
	x:= make(chan int)
	go func ()  {
		x<-100	
	}()
	v, ok := <-x

	fmt.Println("This is the value in X:",v)
	close(x)
	fmt.Println("This is the state of X:",ok)

	//Now reasign the values to the chanels 
	v ,ok= <-x
	fmt.Println("This is the new Values in the Chanels:",v)
	fmt.Println("This is the state of the new Chanels:",ok) 
}
//FIXME: FAN IN || FAN OUT
//Send Chanel
func send_F(e,o chan<-int){
	for i:= 0;i <100;i++{
		if i%2==0{
			e<-i
			}else{
				o<-i
			}
	}
	close(e)
	close(o)
}	
//Recievee Chanel
func recieve_F(e,o <-chan int, fanin chan<- int){
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		for v:= range e{
			fanin<-v
		}
		wg.Done()
	}()
	go func(){
		for v := range o{
			fanin<-v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)

}