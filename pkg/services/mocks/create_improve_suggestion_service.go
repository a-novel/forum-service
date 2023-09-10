// Code generated by mockery v2.20.0. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	models "github.com/a-novel/forum-service/pkg/models"
	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
)

// CreateImproveSuggestionService is an autogenerated mock type for the CreateImproveSuggestionService type
type CreateImproveSuggestionService struct {
	mock.Mock
}

type CreateImproveSuggestionService_Expecter struct {
	mock *mock.Mock
}

func (_m *CreateImproveSuggestionService) EXPECT() *CreateImproveSuggestionService_Expecter {
	return &CreateImproveSuggestionService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, tokenRaw, suggestion, id, now
func (_m *CreateImproveSuggestionService) Create(ctx context.Context, tokenRaw string, suggestion *models.ImproveSuggestionForm, id uuid.UUID, now time.Time) (*models.ImproveSuggestion, error) {
	ret := _m.Called(ctx, tokenRaw, suggestion, id, now)

	var r0 *models.ImproveSuggestion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *models.ImproveSuggestionForm, uuid.UUID, time.Time) (*models.ImproveSuggestion, error)); ok {
		return rf(ctx, tokenRaw, suggestion, id, now)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *models.ImproveSuggestionForm, uuid.UUID, time.Time) *models.ImproveSuggestion); ok {
		r0 = rf(ctx, tokenRaw, suggestion, id, now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ImproveSuggestion)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *models.ImproveSuggestionForm, uuid.UUID, time.Time) error); ok {
		r1 = rf(ctx, tokenRaw, suggestion, id, now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateImproveSuggestionService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type CreateImproveSuggestionService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - tokenRaw string
//   - suggestion *models.ImproveSuggestionForm
//   - id uuid.UUID
//   - now time.Time
func (_e *CreateImproveSuggestionService_Expecter) Create(ctx interface{}, tokenRaw interface{}, suggestion interface{}, id interface{}, now interface{}) *CreateImproveSuggestionService_Create_Call {
	return &CreateImproveSuggestionService_Create_Call{Call: _e.mock.On("Create", ctx, tokenRaw, suggestion, id, now)}
}

func (_c *CreateImproveSuggestionService_Create_Call) Run(run func(ctx context.Context, tokenRaw string, suggestion *models.ImproveSuggestionForm, id uuid.UUID, now time.Time)) *CreateImproveSuggestionService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*models.ImproveSuggestionForm), args[3].(uuid.UUID), args[4].(time.Time))
	})
	return _c
}

func (_c *CreateImproveSuggestionService_Create_Call) Return(_a0 *models.ImproveSuggestion, _a1 error) *CreateImproveSuggestionService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CreateImproveSuggestionService_Create_Call) RunAndReturn(run func(context.Context, string, *models.ImproveSuggestionForm, uuid.UUID, time.Time) (*models.ImproveSuggestion, error)) *CreateImproveSuggestionService_Create_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCreateImproveSuggestionService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCreateImproveSuggestionService creates a new instance of CreateImproveSuggestionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCreateImproveSuggestionService(t mockConstructorTestingTNewCreateImproveSuggestionService) *CreateImproveSuggestionService {
	mock := &CreateImproveSuggestionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
