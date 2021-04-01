# ReportAnnotationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | ID of the annotation provided by the annotation creator. It can be used to identify the annotation as an alternative to it&#39;s generated uuid. It is not used by Bitbucket, but only by the annotation creator for updating or deleting this specific annotation. Needs to be unique. | [optional] 
**Uuid** | Pointer to **string** | The UUID that can be used to identify the annotation. | [optional] 
**AnnotationType** | Pointer to **string** | The type of the report. | [optional] 
**Path** | Pointer to **string** | The path of the file on which this annotation should be placed. This is the path of the file relative to the git repository. If no path is provided, then it will appear in the overview modal on all pull requests where the tip of the branch is the given commit, regardless of which files were modified. | [optional] 
**Line** | Pointer to **int32** | The line number that the annotation should belong to. If no line number is provided, then it will default to 0 and in a pull request it will appear at the top of the file specified by the path field. | [optional] 
**Summary** | Pointer to **string** | The message to display to users. | [optional] 
**Details** | Pointer to **string** | The details to show to users when clicking on the annotation. | [optional] 
**Result** | Pointer to **string** | The state of the report. May be set to PENDING and later updated. | [optional] 
**Severity** | Pointer to **string** | The severity of the annotation. | [optional] 
**Link** | Pointer to **string** | A URL linking to the annotation in an external tool. | [optional] 
**CreatedOn** | Pointer to **time.Time** | The timestamp when the report was created. | [optional] 
**UpdatedOn** | Pointer to **time.Time** | The timestamp when the report was updated. | [optional] 

## Methods

### NewReportAnnotationAllOf

`func NewReportAnnotationAllOf() *ReportAnnotationAllOf`

NewReportAnnotationAllOf instantiates a new ReportAnnotationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportAnnotationAllOfWithDefaults

`func NewReportAnnotationAllOfWithDefaults() *ReportAnnotationAllOf`

NewReportAnnotationAllOfWithDefaults instantiates a new ReportAnnotationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ReportAnnotationAllOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ReportAnnotationAllOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ReportAnnotationAllOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ReportAnnotationAllOf) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUuid

`func (o *ReportAnnotationAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ReportAnnotationAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ReportAnnotationAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ReportAnnotationAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAnnotationType

`func (o *ReportAnnotationAllOf) GetAnnotationType() string`

GetAnnotationType returns the AnnotationType field if non-nil, zero value otherwise.

### GetAnnotationTypeOk

`func (o *ReportAnnotationAllOf) GetAnnotationTypeOk() (*string, bool)`

GetAnnotationTypeOk returns a tuple with the AnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationType

`func (o *ReportAnnotationAllOf) SetAnnotationType(v string)`

SetAnnotationType sets AnnotationType field to given value.

### HasAnnotationType

`func (o *ReportAnnotationAllOf) HasAnnotationType() bool`

HasAnnotationType returns a boolean if a field has been set.

### GetPath

`func (o *ReportAnnotationAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReportAnnotationAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReportAnnotationAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ReportAnnotationAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetLine

`func (o *ReportAnnotationAllOf) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ReportAnnotationAllOf) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *ReportAnnotationAllOf) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *ReportAnnotationAllOf) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetSummary

`func (o *ReportAnnotationAllOf) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ReportAnnotationAllOf) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ReportAnnotationAllOf) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ReportAnnotationAllOf) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDetails

`func (o *ReportAnnotationAllOf) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ReportAnnotationAllOf) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ReportAnnotationAllOf) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ReportAnnotationAllOf) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetResult

`func (o *ReportAnnotationAllOf) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ReportAnnotationAllOf) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ReportAnnotationAllOf) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ReportAnnotationAllOf) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSeverity

`func (o *ReportAnnotationAllOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ReportAnnotationAllOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ReportAnnotationAllOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ReportAnnotationAllOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetLink

`func (o *ReportAnnotationAllOf) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ReportAnnotationAllOf) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ReportAnnotationAllOf) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ReportAnnotationAllOf) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ReportAnnotationAllOf) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ReportAnnotationAllOf) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ReportAnnotationAllOf) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ReportAnnotationAllOf) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *ReportAnnotationAllOf) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *ReportAnnotationAllOf) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *ReportAnnotationAllOf) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *ReportAnnotationAllOf) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


