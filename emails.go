package gitee

import (
	"context"
	"net/http"
)

// EmailsService
// Docs: https://gitee.com/api/v5/swagger#/getV5Emails
type EmailsService service

type Email struct {
	Email string   `json:"email"`
	State string   `json:"state"`
	Scope []string `json:"scope"`
}

func (s *EmailsService) List(ctx context.Context) ([]*Email, error) {
	const apiEndpoint = "/api/v5/emails"
	var respBody []*Email
	if err := s.client.InvokeByCredential(ctx, http.MethodGet, apiEndpoint, nil, &respBody); err != nil {
		return nil, err
	}
	return respBody, nil
}
