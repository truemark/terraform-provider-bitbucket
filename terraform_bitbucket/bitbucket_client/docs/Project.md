# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ProjectLinks**](ProjectLinks.md) |  | [optional] 
**Uuid** | Pointer to **string** | The project&#39;s immutable id. | [optional] 
**Key** | Pointer to **string** | The project&#39;s key. | [optional] 
**Owner** | Pointer to [**Team**](Team.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the project. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsPrivate** | Pointer to **bool** |  Indicates whether the project is publicly accessible, or whether it is private to the team and consequently only visible to team members. Note that private projects cannot contain public repositories. | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**UpdatedOn** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProject

`func NewProject() *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Project) GetLinks() ProjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Project) GetLinksOk() (*ProjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Project) SetLinks(v ProjectLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Project) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetUuid

`func (o *Project) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Project) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Project) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Project) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetKey

`func (o *Project) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Project) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Project) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Project) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOwner

`func (o *Project) GetOwner() Team`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Project) GetOwnerOk() (*Team, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Project) SetOwner(v Team)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Project) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Project) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Project) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPrivate

`func (o *Project) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *Project) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *Project) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *Project) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Project) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Project) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Project) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Project) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *Project) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *Project) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *Project) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *Project) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


