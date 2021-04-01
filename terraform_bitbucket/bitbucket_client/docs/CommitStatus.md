# Commitstatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CommitStatusLinks**](CommitStatusLinks.md) |  | [optional] 
**Uuid** | Pointer to **string** | The commit status&#39; id. | [optional] 
**Key** | Pointer to **string** | An identifier for the status that&#39;s unique to         its type (current \&quot;build\&quot; is the only supported type) and the vendor,         e.g. BB-DEPLOY | [optional] 
**Refname** | Pointer to **string** |  The name of the ref that pointed to this commit at the time the status object was created. Note that this the ref may since have moved off of the commit. This optional field can be useful for build systems whose build triggers and configuration are branch-dependent (e.g. a Pipeline build). It is legitimate for this field to not be set, or even apply (e.g. a static linting job). | [optional] 
**Url** | Pointer to **string** | A URL linking back to the vendor or build system, for providing more information about whatever process produced this status. Accepts context variables &#x60;repository&#x60; and &#x60;commit&#x60; that Bitbucket will evaluate at runtime whenever at runtime. For example, one could use https://foo.com/builds/{repository.full_name} which Bitbucket will turn into https://foo.com/builds/foo/bar at render time. | [optional] 
**State** | Pointer to **string** | Provides some indication of the status of this commit | [optional] 
**Name** | Pointer to **string** | An identifier for the build itself, e.g. BB-DEPLOY-1 | [optional] 
**Description** | Pointer to **string** | A description of the build (e.g. \&quot;Unit tests in Bamboo\&quot;) | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**UpdatedOn** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCommitstatus

`func NewCommitstatus() *Commitstatus`

NewCommitstatus instantiates a new Commitstatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitstatusWithDefaults

`func NewCommitstatusWithDefaults() *Commitstatus`

NewCommitstatusWithDefaults instantiates a new Commitstatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Commitstatus) GetLinks() CommitStatusLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Commitstatus) GetLinksOk() (*CommitStatusLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Commitstatus) SetLinks(v CommitStatusLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Commitstatus) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetUuid

`func (o *Commitstatus) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Commitstatus) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Commitstatus) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Commitstatus) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetKey

`func (o *Commitstatus) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Commitstatus) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Commitstatus) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Commitstatus) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRefname

`func (o *Commitstatus) GetRefname() string`

GetRefname returns the Refname field if non-nil, zero value otherwise.

### GetRefnameOk

`func (o *Commitstatus) GetRefnameOk() (*string, bool)`

GetRefnameOk returns a tuple with the Refname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefname

`func (o *Commitstatus) SetRefname(v string)`

SetRefname sets Refname field to given value.

### HasRefname

`func (o *Commitstatus) HasRefname() bool`

HasRefname returns a boolean if a field has been set.

### GetUrl

`func (o *Commitstatus) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Commitstatus) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Commitstatus) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Commitstatus) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetState

`func (o *Commitstatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Commitstatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Commitstatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Commitstatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *Commitstatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Commitstatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Commitstatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Commitstatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Commitstatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Commitstatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Commitstatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Commitstatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Commitstatus) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Commitstatus) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Commitstatus) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Commitstatus) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *Commitstatus) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *Commitstatus) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *Commitstatus) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *Commitstatus) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


