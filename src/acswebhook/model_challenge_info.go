/*
Authentication webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package acswebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the ChallengeInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ChallengeInfo{}

// ChallengeInfo struct for ChallengeInfo
type ChallengeInfo struct {
	// Indicator informing the Access Control Server (ACS) and the Directory Server (DS) that the authentication has been cancelled. Possible values: * **00**: Data element is absent or value has been sent back with the key `challengeCancel`. * **01**: Cardholder selected **Cancel**. * **02**: 3DS Requestor cancelled Authentication. * **03**: Transaction abandoned. * **04**: Transaction timed out at ACS — other timeouts. * **05**: Transaction timed out at ACS — first CReq not received by ACS. * **06**: Transaction error. * **07**: Unknown. * **08**: Transaction time out at SDK.
	ChallengeCancel *string `json:"challengeCancel,omitempty"`
	// The flow used in the challenge. Possible values:  * **OTP_SMS**: one-time password (OTP) flow * **OOB**: out-of-band (OOB) flow
	Flow string `json:"flow"`
	// The last time of interaction with the challenge.
	LastInteraction time.Time `json:"lastInteraction"`
	// The last four digits of the phone number used in the challenge.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// The number of times the one-time password (OTP) was resent during the challenge.
	Resends *int32 `json:"resends,omitempty"`
	// The number of retries used in the challenge.
	Retries *int32 `json:"retries,omitempty"`
}

// NewChallengeInfo instantiates a new ChallengeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChallengeInfo(flow string, lastInteraction time.Time) *ChallengeInfo {
	this := ChallengeInfo{}
	this.Flow = flow
	this.LastInteraction = lastInteraction
	return &this
}

// NewChallengeInfoWithDefaults instantiates a new ChallengeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChallengeInfoWithDefaults() *ChallengeInfo {
	this := ChallengeInfo{}
	return &this
}

// GetChallengeCancel returns the ChallengeCancel field value if set, zero value otherwise.
func (o *ChallengeInfo) GetChallengeCancel() string {
	if o == nil || common.IsNil(o.ChallengeCancel) {
		var ret string
		return ret
	}
	return *o.ChallengeCancel
}

// GetChallengeCancelOk returns a tuple with the ChallengeCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeInfo) GetChallengeCancelOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChallengeCancel) {
		return nil, false
	}
	return o.ChallengeCancel, true
}

// HasChallengeCancel returns a boolean if a field has been set.
func (o *ChallengeInfo) HasChallengeCancel() bool {
	if o != nil && !common.IsNil(o.ChallengeCancel) {
		return true
	}

	return false
}

// SetChallengeCancel gets a reference to the given string and assigns it to the ChallengeCancel field.
func (o *ChallengeInfo) SetChallengeCancel(v string) {
	o.ChallengeCancel = &v
}

// GetFlow returns the Flow field value
func (o *ChallengeInfo) GetFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Flow
}

// GetFlowOk returns a tuple with the Flow field value
// and a boolean to check if the value has been set.
func (o *ChallengeInfo) GetFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flow, true
}

// SetFlow sets field value
func (o *ChallengeInfo) SetFlow(v string) {
	o.Flow = v
}

// GetLastInteraction returns the LastInteraction field value
func (o *ChallengeInfo) GetLastInteraction() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastInteraction
}

// GetLastInteractionOk returns a tuple with the LastInteraction field value
// and a boolean to check if the value has been set.
func (o *ChallengeInfo) GetLastInteractionOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastInteraction, true
}

// SetLastInteraction sets field value
func (o *ChallengeInfo) SetLastInteraction(v time.Time) {
	o.LastInteraction = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ChallengeInfo) GetPhoneNumber() string {
	if o == nil || common.IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ChallengeInfo) HasPhoneNumber() bool {
	if o != nil && !common.IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *ChallengeInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetResends returns the Resends field value if set, zero value otherwise.
func (o *ChallengeInfo) GetResends() int32 {
	if o == nil || common.IsNil(o.Resends) {
		var ret int32
		return ret
	}
	return *o.Resends
}

// GetResendsOk returns a tuple with the Resends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeInfo) GetResendsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Resends) {
		return nil, false
	}
	return o.Resends, true
}

// HasResends returns a boolean if a field has been set.
func (o *ChallengeInfo) HasResends() bool {
	if o != nil && !common.IsNil(o.Resends) {
		return true
	}

	return false
}

// SetResends gets a reference to the given int32 and assigns it to the Resends field.
func (o *ChallengeInfo) SetResends(v int32) {
	o.Resends = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *ChallengeInfo) GetRetries() int32 {
	if o == nil || common.IsNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeInfo) GetRetriesOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Retries) {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *ChallengeInfo) HasRetries() bool {
	if o != nil && !common.IsNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *ChallengeInfo) SetRetries(v int32) {
	o.Retries = &v
}

func (o ChallengeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChallengeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ChallengeCancel) {
		toSerialize["challengeCancel"] = o.ChallengeCancel
	}
	toSerialize["flow"] = o.Flow
	toSerialize["lastInteraction"] = o.LastInteraction
	if !common.IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !common.IsNil(o.Resends) {
		toSerialize["resends"] = o.Resends
	}
	if !common.IsNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	return toSerialize, nil
}

type NullableChallengeInfo struct {
	value *ChallengeInfo
	isSet bool
}

func (v NullableChallengeInfo) Get() *ChallengeInfo {
	return v.value
}

func (v *NullableChallengeInfo) Set(val *ChallengeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableChallengeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableChallengeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChallengeInfo(val *ChallengeInfo) *NullableChallengeInfo {
	return &NullableChallengeInfo{value: val, isSet: true}
}

func (v NullableChallengeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChallengeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *ChallengeInfo) isValidChallengeCancel() bool {
	var allowedEnumValues = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08"}
	for _, allowed := range allowedEnumValues {
		if o.GetChallengeCancel() == allowed {
			return true
		}
	}
	return false
}
func (o *ChallengeInfo) isValidFlow() bool {
	var allowedEnumValues = []string{"OTP_SMS", "OOB"}
	for _, allowed := range allowedEnumValues {
		if o.GetFlow() == allowed {
			return true
		}
	}
	return false
}
