/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the RiskData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RiskData{}

// RiskData struct for RiskData
type RiskData struct {
	// Contains client-side data, like the device fingerprint, cookies, and specific browser settings.
	ClientData *string `json:"clientData,omitempty"`
	// Any custom fields used as part of the input to configured risk rules.
	CustomFields *map[string]string `json:"customFields,omitempty"`
	// An integer value that is added to the normal fraud score. The value can be either positive or negative.
	FraudOffset *int32 `json:"fraudOffset,omitempty"`
	// The risk profile to assign to this payment. When left empty, the merchant-level account's default risk profile will be applied.
	ProfileReference *string `json:"profileReference,omitempty"`
}

// NewRiskData instantiates a new RiskData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskData() *RiskData {
	this := RiskData{}
	return &this
}

// NewRiskDataWithDefaults instantiates a new RiskData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskDataWithDefaults() *RiskData {
	this := RiskData{}
	return &this
}

// GetClientData returns the ClientData field value if set, zero value otherwise.
func (o *RiskData) GetClientData() string {
	if o == nil || common.IsNil(o.ClientData) {
		var ret string
		return ret
	}
	return *o.ClientData
}

// GetClientDataOk returns a tuple with the ClientData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskData) GetClientDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientData) {
		return nil, false
	}
	return o.ClientData, true
}

// HasClientData returns a boolean if a field has been set.
func (o *RiskData) HasClientData() bool {
	if o != nil && !common.IsNil(o.ClientData) {
		return true
	}

	return false
}

// SetClientData gets a reference to the given string and assigns it to the ClientData field.
func (o *RiskData) SetClientData(v string) {
	o.ClientData = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *RiskData) GetCustomFields() map[string]string {
	if o == nil || common.IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskData) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *RiskData) HasCustomFields() bool {
	if o != nil && !common.IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *RiskData) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

// GetFraudOffset returns the FraudOffset field value if set, zero value otherwise.
func (o *RiskData) GetFraudOffset() int32 {
	if o == nil || common.IsNil(o.FraudOffset) {
		var ret int32
		return ret
	}
	return *o.FraudOffset
}

// GetFraudOffsetOk returns a tuple with the FraudOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskData) GetFraudOffsetOk() (*int32, bool) {
	if o == nil || common.IsNil(o.FraudOffset) {
		return nil, false
	}
	return o.FraudOffset, true
}

// HasFraudOffset returns a boolean if a field has been set.
func (o *RiskData) HasFraudOffset() bool {
	if o != nil && !common.IsNil(o.FraudOffset) {
		return true
	}

	return false
}

// SetFraudOffset gets a reference to the given int32 and assigns it to the FraudOffset field.
func (o *RiskData) SetFraudOffset(v int32) {
	o.FraudOffset = &v
}

// GetProfileReference returns the ProfileReference field value if set, zero value otherwise.
func (o *RiskData) GetProfileReference() string {
	if o == nil || common.IsNil(o.ProfileReference) {
		var ret string
		return ret
	}
	return *o.ProfileReference
}

// GetProfileReferenceOk returns a tuple with the ProfileReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskData) GetProfileReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProfileReference) {
		return nil, false
	}
	return o.ProfileReference, true
}

// HasProfileReference returns a boolean if a field has been set.
func (o *RiskData) HasProfileReference() bool {
	if o != nil && !common.IsNil(o.ProfileReference) {
		return true
	}

	return false
}

// SetProfileReference gets a reference to the given string and assigns it to the ProfileReference field.
func (o *RiskData) SetProfileReference(v string) {
	o.ProfileReference = &v
}

func (o RiskData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ClientData) {
		toSerialize["clientData"] = o.ClientData
	}
	if !common.IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !common.IsNil(o.FraudOffset) {
		toSerialize["fraudOffset"] = o.FraudOffset
	}
	if !common.IsNil(o.ProfileReference) {
		toSerialize["profileReference"] = o.ProfileReference
	}
	return toSerialize, nil
}

type NullableRiskData struct {
	value *RiskData
	isSet bool
}

func (v NullableRiskData) Get() *RiskData {
	return v.value
}

func (v *NullableRiskData) Set(val *RiskData) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskData) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskData(val *RiskData) *NullableRiskData {
	return &NullableRiskData{value: val, isSet: true}
}

func (v NullableRiskData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
