# RepositoryPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Permission** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 

## Methods

### NewRepositoryPermission

`func NewRepositoryPermission(type_ string, ) *RepositoryPermission`

NewRepositoryPermission instantiates a new RepositoryPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryPermissionWithDefaults

`func NewRepositoryPermissionWithDefaults() *RepositoryPermission`

NewRepositoryPermissionWithDefaults instantiates a new RepositoryPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RepositoryPermission) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepositoryPermission) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepositoryPermission) SetType(v string)`

SetType sets Type field to given value.


### GetPermission

`func (o *RepositoryPermission) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *RepositoryPermission) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *RepositoryPermission) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *RepositoryPermission) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUser

`func (o *RepositoryPermission) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RepositoryPermission) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RepositoryPermission) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *RepositoryPermission) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRepository

`func (o *RepositoryPermission) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RepositoryPermission) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RepositoryPermission) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RepositoryPermission) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


