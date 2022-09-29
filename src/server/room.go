package main

import (
	"sync"
)

type room struct {
	roomUsers mapCounter
	roomName string
}

type mapCounter struct {
	blocker sync.RWMutex
	users map[string]*ServerProcessor
}




