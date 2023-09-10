package framework

import (
	"context"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/go-framework/client"
	"net/http"
	"net/url"
)

type Client interface {
	VoteImproveRequest(ctx context.Context, form models.UpdateImproveRequestVotesForm) error
	VoteImproveSuggestion(ctx context.Context, form models.UpdateImproveSuggestionVotesForm) error
	GetImproveRequest(ctx context.Context, query models.GetImproveRequestQuery) (*models.ImproveRequestPreview, error)
	GetImproveSuggestion(ctx context.Context, query models.GetImproveSuggestionQuery) (*models.ImproveSuggestion, error)
}

type clientImpl struct {
	url *url.URL
}

func NewClient(url *url.URL) Client {
	return &clientImpl{url: url}
}

func (a *clientImpl) VoteImproveRequest(ctx context.Context, form models.UpdateImproveRequestVotesForm) error {
	return client.MakeHTTPCall(ctx, client.HTTPCallConfig{
		Path:            a.url.JoinPath("/improve-request/vote"),
		Method:          http.MethodPost,
		Body:            form,
		SuccessStatuses: []int{http.StatusNoContent},
		Client:          http.DefaultClient,
	}, nil)
}

func (a *clientImpl) VoteImproveSuggestion(ctx context.Context, form models.UpdateImproveSuggestionVotesForm) error {
	return client.MakeHTTPCall(ctx, client.HTTPCallConfig{
		Path:            a.url.JoinPath("/improve-suggestion/vote"),
		Method:          http.MethodPost,
		Body:            form,
		SuccessStatuses: []int{http.StatusNoContent},
		Client:          http.DefaultClient,
	}, nil)
}

func (a *clientImpl) GetImproveRequest(ctx context.Context, query models.GetImproveRequestQuery) (*models.ImproveRequestPreview, error) {
	pathQuery := new(url.Values)
	pathQuery.Set("id", string(query.ID))

	path := a.url.JoinPath("/improve-request")
	path.RawQuery = pathQuery.Encode()

	output := new(models.ImproveRequestPreview)

	return output, client.MakeHTTPCall(ctx, client.HTTPCallConfig{
		Path:            path,
		Method:          http.MethodGet,
		SuccessStatuses: []int{http.StatusOK},
		Client:          http.DefaultClient,
	}, output)
}

func (a *clientImpl) GetImproveSuggestion(ctx context.Context, query models.GetImproveSuggestionQuery) (*models.ImproveSuggestion, error) {
	pathQuery := new(url.Values)
	pathQuery.Set("id", string(query.ID))

	path := a.url.JoinPath("/improve-suggestion")
	path.RawQuery = pathQuery.Encode()

	output := new(models.ImproveSuggestion)

	return output, client.MakeHTTPCall(ctx, client.HTTPCallConfig{
		Path:            path,
		Method:          http.MethodGet,
		SuccessStatuses: []int{http.StatusOK},
		Client:          http.DefaultClient,
	}, output)
}
