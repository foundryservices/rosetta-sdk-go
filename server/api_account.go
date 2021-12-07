// Copyright 2021 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/foundryservices/rosetta-sdk-go/asserter"
	"github.com/foundryservices/rosetta-sdk-go/types"
)

// A AccountAPIController binds http requests to an api service and writes the service results to
// the http response
type AccountAPIController struct {
	service  AccountAPIServicer
	asserter *asserter.Asserter
}

// NewAccountAPIController creates a default api controller
func NewAccountAPIController(
	s AccountAPIServicer,
	asserter *asserter.Asserter,
) Router {
	return &AccountAPIController{
		service:  s,
		asserter: asserter,
	}
}

// Routes returns all of the api route for the AccountAPIController
func (c *AccountAPIController) Routes() Routes {
	return Routes{
		{
			"AccountBalance",
			strings.ToUpper("Post"),
			"/account/balance",
			c.AccountBalance,
		},
		{
			"AccountCoins",
			strings.ToUpper("Post"),
			"/account/coins",
			c.AccountCoins,
		},
	}
}

// AccountBalance - Get an Account's Balance
func (c *AccountAPIController) AccountBalance(w http.ResponseWriter, r *http.Request) {
	accountBalanceRequest := &types.AccountBalanceRequest{}
	if err := json.NewDecoder(r.Body).Decode(&accountBalanceRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that AccountBalanceRequest is correct
	if err := c.asserter.AccountBalanceRequest(accountBalanceRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.AccountBalance(r.Context(), accountBalanceRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}

// AccountCoins - Get an Account's Unspent Coins
func (c *AccountAPIController) AccountCoins(w http.ResponseWriter, r *http.Request) {
	accountCoinsRequest := &types.AccountCoinsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&accountCoinsRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that AccountCoinsRequest is correct
	if err := c.asserter.AccountCoinsRequest(accountCoinsRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.AccountCoins(r.Context(), accountCoinsRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}
