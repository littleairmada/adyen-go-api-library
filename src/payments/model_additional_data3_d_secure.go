/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the AdditionalData3DSecure type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalData3DSecure{}

// AdditionalData3DSecure struct for AdditionalData3DSecure
type AdditionalData3DSecure struct {
	// Indicates if you are able to process 3D Secure 2 transactions natively on your payment page. Send this parameter when you are using `/payments` endpoint with any of our [native 3D Secure 2 solutions](https://docs.adyen.com/online-payments/3d-secure/native-3ds2).   > This parameter only indicates readiness to support native 3D Secure 2 authentication. To specify if you _want_ to perform 3D Secure, use [Dynamic 3D Secure](/risk-management/dynamic-3d-secure) or send the `executeThreeD` parameter.  Possible values: * **true** - Ready to support native 3D Secure 2 authentication. Setting this to true does not mean always applying 3D Secure 2. Adyen still selects the version of 3D Secure based on configuration to optimize authorisation rates and improve the shopper's experience. * **false** – Not ready to support native 3D Secure 2 authentication. Adyen will not offer 3D Secure 2 to your shopper regardless of your configuration.
	Allow3DS2 *string `json:"allow3DS2,omitempty"`
	// Dimensions of the 3DS2 challenge window to be displayed to the cardholder.  Possible values:  * **01** - size of 250x400  * **02** - size of 390x400 * **03** - size of 500x600 * **04** - size of 600x400 * **05** - Fullscreen
	ChallengeWindowSize *string `json:"challengeWindowSize,omitempty"`
	// Indicates if you want to perform 3D Secure authentication on a transaction.   > Alternatively, you can use [Dynamic 3D Secure](/risk-management/dynamic-3d-secure) to configure rules for applying 3D Secure.  Possible values: * **true** – Perform 3D Secure authentication. * **false** – Don't perform 3D Secure authentication. Note that this setting results in refusals if the issuer mandates 3D Secure because of the PSD2 directive  or other, national regulations.
	ExecuteThreeD *string `json:"executeThreeD,omitempty"`
	// In case of Secure+, this field must be set to **CUPSecurePlus**.
	MpiImplementationType *string `json:"mpiImplementationType,omitempty"`
	// Indicates the [exemption type](https://docs.adyen.com/payments-fundamentals/psd2-sca-compliance-and-implementation-guide#specifypreferenceinyourapirequest) that you want to request for the transaction.   Possible values: * **lowValue**  * **secureCorporate**  * **trustedBeneficiary**  * **transactionRiskAnalysis**
	ScaExemption *string `json:"scaExemption,omitempty"`
	// Indicates your preference for the 3D Secure version.  > If you use this parameter, you override the checks from Adyen's Authentication Engine. We recommend to use this field only if you have an extensive knowledge of 3D Secure.  Possible values: * **1.0.2**: Apply 3D Secure version 1.0.2.  * **2.1.0**: Apply 3D Secure version 2.1.0.  * **2.2.0**: Apply 3D Secure version 2.2.0. If the issuer does not support version 2.2.0, we will fall back to 2.1.0.  The following rules apply: * If you prefer 2.1.0 or 2.2.0 but we receive a negative `transStatus` in the `ARes`, we will apply the fallback policy configured in your account. For example, if the configuration is to fall back to 3D Secure 1, we will apply version 1.0.2. * If you prefer 2.1.0 or 2.2.0 but the BIN is not enrolled, you will receive an error.
	ThreeDSVersion *string `json:"threeDSVersion,omitempty"`
}

// NewAdditionalData3DSecure instantiates a new AdditionalData3DSecure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalData3DSecure() *AdditionalData3DSecure {
	this := AdditionalData3DSecure{}
	return &this
}

// NewAdditionalData3DSecureWithDefaults instantiates a new AdditionalData3DSecure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalData3DSecureWithDefaults() *AdditionalData3DSecure {
	this := AdditionalData3DSecure{}
	return &this
}

// GetAllow3DS2 returns the Allow3DS2 field value if set, zero value otherwise.
func (o *AdditionalData3DSecure) GetAllow3DS2() string {
	if o == nil || common.IsNil(o.Allow3DS2) {
		var ret string
		return ret
	}
	return *o.Allow3DS2
}

// GetAllow3DS2Ok returns a tuple with the Allow3DS2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalData3DSecure) GetAllow3DS2Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Allow3DS2) {
		return nil, false
	}
	return o.Allow3DS2, true
}

// HasAllow3DS2 returns a boolean if a field has been set.
func (o *AdditionalData3DSecure) HasAllow3DS2() bool {
	if o != nil && !common.IsNil(o.Allow3DS2) {
		return true
	}

	return false
}

// SetAllow3DS2 gets a reference to the given string and assigns it to the Allow3DS2 field.
func (o *AdditionalData3DSecure) SetAllow3DS2(v string) {
	o.Allow3DS2 = &v
}

// GetChallengeWindowSize returns the ChallengeWindowSize field value if set, zero value otherwise.
func (o *AdditionalData3DSecure) GetChallengeWindowSize() string {
	if o == nil || common.IsNil(o.ChallengeWindowSize) {
		var ret string
		return ret
	}
	return *o.ChallengeWindowSize
}

// GetChallengeWindowSizeOk returns a tuple with the ChallengeWindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalData3DSecure) GetChallengeWindowSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChallengeWindowSize) {
		return nil, false
	}
	return o.ChallengeWindowSize, true
}

// HasChallengeWindowSize returns a boolean if a field has been set.
func (o *AdditionalData3DSecure) HasChallengeWindowSize() bool {
	if o != nil && !common.IsNil(o.ChallengeWindowSize) {
		return true
	}

	return false
}

