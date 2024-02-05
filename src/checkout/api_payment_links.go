/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// PaymentLinksApi service
type PaymentLinksApi common.Service

// All parameters accepted by PaymentLinksApi.GetPaymentLink
type PaymentLinksApiGetPaymentLinkInput struct {
	linkId string
}

/*
Prepare a request for GetPaymentLink
@param linkId Unique identifier of the payment link.
@return PaymentLinksApiGetPaymentLinkInput
*/
func (a *PaymentLinksApi) GetPaymentLinkInput(linkId string) PaymentLinksApiGetPaymentLinkInput {
	return PaymentLinksApiGetPaymentLinkInput{
		linkId: linkId,
	}
}

/*
GetPaymentLink Get a payment link

Retrieves the payment link details using the payment link `id`.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentLinksApiGetPaymentLinkInput - Request parameters, see GetPaymentLinkInput
@return PaymentLinkResponse, *http.Response, error
*/
func (a *PaymentLinksApi) GetPaymentLink(ctx context.Context, r PaymentLinksApiGetPaymentLinkInput) (PaymentLinkResponse, *http.Response, error) {
	res := &PaymentLinkResponse{}
	path := "/paymentLinks/{linkId}"
	path = strings.Replace(path, "{"+"linkId"+"}", url.PathEscape(common.ParameterValueToString(r.linkId, "linkId")), -1)
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

	return *res, httpRes, err
}

// All parameters accepted by PaymentLinksApi.PaymentLinks
type PaymentLinksApiPaymentLinksInput struct {
	idempotencyKey     *string
	paymentLinkRequest *PaymentLinkRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r PaymentLinksApiPaymentLinksInput) IdempotencyKey(idempotencyKey string) PaymentLinksApiPaymentLinksInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r PaymentLinksApiPaymentLinksInput) PaymentLinkRequest(paymentLinkRequest PaymentLinkRequest) PaymentLinksApiPaymentLinksInput {
	r.paymentLinkRequest = &paymentLinkRequest
	return r
}

/*
Prepare a request for PaymentLinks

@return PaymentLinksApiPaymentLinksInput
*/
func (a *PaymentLinksApi) PaymentLinksInput() PaymentLinksApiPaymentLinksInput {
	return PaymentLinksApiPaymentLinksInput{}
}

/*
PaymentLinks Create a payment link

Creates a payment link to our hosted payment form where shoppers can pay. The list of payment methods presented to the shopper depends on the `currency` and `country` parameters sent in the request.

For more information, refer to [Pay by Link documentation](https://docs.adyen.com/online-payments/pay-by-link#create-payment-links-through-api).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentLinksApiPaymentLinksInput - Request parameters, see PaymentLinksInput
@return PaymentLinkResponse, *http.Response, error
*/
func (a *PaymentLinksApi) PaymentLinks(ctx context.Context, r PaymentLinksApiPaymentLinksInput) (PaymentLinkResponse, *http.Response, error) {
	res := &PaymentLinkResponse{}
	path := "/paymentLinks"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentLinkRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by PaymentLinksApi.UpdatePaymentLink
type PaymentLinksApiUpdatePaymentLinkInput struct {
	linkId                   string
	updatePaymentLinkRequest *UpdatePaymentLinkRequest
}

func (r PaymentLinksApiUpdatePaymentLinkInput) UpdatePaymentLinkRequest(updatePaymentLinkRequest UpdatePaymentLinkRequest) PaymentLinksApiUpdatePaymentLinkInput {
	r.updatePaymentLinkRequest = &updatePaymentLinkRequest
	return r
}

/*
Prepare a request for UpdatePaymentLink
@param linkId Unique identifier of the payment link.
@return PaymentLinksApiUpdatePaymentLinkInput
*/
func (a *PaymentLinksApi) UpdatePaymentLinkInput(linkId string) PaymentLinksApiUpdatePaymentLinkInput {
	return PaymentLinksApiUpdatePaymentLinkInput{
		linkId: linkId,
	}
}

/*
UpdatePaymentLink Update the status of a payment link

Updates the status of a payment link. Use this endpoint to [force the expiry of a payment link](https://docs.adyen.com/online-payments/pay-by-link#update-payment-link-status).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentLinksApiUpdatePaymentLinkInput - Request parameters, see UpdatePaymentLinkInput
@return PaymentLinkResponse, *http.Response, error
*/
func (a *PaymentLinksApi) UpdatePaymentLink(ctx context.Context, r PaymentLinksApiUpdatePaymentLinkInput) (PaymentLinkResponse, *http.Response, error) {
	res := &PaymentLinkResponse{}
	path := "/paymentLinks/{linkId}"
	path = strings.Replace(path, "{"+"linkId"+"}", url.PathEscape(common.ParameterValueToString(r.linkId, "linkId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updatePaymentLinkRequest,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
