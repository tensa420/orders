package order

import (
	"order/internal/api/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
)

type ServerSuite struct {
	suite.Suite
	service *mocks.OrderService
	ctx     context.Context
	server  *Server
}

func (s *ServerSuite) SetupSuite() {
	s.ctx = context.Background()
	s.service = mocks.NewOrderService(s.T())
	s.server = NewOrderServer(s.service)
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}
