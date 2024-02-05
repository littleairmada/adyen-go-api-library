/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the ThreeDS2ResultRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThreeDS2ResultRequest{}

// ThreeDS2ResultRequest struct for ThreeDS2ResultRequest
type ThreeDS2ResultRequest struct {
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The pspReference returned in the /authorise call.
	PspReference string `json:"pspReference"`
}

// NewThreeDS2ResultRequest instantiates a new ThreeDS2ResultRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDS2ResultRequest(merchantAccount string, pspReference string) *ThreeDS2ResultRequest {
	this := ThreeDS2ResultRequest{}
	this.MerchantAccount = merchantAccount
	this.PspReference = pspReference
	return &this
}

// NewThreeDS2ResultRequestWithDefaults instantiates a new ThreeDS2ResultRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDS2ResultRequestWithDefaults() *ThreeDS2ResultRequest {
	this := ThreeDS2ResultRequest{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *ThreeDS2ResultRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResultRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *ThreeDS2ResultRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetPspReference returns the PspReference field value
func (o *ThreeDS2ResultRequest) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResultRequest) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *ThreeDS2ResultRequest) SetPspReference(v string) {
	o.PspReference = v
}

func (o ThreeDS2ResultRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDS2ResultRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["pspReference"] = o.PspReference
	return toSerialize, nil
}

type NullableThreeDS2ResultRequest struct {
	value *ThreeDS2ResultRequest
	isSet bool
}

func (v NullableThreeDS2ResultRequest) Get() *ThreeDS2ResultRequest {
	return v.value
}

func (v *NullableThreeDS2ResultRequest) Set(val *ThreeDS2ResultRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDS2ResultRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDS2ResultRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDS2ResultRequest(val *ThreeDS2ResultRequest) *NullableThreeDS2ResultRequest {
	return &NullableThreeDS2ResultRequest{value: val, isSet: true}
}

func (v NullableThreeDS2ResultRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDS2ResultRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
