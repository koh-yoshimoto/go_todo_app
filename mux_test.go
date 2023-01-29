package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/koh-yoshimoto/go_todo_app/config"
)

func TestNewMux(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		t.Fatalf("%v", err)
	}
	sut, cleanup, err := NewMux(ctx, cfg)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Cleanup(cleanup)

	sut.ServeHTTP(w, r)
	resp := w.Result()
	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Error("want status code 200, but", resp.StatusCode)
	}
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	want := `{"status": "ok"}`
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
