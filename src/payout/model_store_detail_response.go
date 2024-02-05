/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the StoreDetailResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoreDetailResponse{}

// StoreDetailResponse struct for StoreDetailResponse
type StoreDetailResponse struct {
	// This field contains additional data, which may be returned in a particular response.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// A new reference to uniquely identify this request.
	PspReference string `json:"pspReference"`
	// The token which you can use later on for submitting the payout.
	RecurringDetailReference string `json:"recurringDetailReference"`
	// The result code of the transaction. `Success` indicates that the details were stored successfully.
	ResultCode string `json:"resultCode"`
}

// NewStoreDetailResponse instantiates a new StoreDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreDetailResponse(pspReference string, recurringDetailReference string, resultCode string) *StoreDetailResponse {
	this := StoreDetailResponse{}
	this.PspReference = pspReference
	this.RecurringDetailReference = recurringDetailReference
	this.ResultCode = resultCode
	return &this
}

// NewStoreDetailResponseWithDefaults instantiates a new StoreDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreDetailResponseWithDefaults() *StoreDetailResponse {
	this := StoreDetailResponse{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *StoreDetailResponse) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailResponse) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *StoreDetailResponse) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *StoreDetailResponse) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetPspReference returns the PspReference field value
func (o *StoreDetailResponse) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *StoreDetailResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *StoreDetailResponse) SetPspReference(v string) {
	o.PspReference = v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value
func (o *StoreDetailResponse) GetRecurringDetailReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value
// and a boolean to check if the value has been set.
func (o *StoreDetailResponse) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecurringDetailReference, true
}

// SetRecurringDetailReference sets field value
func (o *StoreDetailResponse) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = v
}

// GetResultCode returns the ResultCode field value
func (o *StoreDetailResponse) GetResultCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value
// and a boolean to check if the value has been set.
func (o *StoreDetailResponse) GetResultCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultCode, true
}

// SetResultCode sets field value
func (o *StoreDetailResponse) SetResultCode(v string) {
	o.ResultCode = v
}

func (o StoreDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["pspReference"] = o.PspReference
	toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	toSerialize["resultCode"] = o.ResultCode
	return toSerialize, nil
}

type NullableStoreDetailResponse struct {
	value *StoreDetailResponse
	isSet bool
}

func (v NullableStoreDetailResponse) Get() *StoreDetailResponse {
	return v.value
}

func (v *NullableStoreDetailResponse) Set(val *StoreDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreDetailResponse(val *StoreDetailResponse) *NullableStoreDetailResponse {
	return &NullableStoreDetailResponse{value: val, isSet: true}
}

func (v NullableStoreDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
