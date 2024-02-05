/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// BankAccountValidationApi service
type BankAccountValidationApi common.Service

// All parameters accepted by BankAccountValidationApi.ValidateBankAccountIdentification
type BankAccountValidationApiValidateBankAccountIdentificationInput struct {
	bankAccountIdentificationValidationRequest *BankAccountIdentificationValidationRequest
}

func (r BankAccountValidationApiValidateBankAccountIdentificationInput) BankAccountIdentificationValidationRequest(bankAccountIdentificationValidationRequest BankAccountIdentificationValidationRequest) BankAccountValidationApiValidateBankAccountIdentificationInput {
	r.bankAccountIdentificationValidationRequest = &bankAccountIdentificationValidationRequest
	return r
}

/*
Prepare a request for ValidateBankAccountIdentification

@return BankAccountValidationApiValidateBankAccountIdentificationInput
*/
func (a *BankAccountValidationApi) ValidateBankAccountIdentificationInput() BankAccountValidationApiValidateBankAccountIdentificationInput {
	return BankAccountValidationApiValidateBankAccountIdentificationInput{}
}

/*
ValidateBankAccountIdentification Validate a bank account

Validates bank account identification details. You can use this endpoint to validate bank account details before you [make a transfer](https://docs.adyen.com/api-explorer/transfers/latest/post/transfers) or [create a transfer instrument](https://docs.adyen.com/api-explorer/legalentity/latest/post/transferInstruments).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BankAccountValidationApiValidateBankAccountIdentificationInput - Request parameters, see ValidateBankAccountIdentificationInput
@return *http.Response, error
*/
func (a *BankAccountValidationApi) ValidateBankAccountIdentification(ctx context.Context, r BankAccountValidationApiValidateBankAccountIdentificationInput) (*http.Response, error) {
	var res interface{}
	path := "/validateBankAccountIdentification"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.bankAccountIdentificationValidationRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	return httpRes, err
}
