/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the BusinessLines type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BusinessLines{}

// BusinessLines struct for BusinessLines
type BusinessLines struct {
	// List of business lines.
	BusinessLines []BusinessLine `json:"businessLines"`
}

// NewBusinessLines instantiates a new BusinessLines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessLines(businessLines []BusinessLine) *BusinessLines {
	this := BusinessLines{}
	this.BusinessLines = businessLines
	return &this
}

// NewBusinessLinesWithDefaults instantiates a new BusinessLines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessLinesWithDefaults() *BusinessLines {
	this := BusinessLines{}
	return &this
}

// GetBusinessLines returns the BusinessLines field value
func (o *BusinessLines) GetBusinessLines() []BusinessLine {
	if o == nil {
		var ret []BusinessLine
		return ret
	}

	return o.BusinessLines
}

// GetBusinessLinesOk returns a tuple with the BusinessLines field value
// and a boolean to check if the value has been set.
func (o *BusinessLines) GetBusinessLinesOk() ([]BusinessLine, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessLines, true
}

// SetBusinessLines sets field value
func (o *BusinessLines) SetBusinessLines(v []BusinessLine) {
	o.BusinessLines = v
}

func (o BusinessLines) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessLines) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["businessLines"] = o.BusinessLines
	return toSerialize, nil
}

type NullableBusinessLines struct {
	value *BusinessLines
	isSet bool
}

func (v NullableBusinessLines) Get() *BusinessLines {
	return v.value
}

func (v *NullableBusinessLines) Set(val *BusinessLines) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessLines) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessLines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessLines(val *BusinessLines) *NullableBusinessLines {
	return &NullableBusinessLines{value: val, isSet: true}
}

func (v NullableBusinessLines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessLines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
