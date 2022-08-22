package loom

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.octolab.org/unsafe"

	xio "go.octolab.org/tact/loop/internal/pkg/io"
)

//go:embed operations/*.gql
var operations embed.FS

type operation string

const (
	UserWorkspaceMemberships operation = "userWorkspaceMemberships"
)

func NewClient(client HttpClient, endpoint, token string) (*Client, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("init http request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{
		Name:  "connect.sid",
		Value: token,
	})

	return &Client{c: client, r: req}, nil
}

type Client struct {
	c HttpClient
	r *http.Request
}

func (c Client) Do(
	ctx context.Context,
	operation operation,
	vars map[string]interface{},
	out interface{},
) error {
	file, err := operations.Open(fmt.Sprintf("operations/%s.gql", operation))
	if err != nil {
		return fmt.Errorf("unknown operation %s: %w", operation, err)
	}
	defer xio.WillClose(file, unsafe.Ignore)

	query, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("read operation %s: %w", operation, err)
	}

	payload := struct {
		OperationName string                 `json:"operationName"`
		Query         string                 `json:"query"`
		Variables     map[string]interface{} `json:"variables"`
	}{
		OperationName: string(operation),
		Query:         string(query),
		Variables:     vars,
	}
	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(payload); err != nil {
		return fmt.Errorf("encode operation body %s: %w", operation, err)
	}
	req := c.r.Clone(ctx)
	req.Body = io.NopCloser(body)

	resp, err := c.c.Do(req)
	if err != nil {
		return fmt.Errorf("execute operation %s: %w", operation, err)
	}
	defer xio.WillDiscard(resp.Body, unsafe.Ignore)

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return fmt.Errorf("decode response %s: %w", operation, err)
	}
	return nil
}
