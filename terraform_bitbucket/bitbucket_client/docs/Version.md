# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**BranchingModelSettingsLinks**](BranchingModelSettingsLinks.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 

## Methods

### NewVersion

`func NewVersion() *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Version) GetLinks() BranchingModelSettingsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Version) GetLinksOk() (*BranchingModelSettingsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Version) SetLinks(v BranchingModelSettingsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Version) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *Version) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Version) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Version) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Version) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *Version) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Version) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Version) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Version) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


