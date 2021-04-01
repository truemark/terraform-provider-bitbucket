# ReportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of data contained in the value field. If not provided, then the value will be detected as a boolean, number or string. | [optional] 
**Title** | Pointer to **string** | A string describing what this data field represents. | [optional] 
**Value** | Pointer to **map[string]interface{}** | The value of the data element. | [optional] 

## Methods

### NewReportData

`func NewReportData() *ReportData`

NewReportData instantiates a new ReportData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportDataWithDefaults

`func NewReportDataWithDefaults() *ReportData`

NewReportDataWithDefaults instantiates a new ReportData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReportData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReportData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *ReportData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReportData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReportData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReportData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetValue

`func (o *ReportData) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReportData) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReportData) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ReportData) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


