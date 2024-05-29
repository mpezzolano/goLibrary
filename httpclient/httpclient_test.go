package httpclient

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, client"))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	body, err := Get(server.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if body != "Hello, client" {
		t.Fatalf("Expected body to be 'Hello, client', got %v", body)
	}
}

func TestPost(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	payload := "Hello, server"
	body, err := Post(server.URL, "text/plain", []byte(payload))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if body != payload {
		t.Fatalf("Expected body to be '%v', got %v", payload, body)
	}
}

func TestPatch(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	payload := "Hello, server"
	body, err := Patch(server.URL, "text/plain", []byte(payload))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if body != payload {
		t.Fatalf("Expected body to be '%v', got %v", payload, body)
	}
}

func TestPut(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	payload := "Hello, server"
	body, err := Put(server.URL, "text/plain", []byte(payload))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if body != payload {
		t.Fatalf("Expected body to be '%v', got %v", payload, body)
	}
}

func TestDelete(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Deleted"))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	body, err := Delete(server.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if body != "Deleted" {
		t.Fatalf("Expected body to be 'Deleted', got %v", body)
	}
}
