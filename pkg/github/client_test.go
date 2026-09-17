package github

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRequest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s", r.Method)
		}
		if r.URL.Path != "/repos/bonsai/LAB-Rapid-Prototype/issues" {
			t.Fatalf("path = %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-token" {
			t.Fatalf("authorization = %s", got)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"number":42,"html_url":"https://github.com/bonsai/LAB-Rapid-Prototype/issues/42","title":"hello"}`))
	}))
	defer srv.Close()

	c := New("test-token", "bonsai", "LAB-Rapid-Prototype")
	c.BaseURL = srv.URL
	c.HTTP = srv.Client()

	issue, err := c.CreateRequest(context.Background(), Request{Title: "hello", Body: "world", Source: "llm"})
	if err != nil {
		t.Fatal(err)
	}
	if issue.Number != 42 || issue.Title != "hello" {
		t.Fatalf("issue = %+v", issue)
	}
}

func TestCreateRequestRequiresFields(t *testing.T) {
	c := New("", "bonsai", "LAB-Rapid-Prototype")
	_, err := c.CreateRequest(context.Background(), Request{Body: "world"})
	if err == nil {
		t.Fatal("expected validation error")
	}
}
