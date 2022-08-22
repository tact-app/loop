package loom

import (
	"context"
	"net/http"
	"os"

	"go.octolab.org/tact/loop/internal/service/loom/dto"
)

func ExampleClient_Do() {
	var resp dto.UserWorkspaceMemberships
	cl, _ := NewClient(new(http.Client), "https://www.loom.com/graphql", os.Getenv("LOOM_TOKEN"))
	_ = cl.Do(context.Background(), UserWorkspaceMemberships, nil, &resp)

	// output:
	//
}
