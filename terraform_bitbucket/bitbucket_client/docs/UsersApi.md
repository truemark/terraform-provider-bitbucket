# \UsersApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsWorkspaceRepositoriesGet**](UsersApi.md#TeamsWorkspaceRepositoriesGet) | **Get** /teams/{workspace}/repositories | 
[**UserEmailsEmailGet**](UsersApi.md#UserEmailsEmailGet) | **Get** /user/emails/{email} | 
[**UserEmailsGet**](UsersApi.md#UserEmailsGet) | **Get** /user/emails | 
[**UserGet**](UsersApi.md#UserGet) | **Get** /user | 
[**UsersSelectedUserGet**](UsersApi.md#UsersSelectedUserGet) | **Get** /users/{selected_user} | 
[**UsersSelectedUserHooksGet**](UsersApi.md#UsersSelectedUserHooksGet) | **Get** /users/{selected_user}/hooks | 
[**UsersSelectedUserHooksPost**](UsersApi.md#UsersSelectedUserHooksPost) | **Post** /users/{selected_user}/hooks | 
[**UsersSelectedUserHooksUidDelete**](UsersApi.md#UsersSelectedUserHooksUidDelete) | **Delete** /users/{selected_user}/hooks/{uid} | 
[**UsersSelectedUserHooksUidGet**](UsersApi.md#UsersSelectedUserHooksUidGet) | **Get** /users/{selected_user}/hooks/{uid} | 
[**UsersSelectedUserHooksUidPut**](UsersApi.md#UsersSelectedUserHooksUidPut) | **Put** /users/{selected_user}/hooks/{uid} | 
[**UsersUsernameMembersGet**](UsersApi.md#UsersUsernameMembersGet) | **Get** /users/{username}/members | 
[**UsersWorkspaceRepositoriesGet**](UsersApi.md#UsersWorkspaceRepositoriesGet) | **Get** /users/{workspace}/repositories | 



## TeamsWorkspaceRepositoriesGet

> ModelError TeamsWorkspaceRepositoriesGet(ctx, workspace).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.TeamsWorkspaceRepositoriesGet(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.TeamsWorkspaceRepositoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsWorkspaceRepositoriesGet`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.TeamsWorkspaceRepositoriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsWorkspaceRepositoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelError**](ModelError.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserEmailsEmailGet

> ModelError UserEmailsEmailGet(ctx, email).Execute()





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
    email := "email_example" // string | Email address of the user.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UserEmailsEmailGet(context.Background(), email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UserEmailsEmailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserEmailsEmailGet`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UserEmailsEmailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserEmailsEmailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelError**](ModelError.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserEmailsGet

> ModelError UserEmailsGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UserEmailsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UserEmailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserEmailsGet`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UserEmailsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserEmailsGetRequest struct via the builder pattern


### Return type

[**ModelError**](ModelError.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGet

> User UserGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UserGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UserGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGet`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UserGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSelectedUserGet

> User UsersSelectedUserGet(ctx, selectedUser).Execute()





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
    resp, r, err := api_client.UsersApi.UsersSelectedUserGet(context.Background(), selectedUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersSelectedUserGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserGet`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersSelectedUserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersSelectedUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSelectedUserHooksGet

> PaginatedWebhookSubscriptions UsersSelectedUserHooksGet(ctx, selectedUser).Execute()





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
    resp, r, err := api_client.UsersApi.UsersSelectedUserHooksGet(context.Background(), selectedUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersSelectedUserHooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserHooksGet`: PaginatedWebhookSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersSelectedUserHooksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersSelectedUserHooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedWebhookSubscriptions**](PaginatedWebhookSubscriptions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSelectedUserHooksPost

> WebhookSubscription UsersSelectedUserHooksPost(ctx, selectedUser).Execute()





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
    resp, r, err := api_client.UsersApi.UsersSelectedUserHooksPost(context.Background(), selectedUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersSelectedUserHooksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserHooksPost`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersSelectedUserHooksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersSelectedUserHooksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookSubscription**](WebhookSubscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSelectedUserHooksUidDelete

> UsersSelectedUserHooksUidDelete(ctx, selectedUser, uid).Execute()





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
    uid := "uid_example" // string | Installed webhook's ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersSelectedUserHooksUidDelete(context.Background(), selectedUser, uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersSelectedUserHooksUidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 
**uid** | **string** | Installed webhook&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersSelectedUserHooksUidDeleteRequest struct via the builder pattern


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


## UsersSelectedUserHooksUidGet

> WebhookSubscription UsersSelectedUserHooksUidGet(ctx, selectedUser, uid).Execute()





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
    uid := "uid_example" // string | Installed webhook's ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersSelectedUserHooksUidGet(context.Background(), selectedUser, uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersSelectedUserHooksUidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserHooksUidGet`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersSelectedUserHooksUidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 
**uid** | **string** | Installed webhook&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersSelectedUserHooksUidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WebhookSubscription**](WebhookSubscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSelectedUserHooksUidPut

> WebhookSubscription UsersSelectedUserHooksUidPut(ctx, selectedUser, uid).Execute()





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
    uid := "uid_example" // string | Installed webhook's ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersSelectedUserHooksUidPut(context.Background(), selectedUser, uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersSelectedUserHooksUidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserHooksUidPut`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersSelectedUserHooksUidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 
**uid** | **string** | Installed webhook&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersSelectedUserHooksUidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WebhookSubscription**](WebhookSubscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsernameMembersGet

> User UsersUsernameMembersGet(ctx, username).Execute()





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
    username := "username_example" // string | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`. An account is either a team or user. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersUsernameMembersGet(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsernameMembersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsernameMembersGet`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsernameMembersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsernameMembersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersWorkspaceRepositoriesGet

> ModelError UsersWorkspaceRepositoriesGet(ctx, workspace).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersWorkspaceRepositoriesGet(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersWorkspaceRepositoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersWorkspaceRepositoriesGet`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersWorkspaceRepositoriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersWorkspaceRepositoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelError**](ModelError.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

