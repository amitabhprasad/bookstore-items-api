package services

import (
	"net/http"

	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/domains/items"
	"github.com/amitabhprasad/bookstore-util-go/rest_errors"
)

var (
	ItemService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {

	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me !", http.StatusNotImplemented, "not_implemented", nil)
}
