# PipelineRefTargetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | Pointer to **string** | The type of reference (branch/tag). | [optional] 
**RefName** | Pointer to **string** | The name of the reference. | [optional] 
**Commit** | Pointer to [**Commit**](Commit.md) |  | [optional] 
**Selector** | Pointer to [**PipelineSelector**](PipelineSelector.md) |  | [optional] 

## Methods

### NewPipelineRefTargetAllOf

`func NewPipelineRefTargetAllOf() *PipelineRefTargetAllOf`

NewPipelineRefTargetAllOf instantiates a new PipelineRefTargetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineRefTargetAllOfWithDefaults

`func NewPipelineRefTargetAllOfWithDefaults() *PipelineRefTargetAllOf`

NewPipelineRefTargetAllOfWithDefaults instantiates a new PipelineRefTargetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *PipelineRefTargetAllOf) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *PipelineRefTargetAllOf) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *PipelineRefTargetAllOf) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *PipelineRefTargetAllOf) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefName

`func (o *PipelineRefTargetAllOf) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *PipelineRefTargetAllOf) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *PipelineRefTargetAllOf) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *PipelineRefTargetAllOf) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetCommit

`func (o *PipelineRefTargetAllOf) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *PipelineRefTargetAllOf) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *PipelineRefTargetAllOf) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *PipelineRefTargetAllOf) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetSelector

`func (o *PipelineRefTargetAllOf) GetSelector() PipelineSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *PipelineRefTargetAllOf) GetSelectorOk() (*PipelineSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *PipelineRefTargetAllOf) SetSelector(v PipelineSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *PipelineRefTargetAllOf) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


