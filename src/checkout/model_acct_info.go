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

// checks if the AcctInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcctInfo{}

// AcctInfo struct for AcctInfo
type AcctInfo struct {
	// Length of time that the cardholder has had the account with the 3DS Requestor.  Allowed values: * **01** — No account * **02** — Created during this transaction * **03** — Less than 30 days * **04** — 30–60 days * **05** — More than 60 days
	ChAccAgeInd *string `json:"chAccAgeInd,omitempty"`
	// Date that the cardholder’s account with the 3DS Requestor was last changed, including Billing or Shipping address, new payment account, or new user(s) added.  Format: **YYYYMMDD**
	ChAccChange *string `json:"chAccChange,omitempty"`
	// Length of time since the cardholder’s account information with the 3DS Requestor was last changed, including Billing or Shipping address, new payment account, or new user(s) added.  Allowed values: * **01** — Changed during this transaction * **02** — Less than 30 days * **03** — 30–60 days * **04** — More than 60 days
	ChAccChangeInd *string `json:"chAccChangeInd,omitempty"`
	// Date that cardholder’s account with the 3DS Requestor had a password change or account reset.  Format: **YYYYMMDD**
	ChAccPwChange *string `json:"chAccPwChange,omitempty"`
	// Indicates the length of time since the cardholder’s account with the 3DS Requestor had a password change or account reset.  Allowed values: * **01** — No change * **02** — Changed during this transaction * **03** — Less than 30 days * **04** — 30–60 days * **05** — More than 60 days
	ChAccPwChangeInd *string `json:"chAccPwChangeInd,omitempty"`
	// Date that the cardholder opened the account with the 3DS Requestor.  Format: **YYYYMMDD**
	ChAccString *string `json:"chAccString,omitempty"`
	// Number of purchases with this cardholder account during the previous six months. Max length: 4 characters.
	NbPurchaseAccount *string `json:"nbPurchaseAccount,omitempty"`
	// String that the payment account was enrolled in the cardholder’s account with the 3DS Requestor.  Format: **YYYYMMDD**
	PaymentAccAge *string `json:"paymentAccAge,omitempty"`
	// Indicates the length of time that the payment account was enrolled in the cardholder’s account with the 3DS Requestor.  Allowed values: * **01** — No account (guest checkout) * **02** — During this transaction * **03** — Less than 30 days * **04** — 30–60 days * **05** — More than 60 days
	PaymentAccInd *string `json:"paymentAccInd,omitempty"`
	// Number of Add Card attempts in the last 24 hours. Max length: 3 characters.
	ProvisionAttemptsDay *string `json:"provisionAttemptsDay,omitempty"`
	// String when the shipping address used for this transaction was first used with the 3DS Requestor.  Format: **YYYYMMDD**
	ShipAddressUsage *string `json:"shipAddressUsage,omitempty"`
	// Indicates when the shipping address used for this transaction was first used with the 3DS Requestor.  Allowed values: * **01** — This transaction * **02** — Less than 30 days * **03** — 30–60 days * **04** — More than 60 days
	ShipAddressUsageInd *string `json:"shipAddressUsageInd,omitempty"`
	// Indicates if the Cardholder Name on the account is identical to the shipping Name used for this transaction.  Allowed values: * **01** — Account Name identical to shipping Name * **02** — Account Name different to shipping Name
	ShipNameIndicator *string `json:"shipNameIndicator,omitempty"`
	// Indicates whether the 3DS Requestor has experienced suspicious activity (including previous fraud) on the cardholder account.  Allowed values: * **01** — No suspicious activity has been observed * **02** — Suspicious activity has been observed
	SuspiciousAccActivity *string `json:"suspiciousAccActivity,omitempty"`
	// Number of transactions (successful and abandoned) for this cardholder account with the 3DS Requestor across all payment accounts in the previous 24 hours. Max length: 3 characters.
	TxnActivityDay *string `json:"txnActivityDay,omitempty"`
	// Number of transactions (successful and abandoned) for this cardholder account with the 3DS Requestor across all payment accounts in the previous year. Max length: 3 characters.
	TxnActivityYear *string `json:"txnActivityYear,omitempty"`
}

