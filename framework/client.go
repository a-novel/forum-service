package framework

import (
	"context"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/go-apis"
	"net/http"
	"net/url"
)

type Client interface {
	VoteImproveRequest(ctx context.Context, form models.UpdateImproveRequestVotesForm) error
	VoteImproveSuggestion(ctx context.Context, form models.UpdateImproveSuggestionVotesForm) error
	GetImproveRequest(ctx context.Context, query models.GetImproveRequestQuery) (*models.ImproveRequestPreview, error)
	GetImproveSuggestion(ctx context.Context, query models.GetImproveSuggestionQuery) (*models.ImproveSuggestion, error)

	Ping() error
}

type clientImpl struct {
	url *url.URL
}

func NewClient(url *url.URL) Client {
	return &clientImpl{url: url}
}

func (a *clientImpl) VoteImproveRequest(ctx context.Context, form models.UpdateImproveRequestVotesForm) error {
	return apis.CallHTTP(ctx, apis.CallHTTPConfig{
		Path:            a.url.JoinPath("/improve-request/vote"),
		Method:          http.MethodPost,
		Body:            form,
		SuccessStatuses: []int{http.StatusNoContent},
	}, nil)
}

func (a *clientImpl) VoteImproveSuggestion(ctx context.Context, form models.UpdateImproveSuggestionVotesForm) error {
	return apis.CallHTTP(ctx, apis.CallHTTPConfig{
		Path:            a.url.JoinPath("/improve-suggestion/vote"),
		Method:          http.MethodPost,
		Body:            form,
		SuccessStatuses: []int{http.StatusNoContent},
	}, nil)
}

func (a *clientImpl) GetImproveRequest(ctx context.Context, query models.GetImproveRequestQuery) (*models.ImproveRequestPreview, error) {
	pathQuery := new(url.Values)
	pathQuery.Set("id", string(query.ID))

	path := a.url.JoinPath("/improve-request")
	path.RawQuery = pathQuery.Encode()

	output := new(models.ImproveRequestPreview)

	return output, apis.CallHTTP(ctx, apis.CallHTTPConfig{
		Path:            path,
		Method:          http.MethodGet,
		SuccessStatuses: []int{http.StatusOK},
	}, output)
}

func (a *clientImpl) GetImproveSuggestion(ctx context.Context, query models.GetImproveSuggestionQuery) (*models.ImproveSuggestion, error) {
	pathQuery := new(url.Values)
	pathQuery.Set("id", string(query.ID))

	path := a.url.JoinPath("/improve-suggestion")
	path.RawQuery = pathQuery.Encode()

	output := new(models.ImproveSuggestion)

	return output, apis.CallHTTP(ctx, apis.CallHTTPConfig{
		Path:            path,
		Method:          http.MethodGet,
		SuccessStatuses: []int{http.StatusOK},
	}, output)
}

func (a *clientImpl) Ping() error {
	return apis.CallHTTP(context.Background(), apis.CallHTTPConfig{
		Path:            a.url.JoinPath("/ping"),
		Method:          http.MethodGet,
		SuccessStatuses: []int{http.StatusOK},
	}, nil)
}
