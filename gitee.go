package gitee

import (
	"context"
	"net/http"
	"strings"

	"github.com/zdz1715/ghttp"
)

type service struct {
	client *Client
}

type Options struct {
	ClientOpts         []ghttp.ClientOption
	InitSkipCredential bool
}

type Client struct {
	cc   *ghttp.Client
	opts *Options

	common service

	OAuth       *OAuthService
	Email       *EmailsService
	PullRequest *PullRequestsService
}

func NewClient(credential Credential, opts *Options) (*Client, error) {
	if opts == nil {
		opts = &Options{}
	}

	clientOptions := []ghttp.ClientOption{
		ghttp.WithEndpoint(CloudEndpoint),
	}

	if len(opts.ClientOpts) > 0 {
		clientOptions = append(clientOptions, opts.ClientOpts...)
	}
	// 覆盖错误
	clientOptions = append(clientOptions, ghttp.WithNot2xxError(&Error{}))

	cc, err := ghttp.NewClient(context.Background(), clientOptions...)
	if err != nil {
		return nil, err
	}

	c := &Client{
		cc:   cc,
		opts: opts,
	}

	c.common.client = c

	c.OAuth = &OAuthService{client: c.common.client}
	c.Email = (*EmailsService)(&c.common)
	c.PullRequest = (*PullRequestsService)(&c.common)

	if opts.InitSkipCredential {
		return c, nil
	}

	if err = c.SetCredential(credential); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) SetCredential(credential Credential) error {
	if credential == nil {
		return ErrCredential
	}

	if err := credential.Valid(); err != nil {
		return err
	}

	if err := c.cc.SetEndpoint(credential.GetEndpoint()); err != nil {
		return err
	}

	if c.OAuth != nil {
		c.OAuth.credential = credential
	}

	return nil
}

func (c *Client) InvokeByCredential(ctx context.Context, method, path string, args interface{}, reply interface{}) error {
	accessToken, err := c.OAuth.GetAccessToken(ctx)
	if err != nil {
		return err
	}

	callOpts, err := c.OAuth.credential.GenerateCallOptions(accessToken)
	if err != nil {
		return err
	}

	return c.Invoke(ctx, method, path, args, reply, callOpts)
}

func (c *Client) Invoke(ctx context.Context, method, path string, args interface{}, reply interface{}, opts ...*ghttp.CallOptions) error {
	callOpts := new(ghttp.CallOptions)
	if len(opts) > 0 && opts[0] != nil {
		callOpts = opts[0]
	}

	if method == http.MethodGet && args != nil {
		callOpts.Query = args
		args = nil
	}

	_, err := c.cc.Invoke(ctx, method, path, args, reply, callOpts)
	return err
}

type Error struct {
	Message  string   `json:"message"`
	Messages []string `json:"messages"`

	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (e *Error) String() string {
	if e.ErrorDescription != "" {
		return e.ErrorDescription
	}
	if e.Error != "" {
		return e.Error
	}
	if e.Message != "" {
		return e.Message
	}
	if len(e.Messages) > 0 {
		return strings.Join(e.Messages, ";")
	}
	return ""
}

func (e *Error) Reset() {
	e.Message = ""
	e.Error = ""
	e.ErrorDescription = ""
	e.Messages = make([]string, 0)
}
