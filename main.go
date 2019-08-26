package main

import "time"

func main() {
	deps := Deps{
		date: time.Now(),
	}

	svc1 := Service{&deps}
	svc2 := AnotherService{&deps}

	svc1.Print()
	svc2.PrintDate()
}
