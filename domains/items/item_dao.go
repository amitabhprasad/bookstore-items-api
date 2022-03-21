package items

import (
	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/client/elasticsearch"
	"github.com/amitabhprasad/bookstore-util-go/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when saving item", err)
	}
	i.Id = result.Id
	return nil
}
