package login_converter

import "strings"

// для всех интерфейсов на уровне пакета: --all
// для конкретного интерфейса: --name=<НазваниеИнтерфейса>

//go:generate go run github.com/vektra/mockery/v2@v2.35.3 --name=ILoginConverter
type ILoginConverter interface {
	ToFullID(id string) string
	IsFullID(id string) bool
}

type LoginConverterService struct {
	serverName string
}

func NewLoginConverterService(serverName string) ILoginConverter {
	return &LoginConverterService{
		serverName: serverName,
	}
}

func (s *LoginConverterService) ToFullID(id string) string {
	return id + "@" + s.serverName
}

func (s *LoginConverterService) IsFullID(id string) bool {
	return strings.HasPrefix(id, "@"+s.serverName)
}
