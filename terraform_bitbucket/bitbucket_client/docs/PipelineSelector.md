# PipelineSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of selector. | 
**Pattern** | Pointer to **string** | The name of the matching pipeline definition. | [optional] 

## Methods

### NewPipelineSelector

`func NewPipelineSelector(type_ string, ) *PipelineSelector`

NewPipelineSelector instantiates a new PipelineSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineSelectorWithDefaults

`func NewPipelineSelectorWithDefaults() *PipelineSelector`

NewPipelineSelectorWithDefaults instantiates a new PipelineSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PipelineSelector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PipelineSelector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PipelineSelector) SetType(v string)`

SetType sets Type field to given value.


### GetPattern

`func (o *PipelineSelector) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *PipelineSelector) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *PipelineSelector) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *PipelineSelector) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


