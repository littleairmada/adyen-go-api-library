/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// GeneralApi service
type GeneralApi common.Service

// All parameters accepted by GeneralApi.Authorise
type GeneralApiAuthoriseInput struct {
	paymentRequest *PaymentRequest
}

func (r GeneralApiAuthoriseInput) PaymentRequest(paymentRequest PaymentRequest) GeneralApiAuthoriseInput {
	r.paymentRequest = &paymentRequest
	return r
}

/*
Prepare a request for Authorise

@return GeneralApiAuthoriseInput
*/
func (a *GeneralApi) AuthoriseInput() GeneralApiAuthoriseInput {
	return GeneralApiAuthoriseInput{}
}

/*
Authorise Create an authorisation

Creates a payment with a unique reference (`pspReference`) and attempts to obtain an authorisation hold. For cards, this amount can be captured or cancelled later. Non-card payment methods typically don't support this and will automatically capture as part of the authorisation.
> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/payments`](https://docs.adyen.com/api-explorer/#/CheckoutService/payments) endpoint under Checkout API instead.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiAuthoriseInput - Request parameters, see AuthoriseInput
@return PaymentResult, *http.Response, error
*/
func (a *GeneralApi) Authorise(ctx context.Context, r GeneralApiAuthoriseInput) (PaymentResult, *http.Response, error) {
	res := &PaymentResult{}
	path := "/authorise"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.Authorise3d
type GeneralApiAuthorise3dInput struct {
	paymentRequest3d *PaymentRequest3d
}

func (r GeneralApiAuthorise3dInput) PaymentRequest3d(paymentRequest3d PaymentRequest3d) GeneralApiAuthorise3dInput {
	r.paymentRequest3d = &paymentRequest3d
	return r
}

/*
Prepare a request for Authorise3d

@return GeneralApiAuthorise3dInput
*/
func (a *GeneralApi) Authorise3dInput() GeneralApiAuthorise3dInput {
	return GeneralApiAuthorise3dInput{}
}

/*
Authorise3d Complete a 3DS authorisation

For an authenticated 3D Secure session, completes the payment authorisation. This endpoint must receive the `md` and `paResponse` parameters that you get from the card issuer after a shopper pays via 3D Secure.

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce/3d-secure). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/payments/details`](https://docs.adyen.com/api-explorer/#/CheckoutService/payments/details) endpoint under Checkout API instead.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiAuthorise3dInput - Request parameters, see Authorise3dInput
@return PaymentResult, *http.Response, error
*/
func (a *GeneralApi) Authorise3d(ctx context.Context, r GeneralApiAuthorise3dInput) (PaymentResult, *http.Response, error) {
	res := &PaymentResult{}
	path := "/authorise3d"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentRequest3d,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.Authorise3ds2
type GeneralApiAuthorise3ds2Input struct {
	paymentRequest3ds2 *PaymentRequest3ds2
}

func (r GeneralApiAuthorise3ds2Input) PaymentRequest3ds2(paymentRequest3ds2 PaymentRequest3ds2) GeneralApiAuthorise3ds2Input {
	r.paymentRequest3ds2 = &paymentRequest3ds2
	return r
}

/*
Prepare a request for Authorise3ds2

@return GeneralApiAuthorise3ds2Input
*/
func (a *GeneralApi) Authorise3ds2Input() GeneralApiAuthorise3ds2Input {
	return GeneralApiAuthorise3ds2Input{}
}

/*
Authorise3ds2 Complete a 3DS2 authorisation

For an authenticated 3D Secure 2 session, completes the payment authorisation. This endpoint must receive the `threeDS2Token` and `threeDS2Result` parameters.

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce/3d-secure). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/payments/details`](https://docs.adyen.com/api-explorer/#/CheckoutService/payments/details) endpoint under Checkout API instead.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiAuthorise3ds2Input - Request parameters, see Authorise3ds2Input
@return PaymentResult, *http.Response, error
*/
func (a *GeneralApi) Authorise3ds2(ctx context.Context, r GeneralApiAuthorise3ds2Input) (PaymentResult, *http.Response, error) {
	res := &PaymentResult{}
	path := "/authorise3ds2"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentRequest3ds2,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.GetAuthenticationResult
type GeneralApiGetAuthenticationResultInput struct {
	authenticationResultRequest *AuthenticationResultRequest
}

func (r GeneralApiGetAuthenticationResultInput) AuthenticationResultRequest(authenticationResultRequest AuthenticationResultRequest) GeneralApiGetAuthenticationResultInput {
	r.authenticationResultRequest = &authenticationResultRequest
	return r
}

/*
Prepare a request for GetAuthenticationResult

@return GeneralApiGetAuthenticationResultInput
*/
func (a *GeneralApi) GetAuthenticationResultInput() GeneralApiGetAuthenticationResultInput {
	return GeneralApiGetAuthenticationResultInput{}
}

/*
GetAuthenticationResult Get the 3DS authentication result

Return the authentication result after doing a 3D Secure authentication only.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiGetAuthenticationResultInput - Request parameters, see GetAuthenticationResultInput
@return AuthenticationResultResponse, *http.Response, error
*/
func (a *GeneralApi) GetAuthenticationResult(ctx context.Context, r GeneralApiGetAuthenticationResultInput) (AuthenticationResultResponse, *http.Response, error) {
	res := &AuthenticationResultResponse{}
	path := "/getAuthenticationResult"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.authenticationResultRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.Retrieve3ds2Result
type GeneralApiRetrieve3ds2ResultInput struct {
	threeDS2ResultRequest *ThreeDS2ResultRequest
}

func (r GeneralApiRetrieve3ds2ResultInput) ThreeDS2ResultRequest(threeDS2ResultRequest ThreeDS2ResultRequest) GeneralApiRetrieve3ds2ResultInput {
	r.threeDS2ResultRequest = &threeDS2ResultRequest
	return r
}

/*
Prepare a request for Retrieve3ds2Result

@return GeneralApiRetrieve3ds2ResultInput
*/
func (a *GeneralApi) Retrieve3ds2ResultInput() GeneralApiRetrieve3ds2ResultInput {
	return GeneralApiRetrieve3ds2ResultInput{}
}

/*
Retrieve3ds2Result Get the 3DS2 authentication result

Retrieves the `threeDS2Result` after doing a 3D Secure 2 authentication only.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiRetrieve3ds2ResultInput - Request parameters, see Retrieve3ds2ResultInput
@return ThreeDS2ResultResponse, *http.Response, error
*/
func (a *GeneralApi) Retrieve3ds2Result(ctx context.Context, r GeneralApiRetrieve3ds2ResultInput) (ThreeDS2ResultResponse, *http.Response, error) {
	res := &ThreeDS2ResultResponse{}
	path := "/retrieve3ds2Result"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.threeDS2ResultRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
