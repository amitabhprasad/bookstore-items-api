package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/domains/items"
	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/domains/queries"
	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/services"
	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/utils/http_utils"
	"github.com/amitabhprasad/bookstore-oauth2-go/oauth2"
	"github.com/amitabhprasad/bookstore-util-go/rest_errors"
	"github.com/gorilla/mux"
)

var (
	ItemController itemControllerInterface = &itemController{}
)

type itemController struct {
}

type itemControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}

func (i *itemController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth2.AuthenticateRequest(r); err != nil {
		http_utils.ResponseError(w, err)
		return
	}
	sellerId := oauth2.GetCallerId(r)
	if sellerId == 0 {
		responseErr := rest_errors.NewUnAuthorizedError("unable to reterieve user information from given token")
		http_utils.ResponseError(w, responseErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.ResponseError(w, responseErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		responseErr := rest_errors.NewBadRequestError("unable to marshal data to item")
		http_utils.ResponseError(w, responseErr)
		return
	}
	itemRequest.Seller = fmt.Sprintf("%d", sellerId)
	result, createErr := services.ItemService.Create(itemRequest)
	if createErr != nil {
		http_utils.ResponseError(w, createErr)
		return
	}
	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (i *itemController) Get(w http.ResponseWriter, r *http.Request) {
	//YQtuwH8BrUNaQaN5zuyc
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])
	item, err := services.ItemService.Get(itemId)
	if err != nil {
		http_utils.ResponseError(w, err)
		return
	}
	fmt.Println("before item ", item.AvailableQuantity)
	item.Id = itemId
	fmt.Println("item ", item.Status)
	http_utils.ResponseJson(w, http.StatusOK, item)
}

func (i *itemController) Search(w http.ResponseWriter, r *http.Request) {
	var query queries.EsQuery
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, apiErr)
		return
	}
	defer r.Body.Close()
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, apiErr)
		return
	}
	items, searchError := services.ItemService.Search(query)
	if searchError != nil {
		http_utils.ResponseError(w, searchError)
		return
	}
	http_utils.ResponseJson(w, http.StatusOK, items)
}