// NewAcctInfo instantiates a new AcctInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcctInfo() *AcctInfo {
	this := AcctInfo{}
	return &this
}

// NewAcctInfoWithDefaults instantiates a new AcctInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcctInfoWithDefaults() *AcctInfo {
	this := AcctInfo{}
	return &this
}

// GetChAccAgeInd returns the ChAccAgeInd field value if set, zero value otherwise.
func (o *AcctInfo) GetChAccAgeInd() string {
	if o == nil || common.IsNil(o.ChAccAgeInd) {
		var ret string
		return ret
	}
	return *o.ChAccAgeInd
}

// GetChAccAgeIndOk returns a tuple with the ChAccAgeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetChAccAgeIndOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChAccAgeInd) {
		return nil, false
	}
	return o.ChAccAgeInd, true
}

// HasChAccAgeInd returns a boolean if a field has been set.
func (o *AcctInfo) HasChAccAgeInd() bool {
	if o != nil && !common.IsNil(o.ChAccAgeInd) {
		return true
	}

	return false
}

// SetChAccAgeInd gets a reference to the given string and assigns it to the ChAccAgeInd field.
func (o *AcctInfo) SetChAccAgeInd(v string) {
	o.ChAccAgeInd = &v
}

// GetChAccChange returns the ChAccChange field value if set, zero value otherwise.
func (o *AcctInfo) GetChAccChange() string {
	if o == nil || common.IsNil(o.ChAccChange) {
		var ret string
		return ret
	}
	return *o.ChAccChange
}

// GetChAccChangeOk returns a tuple with the ChAccChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetChAccChangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChAccChange) {
		return nil, false
	}
	return o.ChAccChange, true
}

// HasChAccChange returns a boolean if a field has been set.
func (o *AcctInfo) HasChAccChange() bool {
	if o != nil && !common.IsNil(o.ChAccChange) {
		return true
	}

	return false
}

// SetChAccChange gets a reference to the given string and assigns it to the ChAccChange field.
func (o *AcctInfo) SetChAccChange(v string) {
	o.ChAccChange = &v
}

// GetChAccChangeInd returns the ChAccChangeInd field value if set, zero value otherwise.
func (o *AcctInfo) GetChAccChangeInd() string {
	if o == nil || common.IsNil(o.ChAccChangeInd) {
		var ret string
		return ret
	}
	return *o.ChAccChangeInd
}

// GetChAccChangeIndOk returns a tuple with the ChAccChangeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetChAccChangeIndOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChAccChangeInd) {
		return nil, false
	}
	return o.ChAccChangeInd, true
}

// HasChAccChangeInd returns a boolean if a field has been set.
func (o *AcctInfo) HasChAccChangeInd() bool {
	if o != nil && !common.IsNil(o.ChAccChangeInd) {
		return true
	}

	return false
}

// SetChAccChangeInd gets a reference to the given string and assigns it to the ChAccChangeInd field.
func (o *AcctInfo) SetChAccChangeInd(v string) {
	o.ChAccChangeInd = &v
}

// GetChAccPwChange returns the ChAccPwChange field value if set, zero value otherwise.
func (o *AcctInfo) GetChAccPwChange() string {
	if o == nil || common.IsNil(o.ChAccPwChange) {
		var ret string
		return ret
	}
	return *o.ChAccPwChange
}

// GetChAccPwChangeOk returns a tuple with the ChAccPwChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetChAccPwChangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChAccPwChange) {
		return nil, false
	}
	return o.ChAccPwChange, true
}

// HasChAccPwChange returns a boolean if a field has been set.
func (o *AcctInfo) HasChAccPwChange() bool {
	if o != nil && !common.IsNil(o.ChAccPwChange) {
		return true
	}

	return false
}

// SetChAccPwChange gets a reference to the given string and assigns it to the ChAccPwChange field.
func (o *AcctInfo) SetChAccPwChange(v string) {
	o.ChAccPwChange = &v
}

