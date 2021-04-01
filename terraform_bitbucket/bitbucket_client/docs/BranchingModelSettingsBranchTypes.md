# BranchingModelSettingsBranchTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the branch type is enabled or not. A disabled branch type may contain an invalid &#x60;prefix&#x60;. | [optional] 
**Kind** | **string** | The kind of the branch type. | 
**Prefix** | Pointer to **string** | The prefix for this branch type. A branch with this prefix will be classified as per &#x60;kind&#x60;. The &#x60;prefix&#x60; of an enabled branch type must be a valid branch prefix.Additionally, it cannot be blank, empty or &#x60;null&#x60;. The &#x60;prefix&#x60; for a disabled branch type can be empty or invalid. | [optional] 

## Methods

### NewBranchingModelSettingsBranchTypes

`func NewBranchingModelSettingsBranchTypes(kind string, ) *BranchingModelSettingsBranchTypes`

NewBranchingModelSettingsBranchTypes instantiates a new BranchingModelSettingsBranchTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchingModelSettingsBranchTypesWithDefaults

`func NewBranchingModelSettingsBranchTypesWithDefaults() *BranchingModelSettingsBranchTypes`

NewBranchingModelSettingsBranchTypesWithDefaults instantiates a new BranchingModelSettingsBranchTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BranchingModelSettingsBranchTypes) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BranchingModelSettingsBranchTypes) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BranchingModelSettingsBranchTypes) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BranchingModelSettingsBranchTypes) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetKind

`func (o *BranchingModelSettingsBranchTypes) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BranchingModelSettingsBranchTypes) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BranchingModelSettingsBranchTypes) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrefix

`func (o *BranchingModelSettingsBranchTypes) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *BranchingModelSettingsBranchTypes) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *BranchingModelSettingsBranchTypes) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *BranchingModelSettingsBranchTypes) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


