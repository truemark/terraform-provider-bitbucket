# IssueJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | The status of the import/export job | [optional] 
**Phase** | Pointer to **string** | The phase of the import/export job | [optional] 
**Total** | Pointer to **int32** | The total number of issues being imported/exported | [optional] 
**Count** | Pointer to **int32** | The total number of issues already imported/exported | [optional] 
**Pct** | Pointer to **float32** | The percentage of issues already imported/exported | [optional] 

## Methods

### NewIssueJobStatus

`func NewIssueJobStatus() *IssueJobStatus`

NewIssueJobStatus instantiates a new IssueJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueJobStatusWithDefaults

`func NewIssueJobStatusWithDefaults() *IssueJobStatus`

NewIssueJobStatusWithDefaults instantiates a new IssueJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IssueJobStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueJobStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueJobStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IssueJobStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *IssueJobStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IssueJobStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IssueJobStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IssueJobStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPhase

`func (o *IssueJobStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *IssueJobStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *IssueJobStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *IssueJobStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetTotal

`func (o *IssueJobStatus) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IssueJobStatus) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IssueJobStatus) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IssueJobStatus) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCount

`func (o *IssueJobStatus) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IssueJobStatus) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IssueJobStatus) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IssueJobStatus) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPct

`func (o *IssueJobStatus) GetPct() float32`

GetPct returns the Pct field if non-nil, zero value otherwise.

### GetPctOk

`func (o *IssueJobStatus) GetPctOk() (*float32, bool)`

GetPctOk returns a tuple with the Pct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPct

`func (o *IssueJobStatus) SetPct(v float32)`

SetPct sets Pct field to given value.

### HasPct

`func (o *IssueJobStatus) HasPct() bool`

HasPct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


