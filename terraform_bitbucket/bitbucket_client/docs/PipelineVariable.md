# PipelineVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID identifying the variable. | [optional] 
**Key** | Pointer to **string** | The unique name of the variable. | [optional] 
**Value** | Pointer to **string** | The value of the variable. If the variable is secured, this will be empty. | [optional] 
**Secured** | Pointer to **bool** | If true, this variable will be treated as secured. The value will never be exposed in the logs or the REST API. | [optional] 

## Methods

### NewPipelineVariable

`func NewPipelineVariable() *PipelineVariable`

NewPipelineVariable instantiates a new PipelineVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineVariableWithDefaults

`func NewPipelineVariableWithDefaults() *PipelineVariable`

NewPipelineVariableWithDefaults instantiates a new PipelineVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PipelineVariable) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PipelineVariable) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PipelineVariable) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PipelineVariable) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetKey

`func (o *PipelineVariable) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PipelineVariable) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PipelineVariable) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PipelineVariable) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *PipelineVariable) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PipelineVariable) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PipelineVariable) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PipelineVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSecured

`func (o *PipelineVariable) GetSecured() bool`

GetSecured returns the Secured field if non-nil, zero value otherwise.

### GetSecuredOk

`func (o *PipelineVariable) GetSecuredOk() (*bool, bool)`

GetSecuredOk returns a tuple with the Secured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecured

`func (o *PipelineVariable) SetSecured(v bool)`

SetSecured sets Secured field to given value.

### HasSecured

`func (o *PipelineVariable) HasSecured() bool`

HasSecured returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


