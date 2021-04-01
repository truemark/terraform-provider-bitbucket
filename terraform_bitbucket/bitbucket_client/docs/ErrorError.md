# ErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Detail** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** | Optional structured data that is endpoint-specific. | [optional] 

## Methods

### NewErrorError

`func NewErrorError(message string, ) *ErrorError`

NewErrorError instantiates a new ErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorErrorWithDefaults

`func NewErrorErrorWithDefaults() *ErrorError`

NewErrorErrorWithDefaults instantiates a new ErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetail

`func (o *ErrorError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ErrorError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetData

`func (o *ErrorError) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ErrorError) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ErrorError) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ErrorError) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


