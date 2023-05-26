/*
Transfers API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the Address2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Address2{}

// Address2 struct for Address2
type Address2 struct {
	// The name of the city.
	City *string `json:"city,omitempty"`
	// The two-character ISO 3166-1 alpha-2 country code. For example, **US**, **NL**, or **GB**.
	Country string `json:"country"`
	// First line of the street address.
	Line1 *string `json:"line1,omitempty"`
	// Second line of the street address.
	Line2 *string `json:"line2,omitempty"`
	// The postal code. Maximum length: * 5 digits for an address in the US. * 10 characters for an address in all other countries.
	PostalCode *string `json:"postalCode,omitempty"`
	// The two-letter ISO 3166-2 state or province code. For example, **CA** in the US or **ON** in Canada. > Required for the US and Canada.
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
}

// NewAddress2 instantiates a new Address2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress2(country string) *Address2 {
	this := Address2{}
	this.Country = country
	return &this
}

// NewAddress2WithDefaults instantiates a new Address2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddress2WithDefaults() *Address2 {
	this := Address2{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Address2) GetCity() string {
	if o == nil || common.IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Address2) HasCity() bool {
	if o != nil && !common.IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Address2) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value
func (o *Address2) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *Address2) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *Address2) SetCountry(v string) {
	o.Country = v
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *Address2) GetLine1() string {
	if o == nil || common.IsNil(o.Line1) {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetLine1Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Line1) {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *Address2) HasLine1() bool {
	if o != nil && !common.IsNil(o.Line1) {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *Address2) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *Address2) GetLine2() string {
	if o == nil || common.IsNil(o.Line2) {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetLine2Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Line2) {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *Address2) HasLine2() bool {
	if o != nil && !common.IsNil(o.Line2) {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *Address2) SetLine2(v string) {
	o.Line2 = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Address2) GetPostalCode() string {
	if o == nil || common.IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Address2) HasPostalCode() bool {
	if o != nil && !common.IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Address2) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise.
func (o *Address2) GetStateOrProvince() string {
	if o == nil || common.IsNil(o.StateOrProvince) {
		var ret string
		return ret
	}
	return *o.StateOrProvince
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetStateOrProvinceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StateOrProvince) {
		return nil, false
	}
	return o.StateOrProvince, true
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *Address2) HasStateOrProvince() bool {
	if o != nil && !common.IsNil(o.StateOrProvince) {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given string and assigns it to the StateOrProvince field.
func (o *Address2) SetStateOrProvince(v string) {
	o.StateOrProvince = &v
}

func (o Address2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	toSerialize["country"] = o.Country
	if !common.IsNil(o.Line1) {
		toSerialize["line1"] = o.Line1
	}
	if !common.IsNil(o.Line2) {
		toSerialize["line2"] = o.Line2
	}
	if !common.IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !common.IsNil(o.StateOrProvince) {
		toSerialize["stateOrProvince"] = o.StateOrProvince
	}
	return toSerialize, nil
}

type NullableAddress2 struct {
	value *Address2
	isSet bool
}

func (v NullableAddress2) Get() *Address2 {
	return v.value
}

func (v *NullableAddress2) Set(val *Address2) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress2) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress2(val *Address2) *NullableAddress2 {
	return &NullableAddress2{value: val, isSet: true}
}

func (v NullableAddress2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
