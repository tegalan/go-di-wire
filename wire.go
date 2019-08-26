package main

import (
	"time"

	"github.com/google/wire"
)

// ProvideDeps ...
func ProvideDeps() *Deps {
	return &Deps{
		date: time.Now(),
	}
}

// ProvideService ...
func ProvideService(deps *Deps) Service {
	return Service{deps}
}

// ProvideAnotherService ...
func ProvideAnotherService(deps *Deps) AnotherService {
	return AnotherService{deps}
}

// InitService ...
func InitService() Service {
	wire.Build(ProvideDeps, ProvideService)
	return Service{}
}

// InitAnotherService ...
func InitAnotherService() AnotherService {
	wire.Build(ProvideDeps, ProvideAnotherService)
	return AnotherService{}
}
