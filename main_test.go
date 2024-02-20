package main

import (
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	var(
		got string
		want string
	)
	data, err := Handler()
	if err != nil{
		log.Fatal("failed to get data")
	}
	got = data.Message
	want = "hello"
	if got != want {
		t.Errorf("want %v, but got %v", want, got)
	}
}
