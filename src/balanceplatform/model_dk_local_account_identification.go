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

// checks if the DKLocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DKLocalAccountIdentification{}

// DKLocalAccountIdentification struct for DKLocalAccountIdentification
type DKLocalAccountIdentification struct {
	// The 4-10 digits bank account number (Kontonummer) (without separators or whitespace).
	AccountNumber string `json:"accountNumber"`
	// The 4-digit bank code (Registreringsnummer) (without separators or whitespace).
	BankCode string `json:"bankCode"`
	// **dkLocal**
	Type string `json:"type"`
}

// NewDKLocalAccountIdentification instantiates a new DKLocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDKLocalAccountIdentification(accountNumber string, bankCode string, type_ string) *DKLocalAccountIdentification {
	this := DKLocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.BankCode = bankCode
	this.Type = type_
	return &this
}

// NewDKLocalAccountIdentificationWithDefaults instantiates a new DKLocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDKLocalAccountIdentificationWithDefaults() *DKLocalAccountIdentification {
	this := DKLocalAccountIdentification{}
	var type_ string = "dkLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *DKLocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *DKLocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *DKLocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetBankCode returns the BankCode field value
func (o *DKLocalAccountIdentification) GetBankCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value
// and a boolean to check if the value has been set.
func (o *DKLocalAccountIdentification) GetBankCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankCode, true
}

// SetBankCode sets field value
func (o *DKLocalAccountIdentification) SetBankCode(v string) {
	o.BankCode = v
}

// GetType returns the Type field value
func (o *DKLocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DKLocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DKLocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o DKLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DKLocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["bankCode"] = o.BankCode
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableDKLocalAccountIdentification struct {
	value *DKLocalAccountIdentification
	isSet bool
}

func (v NullableDKLocalAccountIdentification) Get() *DKLocalAccountIdentification {
	return v.value
}

func (v *NullableDKLocalAccountIdentification) Set(val *DKLocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableDKLocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableDKLocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDKLocalAccountIdentification(val *DKLocalAccountIdentification) *NullableDKLocalAccountIdentification {
	return &NullableDKLocalAccountIdentification{value: val, isSet: true}
}

func (v NullableDKLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDKLocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *DKLocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"dkLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
