package sentences_mock

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

type MockSentences struct {
	mock.Mock
}

func (m *MockSentences) DeleteRecord(c *fiber.Ctx, tableName, idColumn, idValue string, entity interface{}) error {
	args := m.Called(c, tableName, idColumn, idValue, entity)
	return args.Error(0)
}

func (m *MockSentences) PaginateAndQuery(c *fiber.Ctx, tableName string) (string, error) {
	args := m.Called(c, tableName)
	return args.String(0), args.Error(1)
}
