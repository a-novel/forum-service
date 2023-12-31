// Code generated by mockery v2.33.2. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	models "github.com/a-novel/forum-service/pkg/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ListImproveSuggestionsService is an autogenerated mock type for the ListImproveSuggestionsService type
type ListImproveSuggestionsService struct {
	mock.Mock
}

type ListImproveSuggestionsService_Expecter struct {
	mock *mock.Mock
}

func (_m *ListImproveSuggestionsService) EXPECT() *ListImproveSuggestionsService_Expecter {
	return &ListImproveSuggestionsService_Expecter{mock: &_m.Mock}
}

// List provides a mock function with given fields: ctx, ids
func (_m *ListImproveSuggestionsService) List(ctx context.Context, ids []uuid.UUID) ([]*models.ImproveSuggestion, error) {
	ret := _m.Called(ctx, ids)

	var r0 []*models.ImproveSuggestion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) ([]*models.ImproveSuggestion, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) []*models.ImproveSuggestion); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ImproveSuggestion)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uuid.UUID) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListImproveSuggestionsService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ListImproveSuggestionsService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []uuid.UUID
func (_e *ListImproveSuggestionsService_Expecter) List(ctx interface{}, ids interface{}) *ListImproveSuggestionsService_List_Call {
	return &ListImproveSuggestionsService_List_Call{Call: _e.mock.On("List", ctx, ids)}
}

func (_c *ListImproveSuggestionsService_List_Call) Run(run func(ctx context.Context, ids []uuid.UUID)) *ListImproveSuggestionsService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uuid.UUID))
	})
	return _c
}

func (_c *ListImproveSuggestionsService_List_Call) Return(_a0 []*models.ImproveSuggestion, _a1 error) *ListImproveSuggestionsService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ListImproveSuggestionsService_List_Call) RunAndReturn(run func(context.Context, []uuid.UUID) ([]*models.ImproveSuggestion, error)) *ListImproveSuggestionsService_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewListImproveSuggestionsService creates a new instance of ListImproveSuggestionsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewListImproveSuggestionsService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ListImproveSuggestionsService {
	mock := &ListImproveSuggestionsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
