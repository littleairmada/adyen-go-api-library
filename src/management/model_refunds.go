/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the Refunds type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Refunds{}

// Refunds struct for Refunds
type Refunds struct {
	Referenced *Referenced `json:"referenced,omitempty"`
}

// NewRefunds instantiates a new Refunds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefunds() *Refunds {
	this := Refunds{}
	return &this
}

// NewRefundsWithDefaults instantiates a new Refunds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundsWithDefaults() *Refunds {
	this := Refunds{}
	return &this
}

// GetReferenced returns the Referenced field value if set, zero value otherwise.
func (o *Refunds) GetReferenced() Referenced {
	if o == nil || common.IsNil(o.Referenced) {
		var ret Referenced
		return ret
	}
	return *o.Referenced
}

// GetReferencedOk returns a tuple with the Referenced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refunds) GetReferencedOk() (*Referenced, bool) {
	if o == nil || common.IsNil(o.Referenced) {
		return nil, false
	}
	return o.Referenced, true
}

// HasReferenced returns a boolean if a field has been set.
func (o *Refunds) HasReferenced() bool {
	if o != nil && !common.IsNil(o.Referenced) {
		return true
	}

	return false
}

// SetReferenced gets a reference to the given Referenced and assigns it to the Referenced field.
func (o *Refunds) SetReferenced(v Referenced) {
	o.Referenced = &v
}

func (o Refunds) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Refunds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Referenced) {
		toSerialize["referenced"] = o.Referenced
	}
	return toSerialize, nil
}

type NullableRefunds struct {
	value *Refunds
	isSet bool
}

func (v NullableRefunds) Get() *Refunds {
	return v.value
}

func (v *NullableRefunds) Set(val *Refunds) {
	v.value = val
	v.isSet = true
}

func (v NullableRefunds) IsSet() bool {
	return v.isSet
}

func (v *NullableRefunds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefunds(val *Refunds) *NullableRefunds {
	return &NullableRefunds{value: val, isSet: true}
}

func (v NullableRefunds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefunds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
