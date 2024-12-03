/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the ThreeDS2ResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThreeDS2ResponseData{}

// ThreeDS2ResponseData struct for ThreeDS2ResponseData
type ThreeDS2ResponseData struct {
	AcsChallengeMandated *string `json:"acsChallengeMandated,omitempty"`
	AcsOperatorID        *string `json:"acsOperatorID,omitempty"`
	AcsReferenceNumber   *string `json:"acsReferenceNumber,omitempty"`
	AcsSignedContent     *string `json:"acsSignedContent,omitempty"`
	AcsTransID           *string `json:"acsTransID,omitempty"`
	AcsURL               *string `json:"acsURL,omitempty"`
	AuthenticationType   *string `json:"authenticationType,omitempty"`
	CardHolderInfo       *string `json:"cardHolderInfo,omitempty"`
	CavvAlgorithm        *string `json:"cavvAlgorithm,omitempty"`
	ChallengeIndicator   *string `json:"challengeIndicator,omitempty"`
	DsReferenceNumber    *string `json:"dsReferenceNumber,omitempty"`
	DsTransID            *string `json:"dsTransID,omitempty"`
	ExemptionIndicator   *string `json:"exemptionIndicator,omitempty"`
	MessageVersion       *string `json:"messageVersion,omitempty"`
	RiskScore            *string `json:"riskScore,omitempty"`
	SdkEphemPubKey       *string `json:"sdkEphemPubKey,omitempty"`
	ThreeDSServerTransID *string `json:"threeDSServerTransID,omitempty"`
	TransStatus          *string `json:"transStatus,omitempty"`
	TransStatusReason    *string `json:"transStatusReason,omitempty"`
}

// NewThreeDS2ResponseData instantiates a new ThreeDS2ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDS2ResponseData() *ThreeDS2ResponseData {
	this := ThreeDS2ResponseData{}
	return &this
}

// NewThreeDS2ResponseDataWithDefaults instantiates a new ThreeDS2ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDS2ResponseDataWithDefaults() *ThreeDS2ResponseData {
	this := ThreeDS2ResponseData{}
	return &this
}

// GetAcsChallengeMandated returns the AcsChallengeMandated field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetAcsChallengeMandated() string {
	if o == nil || common.IsNil(o.AcsChallengeMandated) {
		var ret string
		return ret
	}
	return *o.AcsChallengeMandated
}

// GetAcsChallengeMandatedOk returns a tuple with the AcsChallengeMandated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetAcsChallengeMandatedOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcsChallengeMandated) {
		return nil, false
	}
	return o.AcsChallengeMandated, true
}

// HasAcsChallengeMandated returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasAcsChallengeMandated() bool {
	if o != nil && !common.IsNil(o.AcsChallengeMandated) {
		return true
	}

	return false
}

// SetAcsChallengeMandated gets a reference to the given string and assigns it to the AcsChallengeMandated field.
func (o *ThreeDS2ResponseData) SetAcsChallengeMandated(v string) {
	o.AcsChallengeMandated = &v
}

// GetAcsOperatorID returns the AcsOperatorID field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetAcsOperatorID() string {
	if o == nil || common.IsNil(o.AcsOperatorID) {
		var ret string
		return ret
	}
	return *o.AcsOperatorID
}

// GetAcsOperatorIDOk returns a tuple with the AcsOperatorID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetAcsOperatorIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcsOperatorID) {
		return nil, false
	}
	return o.AcsOperatorID, true
}

// HasAcsOperatorID returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasAcsOperatorID() bool {
	if o != nil && !common.IsNil(o.AcsOperatorID) {
		return true
	}

	return false
}

// SetAcsOperatorID gets a reference to the given string and assigns it to the AcsOperatorID field.
func (o *ThreeDS2ResponseData) SetAcsOperatorID(v string) {
	o.AcsOperatorID = &v
}

// GetAcsReferenceNumber returns the AcsReferenceNumber field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetAcsReferenceNumber() string {
	if o == nil || common.IsNil(o.AcsReferenceNumber) {
		var ret string
		return ret
	}
	return *o.AcsReferenceNumber
}

