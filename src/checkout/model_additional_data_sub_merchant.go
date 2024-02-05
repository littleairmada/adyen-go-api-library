/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the AdditionalDataSubMerchant type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataSubMerchant{}

// AdditionalDataSubMerchant struct for AdditionalDataSubMerchant
type AdditionalDataSubMerchant struct {
	// Required for transactions performed by registered payment facilitators. Indicates the number of sub-merchants contained in the request. For example, **3**.
	SubMerchantNumberOfSubSellers *string `json:"subMerchant.numberOfSubSellers,omitempty"`
	// Required for transactions performed by registered payment facilitators. The city of the sub-merchant's address. * Format: Alphanumeric * Maximum length: 13 characters
	SubMerchantSubSellerSubSellerNrCity *string `json:"subMerchant.subSeller[subSellerNr].city,omitempty"`
	// Required for transactions performed by registered payment facilitators. The three-letter country code of the sub-merchant's address. For example, **BRA** for Brazil.  * Format: [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) * Fixed length: 3 characters
	SubMerchantSubSellerSubSellerNrCountry *string `json:"subMerchant.subSeller[subSellerNr].country,omitempty"`
	// Required for transactions performed by registered payment facilitators. A unique identifier that you create for the sub-merchant, used by schemes to identify the sub-merchant.  * Format: Alphanumeric * Maximum length: 15 characters
	SubMerchantSubSellerSubSellerNrId *string `json:"subMerchant.subSeller[subSellerNr].id,omitempty"`
	// Required for transactions performed by registered payment facilitators. The sub-merchant's 4-digit Merchant Category Code (MCC).  * Format: Numeric * Fixed length: 4 digits
	SubMerchantSubSellerSubSellerNrMcc *string `json:"subMerchant.subSeller[subSellerNr].mcc,omitempty"`
	// Required for transactions performed by registered payment facilitators. The name of the sub-merchant. Based on scheme specifications, this value will overwrite the shopper statement  that will appear in the card statement. * Format: Alphanumeric * Maximum length: 22 characters
	SubMerchantSubSellerSubSellerNrName *string `json:"subMerchant.subSeller[subSellerNr].name,omitempty"`
	// Required for transactions performed by registered payment facilitators. The postal code of the sub-merchant's address, without dashes. * Format: Numeric * Fixed length: 8 digits
	SubMerchantSubSellerSubSellerNrPostalCode *string `json:"subMerchant.subSeller[subSellerNr].postalCode,omitempty"`
	// Required for transactions performed by registered payment facilitators. The state code of the sub-merchant's address, if applicable to the country. * Format: Alphanumeric * Maximum length: 2 characters
	SubMerchantSubSellerSubSellerNrState *string `json:"subMerchant.subSeller[subSellerNr].state,omitempty"`
	// Required for transactions performed by registered payment facilitators. The street name and house number of the sub-merchant's address. * Format: Alphanumeric * Maximum length: 60 characters
	SubMerchantSubSellerSubSellerNrStreet *string `json:"subMerchant.subSeller[subSellerNr].street,omitempty"`
	// Required for transactions performed by registered payment facilitators. The tax ID of the sub-merchant. * Format: Numeric * Fixed length: 11 digits for the CPF or 14 digits for the CNPJ
	SubMerchantSubSellerSubSellerNrTaxId *string `json:"subMerchant.subSeller[subSellerNr].taxId,omitempty"`
}

// NewAdditionalDataSubMerchant instantiates a new AdditionalDataSubMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataSubMerchant() *AdditionalDataSubMerchant {
	this := AdditionalDataSubMerchant{}
	return &this
}

// NewAdditionalDataSubMerchantWithDefaults instantiates a new AdditionalDataSubMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataSubMerchantWithDefaults() *AdditionalDataSubMerchant {
	this := AdditionalDataSubMerchant{}
	return &this
}

// GetSubMerchantNumberOfSubSellers returns the SubMerchantNumberOfSubSellers field value if set, zero value otherwise.
func (o *AdditionalDataSubMerchant) GetSubMerchantNumberOfSubSellers() string {
	if o == nil || common.IsNil(o.SubMerchantNumberOfSubSellers) {
		var ret string
		return ret
	}
	return *o.SubMerchantNumberOfSubSellers
}

// GetSubMerchantNumberOfSubSellersOk returns a tuple with the SubMerchantNumberOfSubSellers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataSubMerchant) GetSubMerchantNumberOfSubSellersOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantNumberOfSubSellers) {
		return nil, false
	}
	return o.SubMerchantNumberOfSubSellers, true
}

