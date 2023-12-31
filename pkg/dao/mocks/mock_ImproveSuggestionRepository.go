// Code generated by mockery v2.33.2. DO NOT EDIT.

package daomocks

import (
	context "context"

	dao "github.com/a-novel/forum-service/pkg/dao"
	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
)

// ImproveSuggestionRepository is an autogenerated mock type for the ImproveSuggestionRepository type
type ImproveSuggestionRepository struct {
	mock.Mock
}

type ImproveSuggestionRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ImproveSuggestionRepository) EXPECT() *ImproveSuggestionRepository_Expecter {
	return &ImproveSuggestionRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, data, userID, sourceID, id, now
func (_m *ImproveSuggestionRepository) Create(ctx context.Context, data *dao.ImproveSuggestionModelCore, userID uuid.UUID, sourceID uuid.UUID, id uuid.UUID, now time.Time) (*dao.ImproveSuggestionModel, error) {
	ret := _m.Called(ctx, data, userID, sourceID, id, now)

	var r0 *dao.ImproveSuggestionModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dao.ImproveSuggestionModelCore, uuid.UUID, uuid.UUID, uuid.UUID, time.Time) (*dao.ImproveSuggestionModel, error)); ok {
		return rf(ctx, data, userID, sourceID, id, now)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dao.ImproveSuggestionModelCore, uuid.UUID, uuid.UUID, uuid.UUID, time.Time) *dao.ImproveSuggestionModel); ok {
		r0 = rf(ctx, data, userID, sourceID, id, now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.ImproveSuggestionModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dao.ImproveSuggestionModelCore, uuid.UUID, uuid.UUID, uuid.UUID, time.Time) error); ok {
		r1 = rf(ctx, data, userID, sourceID, id, now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImproveSuggestionRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ImproveSuggestionRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - data *dao.ImproveSuggestionModelCore
//   - userID uuid.UUID
//   - sourceID uuid.UUID
//   - id uuid.UUID
//   - now time.Time
func (_e *ImproveSuggestionRepository_Expecter) Create(ctx interface{}, data interface{}, userID interface{}, sourceID interface{}, id interface{}, now interface{}) *ImproveSuggestionRepository_Create_Call {
	return &ImproveSuggestionRepository_Create_Call{Call: _e.mock.On("Create", ctx, data, userID, sourceID, id, now)}
}

func (_c *ImproveSuggestionRepository_Create_Call) Run(run func(ctx context.Context, data *dao.ImproveSuggestionModelCore, userID uuid.UUID, sourceID uuid.UUID, id uuid.UUID, now time.Time)) *ImproveSuggestionRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*dao.ImproveSuggestionModelCore), args[2].(uuid.UUID), args[3].(uuid.UUID), args[4].(uuid.UUID), args[5].(time.Time))
	})
	return _c
}

func (_c *ImproveSuggestionRepository_Create_Call) Return(_a0 *dao.ImproveSuggestionModel, _a1 error) *ImproveSuggestionRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ImproveSuggestionRepository_Create_Call) RunAndReturn(run func(context.Context, *dao.ImproveSuggestionModelCore, uuid.UUID, uuid.UUID, uuid.UUID, time.Time) (*dao.ImproveSuggestionModel, error)) *ImproveSuggestionRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ImproveSuggestionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImproveSuggestionRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ImproveSuggestionRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *ImproveSuggestionRepository_Expecter) Delete(ctx interface{}, id interface{}) *ImproveSuggestionRepository_Delete_Call {
	return &ImproveSuggestionRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *ImproveSuggestionRepository_Delete_Call) Run(run func(ctx context.Context, id uuid.UUID)) *ImproveSuggestionRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *ImproveSuggestionRepository_Delete_Call) Return(_a0 error) *ImproveSuggestionRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ImproveSuggestionRepository_Delete_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *ImproveSuggestionRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *ImproveSuggestionRepository) Get(ctx context.Context, id uuid.UUID) (*dao.ImproveSuggestionModel, error) {
	ret := _m.Called(ctx, id)

	var r0 *dao.ImproveSuggestionModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*dao.ImproveSuggestionModel, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *dao.ImproveSuggestionModel); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.ImproveSuggestionModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImproveSuggestionRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ImproveSuggestionRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *ImproveSuggestionRepository_Expecter) Get(ctx interface{}, id interface{}) *ImproveSuggestionRepository_Get_Call {
	return &ImproveSuggestionRepository_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *ImproveSuggestionRepository_Get_Call) Run(run func(ctx context.Context, id uuid.UUID)) *ImproveSuggestionRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *ImproveSuggestionRepository_Get_Call) Return(_a0 *dao.ImproveSuggestionModel, _a1 error) *ImproveSuggestionRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ImproveSuggestionRepository_Get_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*dao.ImproveSuggestionModel, error)) *ImproveSuggestionRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, ids
