# \WebhooksApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HookEventsGet**](WebhooksApi.md#HookEventsGet) | **Get** /hook_events | 
[**HookEventsSubjectTypeGet**](WebhooksApi.md#HookEventsSubjectTypeGet) | **Get** /hook_events/{subject_type} | 
[**RepositoriesWorkspaceRepoSlugHooksGet**](WebhooksApi.md#RepositoriesWorkspaceRepoSlugHooksGet) | **Get** /repositories/{workspace}/{repo_slug}/hooks | 
[**RepositoriesWorkspaceRepoSlugHooksPost**](WebhooksApi.md#RepositoriesWorkspaceRepoSlugHooksPost) | **Post** /repositories/{workspace}/{repo_slug}/hooks | 
[**RepositoriesWorkspaceRepoSlugHooksUidDelete**](WebhooksApi.md#RepositoriesWorkspaceRepoSlugHooksUidDelete) | **Delete** /repositories/{workspace}/{repo_slug}/hooks/{uid} | 
[**RepositoriesWorkspaceRepoSlugHooksUidGet**](WebhooksApi.md#RepositoriesWorkspaceRepoSlugHooksUidGet) | **Get** /repositories/{workspace}/{repo_slug}/hooks/{uid} | 
[**RepositoriesWorkspaceRepoSlugHooksUidPut**](WebhooksApi.md#RepositoriesWorkspaceRepoSlugHooksUidPut) | **Put** /repositories/{workspace}/{repo_slug}/hooks/{uid} | 
[**TeamsUsernameHooksGet**](WebhooksApi.md#TeamsUsernameHooksGet) | **Get** /teams/{username}/hooks | 
[**TeamsUsernameHooksPost**](WebhooksApi.md#TeamsUsernameHooksPost) | **Post** /teams/{username}/hooks | 
[**TeamsUsernameHooksUidDelete**](WebhooksApi.md#TeamsUsernameHooksUidDelete) | **Delete** /teams/{username}/hooks/{uid} | 
[**TeamsUsernameHooksUidGet**](WebhooksApi.md#TeamsUsernameHooksUidGet) | **Get** /teams/{username}/hooks/{uid} | 
[**TeamsUsernameHooksUidPut**](WebhooksApi.md#TeamsUsernameHooksUidPut) | **Put** /teams/{username}/hooks/{uid} | 
[**UsersSelectedUserHooksGet**](WebhooksApi.md#UsersSelectedUserHooksGet) | **Get** /users/{selected_user}/hooks | 
[**UsersSelectedUserHooksPost**](WebhooksApi.md#UsersSelectedUserHooksPost) | **Post** /users/{selected_user}/hooks | 
[**UsersSelectedUserHooksUidDelete**](WebhooksApi.md#UsersSelectedUserHooksUidDelete) | **Delete** /users/{selected_user}/hooks/{uid} | 
[**UsersSelectedUserHooksUidGet**](WebhooksApi.md#UsersSelectedUserHooksUidGet) | **Get** /users/{selected_user}/hooks/{uid} | 
[**UsersSelectedUserHooksUidPut**](WebhooksApi.md#UsersSelectedUserHooksUidPut) | **Put** /users/{selected_user}/hooks/{uid} | 
[**WorkspacesWorkspaceHooksGet**](WebhooksApi.md#WorkspacesWorkspaceHooksGet) | **Get** /workspaces/{workspace}/hooks | 
[**WorkspacesWorkspaceHooksPost**](WebhooksApi.md#WorkspacesWorkspaceHooksPost) | **Post** /workspaces/{workspace}/hooks | 
[**WorkspacesWorkspaceHooksUidDelete**](WebhooksApi.md#WorkspacesWorkspaceHooksUidDelete) | **Delete** /workspaces/{workspace}/hooks/{uid} | 
[**WorkspacesWorkspaceHooksUidGet**](WebhooksApi.md#WorkspacesWorkspaceHooksUidGet) | **Get** /workspaces/{workspace}/hooks/{uid} | 
[**WorkspacesWorkspaceHooksUidPut**](WebhooksApi.md#WorkspacesWorkspaceHooksUidPut) | **Put** /workspaces/{workspace}/hooks/{uid} | 



## HookEventsGet

> SubjectTypes HookEventsGet(ctx).Execute()





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
    resp, r, err := api_client.WebhooksApi.HookEventsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.HookEventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HookEventsGet`: SubjectTypes
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.HookEventsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHookEventsGetRequest struct via the builder pattern


### Return type

[**SubjectTypes**](SubjectTypes.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HookEventsSubjectTypeGet

> PaginatedHookEvents HookEventsSubjectTypeGet(ctx, subjectType).Execute()





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
    subjectType := "subjectType_example" // string | A resource or subject type.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.HookEventsSubjectTypeGet(context.Background(), subjectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.HookEventsSubjectTypeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HookEventsSubjectTypeGet`: PaginatedHookEvents
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.HookEventsSubjectTypeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectType** | **string** | A resource or subject type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHookEventsSubjectTypeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedHookEvents**](PaginatedHookEvents.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugHooksGet

> PaginatedWebhookSubscriptions RepositoriesWorkspaceRepoSlugHooksGet(ctx, repoSlug, workspace).Execute()





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
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.RepositoriesWorkspaceRepoSlugHooksGet(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.RepositoriesWorkspaceRepoSlugHooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugHooksGet`: PaginatedWebhookSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.RepositoriesWorkspaceRepoSlugHooksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugHooksGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugHooksPost

> WebhookSubscription RepositoriesWorkspaceRepoSlugHooksPost(ctx, repoSlug, workspace).Execute()





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
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.RepositoriesWorkspaceRepoSlugHooksPost(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.RepositoriesWorkspaceRepoSlugHooksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugHooksPost`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.RepositoriesWorkspaceRepoSlugHooksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugHooksPostRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugHooksUidDelete

> RepositoriesWorkspaceRepoSlugHooksUidDelete(ctx, repoSlug, uid, workspace).Execute()





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
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    uid := "uid_example" // string | Installed webhook's ID
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.RepositoriesWorkspaceRepoSlugHooksUidDelete(context.Background(), repoSlug, uid, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.RepositoriesWorkspaceRepoSlugHooksUidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**uid** | **string** | Installed webhook&#39;s ID | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugHooksUidDeleteRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugHooksUidGet

> WebhookSubscription RepositoriesWorkspaceRepoSlugHooksUidGet(ctx, repoSlug, uid, workspace).Execute()





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
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    uid := "uid_example" // string | Installed webhook's ID
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.RepositoriesWorkspaceRepoSlugHooksUidGet(context.Background(), repoSlug, uid, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.RepositoriesWorkspaceRepoSlugHooksUidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugHooksUidGet`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.RepositoriesWorkspaceRepoSlugHooksUidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**uid** | **string** | Installed webhook&#39;s ID | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugHooksUidGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugHooksUidPut

> WebhookSubscription RepositoriesWorkspaceRepoSlugHooksUidPut(ctx, repoSlug, uid, workspace).Execute()





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
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    uid := "uid_example" // string | Installed webhook's ID
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.RepositoriesWorkspaceRepoSlugHooksUidPut(context.Background(), repoSlug, uid, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.RepositoriesWorkspaceRepoSlugHooksUidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugHooksUidPut`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.RepositoriesWorkspaceRepoSlugHooksUidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**uid** | **string** | Installed webhook&#39;s ID | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugHooksUidPutRequest struct via the builder pattern


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


## TeamsUsernameHooksGet

> PaginatedWebhookSubscriptions TeamsUsernameHooksGet(ctx, username).Execute()





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
    resp, r, err := api_client.WebhooksApi.TeamsUsernameHooksGet(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.TeamsUsernameHooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameHooksGet`: PaginatedWebhookSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.TeamsUsernameHooksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameHooksGetRequest struct via the builder pattern


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


## TeamsUsernameHooksPost

> WebhookSubscription TeamsUsernameHooksPost(ctx, username).Execute()





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
    resp, r, err := api_client.WebhooksApi.TeamsUsernameHooksPost(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.TeamsUsernameHooksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameHooksPost`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.TeamsUsernameHooksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameHooksPostRequest struct via the builder pattern


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


## TeamsUsernameHooksUidDelete

> TeamsUsernameHooksUidDelete(ctx, uid, username).Execute()





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
    uid := "uid_example" // string | Installed webhook's ID
    username := "username_example" // string | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`. An account is either a team or user. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.TeamsUsernameHooksUidDelete(context.Background(), uid, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.TeamsUsernameHooksUidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Installed webhook&#39;s ID | 
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameHooksUidDeleteRequest struct via the builder pattern


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


## TeamsUsernameHooksUidGet

> WebhookSubscription TeamsUsernameHooksUidGet(ctx, uid, username).Execute()





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
    uid := "uid_example" // string | Installed webhook's ID
    username := "username_example" // string | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`. An account is either a team or user. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.TeamsUsernameHooksUidGet(context.Background(), uid, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.TeamsUsernameHooksUidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameHooksUidGet`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.TeamsUsernameHooksUidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Installed webhook&#39;s ID | 
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameHooksUidGetRequest struct via the builder pattern


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


## TeamsUsernameHooksUidPut

> WebhookSubscription TeamsUsernameHooksUidPut(ctx, uid, username).Execute()





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
    uid := "uid_example" // string | Installed webhook's ID
    username := "username_example" // string | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`. An account is either a team or user. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.TeamsUsernameHooksUidPut(context.Background(), uid, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.TeamsUsernameHooksUidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameHooksUidPut`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.TeamsUsernameHooksUidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Installed webhook&#39;s ID | 
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameHooksUidPutRequest struct via the builder pattern


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
    resp, r, err := api_client.WebhooksApi.UsersSelectedUserHooksGet(context.Background(), selectedUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.UsersSelectedUserHooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserHooksGet`: PaginatedWebhookSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.UsersSelectedUserHooksGet`: %v\n", resp)
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
    resp, r, err := api_client.WebhooksApi.UsersSelectedUserHooksPost(context.Background(), selectedUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.UsersSelectedUserHooksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserHooksPost`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.UsersSelectedUserHooksPost`: %v\n", resp)
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
    resp, r, err := api_client.WebhooksApi.UsersSelectedUserHooksUidDelete(context.Background(), selectedUser, uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.UsersSelectedUserHooksUidDelete``: %v\n", err)
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
    resp, r, err := api_client.WebhooksApi.UsersSelectedUserHooksUidGet(context.Background(), selectedUser, uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.UsersSelectedUserHooksUidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserHooksUidGet`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.UsersSelectedUserHooksUidGet`: %v\n", resp)
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
    resp, r, err := api_client.WebhooksApi.UsersSelectedUserHooksUidPut(context.Background(), selectedUser, uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.UsersSelectedUserHooksUidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSelectedUserHooksUidPut`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.UsersSelectedUserHooksUidPut`: %v\n", resp)
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


## WorkspacesWorkspaceHooksGet

> PaginatedWebhookSubscriptions WorkspacesWorkspaceHooksGet(ctx, workspace).Execute()





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
    resp, r, err := api_client.WebhooksApi.WorkspacesWorkspaceHooksGet(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WorkspacesWorkspaceHooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceHooksGet`: PaginatedWebhookSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WorkspacesWorkspaceHooksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceHooksGetRequest struct via the builder pattern


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


## WorkspacesWorkspaceHooksPost

> WebhookSubscription WorkspacesWorkspaceHooksPost(ctx, workspace).Execute()





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
    resp, r, err := api_client.WebhooksApi.WorkspacesWorkspaceHooksPost(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WorkspacesWorkspaceHooksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceHooksPost`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WorkspacesWorkspaceHooksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceHooksPostRequest struct via the builder pattern


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


## WorkspacesWorkspaceHooksUidDelete

> WorkspacesWorkspaceHooksUidDelete(ctx, uid, workspace).Execute()





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
    uid := "uid_example" // string | Installed webhook's ID
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.WorkspacesWorkspaceHooksUidDelete(context.Background(), uid, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WorkspacesWorkspaceHooksUidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Installed webhook&#39;s ID | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceHooksUidDeleteRequest struct via the builder pattern


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


## WorkspacesWorkspaceHooksUidGet

> WebhookSubscription WorkspacesWorkspaceHooksUidGet(ctx, uid, workspace).Execute()





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
    uid := "uid_example" // string | Installed webhook's ID
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.WorkspacesWorkspaceHooksUidGet(context.Background(), uid, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WorkspacesWorkspaceHooksUidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceHooksUidGet`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WorkspacesWorkspaceHooksUidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Installed webhook&#39;s ID | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceHooksUidGetRequest struct via the builder pattern


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


## WorkspacesWorkspaceHooksUidPut

> WebhookSubscription WorkspacesWorkspaceHooksUidPut(ctx, uid, workspace).Execute()





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
    uid := "uid_example" // string | Installed webhook's ID
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.WorkspacesWorkspaceHooksUidPut(context.Background(), uid, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WorkspacesWorkspaceHooksUidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceHooksUidPut`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WorkspacesWorkspaceHooksUidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Installed webhook&#39;s ID | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceHooksUidPutRequest struct via the builder pattern


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

