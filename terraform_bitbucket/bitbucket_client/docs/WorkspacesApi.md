# \WorkspacesApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserPermissionsWorkspacesGet**](WorkspacesApi.md#UserPermissionsWorkspacesGet) | **Get** /user/permissions/workspaces | 
[**WorkspacesGet**](WorkspacesApi.md#WorkspacesGet) | **Get** /workspaces | 
[**WorkspacesWorkspaceGet**](WorkspacesApi.md#WorkspacesWorkspaceGet) | **Get** /workspaces/{workspace} | 
[**WorkspacesWorkspaceHooksGet**](WorkspacesApi.md#WorkspacesWorkspaceHooksGet) | **Get** /workspaces/{workspace}/hooks | 
[**WorkspacesWorkspaceHooksPost**](WorkspacesApi.md#WorkspacesWorkspaceHooksPost) | **Post** /workspaces/{workspace}/hooks | 
[**WorkspacesWorkspaceHooksUidDelete**](WorkspacesApi.md#WorkspacesWorkspaceHooksUidDelete) | **Delete** /workspaces/{workspace}/hooks/{uid} | 
[**WorkspacesWorkspaceHooksUidGet**](WorkspacesApi.md#WorkspacesWorkspaceHooksUidGet) | **Get** /workspaces/{workspace}/hooks/{uid} | 
[**WorkspacesWorkspaceHooksUidPut**](WorkspacesApi.md#WorkspacesWorkspaceHooksUidPut) | **Put** /workspaces/{workspace}/hooks/{uid} | 
[**WorkspacesWorkspaceMembersGet**](WorkspacesApi.md#WorkspacesWorkspaceMembersGet) | **Get** /workspaces/{workspace}/members | 
[**WorkspacesWorkspaceMembersMemberGet**](WorkspacesApi.md#WorkspacesWorkspaceMembersMemberGet) | **Get** /workspaces/{workspace}/members/{member} | 
[**WorkspacesWorkspacePermissionsGet**](WorkspacesApi.md#WorkspacesWorkspacePermissionsGet) | **Get** /workspaces/{workspace}/permissions | 
[**WorkspacesWorkspacePermissionsRepositoriesGet**](WorkspacesApi.md#WorkspacesWorkspacePermissionsRepositoriesGet) | **Get** /workspaces/{workspace}/permissions/repositories | 
[**WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet**](WorkspacesApi.md#WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet) | **Get** /workspaces/{workspace}/permissions/repositories/{repo_slug} | 
[**WorkspacesWorkspaceProjectsGet**](WorkspacesApi.md#WorkspacesWorkspaceProjectsGet) | **Get** /workspaces/{workspace}/projects | 
[**WorkspacesWorkspaceProjectsProjectKeyGet**](WorkspacesApi.md#WorkspacesWorkspaceProjectsProjectKeyGet) | **Get** /workspaces/{workspace}/projects/{project_key} | 



## UserPermissionsWorkspacesGet

> PaginatedWorkspaceMemberships UserPermissionsWorkspacesGet(ctx).Q(q).Sort(sort).Execute()





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
    q := "q_example" // string |  Query string to narrow down the response. See [filtering and sorting](../../../meta/filtering) for details. (optional)
    sort := "sort_example" // string |  Name of a response property to sort results. See [filtering and sorting](../../../meta/filtering#query-sort) for details.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesApi.UserPermissionsWorkspacesGet(context.Background()).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.UserPermissionsWorkspacesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserPermissionsWorkspacesGet`: PaginatedWorkspaceMemberships
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.UserPermissionsWorkspacesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserPermissionsWorkspacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  Query string to narrow down the response. See [filtering and sorting](../../../meta/filtering) for details. | 
 **sort** | **string** |  Name of a response property to sort results. See [filtering and sorting](../../../meta/filtering#query-sort) for details.  | 

### Return type

[**PaginatedWorkspaceMemberships**](PaginatedWorkspaceMemberships.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesGet

> PaginatedWorkspaces WorkspacesGet(ctx).Role(role).Q(q).Sort(sort).Execute()





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
    role := "role_example" // string |              Filters the workspaces based on the authenticated user's role on each workspace.              * **member**: returns a list of all the workspaces which the caller is a member of                 at least one workspace group or repository             * **collaborator**: returns a list of workspaces which the caller has write access                 to at least one repository in the workspace             * **owner**: returns a list of workspaces which the caller has administrator access              (optional)
    q := "q_example" // string |  Query string to narrow down the response. See [filtering and sorting](../meta/filtering) for details. (optional)
    sort := "sort_example" // string |  Name of a response property to sort results. See [filtering and sorting](../meta/filtering#query-sort) for details.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesApi.WorkspacesGet(context.Background()).Role(role).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesGet`: PaginatedWorkspaces
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **string** |              Filters the workspaces based on the authenticated user&#39;s role on each workspace.              * **member**: returns a list of all the workspaces which the caller is a member of                 at least one workspace group or repository             * **collaborator**: returns a list of workspaces which the caller has write access                 to at least one repository in the workspace             * **owner**: returns a list of workspaces which the caller has administrator access              | 
 **q** | **string** |  Query string to narrow down the response. See [filtering and sorting](../meta/filtering) for details. | 
 **sort** | **string** |  Name of a response property to sort results. See [filtering and sorting](../meta/filtering#query-sort) for details.  | 

### Return type

[**PaginatedWorkspaces**](PaginatedWorkspaces.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceGet

> Workspace WorkspacesWorkspaceGet(ctx, workspace).Execute()





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
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspaceGet(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspaceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceGet`: Workspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesWorkspaceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Workspace**](Workspace.md)

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
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspaceHooksGet(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspaceHooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceHooksGet`: PaginatedWebhookSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesWorkspaceHooksGet`: %v\n", resp)
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
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspaceHooksPost(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspaceHooksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceHooksPost`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesWorkspaceHooksPost`: %v\n", resp)
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
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspaceHooksUidDelete(context.Background(), uid, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspaceHooksUidDelete``: %v\n", err)
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
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspaceHooksUidGet(context.Background(), uid, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspaceHooksUidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceHooksUidGet`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesWorkspaceHooksUidGet`: %v\n", resp)
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
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspaceHooksUidPut(context.Background(), uid, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspaceHooksUidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceHooksUidPut`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesWorkspaceHooksUidPut`: %v\n", resp)
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


## WorkspacesWorkspaceMembersGet

> PaginatedWorkspaceMemberships WorkspacesWorkspaceMembersGet(ctx, workspace).Execute()





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
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspaceMembersGet(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspaceMembersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceMembersGet`: PaginatedWorkspaceMemberships
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesWorkspaceMembersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceMembersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedWorkspaceMemberships**](PaginatedWorkspaceMemberships.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceMembersMemberGet

> WorkspaceMembership WorkspacesWorkspaceMembersMemberGet(ctx, member, workspace).Execute()





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
    member := "member_example" // string | Member's UUID or Atlassian ID.
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspaceMembersMemberGet(context.Background(), member, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspaceMembersMemberGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceMembersMemberGet`: WorkspaceMembership
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesWorkspaceMembersMemberGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**member** | **string** | Member&#39;s UUID or Atlassian ID. | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceMembersMemberGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkspaceMembership**](WorkspaceMembership.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspacePermissionsGet

> PaginatedWorkspaceMemberships WorkspacesWorkspacePermissionsGet(ctx, workspace).Q(q).Execute()





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
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../../meta/filtering). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspacePermissionsGet(context.Background(), workspace).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspacePermissionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspacePermissionsGet`: PaginatedWorkspaceMemberships
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesWorkspacePermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspacePermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** |  Query string to narrow down the response as per [filtering and sorting](../../../meta/filtering). | 

### Return type

[**PaginatedWorkspaceMemberships**](PaginatedWorkspaceMemberships.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspacePermissionsRepositoriesGet

> PaginatedRepositoryPermissions WorkspacesWorkspacePermissionsRepositoriesGet(ctx, workspace).Q(q).Sort(sort).Execute()





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
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../../../meta/filtering). (optional)
    sort := "sort_example" // string |  Name of a response property sort the result by as per [filtering and sorting](../../../../meta/filtering#query-sort).  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspacePermissionsRepositoriesGet(context.Background(), workspace).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspacePermissionsRepositoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspacePermissionsRepositoriesGet`: PaginatedRepositoryPermissions
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesWorkspacePermissionsRepositoriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspacePermissionsRepositoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** |  Query string to narrow down the response as per [filtering and sorting](../../../../meta/filtering). | 
 **sort** | **string** |  Name of a response property sort the result by as per [filtering and sorting](../../../../meta/filtering#query-sort).  | 

### Return type

[**PaginatedRepositoryPermissions**](PaginatedRepositoryPermissions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet

> PaginatedRepositoryPermissions WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet(ctx, repoSlug, workspace).Q(q).Sort(sort).Execute()





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
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../../../meta/filtering). (optional)
    sort := "sort_example" // string |  Name of a response property sort the result by as per [filtering and sorting](../../../../meta/filtering#query-sort).  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet(context.Background(), repoSlug, workspace).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet`: PaginatedRepositoryPermissions
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspacePermissionsRepositoriesRepoSlugGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** |  Query string to narrow down the response as per [filtering and sorting](../../../../meta/filtering). | 
 **sort** | **string** |  Name of a response property sort the result by as per [filtering and sorting](../../../../meta/filtering#query-sort).  | 

### Return type

[**PaginatedRepositoryPermissions**](PaginatedRepositoryPermissions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceProjectsGet

> PaginatedProjects WorkspacesWorkspaceProjectsGet(ctx, workspace).Execute()





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
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspaceProjectsGet(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspaceProjectsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceProjectsGet`: PaginatedProjects
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesWorkspaceProjectsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedProjects**](PaginatedProjects.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceProjectsProjectKeyGet

> Project WorkspacesWorkspaceProjectsProjectKeyGet(ctx, projectKey, workspace).Execute()





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
    projectKey := "projectKey_example" // string | The project in question. This is the actual `key` assigned to the project. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkspacesApi.WorkspacesWorkspaceProjectsProjectKeyGet(context.Background(), projectKey, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.WorkspacesWorkspaceProjectsProjectKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceProjectsProjectKeyGet`: Project
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.WorkspacesWorkspaceProjectsProjectKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project in question. This is the actual &#x60;key&#x60; assigned to the project.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceProjectsProjectKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

