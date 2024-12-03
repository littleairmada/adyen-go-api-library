/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the Payment type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Payment{}

// Payment struct for Payment
type Payment struct {
	// The default currency for contactless payments on the payment terminal, as the three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.
	ContactlessCurrency *string `json:"contactlessCurrency,omitempty"`
	// Hides the minor units for the listed [ISO currency codes](https://en.wikipedia.org/wiki/ISO_4217).
	HideMinorUnitsInCurrencies []string `json:"hideMinorUnitsInCurrencies,omitempty"`
}

// NewPayment instantiates a new Payment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayment() *Payment {
	this := Payment{}
	return &this
}

// NewPaymentWithDefaults instantiates a new Payment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentWithDefaults() *Payment {
	this := Payment{}
	return &this
}

// GetContactlessCurrency returns the ContactlessCurrency field value if set, zero value otherwise.
func (o *Payment) GetContactlessCurrency() string {
	if o == nil || common.IsNil(o.ContactlessCurrency) {
		var ret string
		return ret
	}
	return *o.ContactlessCurrency
}

// GetContactlessCurrencyOk returns a tuple with the ContactlessCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetContactlessCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContactlessCurrency) {
		return nil, false
	}
	return o.ContactlessCurrency, true
}

// HasContactlessCurrency returns a boolean if a field has been set.
func (o *Payment) HasContactlessCurrency() bool {
	if o != nil && !common.IsNil(o.ContactlessCurrency) {
		return true
	}

	return false
}

// SetContactlessCurrency gets a reference to the given string and assigns it to the ContactlessCurrency field.
func (o *Payment) SetContactlessCurrency(v string) {
	o.ContactlessCurrency = &v
}

// GetHideMinorUnitsInCurrencies returns the HideMinorUnitsInCurrencies field value if set, zero value otherwise.
func (o *Payment) GetHideMinorUnitsInCurrencies() []string {
	if o == nil || common.IsNil(o.HideMinorUnitsInCurrencies) {
		var ret []string
		return ret
	}
	return o.HideMinorUnitsInCurrencies
}

// GetHideMinorUnitsInCurrenciesOk returns a tuple with the HideMinorUnitsInCurrencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetHideMinorUnitsInCurrenciesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.HideMinorUnitsInCurrencies) {
		return nil, false
	}
	return o.HideMinorUnitsInCurrencies, true
}

// HasHideMinorUnitsInCurrencies returns a boolean if a field has been set.
func (o *Payment) HasHideMinorUnitsInCurrencies() bool {
	if o != nil && !common.IsNil(o.HideMinorUnitsInCurrencies) {
		return true
	}

	return false
}

// SetHideMinorUnitsInCurrencies gets a reference to the given []string and assigns it to the HideMinorUnitsInCurrencies field.
func (o *Payment) SetHideMinorUnitsInCurrencies(v []string) {
	o.HideMinorUnitsInCurrencies = v
}

func (o Payment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Payment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ContactlessCurrency) {
		toSerialize["contactlessCurrency"] = o.ContactlessCurrency
	}
	if !common.IsNil(o.HideMinorUnitsInCurrencies) {
		toSerialize["hideMinorUnitsInCurrencies"] = o.HideMinorUnitsInCurrencies
	}
	return toSerialize, nil
}

type NullablePayment struct {
	value *Payment
	isSet bool
}

func (v NullablePayment) Get() *Payment {
	return v.value
}

func (v *NullablePayment) Set(val *Payment) {
	v.value = val
	v.isSet = true
}

func (v NullablePayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayment(val *Payment) *NullablePayment {
	return &NullablePayment{value: val, isSet: true}
}

func (v NullablePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



