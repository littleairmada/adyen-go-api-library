/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the CZLocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CZLocalAccountIdentification{}

// CZLocalAccountIdentification struct for CZLocalAccountIdentification
type CZLocalAccountIdentification struct {
	// The 2- to 16-digit bank account number (Číslo účtu) in the following format:  - The optional prefix (předčíslí).  - The required second part (základní část) which must be at least two non-zero digits.  Examples:  - **19-123457** (with prefix)  - **123457** (without prefix)  - **000019-0000123457** (with prefix, normalized)  - **000000-0000123457** (without prefix, normalized)
	AccountNumber string `json:"accountNumber"`
	// The 4-digit bank code (Kód banky), without separators or whitespace.
	BankCode string `json:"bankCode"`
	// **czLocal**
	Type string `json:"type"`
}

// NewCZLocalAccountIdentification instantiates a new CZLocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCZLocalAccountIdentification(accountNumber string, bankCode string, type_ string) *CZLocalAccountIdentification {
	this := CZLocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.BankCode = bankCode
	this.Type = type_
	return &this
}

// NewCZLocalAccountIdentificationWithDefaults instantiates a new CZLocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCZLocalAccountIdentificationWithDefaults() *CZLocalAccountIdentification {
	this := CZLocalAccountIdentification{}
	var type_ string = "czLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *CZLocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *CZLocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *CZLocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetBankCode returns the BankCode field value
func (o *CZLocalAccountIdentification) GetBankCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value
// and a boolean to check if the value has been set.
func (o *CZLocalAccountIdentification) GetBankCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankCode, true
}

// SetBankCode sets field value
func (o *CZLocalAccountIdentification) SetBankCode(v string) {
	o.BankCode = v
}

// GetType returns the Type field value
func (o *CZLocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CZLocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CZLocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o CZLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CZLocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["bankCode"] = o.BankCode
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCZLocalAccountIdentification struct {
	value *CZLocalAccountIdentification
	isSet bool
}

func (v NullableCZLocalAccountIdentification) Get() *CZLocalAccountIdentification {
	return v.value
}

func (v *NullableCZLocalAccountIdentification) Set(val *CZLocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableCZLocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableCZLocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCZLocalAccountIdentification(val *CZLocalAccountIdentification) *NullableCZLocalAccountIdentification {
	return &NullableCZLocalAccountIdentification{value: val, isSet: true}
}

func (v NullableCZLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCZLocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CZLocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"czLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
