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

// checks if the AdditionalDataLevel23 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataLevel23{}

// AdditionalDataLevel23 struct for AdditionalDataLevel23
type AdditionalDataLevel23 struct {
	// The customer code. * Encoding: ASCII * Max length: 25 characters * Must not start with a space or be all spaces * Must not be all zeros.
	EnhancedSchemeDataCustomerReference *string `json:"enhancedSchemeData.customerReference,omitempty"`
	// The three-letter [ISO 3166-1 alpha-3 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) for the destination address. * Encoding: ASCII * Fixed length: 3 characters
	EnhancedSchemeDataDestinationCountryCode *string `json:"enhancedSchemeData.destinationCountryCode,omitempty"`
	// The postal code of the destination address. * Encoding: ASCII * Max length: 10 characters * Must not start with a space
	EnhancedSchemeDataDestinationPostalCode *string `json:"enhancedSchemeData.destinationPostalCode,omitempty"`
	// Destination state or province code. * Encoding: ASCII * Max length: 3 characters * Must not start with a space
	EnhancedSchemeDataDestinationStateProvinceCode *string `json:"enhancedSchemeData.destinationStateProvinceCode,omitempty"`
	// The duty amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * For example, 2000 means USD 20.00. * Encoding: Numeric * Max length: 12 characters
	EnhancedSchemeDataDutyAmount *string `json:"enhancedSchemeData.dutyAmount,omitempty"`
	// The shipping amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * For example, 2000 means USD 20.00. * Encoding: Numeric *Max length: 12 characters
	EnhancedSchemeDataFreightAmount *string `json:"enhancedSchemeData.freightAmount,omitempty"`
	// The [UNSPC commodity code](https://www.unspsc.org/) of the item. * Encoding: ASCII * Max length: 12 characters * Must not start with a space or be all spaces * Must not be all zeros.
	EnhancedSchemeDataItemDetailLineItemNrCommodityCode *string `json:"enhancedSchemeData.itemDetailLine[itemNr].commodityCode,omitempty"`
	// A description of the item. * Encoding: ASCII * Max length: 26 characters * Must not start with a space or be all spaces * Must not be all zeros.
	EnhancedSchemeDataItemDetailLineItemNrDescription *string `json:"enhancedSchemeData.itemDetailLine[itemNr].description,omitempty"`
	// The discount amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * For example, 2000 means USD 20.00. * Encoding: Numeric * Max length: 12 characters
	EnhancedSchemeDataItemDetailLineItemNrDiscountAmount *string `json:"enhancedSchemeData.itemDetailLine[itemNr].discountAmount,omitempty"`
	// The product code. * Encoding: ASCII. * Max length: 12 characters * Must not start with a space or be all spaces * Must not be all zeros.
	EnhancedSchemeDataItemDetailLineItemNrProductCode *string `json:"enhancedSchemeData.itemDetailLine[itemNr].productCode,omitempty"`
	// The number of items. Must be an integer greater than zero. * Encoding: Numeric * Max length: 12 characters * Must not start with a space or be all spaces
	EnhancedSchemeDataItemDetailLineItemNrQuantity *string `json:"enhancedSchemeData.itemDetailLine[itemNr].quantity,omitempty"`
	// The total amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * For example, 2000 means USD 20.00. * Max length: 12 characters * Must not start with a space or be all spaces * Must not be all zeros.
	EnhancedSchemeDataItemDetailLineItemNrTotalAmount *string `json:"enhancedSchemeData.itemDetailLine[itemNr].totalAmount,omitempty"`
	// The unit of measurement for an item. * Encoding: ASCII  Max length: 3 characters * Must not start with a space or be all spaces * Must not be all zeros.
	EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure *string `json:"enhancedSchemeData.itemDetailLine[itemNr].unitOfMeasure,omitempty"`
	// The unit price in [minor units](https://docs.adyen.com/development-resources/currency-codes). * For example, 2000 means USD 20.00. * Encoding: Numeric * Max length: 12 characters * Must not be all zeros.
	EnhancedSchemeDataItemDetailLineItemNrUnitPrice *string `json:"enhancedSchemeData.itemDetailLine[itemNr].unitPrice,omitempty"`
	// The order date. * Format: `ddMMyy` * Encoding: ASCII * Max length: 6 characters
	EnhancedSchemeDataOrderDate *string `json:"enhancedSchemeData.orderDate,omitempty"`
	// The postal code of the address the item is shipped from. * Encoding: ASCII * Max length: 10 characters * Must not start with a space or be all spaces * Must not be all zeros.
	EnhancedSchemeDataShipFromPostalCode *string `json:"enhancedSchemeData.shipFromPostalCode,omitempty"`
	// The total tax amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * For example, 2000 means USD 20.00. *Encoding: Numeric *Max length: 12 characters * Must not be all zeros.
	EnhancedSchemeDataTotalTaxAmount *string `json:"enhancedSchemeData.totalTaxAmount,omitempty"`
}

