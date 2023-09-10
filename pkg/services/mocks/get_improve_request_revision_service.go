// Code generated by mockery v2.20.0. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	models "github.com/a-novel/forum-service/pkg/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// GetImproveRequestRevisionService is an autogenerated mock type for the GetImproveRequestRevisionService type
type GetImproveRequestRevisionService struct {
	mock.Mock
}

type GetImproveRequestRevisionService_Expecter struct {
	mock *mock.Mock
}

func (_m *GetImproveRequestRevisionService) EXPECT() *GetImproveRequestRevisionService_Expecter {
	return &GetImproveRequestRevisionService_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id
func (_m *GetImproveRequestRevisionService) Get(ctx context.Context, id uuid.UUID) (*models.ImproveRequestRevision, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.ImproveRequestRevision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*models.ImproveRequestRevision, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *models.ImproveRequestRevision); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ImproveRequestRevision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImproveRequestRevisionService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type GetImproveRequestRevisionService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *GetImproveRequestRevisionService_Expecter) Get(ctx interface{}, id interface{}) *GetImproveRequestRevisionService_Get_Call {
	return &GetImproveRequestRevisionService_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *GetImproveRequestRevisionService_Get_Call) Run(run func(ctx context.Context, id uuid.UUID)) *GetImproveRequestRevisionService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *GetImproveRequestRevisionService_Get_Call) Return(_a0 *models.ImproveRequestRevision, _a1 error) *GetImproveRequestRevisionService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GetImproveRequestRevisionService_Get_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*models.ImproveRequestRevision, error)) *GetImproveRequestRevisionService_Get_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewGetImproveRequestRevisionService interface {
	mock.TestingT
	Cleanup(func())
}

// NewGetImproveRequestRevisionService creates a new instance of GetImproveRequestRevisionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGetImproveRequestRevisionService(t mockConstructorTestingTNewGetImproveRequestRevisionService) *GetImproveRequestRevisionService {
	mock := &GetImproveRequestRevisionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
