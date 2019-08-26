package main

func main() {

	svc1 := InitService()
	svc2 := InitAnotherService()

	svc1.Print()
	svc2.PrintDate()
}