// NewAdditionalDataLevel23 instantiates a new AdditionalDataLevel23 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataLevel23() *AdditionalDataLevel23 {
	this := AdditionalDataLevel23{}
	return &this
}

// NewAdditionalDataLevel23WithDefaults instantiates a new AdditionalDataLevel23 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataLevel23WithDefaults() *AdditionalDataLevel23 {
	this := AdditionalDataLevel23{}
	return &this
}

// GetEnhancedSchemeDataCustomerReference returns the EnhancedSchemeDataCustomerReference field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataCustomerReference() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataCustomerReference) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataCustomerReference
}

// GetEnhancedSchemeDataCustomerReferenceOk returns a tuple with the EnhancedSchemeDataCustomerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataCustomerReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataCustomerReference) {
		return nil, false
	}
	return o.EnhancedSchemeDataCustomerReference, true
}

// HasEnhancedSchemeDataCustomerReference returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataCustomerReference() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataCustomerReference) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataCustomerReference gets a reference to the given string and assigns it to the EnhancedSchemeDataCustomerReference field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataCustomerReference(v string) {
	o.EnhancedSchemeDataCustomerReference = &v
}

// GetEnhancedSchemeDataDestinationCountryCode returns the EnhancedSchemeDataDestinationCountryCode field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDestinationCountryCode() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataDestinationCountryCode) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataDestinationCountryCode
}

// GetEnhancedSchemeDataDestinationCountryCodeOk returns a tuple with the EnhancedSchemeDataDestinationCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDestinationCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataDestinationCountryCode) {
		return nil, false
	}
	return o.EnhancedSchemeDataDestinationCountryCode, true
}

// HasEnhancedSchemeDataDestinationCountryCode returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataDestinationCountryCode() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataDestinationCountryCode) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataDestinationCountryCode gets a reference to the given string and assigns it to the EnhancedSchemeDataDestinationCountryCode field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataDestinationCountryCode(v string) {
	o.EnhancedSchemeDataDestinationCountryCode = &v
}

// GetEnhancedSchemeDataDestinationPostalCode returns the EnhancedSchemeDataDestinationPostalCode field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDestinationPostalCode() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataDestinationPostalCode) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataDestinationPostalCode
}

// GetEnhancedSchemeDataDestinationPostalCodeOk returns a tuple with the EnhancedSchemeDataDestinationPostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDestinationPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataDestinationPostalCode) {
		return nil, false
	}
	return o.EnhancedSchemeDataDestinationPostalCode, true
}

// HasEnhancedSchemeDataDestinationPostalCode returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataDestinationPostalCode() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataDestinationPostalCode) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataDestinationPostalCode gets a reference to the given string and assigns it to the EnhancedSchemeDataDestinationPostalCode field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataDestinationPostalCode(v string) {
	o.EnhancedSchemeDataDestinationPostalCode = &v
}

