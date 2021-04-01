# DeploymentEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID identifying the environment. | [optional] 
**Name** | Pointer to **string** | The name of the environment. | [optional] 

## Methods

### NewDeploymentEnvironment

`func NewDeploymentEnvironment() *DeploymentEnvironment`

NewDeploymentEnvironment instantiates a new DeploymentEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentEnvironmentWithDefaults

`func NewDeploymentEnvironmentWithDefaults() *DeploymentEnvironment`

NewDeploymentEnvironmentWithDefaults instantiates a new DeploymentEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DeploymentEnvironment) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeploymentEnvironment) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeploymentEnvironment) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeploymentEnvironment) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *DeploymentEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


