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

// checks if the AULocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AULocalAccountIdentification{}

// AULocalAccountIdentification struct for AULocalAccountIdentification
type AULocalAccountIdentification struct {
	// The bank account number, without separators or whitespace.
	AccountNumber string `json:"accountNumber"`
	// The 6-digit [Bank State Branch (BSB) code](https://en.wikipedia.org/wiki/Bank_state_branch), without separators or whitespace.
	BsbCode string `json:"bsbCode"`
	// **auLocal**
	Type string `json:"type"`
}

// NewAULocalAccountIdentification instantiates a new AULocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAULocalAccountIdentification(accountNumber string, bsbCode string, type_ string) *AULocalAccountIdentification {
	this := AULocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.BsbCode = bsbCode
	this.Type = type_
	return &this
}

// NewAULocalAccountIdentificationWithDefaults instantiates a new AULocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAULocalAccountIdentificationWithDefaults() *AULocalAccountIdentification {
	this := AULocalAccountIdentification{}
	var type_ string = "auLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *AULocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *AULocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *AULocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetBsbCode returns the BsbCode field value
func (o *AULocalAccountIdentification) GetBsbCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BsbCode
}

// GetBsbCodeOk returns a tuple with the BsbCode field value
// and a boolean to check if the value has been set.
func (o *AULocalAccountIdentification) GetBsbCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BsbCode, true
}

// SetBsbCode sets field value
func (o *AULocalAccountIdentification) SetBsbCode(v string) {
	o.BsbCode = v
}

// GetType returns the Type field value
func (o *AULocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AULocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AULocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o AULocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AULocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["bsbCode"] = o.BsbCode
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAULocalAccountIdentification struct {
	value *AULocalAccountIdentification
	isSet bool
}

func (v NullableAULocalAccountIdentification) Get() *AULocalAccountIdentification {
	return v.value
}

func (v *NullableAULocalAccountIdentification) Set(val *AULocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableAULocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableAULocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAULocalAccountIdentification(val *AULocalAccountIdentification) *NullableAULocalAccountIdentification {
	return &NullableAULocalAccountIdentification{value: val, isSet: true}
}

func (v NullableAULocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAULocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AULocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"auLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
