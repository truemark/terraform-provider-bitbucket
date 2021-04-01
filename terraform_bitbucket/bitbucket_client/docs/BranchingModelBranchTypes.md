# BranchingModelBranchTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind of branch. | 
**Prefix** | **string** | The prefix for this branch type. A branch with this prefix will be classified as per &#x60;kind&#x60;. The prefix must be a valid prefix for a branch and must always exist. It cannot be blank, empty or &#x60;null&#x60;. | 

## Methods

### NewBranchingModelBranchTypes

`func NewBranchingModelBranchTypes(kind string, prefix string, ) *BranchingModelBranchTypes`

NewBranchingModelBranchTypes instantiates a new BranchingModelBranchTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchingModelBranchTypesWithDefaults

`func NewBranchingModelBranchTypesWithDefaults() *BranchingModelBranchTypes`

NewBranchingModelBranchTypesWithDefaults instantiates a new BranchingModelBranchTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *BranchingModelBranchTypes) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BranchingModelBranchTypes) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BranchingModelBranchTypes) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrefix

`func (o *BranchingModelBranchTypes) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *BranchingModelBranchTypes) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *BranchingModelBranchTypes) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


