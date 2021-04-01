# WebhookSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The webhook&#39;s id | [optional] 
**Url** | Pointer to **string** | The URL events get delivered to. | [optional] 
**Description** | Pointer to **string** | A user-defined description of the webhook. | [optional] 
**SubjectType** | Pointer to **string** | The type of entity, which is &#x60;repository&#x60; in the case of webhook subscriptions on repositories. | [optional] 
**Subject** | Pointer to [**Object**](Object.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Events** | Pointer to **[]string** | The events this webhook is subscribed to. | [optional] 

## Methods

### NewWebhookSubscription

`func NewWebhookSubscription() *WebhookSubscription`

NewWebhookSubscription instantiates a new WebhookSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionWithDefaults

`func NewWebhookSubscriptionWithDefaults() *WebhookSubscription`

NewWebhookSubscriptionWithDefaults instantiates a new WebhookSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *WebhookSubscription) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *WebhookSubscription) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *WebhookSubscription) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *WebhookSubscription) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookSubscription) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookSubscription) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookSubscription) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookSubscription) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDescription

`func (o *WebhookSubscription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookSubscription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookSubscription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookSubscription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubjectType

`func (o *WebhookSubscription) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *WebhookSubscription) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *WebhookSubscription) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *WebhookSubscription) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### GetSubject

`func (o *WebhookSubscription) GetSubject() Object`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WebhookSubscription) GetSubjectOk() (*Object, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WebhookSubscription) SetSubject(v Object)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *WebhookSubscription) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetActive

`func (o *WebhookSubscription) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookSubscription) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookSubscription) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookSubscription) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookSubscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookSubscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookSubscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebhookSubscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvents

`func (o *WebhookSubscription) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookSubscription) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookSubscription) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhookSubscription) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


