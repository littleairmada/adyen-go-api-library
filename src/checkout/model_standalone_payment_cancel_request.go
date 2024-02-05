/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the StandalonePaymentCancelRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StandalonePaymentCancelRequest{}

// StandalonePaymentCancelRequest struct for StandalonePaymentCancelRequest
type StandalonePaymentCancelRequest struct {
	ApplicationInfo *ApplicationInfo `json:"applicationInfo,omitempty"`
	// The merchant account that is used to process the payment.
	MerchantAccount string `json:"merchantAccount"`
	// The [`reference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__reqParam_reference) of the payment that you want to cancel.
	PaymentReference string `json:"paymentReference"`
	// Your reference for the cancel request. Maximum length: 80 characters.
	Reference *string `json:"reference,omitempty"`
}

// NewStandalonePaymentCancelRequest instantiates a new StandalonePaymentCancelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandalonePaymentCancelRequest(merchantAccount string, paymentReference string) *StandalonePaymentCancelRequest {
	this := StandalonePaymentCancelRequest{}
	this.MerchantAccount = merchantAccount
	this.PaymentReference = paymentReference
	return &this
}

// NewStandalonePaymentCancelRequestWithDefaults instantiates a new StandalonePaymentCancelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandalonePaymentCancelRequestWithDefaults() *StandalonePaymentCancelRequest {
	this := StandalonePaymentCancelRequest{}
	return &this
}

// GetApplicationInfo returns the ApplicationInfo field value if set, zero value otherwise.
func (o *StandalonePaymentCancelRequest) GetApplicationInfo() ApplicationInfo {
	if o == nil || common.IsNil(o.ApplicationInfo) {
		var ret ApplicationInfo
		return ret
	}
	return *o.ApplicationInfo
}

// GetApplicationInfoOk returns a tuple with the ApplicationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandalonePaymentCancelRequest) GetApplicationInfoOk() (*ApplicationInfo, bool) {
	if o == nil || common.IsNil(o.ApplicationInfo) {
		return nil, false
	}
	return o.ApplicationInfo, true
}

// HasApplicationInfo returns a boolean if a field has been set.
func (o *StandalonePaymentCancelRequest) HasApplicationInfo() bool {
	if o != nil && !common.IsNil(o.ApplicationInfo) {
		return true
	}

	return false
}

// SetApplicationInfo gets a reference to the given ApplicationInfo and assigns it to the ApplicationInfo field.
func (o *StandalonePaymentCancelRequest) SetApplicationInfo(v ApplicationInfo) {
	o.ApplicationInfo = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *StandalonePaymentCancelRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *StandalonePaymentCancelRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *StandalonePaymentCancelRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetPaymentReference returns the PaymentReference field value
func (o *StandalonePaymentCancelRequest) GetPaymentReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentReference
}

// GetPaymentReferenceOk returns a tuple with the PaymentReference field value
// and a boolean to check if the value has been set.
func (o *StandalonePaymentCancelRequest) GetPaymentReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentReference, true
}

// SetPaymentReference sets field value
func (o *StandalonePaymentCancelRequest) SetPaymentReference(v string) {
	o.PaymentReference = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *StandalonePaymentCancelRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandalonePaymentCancelRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *StandalonePaymentCancelRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *StandalonePaymentCancelRequest) SetReference(v string) {
	o.Reference = &v
}

func (o StandalonePaymentCancelRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandalonePaymentCancelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ApplicationInfo) {
		toSerialize["applicationInfo"] = o.ApplicationInfo
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["paymentReference"] = o.PaymentReference
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	return toSerialize, nil
}

type NullableStandalonePaymentCancelRequest struct {
	value *StandalonePaymentCancelRequest
	isSet bool
}

func (v NullableStandalonePaymentCancelRequest) Get() *StandalonePaymentCancelRequest {
	return v.value
}

func (v *NullableStandalonePaymentCancelRequest) Set(val *StandalonePaymentCancelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStandalonePaymentCancelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStandalonePaymentCancelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandalonePaymentCancelRequest(val *StandalonePaymentCancelRequest) *NullableStandalonePaymentCancelRequest {
	return &NullableStandalonePaymentCancelRequest{value: val, isSet: true}
}

func (v NullableStandalonePaymentCancelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandalonePaymentCancelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
