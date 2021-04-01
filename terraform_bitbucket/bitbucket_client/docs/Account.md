# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AccountLinks**](AccountLinks.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** | Account name defined by the owner. Should be used instead of the \&quot;username\&quot; field. Note that \&quot;nickname\&quot; cannot be used in place of \&quot;username\&quot; in URLs and queries, as \&quot;nickname\&quot; is not guaranteed to be unique. | [optional] 
**AccountStatus** | Pointer to **string** | The status of the account. Currently the only possible value is \&quot;active\&quot;, but more values may be added in the future. | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Has2faEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Account) GetLinks() AccountLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Account) GetLinksOk() (*AccountLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Account) SetLinks(v AccountLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Account) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetUsername

`func (o *Account) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Account) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Account) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Account) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetNickname

`func (o *Account) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *Account) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *Account) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *Account) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetAccountStatus

`func (o *Account) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *Account) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *Account) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *Account) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetDisplayName

`func (o *Account) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Account) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Account) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Account) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetWebsite

`func (o *Account) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Account) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Account) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Account) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Account) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Account) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Account) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Account) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUuid

`func (o *Account) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Account) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Account) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Account) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetHas2faEnabled

`func (o *Account) GetHas2faEnabled() bool`

GetHas2faEnabled returns the Has2faEnabled field if non-nil, zero value otherwise.

### GetHas2faEnabledOk

`func (o *Account) GetHas2faEnabledOk() (*bool, bool)`

GetHas2faEnabledOk returns a tuple with the Has2faEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHas2faEnabled

`func (o *Account) SetHas2faEnabled(v bool)`

SetHas2faEnabled sets Has2faEnabled field to given value.

### HasHas2faEnabled

`func (o *Account) HasHas2faEnabled() bool`

HasHas2faEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


