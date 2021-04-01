# BranchingModelDevelopment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to [**Branch**](Branch.md) |  | [optional] 
**Name** | **string** | Name of the target branch. Will be listed here even when the target branch does not exist. Will be &#x60;null&#x60; if targeting the main branch and the repository is empty. | 
**UseMainbranch** | **bool** | Indicates if the setting points at an explicit branch (&#x60;false&#x60;) or tracks the main branch (&#x60;true&#x60;). | 

## Methods

### NewBranchingModelDevelopment

`func NewBranchingModelDevelopment(name string, useMainbranch bool, ) *BranchingModelDevelopment`

NewBranchingModelDevelopment instantiates a new BranchingModelDevelopment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchingModelDevelopmentWithDefaults

`func NewBranchingModelDevelopmentWithDefaults() *BranchingModelDevelopment`

NewBranchingModelDevelopmentWithDefaults instantiates a new BranchingModelDevelopment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *BranchingModelDevelopment) GetBranch() Branch`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *BranchingModelDevelopment) GetBranchOk() (*Branch, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *BranchingModelDevelopment) SetBranch(v Branch)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *BranchingModelDevelopment) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetName

`func (o *BranchingModelDevelopment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BranchingModelDevelopment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BranchingModelDevelopment) SetName(v string)`

SetName sets Name field to given value.


### GetUseMainbranch

`func (o *BranchingModelDevelopment) GetUseMainbranch() bool`

GetUseMainbranch returns the UseMainbranch field if non-nil, zero value otherwise.

### GetUseMainbranchOk

`func (o *BranchingModelDevelopment) GetUseMainbranchOk() (*bool, bool)`

GetUseMainbranchOk returns a tuple with the UseMainbranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMainbranch

`func (o *BranchingModelDevelopment) SetUseMainbranch(v bool)`

SetUseMainbranch sets UseMainbranch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