// GetAcsReferenceNumberOk returns a tuple with the AcsReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetAcsReferenceNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcsReferenceNumber) {
		return nil, false
	}
	return o.AcsReferenceNumber, true
}

// HasAcsReferenceNumber returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasAcsReferenceNumber() bool {
	if o != nil && !common.IsNil(o.AcsReferenceNumber) {
		return true
	}

	return false
}

// SetAcsReferenceNumber gets a reference to the given string and assigns it to the AcsReferenceNumber field.
func (o *ThreeDS2ResponseData) SetAcsReferenceNumber(v string) {
	o.AcsReferenceNumber = &v
}

// GetAcsSignedContent returns the AcsSignedContent field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetAcsSignedContent() string {
	if o == nil || common.IsNil(o.AcsSignedContent) {
		var ret string
		return ret
	}
	return *o.AcsSignedContent
}

// GetAcsSignedContentOk returns a tuple with the AcsSignedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetAcsSignedContentOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcsSignedContent) {
		return nil, false
	}
	return o.AcsSignedContent, true
}

// HasAcsSignedContent returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasAcsSignedContent() bool {
	if o != nil && !common.IsNil(o.AcsSignedContent) {
		return true
	}

	return false
}

// SetAcsSignedContent gets a reference to the given string and assigns it to the AcsSignedContent field.
func (o *ThreeDS2ResponseData) SetAcsSignedContent(v string) {
	o.AcsSignedContent = &v
}

// GetAcsTransID returns the AcsTransID field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetAcsTransID() string {
	if o == nil || common.IsNil(o.AcsTransID) {
		var ret string
		return ret
	}
	return *o.AcsTransID
}

// GetAcsTransIDOk returns a tuple with the AcsTransID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetAcsTransIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcsTransID) {
		return nil, false
	}
	return o.AcsTransID, true
}

// HasAcsTransID returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasAcsTransID() bool {
	if o != nil && !common.IsNil(o.AcsTransID) {
		return true
	}

	return false
}

// SetAcsTransID gets a reference to the given string and assigns it to the AcsTransID field.
func (o *ThreeDS2ResponseData) SetAcsTransID(v string) {
	o.AcsTransID = &v
}

// GetAcsURL returns the AcsURL field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetAcsURL() string {
	if o == nil || common.IsNil(o.AcsURL) {
		var ret string
		return ret
	}
	return *o.AcsURL
}

// GetAcsURLOk returns a tuple with the AcsURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetAcsURLOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcsURL) {
		return nil, false
	}
	return o.AcsURL, true
}

// HasAcsURL returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasAcsURL() bool {
	if o != nil && !common.IsNil(o.AcsURL) {
		return true
	}

	return false
}

// SetAcsURL gets a reference to the given string and assigns it to the AcsURL field.
func (o *ThreeDS2ResponseData) SetAcsURL(v string) {
	o.AcsURL = &v
}

// GetAuthenticationType returns the AuthenticationType field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetAuthenticationType() string {
	if o == nil || common.IsNil(o.AuthenticationType) {
		var ret string
		return ret
	}
	return *o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetAuthenticationTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AuthenticationType) {
		return nil, false
	}
	return o.AuthenticationType, true
}

// HasAuthenticationType returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasAuthenticationType() bool {
	if o != nil && !common.IsNil(o.AuthenticationType) {
		return true
	}

	return false
}

// SetAuthenticationType gets a reference to the given string and assigns it to the AuthenticationType field.
func (o *ThreeDS2ResponseData) SetAuthenticationType(v string) {
	o.AuthenticationType = &v
}

// GetCardHolderInfo returns the CardHolderInfo field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetCardHolderInfo() string {
	if o == nil || common.IsNil(o.CardHolderInfo) {
		var ret string
		return ret
	}
	return *o.CardHolderInfo
}

// GetCardHolderInfoOk returns a tuple with the CardHolderInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetCardHolderInfoOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardHolderInfo) {
		return nil, false
	}
	return o.CardHolderInfo, true
}

// HasCardHolderInfo returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasCardHolderInfo() bool {
	if o != nil && !common.IsNil(o.CardHolderInfo) {
		return true
	}

	return false
}

