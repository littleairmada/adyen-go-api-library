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

// checks if the PinChangeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PinChangeResponse{}

// PinChangeResponse struct for PinChangeResponse
type PinChangeResponse struct {
	// The pin change status.
	Status string `json:"status"`
}

// NewPinChangeResponse instantiates a new PinChangeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPinChangeResponse(status string) *PinChangeResponse {
	this := PinChangeResponse{}
	this.Status = status
	return &this
}

// NewPinChangeResponseWithDefaults instantiates a new PinChangeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinChangeResponseWithDefaults() *PinChangeResponse {
	this := PinChangeResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *PinChangeResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PinChangeResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PinChangeResponse) SetStatus(v string) {
	o.Status = v
}

func (o PinChangeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PinChangeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullablePinChangeResponse struct {
	value *PinChangeResponse
	isSet bool
}

func (v NullablePinChangeResponse) Get() *PinChangeResponse {
	return v.value
}

func (v *NullablePinChangeResponse) Set(val *PinChangeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePinChangeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePinChangeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePinChangeResponse(val *PinChangeResponse) *NullablePinChangeResponse {
	return &NullablePinChangeResponse{value: val, isSet: true}
}

func (v NullablePinChangeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePinChangeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PinChangeResponse) isValidStatus() bool {
	var allowedEnumValues = []string{"completed", "pending", "unavailable"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
