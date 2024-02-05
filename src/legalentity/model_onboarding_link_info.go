/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the OnboardingLinkInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OnboardingLinkInfo{}

// OnboardingLinkInfo struct for OnboardingLinkInfo
type OnboardingLinkInfo struct {
	// The language that will be used for the page, specified by a combination of two letter [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language and [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. See [possible values](https://docs.adyen.com/marketplaces-and-platforms/collect-verification-details/hosted#supported-languages).   If not specified in the request or if the language is not supported, the page uses the browser language. If the browser language is not supported, the page uses **en-US** by default.
	Locale *string `json:"locale,omitempty"`
	// The URL where the user is redirected after they complete hosted onboarding.
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// Boolean key-value pairs indicating the settings for the hosted onboarding page. The keys are the settings.  Possible keys:  By default, these values are set to **true**. Set to **false** to not allow the action.  - **changeLegalEntityType**: The user can change their legal entity type.  - **editPrefilledCountry**: The user can change the country of their legal entity's address, for example the registered address of an organization.  By default, these values are set to **false**. Set to **true** to allow the action.  - **allowBankAccountFormatSelection**: The user can select the format for their payout account if applicable.  - **allowIntraRegionCrossBorderPayout**: The user can select a payout account in a different EU/EEA country than the country of their legal entity.  By default, these value are set to **false**. Set the following values to **true** to require the user to sign PCI questionnaires based on their sales channels. The user must sign PCI questionnaires for all relevant sales channels.  - **requirePciSignEcommerce**  - **requirePciSignPos**  - **requirePciSignEcomMoto**  - **requirePciSignPosMoto**
	Settings *map[string]bool `json:"settings,omitempty"`
	// The unique identifier of the hosted onboarding theme.
	ThemeId *string `json:"themeId,omitempty"`
}

// NewOnboardingLinkInfo instantiates a new OnboardingLinkInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnboardingLinkInfo() *OnboardingLinkInfo {
	this := OnboardingLinkInfo{}
	return &this
}

// NewOnboardingLinkInfoWithDefaults instantiates a new OnboardingLinkInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnboardingLinkInfoWithDefaults() *OnboardingLinkInfo {
	this := OnboardingLinkInfo{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *OnboardingLinkInfo) GetLocale() string {
	if o == nil || common.IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkInfo) GetLocaleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *OnboardingLinkInfo) HasLocale() bool {
	if o != nil && !common.IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *OnboardingLinkInfo) SetLocale(v string) {
	o.Locale = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *OnboardingLinkInfo) GetRedirectUrl() string {
	if o == nil || common.IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkInfo) GetRedirectUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *OnboardingLinkInfo) HasRedirectUrl() bool {
	if o != nil && !common.IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *OnboardingLinkInfo) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *OnboardingLinkInfo) GetSettings() map[string]bool {
	if o == nil || common.IsNil(o.Settings) {
		var ret map[string]bool
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkInfo) GetSettingsOk() (*map[string]bool, bool) {
	if o == nil || common.IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *OnboardingLinkInfo) HasSettings() bool {
	if o != nil && !common.IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]bool and assigns it to the Settings field.
func (o *OnboardingLinkInfo) SetSettings(v map[string]bool) {
	o.Settings = &v
}

// GetThemeId returns the ThemeId field value if set, zero value otherwise.
func (o *OnboardingLinkInfo) GetThemeId() string {
	if o == nil || common.IsNil(o.ThemeId) {
		var ret string
		return ret
	}
	return *o.ThemeId
}

// GetThemeIdOk returns a tuple with the ThemeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkInfo) GetThemeIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThemeId) {
		return nil, false
	}
	return o.ThemeId, true
}

// HasThemeId returns a boolean if a field has been set.
func (o *OnboardingLinkInfo) HasThemeId() bool {
	if o != nil && !common.IsNil(o.ThemeId) {
		return true
	}

	return false
}

// SetThemeId gets a reference to the given string and assigns it to the ThemeId field.
func (o *OnboardingLinkInfo) SetThemeId(v string) {
	o.ThemeId = &v
}

func (o OnboardingLinkInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnboardingLinkInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !common.IsNil(o.RedirectUrl) {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if !common.IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !common.IsNil(o.ThemeId) {
		toSerialize["themeId"] = o.ThemeId
	}
	return toSerialize, nil
}

type NullableOnboardingLinkInfo struct {
	value *OnboardingLinkInfo
	isSet bool
}

func (v NullableOnboardingLinkInfo) Get() *OnboardingLinkInfo {
	return v.value
}

func (v *NullableOnboardingLinkInfo) Set(val *OnboardingLinkInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOnboardingLinkInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOnboardingLinkInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnboardingLinkInfo(val *OnboardingLinkInfo) *NullableOnboardingLinkInfo {
	return &NullableOnboardingLinkInfo{value: val, isSet: true}
}

func (v NullableOnboardingLinkInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnboardingLinkInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
