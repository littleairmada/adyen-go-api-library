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

// checks if the Repayment type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Repayment{}

// Repayment struct for Repayment
type Repayment struct {
	// The repayment that is deducted daily from incoming net volume, in [basis points](https://www.investopedia.com/terms/b/basispoint.asp).
	BasisPoints int32               `json:"basisPoints"`
	Term        *RepaymentTerm      `json:"term,omitempty"`
	Threshold   *ThresholdRepayment `json:"threshold,omitempty"`
}

// NewRepayment instantiates a new Repayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepayment(basisPoints int32) *Repayment {
	this := Repayment{}
	this.BasisPoints = basisPoints
	return &this
}

// NewRepaymentWithDefaults instantiates a new Repayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepaymentWithDefaults() *Repayment {
	this := Repayment{}
	return &this
}

// GetBasisPoints returns the BasisPoints field value
func (o *Repayment) GetBasisPoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BasisPoints
}

// GetBasisPointsOk returns a tuple with the BasisPoints field value
// and a boolean to check if the value has been set.
func (o *Repayment) GetBasisPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasisPoints, true
}

// SetBasisPoints sets field value
func (o *Repayment) SetBasisPoints(v int32) {
	o.BasisPoints = v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *Repayment) GetTerm() RepaymentTerm {
	if o == nil || common.IsNil(o.Term) {
		var ret RepaymentTerm
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repayment) GetTermOk() (*RepaymentTerm, bool) {
	if o == nil || common.IsNil(o.Term) {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *Repayment) HasTerm() bool {
	if o != nil && !common.IsNil(o.Term) {
		return true
	}

	return false
}

// SetTerm gets a reference to the given RepaymentTerm and assigns it to the Term field.
func (o *Repayment) SetTerm(v RepaymentTerm) {
	o.Term = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *Repayment) GetThreshold() ThresholdRepayment {
	if o == nil || common.IsNil(o.Threshold) {
		var ret ThresholdRepayment
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repayment) GetThresholdOk() (*ThresholdRepayment, bool) {
	if o == nil || common.IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *Repayment) HasThreshold() bool {
	if o != nil && !common.IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given ThresholdRepayment and assigns it to the Threshold field.
func (o *Repayment) SetThreshold(v ThresholdRepayment) {
	o.Threshold = &v
}

func (o Repayment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Repayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["basisPoints"] = o.BasisPoints
	if !common.IsNil(o.Term) {
		toSerialize["term"] = o.Term
	}
	if !common.IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	return toSerialize, nil
}

type NullableRepayment struct {
	value *Repayment
	isSet bool
}

func (v NullableRepayment) Get() *Repayment {
	return v.value
}

func (v *NullableRepayment) Set(val *Repayment) {
	v.value = val
	v.isSet = true
}

func (v NullableRepayment) IsSet() bool {
	return v.isSet
}

func (v *NullableRepayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepayment(val *Repayment) *NullableRepayment {
	return &NullableRepayment{value: val, isSet: true}
}

func (v NullableRepayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
