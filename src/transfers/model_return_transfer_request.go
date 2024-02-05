/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the ReturnTransferRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReturnTransferRequest{}

// ReturnTransferRequest struct for ReturnTransferRequest
type ReturnTransferRequest struct {
	Amount Amount `json:"amount"`
	// Your internal reference for the return. If you don't provide this in the request, Adyen generates a unique reference. This reference is used in all communication with you about the instruction status.  We recommend using a unique value per instruction. If you need to provide multiple references for a transaction, separate them with hyphens (\"-\").
	Reference *string `json:"reference,omitempty"`
}

// NewReturnTransferRequest instantiates a new ReturnTransferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnTransferRequest(amount Amount) *ReturnTransferRequest {
	this := ReturnTransferRequest{}
	this.Amount = amount
	return &this
}

// NewReturnTransferRequestWithDefaults instantiates a new ReturnTransferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnTransferRequestWithDefaults() *ReturnTransferRequest {
	this := ReturnTransferRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ReturnTransferRequest) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ReturnTransferRequest) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ReturnTransferRequest) SetAmount(v Amount) {
	o.Amount = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ReturnTransferRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnTransferRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ReturnTransferRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ReturnTransferRequest) SetReference(v string) {
	o.Reference = &v
}

func (o ReturnTransferRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnTransferRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	return toSerialize, nil
}

type NullableReturnTransferRequest struct {
	value *ReturnTransferRequest
	isSet bool
}

func (v NullableReturnTransferRequest) Get() *ReturnTransferRequest {
	return v.value
}

func (v *NullableReturnTransferRequest) Set(val *ReturnTransferRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnTransferRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnTransferRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnTransferRequest(val *ReturnTransferRequest) *NullableReturnTransferRequest {
	return &NullableReturnTransferRequest{value: val, isSet: true}
}

func (v NullableReturnTransferRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnTransferRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
