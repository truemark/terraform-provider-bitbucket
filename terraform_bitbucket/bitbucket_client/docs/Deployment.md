# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID identifying the deployment. | [optional] 
**State** | Pointer to [**DeploymentState**](DeploymentState.md) |  | [optional] 
**Environment** | Pointer to [**DeploymentEnvironment**](DeploymentEnvironment.md) |  | [optional] 
**Release** | Pointer to [**DeploymentRelease**](DeploymentRelease.md) |  | [optional] 

## Methods

### NewDeployment

`func NewDeployment() *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Deployment) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Deployment) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Deployment) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Deployment) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *Deployment) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Deployment) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Deployment) SetState(v DeploymentState)`

SetState sets State field to given value.

### HasState

`func (o *Deployment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEnvironment

`func (o *Deployment) GetEnvironment() DeploymentEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Deployment) GetEnvironmentOk() (*DeploymentEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Deployment) SetEnvironment(v DeploymentEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Deployment) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRelease

`func (o *Deployment) GetRelease() DeploymentRelease`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *Deployment) GetReleaseOk() (*DeploymentRelease, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *Deployment) SetRelease(v DeploymentRelease)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *Deployment) HasRelease() bool`

HasRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


