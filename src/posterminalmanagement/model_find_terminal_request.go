/*
POS Terminal Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posterminalmanagement

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the FindTerminalRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FindTerminalRequest{}

// FindTerminalRequest struct for FindTerminalRequest
type FindTerminalRequest struct {
	// The unique terminal ID in the format `[Device model]-[Serial number]`.   For example, **V400m-324689776**.
	Terminal string `json:"terminal"`
}

// NewFindTerminalRequest instantiates a new FindTerminalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindTerminalRequest(terminal string) *FindTerminalRequest {
	this := FindTerminalRequest{}
	this.Terminal = terminal
	return &this
}

// NewFindTerminalRequestWithDefaults instantiates a new FindTerminalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindTerminalRequestWithDefaults() *FindTerminalRequest {
	this := FindTerminalRequest{}
	return &this
}

// GetTerminal returns the Terminal field value
func (o *FindTerminalRequest) GetTerminal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Terminal
}

// GetTerminalOk returns a tuple with the Terminal field value
// and a boolean to check if the value has been set.
func (o *FindTerminalRequest) GetTerminalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Terminal, true
}

// SetTerminal sets field value
func (o *FindTerminalRequest) SetTerminal(v string) {
	o.Terminal = v
}

func (o FindTerminalRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindTerminalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["terminal"] = o.Terminal
	return toSerialize, nil
}

type NullableFindTerminalRequest struct {
	value *FindTerminalRequest
	isSet bool
}

func (v NullableFindTerminalRequest) Get() *FindTerminalRequest {
	return v.value
}

func (v *NullableFindTerminalRequest) Set(val *FindTerminalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFindTerminalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFindTerminalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindTerminalRequest(val *FindTerminalRequest) *NullableFindTerminalRequest {
	return &NullableFindTerminalRequest{value: val, isSet: true}
}

func (v NullableFindTerminalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindTerminalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



