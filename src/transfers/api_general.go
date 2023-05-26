/*
Transfers API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"context"
	_context "context"
	_nethttp "net/http"
	"net/url"
	"strings"
	"time"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// GeneralApi GeneralApi service
type GeneralApi common.Service

type GeneralApiGetAllTransactionsConfig struct {
	ctx                 context.Context
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
func (r GeneralApiGetAllTransactionsConfig) CreatedSince(createdSince time.Time) GeneralApiGetAllTransactionsConfig {
	r.createdSince = &createdSince
	return r
}

// Only include transactions that have been created on or before this point in time. The value must be in ISO 8601 format. For example, **2021-05-30T15:07:40Z**.
func (r GeneralApiGetAllTransactionsConfig) CreatedUntil(createdUntil time.Time) GeneralApiGetAllTransactionsConfig {
	r.createdUntil = &createdUntil
	return r
}

// Unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id).
func (r GeneralApiGetAllTransactionsConfig) BalancePlatform(balancePlatform string) GeneralApiGetAllTransactionsConfig {
	r.balancePlatform = &balancePlatform
	return r
}

// Unique identifier of the [payment instrument](https://docs.adyen.com/api-explorer/balanceplatform/latest/get/paymentInstruments/_id_).
func (r GeneralApiGetAllTransactionsConfig) PaymentInstrumentId(paymentInstrumentId string) GeneralApiGetAllTransactionsConfig {
	r.paymentInstrumentId = &paymentInstrumentId
	return r
}

// Unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/accountHolders/{id}__queryParam_id).
func (r GeneralApiGetAllTransactionsConfig) AccountHolderId(accountHolderId string) GeneralApiGetAllTransactionsConfig {
	r.accountHolderId = &accountHolderId
	return r
}

// Unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balanceAccounts/{id}__queryParam_id).
func (r GeneralApiGetAllTransactionsConfig) BalanceAccountId(balanceAccountId string) GeneralApiGetAllTransactionsConfig {
	r.balanceAccountId = &balanceAccountId
	return r
}

// The &#x60;cursor&#x60; returned in the links of the previous response.
func (r GeneralApiGetAllTransactionsConfig) Cursor(cursor string) GeneralApiGetAllTransactionsConfig {
	r.cursor = &cursor
	return r
}

// The number of items returned per page, maximum of 100 items. By default, the response returns 10 items per page.
func (r GeneralApiGetAllTransactionsConfig) Limit(limit int32) GeneralApiGetAllTransactionsConfig {
	r.limit = &limit
	return r
}

/*
GetAllTransactions Get all transactions

Returns all transactions related to a balance account with a payment instrument of type **bankAccount**.

This endpoint supports cursor-based pagination. The response returns the first page of results, and returns links to the next page when applicable. You can use the links to page through the results. The response also returns links to the previous page when applicable.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GeneralApiGetAllTransactionsConfig
*/
func (a *GeneralApi) GetAllTransactionsConfig(ctx context.Context) GeneralApiGetAllTransactionsConfig {
	return GeneralApiGetAllTransactionsConfig{
		ctx: ctx,
	}
}

/*
Get all transactions
Returns all transactions related to a balance account with a payment instrument of type **bankAccount**.  This endpoint supports cursor-based pagination. The response returns the first page of results, and returns links to the next page when applicable. You can use the links to page through the results. The response also returns links to the previous page when applicable.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TransactionSearchResponse
*/

func (a *GeneralApi) GetAllTransactions(r GeneralApiGetAllTransactionsConfig) (TransactionSearchResponse, *_nethttp.Response, error) {
	res := &TransactionSearchResponse{}
	path := "/transactions"
	queryString := url.Values{}
	if r.balancePlatform != nil {
		common.ParameterAddToQuery(queryString, "balancePlatform", r.balancePlatform, "")
	}
	if r.paymentInstrumentId != nil {
		common.ParameterAddToQuery(queryString, "paymentInstrumentId", r.paymentInstrumentId, "")
	}
	if r.accountHolderId != nil {
		common.ParameterAddToQuery(queryString, "accountHolderId", r.accountHolderId, "")
	}
	if r.balanceAccountId != nil {
		common.ParameterAddToQuery(queryString, "balanceAccountId", r.balanceAccountId, "")
	}
	if r.cursor != nil {
		common.ParameterAddToQuery(queryString, "cursor", r.cursor, "")
	}
	if r.createdSince != nil {
		common.ParameterAddToQuery(queryString, "createdSince", r.createdSince, "")
	}
	if r.createdUntil != nil {
		common.ParameterAddToQuery(queryString, "createdUntil", r.createdUntil, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryString, "limit", r.limit, "")
	}
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path+"?"+queryString.Encode(), []_context.Context{r.ctx})
	return *res, httpRes, err
}

type GeneralApiGetTransactionConfig struct {
	ctx context.Context
	id  string
}

/*
GetTransaction Get a transaction

Returns a transaction.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Unique identifier of the transaction.
	@return GeneralApiGetTransactionConfig
*/
func (a *GeneralApi) GetTransactionConfig(ctx context.Context, id string) GeneralApiGetTransactionConfig {
	return GeneralApiGetTransactionConfig{
		ctx: ctx,
		id:  id,
	}
}

/*
Get a transaction
Returns a transaction.
 * @param id Unique identifier of the transaction.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Transaction
*/

func (a *GeneralApi) GetTransaction(r GeneralApiGetTransactionConfig) (Transaction, *_nethttp.Response, error) {
	res := &Transaction{}
	path := "/transactions/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type GeneralApiTransferFundsConfig struct {
	ctx          context.Context
	transferInfo *TransferInfo
}

func (r GeneralApiTransferFundsConfig) TransferInfo(transferInfo TransferInfo) GeneralApiTransferFundsConfig {
	r.transferInfo = &transferInfo
	return r
}

/*
TransferFunds Transfer funds

Starts a request to transfer funds to [balance accounts](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts), [transfer instruments](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments), or third-party bank accounts. Adyen sends the outcome of the transfer request through webhooks.

To use this endpoint, you need an additional role for your API credential and transfers must be enabled for the source balance account. Your Adyen contact will set these up for you.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GeneralApiTransferFundsConfig
*/
func (a *GeneralApi) TransferFundsConfig(ctx context.Context) GeneralApiTransferFundsConfig {
	return GeneralApiTransferFundsConfig{
		ctx: ctx,
	}
}

/*
Transfer funds
Starts a request to transfer funds to [balance accounts](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts), [transfer instruments](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments), or third-party bank accounts. Adyen sends the outcome of the transfer request through webhooks.  To use this endpoint, you need an additional role for your API credential and transfers must be enabled for the source balance account. Your Adyen contact will set these up for you.
 * @param req TransferInfo - reference of TransferInfo).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Transfer
*/

func (a *GeneralApi) TransferFunds(r GeneralApiTransferFundsConfig) (Transfer, *_nethttp.Response, error) {
	res := &Transfer{}
	path := "/transfers"
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, r.transferInfo, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}