// SetCardHolderInfo gets a reference to the given string and assigns it to the CardHolderInfo field.
func (o *ThreeDS2ResponseData) SetCardHolderInfo(v string) {
	o.CardHolderInfo = &v
}

// GetCavvAlgorithm returns the CavvAlgorithm field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetCavvAlgorithm() string {
	if o == nil || common.IsNil(o.CavvAlgorithm) {
		var ret string
		return ret
	}
	return *o.CavvAlgorithm
}

// GetCavvAlgorithmOk returns a tuple with the CavvAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetCavvAlgorithmOk() (*string, bool) {
	if o == nil || common.IsNil(o.CavvAlgorithm) {
		return nil, false
	}
	return o.CavvAlgorithm, true
}

// HasCavvAlgorithm returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasCavvAlgorithm() bool {
	if o != nil && !common.IsNil(o.CavvAlgorithm) {
		return true
	}

	return false
}

// SetCavvAlgorithm gets a reference to the given string and assigns it to the CavvAlgorithm field.
func (o *ThreeDS2ResponseData) SetCavvAlgorithm(v string) {
	o.CavvAlgorithm = &v
}

// GetChallengeIndicator returns the ChallengeIndicator field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetChallengeIndicator() string {
	if o == nil || common.IsNil(o.ChallengeIndicator) {
		var ret string
		return ret
	}
	return *o.ChallengeIndicator
}

// GetChallengeIndicatorOk returns a tuple with the ChallengeIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetChallengeIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChallengeIndicator) {
		return nil, false
	}
	return o.ChallengeIndicator, true
}

// HasChallengeIndicator returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasChallengeIndicator() bool {
	if o != nil && !common.IsNil(o.ChallengeIndicator) {
		return true
	}

	return false
}

// SetChallengeIndicator gets a reference to the given string and assigns it to the ChallengeIndicator field.
func (o *ThreeDS2ResponseData) SetChallengeIndicator(v string) {
	o.ChallengeIndicator = &v
}

// GetDsReferenceNumber returns the DsReferenceNumber field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetDsReferenceNumber() string {
	if o == nil || common.IsNil(o.DsReferenceNumber) {
		var ret string
		return ret
	}
	return *o.DsReferenceNumber
}

// GetDsReferenceNumberOk returns a tuple with the DsReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetDsReferenceNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.DsReferenceNumber) {
		return nil, false
	}
	return o.DsReferenceNumber, true
}

// HasDsReferenceNumber returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasDsReferenceNumber() bool {
	if o != nil && !common.IsNil(o.DsReferenceNumber) {
		return true
	}

	return false
}

// SetDsReferenceNumber gets a reference to the given string and assigns it to the DsReferenceNumber field.
func (o *ThreeDS2ResponseData) SetDsReferenceNumber(v string) {
	o.DsReferenceNumber = &v
}

// GetDsTransID returns the DsTransID field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetDsTransID() string {
	if o == nil || common.IsNil(o.DsTransID) {
		var ret string
		return ret
	}
	return *o.DsTransID
}

// GetDsTransIDOk returns a tuple with the DsTransID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetDsTransIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.DsTransID) {
		return nil, false
	}
	return o.DsTransID, true
}

// HasDsTransID returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasDsTransID() bool {
	if o != nil && !common.IsNil(o.DsTransID) {
		return true
	}

	return false
}

// SetDsTransID gets a reference to the given string and assigns it to the DsTransID field.
func (o *ThreeDS2ResponseData) SetDsTransID(v string) {
	o.DsTransID = &v
}

// GetExemptionIndicator returns the ExemptionIndicator field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetExemptionIndicator() string {
	if o == nil || common.IsNil(o.ExemptionIndicator) {
		var ret string
		return ret
	}
	return *o.ExemptionIndicator
}

// GetExemptionIndicatorOk returns a tuple with the ExemptionIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetExemptionIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExemptionIndicator) {
		return nil, false
	}
	return o.ExemptionIndicator, true
}

