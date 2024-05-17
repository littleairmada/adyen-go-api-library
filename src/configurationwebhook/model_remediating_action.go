/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the RemediatingAction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RemediatingAction{}

// RemediatingAction struct for RemediatingAction
type RemediatingAction struct {
	// The remediating action code.
	Code *string `json:"code,omitempty"`
	// A description of how you can resolve the verification error.
	Message *string `json:"message,omitempty"`
}

// NewRemediatingAction instantiates a new RemediatingAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemediatingAction() *RemediatingAction {
	this := RemediatingAction{}
	return &this
}

// NewRemediatingActionWithDefaults instantiates a new RemediatingAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemediatingActionWithDefaults() *RemediatingAction {
	this := RemediatingAction{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RemediatingAction) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediatingAction) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RemediatingAction) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *RemediatingAction) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RemediatingAction) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediatingAction) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RemediatingAction) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RemediatingAction) SetMessage(v string) {
	o.Message = &v
}

func (o RemediatingAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemediatingAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableRemediatingAction struct {
	value *RemediatingAction
	isSet bool
}

func (v NullableRemediatingAction) Get() *RemediatingAction {
	return v.value
}

func (v *NullableRemediatingAction) Set(val *RemediatingAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRemediatingAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRemediatingAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemediatingAction(val *RemediatingAction) *NullableRemediatingAction {
	return &NullableRemediatingAction{value: val, isSet: true}
}

func (v NullableRemediatingAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemediatingAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
