# ReportAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID that can be used to identify the report. | [optional] 
**Title** | Pointer to **string** | The title of the report. | [optional] 
**Details** | Pointer to **string** | A string to describe the purpose of the report. | [optional] 
**ExternalId** | Pointer to **string** | ID of the report provided by the report creator. It can be used to identify the report as an alternative to it&#39;s generated uuid. It is not used by Bitbucket, but only by the report creator for updating or deleting this specific report. Needs to be unique. | [optional] 
**Reporter** | Pointer to **string** | A string to describe the tool or company who created the report. | [optional] 
**Link** | Pointer to **string** | A URL linking to the results of the report in an external tool. | [optional] 
**RemoteLinkEnabled** | Pointer to **bool** | If enabled, a remote link is created in Jira for the issue associated with the commit the report belongs to. | [optional] 
**LogoUrl** | Pointer to **string** | A URL to the report logo. If none is provided, the default insights logo will be used. | [optional] 
**ReportType** | Pointer to **string** | The type of the report. | [optional] 
**Result** | Pointer to **string** | The state of the report. May be set to PENDING and later updated. | [optional] 
**Data** | Pointer to [**[]ReportData**](ReportData.md) | An array of data fields to display information on the report. Maximum 10. | [optional] 
**CreatedOn** | Pointer to **time.Time** | The timestamp when the report was created. | [optional] 
**UpdatedOn** | Pointer to **time.Time** | The timestamp when the report was updated. | [optional] 

## Methods

### NewReportAllOf

`func NewReportAllOf() *ReportAllOf`

NewReportAllOf instantiates a new ReportAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportAllOfWithDefaults

`func NewReportAllOfWithDefaults() *ReportAllOf`

NewReportAllOfWithDefaults instantiates a new ReportAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ReportAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ReportAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ReportAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ReportAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetTitle

`func (o *ReportAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReportAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReportAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReportAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetails

`func (o *ReportAllOf) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ReportAllOf) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ReportAllOf) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ReportAllOf) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetExternalId

`func (o *ReportAllOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ReportAllOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ReportAllOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ReportAllOf) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetReporter

`func (o *ReportAllOf) GetReporter() string`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *ReportAllOf) GetReporterOk() (*string, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *ReportAllOf) SetReporter(v string)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *ReportAllOf) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### GetLink

`func (o *ReportAllOf) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ReportAllOf) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ReportAllOf) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ReportAllOf) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetRemoteLinkEnabled

`func (o *ReportAllOf) GetRemoteLinkEnabled() bool`

GetRemoteLinkEnabled returns the RemoteLinkEnabled field if non-nil, zero value otherwise.

### GetRemoteLinkEnabledOk

`func (o *ReportAllOf) GetRemoteLinkEnabledOk() (*bool, bool)`

GetRemoteLinkEnabledOk returns a tuple with the RemoteLinkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLinkEnabled

`func (o *ReportAllOf) SetRemoteLinkEnabled(v bool)`

SetRemoteLinkEnabled sets RemoteLinkEnabled field to given value.

### HasRemoteLinkEnabled

`func (o *ReportAllOf) HasRemoteLinkEnabled() bool`

HasRemoteLinkEnabled returns a boolean if a field has been set.

### GetLogoUrl

`func (o *ReportAllOf) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *ReportAllOf) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *ReportAllOf) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *ReportAllOf) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetReportType

`func (o *ReportAllOf) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportAllOf) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportAllOf) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *ReportAllOf) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetResult

`func (o *ReportAllOf) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ReportAllOf) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ReportAllOf) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ReportAllOf) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetData

`func (o *ReportAllOf) GetData() []ReportData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReportAllOf) GetDataOk() (*[]ReportData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReportAllOf) SetData(v []ReportData)`

SetData sets Data field to given value.

### HasData

`func (o *ReportAllOf) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ReportAllOf) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ReportAllOf) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ReportAllOf) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ReportAllOf) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *ReportAllOf) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *ReportAllOf) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *ReportAllOf) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *ReportAllOf) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