// GetEnhancedSchemeDataDestinationStateProvinceCode returns the EnhancedSchemeDataDestinationStateProvinceCode field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDestinationStateProvinceCode() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataDestinationStateProvinceCode) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataDestinationStateProvinceCode
}

// GetEnhancedSchemeDataDestinationStateProvinceCodeOk returns a tuple with the EnhancedSchemeDataDestinationStateProvinceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDestinationStateProvinceCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataDestinationStateProvinceCode) {
		return nil, false
	}
	return o.EnhancedSchemeDataDestinationStateProvinceCode, true
}

// HasEnhancedSchemeDataDestinationStateProvinceCode returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataDestinationStateProvinceCode() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataDestinationStateProvinceCode) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataDestinationStateProvinceCode gets a reference to the given string and assigns it to the EnhancedSchemeDataDestinationStateProvinceCode field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataDestinationStateProvinceCode(v string) {
	o.EnhancedSchemeDataDestinationStateProvinceCode = &v
}

// GetEnhancedSchemeDataDutyAmount returns the EnhancedSchemeDataDutyAmount field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDutyAmount() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataDutyAmount) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataDutyAmount
}

// GetEnhancedSchemeDataDutyAmountOk returns a tuple with the EnhancedSchemeDataDutyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDutyAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataDutyAmount) {
		return nil, false
	}
	return o.EnhancedSchemeDataDutyAmount, true
}

// HasEnhancedSchemeDataDutyAmount returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataDutyAmount() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataDutyAmount) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataDutyAmount gets a reference to the given string and assigns it to the EnhancedSchemeDataDutyAmount field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataDutyAmount(v string) {
	o.EnhancedSchemeDataDutyAmount = &v
}

// GetEnhancedSchemeDataFreightAmount returns the EnhancedSchemeDataFreightAmount field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataFreightAmount() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataFreightAmount) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataFreightAmount
}

// GetEnhancedSchemeDataFreightAmountOk returns a tuple with the EnhancedSchemeDataFreightAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataFreightAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataFreightAmount) {
		return nil, false
	}
	return o.EnhancedSchemeDataFreightAmount, true
}

// HasEnhancedSchemeDataFreightAmount returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataFreightAmount() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataFreightAmount) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataFreightAmount gets a reference to the given string and assigns it to the EnhancedSchemeDataFreightAmount field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataFreightAmount(v string) {
	o.EnhancedSchemeDataFreightAmount = &v
}

// GetEnhancedSchemeDataItemDetailLineItemNrCommodityCode returns the EnhancedSchemeDataItemDetailLineItemNrCommodityCode field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrCommodityCode() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrCommodityCode) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataItemDetailLineItemNrCommodityCode
}

// GetEnhancedSchemeDataItemDetailLineItemNrCommodityCodeOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrCommodityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrCommodityCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrCommodityCode) {
		return nil, false
	}
	return o.EnhancedSchemeDataItemDetailLineItemNrCommodityCode, true
}

// HasEnhancedSchemeDataItemDetailLineItemNrCommodityCode returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrCommodityCode() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrCommodityCode) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataItemDetailLineItemNrCommodityCode gets a reference to the given string and assigns it to the EnhancedSchemeDataItemDetailLineItemNrCommodityCode field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrCommodityCode(v string) {
	o.EnhancedSchemeDataItemDetailLineItemNrCommodityCode = &v
}

// GetEnhancedSchemeDataItemDetailLineItemNrDescription returns the EnhancedSchemeDataItemDetailLineItemNrDescription field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrDescription() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrDescription) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataItemDetailLineItemNrDescription
}

// GetEnhancedSchemeDataItemDetailLineItemNrDescriptionOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrDescription) {
		return nil, false
	}
	return o.EnhancedSchemeDataItemDetailLineItemNrDescription, true
}

// HasEnhancedSchemeDataItemDetailLineItemNrDescription returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrDescription() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrDescription) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataItemDetailLineItemNrDescription gets a reference to the given string and assigns it to the EnhancedSchemeDataItemDetailLineItemNrDescription field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrDescription(v string) {
	o.EnhancedSchemeDataItemDetailLineItemNrDescription = &v
}

