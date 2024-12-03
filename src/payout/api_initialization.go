/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// InitializationApi service
type InitializationApi common.Service

// All parameters accepted by InitializationApi.StoreDetail
type InitializationApiStoreDetailInput struct {
	storeDetailRequest *StoreDetailRequest
}

func (r InitializationApiStoreDetailInput) StoreDetailRequest(storeDetailRequest StoreDetailRequest) InitializationApiStoreDetailInput {
	r.storeDetailRequest = &storeDetailRequest
	return r
}

/*
Prepare a request for StoreDetail

@return InitializationApiStoreDetailInput
*/
func (a *InitializationApi) StoreDetailInput() InitializationApiStoreDetailInput {
	return InitializationApiStoreDetailInput{}
}

/*
StoreDetail Store payout details

> This endpoint is **deprecated** and no longer supports new integrations. Do one of the following:
>- If you are building a new integration, use the [Transfers API](https://docs.adyen.com/api-explorer/transfers/latest/overview) instead.
> - If you are already using the Payout API, reach out to your Adyen contact to learn how to migrate to the Transfers API.
>
> With the Transfers API, you can:
> - Handle multiple payout use cases with a single API.
> - Use new payout functionalities, such as instant payouts to bank accounts.
> - Receive webhooks with more details and defined transfer states.
>
> For more information about the payout features of the Transfers API, see our [Payouts](https://docs.adyen.com/payouts/payout-service) documentation.


Stores payment details under the `PAYOUT` recurring contract. These payment details can be used later to submit a payout via the `/submitThirdParty` call.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r InitializationApiStoreDetailInput - Request parameters, see StoreDetailInput
@return StoreDetailResponse, *http.Response, error
*/
func (a *InitializationApi) StoreDetail(ctx context.Context, r InitializationApiStoreDetailInput) (StoreDetailResponse, *http.Response, error) {
	res := &StoreDetailResponse{}
	path := "/storeDetail"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.storeDetailRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by InitializationApi.StoreDetailAndSubmitThirdParty
type InitializationApiStoreDetailAndSubmitThirdPartyInput struct {
	storeDetailAndSubmitRequest *StoreDetailAndSubmitRequest
}

func (r InitializationApiStoreDetailAndSubmitThirdPartyInput) StoreDetailAndSubmitRequest(storeDetailAndSubmitRequest StoreDetailAndSubmitRequest) InitializationApiStoreDetailAndSubmitThirdPartyInput {
	r.storeDetailAndSubmitRequest = &storeDetailAndSubmitRequest
	return r
}

/*
Prepare a request for StoreDetailAndSubmitThirdParty

@return InitializationApiStoreDetailAndSubmitThirdPartyInput
*/
func (a *InitializationApi) StoreDetailAndSubmitThirdPartyInput() InitializationApiStoreDetailAndSubmitThirdPartyInput {
	return InitializationApiStoreDetailAndSubmitThirdPartyInput{}
}

/*
StoreDetailAndSubmitThirdParty Store details and submit a payout

> This endpoint is **deprecated** and no longer supports new integrations. Do one of the following:
>- If you are building a new integration, use the POST [/transfers](https://docs.adyen.com/api-explorer/transfers/latest/post/transfers) endpoint instead.
> - If you are already using the Payout API, reach out to your Adyen contact to learn how to migrate to the Transfers API.
>
> With the Transfers API, you can:
> - Handle multiple payout use cases with a single API.
> - Use new payout functionalities, such as instant payouts to bank accounts.
> - Receive webhooks with more details and defined transfer states.
>
> For more information about the payout features of the Transfers API, see our [Payouts](https://docs.adyen.com/payouts/payout-service) documentation.


Submits a payout and stores its details for subsequent payouts.

The submitted payout must be confirmed or declined either by a reviewer or via `/confirmThirdParty` or `/declineThirdParty` calls.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r InitializationApiStoreDetailAndSubmitThirdPartyInput - Request parameters, see StoreDetailAndSubmitThirdPartyInput
@return StoreDetailAndSubmitResponse, *http.Response, error
*/
func (a *InitializationApi) StoreDetailAndSubmitThirdParty(ctx context.Context, r InitializationApiStoreDetailAndSubmitThirdPartyInput) (StoreDetailAndSubmitResponse, *http.Response, error) {
	res := &StoreDetailAndSubmitResponse{}
	path := "/storeDetailAndSubmitThirdParty"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.storeDetailAndSubmitRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by InitializationApi.SubmitThirdParty
type InitializationApiSubmitThirdPartyInput struct {
	submitRequest *SubmitRequest
}

func (r InitializationApiSubmitThirdPartyInput) SubmitRequest(submitRequest SubmitRequest) InitializationApiSubmitThirdPartyInput {
	r.submitRequest = &submitRequest
	return r
}

/*
Prepare a request for SubmitThirdParty

@return InitializationApiSubmitThirdPartyInput
*/
func (a *InitializationApi) SubmitThirdPartyInput() InitializationApiSubmitThirdPartyInput {
	return InitializationApiSubmitThirdPartyInput{}
}

/*
SubmitThirdParty Submit a payout

> This endpoint is **deprecated** and no longer supports new integrations. Do one of the following:
>- If you are building a new integration, use the POST [/transfers](https://docs.adyen.com/api-explorer/transfers/latest/post/transfers) endpoint instead.
> - If you are already using the Payout API, reach out to your Adyen contact to learn how to migrate to the Transfers API.
>
> With the Transfers API, you can:
> - Handle multiple payout use cases with a single API.
> - Use new payout functionalities, such as instant payouts to bank accounts.
> - Receive webhooks with more details and defined transfer states.
>
> For more information about the payout features of the Transfers API, see our [Payouts](https://docs.adyen.com/payouts/payout-service) documentation.


Submits a payout using the previously stored payment details. To store payment details, use the `/storeDetail` API call.

The submitted payout must be confirmed or declined either by a reviewer or via `/confirmThirdParty` or `/declineThirdParty` calls.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r InitializationApiSubmitThirdPartyInput - Request parameters, see SubmitThirdPartyInput
@return SubmitResponse, *http.Response, error
*/
func (a *InitializationApi) SubmitThirdParty(ctx context.Context, r InitializationApiSubmitThirdPartyInput) (SubmitResponse, *http.Response, error) {
	res := &SubmitResponse{}
	path := "/submitThirdParty"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.submitRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
