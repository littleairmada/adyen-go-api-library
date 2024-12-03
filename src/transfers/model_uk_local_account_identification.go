/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the UKLocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UKLocalAccountIdentification{}

// UKLocalAccountIdentification struct for UKLocalAccountIdentification
type UKLocalAccountIdentification struct {
	// The 8-digit bank account number, without separators or whitespace.
	AccountNumber string `json:"accountNumber"`
	// The 6-digit [sort code](https://en.wikipedia.org/wiki/Sort_code), without separators or whitespace.
	SortCode string `json:"sortCode"`
	// **ukLocal**
	Type string `json:"type"`
}

// NewUKLocalAccountIdentification instantiates a new UKLocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUKLocalAccountIdentification(accountNumber string, sortCode string, type_ string) *UKLocalAccountIdentification {
	this := UKLocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.SortCode = sortCode
	this.Type = type_
	return &this
}

// NewUKLocalAccountIdentificationWithDefaults instantiates a new UKLocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUKLocalAccountIdentificationWithDefaults() *UKLocalAccountIdentification {
	this := UKLocalAccountIdentification{}
	var type_ string = "ukLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *UKLocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *UKLocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *UKLocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetSortCode returns the SortCode field value
func (o *UKLocalAccountIdentification) GetSortCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SortCode
}

// GetSortCodeOk returns a tuple with the SortCode field value
// and a boolean to check if the value has been set.
func (o *UKLocalAccountIdentification) GetSortCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortCode, true
}

// SetSortCode sets field value
func (o *UKLocalAccountIdentification) SetSortCode(v string) {
	o.SortCode = v
}

// GetType returns the Type field value
func (o *UKLocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UKLocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UKLocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o UKLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UKLocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["sortCode"] = o.SortCode
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableUKLocalAccountIdentification struct {
	value *UKLocalAccountIdentification
	isSet bool
}

func (v NullableUKLocalAccountIdentification) Get() *UKLocalAccountIdentification {
	return v.value
}

func (v *NullableUKLocalAccountIdentification) Set(val *UKLocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableUKLocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableUKLocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUKLocalAccountIdentification(val *UKLocalAccountIdentification) *NullableUKLocalAccountIdentification {
	return &NullableUKLocalAccountIdentification{value: val, isSet: true}
}

func (v NullableUKLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUKLocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UKLocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"ukLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
