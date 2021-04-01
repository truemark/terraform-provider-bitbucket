# IssueContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Raw** | Pointer to **string** | The text as it was typed by a user. | [optional] 
**Markup** | Pointer to **string** | The type of markup language the raw content is to be interpreted in. | [optional] 
**Html** | Pointer to **string** | The user&#39;s content rendered as HTML. | [optional] 

## Methods

### NewIssueContent

`func NewIssueContent() *IssueContent`

NewIssueContent instantiates a new IssueContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueContentWithDefaults

`func NewIssueContentWithDefaults() *IssueContent`

NewIssueContentWithDefaults instantiates a new IssueContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaw

`func (o *IssueContent) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *IssueContent) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *IssueContent) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *IssueContent) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetMarkup

`func (o *IssueContent) GetMarkup() string`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *IssueContent) GetMarkupOk() (*string, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *IssueContent) SetMarkup(v string)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *IssueContent) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### GetHtml

`func (o *IssueContent) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *IssueContent) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *IssueContent) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *IssueContent) HasHtml() bool`

HasHtml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


