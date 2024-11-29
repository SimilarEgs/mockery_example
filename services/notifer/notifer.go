package notifer

import (
	"errors"
	"github.com/SimilarEgs/mockery_example/services/login_converter"
	"log"
)

// для всех интерфейсов на уровне пакета: --all
// для конкретного интерфейса: --name=<НазваниеИнтерфейса>

//go:generate go run github.com/vektra/mockery/v2@v2.35.3 --name=INotifier
type INotifier interface {
	NotifyOfUserCreated(id string) error
}

type NotifierService struct {
	URL            string
	LoginConverter login_converter.ILoginConverter
}

func NewNotifierService(url string, LoginConverter login_converter.ILoginConverter) INotifier {
	return &NotifierService{
		URL:            url,
		LoginConverter: LoginConverter,
	}
}

func (s *NotifierService) NotifyOfUserCreated(id string) error {
	if !s.LoginConverter.IsFullID(id) {
		return errors.New("notifier.id_not_full")
	}

	log.Printf("notifier: sending notify of user <%s> creation to - %s\n", id, s.URL)

	return nil
}
