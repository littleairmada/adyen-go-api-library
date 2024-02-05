/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the IdName type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IdName{}

// IdName struct for IdName
type IdName struct {
	// The identifier of the terminal model.
	Id *string `json:"id,omitempty"`
	// The name of the terminal model.
	Name *string `json:"name,omitempty"`
}

// NewIdName instantiates a new IdName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdName() *IdName {
	this := IdName{}
	return &this
}

// NewIdNameWithDefaults instantiates a new IdName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdNameWithDefaults() *IdName {
	this := IdName{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdName) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdName) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdName) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdName) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdName) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdName) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdName) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdName) SetName(v string) {
	o.Name = &v
}

func (o IdName) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableIdName struct {
	value *IdName
	isSet bool
}

func (v NullableIdName) Get() *IdName {
	return v.value
}

func (v *NullableIdName) Set(val *IdName) {
	v.value = val
	v.isSet = true
}

func (v NullableIdName) IsSet() bool {
	return v.isSet
}

func (v *NullableIdName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdName(val *IdName) *NullableIdName {
	return &NullableIdName{value: val, isSet: true}
}

func (v NullableIdName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
