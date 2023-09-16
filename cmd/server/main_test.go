package main

import (
	"net/http/httptest"
	"testing"
)

func TestHelloHandlerRespondsWithHelloWorld(t *testing.T) {
	// setup request and recorder
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	want := "Hello, world!"

	// call hello handler
	hello(rec, req)

	got := rec.Body.String()
	status := rec.Code
	// assert response is "Hello, world!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	// assert status code is 200
	if status != 200 {
		t.Errorf("status code got %d, want %d", status, 200)
	}
	// assert content type is text/plain
	if ctype := rec.Header().Get("Content-Type"); ctype != "text/plain" {
		t.Errorf("content type header got %q, want %q", ctype, "text/plain")
	}

}

func TestAboutHandler(t *testing.T) {
	// setup request and recorder
	req := httptest.NewRequest("GET", "/about", nil)
	rec := httptest.NewRecorder()

	want := "About"

	// call about handler
	about(rec, req)

	got := rec.Body.String()
	status := rec.Code
	// assert response is "About"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	// assert status code is 200
	if status != 200 {
		t.Errorf("status code got %d, want %d", status, 200)
	}
	// assert content type is text/plain
	if ctype := rec.Header().Get("Content-Type"); ctype != "text/plain" {
		t.Errorf("content type header got %q, want %q", ctype, "text/plain")
	}
}
