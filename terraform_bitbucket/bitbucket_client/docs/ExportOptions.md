# ExportOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ProjectKey** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**SendEmail** | Pointer to **bool** |  | [optional] 
**IncludeAttachments** | Pointer to **bool** |  | [optional] 

## Methods

### NewExportOptions

`func NewExportOptions(type_ string, ) *ExportOptions`

NewExportOptions instantiates a new ExportOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportOptionsWithDefaults

`func NewExportOptionsWithDefaults() *ExportOptions`

NewExportOptionsWithDefaults instantiates a new ExportOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExportOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportOptions) SetType(v string)`

SetType sets Type field to given value.


### GetProjectKey

`func (o *ExportOptions) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *ExportOptions) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *ExportOptions) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *ExportOptions) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetProjectName

`func (o *ExportOptions) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ExportOptions) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ExportOptions) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ExportOptions) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetSendEmail

`func (o *ExportOptions) GetSendEmail() bool`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *ExportOptions) GetSendEmailOk() (*bool, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *ExportOptions) SetSendEmail(v bool)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *ExportOptions) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.

### GetIncludeAttachments

`func (o *ExportOptions) GetIncludeAttachments() bool`

GetIncludeAttachments returns the IncludeAttachments field if non-nil, zero value otherwise.

### GetIncludeAttachmentsOk

`func (o *ExportOptions) GetIncludeAttachmentsOk() (*bool, bool)`

GetIncludeAttachmentsOk returns a tuple with the IncludeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttachments

`func (o *ExportOptions) SetIncludeAttachments(v bool)`

SetIncludeAttachments sets IncludeAttachments field to given value.

### HasIncludeAttachments

`func (o *ExportOptions) HasIncludeAttachments() bool`

HasIncludeAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


