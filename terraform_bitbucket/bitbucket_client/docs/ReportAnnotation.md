# ReportAnnotation

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

### NewReportAnnotation

`func NewReportAnnotation() *ReportAnnotation`

NewReportAnnotation instantiates a new ReportAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportAnnotationWithDefaults

`func NewReportAnnotationWithDefaults() *ReportAnnotation`

NewReportAnnotationWithDefaults instantiates a new ReportAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ReportAnnotation) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ReportAnnotation) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ReportAnnotation) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ReportAnnotation) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUuid

`func (o *ReportAnnotation) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ReportAnnotation) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ReportAnnotation) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ReportAnnotation) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAnnotationType

`func (o *ReportAnnotation) GetAnnotationType() string`

GetAnnotationType returns the AnnotationType field if non-nil, zero value otherwise.

### GetAnnotationTypeOk

`func (o *ReportAnnotation) GetAnnotationTypeOk() (*string, bool)`

GetAnnotationTypeOk returns a tuple with the AnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationType

`func (o *ReportAnnotation) SetAnnotationType(v string)`

SetAnnotationType sets AnnotationType field to given value.

### HasAnnotationType

`func (o *ReportAnnotation) HasAnnotationType() bool`

HasAnnotationType returns a boolean if a field has been set.

### GetPath

`func (o *ReportAnnotation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReportAnnotation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReportAnnotation) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ReportAnnotation) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetLine

`func (o *ReportAnnotation) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ReportAnnotation) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *ReportAnnotation) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *ReportAnnotation) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetSummary

`func (o *ReportAnnotation) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ReportAnnotation) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ReportAnnotation) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ReportAnnotation) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDetails

`func (o *ReportAnnotation) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ReportAnnotation) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ReportAnnotation) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ReportAnnotation) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetResult

`func (o *ReportAnnotation) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ReportAnnotation) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ReportAnnotation) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ReportAnnotation) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSeverity

`func (o *ReportAnnotation) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ReportAnnotation) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ReportAnnotation) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ReportAnnotation) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetLink

`func (o *ReportAnnotation) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ReportAnnotation) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ReportAnnotation) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ReportAnnotation) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ReportAnnotation) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ReportAnnotation) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ReportAnnotation) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ReportAnnotation) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *ReportAnnotation) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *ReportAnnotation) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *ReportAnnotation) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *ReportAnnotation) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


