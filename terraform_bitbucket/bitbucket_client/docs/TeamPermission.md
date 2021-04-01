# TeamPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Permission** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Team** | Pointer to [**Team**](Team.md) |  | [optional] 

## Methods

### NewTeamPermission

`func NewTeamPermission(type_ string, ) *TeamPermission`

NewTeamPermission instantiates a new TeamPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPermissionWithDefaults

`func NewTeamPermissionWithDefaults() *TeamPermission`

NewTeamPermissionWithDefaults instantiates a new TeamPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TeamPermission) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamPermission) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamPermission) SetType(v string)`

SetType sets Type field to given value.


### GetPermission

`func (o *TeamPermission) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *TeamPermission) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *TeamPermission) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *TeamPermission) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUser

`func (o *TeamPermission) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TeamPermission) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TeamPermission) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *TeamPermission) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTeam

`func (o *TeamPermission) GetTeam() Team`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *TeamPermission) GetTeamOk() (*Team, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *TeamPermission) SetTeam(v Team)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *TeamPermission) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