// HasExemptionIndicator returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasExemptionIndicator() bool {
	if o != nil && !common.IsNil(o.ExemptionIndicator) {
		return true
	}

	return false
}

// SetExemptionIndicator gets a reference to the given string and assigns it to the ExemptionIndicator field.
func (o *ThreeDS2ResponseData) SetExemptionIndicator(v string) {
	o.ExemptionIndicator = &v
}

// GetMessageVersion returns the MessageVersion field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetMessageVersion() string {
	if o == nil || common.IsNil(o.MessageVersion) {
		var ret string
		return ret
	}
	return *o.MessageVersion
}

// GetMessageVersionOk returns a tuple with the MessageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetMessageVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.MessageVersion) {
		return nil, false
	}
	return o.MessageVersion, true
}

// HasMessageVersion returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasMessageVersion() bool {
	if o != nil && !common.IsNil(o.MessageVersion) {
		return true
	}

	return false
}

// SetMessageVersion gets a reference to the given string and assigns it to the MessageVersion field.
func (o *ThreeDS2ResponseData) SetMessageVersion(v string) {
	o.MessageVersion = &v
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetRiskScore() string {
	if o == nil || common.IsNil(o.RiskScore) {
		var ret string
		return ret
	}
	return *o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetRiskScoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskScore) {
		return nil, false
	}
	return o.RiskScore, true
}

// HasRiskScore returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasRiskScore() bool {
	if o != nil && !common.IsNil(o.RiskScore) {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given string and assigns it to the RiskScore field.
func (o *ThreeDS2ResponseData) SetRiskScore(v string) {
	o.RiskScore = &v
}

// GetSdkEphemPubKey returns the SdkEphemPubKey field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetSdkEphemPubKey() string {
	if o == nil || common.IsNil(o.SdkEphemPubKey) {
		var ret string
		return ret
	}
	return *o.SdkEphemPubKey
}

// GetSdkEphemPubKeyOk returns a tuple with the SdkEphemPubKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetSdkEphemPubKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.SdkEphemPubKey) {
		return nil, false
	}
	return o.SdkEphemPubKey, true
}

// HasSdkEphemPubKey returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasSdkEphemPubKey() bool {
	if o != nil && !common.IsNil(o.SdkEphemPubKey) {
		return true
	}

	return false
}

// SetSdkEphemPubKey gets a reference to the given string and assigns it to the SdkEphemPubKey field.
func (o *ThreeDS2ResponseData) SetSdkEphemPubKey(v string) {
	o.SdkEphemPubKey = &v
}

// GetThreeDSServerTransID returns the ThreeDSServerTransID field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetThreeDSServerTransID() string {
	if o == nil || common.IsNil(o.ThreeDSServerTransID) {
		var ret string
		return ret
	}
	return *o.ThreeDSServerTransID
}

// GetThreeDSServerTransIDOk returns a tuple with the ThreeDSServerTransID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetThreeDSServerTransIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSServerTransID) {
		return nil, false
	}
	return o.ThreeDSServerTransID, true
}

// HasThreeDSServerTransID returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasThreeDSServerTransID() bool {
	if o != nil && !common.IsNil(o.ThreeDSServerTransID) {
		return true
	}

	return false
}

// SetThreeDSServerTransID gets a reference to the given string and assigns it to the ThreeDSServerTransID field.
func (o *ThreeDS2ResponseData) SetThreeDSServerTransID(v string) {
	o.ThreeDSServerTransID = &v
}

// GetTransStatus returns the TransStatus field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetTransStatus() string {
	if o == nil || common.IsNil(o.TransStatus) {
		var ret string
		return ret
	}
	return *o.TransStatus
}

// GetTransStatusOk returns a tuple with the TransStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetTransStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransStatus) {
		return nil, false
	}
	return o.TransStatus, true
}

// HasTransStatus returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasTransStatus() bool {
	if o != nil && !common.IsNil(o.TransStatus) {
		return true
	}

	return false
}

// SetTransStatus gets a reference to the given string and assigns it to the TransStatus field.
func (o *ThreeDS2ResponseData) SetTransStatus(v string) {
	o.TransStatus = &v
}