// SetChallengeWindowSize gets a reference to the given string and assigns it to the ChallengeWindowSize field.
func (o *AdditionalData3DSecure) SetChallengeWindowSize(v string) {
	o.ChallengeWindowSize = &v
}

// GetExecuteThreeD returns the ExecuteThreeD field value if set, zero value otherwise.
func (o *AdditionalData3DSecure) GetExecuteThreeD() string {
	if o == nil || common.IsNil(o.ExecuteThreeD) {
		var ret string
		return ret
	}
	return *o.ExecuteThreeD
}

// GetExecuteThreeDOk returns a tuple with the ExecuteThreeD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalData3DSecure) GetExecuteThreeDOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecuteThreeD) {
		return nil, false
	}
	return o.ExecuteThreeD, true
}

// HasExecuteThreeD returns a boolean if a field has been set.
func (o *AdditionalData3DSecure) HasExecuteThreeD() bool {
	if o != nil && !common.IsNil(o.ExecuteThreeD) {
		return true
	}

	return false
}

// SetExecuteThreeD gets a reference to the given string and assigns it to the ExecuteThreeD field.
func (o *AdditionalData3DSecure) SetExecuteThreeD(v string) {
	o.ExecuteThreeD = &v
}

// GetMpiImplementationType returns the MpiImplementationType field value if set, zero value otherwise.
func (o *AdditionalData3DSecure) GetMpiImplementationType() string {
	if o == nil || common.IsNil(o.MpiImplementationType) {
		var ret string
		return ret
	}
	return *o.MpiImplementationType
}

// GetMpiImplementationTypeOk returns a tuple with the MpiImplementationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalData3DSecure) GetMpiImplementationTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MpiImplementationType) {
		return nil, false
	}
	return o.MpiImplementationType, true
}

// HasMpiImplementationType returns a boolean if a field has been set.
func (o *AdditionalData3DSecure) HasMpiImplementationType() bool {
	if o != nil && !common.IsNil(o.MpiImplementationType) {
		return true
	}

	return false
}

// SetMpiImplementationType gets a reference to the given string and assigns it to the MpiImplementationType field.
func (o *AdditionalData3DSecure) SetMpiImplementationType(v string) {
	o.MpiImplementationType = &v
}

// GetScaExemption returns the ScaExemption field value if set, zero value otherwise.
func (o *AdditionalData3DSecure) GetScaExemption() string {
	if o == nil || common.IsNil(o.ScaExemption) {
		var ret string
		return ret
	}
	return *o.ScaExemption
}

// GetScaExemptionOk returns a tuple with the ScaExemption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalData3DSecure) GetScaExemptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.ScaExemption) {
		return nil, false
	}
	return o.ScaExemption, true
}

// HasScaExemption returns a boolean if a field has been set.
func (o *AdditionalData3DSecure) HasScaExemption() bool {
	if o != nil && !common.IsNil(o.ScaExemption) {
		return true
	}

	return false
}

// SetScaExemption gets a reference to the given string and assigns it to the ScaExemption field.
func (o *AdditionalData3DSecure) SetScaExemption(v string) {
	o.ScaExemption = &v
}

// GetThreeDSVersion returns the ThreeDSVersion field value if set, zero value otherwise.
func (o *AdditionalData3DSecure) GetThreeDSVersion() string {
	if o == nil || common.IsNil(o.ThreeDSVersion) {
		var ret string
		return ret
	}
	return *o.ThreeDSVersion
}

// GetThreeDSVersionOk returns a tuple with the ThreeDSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalData3DSecure) GetThreeDSVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSVersion) {
		return nil, false
	}
	return o.ThreeDSVersion, true
}

// HasThreeDSVersion returns a boolean if a field has been set.
func (o *AdditionalData3DSecure) HasThreeDSVersion() bool {
	if o != nil && !common.IsNil(o.ThreeDSVersion) {
		return true
	}

	return false
}

// SetThreeDSVersion gets a reference to the given string and assigns it to the ThreeDSVersion field.
func (o *AdditionalData3DSecure) SetThreeDSVersion(v string) {
	o.ThreeDSVersion = &v
}

func (o AdditionalData3DSecure) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalData3DSecure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Allow3DS2) {
		toSerialize["allow3DS2"] = o.Allow3DS2
	}
	if !common.IsNil(o.ChallengeWindowSize) {
		toSerialize["challengeWindowSize"] = o.ChallengeWindowSize
	}
	if !common.IsNil(o.ExecuteThreeD) {
		toSerialize["executeThreeD"] = o.ExecuteThreeD
	}
	if !common.IsNil(o.MpiImplementationType) {
		toSerialize["mpiImplementationType"] = o.MpiImplementationType
	}
	if !common.IsNil(o.ScaExemption) {
		toSerialize["scaExemption"] = o.ScaExemption
	}
	if !common.IsNil(o.ThreeDSVersion) {
		toSerialize["threeDSVersion"] = o.ThreeDSVersion
	}
	return toSerialize, nil
}

type NullableAdditionalData3DSecure struct {
	value *AdditionalData3DSecure
	isSet bool
}

func (v NullableAdditionalData3DSecure) Get() *AdditionalData3DSecure {
	return v.value
}

func (v *NullableAdditionalData3DSecure) Set(val *AdditionalData3DSecure) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalData3DSecure) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalData3DSecure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalData3DSecure(val *AdditionalData3DSecure) *NullableAdditionalData3DSecure {
	return &NullableAdditionalData3DSecure{value: val, isSet: true}
}

func (v NullableAdditionalData3DSecure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalData3DSecure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AdditionalData3DSecure) isValidChallengeWindowSize() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04", "05"}
	for _, allowed := range allowedEnumValues {
		if o.GetChallengeWindowSize() == allowed {
			return true
		}
	}
	return false
}
