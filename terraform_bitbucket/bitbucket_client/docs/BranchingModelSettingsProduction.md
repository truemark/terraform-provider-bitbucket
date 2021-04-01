# BranchingModelSettingsProduction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsValid** | Pointer to **bool** | Indicates if the configured branch is valid, that is, if the configured branch actually exists currently. Is always &#x60;true&#x60; when &#x60;use_mainbranch&#x60; is &#x60;true&#x60; (even if the main branch does not exist). This field is read-only. This field is ignored when updating/creating settings. | [optional] 
**Name** | Pointer to **string** | The configured branch. It must be &#x60;null&#x60; when &#x60;use_mainbranch&#x60; is &#x60;true&#x60;. Otherwise it must be a non-empty value. It is possible for the configured branch to not exist (e.g. it was deleted after the settings are set). In this case &#x60;is_valid&#x60; will be &#x60;false&#x60;. The branch must exist when updating/setting the &#x60;name&#x60; or an error will occur. | [optional] 
**UseMainbranch** | Pointer to **bool** | Indicates if the setting points at an explicit branch (&#x60;false&#x60;) or tracks the main branch (&#x60;true&#x60;). When &#x60;true&#x60; the &#x60;name&#x60; must be &#x60;null&#x60; or not provided. When &#x60;false&#x60; the &#x60;name&#x60; must contain a non-empty branch name. | [optional] 
**Enabled** | Pointer to **bool** | Indicates if branch is enabled or not. | [optional] 

## Methods

### NewBranchingModelSettingsProduction

`func NewBranchingModelSettingsProduction() *BranchingModelSettingsProduction`

NewBranchingModelSettingsProduction instantiates a new BranchingModelSettingsProduction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchingModelSettingsProductionWithDefaults

`func NewBranchingModelSettingsProductionWithDefaults() *BranchingModelSettingsProduction`

NewBranchingModelSettingsProductionWithDefaults instantiates a new BranchingModelSettingsProduction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsValid

`func (o *BranchingModelSettingsProduction) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *BranchingModelSettingsProduction) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *BranchingModelSettingsProduction) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *BranchingModelSettingsProduction) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetName

`func (o *BranchingModelSettingsProduction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BranchingModelSettingsProduction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BranchingModelSettingsProduction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BranchingModelSettingsProduction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUseMainbranch

`func (o *BranchingModelSettingsProduction) GetUseMainbranch() bool`

GetUseMainbranch returns the UseMainbranch field if non-nil, zero value otherwise.

### GetUseMainbranchOk

`func (o *BranchingModelSettingsProduction) GetUseMainbranchOk() (*bool, bool)`

GetUseMainbranchOk returns a tuple with the UseMainbranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMainbranch

`func (o *BranchingModelSettingsProduction) SetUseMainbranch(v bool)`

SetUseMainbranch sets UseMainbranch field to given value.

### HasUseMainbranch

`func (o *BranchingModelSettingsProduction) HasUseMainbranch() bool`

HasUseMainbranch returns a boolean if a field has been set.

### GetEnabled

`func (o *BranchingModelSettingsProduction) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BranchingModelSettingsProduction) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BranchingModelSettingsProduction) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BranchingModelSettingsProduction) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


