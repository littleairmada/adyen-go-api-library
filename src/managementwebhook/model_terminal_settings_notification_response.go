/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the TerminalSettingsNotificationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalSettingsNotificationResponse{}

// TerminalSettingsNotificationResponse struct for TerminalSettingsNotificationResponse
type TerminalSettingsNotificationResponse struct {
	// Respond with any **2xx** HTTP status code to [accept the webhook](https://docs.adyen.com/development-resources/webhooks#accept-notifications).
	NotificationResponse *string `json:"notificationResponse,omitempty"`
}

// NewTerminalSettingsNotificationResponse instantiates a new TerminalSettingsNotificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalSettingsNotificationResponse() *TerminalSettingsNotificationResponse {
	this := TerminalSettingsNotificationResponse{}
	return &this
}

// NewTerminalSettingsNotificationResponseWithDefaults instantiates a new TerminalSettingsNotificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalSettingsNotificationResponseWithDefaults() *TerminalSettingsNotificationResponse {
	this := TerminalSettingsNotificationResponse{}
	return &this
}

// GetNotificationResponse returns the NotificationResponse field value if set, zero value otherwise.
func (o *TerminalSettingsNotificationResponse) GetNotificationResponse() string {
	if o == nil || common.IsNil(o.NotificationResponse) {
		var ret string
		return ret
	}
	return *o.NotificationResponse
}

// GetNotificationResponseOk returns a tuple with the NotificationResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettingsNotificationResponse) GetNotificationResponseOk() (*string, bool) {
	if o == nil || common.IsNil(o.NotificationResponse) {
		return nil, false
	}
	return o.NotificationResponse, true
}

// HasNotificationResponse returns a boolean if a field has been set.
func (o *TerminalSettingsNotificationResponse) HasNotificationResponse() bool {
	if o != nil && !common.IsNil(o.NotificationResponse) {
		return true
	}

	return false
}

// SetNotificationResponse gets a reference to the given string and assigns it to the NotificationResponse field.
func (o *TerminalSettingsNotificationResponse) SetNotificationResponse(v string) {
	o.NotificationResponse = &v
}

func (o TerminalSettingsNotificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalSettingsNotificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NotificationResponse) {
		toSerialize["notificationResponse"] = o.NotificationResponse
	}
	return toSerialize, nil
}

type NullableTerminalSettingsNotificationResponse struct {
	value *TerminalSettingsNotificationResponse
	isSet bool
}

func (v NullableTerminalSettingsNotificationResponse) Get() *TerminalSettingsNotificationResponse {
	return v.value
}

func (v *NullableTerminalSettingsNotificationResponse) Set(val *TerminalSettingsNotificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalSettingsNotificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalSettingsNotificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalSettingsNotificationResponse(val *TerminalSettingsNotificationResponse) *NullableTerminalSettingsNotificationResponse {
	return &NullableTerminalSettingsNotificationResponse{value: val, isSet: true}
}

func (v NullableTerminalSettingsNotificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalSettingsNotificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
