/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the DefendDisputeRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DefendDisputeRequest{}

// DefendDisputeRequest struct for DefendDisputeRequest
type DefendDisputeRequest struct {
	// The defense reason code that was selected to defend this dispute.
	DefenseReasonCode string `json:"defenseReasonCode"`
	// The PSP reference assigned to the dispute.
	DisputePspReference string `json:"disputePspReference"`
	// The merchant account identifier, for which you want to process the dispute transaction.
	MerchantAccountCode string `json:"merchantAccountCode"`
}

// NewDefendDisputeRequest instantiates a new DefendDisputeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefendDisputeRequest(defenseReasonCode string, disputePspReference string, merchantAccountCode string) *DefendDisputeRequest {
	this := DefendDisputeRequest{}
	this.DefenseReasonCode = defenseReasonCode
	this.DisputePspReference = disputePspReference
	this.MerchantAccountCode = merchantAccountCode
	return &this
}

// NewDefendDisputeRequestWithDefaults instantiates a new DefendDisputeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefendDisputeRequestWithDefaults() *DefendDisputeRequest {
	this := DefendDisputeRequest{}
	return &this
}

// GetDefenseReasonCode returns the DefenseReasonCode field value
func (o *DefendDisputeRequest) GetDefenseReasonCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefenseReasonCode
}

// GetDefenseReasonCodeOk returns a tuple with the DefenseReasonCode field value
// and a boolean to check if the value has been set.
func (o *DefendDisputeRequest) GetDefenseReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefenseReasonCode, true
}

// SetDefenseReasonCode sets field value
func (o *DefendDisputeRequest) SetDefenseReasonCode(v string) {
	o.DefenseReasonCode = v
}

// GetDisputePspReference returns the DisputePspReference field value
func (o *DefendDisputeRequest) GetDisputePspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisputePspReference
}

// GetDisputePspReferenceOk returns a tuple with the DisputePspReference field value
// and a boolean to check if the value has been set.
func (o *DefendDisputeRequest) GetDisputePspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisputePspReference, true
}

// SetDisputePspReference sets field value
func (o *DefendDisputeRequest) SetDisputePspReference(v string) {
	o.DisputePspReference = v
}

// GetMerchantAccountCode returns the MerchantAccountCode field value
func (o *DefendDisputeRequest) GetMerchantAccountCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccountCode
}

// GetMerchantAccountCodeOk returns a tuple with the MerchantAccountCode field value
// and a boolean to check if the value has been set.
func (o *DefendDisputeRequest) GetMerchantAccountCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccountCode, true
}

// SetMerchantAccountCode sets field value
func (o *DefendDisputeRequest) SetMerchantAccountCode(v string) {
	o.MerchantAccountCode = v
}

func (o DefendDisputeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefendDisputeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defenseReasonCode"] = o.DefenseReasonCode
	toSerialize["disputePspReference"] = o.DisputePspReference
	toSerialize["merchantAccountCode"] = o.MerchantAccountCode
	return toSerialize, nil
}

type NullableDefendDisputeRequest struct {
	value *DefendDisputeRequest
	isSet bool
}

func (v NullableDefendDisputeRequest) Get() *DefendDisputeRequest {
	return v.value
}

func (v *NullableDefendDisputeRequest) Set(val *DefendDisputeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDefendDisputeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDefendDisputeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefendDisputeRequest(val *DefendDisputeRequest) *NullableDefendDisputeRequest {
	return &NullableDefendDisputeRequest{value: val, isSet: true}
}

func (v NullableDefendDisputeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefendDisputeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
