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

// checks if the BalancePlatform type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalancePlatform{}

// BalancePlatform struct for BalancePlatform
type BalancePlatform struct {
	// Your description of the balance platform, maximum 300 characters.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the balance platform.
	Id string `json:"id"`
	// The status of the balance platform.  Possible values: **Active**, **Inactive**, **Closed**, **Suspended**.
	Status *string `json:"status,omitempty"`
}

// NewBalancePlatform instantiates a new BalancePlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalancePlatform(id string) *BalancePlatform {
	this := BalancePlatform{}
	this.Id = id
	return &this
}

// NewBalancePlatformWithDefaults instantiates a new BalancePlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalancePlatformWithDefaults() *BalancePlatform {
	this := BalancePlatform{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BalancePlatform) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePlatform) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BalancePlatform) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BalancePlatform) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *BalancePlatform) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BalancePlatform) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BalancePlatform) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BalancePlatform) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePlatform) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BalancePlatform) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BalancePlatform) SetStatus(v string) {
	o.Status = &v
}

func (o BalancePlatform) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalancePlatform) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableBalancePlatform struct {
	value *BalancePlatform
	isSet bool
}

func (v NullableBalancePlatform) Get() *BalancePlatform {
	return v.value
}

func (v *NullableBalancePlatform) Set(val *BalancePlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableBalancePlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableBalancePlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalancePlatform(val *BalancePlatform) *NullableBalancePlatform {
	return &NullableBalancePlatform{value: val, isSet: true}
}

func (v NullableBalancePlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalancePlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
