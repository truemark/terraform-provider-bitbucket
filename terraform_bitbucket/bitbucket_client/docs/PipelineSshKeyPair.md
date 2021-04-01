# PipelineSshKeyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | Pointer to **string** | The SSH private key. This value will be empty when retrieving the SSH key pair. | [optional] 
**PublicKey** | Pointer to **string** | The SSH public key. | [optional] 

## Methods

### NewPipelineSshKeyPair

`func NewPipelineSshKeyPair() *PipelineSshKeyPair`

NewPipelineSshKeyPair instantiates a new PipelineSshKeyPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineSshKeyPairWithDefaults

`func NewPipelineSshKeyPairWithDefaults() *PipelineSshKeyPair`

NewPipelineSshKeyPairWithDefaults instantiates a new PipelineSshKeyPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *PipelineSshKeyPair) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PipelineSshKeyPair) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PipelineSshKeyPair) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *PipelineSshKeyPair) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *PipelineSshKeyPair) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PipelineSshKeyPair) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PipelineSshKeyPair) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *PipelineSshKeyPair) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


