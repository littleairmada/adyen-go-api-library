/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the Balance type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Balance{}

// Balance struct for Balance
type Balance struct {
	// The balance available for use.
	Available int64 `json:"available"`
	// The sum of transactions that have already been settled.
	Balance int64 `json:"balance"`
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) of the balance.
	Currency string `json:"currency"`
	// The sum of the transactions that will be settled in the future.
	Pending *int64 `json:"pending,omitempty"`
	// The balance currently held in reserve.
	Reserved int64 `json:"reserved"`
}

// NewBalance instantiates a new Balance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalance(available int64, balance int64, currency string, reserved int64) *Balance {
	this := Balance{}
	this.Available = available
	this.Balance = balance
	this.Currency = currency
	this.Reserved = reserved
	return &this
}

// NewBalanceWithDefaults instantiates a new Balance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceWithDefaults() *Balance {
	this := Balance{}
	return &this
}

// GetAvailable returns the Available field value
func (o *Balance) GetAvailable() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *Balance) GetAvailableOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *Balance) SetAvailable(v int64) {
	o.Available = v
}

// GetBalance returns the Balance field value
func (o *Balance) GetBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *Balance) GetBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *Balance) SetBalance(v int64) {
	o.Balance = v
}

// GetCurrency returns the Currency field value
func (o *Balance) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Balance) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Balance) SetCurrency(v string) {
	o.Currency = v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *Balance) GetPending() int64 {
	if o == nil || common.IsNil(o.Pending) {
		var ret int64
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balance) GetPendingOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *Balance) HasPending() bool {
	if o != nil && !common.IsNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given int64 and assigns it to the Pending field.
func (o *Balance) SetPending(v int64) {
	o.Pending = &v
}

// GetReserved returns the Reserved field value
func (o *Balance) GetReserved() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value
// and a boolean to check if the value has been set.
func (o *Balance) GetReservedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reserved, true
}

// SetReserved sets field value
func (o *Balance) SetReserved(v int64) {
	o.Reserved = v
}

func (o Balance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Balance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["available"] = o.Available
	toSerialize["balance"] = o.Balance
	toSerialize["currency"] = o.Currency
	if !common.IsNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	toSerialize["reserved"] = o.Reserved
	return toSerialize, nil
}

type NullableBalance struct {
	value *Balance
	isSet bool
}

func (v NullableBalance) Get() *Balance {
	return v.value
}

func (v *NullableBalance) Set(val *Balance) {
	v.value = val
	v.isSet = true
}

func (v NullableBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalance(val *Balance) *NullableBalance {
	return &NullableBalance{value: val, isSet: true}
}

func (v NullableBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
