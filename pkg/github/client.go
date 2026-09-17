package github

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Client is a small GitHub REST client for the Factory Lab.
type Client struct {
	Token   string
	Owner   string
	Repo    string
	BaseURL string
	HTTP    *http.Client
}

type Request struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Source string `json:"source,omitempty"`
}

type Issue struct {
	Number int    `json:"number"`
	HTMLURL string `json:"html_url"`
	Title  string `json:"title"`
}

func New(token, owner, repo string) *Client {
	return &Client{
		Token: token, Owner: owner, Repo: repo,
		BaseURL: "https://api.github.com",
		HTTP: http.DefaultClient,
	}
}

func (c *Client) CreateRequest(ctx context.Context, req Request) (Issue, error) {
	if strings.TrimSpace(req.Title) == "" || strings.TrimSpace(req.Body) == "" {
		return Issue{}, fmt.Errorf("title and body are required")
	}

	payload := struct {
		Title  string   `json:"title"`
		Body   string   `json:"body"`
		Labels []string `json:"labels"`
	}{req.Title, req.Body, []string{"prototype", "request"}}

	data, err := json.Marshal(payload)
	if err != nil {
		return Issue{}, err
	}

	url := fmt.Sprintf("%s/repos/%s/%s/issues", strings.TrimRight(c.BaseURL, "/"), c.Owner, c.Repo)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return Issue{}, err
	}
	httpReq.Header.Set("Accept", "application/vnd.github+json")
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if c.Token != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.HTTP.Do(httpReq)
	if err != nil {
		return Issue{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return Issue{}, fmt.Errorf("github API: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return Issue{}, err
	}
	return issue, nil
}