// GetEnhancedSchemeDataItemDetailLineItemNrDiscountAmount returns the EnhancedSchemeDataItemDetailLineItemNrDiscountAmount field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrDiscountAmount() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrDiscountAmount) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataItemDetailLineItemNrDiscountAmount
}

// GetEnhancedSchemeDataItemDetailLineItemNrDiscountAmountOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrDiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrDiscountAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrDiscountAmount) {
		return nil, false
	}
	return o.EnhancedSchemeDataItemDetailLineItemNrDiscountAmount, true
}

// HasEnhancedSchemeDataItemDetailLineItemNrDiscountAmount returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrDiscountAmount() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrDiscountAmount) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataItemDetailLineItemNrDiscountAmount gets a reference to the given string and assigns it to the EnhancedSchemeDataItemDetailLineItemNrDiscountAmount field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrDiscountAmount(v string) {
	o.EnhancedSchemeDataItemDetailLineItemNrDiscountAmount = &v
}

// GetEnhancedSchemeDataItemDetailLineItemNrProductCode returns the EnhancedSchemeDataItemDetailLineItemNrProductCode field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrProductCode() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrProductCode) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataItemDetailLineItemNrProductCode
}

// GetEnhancedSchemeDataItemDetailLineItemNrProductCodeOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrProductCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrProductCode) {
		return nil, false
	}
	return o.EnhancedSchemeDataItemDetailLineItemNrProductCode, true
}

// HasEnhancedSchemeDataItemDetailLineItemNrProductCode returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrProductCode() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrProductCode) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataItemDetailLineItemNrProductCode gets a reference to the given string and assigns it to the EnhancedSchemeDataItemDetailLineItemNrProductCode field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrProductCode(v string) {
	o.EnhancedSchemeDataItemDetailLineItemNrProductCode = &v
}

// GetEnhancedSchemeDataItemDetailLineItemNrQuantity returns the EnhancedSchemeDataItemDetailLineItemNrQuantity field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrQuantity() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrQuantity) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataItemDetailLineItemNrQuantity
}

// GetEnhancedSchemeDataItemDetailLineItemNrQuantityOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrQuantity) {
		return nil, false
	}
	return o.EnhancedSchemeDataItemDetailLineItemNrQuantity, true
}

// HasEnhancedSchemeDataItemDetailLineItemNrQuantity returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrQuantity() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrQuantity) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataItemDetailLineItemNrQuantity gets a reference to the given string and assigns it to the EnhancedSchemeDataItemDetailLineItemNrQuantity field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrQuantity(v string) {
	o.EnhancedSchemeDataItemDetailLineItemNrQuantity = &v
}

// GetEnhancedSchemeDataItemDetailLineItemNrTotalAmount returns the EnhancedSchemeDataItemDetailLineItemNrTotalAmount field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrTotalAmount() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrTotalAmount) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataItemDetailLineItemNrTotalAmount
}

// GetEnhancedSchemeDataItemDetailLineItemNrTotalAmountOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrTotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrTotalAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrTotalAmount) {
		return nil, false
	}
	return o.EnhancedSchemeDataItemDetailLineItemNrTotalAmount, true
}

// HasEnhancedSchemeDataItemDetailLineItemNrTotalAmount returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrTotalAmount() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrTotalAmount) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataItemDetailLineItemNrTotalAmount gets a reference to the given string and assigns it to the EnhancedSchemeDataItemDetailLineItemNrTotalAmount field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrTotalAmount(v string) {
	o.EnhancedSchemeDataItemDetailLineItemNrTotalAmount = &v
}

// GetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure returns the EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure
}

// GetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasureOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasureOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure) {
		return nil, false
	}
	return o.EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure, true
}

// HasEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure gets a reference to the given string and assigns it to the EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure(v string) {
	o.EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure = &v
}