// GetChAccPwChangeInd returns the ChAccPwChangeInd field value if set, zero value otherwise.
func (o *AcctInfo) GetChAccPwChangeInd() string {
	if o == nil || common.IsNil(o.ChAccPwChangeInd) {
		var ret string
		return ret
	}
	return *o.ChAccPwChangeInd
}

// GetChAccPwChangeIndOk returns a tuple with the ChAccPwChangeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetChAccPwChangeIndOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChAccPwChangeInd) {
		return nil, false
	}
	return o.ChAccPwChangeInd, true
}

// HasChAccPwChangeInd returns a boolean if a field has been set.
func (o *AcctInfo) HasChAccPwChangeInd() bool {
	if o != nil && !common.IsNil(o.ChAccPwChangeInd) {
		return true
	}

	return false
}

// SetChAccPwChangeInd gets a reference to the given string and assigns it to the ChAccPwChangeInd field.
func (o *AcctInfo) SetChAccPwChangeInd(v string) {
	o.ChAccPwChangeInd = &v
}

// GetChAccString returns the ChAccString field value if set, zero value otherwise.
func (o *AcctInfo) GetChAccString() string {
	if o == nil || common.IsNil(o.ChAccString) {
		var ret string
		return ret
	}
	return *o.ChAccString
}

// GetChAccStringOk returns a tuple with the ChAccString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetChAccStringOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChAccString) {
		return nil, false
	}
	return o.ChAccString, true
}

// HasChAccString returns a boolean if a field has been set.
func (o *AcctInfo) HasChAccString() bool {
	if o != nil && !common.IsNil(o.ChAccString) {
		return true
	}

	return false
}

// SetChAccString gets a reference to the given string and assigns it to the ChAccString field.
func (o *AcctInfo) SetChAccString(v string) {
	o.ChAccString = &v
}

// GetNbPurchaseAccount returns the NbPurchaseAccount field value if set, zero value otherwise.
func (o *AcctInfo) GetNbPurchaseAccount() string {
	if o == nil || common.IsNil(o.NbPurchaseAccount) {
		var ret string
		return ret
	}
	return *o.NbPurchaseAccount
}

// GetNbPurchaseAccountOk returns a tuple with the NbPurchaseAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetNbPurchaseAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.NbPurchaseAccount) {
		return nil, false
	}
	return o.NbPurchaseAccount, true
}

// HasNbPurchaseAccount returns a boolean if a field has been set.
func (o *AcctInfo) HasNbPurchaseAccount() bool {
	if o != nil && !common.IsNil(o.NbPurchaseAccount) {
		return true
	}

	return false
}

// SetNbPurchaseAccount gets a reference to the given string and assigns it to the NbPurchaseAccount field.
func (o *AcctInfo) SetNbPurchaseAccount(v string) {
	o.NbPurchaseAccount = &v
}

// GetPaymentAccAge returns the PaymentAccAge field value if set, zero value otherwise.
func (o *AcctInfo) GetPaymentAccAge() string {
	if o == nil || common.IsNil(o.PaymentAccAge) {
		var ret string
		return ret
	}
	return *o.PaymentAccAge
}

// GetPaymentAccAgeOk returns a tuple with the PaymentAccAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetPaymentAccAgeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentAccAge) {
		return nil, false
	}
	return o.PaymentAccAge, true
}

// HasPaymentAccAge returns a boolean if a field has been set.
func (o *AcctInfo) HasPaymentAccAge() bool {
	if o != nil && !common.IsNil(o.PaymentAccAge) {
		return true
	}

	return false
}

// SetPaymentAccAge gets a reference to the given string and assigns it to the PaymentAccAge field.
func (o *AcctInfo) SetPaymentAccAge(v string) {
	o.PaymentAccAge = &v
}

// GetPaymentAccInd returns the PaymentAccInd field value if set, zero value otherwise.
func (o *AcctInfo) GetPaymentAccInd() string {
	if o == nil || common.IsNil(o.PaymentAccInd) {
		var ret string
		return ret
	}
	return *o.PaymentAccInd
}

