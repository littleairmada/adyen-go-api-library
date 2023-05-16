/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the RequestActivationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RequestActivationResponse{}

// RequestActivationResponse struct for RequestActivationResponse
type RequestActivationResponse struct {
	// The unique identifier of the company account.
	CompanyId *string `json:"companyId,omitempty"`
	// The unique identifier of the merchant account you requested to activate.
	MerchantId *string `json:"merchantId,omitempty"`
}

// NewRequestActivationResponse instantiates a new RequestActivationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestActivationResponse() *RequestActivationResponse {
	this := RequestActivationResponse{}
	return &this
}

// NewRequestActivationResponseWithDefaults instantiates a new RequestActivationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestActivationResponseWithDefaults() *RequestActivationResponse {
	this := RequestActivationResponse{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *RequestActivationResponse) GetCompanyId() string {
	if o == nil || common.IsNil(o.CompanyId) {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestActivationResponse) GetCompanyIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *RequestActivationResponse) HasCompanyId() bool {
	if o != nil && !common.IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *RequestActivationResponse) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *RequestActivationResponse) GetMerchantId() string {
	if o == nil || common.IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestActivationResponse) GetMerchantIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *RequestActivationResponse) HasMerchantId() bool {
	if o != nil && !common.IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *RequestActivationResponse) SetMerchantId(v string) {
	o.MerchantId = &v
}

func (o RequestActivationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestActivationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CompanyId) {
		toSerialize["companyId"] = o.CompanyId
	}
	if !common.IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	return toSerialize, nil
}

type NullableRequestActivationResponse struct {
	value *RequestActivationResponse
	isSet bool
}

func (v NullableRequestActivationResponse) Get() *RequestActivationResponse {
	return v.value
}

func (v *NullableRequestActivationResponse) Set(val *RequestActivationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestActivationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestActivationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestActivationResponse(val *RequestActivationResponse) *NullableRequestActivationResponse {
	return &NullableRequestActivationResponse{value: val, isSet: true}
}

func (v NullableRequestActivationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestActivationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
