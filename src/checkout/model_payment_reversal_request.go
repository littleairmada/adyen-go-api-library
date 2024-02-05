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

// checks if the PaymentReversalRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentReversalRequest{}

// PaymentReversalRequest struct for PaymentReversalRequest
type PaymentReversalRequest struct {
	ApplicationInfo *ApplicationInfo `json:"applicationInfo,omitempty"`
	// The merchant account that is used to process the payment.
	MerchantAccount string `json:"merchantAccount"`
	// Your reference for the reversal request. Maximum length: 80 characters.
	Reference *string `json:"reference,omitempty"`
}

// NewPaymentReversalRequest instantiates a new PaymentReversalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentReversalRequest(merchantAccount string) *PaymentReversalRequest {
	this := PaymentReversalRequest{}
	this.MerchantAccount = merchantAccount
	return &this
}

// NewPaymentReversalRequestWithDefaults instantiates a new PaymentReversalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentReversalRequestWithDefaults() *PaymentReversalRequest {
	this := PaymentReversalRequest{}
	return &this
}

// GetApplicationInfo returns the ApplicationInfo field value if set, zero value otherwise.
func (o *PaymentReversalRequest) GetApplicationInfo() ApplicationInfo {
	if o == nil || common.IsNil(o.ApplicationInfo) {
		var ret ApplicationInfo
		return ret
	}
	return *o.ApplicationInfo
}

// GetApplicationInfoOk returns a tuple with the ApplicationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentReversalRequest) GetApplicationInfoOk() (*ApplicationInfo, bool) {
	if o == nil || common.IsNil(o.ApplicationInfo) {
		return nil, false
	}
	return o.ApplicationInfo, true
}

// HasApplicationInfo returns a boolean if a field has been set.
func (o *PaymentReversalRequest) HasApplicationInfo() bool {
	if o != nil && !common.IsNil(o.ApplicationInfo) {
		return true
	}

	return false
}

// SetApplicationInfo gets a reference to the given ApplicationInfo and assigns it to the ApplicationInfo field.
func (o *PaymentReversalRequest) SetApplicationInfo(v ApplicationInfo) {
	o.ApplicationInfo = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *PaymentReversalRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *PaymentReversalRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *PaymentReversalRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentReversalRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentReversalRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentReversalRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentReversalRequest) SetReference(v string) {
	o.Reference = &v
}

func (o PaymentReversalRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentReversalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ApplicationInfo) {
		toSerialize["applicationInfo"] = o.ApplicationInfo
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	return toSerialize, nil
}

type NullablePaymentReversalRequest struct {
	value *PaymentReversalRequest
	isSet bool
}

func (v NullablePaymentReversalRequest) Get() *PaymentReversalRequest {
	return v.value
}

func (v *NullablePaymentReversalRequest) Set(val *PaymentReversalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentReversalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentReversalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentReversalRequest(val *PaymentReversalRequest) *NullablePaymentReversalRequest {
	return &NullablePaymentReversalRequest{value: val, isSet: true}
}

func (v NullablePaymentReversalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentReversalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
