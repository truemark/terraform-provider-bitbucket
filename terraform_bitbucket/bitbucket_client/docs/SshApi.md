# \SshApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersSelectedUserSshKeysGet**](SshApi.md#UsersSelectedUserSshKeysGet) | **Get** /users/{selected_user}/ssh-keys | 
[**UsersSelectedUserSshKeysKeyIdDelete**](SshApi.md#UsersSelectedUserSshKeysKeyIdDelete) | **Delete** /users/{selected_user}/ssh-keys/{key_id} | 
[**UsersSelectedUserSshKeysKeyIdGet**](SshApi.md#UsersSelectedUserSshKeysKeyIdGet) | **Get** /users/{selected_user}/ssh-keys/{key_id} | 
[**UsersSelectedUserSshKeysKeyIdPut**](SshApi.md#UsersSelectedUserSshKeysKeyIdPut) | **Put** /users/{selected_user}/ssh-keys/{key_id} | 
[**UsersSelectedUserSshKeysPost**](SshApi.md#UsersSelectedUserSshKeysPost) | **Post** /users/{selected_user}/ssh-keys | 



## UsersSelectedUserSshKeysGet

> PaginatedSshUserKeys UsersSelectedUserSshKeysGet(ctx, selectedUser).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    selectedUser := "selectedUser_example" // string | This can either be the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`, OR an Atlassian Account ID. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SshApi.UsersSelectedUserSshKeysGet(context.Background(), selectedUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshApi.UsersSelectedUserSshKeysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserSshKeysGet`: PaginatedSshUserKeys
    fmt.Fprintf(os.Stdout, "Response from `SshApi.UsersSelectedUserSshKeysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersSelectedUserSshKeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedSshUserKeys**](PaginatedSshUserKeys.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSelectedUserSshKeysKeyIdDelete

> UsersSelectedUserSshKeysKeyIdDelete(ctx, keyId, selectedUser).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keyId := "keyId_example" // string | The SSH key's UUID value.
    selectedUser := "selectedUser_example" // string | This can either be the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`, OR an Atlassian Account ID. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SshApi.UsersSelectedUserSshKeysKeyIdDelete(context.Background(), keyId, selectedUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshApi.UsersSelectedUserSshKeysKeyIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The SSH key&#39;s UUID value. | 
**selectedUser** | **string** | This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersSelectedUserSshKeysKeyIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSelectedUserSshKeysKeyIdGet

> SshAccountKey UsersSelectedUserSshKeysKeyIdGet(ctx, keyId, selectedUser).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keyId := "keyId_example" // string | The SSH key's UUID value.
    selectedUser := "selectedUser_example" // string | This can either be the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`, OR an Atlassian Account ID. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SshApi.UsersSelectedUserSshKeysKeyIdGet(context.Background(), keyId, selectedUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshApi.UsersSelectedUserSshKeysKeyIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserSshKeysKeyIdGet`: SshAccountKey
    fmt.Fprintf(os.Stdout, "Response from `SshApi.UsersSelectedUserSshKeysKeyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The SSH key&#39;s UUID value. | 
**selectedUser** | **string** | This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersSelectedUserSshKeysKeyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SshAccountKey**](SshAccountKey.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSelectedUserSshKeysKeyIdPut

> SshAccountKey UsersSelectedUserSshKeysKeyIdPut(ctx, keyId, selectedUser).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keyId := "keyId_example" // string | The SSH key's UUID value.
    selectedUser := "selectedUser_example" // string | This can either be the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`, OR an Atlassian Account ID. 
    body := *openapiclient.NewSshAccountKey("Type_example") // SshAccountKey | The updated SSH key object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SshApi.UsersSelectedUserSshKeysKeyIdPut(context.Background(), keyId, selectedUser).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshApi.UsersSelectedUserSshKeysKeyIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserSshKeysKeyIdPut`: SshAccountKey
    fmt.Fprintf(os.Stdout, "Response from `SshApi.UsersSelectedUserSshKeysKeyIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The SSH key&#39;s UUID value. | 
**selectedUser** | **string** | This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersSelectedUserSshKeysKeyIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SshAccountKey**](SshAccountKey.md) | The updated SSH key object | 

### Return type

[**SshAccountKey**](SshAccountKey.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSelectedUserSshKeysPost

> SshAccountKey UsersSelectedUserSshKeysPost(ctx, selectedUser).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    selectedUser := "selectedUser_example" // string | This can either be the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`, OR an Atlassian Account ID. 
    body := *openapiclient.NewSshAccountKey("Type_example") // SshAccountKey | The new SSH key object. Note that the username property has been deprecated due to [privacy changes](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-changes-gdpr/#removal-of-usernames-from-user-referencing-apis). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SshApi.UsersSelectedUserSshKeysPost(context.Background(), selectedUser).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshApi.UsersSelectedUserSshKeysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserSshKeysPost`: SshAccountKey
    fmt.Fprintf(os.Stdout, "Response from `SshApi.UsersSelectedUserSshKeysPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersSelectedUserSshKeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SshAccountKey**](SshAccountKey.md) | The new SSH key object. Note that the username property has been deprecated due to [privacy changes](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-changes-gdpr/#removal-of-usernames-from-user-referencing-apis). | 

### Return type

[**SshAccountKey**](SshAccountKey.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

