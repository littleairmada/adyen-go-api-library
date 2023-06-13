/*
Transfer webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the TransactionRulesResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionRulesResult{}

// TransactionRulesResult struct for TransactionRulesResult
type TransactionRulesResult struct {
	// The advice given by the Risk analysis.
	Advice *string `json:"advice,omitempty"`
	// Indicates whether the transaction passed the evaluation for all transaction rules.
	AllRulesPassed *bool `json:"allRulesPassed,omitempty"`
	// Array containing all the transaction rules that the transaction violated. This list is only sent when `allRulesPassed` is **false**.
	FailedTransactionRules []TransactionEventViolation `json:"failedTransactionRules,omitempty"`
	// The score of the Risk analysis.
	Score *int32 `json:"score,omitempty"`
}

// NewTransactionRulesResult instantiates a new TransactionRulesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRulesResult() *TransactionRulesResult {
	this := TransactionRulesResult{}
	return &this
}

// NewTransactionRulesResultWithDefaults instantiates a new TransactionRulesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRulesResultWithDefaults() *TransactionRulesResult {
	this := TransactionRulesResult{}
	return &this
}

// GetAdvice returns the Advice field value if set, zero value otherwise.
func (o *TransactionRulesResult) GetAdvice() string {
	if o == nil || common.IsNil(o.Advice) {
		var ret string
		return ret
	}
	return *o.Advice
}

// GetAdviceOk returns a tuple with the Advice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRulesResult) GetAdviceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Advice) {
		return nil, false
	}
	return o.Advice, true
}

// HasAdvice returns a boolean if a field has been set.
func (o *TransactionRulesResult) HasAdvice() bool {
	if o != nil && !common.IsNil(o.Advice) {
		return true
	}

	return false
}

// SetAdvice gets a reference to the given string and assigns it to the Advice field.
func (o *TransactionRulesResult) SetAdvice(v string) {
	o.Advice = &v
}

// GetAllRulesPassed returns the AllRulesPassed field value if set, zero value otherwise.
func (o *TransactionRulesResult) GetAllRulesPassed() bool {
	if o == nil || common.IsNil(o.AllRulesPassed) {
		var ret bool
		return ret
	}
	return *o.AllRulesPassed
}

// GetAllRulesPassedOk returns a tuple with the AllRulesPassed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRulesResult) GetAllRulesPassedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AllRulesPassed) {
		return nil, false
	}
	return o.AllRulesPassed, true
}

// HasAllRulesPassed returns a boolean if a field has been set.
func (o *TransactionRulesResult) HasAllRulesPassed() bool {
	if o != nil && !common.IsNil(o.AllRulesPassed) {
		return true
	}

	return false
}

// SetAllRulesPassed gets a reference to the given bool and assigns it to the AllRulesPassed field.
func (o *TransactionRulesResult) SetAllRulesPassed(v bool) {
	o.AllRulesPassed = &v
}

// GetFailedTransactionRules returns the FailedTransactionRules field value if set, zero value otherwise.
func (o *TransactionRulesResult) GetFailedTransactionRules() []TransactionEventViolation {
	if o == nil || common.IsNil(o.FailedTransactionRules) {
		var ret []TransactionEventViolation
		return ret
	}
	return o.FailedTransactionRules
}

// GetFailedTransactionRulesOk returns a tuple with the FailedTransactionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRulesResult) GetFailedTransactionRulesOk() ([]TransactionEventViolation, bool) {
	if o == nil || common.IsNil(o.FailedTransactionRules) {
		return nil, false
	}
	return o.FailedTransactionRules, true
}

// HasFailedTransactionRules returns a boolean if a field has been set.
func (o *TransactionRulesResult) HasFailedTransactionRules() bool {
	if o != nil && !common.IsNil(o.FailedTransactionRules) {
		return true
	}

	return false
}

// SetFailedTransactionRules gets a reference to the given []TransactionEventViolation and assigns it to the FailedTransactionRules field.
func (o *TransactionRulesResult) SetFailedTransactionRules(v []TransactionEventViolation) {
	o.FailedTransactionRules = v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *TransactionRulesResult) GetScore() int32 {
	if o == nil || common.IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRulesResult) GetScoreOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *TransactionRulesResult) HasScore() bool {
	if o != nil && !common.IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *TransactionRulesResult) SetScore(v int32) {
	o.Score = &v
}

func (o TransactionRulesResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRulesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Advice) {
		toSerialize["advice"] = o.Advice
	}
	if !common.IsNil(o.AllRulesPassed) {
		toSerialize["allRulesPassed"] = o.AllRulesPassed
	}
	if !common.IsNil(o.FailedTransactionRules) {
		toSerialize["failedTransactionRules"] = o.FailedTransactionRules
	}
	if !common.IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	return toSerialize, nil
}

type NullableTransactionRulesResult struct {
	value *TransactionRulesResult
	isSet bool
}

func (v NullableTransactionRulesResult) Get() *TransactionRulesResult {
	return v.value
}

func (v *NullableTransactionRulesResult) Set(val *TransactionRulesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRulesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRulesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRulesResult(val *TransactionRulesResult) *NullableTransactionRulesResult {
	return &NullableTransactionRulesResult{value: val, isSet: true}
}

func (v NullableTransactionRulesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRulesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
