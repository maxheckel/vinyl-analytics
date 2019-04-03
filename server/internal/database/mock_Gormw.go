// Code generated by mockery v1.0.0. DO NOT EDIT.

package database

import gorm "github.com/jinzhu/gorm"
import mock "github.com/stretchr/testify/mock"
import sql "database/sql"

// MockGormw is an autogenerated mock type for the Gormw type
type MockGormw struct {
	mock.Mock
}

// AddError provides a mock function with given fields: err
func (_m *MockGormw) AddError(err error) error {
	ret := _m.Called(err)

	var r0 error
	if rf, ok := ret.Get(0).(func(error) error); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddForeignKey provides a mock function with given fields: field, dest, onDelete, onUpdate
func (_m *MockGormw) AddForeignKey(field string, dest string, onDelete string, onUpdate string) Gormw {
	ret := _m.Called(field, dest, onDelete, onUpdate)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, string, string, string) Gormw); ok {
		r0 = rf(field, dest, onDelete, onUpdate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// AddIndex provides a mock function with given fields: indexName, column
func (_m *MockGormw) AddIndex(indexName string, column ...string) Gormw {
	_va := make([]interface{}, len(column))
	for _i := range column {
		_va[_i] = column[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, indexName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, ...string) Gormw); ok {
		r0 = rf(indexName, column...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// AddUniqueIndex provides a mock function with given fields: indexName, column
func (_m *MockGormw) AddUniqueIndex(indexName string, column ...string) Gormw {
	_va := make([]interface{}, len(column))
	for _i := range column {
		_va[_i] = column[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, indexName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, ...string) Gormw); ok {
		r0 = rf(indexName, column...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Assign provides a mock function with given fields: attrs
func (_m *MockGormw) Assign(attrs ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, attrs...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(...interface{}) Gormw); ok {
		r0 = rf(attrs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Association provides a mock function with given fields: column
func (_m *MockGormw) Association(column string) *gorm.Association {
	ret := _m.Called(column)

	var r0 *gorm.Association
	if rf, ok := ret.Get(0).(func(string) *gorm.Association); ok {
		r0 = rf(column)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.Association)
		}
	}

	return r0
}

// Attrs provides a mock function with given fields: attrs
func (_m *MockGormw) Attrs(attrs ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, attrs...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(...interface{}) Gormw); ok {
		r0 = rf(attrs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// AutoMigrate provides a mock function with given fields: values
func (_m *MockGormw) AutoMigrate(values ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(...interface{}) Gormw); ok {
		r0 = rf(values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Begin provides a mock function with given fields:
func (_m *MockGormw) Begin() Gormw {
	ret := _m.Called()

	var r0 Gormw
	if rf, ok := ret.Get(0).(func() Gormw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Callback provides a mock function with given fields:
func (_m *MockGormw) Callback() *gorm.Callback {
	ret := _m.Called()

	var r0 *gorm.Callback
	if rf, ok := ret.Get(0).(func() *gorm.Callback); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.Callback)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *MockGormw) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Commit provides a mock function with given fields:
func (_m *MockGormw) Commit() Gormw {
	ret := _m.Called()

	var r0 Gormw
	if rf, ok := ret.Get(0).(func() Gormw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// CommonDB provides a mock function with given fields:
func (_m *MockGormw) CommonDB() gorm.SQLCommon {
	ret := _m.Called()

	var r0 gorm.SQLCommon
	if rf, ok := ret.Get(0).(func() gorm.SQLCommon); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gorm.SQLCommon)
		}
	}

	return r0
}

// Count provides a mock function with given fields: value
func (_m *MockGormw) Count(value interface{}) Gormw {
	ret := _m.Called(value)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}) Gormw); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Create provides a mock function with given fields: value
func (_m *MockGormw) Create(value interface{}) Gormw {
	ret := _m.Called(value)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}) Gormw); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// CreateTable provides a mock function with given fields: values
func (_m *MockGormw) CreateTable(values ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(...interface{}) Gormw); ok {
		r0 = rf(values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// DB provides a mock function with given fields:
func (_m *MockGormw) DB() *sql.DB {
	ret := _m.Called()

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	return r0
}

// Debug provides a mock function with given fields:
func (_m *MockGormw) Debug() Gormw {
	ret := _m.Called()

	var r0 Gormw
	if rf, ok := ret.Get(0).(func() Gormw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: value, where
func (_m *MockGormw) Delete(value interface{}, where ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, where...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) Gormw); ok {
		r0 = rf(value, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// DropColumn provides a mock function with given fields: column
func (_m *MockGormw) DropColumn(column string) Gormw {
	ret := _m.Called(column)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string) Gormw); ok {
		r0 = rf(column)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// DropTable provides a mock function with given fields: values
func (_m *MockGormw) DropTable(values ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(...interface{}) Gormw); ok {
		r0 = rf(values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// DropTableIfExists provides a mock function with given fields: values
func (_m *MockGormw) DropTableIfExists(values ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(...interface{}) Gormw); ok {
		r0 = rf(values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Error provides a mock function with given fields:
func (_m *MockGormw) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exec provides a mock function with given fields: _a0, values
func (_m *MockGormw) Exec(_a0 string, values ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, ...interface{}) Gormw); ok {
		r0 = rf(_a0, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Find provides a mock function with given fields: out, where
func (_m *MockGormw) Find(out interface{}, where ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, out)
	_ca = append(_ca, where...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) Gormw); ok {
		r0 = rf(out, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// First provides a mock function with given fields: out, where
func (_m *MockGormw) First(out interface{}, where ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, out)
	_ca = append(_ca, where...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) Gormw); ok {
		r0 = rf(out, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// FirstOrCreate provides a mock function with given fields: out, where
func (_m *MockGormw) FirstOrCreate(out interface{}, where ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, out)
	_ca = append(_ca, where...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) Gormw); ok {
		r0 = rf(out, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// FirstOrInit provides a mock function with given fields: out, where
func (_m *MockGormw) FirstOrInit(out interface{}, where ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, out)
	_ca = append(_ca, where...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) Gormw); ok {
		r0 = rf(out, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Get provides a mock function with given fields: name
func (_m *MockGormw) Get(name string) (interface{}, bool) {
	ret := _m.Called(name)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetErrors provides a mock function with given fields:
func (_m *MockGormw) GetErrors() []error {
	ret := _m.Called()

	var r0 []error
	if rf, ok := ret.Get(0).(func() []error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}

// Group provides a mock function with given fields: query
func (_m *MockGormw) Group(query string) Gormw {
	ret := _m.Called(query)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string) Gormw); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// HasTable provides a mock function with given fields: value
func (_m *MockGormw) HasTable(value interface{}) bool {
	ret := _m.Called(value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Having provides a mock function with given fields: query, values
func (_m *MockGormw) Having(query string, values ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, ...interface{}) Gormw); ok {
		r0 = rf(query, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// InstantSet provides a mock function with given fields: name, value
func (_m *MockGormw) InstantSet(name string, value interface{}) Gormw {
	ret := _m.Called(name, value)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, interface{}) Gormw); ok {
		r0 = rf(name, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Joins provides a mock function with given fields: query, args
func (_m *MockGormw) Joins(query string, args ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, ...interface{}) Gormw); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Last provides a mock function with given fields: out, where
func (_m *MockGormw) Last(out interface{}, where ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, out)
	_ca = append(_ca, where...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) Gormw); ok {
		r0 = rf(out, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Limit provides a mock function with given fields: value
func (_m *MockGormw) Limit(value int) Gormw {
	ret := _m.Called(value)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(int) Gormw); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// LogMode provides a mock function with given fields: enable
func (_m *MockGormw) LogMode(enable bool) Gormw {
	ret := _m.Called(enable)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(bool) Gormw); ok {
		r0 = rf(enable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Model provides a mock function with given fields: value
func (_m *MockGormw) Model(value interface{}) Gormw {
	ret := _m.Called(value)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}) Gormw); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// ModifyColumn provides a mock function with given fields: column, typ
func (_m *MockGormw) ModifyColumn(column string, typ string) Gormw {
	ret := _m.Called(column, typ)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, string) Gormw); ok {
		r0 = rf(column, typ)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// New provides a mock function with given fields:
func (_m *MockGormw) New() Gormw {
	ret := _m.Called()

	var r0 Gormw
	if rf, ok := ret.Get(0).(func() Gormw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// NewRecord provides a mock function with given fields: value
func (_m *MockGormw) NewRecord(value interface{}) bool {
	ret := _m.Called(value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewScope provides a mock function with given fields: value
func (_m *MockGormw) NewScope(value interface{}) *gorm.Scope {
	ret := _m.Called(value)

	var r0 *gorm.Scope
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.Scope); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.Scope)
		}
	}

	return r0
}

// Not provides a mock function with given fields: query, args
func (_m *MockGormw) Not(query interface{}, args ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) Gormw); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Offset provides a mock function with given fields: value
func (_m *MockGormw) Offset(value int) Gormw {
	ret := _m.Called(value)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(int) Gormw); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Omit provides a mock function with given fields: columns
func (_m *MockGormw) Omit(columns ...string) Gormw {
	_va := make([]interface{}, len(columns))
	for _i := range columns {
		_va[_i] = columns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(...string) Gormw); ok {
		r0 = rf(columns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Or provides a mock function with given fields: query, args
func (_m *MockGormw) Or(query interface{}, args ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) Gormw); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Order provides a mock function with given fields: value, reorder
func (_m *MockGormw) Order(value string, reorder ...bool) Gormw {
	_va := make([]interface{}, len(reorder))
	for _i := range reorder {
		_va[_i] = reorder[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, ...bool) Gormw); ok {
		r0 = rf(value, reorder...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Pluck provides a mock function with given fields: column, value
func (_m *MockGormw) Pluck(column string, value interface{}) Gormw {
	ret := _m.Called(column, value)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, interface{}) Gormw); ok {
		r0 = rf(column, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Preload provides a mock function with given fields: column, conditions
func (_m *MockGormw) Preload(column string, conditions ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, column)
	_ca = append(_ca, conditions...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, ...interface{}) Gormw); ok {
		r0 = rf(column, conditions...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Raw provides a mock function with given fields: _a0, values
func (_m *MockGormw) Raw(_a0 string, values ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, ...interface{}) Gormw); ok {
		r0 = rf(_a0, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// RecordNotFound provides a mock function with given fields:
func (_m *MockGormw) RecordNotFound() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Related provides a mock function with given fields: value, foreignKeys
func (_m *MockGormw) Related(value interface{}, foreignKeys ...string) Gormw {
	_va := make([]interface{}, len(foreignKeys))
	for _i := range foreignKeys {
		_va[_i] = foreignKeys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}, ...string) Gormw); ok {
		r0 = rf(value, foreignKeys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// RemoveIndex provides a mock function with given fields: indexName
func (_m *MockGormw) RemoveIndex(indexName string) Gormw {
	ret := _m.Called(indexName)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string) Gormw); ok {
		r0 = rf(indexName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Rollback provides a mock function with given fields:
func (_m *MockGormw) Rollback() Gormw {
	ret := _m.Called()

	var r0 Gormw
	if rf, ok := ret.Get(0).(func() Gormw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Row provides a mock function with given fields:
func (_m *MockGormw) Row() *sql.Row {
	ret := _m.Called()

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func() *sql.Row); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// Rows provides a mock function with given fields:
func (_m *MockGormw) Rows() (*sql.Rows, error) {
	ret := _m.Called()

	var r0 *sql.Rows
	if rf, ok := ret.Get(0).(func() *sql.Rows); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RowsAffected provides a mock function with given fields:
func (_m *MockGormw) RowsAffected() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Save provides a mock function with given fields: value
func (_m *MockGormw) Save(value interface{}) Gormw {
	ret := _m.Called(value)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}) Gormw); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Scan provides a mock function with given fields: dest
func (_m *MockGormw) Scan(dest interface{}) Gormw {
	ret := _m.Called(dest)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}) Gormw); ok {
		r0 = rf(dest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// ScanRows provides a mock function with given fields: rows, result
func (_m *MockGormw) ScanRows(rows *sql.Rows, result interface{}) error {
	ret := _m.Called(rows, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sql.Rows, interface{}) error); ok {
		r0 = rf(rows, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scopes provides a mock function with given fields: funcs
func (_m *MockGormw) Scopes(funcs ...func(*gorm.DB) *gorm.DB) Gormw {
	_va := make([]interface{}, len(funcs))
	for _i := range funcs {
		_va[_i] = funcs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(...func(*gorm.DB) *gorm.DB) Gormw); ok {
		r0 = rf(funcs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Select provides a mock function with given fields: query, args
func (_m *MockGormw) Select(query interface{}, args ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) Gormw); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Set provides a mock function with given fields: name, value
func (_m *MockGormw) Set(name string, value interface{}) Gormw {
	ret := _m.Called(name, value)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string, interface{}) Gormw); ok {
		r0 = rf(name, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// SetJoinTableHandler provides a mock function with given fields: source, column, handler
func (_m *MockGormw) SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface) {
	_m.Called(source, column, handler)
}

// SetLogger provides a mock function with given fields: l
func (_m *MockGormw) SetLogger(l gorm.Logger) {
	_m.Called(l)
}

// SingularTable provides a mock function with given fields: enable
func (_m *MockGormw) SingularTable(enable bool) {
	_m.Called(enable)
}

// Table provides a mock function with given fields: name
func (_m *MockGormw) Table(name string) Gormw {
	ret := _m.Called(name)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(string) Gormw); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Unscoped provides a mock function with given fields:
func (_m *MockGormw) Unscoped() Gormw {
	ret := _m.Called()

	var r0 Gormw
	if rf, ok := ret.Get(0).(func() Gormw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Update provides a mock function with given fields: attrs
func (_m *MockGormw) Update(attrs ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, attrs...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(...interface{}) Gormw); ok {
		r0 = rf(attrs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// UpdateColumn provides a mock function with given fields: attrs
func (_m *MockGormw) UpdateColumn(attrs ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, attrs...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(...interface{}) Gormw); ok {
		r0 = rf(attrs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// UpdateColumns provides a mock function with given fields: values
func (_m *MockGormw) UpdateColumns(values interface{}) Gormw {
	ret := _m.Called(values)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}) Gormw); ok {
		r0 = rf(values)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Updates provides a mock function with given fields: values, ignoreProtectedAttrs
func (_m *MockGormw) Updates(values interface{}, ignoreProtectedAttrs ...bool) Gormw {
	_va := make([]interface{}, len(ignoreProtectedAttrs))
	for _i := range ignoreProtectedAttrs {
		_va[_i] = ignoreProtectedAttrs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, values)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}, ...bool) Gormw); ok {
		r0 = rf(values, ignoreProtectedAttrs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}

// Where provides a mock function with given fields: query, args
func (_m *MockGormw) Where(query interface{}, args ...interface{}) Gormw {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 Gormw
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) Gormw); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Gormw)
		}
	}

	return r0
}
