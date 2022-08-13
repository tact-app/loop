package loom

import (
	"context"
	"os"
)

func ExampleClient_Do() {
	cl, _ := NewClient("https://www.loom.com/graphql", os.Getenv("LOOM_TOKEN"))
	_ = cl.Do(context.Background(), UserWorkspaceMemberships, nil)

	// output:
	//
}