// HasSubMerchantNumberOfSubSellers returns a boolean if a field has been set.
func (o *AdditionalDataSubMerchant) HasSubMerchantNumberOfSubSellers() bool {
	if o != nil && !common.IsNil(o.SubMerchantNumberOfSubSellers) {
		return true
	}

	return false
}

// SetSubMerchantNumberOfSubSellers gets a reference to the given string and assigns it to the SubMerchantNumberOfSubSellers field.
func (o *AdditionalDataSubMerchant) SetSubMerchantNumberOfSubSellers(v string) {
	o.SubMerchantNumberOfSubSellers = &v
}

// GetSubMerchantSubSellerSubSellerNrCity returns the SubMerchantSubSellerSubSellerNrCity field value if set, zero value otherwise.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrCity() string {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrCity) {
		var ret string
		return ret
	}
	return *o.SubMerchantSubSellerSubSellerNrCity
}

// GetSubMerchantSubSellerSubSellerNrCityOk returns a tuple with the SubMerchantSubSellerSubSellerNrCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrCity) {
		return nil, false
	}
	return o.SubMerchantSubSellerSubSellerNrCity, true
}

// HasSubMerchantSubSellerSubSellerNrCity returns a boolean if a field has been set.
func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrCity() bool {
	if o != nil && !common.IsNil(o.SubMerchantSubSellerSubSellerNrCity) {
		return true
	}

	return false
}

// SetSubMerchantSubSellerSubSellerNrCity gets a reference to the given string and assigns it to the SubMerchantSubSellerSubSellerNrCity field.
func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrCity(v string) {
	o.SubMerchantSubSellerSubSellerNrCity = &v
}

// GetSubMerchantSubSellerSubSellerNrCountry returns the SubMerchantSubSellerSubSellerNrCountry field value if set, zero value otherwise.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrCountry() string {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrCountry) {
		var ret string
		return ret
	}
	return *o.SubMerchantSubSellerSubSellerNrCountry
}

// GetSubMerchantSubSellerSubSellerNrCountryOk returns a tuple with the SubMerchantSubSellerSubSellerNrCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrCountry) {
		return nil, false
	}
	return o.SubMerchantSubSellerSubSellerNrCountry, true
}

// HasSubMerchantSubSellerSubSellerNrCountry returns a boolean if a field has been set.
func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrCountry() bool {
	if o != nil && !common.IsNil(o.SubMerchantSubSellerSubSellerNrCountry) {
		return true
	}

	return false
}

// SetSubMerchantSubSellerSubSellerNrCountry gets a reference to the given string and assigns it to the SubMerchantSubSellerSubSellerNrCountry field.
func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrCountry(v string) {
	o.SubMerchantSubSellerSubSellerNrCountry = &v
}

// GetSubMerchantSubSellerSubSellerNrId returns the SubMerchantSubSellerSubSellerNrId field value if set, zero value otherwise.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrId() string {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrId) {
		var ret string
		return ret
	}
	return *o.SubMerchantSubSellerSubSellerNrId
}

// GetSubMerchantSubSellerSubSellerNrIdOk returns a tuple with the SubMerchantSubSellerSubSellerNrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrId) {
		return nil, false
	}
	return o.SubMerchantSubSellerSubSellerNrId, true
}

// HasSubMerchantSubSellerSubSellerNrId returns a boolean if a field has been set.
func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrId() bool {
	if o != nil && !common.IsNil(o.SubMerchantSubSellerSubSellerNrId) {
		return true
	}

	return false
}

// SetSubMerchantSubSellerSubSellerNrId gets a reference to the given string and assigns it to the SubMerchantSubSellerSubSellerNrId field.
func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrId(v string) {
	o.SubMerchantSubSellerSubSellerNrId = &v
}

// GetSubMerchantSubSellerSubSellerNrMcc returns the SubMerchantSubSellerSubSellerNrMcc field value if set, zero value otherwise.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrMcc() string {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrMcc) {
		var ret string
		return ret
	}
	return *o.SubMerchantSubSellerSubSellerNrMcc
}

// GetSubMerchantSubSellerSubSellerNrMccOk returns a tuple with the SubMerchantSubSellerSubSellerNrMcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrMccOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrMcc) {
		return nil, false
	}
	return o.SubMerchantSubSellerSubSellerNrMcc, true
}

// HasSubMerchantSubSellerSubSellerNrMcc returns a boolean if a field has been set.
func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrMcc() bool {
	if o != nil && !common.IsNil(o.SubMerchantSubSellerSubSellerNrMcc) {
		return true
	}

	return false
}

