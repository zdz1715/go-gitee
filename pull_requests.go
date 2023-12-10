package gitee

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// PullRequestsService
// GitLab API docs: https://gitee.com/api/v5/swagger#/getV5ReposOwnerRepoPulls
type PullRequestsService service

type PullRequest struct {
	ID                int           `json:"id"`
	Assignees         []User        `json:"assignees"`
	AssigneesNumber   int           `json:"assignees_number"`
	Base              PullBranch    `json:"base"`
	CanMergeCheck     bool          `json:"can_merge_check"`
	CommentsURL       string        `json:"comments_url"`
	CommitsURL        string        `json:"commits_url"`
	CreatedAt         *time.Time    `json:"created_at"`
	DiffURL           string        `json:"diff_url"`
	Draft             bool          `json:"draft"`
	Head              PullBranch    `json:"head"`
	HtmlURL           string        `json:"html_url"`
	IssueURL          string        `json:"issue_url"`
	Labels            []interface{} `json:"labels"`
	Locked            bool          `json:"locked"`
	Mergeable         bool          `json:"mergeable"`
	Number            int           `json:"number"`
	PatchUR           string        `json:"patch_url"`
	RefPullRequests   []interface{} `json:"ref_pull_requests"`
	ReviewCommentURL  string        `json:"review_comment_url"`
	ReviewCommentsURL string        `json:"review_comments_url"`
	State             string        `json:"state"`
	Testers           []User        `json:"testers"`
	TestersNumber     int           `json:"testers_number"`
	Title             string        `json:"title"`
	UpdatedAt         *time.Time    `json:"updated_at"`
	URL               string        `json:"url"`
	User              User          `json:"user"`
}

type PullBranch struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	Repo  Repo   `json:"repo"`
	Sha   string `json:"sha"`
	User  User   `json:"user"`
}

// CreatePullRequestOptions represents the available CreatePullRequest()
// options.
type CreatePullRequestOptions struct {
	Title                 *string `json:"title,omitempty" query:"title"`
	Description           *string `json:"description,omitempty" query:"description"`
	Head                  *string `json:"head,omitempty" query:"head"`
	Base                  *string `json:"base,omitempty" query:"base"`
	Body                  *string `json:"body,omitempty" query:"body"`
	MilestoneNumber       *int    `json:"milestone_number,omitempty" query:"milestone_number"`
	Labels                *string `json:"labels,omitempty" query:"labels"`
	Issue                 *string `json:"issue,omitempty" query:"issue"`
	Assignees             *string `json:"assignees,omitempty" query:"assignees"`
	Testers               *string `json:"testers,omitempty" query:"testers"`
	AssigneeNumber        *int    `json:"assignee_number,omitempty" query:"assignee_number"`
	TestersNumber         *int    `json:"testers_number,omitempty" query:"testers_number"`
	RefPullRequestNumbers *string `json:"ref_pull_request_numbers,omitempty" query:"ref_pull_request_numbers"`
	PruneSourceBranch     *bool   `json:"prune_source_branch,omitempty" query:"prune_source_branch"`
	CloseRelatedIssue     *bool   `json:"close_related_issue,omitempty" query:"close_related_issue"`
	Draft                 *bool   `json:"draft,omitempty" query:"draft"`
	Squash                *bool   `json:"squash,omitempty" query:"squash"`
}

// CreatePullRequest
// Gitee API docs:
// https://gitee.com/api/v5/swagger#/postV5ReposOwnerRepoPulls
func (s *PullRequestsService) CreatePullRequest(ctx context.Context, fullName string, opts *CreatePullRequestOptions) (*PullRequest, error) {
	apiEndpoint := fmt.Sprintf("/api/v5/repos/%s/pulls", fullName)
	var v PullRequest
	if err := s.client.InvokeByCredential(ctx, http.MethodPost, apiEndpoint, opts, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// AcceptPullRequestOptions represents the available AcceptPullRequest()
// options.

type AcceptPullRequestOptions struct {
	Title             *string      `json:"title,omitempty" query:"title"`
	Description       *string      `json:"description,omitempty" query:"description"`
	MergeMethod       *MergeMethod `json:"merge_method,omitempty" query:"merge_method"`
	PruneSourceBranch *bool        `json:"prune_source_branch,omitempty" query:"prune_source_branch"`
}

type AcceptPullRequest struct {
	Merged  bool   `json:"merged"`
	Message string `json:"message"`
	Sha     string `json:"sha"`
}

// AcceptPullRequest
// Gitee API docs:
// https://gitee.com/api/v5/swagger#/putV5ReposOwnerRepoPullsNumberMerge
func (s *PullRequestsService) AcceptPullRequest(ctx context.Context, fullName string, iid int, opts *AcceptPullRequestOptions) (*AcceptPullRequest, error) {
	apiEndpoint := fmt.Sprintf("/api/v5/repos/%s/pulls/%d/merge", fullName, iid)
	var v AcceptPullRequest
	if err := s.client.InvokeByCredential(ctx, http.MethodPut, apiEndpoint, opts, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
