# PipelineStepError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The error key. | [optional] 
**Message** | Pointer to **string** | The error message. | [optional] 

## Methods

### NewPipelineStepError

`func NewPipelineStepError() *PipelineStepError`

NewPipelineStepError instantiates a new PipelineStepError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStepErrorWithDefaults

`func NewPipelineStepErrorWithDefaults() *PipelineStepError`

NewPipelineStepErrorWithDefaults instantiates a new PipelineStepError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PipelineStepError) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PipelineStepError) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PipelineStepError) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PipelineStepError) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMessage

`func (o *PipelineStepError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PipelineStepError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PipelineStepError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PipelineStepError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


