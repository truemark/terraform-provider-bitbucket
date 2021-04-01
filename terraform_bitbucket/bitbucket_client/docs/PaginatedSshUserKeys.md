# PaginatedSshUserKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** | Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute. | [optional] 
**Page** | Pointer to **int32** | Page number of the current results. This is an optional element that is not provided in all responses. | [optional] 
**Pagelen** | Pointer to **int32** | Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values. | [optional] 
**Next** | Pointer to **string** | Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs. | [optional] 
**Previous** | Pointer to **string** | Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs. | [optional] 
**Values** | Pointer to [**[]SshAccountKey**](SshAccountKey.md) |  | [optional] 

## Methods

### NewPaginatedSshUserKeys

`func NewPaginatedSshUserKeys() *PaginatedSshUserKeys`

NewPaginatedSshUserKeys instantiates a new PaginatedSshUserKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedSshUserKeysWithDefaults

`func NewPaginatedSshUserKeysWithDefaults() *PaginatedSshUserKeys`

NewPaginatedSshUserKeysWithDefaults instantiates a new PaginatedSshUserKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *PaginatedSshUserKeys) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PaginatedSshUserKeys) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PaginatedSshUserKeys) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PaginatedSshUserKeys) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPage

`func (o *PaginatedSshUserKeys) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginatedSshUserKeys) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginatedSshUserKeys) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaginatedSshUserKeys) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPagelen

`func (o *PaginatedSshUserKeys) GetPagelen() int32`

GetPagelen returns the Pagelen field if non-nil, zero value otherwise.

### GetPagelenOk

`func (o *PaginatedSshUserKeys) GetPagelenOk() (*int32, bool)`

GetPagelenOk returns a tuple with the Pagelen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagelen

`func (o *PaginatedSshUserKeys) SetPagelen(v int32)`

SetPagelen sets Pagelen field to given value.

### HasPagelen

`func (o *PaginatedSshUserKeys) HasPagelen() bool`

HasPagelen returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedSshUserKeys) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedSshUserKeys) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedSshUserKeys) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedSshUserKeys) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedSshUserKeys) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedSshUserKeys) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedSshUserKeys) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedSshUserKeys) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetValues

`func (o *PaginatedSshUserKeys) GetValues() []SshAccountKey`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PaginatedSshUserKeys) GetValuesOk() (*[]SshAccountKey, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PaginatedSshUserKeys) SetValues(v []SshAccountKey)`

SetValues sets Values field to given value.

### HasValues

`func (o *PaginatedSshUserKeys) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


