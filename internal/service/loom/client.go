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

//go:embed ops/*.gql
var operations embed.FS

type Operation string

func (op Operation) Name() string {
	name := string(op)
	if alias, has := aliases[name]; has {
		return alias
	}
	return name
}

const (
	Workspaces     Operation = "get-workspaces"
	ArchivedSpaces Operation = "get-archived-spaces"
	PrivateSpaces  Operation = "get-private-spaces"
	PublicSpaces   Operation = "get-public-spaces"
	Folders        Operation = "get-folders"
)

var aliases = map[string]string{
	"get-workspaces":      "userWorkspaceMemberships",
	"get-archived-spaces": "GetWorkspaceArchivedSpaces",
	"get-private-spaces":  "GetMyClosedSpaceMemberships",
	"get-public-spaces":   "getOpenSpaces",
	"get-folders":         "GetPublishedFolders",
}

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

func (c *Client) Do(ctx context.Context, op Operation, vars Vars, out Pointer) error {
	file, err := operations.Open(fmt.Sprintf("ops/%s.gql", op))
	if err != nil {
		return fmt.Errorf("unknown operation %s: %w", op, err)
	}
	defer xio.WillClose(file, unsafe.Ignore)

	query, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("read operation %s: %w", op, err)
	}

	payload := struct {
		OperationName string `json:"operationName"`
		Query         string `json:"query"`
		Variables     Vars   `json:"variables"`
	}{
		OperationName: op.Name(),
		Query:         string(query),
		Variables:     vars,
	}
	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(payload); err != nil {
		return fmt.Errorf("encode operation body %s: %w", op, err)
	}
	req := c.r.Clone(ctx)
	req.Body = io.NopCloser(body)

	resp, err := c.c.Do(req)
	if err != nil {
		return fmt.Errorf("execute operation %s: %w", op, err)
	}
	defer xio.WillDiscard(resp.Body, unsafe.Ignore)

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return fmt.Errorf("decode response %s: %w", op, err)
	}
	return nil
}
