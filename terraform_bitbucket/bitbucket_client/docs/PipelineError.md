# PipelineError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The error key. | [optional] 
**Message** | Pointer to **string** | The error message. | [optional] 

## Methods

### NewPipelineError

`func NewPipelineError() *PipelineError`

NewPipelineError instantiates a new PipelineError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineErrorWithDefaults

`func NewPipelineErrorWithDefaults() *PipelineError`

NewPipelineErrorWithDefaults instantiates a new PipelineError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PipelineError) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PipelineError) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PipelineError) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PipelineError) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMessage

`func (o *PipelineError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PipelineError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PipelineError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PipelineError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