// GetPaymentAccIndOk returns a tuple with the PaymentAccInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetPaymentAccIndOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentAccInd) {
		return nil, false
	}
	return o.PaymentAccInd, true
}

// HasPaymentAccInd returns a boolean if a field has been set.
func (o *AcctInfo) HasPaymentAccInd() bool {
	if o != nil && !common.IsNil(o.PaymentAccInd) {
		return true
	}

	return false
}

// SetPaymentAccInd gets a reference to the given string and assigns it to the PaymentAccInd field.
func (o *AcctInfo) SetPaymentAccInd(v string) {
	o.PaymentAccInd = &v
}

// GetProvisionAttemptsDay returns the ProvisionAttemptsDay field value if set, zero value otherwise.
func (o *AcctInfo) GetProvisionAttemptsDay() string {
	if o == nil || common.IsNil(o.ProvisionAttemptsDay) {
		var ret string
		return ret
	}
	return *o.ProvisionAttemptsDay
}

// GetProvisionAttemptsDayOk returns a tuple with the ProvisionAttemptsDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetProvisionAttemptsDayOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProvisionAttemptsDay) {
		return nil, false
	}
	return o.ProvisionAttemptsDay, true
}

// HasProvisionAttemptsDay returns a boolean if a field has been set.
func (o *AcctInfo) HasProvisionAttemptsDay() bool {
	if o != nil && !common.IsNil(o.ProvisionAttemptsDay) {
		return true
	}

	return false
}

// SetProvisionAttemptsDay gets a reference to the given string and assigns it to the ProvisionAttemptsDay field.
func (o *AcctInfo) SetProvisionAttemptsDay(v string) {
	o.ProvisionAttemptsDay = &v
}

// GetShipAddressUsage returns the ShipAddressUsage field value if set, zero value otherwise.
func (o *AcctInfo) GetShipAddressUsage() string {
	if o == nil || common.IsNil(o.ShipAddressUsage) {
		var ret string
		return ret
	}
	return *o.ShipAddressUsage
}

// GetShipAddressUsageOk returns a tuple with the ShipAddressUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetShipAddressUsageOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShipAddressUsage) {
		return nil, false
	}
	return o.ShipAddressUsage, true
}

// HasShipAddressUsage returns a boolean if a field has been set.
func (o *AcctInfo) HasShipAddressUsage() bool {
	if o != nil && !common.IsNil(o.ShipAddressUsage) {
		return true
	}

	return false
}

// SetShipAddressUsage gets a reference to the given string and assigns it to the ShipAddressUsage field.
func (o *AcctInfo) SetShipAddressUsage(v string) {
	o.ShipAddressUsage = &v
}

// GetShipAddressUsageInd returns the ShipAddressUsageInd field value if set, zero value otherwise.
func (o *AcctInfo) GetShipAddressUsageInd() string {
	if o == nil || common.IsNil(o.ShipAddressUsageInd) {
		var ret string
		return ret
	}
	return *o.ShipAddressUsageInd
}

// GetShipAddressUsageIndOk returns a tuple with the ShipAddressUsageInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetShipAddressUsageIndOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShipAddressUsageInd) {
		return nil, false
	}
	return o.ShipAddressUsageInd, true
}

// HasShipAddressUsageInd returns a boolean if a field has been set.
func (o *AcctInfo) HasShipAddressUsageInd() bool {
	if o != nil && !common.IsNil(o.ShipAddressUsageInd) {
		return true
	}

	return false
}

// SetShipAddressUsageInd gets a reference to the given string and assigns it to the ShipAddressUsageInd field.
func (o *AcctInfo) SetShipAddressUsageInd(v string) {
	o.ShipAddressUsageInd = &v
}

// GetShipNameIndicator returns the ShipNameIndicator field value if set, zero value otherwise.
func (o *AcctInfo) GetShipNameIndicator() string {
	if o == nil || common.IsNil(o.ShipNameIndicator) {
		var ret string
		return ret
	}
	return *o.ShipNameIndicator
}

