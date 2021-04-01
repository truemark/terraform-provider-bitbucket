# Diffstat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**LinesAdded** | Pointer to **int32** |  | [optional] 
**LinesRemoved** | Pointer to **int32** |  | [optional] 
**Old** | Pointer to [**CommitFile**](CommitFile.md) |  | [optional] 
**New** | Pointer to [**CommitFile**](CommitFile.md) |  | [optional] 

## Methods

### NewDiffstat

`func NewDiffstat(type_ string, ) *Diffstat`

NewDiffstat instantiates a new Diffstat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiffstatWithDefaults

`func NewDiffstatWithDefaults() *Diffstat`

NewDiffstatWithDefaults instantiates a new Diffstat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Diffstat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Diffstat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Diffstat) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *Diffstat) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Diffstat) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Diffstat) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Diffstat) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinesAdded

`func (o *Diffstat) GetLinesAdded() int32`

GetLinesAdded returns the LinesAdded field if non-nil, zero value otherwise.

### GetLinesAddedOk

`func (o *Diffstat) GetLinesAddedOk() (*int32, bool)`

GetLinesAddedOk returns a tuple with the LinesAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesAdded

`func (o *Diffstat) SetLinesAdded(v int32)`

SetLinesAdded sets LinesAdded field to given value.

### HasLinesAdded

`func (o *Diffstat) HasLinesAdded() bool`

HasLinesAdded returns a boolean if a field has been set.

### GetLinesRemoved

`func (o *Diffstat) GetLinesRemoved() int32`

GetLinesRemoved returns the LinesRemoved field if non-nil, zero value otherwise.

### GetLinesRemovedOk

`func (o *Diffstat) GetLinesRemovedOk() (*int32, bool)`

GetLinesRemovedOk returns a tuple with the LinesRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesRemoved

`func (o *Diffstat) SetLinesRemoved(v int32)`

SetLinesRemoved sets LinesRemoved field to given value.

### HasLinesRemoved

`func (o *Diffstat) HasLinesRemoved() bool`

HasLinesRemoved returns a boolean if a field has been set.

### GetOld

`func (o *Diffstat) GetOld() CommitFile`

GetOld returns the Old field if non-nil, zero value otherwise.

### GetOldOk

`func (o *Diffstat) GetOldOk() (*CommitFile, bool)`

GetOldOk returns a tuple with the Old field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOld

`func (o *Diffstat) SetOld(v CommitFile)`

SetOld sets Old field to given value.

### HasOld

`func (o *Diffstat) HasOld() bool`

HasOld returns a boolean if a field has been set.

### GetNew

`func (o *Diffstat) GetNew() CommitFile`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *Diffstat) GetNewOk() (*CommitFile, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *Diffstat) SetNew(v CommitFile)`

SetNew sets New field to given value.

### HasNew

`func (o *Diffstat) HasNew() bool`

HasNew returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


