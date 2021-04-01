# PipelineSshPublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyType** | Pointer to **string** | The type of the public key. | [optional] 
**Key** | Pointer to **string** | The base64 encoded public key. | [optional] 
**Md5Fingerprint** | Pointer to **string** | The MD5 fingerprint of the public key. | [optional] 
**Sha256Fingerprint** | Pointer to **string** | The SHA-256 fingerprint of the public key. | [optional] 

## Methods

### NewPipelineSshPublicKey

`func NewPipelineSshPublicKey() *PipelineSshPublicKey`

NewPipelineSshPublicKey instantiates a new PipelineSshPublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineSshPublicKeyWithDefaults

`func NewPipelineSshPublicKeyWithDefaults() *PipelineSshPublicKey`

NewPipelineSshPublicKeyWithDefaults instantiates a new PipelineSshPublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyType

`func (o *PipelineSshPublicKey) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PipelineSshPublicKey) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PipelineSshPublicKey) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *PipelineSshPublicKey) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetKey

`func (o *PipelineSshPublicKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PipelineSshPublicKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PipelineSshPublicKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PipelineSshPublicKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMd5Fingerprint

`func (o *PipelineSshPublicKey) GetMd5Fingerprint() string`

GetMd5Fingerprint returns the Md5Fingerprint field if non-nil, zero value otherwise.

### GetMd5FingerprintOk

`func (o *PipelineSshPublicKey) GetMd5FingerprintOk() (*string, bool)`

GetMd5FingerprintOk returns a tuple with the Md5Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Fingerprint

`func (o *PipelineSshPublicKey) SetMd5Fingerprint(v string)`

SetMd5Fingerprint sets Md5Fingerprint field to given value.

### HasMd5Fingerprint

`func (o *PipelineSshPublicKey) HasMd5Fingerprint() bool`

HasMd5Fingerprint returns a boolean if a field has been set.

### GetSha256Fingerprint

`func (o *PipelineSshPublicKey) GetSha256Fingerprint() string`

GetSha256Fingerprint returns the Sha256Fingerprint field if non-nil, zero value otherwise.

### GetSha256FingerprintOk

`func (o *PipelineSshPublicKey) GetSha256FingerprintOk() (*string, bool)`

GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Fingerprint

`func (o *PipelineSshPublicKey) SetSha256Fingerprint(v string)`

SetSha256Fingerprint sets Sha256Fingerprint field to given value.

### HasSha256Fingerprint

`func (o *PipelineSshPublicKey) HasSha256Fingerprint() bool`

HasSha256Fingerprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


