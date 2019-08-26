package main

import (
	"log"
	"time"
)

// Deps ...
type Deps struct {
	date time.Time
}

// Service ...
type Service struct {
	Deps *Deps
}

// AnotherService ...
type AnotherService struct {
	Deps *Deps
}

// Print Function
func (s *Service) Print() {
	log.Println("SVC1", s.Deps.date)
}

// PrintDate function
func (s *AnotherService) PrintDate() {
	log.Println("SVC2", s.Deps.date)
}
