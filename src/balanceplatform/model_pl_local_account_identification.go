/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the PLLocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PLLocalAccountIdentification{}

// PLLocalAccountIdentification struct for PLLocalAccountIdentification
type PLLocalAccountIdentification struct {
	// The 26-digit bank account number ([Numer rachunku](https://pl.wikipedia.org/wiki/Numer_Rachunku_Bankowego)), without separators or whitespace.
	AccountNumber string `json:"accountNumber"`
	// **plLocal**
	Type string `json:"type"`
}

// NewPLLocalAccountIdentification instantiates a new PLLocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPLLocalAccountIdentification(accountNumber string, type_ string) *PLLocalAccountIdentification {
	this := PLLocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.Type = type_
	return &this
}

// NewPLLocalAccountIdentificationWithDefaults instantiates a new PLLocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPLLocalAccountIdentificationWithDefaults() *PLLocalAccountIdentification {
	this := PLLocalAccountIdentification{}
	var type_ string = "plLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *PLLocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *PLLocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *PLLocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetType returns the Type field value
func (o *PLLocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PLLocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PLLocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o PLLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PLLocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullablePLLocalAccountIdentification struct {
	value *PLLocalAccountIdentification
	isSet bool
}

func (v NullablePLLocalAccountIdentification) Get() *PLLocalAccountIdentification {
	return v.value
}

func (v *NullablePLLocalAccountIdentification) Set(val *PLLocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullablePLLocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullablePLLocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePLLocalAccountIdentification(val *PLLocalAccountIdentification) *NullablePLLocalAccountIdentification {
	return &NullablePLLocalAccountIdentification{value: val, isSet: true}
}

func (v NullablePLLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePLLocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PLLocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"plLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
