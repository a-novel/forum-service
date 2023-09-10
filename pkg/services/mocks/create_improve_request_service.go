// Code generated by mockery v2.20.0. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	models "github.com/a-novel/forum-service/pkg/models"
	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
)

// CreateImproveRequestService is an autogenerated mock type for the CreateImproveRequestService type
type CreateImproveRequestService struct {
	mock.Mock
}

type CreateImproveRequestService_Expecter struct {
	mock *mock.Mock
}

func (_m *CreateImproveRequestService) EXPECT() *CreateImproveRequestService_Expecter {
	return &CreateImproveRequestService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, tokenRaw, title, content, sourceID, id, now
func (_m *CreateImproveRequestService) Create(ctx context.Context, tokenRaw string, title string, content string, sourceID uuid.UUID, id uuid.UUID, now time.Time) (*models.ImproveRequestPreview, error) {
	ret := _m.Called(ctx, tokenRaw, title, content, sourceID, id, now)

	var r0 *models.ImproveRequestPreview
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, uuid.UUID, uuid.UUID, time.Time) (*models.ImproveRequestPreview, error)); ok {
		return rf(ctx, tokenRaw, title, content, sourceID, id, now)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, uuid.UUID, uuid.UUID, time.Time) *models.ImproveRequestPreview); ok {
		r0 = rf(ctx, tokenRaw, title, content, sourceID, id, now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ImproveRequestPreview)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, uuid.UUID, uuid.UUID, time.Time) error); ok {
		r1 = rf(ctx, tokenRaw, title, content, sourceID, id, now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateImproveRequestService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type CreateImproveRequestService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - tokenRaw string
//   - title string
//   - content string
//   - sourceID uuid.UUID
//   - id uuid.UUID
//   - now time.Time
func (_e *CreateImproveRequestService_Expecter) Create(ctx interface{}, tokenRaw interface{}, title interface{}, content interface{}, sourceID interface{}, id interface{}, now interface{}) *CreateImproveRequestService_Create_Call {
	return &CreateImproveRequestService_Create_Call{Call: _e.mock.On("Create", ctx, tokenRaw, title, content, sourceID, id, now)}
}

func (_c *CreateImproveRequestService_Create_Call) Run(run func(ctx context.Context, tokenRaw string, title string, content string, sourceID uuid.UUID, id uuid.UUID, now time.Time)) *CreateImproveRequestService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(uuid.UUID), args[5].(uuid.UUID), args[6].(time.Time))
	})
	return _c
}

func (_c *CreateImproveRequestService_Create_Call) Return(_a0 *models.ImproveRequestPreview, _a1 error) *CreateImproveRequestService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CreateImproveRequestService_Create_Call) RunAndReturn(run func(context.Context, string, string, string, uuid.UUID, uuid.UUID, time.Time) (*models.ImproveRequestPreview, error)) *CreateImproveRequestService_Create_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCreateImproveRequestService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCreateImproveRequestService creates a new instance of CreateImproveRequestService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCreateImproveRequestService(t mockConstructorTestingTNewCreateImproveRequestService) *CreateImproveRequestService {
	mock := &CreateImproveRequestService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
