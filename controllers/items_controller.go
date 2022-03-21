package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/domains/items"
	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/services"
	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/utils/http_utils"
	"github.com/amitabhprasad/bookstore-oauth2-go/oauth2"
	"github.com/amitabhprasad/bookstore-util-go/rest_errors"
)

var (
	ItemController itemControllerInterface = &itemController{}
)

type itemController struct {
}

type itemControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

func (i *itemController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth2.AuthenticateRequest(r); err != nil {
		http_utils.ResponseError(w, *err)
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

}
