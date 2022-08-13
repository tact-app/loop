package loom

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	UserWorkspaceMemberships operation = "userWorkspaceMemberships"
)

type operation string

//go:embed operations/*.gql
var operations embed.FS

func NewClient(endpoint, token string) (*Client, error) {
	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{
		Name:  "connect.sid",
		Value: token,
	})

	cl := new(http.Client)
	return &Client{c: cl, r: req}, nil
}

type Client struct {
	c *http.Client
	r *http.Request
}

func (c Client) Do(
	ctx context.Context,
	operation operation,
	vars map[string]interface{},
) error {
	file, err := operations.Open(fmt.Sprintf("operations/%s.gql", operation))
	if err != nil {
		return err
	}

	query, err := io.ReadAll(file)
	if err != nil {
		return err
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
		return err
	}
	req := c.r.Clone(ctx)
	req.Body = io.NopCloser(body)

	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer func() { _, _ = io.Copy(io.Discard, resp.Body) }()
	var r map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}
	return nil
}