// GetShipNameIndicatorOk returns a tuple with the ShipNameIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetShipNameIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShipNameIndicator) {
		return nil, false
	}
	return o.ShipNameIndicator, true
}

// HasShipNameIndicator returns a boolean if a field has been set.
func (o *AcctInfo) HasShipNameIndicator() bool {
	if o != nil && !common.IsNil(o.ShipNameIndicator) {
		return true
	}

	return false
}

// SetShipNameIndicator gets a reference to the given string and assigns it to the ShipNameIndicator field.
func (o *AcctInfo) SetShipNameIndicator(v string) {
	o.ShipNameIndicator = &v
}

// GetSuspiciousAccActivity returns the SuspiciousAccActivity field value if set, zero value otherwise.
func (o *AcctInfo) GetSuspiciousAccActivity() string {
	if o == nil || common.IsNil(o.SuspiciousAccActivity) {
		var ret string
		return ret
	}
	return *o.SuspiciousAccActivity
}

// GetSuspiciousAccActivityOk returns a tuple with the SuspiciousAccActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetSuspiciousAccActivityOk() (*string, bool) {
	if o == nil || common.IsNil(o.SuspiciousAccActivity) {
		return nil, false
	}
	return o.SuspiciousAccActivity, true
}

// HasSuspiciousAccActivity returns a boolean if a field has been set.
func (o *AcctInfo) HasSuspiciousAccActivity() bool {
	if o != nil && !common.IsNil(o.SuspiciousAccActivity) {
		return true
	}

	return false
}

// SetSuspiciousAccActivity gets a reference to the given string and assigns it to the SuspiciousAccActivity field.
func (o *AcctInfo) SetSuspiciousAccActivity(v string) {
	o.SuspiciousAccActivity = &v
}

// GetTxnActivityDay returns the TxnActivityDay field value if set, zero value otherwise.
func (o *AcctInfo) GetTxnActivityDay() string {
	if o == nil || common.IsNil(o.TxnActivityDay) {
		var ret string
		return ret
	}
	return *o.TxnActivityDay
}

// GetTxnActivityDayOk returns a tuple with the TxnActivityDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetTxnActivityDayOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxnActivityDay) {
		return nil, false
	}
	return o.TxnActivityDay, true
}

// HasTxnActivityDay returns a boolean if a field has been set.
func (o *AcctInfo) HasTxnActivityDay() bool {
	if o != nil && !common.IsNil(o.TxnActivityDay) {
		return true
	}

	return false
}

// SetTxnActivityDay gets a reference to the given string and assigns it to the TxnActivityDay field.
func (o *AcctInfo) SetTxnActivityDay(v string) {
	o.TxnActivityDay = &v
}

// GetTxnActivityYear returns the TxnActivityYear field value if set, zero value otherwise.
func (o *AcctInfo) GetTxnActivityYear() string {
	if o == nil || common.IsNil(o.TxnActivityYear) {
		var ret string
		return ret
	}
	return *o.TxnActivityYear
}

// GetTxnActivityYearOk returns a tuple with the TxnActivityYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcctInfo) GetTxnActivityYearOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxnActivityYear) {
		return nil, false
	}
	return o.TxnActivityYear, true
}

// HasTxnActivityYear returns a boolean if a field has been set.
func (o *AcctInfo) HasTxnActivityYear() bool {
	if o != nil && !common.IsNil(o.TxnActivityYear) {
		return true
	}

	return false
}

// SetTxnActivityYear gets a reference to the given string and assigns it to the TxnActivityYear field.
func (o *AcctInfo) SetTxnActivityYear(v string) {
	o.TxnActivityYear = &v
}

