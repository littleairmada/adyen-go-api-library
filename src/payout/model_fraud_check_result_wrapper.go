/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the FraudCheckResultWrapper type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FraudCheckResultWrapper{}

// FraudCheckResultWrapper struct for FraudCheckResultWrapper
type FraudCheckResultWrapper struct {
	FraudCheckResult *FraudCheckResult `json:"FraudCheckResult,omitempty"`
}

// NewFraudCheckResultWrapper instantiates a new FraudCheckResultWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFraudCheckResultWrapper() *FraudCheckResultWrapper {
	this := FraudCheckResultWrapper{}
	return &this
}

// NewFraudCheckResultWrapperWithDefaults instantiates a new FraudCheckResultWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFraudCheckResultWrapperWithDefaults() *FraudCheckResultWrapper {
	this := FraudCheckResultWrapper{}
	return &this
}

// GetFraudCheckResult returns the FraudCheckResult field value if set, zero value otherwise.
func (o *FraudCheckResultWrapper) GetFraudCheckResult() FraudCheckResult {
	if o == nil || common.IsNil(o.FraudCheckResult) {
		var ret FraudCheckResult
		return ret
	}
	return *o.FraudCheckResult
}

// GetFraudCheckResultOk returns a tuple with the FraudCheckResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FraudCheckResultWrapper) GetFraudCheckResultOk() (*FraudCheckResult, bool) {
	if o == nil || common.IsNil(o.FraudCheckResult) {
		return nil, false
	}
	return o.FraudCheckResult, true
}

// HasFraudCheckResult returns a boolean if a field has been set.
func (o *FraudCheckResultWrapper) HasFraudCheckResult() bool {
	if o != nil && !common.IsNil(o.FraudCheckResult) {
		return true
	}

	return false
}

// SetFraudCheckResult gets a reference to the given FraudCheckResult and assigns it to the FraudCheckResult field.
func (o *FraudCheckResultWrapper) SetFraudCheckResult(v FraudCheckResult) {
	o.FraudCheckResult = &v
}

func (o FraudCheckResultWrapper) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FraudCheckResultWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FraudCheckResult) {
		toSerialize["FraudCheckResult"] = o.FraudCheckResult
	}
	return toSerialize, nil
}

type NullableFraudCheckResultWrapper struct {
	value *FraudCheckResultWrapper
	isSet bool
}

func (v NullableFraudCheckResultWrapper) Get() *FraudCheckResultWrapper {
	return v.value
}

func (v *NullableFraudCheckResultWrapper) Set(val *FraudCheckResultWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableFraudCheckResultWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableFraudCheckResultWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFraudCheckResultWrapper(val *FraudCheckResultWrapper) *NullableFraudCheckResultWrapper {
	return &NullableFraudCheckResultWrapper{value: val, isSet: true}
}

func (v NullableFraudCheckResultWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFraudCheckResultWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