// SetSubMerchantSubSellerSubSellerNrMcc gets a reference to the given string and assigns it to the SubMerchantSubSellerSubSellerNrMcc field.
func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrMcc(v string) {
	o.SubMerchantSubSellerSubSellerNrMcc = &v
}

// GetSubMerchantSubSellerSubSellerNrName returns the SubMerchantSubSellerSubSellerNrName field value if set, zero value otherwise.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrName() string {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrName) {
		var ret string
		return ret
	}
	return *o.SubMerchantSubSellerSubSellerNrName
}

// GetSubMerchantSubSellerSubSellerNrNameOk returns a tuple with the SubMerchantSubSellerSubSellerNrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrName) {
		return nil, false
	}
	return o.SubMerchantSubSellerSubSellerNrName, true
}

// HasSubMerchantSubSellerSubSellerNrName returns a boolean if a field has been set.
func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrName() bool {
	if o != nil && !common.IsNil(o.SubMerchantSubSellerSubSellerNrName) {
		return true
	}

	return false
}

// SetSubMerchantSubSellerSubSellerNrName gets a reference to the given string and assigns it to the SubMerchantSubSellerSubSellerNrName field.
func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrName(v string) {
	o.SubMerchantSubSellerSubSellerNrName = &v
}

// GetSubMerchantSubSellerSubSellerNrPostalCode returns the SubMerchantSubSellerSubSellerNrPostalCode field value if set, zero value otherwise.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrPostalCode() string {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrPostalCode) {
		var ret string
		return ret
	}
	return *o.SubMerchantSubSellerSubSellerNrPostalCode
}

// GetSubMerchantSubSellerSubSellerNrPostalCodeOk returns a tuple with the SubMerchantSubSellerSubSellerNrPostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrPostalCode) {
		return nil, false
	}
	return o.SubMerchantSubSellerSubSellerNrPostalCode, true
}

// HasSubMerchantSubSellerSubSellerNrPostalCode returns a boolean if a field has been set.
func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrPostalCode() bool {
	if o != nil && !common.IsNil(o.SubMerchantSubSellerSubSellerNrPostalCode) {
		return true
	}

	return false
}

// SetSubMerchantSubSellerSubSellerNrPostalCode gets a reference to the given string and assigns it to the SubMerchantSubSellerSubSellerNrPostalCode field.
func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrPostalCode(v string) {
	o.SubMerchantSubSellerSubSellerNrPostalCode = &v
}

// GetSubMerchantSubSellerSubSellerNrState returns the SubMerchantSubSellerSubSellerNrState field value if set, zero value otherwise.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrState() string {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrState) {
		var ret string
		return ret
	}
	return *o.SubMerchantSubSellerSubSellerNrState
}

// GetSubMerchantSubSellerSubSellerNrStateOk returns a tuple with the SubMerchantSubSellerSubSellerNrState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrStateOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrState) {
		return nil, false
	}
	return o.SubMerchantSubSellerSubSellerNrState, true
}

// HasSubMerchantSubSellerSubSellerNrState returns a boolean if a field has been set.
func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrState() bool {
	if o != nil && !common.IsNil(o.SubMerchantSubSellerSubSellerNrState) {
		return true
	}

	return false
}

// SetSubMerchantSubSellerSubSellerNrState gets a reference to the given string and assigns it to the SubMerchantSubSellerSubSellerNrState field.
func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrState(v string) {
	o.SubMerchantSubSellerSubSellerNrState = &v
}

// GetSubMerchantSubSellerSubSellerNrStreet returns the SubMerchantSubSellerSubSellerNrStreet field value if set, zero value otherwise.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrStreet() string {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrStreet) {
		var ret string
		return ret
	}
	return *o.SubMerchantSubSellerSubSellerNrStreet
}

// GetSubMerchantSubSellerSubSellerNrStreetOk returns a tuple with the SubMerchantSubSellerSubSellerNrStreet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrStreetOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrStreet) {
		return nil, false
	}
	return o.SubMerchantSubSellerSubSellerNrStreet, true
}

// HasSubMerchantSubSellerSubSellerNrStreet returns a boolean if a field has been set.
func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrStreet() bool {
	if o != nil && !common.IsNil(o.SubMerchantSubSellerSubSellerNrStreet) {
		return true
	}

	return false
}

