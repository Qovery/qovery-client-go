# BillingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Email** | **string** | email used for billing, and to receive all invoices by email | 
**Address** | **string** |  | 
**City** | **string** |  | 
**Zip** | **string** |  | 
**State** | Pointer to **string** | only for US | [optional] 
**CountryCode** | **string** | ISO code of the country | 
**Company** | Pointer to **string** | name of the company to bill | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingInfo

`func NewBillingInfo(firstName string, lastName string, email string, address string, city string, zip string, countryCode string, ) *BillingInfo`

NewBillingInfo instantiates a new BillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInfoWithDefaults

`func NewBillingInfoWithDefaults() *BillingInfo`

NewBillingInfoWithDefaults instantiates a new BillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *BillingInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BillingInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BillingInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *BillingInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BillingInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BillingInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *BillingInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillingInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillingInfo) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAddress

`func (o *BillingInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BillingInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BillingInfo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCity

`func (o *BillingInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BillingInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BillingInfo) SetCity(v string)`

SetCity sets City field to given value.


### GetZip

`func (o *BillingInfo) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *BillingInfo) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *BillingInfo) SetZip(v string)`

SetZip sets Zip field to given value.


### GetState

`func (o *BillingInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BillingInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BillingInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BillingInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountryCode

`func (o *BillingInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BillingInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BillingInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCompany

`func (o *BillingInfo) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BillingInfo) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BillingInfo) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BillingInfo) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetVatNumber

`func (o *BillingInfo) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *BillingInfo) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *BillingInfo) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *BillingInfo) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


