# SshKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The SSH key&#39;s immutable ID. | [optional] 
**Key** | Pointer to **string** | The SSH public key value in OpenSSH format. | [optional] 
**Comment** | Pointer to **string** | The comment parsed from the SSH key (if present) | [optional] 
**Label** | Pointer to **string** | The user-defined label for the SSH key | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**LastUsed** | Pointer to **time.Time** |  | [optional] 
**Links** | Pointer to [**BranchingModelSettingsLinks**](BranchingModelSettingsLinks.md) |  | [optional] 

## Methods

### NewSshKey

`func NewSshKey() *SshKey`

NewSshKey instantiates a new SshKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyWithDefaults

`func NewSshKeyWithDefaults() *SshKey`

NewSshKeyWithDefaults instantiates a new SshKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SshKey) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SshKey) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SshKey) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SshKey) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetKey

`func (o *SshKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SshKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SshKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SshKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetComment

`func (o *SshKey) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SshKey) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SshKey) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SshKey) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLabel

`func (o *SshKey) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SshKey) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SshKey) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SshKey) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCreatedOn

`func (o *SshKey) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *SshKey) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *SshKey) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *SshKey) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetLastUsed

`func (o *SshKey) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *SshKey) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *SshKey) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *SshKey) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetLinks

`func (o *SshKey) GetLinks() BranchingModelSettingsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SshKey) GetLinksOk() (*BranchingModelSettingsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SshKey) SetLinks(v BranchingModelSettingsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SshKey) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


