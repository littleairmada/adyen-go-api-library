/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the CardInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CardInfo{}

// CardInfo struct for CardInfo
type CardInfo struct {
	Authentication *Authentication `json:"authentication,omitempty"`
	// The brand of the physical or the virtual card. Possible values: **visa**, **mc**.
	Brand string `json:"brand"`
	// The brand variant of the physical or the virtual card. For example, **visadebit** or **mcprepaid**. >Reach out to your Adyen contact to get the values relevant for your integration.
	BrandVariant string `json:"brandVariant"`
	// The name of the cardholder.  Maximum length: 26 characters.
	CardholderName  string             `json:"cardholderName"`
	Configuration   *CardConfiguration `json:"configuration,omitempty"`
	DeliveryContact *DeliveryContact   `json:"deliveryContact,omitempty"`
	// The form factor of the card. Possible values: **virtual**, **physical**.
	FormFactor string `json:"formFactor"`
	// Allocates a specific product range for either a physical or a virtual card. Possible values: **fullySupported**, **secureCorporate**. >Reach out to your Adyen contact to get the values relevant for your integration.
	ThreeDSecure *string `json:"threeDSecure,omitempty"`
}

// NewCardInfo instantiates a new CardInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfo(brand string, brandVariant string, cardholderName string, formFactor string) *CardInfo {
	this := CardInfo{}
	this.Brand = brand
	this.BrandVariant = brandVariant
	this.CardholderName = cardholderName
	this.FormFactor = formFactor
	return &this
}

// NewCardInfoWithDefaults instantiates a new CardInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoWithDefaults() *CardInfo {
	this := CardInfo{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *CardInfo) GetAuthentication() Authentication {
	if o == nil || common.IsNil(o.Authentication) {
		var ret Authentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetAuthenticationOk() (*Authentication, bool) {
	if o == nil || common.IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *CardInfo) HasAuthentication() bool {
	if o != nil && !common.IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given Authentication and assigns it to the Authentication field.
func (o *CardInfo) SetAuthentication(v Authentication) {
	o.Authentication = &v
}

// GetBrand returns the Brand field value
func (o *CardInfo) GetBrand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Brand
}

// GetBrandOk returns a tuple with the Brand field value
// and a boolean to check if the value has been set.
func (o *CardInfo) GetBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Brand, true
}

// SetBrand sets field value
func (o *CardInfo) SetBrand(v string) {
	o.Brand = v
}

// GetBrandVariant returns the BrandVariant field value
func (o *CardInfo) GetBrandVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandVariant
}

// GetBrandVariantOk returns a tuple with the BrandVariant field value
// and a boolean to check if the value has been set.
func (o *CardInfo) GetBrandVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandVariant, true
}

// SetBrandVariant sets field value
func (o *CardInfo) SetBrandVariant(v string) {
	o.BrandVariant = v
}

// GetCardholderName returns the CardholderName field value
func (o *CardInfo) GetCardholderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardholderName
}

// GetCardholderNameOk returns a tuple with the CardholderName field value
// and a boolean to check if the value has been set.
func (o *CardInfo) GetCardholderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardholderName, true
}

// SetCardholderName sets field value
func (o *CardInfo) SetCardholderName(v string) {
	o.CardholderName = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *CardInfo) GetConfiguration() CardConfiguration {
	if o == nil || common.IsNil(o.Configuration) {
		var ret CardConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetConfigurationOk() (*CardConfiguration, bool) {
	if o == nil || common.IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *CardInfo) HasConfiguration() bool {
	if o != nil && !common.IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given CardConfiguration and assigns it to the Configuration field.
func (o *CardInfo) SetConfiguration(v CardConfiguration) {
	o.Configuration = &v
}

// GetDeliveryContact returns the DeliveryContact field value if set, zero value otherwise.
func (o *CardInfo) GetDeliveryContact() DeliveryContact {
	if o == nil || common.IsNil(o.DeliveryContact) {
		var ret DeliveryContact
		return ret
	}
	return *o.DeliveryContact
}

// GetDeliveryContactOk returns a tuple with the DeliveryContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetDeliveryContactOk() (*DeliveryContact, bool) {
	if o == nil || common.IsNil(o.DeliveryContact) {
		return nil, false
	}
	return o.DeliveryContact, true
}

// HasDeliveryContact returns a boolean if a field has been set.
func (o *CardInfo) HasDeliveryContact() bool {
	if o != nil && !common.IsNil(o.DeliveryContact) {
		return true
	}

	return false
}

// SetDeliveryContact gets a reference to the given DeliveryContact and assigns it to the DeliveryContact field.
func (o *CardInfo) SetDeliveryContact(v DeliveryContact) {
	o.DeliveryContact = &v
}

// GetFormFactor returns the FormFactor field value
func (o *CardInfo) GetFormFactor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormFactor
}

// GetFormFactorOk returns a tuple with the FormFactor field value
// and a boolean to check if the value has been set.
func (o *CardInfo) GetFormFactorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormFactor, true
}

// SetFormFactor sets field value
func (o *CardInfo) SetFormFactor(v string) {
	o.FormFactor = v
}

// GetThreeDSecure returns the ThreeDSecure field value if set, zero value otherwise.
func (o *CardInfo) GetThreeDSecure() string {
	if o == nil || common.IsNil(o.ThreeDSecure) {
		var ret string
		return ret
	}
	return *o.ThreeDSecure
}

// GetThreeDSecureOk returns a tuple with the ThreeDSecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetThreeDSecureOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSecure) {
		return nil, false
	}
	return o.ThreeDSecure, true
}

// HasThreeDSecure returns a boolean if a field has been set.
func (o *CardInfo) HasThreeDSecure() bool {
	if o != nil && !common.IsNil(o.ThreeDSecure) {
		return true
	}

	return false
}

// SetThreeDSecure gets a reference to the given string and assigns it to the ThreeDSecure field.
func (o *CardInfo) SetThreeDSecure(v string) {
	o.ThreeDSecure = &v
}

func (o CardInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	toSerialize["brand"] = o.Brand
	toSerialize["brandVariant"] = o.BrandVariant
	toSerialize["cardholderName"] = o.CardholderName
	if !common.IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !common.IsNil(o.DeliveryContact) {
		toSerialize["deliveryContact"] = o.DeliveryContact
	}
	toSerialize["formFactor"] = o.FormFactor
	if !common.IsNil(o.ThreeDSecure) {
		toSerialize["threeDSecure"] = o.ThreeDSecure
	}
	return toSerialize, nil
}

type NullableCardInfo struct {
	value *CardInfo
	isSet bool
}

func (v NullableCardInfo) Get() *CardInfo {
	return v.value
}

func (v *NullableCardInfo) Set(val *CardInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfo(val *CardInfo) *NullableCardInfo {
	return &NullableCardInfo{value: val, isSet: true}
}

func (v NullableCardInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CardInfo) isValidFormFactor() bool {
	var allowedEnumValues = []string{"physical", "unknown", "virtual"}
	for _, allowed := range allowedEnumValues {
		if o.GetFormFactor() == allowed {
			return true
		}
	}
	return false
}
