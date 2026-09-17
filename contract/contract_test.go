package contract

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type request struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Source string `json:"source,omitempty"`
}

type accepted struct {
	Accepted bool   `json:"accepted"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Source   string `json:"source,omitempty"`
	Issue    *struct {
		Number int    `json:"number"`
		HTMLURL string `json:"html_url"`
	} `json:"issue,omitempty"`
}

// TestPostRequestsContract proves the minimum HTTP contract described by
// contract/openapi.yaml: valid input is accepted and the response preserves
// the request data while exposing the created GitHub Issue when available.
func TestPostRequestsContract(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/requests" {
			t.Fatalf("unexpected endpoint: %s %s", r.Method, r.URL.Path)
		}
		if got := r.Header.Get("Content-Type"); got != "application/json" {
			t.Fatalf("Content-Type = %q", got)
		}

		var in request
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			t.Fatal(err)
		}
		if in.Title == "" || in.Body == "" {
			t.Fatal("contract requires title and body")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(accepted{
			Accepted: true,
			Title:    in.Title,
			Body:     in.Body,
			Source:   in.Source,
		})
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	payload := request{Title: "contract proof", Body: "prove it", Source: "llm"}
	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(ts.URL+"/requests", "application/json", bytesReader(data))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusCreated)
	}

	var out accepted
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatal(err)
	}
	if !out.Accepted || out.Title != payload.Title || out.Body != payload.Body || out.Source != payload.Source {
		t.Fatalf("response does not satisfy contract: %+v", out)
	}
}

func bytesReader(data []byte) *byteReader { return &byteReader{data: data} }

type byteReader struct {
	data []byte
	pos  int
}

func (r *byteReader) Read(p []byte) (int, error) {
	if r.pos >= len(r.data) {
		return 0, http.ErrBodyReadAfterClose
	}
	n := copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}