func (o AcctInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcctInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ChAccAgeInd) {
		toSerialize["chAccAgeInd"] = o.ChAccAgeInd
	}
	if !common.IsNil(o.ChAccChange) {
		toSerialize["chAccChange"] = o.ChAccChange
	}
	if !common.IsNil(o.ChAccChangeInd) {
		toSerialize["chAccChangeInd"] = o.ChAccChangeInd
	}
	if !common.IsNil(o.ChAccPwChange) {
		toSerialize["chAccPwChange"] = o.ChAccPwChange
	}
	if !common.IsNil(o.ChAccPwChangeInd) {
		toSerialize["chAccPwChangeInd"] = o.ChAccPwChangeInd
	}
	if !common.IsNil(o.ChAccString) {
		toSerialize["chAccString"] = o.ChAccString
	}
	if !common.IsNil(o.NbPurchaseAccount) {
		toSerialize["nbPurchaseAccount"] = o.NbPurchaseAccount
	}
	if !common.IsNil(o.PaymentAccAge) {
		toSerialize["paymentAccAge"] = o.PaymentAccAge
	}
	if !common.IsNil(o.PaymentAccInd) {
		toSerialize["paymentAccInd"] = o.PaymentAccInd
	}
	if !common.IsNil(o.ProvisionAttemptsDay) {
		toSerialize["provisionAttemptsDay"] = o.ProvisionAttemptsDay
	}
	if !common.IsNil(o.ShipAddressUsage) {
		toSerialize["shipAddressUsage"] = o.ShipAddressUsage
	}
	if !common.IsNil(o.ShipAddressUsageInd) {
		toSerialize["shipAddressUsageInd"] = o.ShipAddressUsageInd
	}
	if !common.IsNil(o.ShipNameIndicator) {
		toSerialize["shipNameIndicator"] = o.ShipNameIndicator
	}
	if !common.IsNil(o.SuspiciousAccActivity) {
		toSerialize["suspiciousAccActivity"] = o.SuspiciousAccActivity
	}
	if !common.IsNil(o.TxnActivityDay) {
		toSerialize["txnActivityDay"] = o.TxnActivityDay
	}
	if !common.IsNil(o.TxnActivityYear) {
		toSerialize["txnActivityYear"] = o.TxnActivityYear
	}
	return toSerialize, nil
}

type NullableAcctInfo struct {
	value *AcctInfo
	isSet bool
}

func (v NullableAcctInfo) Get() *AcctInfo {
	return v.value
}

func (v *NullableAcctInfo) Set(val *AcctInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAcctInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAcctInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcctInfo(val *AcctInfo) *NullableAcctInfo {
	return &NullableAcctInfo{value: val, isSet: true}
}

func (v NullableAcctInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcctInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AcctInfo) isValidChAccAgeInd() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04", "05"}
	for _, allowed := range allowedEnumValues {
		if o.GetChAccAgeInd() == allowed {
			return true
		}
	}
	return false
}
func (o *AcctInfo) isValidChAccChangeInd() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04"}
	for _, allowed := range allowedEnumValues {
		if o.GetChAccChangeInd() == allowed {
			return true
		}
	}
	return false
}
func (o *AcctInfo) isValidChAccPwChangeInd() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04", "05"}
	for _, allowed := range allowedEnumValues {
		if o.GetChAccPwChangeInd() == allowed {
			return true
		}
	}
	return false
}
func (o *AcctInfo) isValidPaymentAccInd() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04", "05"}
	for _, allowed := range allowedEnumValues {
		if o.GetPaymentAccInd() == allowed {
			return true
		}
	}
	return false
}
func (o *AcctInfo) isValidShipAddressUsageInd() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04"}
	for _, allowed := range allowedEnumValues {
		if o.GetShipAddressUsageInd() == allowed {
			return true
		}
	}
	return false
}
func (o *AcctInfo) isValidShipNameIndicator() bool {
	var allowedEnumValues = []string{"01", "02"}
	for _, allowed := range allowedEnumValues {
		if o.GetShipNameIndicator() == allowed {
			return true
		}
	}
	return false
}
func (o *AcctInfo) isValidSuspiciousAccActivity() bool {
	var allowedEnumValues = []string{"01", "02"}
	for _, allowed := range allowedEnumValues {
		if o.GetSuspiciousAccActivity() == allowed {
			return true
		}
	}
	return false
}
