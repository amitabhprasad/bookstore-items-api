package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/client/elasticsearch"
	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/domains/queries"
	"github.com/amitabhprasad/bookstore-util-go/rest_errors"
)

const (
	indexItems = "items"
	itemType   = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, itemType, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when saving item", err)
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {

	result, err := elasticsearch.Client.Get(indexItems, itemType, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no items found with id %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get item by id %s", i.Id),
			errors.New("data base error"))
	}
	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", err)
	}
	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", err)
	}

	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when when searching documents ", err)
	}
	fmt.Println(result)
	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		var item Item
		bytes, _ := hit.Source.MarshalJSON()
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when parsing response", err)
		}
		item.Id = hit.Id
		items[index] = item
	}
	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("No matching records found based on given criteria")
	}
	return items, nil
}