func (_m *ImproveSuggestionRepository) List(ctx context.Context, ids []uuid.UUID) ([]*dao.ImproveSuggestionModel, error) {
	ret := _m.Called(ctx, ids)

	var r0 []*dao.ImproveSuggestionModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) ([]*dao.ImproveSuggestionModel, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) []*dao.ImproveSuggestionModel); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dao.ImproveSuggestionModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uuid.UUID) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImproveSuggestionRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ImproveSuggestionRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []uuid.UUID
func (_e *ImproveSuggestionRepository_Expecter) List(ctx interface{}, ids interface{}) *ImproveSuggestionRepository_List_Call {
	return &ImproveSuggestionRepository_List_Call{Call: _e.mock.On("List", ctx, ids)}
}

func (_c *ImproveSuggestionRepository_List_Call) Run(run func(ctx context.Context, ids []uuid.UUID)) *ImproveSuggestionRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uuid.UUID))
	})
	return _c
}

func (_c *ImproveSuggestionRepository_List_Call) Return(_a0 []*dao.ImproveSuggestionModel, _a1 error) *ImproveSuggestionRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ImproveSuggestionRepository_List_Call) RunAndReturn(run func(context.Context, []uuid.UUID) ([]*dao.ImproveSuggestionModel, error)) *ImproveSuggestionRepository_List_Call {
	_c.Call.Return(run)
	return _c
}

// Search provides a mock function with given fields: ctx, query, limit, offset
func (_m *ImproveSuggestionRepository) Search(ctx context.Context, query dao.ImproveSuggestionSearchQuery, limit int, offset int) ([]*dao.ImproveSuggestionModel, int, error) {
	ret := _m.Called(ctx, query, limit, offset)

	var r0 []*dao.ImproveSuggestionModel
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, dao.ImproveSuggestionSearchQuery, int, int) ([]*dao.ImproveSuggestionModel, int, error)); ok {
		return rf(ctx, query, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dao.ImproveSuggestionSearchQuery, int, int) []*dao.ImproveSuggestionModel); ok {
		r0 = rf(ctx, query, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dao.ImproveSuggestionModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dao.ImproveSuggestionSearchQuery, int, int) int); ok {
		r1 = rf(ctx, query, limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, dao.ImproveSuggestionSearchQuery, int, int) error); ok {
		r2 = rf(ctx, query, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ImproveSuggestionRepository_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type ImproveSuggestionRepository_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//   - ctx context.Context
//   - query dao.ImproveSuggestionSearchQuery
//   - limit int
//   - offset int
func (_e *ImproveSuggestionRepository_Expecter) Search(ctx interface{}, query interface{}, limit interface{}, offset interface{}) *ImproveSuggestionRepository_Search_Call {
	return &ImproveSuggestionRepository_Search_Call{Call: _e.mock.On("Search", ctx, query, limit, offset)}
}

func (_c *ImproveSuggestionRepository_Search_Call) Run(run func(ctx context.Context, query dao.ImproveSuggestionSearchQuery, limit int, offset int)) *ImproveSuggestionRepository_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dao.ImproveSuggestionSearchQuery), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *ImproveSuggestionRepository_Search_Call) Return(_a0 []*dao.ImproveSuggestionModel, _a1 int, _a2 error) *ImproveSuggestionRepository_Search_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ImproveSuggestionRepository_Search_Call) RunAndReturn(run func(context.Context, dao.ImproveSuggestionSearchQuery, int, int) ([]*dao.ImproveSuggestionModel, int, error)) *ImproveSuggestionRepository_Search_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, data, id, now
func (_m *ImproveSuggestionRepository) Update(ctx context.Context, data *dao.ImproveSuggestionModelCore, id uuid.UUID, now time.Time) (*dao.ImproveSuggestionModel, error) {
	ret := _m.Called(ctx, data, id, now)

	var r0 *dao.ImproveSuggestionModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dao.ImproveSuggestionModelCore, uuid.UUID, time.Time) (*dao.ImproveSuggestionModel, error)); ok {
		return rf(ctx, data, id, now)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dao.ImproveSuggestionModelCore, uuid.UUID, time.Time) *dao.ImproveSuggestionModel); ok {
		r0 = rf(ctx, data, id, now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.ImproveSuggestionModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dao.ImproveSuggestionModelCore, uuid.UUID, time.Time) error); ok {
		r1 = rf(ctx, data, id, now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImproveSuggestionRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ImproveSuggestionRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - data *dao.ImproveSuggestionModelCore
//   - id uuid.UUID
//   - now time.Time
func (_e *ImproveSuggestionRepository_Expecter) Update(ctx interface{}, data interface{}, id interface{}, now interface{}) *ImproveSuggestionRepository_Update_Call {
	return &ImproveSuggestionRepository_Update_Call{Call: _e.mock.On("Update", ctx, data, id, now)}
}

func (_c *ImproveSuggestionRepository_Update_Call) Run(run func(ctx context.Context, data *dao.ImproveSuggestionModelCore, id uuid.UUID, now time.Time)) *ImproveSuggestionRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*dao.ImproveSuggestionModelCore), args[2].(uuid.UUID), args[3].(time.Time))
	})
	return _c
}

