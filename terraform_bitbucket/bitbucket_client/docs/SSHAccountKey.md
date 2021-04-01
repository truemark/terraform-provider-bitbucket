# SshAccountKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to [**Account**](Account.md) |  | [optional] 

## Methods

### NewSshAccountKey

`func NewSshAccountKey() *SshAccountKey`

NewSshAccountKey instantiates a new SshAccountKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshAccountKeyWithDefaults

`func NewSshAccountKeyWithDefaults() *SshAccountKey`

NewSshAccountKeyWithDefaults instantiates a new SshAccountKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *SshAccountKey) GetOwner() Account`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SshAccountKey) GetOwnerOk() (*Account, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SshAccountKey) SetOwner(v Account)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SshAccountKey) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


