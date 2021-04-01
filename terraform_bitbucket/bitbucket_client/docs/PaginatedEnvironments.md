# PaginatedEnvironments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | Page number of the current results. This is an optional element that is not provided in all responses. | [optional] 
**Values** | Pointer to [**[]DeploymentEnvironment**](DeploymentEnvironment.md) | The values of the current page. | [optional] 
**Size** | Pointer to **int32** | Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute. | [optional] 
**Pagelen** | Pointer to **int32** | Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values. | [optional] 
**Next** | Pointer to **string** | Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs. | [optional] 
**Previous** | Pointer to **string** | Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs. | [optional] 

## Methods

### NewPaginatedEnvironments

`func NewPaginatedEnvironments() *PaginatedEnvironments`

NewPaginatedEnvironments instantiates a new PaginatedEnvironments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedEnvironmentsWithDefaults

`func NewPaginatedEnvironmentsWithDefaults() *PaginatedEnvironments`

NewPaginatedEnvironmentsWithDefaults instantiates a new PaginatedEnvironments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PaginatedEnvironments) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginatedEnvironments) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginatedEnvironments) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaginatedEnvironments) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetValues

`func (o *PaginatedEnvironments) GetValues() []DeploymentEnvironment`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PaginatedEnvironments) GetValuesOk() (*[]DeploymentEnvironment, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PaginatedEnvironments) SetValues(v []DeploymentEnvironment)`

SetValues sets Values field to given value.

### HasValues

`func (o *PaginatedEnvironments) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetSize

`func (o *PaginatedEnvironments) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PaginatedEnvironments) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PaginatedEnvironments) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PaginatedEnvironments) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPagelen

`func (o *PaginatedEnvironments) GetPagelen() int32`

GetPagelen returns the Pagelen field if non-nil, zero value otherwise.

### GetPagelenOk

`func (o *PaginatedEnvironments) GetPagelenOk() (*int32, bool)`

GetPagelenOk returns a tuple with the Pagelen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagelen

`func (o *PaginatedEnvironments) SetPagelen(v int32)`

SetPagelen sets Pagelen field to given value.

### HasPagelen

`func (o *PaginatedEnvironments) HasPagelen() bool`

HasPagelen returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedEnvironments) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedEnvironments) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedEnvironments) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedEnvironments) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedEnvironments) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedEnvironments) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedEnvironments) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedEnvironments) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


