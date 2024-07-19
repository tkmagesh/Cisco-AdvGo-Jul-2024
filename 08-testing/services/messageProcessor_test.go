package services

import (
	"testing"

	mocks "github.com/tkmagesh/cisco-advgo-jul-2024/08-testing/mocks/services"
)

/*
type MockMessageService struct {
	invocations map[string]int
}

// implementation of the "MessageService" interface (expected by the MessageProcessor)
func (mms *MockMessageService) Send(msg string) bool {
	mms.invocations["Send"]++
	return true
}

func (mms *MockMessageService) Called(methodName string) bool {
	if count, exists := mms.invocations[methodName]; exists {
		return count > 0
	}
	return false
}

func NewMockMessageService() *MockMessageService {
	return &MockMessageService{
		invocations: make(map[string]int),
	}
}

func Test_MessageProcessor(t *testing.T) {
	// Arrange
	mms := NewMockMessageService()
	sut := NewMessageProcessor(mms)

	// Act
	sut.Process("dummy message")

	// Assert
	if !mms.Called("Send") {
		t.Error("MessageService.Send() was not invoked")
	}
}
*/

func Test_MessageProcessor(t *testing.T) {
	// Arrange
	mms := mocks.NewMessageService(t)
	sut := NewMessageProcessor(mms)
	testMessage := "dummy message"
	mms.On("Send", testMessage).Return(true)

	// Act
	sut.Process(testMessage)

	// Assert
	mms.AssertExpectations(t)
}
