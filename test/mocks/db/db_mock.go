package db_mock

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDb struct {
	mock.Mock
}

func (m *MockDb) Raw(query string, args ...interface{}) *gorm.DB {
	args = append([]interface{}{query}, args...)
	call := m.Called(args...)
	return call.Get(0).(*gorm.DB)
}

func (m *MockDb) Scan(dest interface{}) *gorm.DB {
	arg := m.Called(dest) 
	return arg.Get(0).(*gorm.DB)
}
