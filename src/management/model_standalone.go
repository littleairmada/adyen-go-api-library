/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the Standalone type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Standalone{}

// Standalone struct for Standalone
type Standalone struct {
	// The default currency of the standalone payment terminal as an [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Enable standalone mode.
	EnableStandalone *bool `json:"enableStandalone,omitempty"`
}

// NewStandalone instantiates a new Standalone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandalone() *Standalone {
	this := Standalone{}
	return &this
}

// NewStandaloneWithDefaults instantiates a new Standalone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandaloneWithDefaults() *Standalone {
	this := Standalone{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *Standalone) GetCurrencyCode() string {
	if o == nil || common.IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Standalone) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *Standalone) HasCurrencyCode() bool {
	if o != nil && !common.IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *Standalone) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetEnableStandalone returns the EnableStandalone field value if set, zero value otherwise.
func (o *Standalone) GetEnableStandalone() bool {
	if o == nil || common.IsNil(o.EnableStandalone) {
		var ret bool
		return ret
	}
	return *o.EnableStandalone
}

// GetEnableStandaloneOk returns a tuple with the EnableStandalone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Standalone) GetEnableStandaloneOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnableStandalone) {
		return nil, false
	}
	return o.EnableStandalone, true
}

// HasEnableStandalone returns a boolean if a field has been set.
func (o *Standalone) HasEnableStandalone() bool {
	if o != nil && !common.IsNil(o.EnableStandalone) {
		return true
	}

	return false
}

// SetEnableStandalone gets a reference to the given bool and assigns it to the EnableStandalone field.
func (o *Standalone) SetEnableStandalone(v bool) {
	o.EnableStandalone = &v
}

func (o Standalone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Standalone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !common.IsNil(o.EnableStandalone) {
		toSerialize["enableStandalone"] = o.EnableStandalone
	}
	return toSerialize, nil
}

type NullableStandalone struct {
	value *Standalone
	isSet bool
}

func (v NullableStandalone) Get() *Standalone {
	return v.value
}

func (v *NullableStandalone) Set(val *Standalone) {
	v.value = val
	v.isSet = true
}

func (v NullableStandalone) IsSet() bool {
	return v.isSet
}

func (v *NullableStandalone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandalone(val *Standalone) *NullableStandalone {
	return &NullableStandalone{value: val, isSet: true}
}

func (v NullableStandalone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandalone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
