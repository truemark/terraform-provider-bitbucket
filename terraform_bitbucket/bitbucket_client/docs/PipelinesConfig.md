# PipelinesConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether Pipelines is enabled for the repository. | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 

## Methods

### NewPipelinesConfig

`func NewPipelinesConfig() *PipelinesConfig`

NewPipelinesConfig instantiates a new PipelinesConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelinesConfigWithDefaults

`func NewPipelinesConfigWithDefaults() *PipelinesConfig`

NewPipelinesConfigWithDefaults instantiates a new PipelinesConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PipelinesConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PipelinesConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PipelinesConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PipelinesConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRepository

`func (o *PipelinesConfig) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PipelinesConfig) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PipelinesConfig) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PipelinesConfig) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


