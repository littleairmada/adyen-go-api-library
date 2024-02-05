/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// TransactionsApi service
type TransactionsApi common.Service

// All parameters accepted by TransactionsApi.GetAllTransactions
type TransactionsApiGetAllTransactionsInput struct {
	createdSince        *time.Time
	createdUntil        *time.Time
	balancePlatform     *string
	paymentInstrumentId *string
	accountHolderId     *string
	balanceAccountId    *string
	cursor              *string
	limit               *int32
}

// Only include transactions that have been created on or after this point in time. The value must be in ISO 8601 format. For example, **2021-05-30T15:07:40Z**.
func (r TransactionsApiGetAllTransactionsInput) CreatedSince(createdSince time.Time) TransactionsApiGetAllTransactionsInput {
	r.createdSince = &createdSince
	return r
}

// Only include transactions that have been created on or before this point in time. The value must be in ISO 8601 format. For example, **2021-05-30T15:07:40Z**.
func (r TransactionsApiGetAllTransactionsInput) CreatedUntil(createdUntil time.Time) TransactionsApiGetAllTransactionsInput {
	r.createdUntil = &createdUntil
	return r
}

// The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id).  Required if you don&#39;t provide a &#x60;balanceAccountId&#x60; or &#x60;accountHolderId&#x60;.
func (r TransactionsApiGetAllTransactionsInput) BalancePlatform(balancePlatform string) TransactionsApiGetAllTransactionsInput {
	r.balancePlatform = &balancePlatform
	return r
}

// The unique identifier of the [payment instrument](https://docs.adyen.com/api-explorer/balanceplatform/latest/get/paymentInstruments/_id_).  To use this parameter, you must also provide a &#x60;balanceAccountId&#x60;, &#x60;accountHolderId&#x60;, or &#x60;balancePlatform&#x60;.  The &#x60;paymentInstrumentId&#x60; must be related to the &#x60;balanceAccountId&#x60; or &#x60;accountHolderId&#x60; that you provide.
func (r TransactionsApiGetAllTransactionsInput) PaymentInstrumentId(paymentInstrumentId string) TransactionsApiGetAllTransactionsInput {
	r.paymentInstrumentId = &paymentInstrumentId
	return r
}

// The unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/accountHolders/{id}__queryParam_id).  Required if you don&#39;t provide a &#x60;balanceAccountId&#x60; or &#x60;balancePlatform&#x60;.  If you provide a &#x60;balanceAccountId&#x60;, the &#x60;accountHolderId&#x60; must be related to the &#x60;balanceAccountId&#x60;.
func (r TransactionsApiGetAllTransactionsInput) AccountHolderId(accountHolderId string) TransactionsApiGetAllTransactionsInput {
	r.accountHolderId = &accountHolderId
	return r
}

// The unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balanceAccounts/{id}__queryParam_id).  Required if you don&#39;t provide an &#x60;accountHolderId&#x60; or &#x60;balancePlatform&#x60;.  If you provide an &#x60;accountHolderId&#x60;, the &#x60;balanceAccountId&#x60; must be related to the &#x60;accountHolderId&#x60;.
func (r TransactionsApiGetAllTransactionsInput) BalanceAccountId(balanceAccountId string) TransactionsApiGetAllTransactionsInput {
	r.balanceAccountId = &balanceAccountId
	return r
}

// The &#x60;cursor&#x60; returned in the links of the previous response.
func (r TransactionsApiGetAllTransactionsInput) Cursor(cursor string) TransactionsApiGetAllTransactionsInput {
	r.cursor = &cursor
	return r
}

// The number of items returned per page, maximum of 100 items. By default, the response returns 10 items per page.
func (r TransactionsApiGetAllTransactionsInput) Limit(limit int32) TransactionsApiGetAllTransactionsInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for GetAllTransactions

@return TransactionsApiGetAllTransactionsInput
*/
func (a *TransactionsApi) GetAllTransactionsInput() TransactionsApiGetAllTransactionsInput {
	return TransactionsApiGetAllTransactionsInput{}
}

/*
GetAllTransactions Get all transactions

>Versions 1 and 2 of the Transfers API are deprecated. If you are just starting your implementation, use the latest version.

Returns all the transactions related to a balance account, account holder, or balance platform.

When making this request, you must include at least one of the following:
- `balanceAccountId`
- `accountHolderId`
- `balancePlatform`.

This endpoint supports cursor-based pagination. The response returns the first page of results, and returns links to the next and previous pages when applicable. You can use the links to page through the results.



@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransactionsApiGetAllTransactionsInput - Request parameters, see GetAllTransactionsInput
@return TransactionSearchResponse, *http.Response, error
*/
func (a *TransactionsApi) GetAllTransactions(ctx context.Context, r TransactionsApiGetAllTransactionsInput) (TransactionSearchResponse, *http.Response, error) {
	res := &TransactionSearchResponse{}
	path := "/transactions"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.balancePlatform != nil {
		common.ParameterAddToQuery(queryParams, "balancePlatform", r.balancePlatform, "")
	}
	if r.paymentInstrumentId != nil {
		common.ParameterAddToQuery(queryParams, "paymentInstrumentId", r.paymentInstrumentId, "")
	}
	if r.accountHolderId != nil {
		common.ParameterAddToQuery(queryParams, "accountHolderId", r.accountHolderId, "")
	}
	if r.balanceAccountId != nil {
		common.ParameterAddToQuery(queryParams, "balanceAccountId", r.balanceAccountId, "")
	}
	if r.cursor != nil {
		common.ParameterAddToQuery(queryParams, "cursor", r.cursor, "")
	}
	if r.createdSince != nil {
		common.ParameterAddToQuery(queryParams, "createdSince", r.createdSince, "")
	}
	if r.createdUntil != nil {
		common.ParameterAddToQuery(queryParams, "createdUntil", r.createdUntil, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryParams, "limit", r.limit, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TransactionsApi.GetTransaction
type TransactionsApiGetTransactionInput struct {
	id string
}

/*
Prepare a request for GetTransaction
@param id The unique identifier of the transaction.
@return TransactionsApiGetTransactionInput
*/
func (a *TransactionsApi) GetTransactionInput(id string) TransactionsApiGetTransactionInput {
	return TransactionsApiGetTransactionInput{
		id: id,
	}
}

/*
GetTransaction Get a transaction

>Versions 1 and 2 of the Transfers API are deprecated. If you are just starting your implementation, use the latest version.

Returns a transaction.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransactionsApiGetTransactionInput - Request parameters, see GetTransactionInput
@return Transaction, *http.Response, error
*/
func (a *TransactionsApi) GetTransaction(ctx context.Context, r TransactionsApiGetTransactionInput) (Transaction, *http.Response, error) {
	res := &Transaction{}
	path := "/transactions/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}
