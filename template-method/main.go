package main

import "fmt"

type AbstructWorker struct {
	client Client
}

func (a AbstructWorker) Do() {
	a.client.Setup()
	a.client.Run()
	a.client.Teardown()
}

type Client interface {
	Setup()
	Run()
	Teardown()
}

type ClientImpl1 struct {
	AbstructWorker
}

func NewClientImple() ClientImpl1 {
	c := ClientImpl1{}
	c.client = c
	return c
}

func (c ClientImpl1) Setup() {
	fmt.Println("👹")
}

func (c ClientImpl1) Run() {
	fmt.Println("🌟")
}

func (c ClientImpl1) Teardown() {
	fmt.Println("👹")
}

func main() {
	c := NewClientImple()
	c.Do()
}
