package main

import (
	"time"

	"github.com/google/wire"
)

// Single Instance
var deps *Deps

// ProvideDeps ...
func ProvideDeps() *Deps {
	if deps != nil {
		return deps
	}

	deps = &Deps{
		date: time.Now(),
	}

	return deps
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