// SetSubMerchantSubSellerSubSellerNrStreet gets a reference to the given string and assigns it to the SubMerchantSubSellerSubSellerNrStreet field.
func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrStreet(v string) {
	o.SubMerchantSubSellerSubSellerNrStreet = &v
}

// GetSubMerchantSubSellerSubSellerNrTaxId returns the SubMerchantSubSellerSubSellerNrTaxId field value if set, zero value otherwise.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrTaxId() string {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrTaxId) {
		var ret string
		return ret
	}
	return *o.SubMerchantSubSellerSubSellerNrTaxId
}

// GetSubMerchantSubSellerSubSellerNrTaxIdOk returns a tuple with the SubMerchantSubSellerSubSellerNrTaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataSubMerchant) GetSubMerchantSubSellerSubSellerNrTaxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubMerchantSubSellerSubSellerNrTaxId) {
		return nil, false
	}
	return o.SubMerchantSubSellerSubSellerNrTaxId, true
}

// HasSubMerchantSubSellerSubSellerNrTaxId returns a boolean if a field has been set.
func (o *AdditionalDataSubMerchant) HasSubMerchantSubSellerSubSellerNrTaxId() bool {
	if o != nil && !common.IsNil(o.SubMerchantSubSellerSubSellerNrTaxId) {
		return true
	}

	return false
}

// SetSubMerchantSubSellerSubSellerNrTaxId gets a reference to the given string and assigns it to the SubMerchantSubSellerSubSellerNrTaxId field.
func (o *AdditionalDataSubMerchant) SetSubMerchantSubSellerSubSellerNrTaxId(v string) {
	o.SubMerchantSubSellerSubSellerNrTaxId = &v
}

func (o AdditionalDataSubMerchant) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataSubMerchant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SubMerchantNumberOfSubSellers) {
		toSerialize["subMerchant.numberOfSubSellers"] = o.SubMerchantNumberOfSubSellers
	}
	if !common.IsNil(o.SubMerchantSubSellerSubSellerNrCity) {
		toSerialize["subMerchant.subSeller[subSellerNr].city"] = o.SubMerchantSubSellerSubSellerNrCity
	}
	if !common.IsNil(o.SubMerchantSubSellerSubSellerNrCountry) {
		toSerialize["subMerchant.subSeller[subSellerNr].country"] = o.SubMerchantSubSellerSubSellerNrCountry
	}
	if !common.IsNil(o.SubMerchantSubSellerSubSellerNrId) {
		toSerialize["subMerchant.subSeller[subSellerNr].id"] = o.SubMerchantSubSellerSubSellerNrId
	}
	if !common.IsNil(o.SubMerchantSubSellerSubSellerNrMcc) {
		toSerialize["subMerchant.subSeller[subSellerNr].mcc"] = o.SubMerchantSubSellerSubSellerNrMcc
	}
	if !common.IsNil(o.SubMerchantSubSellerSubSellerNrName) {
		toSerialize["subMerchant.subSeller[subSellerNr].name"] = o.SubMerchantSubSellerSubSellerNrName
	}
	if !common.IsNil(o.SubMerchantSubSellerSubSellerNrPostalCode) {
		toSerialize["subMerchant.subSeller[subSellerNr].postalCode"] = o.SubMerchantSubSellerSubSellerNrPostalCode
	}
	if !common.IsNil(o.SubMerchantSubSellerSubSellerNrState) {
		toSerialize["subMerchant.subSeller[subSellerNr].state"] = o.SubMerchantSubSellerSubSellerNrState
	}
	if !common.IsNil(o.SubMerchantSubSellerSubSellerNrStreet) {
		toSerialize["subMerchant.subSeller[subSellerNr].street"] = o.SubMerchantSubSellerSubSellerNrStreet
	}
	if !common.IsNil(o.SubMerchantSubSellerSubSellerNrTaxId) {
		toSerialize["subMerchant.subSeller[subSellerNr].taxId"] = o.SubMerchantSubSellerSubSellerNrTaxId
	}
	return toSerialize, nil
}

type NullableAdditionalDataSubMerchant struct {
	value *AdditionalDataSubMerchant
	isSet bool
}

func (v NullableAdditionalDataSubMerchant) Get() *AdditionalDataSubMerchant {
	return v.value
}

func (v *NullableAdditionalDataSubMerchant) Set(val *AdditionalDataSubMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataSubMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataSubMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataSubMerchant(val *AdditionalDataSubMerchant) *NullableAdditionalDataSubMerchant {
	return &NullableAdditionalDataSubMerchant{value: val, isSet: true}
}

func (v NullableAdditionalDataSubMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataSubMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
