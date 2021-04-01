# ModelError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Error** | Pointer to [**ErrorError**](ErrorError.md) |  | [optional] 

## Methods

### NewModelError

`func NewModelError(type_ string, ) *ModelError`

NewModelError instantiates a new ModelError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelErrorWithDefaults

`func NewModelErrorWithDefaults() *ModelError`

NewModelErrorWithDefaults instantiates a new ModelError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModelError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelError) SetType(v string)`

SetType sets Type field to given value.


### GetError

`func (o *ModelError) GetError() ErrorError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ModelError) GetErrorOk() (*ErrorError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ModelError) SetError(v ErrorError)`

SetError sets Error field to given value.

### HasError

`func (o *ModelError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


