/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the TotalAmountRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TotalAmountRestriction{}

// TotalAmountRestriction struct for TotalAmountRestriction
type TotalAmountRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string  `json:"operation"`
	Value     *Amount `json:"value,omitempty"`
}

// NewTotalAmountRestriction instantiates a new TotalAmountRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTotalAmountRestriction(operation string) *TotalAmountRestriction {
	this := TotalAmountRestriction{}
	this.Operation = operation
	return &this
}

// NewTotalAmountRestrictionWithDefaults instantiates a new TotalAmountRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotalAmountRestrictionWithDefaults() *TotalAmountRestriction {
	this := TotalAmountRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *TotalAmountRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *TotalAmountRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *TotalAmountRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TotalAmountRestriction) GetValue() Amount {
	if o == nil || common.IsNil(o.Value) {
		var ret Amount
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TotalAmountRestriction) GetValueOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TotalAmountRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given Amount and assigns it to the Value field.
func (o *TotalAmountRestriction) SetValue(v Amount) {
	o.Value = &v
}

func (o TotalAmountRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TotalAmountRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTotalAmountRestriction struct {
	value *TotalAmountRestriction
	isSet bool
}

func (v NullableTotalAmountRestriction) Get() *TotalAmountRestriction {
	return v.value
}

func (v *NullableTotalAmountRestriction) Set(val *TotalAmountRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableTotalAmountRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableTotalAmountRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotalAmountRestriction(val *TotalAmountRestriction) *NullableTotalAmountRestriction {
	return &NullableTotalAmountRestriction{value: val, isSet: true}
}

func (v NullableTotalAmountRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotalAmountRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
