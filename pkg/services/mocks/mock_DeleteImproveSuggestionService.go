// Code generated by mockery v2.33.2. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// DeleteImproveSuggestionService is an autogenerated mock type for the DeleteImproveSuggestionService type
type DeleteImproveSuggestionService struct {
	mock.Mock
}

type DeleteImproveSuggestionService_Expecter struct {
	mock *mock.Mock
}

func (_m *DeleteImproveSuggestionService) EXPECT() *DeleteImproveSuggestionService_Expecter {
	return &DeleteImproveSuggestionService_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, tokenRaw, id
func (_m *DeleteImproveSuggestionService) Delete(ctx context.Context, tokenRaw string, id uuid.UUID) error {
	ret := _m.Called(ctx, tokenRaw, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uuid.UUID) error); ok {
		r0 = rf(ctx, tokenRaw, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteImproveSuggestionService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type DeleteImproveSuggestionService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - tokenRaw string
//   - id uuid.UUID
func (_e *DeleteImproveSuggestionService_Expecter) Delete(ctx interface{}, tokenRaw interface{}, id interface{}) *DeleteImproveSuggestionService_Delete_Call {
	return &DeleteImproveSuggestionService_Delete_Call{Call: _e.mock.On("Delete", ctx, tokenRaw, id)}
}

func (_c *DeleteImproveSuggestionService_Delete_Call) Run(run func(ctx context.Context, tokenRaw string, id uuid.UUID)) *DeleteImproveSuggestionService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uuid.UUID))
	})
	return _c
}

func (_c *DeleteImproveSuggestionService_Delete_Call) Return(_a0 error) *DeleteImproveSuggestionService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeleteImproveSuggestionService_Delete_Call) RunAndReturn(run func(context.Context, string, uuid.UUID) error) *DeleteImproveSuggestionService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// NewDeleteImproveSuggestionService creates a new instance of DeleteImproveSuggestionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeleteImproveSuggestionService(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeleteImproveSuggestionService {
	mock := &DeleteImproveSuggestionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
