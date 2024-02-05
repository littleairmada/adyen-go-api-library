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

// checks if the AccountHolderInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountHolderInfo{}

// AccountHolderInfo struct for AccountHolderInfo
type AccountHolderInfo struct {
	// The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the account holder belongs. Required in the request if your API credentials can be used for multiple balance platforms.
	BalancePlatform *string `json:"balancePlatform,omitempty"`
	// Contains key-value pairs that specify the actions that an account holder can do in your platform. The key is a capability required for your integration. For example, **issueCard** for Issuing. The value is an object containing the settings for the capability.
	Capabilities   *map[string]AccountHolderCapability `json:"capabilities,omitempty"`
	ContactDetails *ContactDetails                     `json:"contactDetails,omitempty"`
	// Your description for the account holder, maximum 300 characters.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id) associated with the account holder. Adyen performs a verification process against the legal entity of the account holder.
	LegalEntityId string `json:"legalEntityId"`
	// A set of key and value pairs for general use. The keys do not have specific names and may be used for storing miscellaneous data as desired. > Note that during an update of metadata, the omission of existing key-value pairs will result in the deletion of those key-value pairs.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The unique identifier of the migrated account holder in the classic integration.
	MigratedAccountHolderCode *string `json:"migratedAccountHolderCode,omitempty"`
	// Your reference for the account holder, maximum 150 characters.
	Reference *string `json:"reference,omitempty"`
	// The time zone of the account holder. For example, **Europe/Amsterdam**. Defaults to the time zone of the balance platform if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
	TimeZone *string `json:"timeZone,omitempty"`
}

// NewAccountHolderInfo instantiates a new AccountHolderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolderInfo(legalEntityId string) *AccountHolderInfo {
	this := AccountHolderInfo{}
	this.LegalEntityId = legalEntityId
	return &this
}

// NewAccountHolderInfoWithDefaults instantiates a new AccountHolderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderInfoWithDefaults() *AccountHolderInfo {
	this := AccountHolderInfo{}
	return &this
}

// GetBalancePlatform returns the BalancePlatform field value if set, zero value otherwise.
func (o *AccountHolderInfo) GetBalancePlatform() string {
	if o == nil || common.IsNil(o.BalancePlatform) {
		var ret string
		return ret
	}
	return *o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderInfo) GetBalancePlatformOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalancePlatform) {
		return nil, false
	}
	return o.BalancePlatform, true
}

// HasBalancePlatform returns a boolean if a field has been set.
func (o *AccountHolderInfo) HasBalancePlatform() bool {
	if o != nil && !common.IsNil(o.BalancePlatform) {
		return true
	}

	return false
}

// SetBalancePlatform gets a reference to the given string and assigns it to the BalancePlatform field.
func (o *AccountHolderInfo) SetBalancePlatform(v string) {
	o.BalancePlatform = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *AccountHolderInfo) GetCapabilities() map[string]AccountHolderCapability {
	if o == nil || common.IsNil(o.Capabilities) {
		var ret map[string]AccountHolderCapability
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderInfo) GetCapabilitiesOk() (*map[string]AccountHolderCapability, bool) {
	if o == nil || common.IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *AccountHolderInfo) HasCapabilities() bool {
	if o != nil && !common.IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given map[string]AccountHolderCapability and assigns it to the Capabilities field.
func (o *AccountHolderInfo) SetCapabilities(v map[string]AccountHolderCapability) {
	o.Capabilities = &v
}

// GetContactDetails returns the ContactDetails field value if set, zero value otherwise.
func (o *AccountHolderInfo) GetContactDetails() ContactDetails {
	if o == nil || common.IsNil(o.ContactDetails) {
		var ret ContactDetails
		return ret
	}
	return *o.ContactDetails
}

// GetContactDetailsOk returns a tuple with the ContactDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderInfo) GetContactDetailsOk() (*ContactDetails, bool) {
	if o == nil || common.IsNil(o.ContactDetails) {
		return nil, false
	}
	return o.ContactDetails, true
}

// HasContactDetails returns a boolean if a field has been set.
func (o *AccountHolderInfo) HasContactDetails() bool {
	if o != nil && !common.IsNil(o.ContactDetails) {
		return true
	}

	return false
}

// SetContactDetails gets a reference to the given ContactDetails and assigns it to the ContactDetails field.
func (o *AccountHolderInfo) SetContactDetails(v ContactDetails) {
	o.ContactDetails = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccountHolderInfo) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccountHolderInfo) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccountHolderInfo) SetDescription(v string) {
	o.Description = &v
}

// GetLegalEntityId returns the LegalEntityId field value
func (o *AccountHolderInfo) GetLegalEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value
// and a boolean to check if the value has been set.
func (o *AccountHolderInfo) GetLegalEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegalEntityId, true
}

// SetLegalEntityId sets field value
func (o *AccountHolderInfo) SetLegalEntityId(v string) {
	o.LegalEntityId = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AccountHolderInfo) GetMetadata() map[string]string {
	if o == nil || common.IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderInfo) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AccountHolderInfo) HasMetadata() bool {
	if o != nil && !common.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *AccountHolderInfo) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMigratedAccountHolderCode returns the MigratedAccountHolderCode field value if set, zero value otherwise.
func (o *AccountHolderInfo) GetMigratedAccountHolderCode() string {
	if o == nil || common.IsNil(o.MigratedAccountHolderCode) {
		var ret string
		return ret
	}
	return *o.MigratedAccountHolderCode
}

// GetMigratedAccountHolderCodeOk returns a tuple with the MigratedAccountHolderCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderInfo) GetMigratedAccountHolderCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MigratedAccountHolderCode) {
		return nil, false
	}
	return o.MigratedAccountHolderCode, true
}

// HasMigratedAccountHolderCode returns a boolean if a field has been set.
func (o *AccountHolderInfo) HasMigratedAccountHolderCode() bool {
	if o != nil && !common.IsNil(o.MigratedAccountHolderCode) {
		return true
	}

	return false
}

// SetMigratedAccountHolderCode gets a reference to the given string and assigns it to the MigratedAccountHolderCode field.
func (o *AccountHolderInfo) SetMigratedAccountHolderCode(v string) {
	o.MigratedAccountHolderCode = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *AccountHolderInfo) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderInfo) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *AccountHolderInfo) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *AccountHolderInfo) SetReference(v string) {
	o.Reference = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *AccountHolderInfo) GetTimeZone() string {
	if o == nil || common.IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderInfo) GetTimeZoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *AccountHolderInfo) HasTimeZone() bool {
	if o != nil && !common.IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *AccountHolderInfo) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o AccountHolderInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountHolderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BalancePlatform) {
		toSerialize["balancePlatform"] = o.BalancePlatform
	}
	if !common.IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !common.IsNil(o.ContactDetails) {
		toSerialize["contactDetails"] = o.ContactDetails
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["legalEntityId"] = o.LegalEntityId
	if !common.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !common.IsNil(o.MigratedAccountHolderCode) {
		toSerialize["migratedAccountHolderCode"] = o.MigratedAccountHolderCode
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	return toSerialize, nil
}

type NullableAccountHolderInfo struct {
	value *AccountHolderInfo
	isSet bool
}

func (v NullableAccountHolderInfo) Get() *AccountHolderInfo {
	return v.value
}

func (v *NullableAccountHolderInfo) Set(val *AccountHolderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderInfo(val *AccountHolderInfo) *NullableAccountHolderInfo {
	return &NullableAccountHolderInfo{value: val, isSet: true}
}

func (v NullableAccountHolderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
