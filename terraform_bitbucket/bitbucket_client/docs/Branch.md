# Branch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Links** | Pointer to [**RefLinks**](RefLinks.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the ref. | [optional] 
**Target** | Pointer to [**Commit**](Commit.md) |  | [optional] 

## Methods

### NewBranch

`func NewBranch(type_ string, ) *Branch`

NewBranch instantiates a new Branch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchWithDefaults

`func NewBranchWithDefaults() *Branch`

NewBranchWithDefaults instantiates a new Branch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Branch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Branch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Branch) SetType(v string)`

SetType sets Type field to given value.


### GetLinks

`func (o *Branch) GetLinks() RefLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Branch) GetLinksOk() (*RefLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Branch) SetLinks(v RefLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Branch) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *Branch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Branch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Branch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Branch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTarget

`func (o *Branch) GetTarget() Commit`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Branch) GetTargetOk() (*Commit, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Branch) SetTarget(v Commit)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Branch) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