// GetEnhancedSchemeDataItemDetailLineItemNrUnitPrice returns the EnhancedSchemeDataItemDetailLineItemNrUnitPrice field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrUnitPrice() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrUnitPrice) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataItemDetailLineItemNrUnitPrice
}

// GetEnhancedSchemeDataItemDetailLineItemNrUnitPriceOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrUnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrUnitPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrUnitPrice) {
		return nil, false
	}
	return o.EnhancedSchemeDataItemDetailLineItemNrUnitPrice, true
}

// HasEnhancedSchemeDataItemDetailLineItemNrUnitPrice returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrUnitPrice() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrUnitPrice) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataItemDetailLineItemNrUnitPrice gets a reference to the given string and assigns it to the EnhancedSchemeDataItemDetailLineItemNrUnitPrice field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrUnitPrice(v string) {
	o.EnhancedSchemeDataItemDetailLineItemNrUnitPrice = &v
}

// GetEnhancedSchemeDataOrderDate returns the EnhancedSchemeDataOrderDate field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataOrderDate() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataOrderDate) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataOrderDate
}

// GetEnhancedSchemeDataOrderDateOk returns a tuple with the EnhancedSchemeDataOrderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataOrderDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataOrderDate) {
		return nil, false
	}
	return o.EnhancedSchemeDataOrderDate, true
}

// HasEnhancedSchemeDataOrderDate returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataOrderDate() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataOrderDate) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataOrderDate gets a reference to the given string and assigns it to the EnhancedSchemeDataOrderDate field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataOrderDate(v string) {
	o.EnhancedSchemeDataOrderDate = &v
}

// GetEnhancedSchemeDataShipFromPostalCode returns the EnhancedSchemeDataShipFromPostalCode field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataShipFromPostalCode() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataShipFromPostalCode) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataShipFromPostalCode
}

// GetEnhancedSchemeDataShipFromPostalCodeOk returns a tuple with the EnhancedSchemeDataShipFromPostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataShipFromPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataShipFromPostalCode) {
		return nil, false
	}
	return o.EnhancedSchemeDataShipFromPostalCode, true
}

// HasEnhancedSchemeDataShipFromPostalCode returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataShipFromPostalCode() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataShipFromPostalCode) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataShipFromPostalCode gets a reference to the given string and assigns it to the EnhancedSchemeDataShipFromPostalCode field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataShipFromPostalCode(v string) {
	o.EnhancedSchemeDataShipFromPostalCode = &v
}

// GetEnhancedSchemeDataTotalTaxAmount returns the EnhancedSchemeDataTotalTaxAmount field value if set, zero value otherwise.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataTotalTaxAmount() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataTotalTaxAmount) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataTotalTaxAmount
}

// GetEnhancedSchemeDataTotalTaxAmountOk returns a tuple with the EnhancedSchemeDataTotalTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLevel23) GetEnhancedSchemeDataTotalTaxAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataTotalTaxAmount) {
		return nil, false
	}
	return o.EnhancedSchemeDataTotalTaxAmount, true
}

// HasEnhancedSchemeDataTotalTaxAmount returns a boolean if a field has been set.
func (o *AdditionalDataLevel23) HasEnhancedSchemeDataTotalTaxAmount() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataTotalTaxAmount) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataTotalTaxAmount gets a reference to the given string and assigns it to the EnhancedSchemeDataTotalTaxAmount field.
func (o *AdditionalDataLevel23) SetEnhancedSchemeDataTotalTaxAmount(v string) {
	o.EnhancedSchemeDataTotalTaxAmount = &v
}

