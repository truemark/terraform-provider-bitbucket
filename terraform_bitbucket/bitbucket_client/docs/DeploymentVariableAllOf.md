# DeploymentVariableAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID identifying the variable. | [optional] 
**Key** | Pointer to **string** | The unique name of the variable. | [optional] 
**Value** | Pointer to **string** | The value of the variable. If the variable is secured, this will be empty. | [optional] 
**Secured** | Pointer to **bool** | If true, this variable will be treated as secured. The value will never be exposed in the logs or the REST API. | [optional] 

## Methods

### NewDeploymentVariableAllOf

`func NewDeploymentVariableAllOf() *DeploymentVariableAllOf`

NewDeploymentVariableAllOf instantiates a new DeploymentVariableAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentVariableAllOfWithDefaults

`func NewDeploymentVariableAllOfWithDefaults() *DeploymentVariableAllOf`

NewDeploymentVariableAllOfWithDefaults instantiates a new DeploymentVariableAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DeploymentVariableAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeploymentVariableAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeploymentVariableAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeploymentVariableAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetKey

`func (o *DeploymentVariableAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DeploymentVariableAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DeploymentVariableAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DeploymentVariableAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *DeploymentVariableAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeploymentVariableAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeploymentVariableAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DeploymentVariableAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSecured

`func (o *DeploymentVariableAllOf) GetSecured() bool`

GetSecured returns the Secured field if non-nil, zero value otherwise.

### GetSecuredOk

`func (o *DeploymentVariableAllOf) GetSecuredOk() (*bool, bool)`

GetSecuredOk returns a tuple with the Secured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecured

`func (o *DeploymentVariableAllOf) SetSecured(v bool)`

SetSecured sets Secured field to given value.

### HasSecured

`func (o *DeploymentVariableAllOf) HasSecured() bool`

HasSecured returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


