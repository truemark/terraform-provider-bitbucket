# PipelineRefTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | Pointer to **string** | The type of reference (branch/tag). | [optional] 
**RefName** | Pointer to **string** | The name of the reference. | [optional] 
**Commit** | Pointer to [**Commit**](Commit.md) |  | [optional] 
**Selector** | Pointer to [**PipelineSelector**](PipelineSelector.md) |  | [optional] 

## Methods

### NewPipelineRefTarget

`func NewPipelineRefTarget() *PipelineRefTarget`

NewPipelineRefTarget instantiates a new PipelineRefTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineRefTargetWithDefaults

`func NewPipelineRefTargetWithDefaults() *PipelineRefTarget`

NewPipelineRefTargetWithDefaults instantiates a new PipelineRefTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *PipelineRefTarget) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *PipelineRefTarget) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *PipelineRefTarget) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *PipelineRefTarget) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefName

`func (o *PipelineRefTarget) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *PipelineRefTarget) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *PipelineRefTarget) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *PipelineRefTarget) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetCommit

`func (o *PipelineRefTarget) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *PipelineRefTarget) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *PipelineRefTarget) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *PipelineRefTarget) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetSelector

`func (o *PipelineRefTarget) GetSelector() PipelineSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *PipelineRefTarget) GetSelectorOk() (*PipelineSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *PipelineRefTarget) SetSelector(v PipelineSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *PipelineRefTarget) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


