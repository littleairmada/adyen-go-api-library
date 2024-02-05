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

// checks if the BankAccount type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BankAccount{}

// BankAccount struct for BankAccount
type BankAccount struct {
	AccountIdentification BankAccountAccountIdentification `json:"accountIdentification"`
}

// NewBankAccount instantiates a new BankAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccount(accountIdentification BankAccountAccountIdentification) *BankAccount {
	this := BankAccount{}
	this.AccountIdentification = accountIdentification
	return &this
}

// NewBankAccountWithDefaults instantiates a new BankAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountWithDefaults() *BankAccount {
	this := BankAccount{}
	return &this
}

// GetAccountIdentification returns the AccountIdentification field value
func (o *BankAccount) GetAccountIdentification() BankAccountAccountIdentification {
	if o == nil {
		var ret BankAccountAccountIdentification
		return ret
	}

	return o.AccountIdentification
}

// GetAccountIdentificationOk returns a tuple with the AccountIdentification field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetAccountIdentificationOk() (*BankAccountAccountIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentification, true
}

// SetAccountIdentification sets field value
func (o *BankAccount) SetAccountIdentification(v BankAccountAccountIdentification) {
	o.AccountIdentification = v
}

func (o BankAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountIdentification"] = o.AccountIdentification
	return toSerialize, nil
}

type NullableBankAccount struct {
	value *BankAccount
	isSet bool
}

func (v NullableBankAccount) Get() *BankAccount {
	return v.value
}

func (v *NullableBankAccount) Set(val *BankAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccount(val *BankAccount) *NullableBankAccount {
	return &NullableBankAccount{value: val, isSet: true}
}

func (v NullableBankAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