func (_c *ImproveSuggestionRepository_Update_Call) Return(_a0 *dao.ImproveSuggestionModel, _a1 error) *ImproveSuggestionRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ImproveSuggestionRepository_Update_Call) RunAndReturn(run func(context.Context, *dao.ImproveSuggestionModelCore, uuid.UUID, time.Time) (*dao.ImproveSuggestionModel, error)) *ImproveSuggestionRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateVotes provides a mock function with given fields: ctx, id, upVotes, downVotes
func (_m *ImproveSuggestionRepository) UpdateVotes(ctx context.Context, id uuid.UUID, upVotes int, downVotes int) error {
	ret := _m.Called(ctx, id, upVotes, downVotes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int, int) error); ok {
		r0 = rf(ctx, id, upVotes, downVotes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImproveSuggestionRepository_UpdateVotes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateVotes'
type ImproveSuggestionRepository_UpdateVotes_Call struct {
	*mock.Call
}

// UpdateVotes is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
//   - upVotes int
//   - downVotes int
func (_e *ImproveSuggestionRepository_Expecter) UpdateVotes(ctx interface{}, id interface{}, upVotes interface{}, downVotes interface{}) *ImproveSuggestionRepository_UpdateVotes_Call {
	return &ImproveSuggestionRepository_UpdateVotes_Call{Call: _e.mock.On("UpdateVotes", ctx, id, upVotes, downVotes)}
}

func (_c *ImproveSuggestionRepository_UpdateVotes_Call) Run(run func(ctx context.Context, id uuid.UUID, upVotes int, downVotes int)) *ImproveSuggestionRepository_UpdateVotes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *ImproveSuggestionRepository_UpdateVotes_Call) Return(_a0 error) *ImproveSuggestionRepository_UpdateVotes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ImproveSuggestionRepository_UpdateVotes_Call) RunAndReturn(run func(context.Context, uuid.UUID, int, int) error) *ImproveSuggestionRepository_UpdateVotes_Call {
	_c.Call.Return(run)
	return _c
}

// Validate provides a mock function with given fields: ctx, validated, id
func (_m *ImproveSuggestionRepository) Validate(ctx context.Context, validated bool, id uuid.UUID) (*dao.ImproveSuggestionModel, error) {
	ret := _m.Called(ctx, validated, id)

	var r0 *dao.ImproveSuggestionModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool, uuid.UUID) (*dao.ImproveSuggestionModel, error)); ok {
		return rf(ctx, validated, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool, uuid.UUID) *dao.ImproveSuggestionModel); ok {
		r0 = rf(ctx, validated, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.ImproveSuggestionModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool, uuid.UUID) error); ok {
		r1 = rf(ctx, validated, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImproveSuggestionRepository_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type ImproveSuggestionRepository_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - ctx context.Context
//   - validated bool
//   - id uuid.UUID
func (_e *ImproveSuggestionRepository_Expecter) Validate(ctx interface{}, validated interface{}, id interface{}) *ImproveSuggestionRepository_Validate_Call {
	return &ImproveSuggestionRepository_Validate_Call{Call: _e.mock.On("Validate", ctx, validated, id)}
}

func (_c *ImproveSuggestionRepository_Validate_Call) Run(run func(ctx context.Context, validated bool, id uuid.UUID)) *ImproveSuggestionRepository_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool), args[2].(uuid.UUID))
	})
	return _c
}

func (_c *ImproveSuggestionRepository_Validate_Call) Return(_a0 *dao.ImproveSuggestionModel, _a1 error) *ImproveSuggestionRepository_Validate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ImproveSuggestionRepository_Validate_Call) RunAndReturn(run func(context.Context, bool, uuid.UUID) (*dao.ImproveSuggestionModel, error)) *ImproveSuggestionRepository_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewImproveSuggestionRepository creates a new instance of ImproveSuggestionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewImproveSuggestionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ImproveSuggestionRepository {
	mock := &ImproveSuggestionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