func (o AdditionalDataLevel23) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataLevel23) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EnhancedSchemeDataCustomerReference) {
		toSerialize["enhancedSchemeData.customerReference"] = o.EnhancedSchemeDataCustomerReference
	}
	if !common.IsNil(o.EnhancedSchemeDataDestinationCountryCode) {
		toSerialize["enhancedSchemeData.destinationCountryCode"] = o.EnhancedSchemeDataDestinationCountryCode
	}
	if !common.IsNil(o.EnhancedSchemeDataDestinationPostalCode) {
		toSerialize["enhancedSchemeData.destinationPostalCode"] = o.EnhancedSchemeDataDestinationPostalCode
	}
	if !common.IsNil(o.EnhancedSchemeDataDestinationStateProvinceCode) {
		toSerialize["enhancedSchemeData.destinationStateProvinceCode"] = o.EnhancedSchemeDataDestinationStateProvinceCode
	}
	if !common.IsNil(o.EnhancedSchemeDataDutyAmount) {
		toSerialize["enhancedSchemeData.dutyAmount"] = o.EnhancedSchemeDataDutyAmount
	}
	if !common.IsNil(o.EnhancedSchemeDataFreightAmount) {
		toSerialize["enhancedSchemeData.freightAmount"] = o.EnhancedSchemeDataFreightAmount
	}
	if !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrCommodityCode) {
		toSerialize["enhancedSchemeData.itemDetailLine[itemNr].commodityCode"] = o.EnhancedSchemeDataItemDetailLineItemNrCommodityCode
	}
	if !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrDescription) {
		toSerialize["enhancedSchemeData.itemDetailLine[itemNr].description"] = o.EnhancedSchemeDataItemDetailLineItemNrDescription
	}
	if !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrDiscountAmount) {
		toSerialize["enhancedSchemeData.itemDetailLine[itemNr].discountAmount"] = o.EnhancedSchemeDataItemDetailLineItemNrDiscountAmount
	}
	if !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrProductCode) {
		toSerialize["enhancedSchemeData.itemDetailLine[itemNr].productCode"] = o.EnhancedSchemeDataItemDetailLineItemNrProductCode
	}
	if !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrQuantity) {
		toSerialize["enhancedSchemeData.itemDetailLine[itemNr].quantity"] = o.EnhancedSchemeDataItemDetailLineItemNrQuantity
	}
	if !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrTotalAmount) {
		toSerialize["enhancedSchemeData.itemDetailLine[itemNr].totalAmount"] = o.EnhancedSchemeDataItemDetailLineItemNrTotalAmount
	}
	if !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure) {
		toSerialize["enhancedSchemeData.itemDetailLine[itemNr].unitOfMeasure"] = o.EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure
	}
	if !common.IsNil(o.EnhancedSchemeDataItemDetailLineItemNrUnitPrice) {
		toSerialize["enhancedSchemeData.itemDetailLine[itemNr].unitPrice"] = o.EnhancedSchemeDataItemDetailLineItemNrUnitPrice
	}
	if !common.IsNil(o.EnhancedSchemeDataOrderDate) {
		toSerialize["enhancedSchemeData.orderDate"] = o.EnhancedSchemeDataOrderDate
	}
	if !common.IsNil(o.EnhancedSchemeDataShipFromPostalCode) {
		toSerialize["enhancedSchemeData.shipFromPostalCode"] = o.EnhancedSchemeDataShipFromPostalCode
	}
	if !common.IsNil(o.EnhancedSchemeDataTotalTaxAmount) {
		toSerialize["enhancedSchemeData.totalTaxAmount"] = o.EnhancedSchemeDataTotalTaxAmount
	}
	return toSerialize, nil
}

type NullableAdditionalDataLevel23 struct {
	value *AdditionalDataLevel23
	isSet bool
}

func (v NullableAdditionalDataLevel23) Get() *AdditionalDataLevel23 {
	return v.value
}

func (v *NullableAdditionalDataLevel23) Set(val *AdditionalDataLevel23) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataLevel23) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataLevel23) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataLevel23(val *AdditionalDataLevel23) *NullableAdditionalDataLevel23 {
	return &NullableAdditionalDataLevel23{value: val, isSet: true}
}

func (v NullableAdditionalDataLevel23) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataLevel23) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
