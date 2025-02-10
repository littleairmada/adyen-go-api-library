/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the Organization type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Organization{}

// Organization struct for Organization
type Organization struct {
	// The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the governing country.
	CountryOfGoverningLaw *string `json:"countryOfGoverningLaw,omitempty"`
	// The date when the organization was incorporated in YYYY-MM-DD format.
	DateOfIncorporation *string `json:"dateOfIncorporation,omitempty"`
	// Your description for the organization.
	Description *string `json:"description,omitempty"`
	// The organization's trading name, if different from the registered legal name.
	DoingBusinessAs *string `json:"doingBusinessAs,omitempty"`
	// The email address of the legal entity.
	Email *string `json:"email,omitempty"`
	// The financial report information of the organization.
	FinancialReports []FinancialReport `json:"financialReports,omitempty"`
	// The organization's legal name.
	LegalName                string       `json:"legalName"`
	Phone                    *PhoneNumber `json:"phone,omitempty"`
	PrincipalPlaceOfBusiness *Address     `json:"principalPlaceOfBusiness,omitempty"`
	RegisteredAddress        Address      `json:"registeredAddress"`
	// The organization's registration number.
	RegistrationNumber *string    `json:"registrationNumber,omitempty"`
	StockData          *StockData `json:"stockData,omitempty"`
	// The tax information of the organization.
	TaxInformation             []TaxInformation            `json:"taxInformation,omitempty"`
	TaxReportingClassification *TaxReportingClassification `json:"taxReportingClassification,omitempty"`
	// Type of organization.  Possible values: **associationIncorporated**, **governmentalOrganization**, **listedPublicCompany**, **nonProfit**, **partnershipIncorporated**, **privateCompany**.
	Type *string `json:"type,omitempty"`
	// The reason the organization has not provided a VAT number.  Possible values: **industryExemption**, **belowTaxThreshold**.
	VatAbsenceReason *string `json:"vatAbsenceReason,omitempty"`
	// The organization's VAT number.
	VatNumber *string  `json:"vatNumber,omitempty"`
	WebData   *WebData `json:"webData,omitempty"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization(legalName string, registeredAddress Address) *Organization {
	this := Organization{}
	this.LegalName = legalName
	this.RegisteredAddress = registeredAddress
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetCountryOfGoverningLaw returns the CountryOfGoverningLaw field value if set, zero value otherwise.
func (o *Organization) GetCountryOfGoverningLaw() string {
	if o == nil || common.IsNil(o.CountryOfGoverningLaw) {
		var ret string
		return ret
	}
	return *o.CountryOfGoverningLaw
}

// GetCountryOfGoverningLawOk returns a tuple with the CountryOfGoverningLaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCountryOfGoverningLawOk() (*string, bool) {
	if o == nil || common.IsNil(o.CountryOfGoverningLaw) {
		return nil, false
	}
	return o.CountryOfGoverningLaw, true
}

// HasCountryOfGoverningLaw returns a boolean if a field has been set.
func (o *Organization) HasCountryOfGoverningLaw() bool {
	if o != nil && !common.IsNil(o.CountryOfGoverningLaw) {
		return true
	}

	return false
}

// SetCountryOfGoverningLaw gets a reference to the given string and assigns it to the CountryOfGoverningLaw field.
func (o *Organization) SetCountryOfGoverningLaw(v string) {
	o.CountryOfGoverningLaw = &v
}

// GetDateOfIncorporation returns the DateOfIncorporation field value if set, zero value otherwise.
func (o *Organization) GetDateOfIncorporation() string {
	if o == nil || common.IsNil(o.DateOfIncorporation) {
		var ret string
		return ret
	}
	return *o.DateOfIncorporation
}

// GetDateOfIncorporationOk returns a tuple with the DateOfIncorporation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetDateOfIncorporationOk() (*string, bool) {
	if o == nil || common.IsNil(o.DateOfIncorporation) {
		return nil, false
	}
	return o.DateOfIncorporation, true
}

// HasDateOfIncorporation returns a boolean if a field has been set.
func (o *Organization) HasDateOfIncorporation() bool {
	if o != nil && !common.IsNil(o.DateOfIncorporation) {
		return true
	}

	return false
}

// SetDateOfIncorporation gets a reference to the given string and assigns it to the DateOfIncorporation field.
func (o *Organization) SetDateOfIncorporation(v string) {
	o.DateOfIncorporation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Organization) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Organization) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Organization) SetDescription(v string) {
	o.Description = &v
}

// GetDoingBusinessAs returns the DoingBusinessAs field value if set, zero value otherwise.
func (o *Organization) GetDoingBusinessAs() string {
	if o == nil || common.IsNil(o.DoingBusinessAs) {
		var ret string
		return ret
	}
	return *o.DoingBusinessAs
}

// GetDoingBusinessAsOk returns a tuple with the DoingBusinessAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetDoingBusinessAsOk() (*string, bool) {
	if o == nil || common.IsNil(o.DoingBusinessAs) {
		return nil, false
	}
	return o.DoingBusinessAs, true
}

// HasDoingBusinessAs returns a boolean if a field has been set.
func (o *Organization) HasDoingBusinessAs() bool {
	if o != nil && !common.IsNil(o.DoingBusinessAs) {
		return true
	}

	return false
}

// SetDoingBusinessAs gets a reference to the given string and assigns it to the DoingBusinessAs field.
func (o *Organization) SetDoingBusinessAs(v string) {
	o.DoingBusinessAs = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Organization) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Organization) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Organization) SetEmail(v string) {
	o.Email = &v
}

// GetFinancialReports returns the FinancialReports field value if set, zero value otherwise.
func (o *Organization) GetFinancialReports() []FinancialReport {
	if o == nil || common.IsNil(o.FinancialReports) {
		var ret []FinancialReport
		return ret
	}
	return o.FinancialReports
}

// GetFinancialReportsOk returns a tuple with the FinancialReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetFinancialReportsOk() ([]FinancialReport, bool) {
	if o == nil || common.IsNil(o.FinancialReports) {
		return nil, false
	}
	return o.FinancialReports, true
}

// HasFinancialReports returns a boolean if a field has been set.
func (o *Organization) HasFinancialReports() bool {
	if o != nil && !common.IsNil(o.FinancialReports) {
		return true
	}

	return false
}

// SetFinancialReports gets a reference to the given []FinancialReport and assigns it to the FinancialReports field.
func (o *Organization) SetFinancialReports(v []FinancialReport) {
	o.FinancialReports = v
}

// GetLegalName returns the LegalName field value
func (o *Organization) GetLegalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value
// and a boolean to check if the value has been set.
func (o *Organization) GetLegalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegalName, true
}

// SetLegalName sets field value
func (o *Organization) SetLegalName(v string) {
	o.LegalName = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Organization) GetPhone() PhoneNumber {
	if o == nil || common.IsNil(o.Phone) {
		var ret PhoneNumber
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetPhoneOk() (*PhoneNumber, bool) {
	if o == nil || common.IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Organization) HasPhone() bool {
	if o != nil && !common.IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given PhoneNumber and assigns it to the Phone field.
func (o *Organization) SetPhone(v PhoneNumber) {
	o.Phone = &v
}

// GetPrincipalPlaceOfBusiness returns the PrincipalPlaceOfBusiness field value if set, zero value otherwise.
func (o *Organization) GetPrincipalPlaceOfBusiness() Address {
	if o == nil || common.IsNil(o.PrincipalPlaceOfBusiness) {
		var ret Address
		return ret
	}
	return *o.PrincipalPlaceOfBusiness
}

// GetPrincipalPlaceOfBusinessOk returns a tuple with the PrincipalPlaceOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetPrincipalPlaceOfBusinessOk() (*Address, bool) {
	if o == nil || common.IsNil(o.PrincipalPlaceOfBusiness) {
		return nil, false
	}
	return o.PrincipalPlaceOfBusiness, true
}

// HasPrincipalPlaceOfBusiness returns a boolean if a field has been set.
func (o *Organization) HasPrincipalPlaceOfBusiness() bool {
	if o != nil && !common.IsNil(o.PrincipalPlaceOfBusiness) {
		return true
	}

	return false
}

// SetPrincipalPlaceOfBusiness gets a reference to the given Address and assigns it to the PrincipalPlaceOfBusiness field.
func (o *Organization) SetPrincipalPlaceOfBusiness(v Address) {
	o.PrincipalPlaceOfBusiness = &v
}

// GetRegisteredAddress returns the RegisteredAddress field value
func (o *Organization) GetRegisteredAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.RegisteredAddress
}

// GetRegisteredAddressOk returns a tuple with the RegisteredAddress field value
// and a boolean to check if the value has been set.
func (o *Organization) GetRegisteredAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegisteredAddress, true
}

// SetRegisteredAddress sets field value
func (o *Organization) SetRegisteredAddress(v Address) {
	o.RegisteredAddress = v
}

// GetRegistrationNumber returns the RegistrationNumber field value if set, zero value otherwise.
func (o *Organization) GetRegistrationNumber() string {
	if o == nil || common.IsNil(o.RegistrationNumber) {
		var ret string
		return ret
	}
	return *o.RegistrationNumber
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetRegistrationNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.RegistrationNumber) {
		return nil, false
	}
	return o.RegistrationNumber, true
}

// HasRegistrationNumber returns a boolean if a field has been set.
func (o *Organization) HasRegistrationNumber() bool {
	if o != nil && !common.IsNil(o.RegistrationNumber) {
		return true
	}

	return false
}

// SetRegistrationNumber gets a reference to the given string and assigns it to the RegistrationNumber field.
func (o *Organization) SetRegistrationNumber(v string) {
	o.RegistrationNumber = &v
}

// GetStockData returns the StockData field value if set, zero value otherwise.
func (o *Organization) GetStockData() StockData {
	if o == nil || common.IsNil(o.StockData) {
		var ret StockData
		return ret
	}
	return *o.StockData
}

// GetStockDataOk returns a tuple with the StockData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetStockDataOk() (*StockData, bool) {
	if o == nil || common.IsNil(o.StockData) {
		return nil, false
	}
	return o.StockData, true
}

// HasStockData returns a boolean if a field has been set.
func (o *Organization) HasStockData() bool {
	if o != nil && !common.IsNil(o.StockData) {
		return true
	}

	return false
}

// SetStockData gets a reference to the given StockData and assigns it to the StockData field.
func (o *Organization) SetStockData(v StockData) {
	o.StockData = &v
}

// GetTaxInformation returns the TaxInformation field value if set, zero value otherwise.
func (o *Organization) GetTaxInformation() []TaxInformation {
	if o == nil || common.IsNil(o.TaxInformation) {
		var ret []TaxInformation
		return ret
	}
	return o.TaxInformation
}

// GetTaxInformationOk returns a tuple with the TaxInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetTaxInformationOk() ([]TaxInformation, bool) {
	if o == nil || common.IsNil(o.TaxInformation) {
		return nil, false
	}
	return o.TaxInformation, true
}

// HasTaxInformation returns a boolean if a field has been set.
func (o *Organization) HasTaxInformation() bool {
	if o != nil && !common.IsNil(o.TaxInformation) {
		return true
	}

	return false
}

// SetTaxInformation gets a reference to the given []TaxInformation and assigns it to the TaxInformation field.
func (o *Organization) SetTaxInformation(v []TaxInformation) {
	o.TaxInformation = v
}

// GetTaxReportingClassification returns the TaxReportingClassification field value if set, zero value otherwise.
func (o *Organization) GetTaxReportingClassification() TaxReportingClassification {
	if o == nil || common.IsNil(o.TaxReportingClassification) {
		var ret TaxReportingClassification
		return ret
	}
	return *o.TaxReportingClassification
}

// GetTaxReportingClassificationOk returns a tuple with the TaxReportingClassification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetTaxReportingClassificationOk() (*TaxReportingClassification, bool) {
	if o == nil || common.IsNil(o.TaxReportingClassification) {
		return nil, false
	}
	return o.TaxReportingClassification, true
}

// HasTaxReportingClassification returns a boolean if a field has been set.
func (o *Organization) HasTaxReportingClassification() bool {
	if o != nil && !common.IsNil(o.TaxReportingClassification) {
		return true
	}

	return false
}

// SetTaxReportingClassification gets a reference to the given TaxReportingClassification and assigns it to the TaxReportingClassification field.
func (o *Organization) SetTaxReportingClassification(v TaxReportingClassification) {
	o.TaxReportingClassification = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Organization) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Organization) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Organization) SetType(v string) {
	o.Type = &v
}

// GetVatAbsenceReason returns the VatAbsenceReason field value if set, zero value otherwise.
func (o *Organization) GetVatAbsenceReason() string {
	if o == nil || common.IsNil(o.VatAbsenceReason) {
		var ret string
		return ret
	}
	return *o.VatAbsenceReason
}

// GetVatAbsenceReasonOk returns a tuple with the VatAbsenceReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetVatAbsenceReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.VatAbsenceReason) {
		return nil, false
	}
	return o.VatAbsenceReason, true
}

// HasVatAbsenceReason returns a boolean if a field has been set.
func (o *Organization) HasVatAbsenceReason() bool {
	if o != nil && !common.IsNil(o.VatAbsenceReason) {
		return true
	}

	return false
}

// SetVatAbsenceReason gets a reference to the given string and assigns it to the VatAbsenceReason field.
func (o *Organization) SetVatAbsenceReason(v string) {
	o.VatAbsenceReason = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *Organization) GetVatNumber() string {
	if o == nil || common.IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetVatNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *Organization) HasVatNumber() bool {
	if o != nil && !common.IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *Organization) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetWebData returns the WebData field value if set, zero value otherwise.
func (o *Organization) GetWebData() WebData {
	if o == nil || common.IsNil(o.WebData) {
		var ret WebData
		return ret
	}
	return *o.WebData
}

// GetWebDataOk returns a tuple with the WebData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetWebDataOk() (*WebData, bool) {
	if o == nil || common.IsNil(o.WebData) {
		return nil, false
	}
	return o.WebData, true
}

// HasWebData returns a boolean if a field has been set.
func (o *Organization) HasWebData() bool {
	if o != nil && !common.IsNil(o.WebData) {
		return true
	}

	return false
}

// SetWebData gets a reference to the given WebData and assigns it to the WebData field.
func (o *Organization) SetWebData(v WebData) {
	o.WebData = &v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CountryOfGoverningLaw) {
		toSerialize["countryOfGoverningLaw"] = o.CountryOfGoverningLaw
	}
	if !common.IsNil(o.DateOfIncorporation) {
		toSerialize["dateOfIncorporation"] = o.DateOfIncorporation
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.DoingBusinessAs) {
		toSerialize["doingBusinessAs"] = o.DoingBusinessAs
	}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.FinancialReports) {
		toSerialize["financialReports"] = o.FinancialReports
	}
	toSerialize["legalName"] = o.LegalName
	if !common.IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !common.IsNil(o.PrincipalPlaceOfBusiness) {
		toSerialize["principalPlaceOfBusiness"] = o.PrincipalPlaceOfBusiness
	}
	toSerialize["registeredAddress"] = o.RegisteredAddress
	if !common.IsNil(o.RegistrationNumber) {
		toSerialize["registrationNumber"] = o.RegistrationNumber
	}
	if !common.IsNil(o.StockData) {
		toSerialize["stockData"] = o.StockData
	}
	if !common.IsNil(o.TaxInformation) {
		toSerialize["taxInformation"] = o.TaxInformation
	}
	if !common.IsNil(o.TaxReportingClassification) {
		toSerialize["taxReportingClassification"] = o.TaxReportingClassification
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.VatAbsenceReason) {
		toSerialize["vatAbsenceReason"] = o.VatAbsenceReason
	}
	if !common.IsNil(o.VatNumber) {
		toSerialize["vatNumber"] = o.VatNumber
	}
	if !common.IsNil(o.WebData) {
		toSerialize["webData"] = o.WebData
	}
	return toSerialize, nil
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Organization) isValidType() bool {
	var allowedEnumValues = []string{"associationIncorporated", "governmentalOrganization", "listedPublicCompany", "nonProfit", "partnershipIncorporated", "privateCompany"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
func (o *Organization) isValidVatAbsenceReason() bool {
	var allowedEnumValues = []string{"industryExemption", "belowTaxThreshold"}
	for _, allowed := range allowedEnumValues {
		if o.GetVatAbsenceReason() == allowed {
			return true
		}
	}
	return false
}
