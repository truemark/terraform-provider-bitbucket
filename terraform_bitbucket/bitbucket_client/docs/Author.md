# Author

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Raw** | Pointer to **string** | The raw author value from the repository. This may be the only value available if the author does not match a user in Bitbucket. | [optional] 
**User** | Pointer to [**Account**](Account.md) |  | [optional] 

## Methods

### NewAuthor

`func NewAuthor() *Author`

NewAuthor instantiates a new Author object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorWithDefaults

`func NewAuthorWithDefaults() *Author`

NewAuthorWithDefaults instantiates a new Author object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaw

`func (o *Author) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *Author) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *Author) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *Author) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetUser

`func (o *Author) GetUser() Account`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Author) GetUserOk() (*Account, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Author) SetUser(v Account)`

SetUser sets User field to given value.

### HasUser

`func (o *Author) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


