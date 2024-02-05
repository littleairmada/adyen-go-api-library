/*
Adyen Stored Value API

API version: 46
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storedvalue

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// GeneralApi service
type GeneralApi common.Service

// All parameters accepted by GeneralApi.ChangeStatus
type GeneralApiChangeStatusInput struct {
	storedValueStatusChangeRequest *StoredValueStatusChangeRequest
}

func (r GeneralApiChangeStatusInput) StoredValueStatusChangeRequest(storedValueStatusChangeRequest StoredValueStatusChangeRequest) GeneralApiChangeStatusInput {
	r.storedValueStatusChangeRequest = &storedValueStatusChangeRequest
	return r
}

/*
Prepare a request for ChangeStatus

@return GeneralApiChangeStatusInput
*/
func (a *GeneralApi) ChangeStatusInput() GeneralApiChangeStatusInput {
	return GeneralApiChangeStatusInput{}
}

/*
ChangeStatus Changes the status of the payment method.

Changes the status of the provided payment method to the specified status.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiChangeStatusInput - Request parameters, see ChangeStatusInput
@return StoredValueStatusChangeResponse, *http.Response, error
*/
func (a *GeneralApi) ChangeStatus(ctx context.Context, r GeneralApiChangeStatusInput) (StoredValueStatusChangeResponse, *http.Response, error) {
	res := &StoredValueStatusChangeResponse{}
	path := "/changeStatus"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.storedValueStatusChangeRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.CheckBalance
type GeneralApiCheckBalanceInput struct {
	storedValueBalanceCheckRequest *StoredValueBalanceCheckRequest
}

func (r GeneralApiCheckBalanceInput) StoredValueBalanceCheckRequest(storedValueBalanceCheckRequest StoredValueBalanceCheckRequest) GeneralApiCheckBalanceInput {
	r.storedValueBalanceCheckRequest = &storedValueBalanceCheckRequest
	return r
}

/*
Prepare a request for CheckBalance

@return GeneralApiCheckBalanceInput
*/
func (a *GeneralApi) CheckBalanceInput() GeneralApiCheckBalanceInput {
	return GeneralApiCheckBalanceInput{}
}

/*
CheckBalance Checks the balance.

Checks the balance of the provided payment method.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiCheckBalanceInput - Request parameters, see CheckBalanceInput
@return StoredValueBalanceCheckResponse, *http.Response, error
*/
func (a *GeneralApi) CheckBalance(ctx context.Context, r GeneralApiCheckBalanceInput) (StoredValueBalanceCheckResponse, *http.Response, error) {
	res := &StoredValueBalanceCheckResponse{}
	path := "/checkBalance"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.storedValueBalanceCheckRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.Issue
type GeneralApiIssueInput struct {
	storedValueIssueRequest *StoredValueIssueRequest
}

func (r GeneralApiIssueInput) StoredValueIssueRequest(storedValueIssueRequest StoredValueIssueRequest) GeneralApiIssueInput {
	r.storedValueIssueRequest = &storedValueIssueRequest
	return r
}

/*
Prepare a request for Issue

@return GeneralApiIssueInput
*/
func (a *GeneralApi) IssueInput() GeneralApiIssueInput {
	return GeneralApiIssueInput{}
}

/*
Issue Issues a new card.

Issues a new card of the given payment method.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiIssueInput - Request parameters, see IssueInput
@return StoredValueIssueResponse, *http.Response, error
*/
func (a *GeneralApi) Issue(ctx context.Context, r GeneralApiIssueInput) (StoredValueIssueResponse, *http.Response, error) {
	res := &StoredValueIssueResponse{}
	path := "/issue"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.storedValueIssueRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.Load
type GeneralApiLoadInput struct {
	storedValueLoadRequest *StoredValueLoadRequest
}

func (r GeneralApiLoadInput) StoredValueLoadRequest(storedValueLoadRequest StoredValueLoadRequest) GeneralApiLoadInput {
	r.storedValueLoadRequest = &storedValueLoadRequest
	return r
}

/*
Prepare a request for Load

@return GeneralApiLoadInput
*/
func (a *GeneralApi) LoadInput() GeneralApiLoadInput {
	return GeneralApiLoadInput{}
}

/*
Load Loads the payment method.

Loads the payment method with the specified funds.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiLoadInput - Request parameters, see LoadInput
@return StoredValueLoadResponse, *http.Response, error
*/
func (a *GeneralApi) Load(ctx context.Context, r GeneralApiLoadInput) (StoredValueLoadResponse, *http.Response, error) {
	res := &StoredValueLoadResponse{}
	path := "/load"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.storedValueLoadRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.MergeBalance
type GeneralApiMergeBalanceInput struct {
	storedValueBalanceMergeRequest *StoredValueBalanceMergeRequest
}

func (r GeneralApiMergeBalanceInput) StoredValueBalanceMergeRequest(storedValueBalanceMergeRequest StoredValueBalanceMergeRequest) GeneralApiMergeBalanceInput {
	r.storedValueBalanceMergeRequest = &storedValueBalanceMergeRequest
	return r
}

/*
Prepare a request for MergeBalance

@return GeneralApiMergeBalanceInput
*/
func (a *GeneralApi) MergeBalanceInput() GeneralApiMergeBalanceInput {
	return GeneralApiMergeBalanceInput{}
}

/*
MergeBalance Merge the balance of two cards.

Increases the balance of the paymentmethod by the full amount left on the source paymentmethod

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiMergeBalanceInput - Request parameters, see MergeBalanceInput
@return StoredValueBalanceMergeResponse, *http.Response, error
*/
func (a *GeneralApi) MergeBalance(ctx context.Context, r GeneralApiMergeBalanceInput) (StoredValueBalanceMergeResponse, *http.Response, error) {
	res := &StoredValueBalanceMergeResponse{}
	path := "/mergeBalance"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.storedValueBalanceMergeRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.VoidTransaction
type GeneralApiVoidTransactionInput struct {
	storedValueVoidRequest *StoredValueVoidRequest
}

func (r GeneralApiVoidTransactionInput) StoredValueVoidRequest(storedValueVoidRequest StoredValueVoidRequest) GeneralApiVoidTransactionInput {
	r.storedValueVoidRequest = &storedValueVoidRequest
	return r
}

/*
Prepare a request for VoidTransaction

@return GeneralApiVoidTransactionInput
*/
func (a *GeneralApi) VoidTransactionInput() GeneralApiVoidTransactionInput {
	return GeneralApiVoidTransactionInput{}
}

/*
VoidTransaction Voids a transaction.

Voids the referenced stored value transaction.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiVoidTransactionInput - Request parameters, see VoidTransactionInput
@return StoredValueVoidResponse, *http.Response, error
*/
func (a *GeneralApi) VoidTransaction(ctx context.Context, r GeneralApiVoidTransactionInput) (StoredValueVoidResponse, *http.Response, error) {
	res := &StoredValueVoidResponse{}
	path := "/voidTransaction"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.storedValueVoidRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