// GetTransStatusReason returns the TransStatusReason field value if set, zero value otherwise.
func (o *ThreeDS2ResponseData) GetTransStatusReason() string {
	if o == nil || common.IsNil(o.TransStatusReason) {
		var ret string
		return ret
	}
	return *o.TransStatusReason
}

// GetTransStatusReasonOk returns a tuple with the TransStatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResponseData) GetTransStatusReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransStatusReason) {
		return nil, false
	}
	return o.TransStatusReason, true
}

// HasTransStatusReason returns a boolean if a field has been set.
func (o *ThreeDS2ResponseData) HasTransStatusReason() bool {
	if o != nil && !common.IsNil(o.TransStatusReason) {
		return true
	}

	return false
}

// SetTransStatusReason gets a reference to the given string and assigns it to the TransStatusReason field.
func (o *ThreeDS2ResponseData) SetTransStatusReason(v string) {
	o.TransStatusReason = &v
}

func (o ThreeDS2ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDS2ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcsChallengeMandated) {
		toSerialize["acsChallengeMandated"] = o.AcsChallengeMandated
	}
	if !common.IsNil(o.AcsOperatorID) {
		toSerialize["acsOperatorID"] = o.AcsOperatorID
	}
	if !common.IsNil(o.AcsReferenceNumber) {
		toSerialize["acsReferenceNumber"] = o.AcsReferenceNumber
	}
	if !common.IsNil(o.AcsSignedContent) {
		toSerialize["acsSignedContent"] = o.AcsSignedContent
	}
	if !common.IsNil(o.AcsTransID) {
		toSerialize["acsTransID"] = o.AcsTransID
	}
	if !common.IsNil(o.AcsURL) {
		toSerialize["acsURL"] = o.AcsURL
	}
	if !common.IsNil(o.AuthenticationType) {
		toSerialize["authenticationType"] = o.AuthenticationType
	}
	if !common.IsNil(o.CardHolderInfo) {
		toSerialize["cardHolderInfo"] = o.CardHolderInfo
	}
	if !common.IsNil(o.CavvAlgorithm) {
		toSerialize["cavvAlgorithm"] = o.CavvAlgorithm
	}
	if !common.IsNil(o.ChallengeIndicator) {
		toSerialize["challengeIndicator"] = o.ChallengeIndicator
	}
	if !common.IsNil(o.DsReferenceNumber) {
		toSerialize["dsReferenceNumber"] = o.DsReferenceNumber
	}
	if !common.IsNil(o.DsTransID) {
		toSerialize["dsTransID"] = o.DsTransID
	}
	if !common.IsNil(o.ExemptionIndicator) {
		toSerialize["exemptionIndicator"] = o.ExemptionIndicator
	}
	if !common.IsNil(o.MessageVersion) {
		toSerialize["messageVersion"] = o.MessageVersion
	}
	if !common.IsNil(o.RiskScore) {
		toSerialize["riskScore"] = o.RiskScore
	}
	if !common.IsNil(o.SdkEphemPubKey) {
		toSerialize["sdkEphemPubKey"] = o.SdkEphemPubKey
	}
	if !common.IsNil(o.ThreeDSServerTransID) {
		toSerialize["threeDSServerTransID"] = o.ThreeDSServerTransID
	}
	if !common.IsNil(o.TransStatus) {
		toSerialize["transStatus"] = o.TransStatus
	}
	if !common.IsNil(o.TransStatusReason) {
		toSerialize["transStatusReason"] = o.TransStatusReason
	}
	return toSerialize, nil
}

type NullableThreeDS2ResponseData struct {
	value *ThreeDS2ResponseData
	isSet bool
}

func (v NullableThreeDS2ResponseData) Get() *ThreeDS2ResponseData {
	return v.value
}

func (v *NullableThreeDS2ResponseData) Set(val *ThreeDS2ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDS2ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDS2ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDS2ResponseData(val *ThreeDS2ResponseData) *NullableThreeDS2ResponseData {
	return &NullableThreeDS2ResponseData{value: val, isSet: true}
}

func (v NullableThreeDS2ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDS2ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
