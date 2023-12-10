package gitee

import (
	"context"
	"fmt"
	"testing"

	"github.com/zdz1715/go-utils/goutils"

	"github.com/zdz1715/ghttp"
)

func TestPullRequestService_CreatePullRequest(t *testing.T) {
	client, err := NewClient(testTokenCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(true),
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	p, err := client.PullRequest.CreatePullRequest(context.Background(), "zdzserver/test", &CreatePullRequestOptions{
		Title: goutils.Ptr("test"),
		Head:  goutils.Ptr("master"),
		Base:  goutils.Ptr("111"),
	})

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("PullRequest: %+v\n", p)
}

func TestPullRequestsService_AcceptPullRequest(t *testing.T) {
	client, err := NewClient(testTokenCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(true),
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	p, err := client.PullRequest.AcceptPullRequest(context.Background(), "zdzserver/test", 9, &AcceptPullRequestOptions{})

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("PullRequest: %+v\n", p)
}
