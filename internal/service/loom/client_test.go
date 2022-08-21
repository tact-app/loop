package loom

import (
	"context"
	"os"

	"go.octolab.org/tact/loop/internal/service/loom/dto"
)

func ExampleClient_Do() {
	var resp dto.UserWorkspaceMemberships
	cl, _ := NewClient("https://www.loom.com/graphql", os.Getenv("LOOM_TOKEN"))
	_ = cl.Do(context.Background(), UserWorkspaceMemberships, nil, &resp)
	_ = resp

	// output:
	//
}
