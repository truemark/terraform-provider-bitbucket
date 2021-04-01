# PipelineImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the image. If the image is hosted on DockerHub the short name can be used, otherwise the fully qualified name is required here. | [optional] 
**Username** | Pointer to **string** | The username needed to authenticate with the Docker registry. Only required when using a private Docker image. | [optional] 
**Password** | Pointer to **string** | The password needed to authenticate with the Docker registry. Only required when using a private Docker image. | [optional] 
**Email** | Pointer to **string** | The email needed to authenticate with the Docker registry. Only required when using a private Docker image. | [optional] 

## Methods

### NewPipelineImage

`func NewPipelineImage() *PipelineImage`

NewPipelineImage instantiates a new PipelineImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineImageWithDefaults

`func NewPipelineImageWithDefaults() *PipelineImage`

NewPipelineImageWithDefaults instantiates a new PipelineImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PipelineImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PipelineImage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *PipelineImage) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PipelineImage) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PipelineImage) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PipelineImage) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *PipelineImage) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PipelineImage) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PipelineImage) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PipelineImage) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmail

`func (o *PipelineImage) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PipelineImage) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PipelineImage) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PipelineImage) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


