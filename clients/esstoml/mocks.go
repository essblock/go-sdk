package esstoml

import "github.com/stretchr/testify/mock"

// MockClient is a mockable esstoml client.
type MockClient struct {
	mock.Mock
}

// GetEssToml is a mocking a method
func (m *MockClient) GetEssToml(domain string) (*Response, error) {
	a := m.Called(domain)
	return a.Get(0).(*Response), a.Error(1)
}

// GetEssTomlByAddress is a mocking a method
func (m *MockClient) GetEssTomlByAddress(address string) (*Response, error) {
	a := m.Called(address)
	return a.Get(0).(*Response), a.Error(1)
}
