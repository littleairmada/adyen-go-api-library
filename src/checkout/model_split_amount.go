/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the SplitAmount type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SplitAmount{}

// SplitAmount struct for SplitAmount
type SplitAmount struct {
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes). By default, this is the original payment currency.
	Currency *string `json:"currency,omitempty"`
	// The value of the split amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes).
	Value int64 `json:"value"`
}

// NewSplitAmount instantiates a new SplitAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitAmount(value int64) *SplitAmount {
	this := SplitAmount{}
	this.Value = value
	return &this
}

// NewSplitAmountWithDefaults instantiates a new SplitAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitAmountWithDefaults() *SplitAmount {
	this := SplitAmount{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *SplitAmount) GetCurrency() string {
	if o == nil || common.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitAmount) GetCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *SplitAmount) HasCurrency() bool {
	if o != nil && !common.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *SplitAmount) SetCurrency(v string) {
	o.Currency = &v
}

// GetValue returns the Value field value
func (o *SplitAmount) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SplitAmount) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SplitAmount) SetValue(v int64) {
	o.Value = v
}

func (o SplitAmount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableSplitAmount struct {
	value *SplitAmount
	isSet bool
}

func (v NullableSplitAmount) Get() *SplitAmount {
	return v.value
}

func (v *NullableSplitAmount) Set(val *SplitAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitAmount(val *SplitAmount) *NullableSplitAmount {
	return &NullableSplitAmount{value: val, isSet: true}
}

func (v NullableSplitAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
