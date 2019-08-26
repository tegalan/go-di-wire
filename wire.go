package main

import "time"

// ProvideDeps ...
func ProvideDeps() *Deps {
	return &Deps{
		date: time.Now(),
	}
}

// ProvideService ...
func ProvideService(deps *Deps) Service {
	return Service{Deps: deps}
}

// InitService ..
func InitService() {

}
