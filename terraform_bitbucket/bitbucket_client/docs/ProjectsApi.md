# \ProjectsApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsUsernameProjectsGet**](ProjectsApi.md#TeamsUsernameProjectsGet) | **Get** /teams/{username}/projects/ | 
[**TeamsUsernameProjectsPost**](ProjectsApi.md#TeamsUsernameProjectsPost) | **Post** /teams/{username}/projects/ | 
[**TeamsUsernameProjectsProjectKeyDelete**](ProjectsApi.md#TeamsUsernameProjectsProjectKeyDelete) | **Delete** /teams/{username}/projects/{project_key} | 
[**TeamsUsernameProjectsProjectKeyGet**](ProjectsApi.md#TeamsUsernameProjectsProjectKeyGet) | **Get** /teams/{username}/projects/{project_key} | 
[**TeamsUsernameProjectsProjectKeyPut**](ProjectsApi.md#TeamsUsernameProjectsProjectKeyPut) | **Put** /teams/{username}/projects/{project_key} | 
[**WorkspacesWorkspaceProjectsPost**](ProjectsApi.md#WorkspacesWorkspaceProjectsPost) | **Post** /workspaces/{workspace}/projects | 
[**WorkspacesWorkspaceProjectsProjectKeyDelete**](ProjectsApi.md#WorkspacesWorkspaceProjectsProjectKeyDelete) | **Delete** /workspaces/{workspace}/projects/{project_key} | 
[**WorkspacesWorkspaceProjectsProjectKeyGet**](ProjectsApi.md#WorkspacesWorkspaceProjectsProjectKeyGet) | **Get** /workspaces/{workspace}/projects/{project_key} | 
[**WorkspacesWorkspaceProjectsProjectKeyPut**](ProjectsApi.md#WorkspacesWorkspaceProjectsProjectKeyPut) | **Put** /workspaces/{workspace}/projects/{project_key} | 



## TeamsUsernameProjectsGet

> PaginatedProjects TeamsUsernameProjectsGet(ctx, username).Execute()





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
    resp, r, err := api_client.ProjectsApi.TeamsUsernameProjectsGet(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.TeamsUsernameProjectsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameProjectsGet`: PaginatedProjects
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.TeamsUsernameProjectsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameProjectsGetRequest struct via the builder pattern


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


## TeamsUsernameProjectsPost

> Project TeamsUsernameProjectsPost(ctx, username).Body(body).Execute()





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
    body := *openapiclient.NewProject() // Project | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.TeamsUsernameProjectsPost(context.Background(), username).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.TeamsUsernameProjectsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameProjectsPost`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.TeamsUsernameProjectsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameProjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Project**](Project.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUsernameProjectsProjectKeyDelete

> TeamsUsernameProjectsProjectKeyDelete(ctx, projectKey, username).Execute()





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
    projectKey := "projectKey_example" // string | The project in question. This can either be the actual `key` assigned to the project or the `UUID` (surrounded by curly-braces (`{}`)). 
    username := "username_example" // string | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`. An account is either a team or user. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.TeamsUsernameProjectsProjectKeyDelete(context.Background(), projectKey, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.TeamsUsernameProjectsProjectKeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project in question. This can either be the actual &#x60;key&#x60; assigned to the project or the &#x60;UUID&#x60; (surrounded by curly-braces (&#x60;{}&#x60;)).  | 
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameProjectsProjectKeyDeleteRequest struct via the builder pattern


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


## TeamsUsernameProjectsProjectKeyGet

> Project TeamsUsernameProjectsProjectKeyGet(ctx, projectKey, username).Execute()





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
    projectKey := "projectKey_example" // string | The project in question. This can either be the actual `key` assigned to the project or the `UUID` (surrounded by curly-braces (`{}`)). 
    username := "username_example" // string | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`. An account is either a team or user. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.TeamsUsernameProjectsProjectKeyGet(context.Background(), projectKey, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.TeamsUsernameProjectsProjectKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameProjectsProjectKeyGet`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.TeamsUsernameProjectsProjectKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project in question. This can either be the actual &#x60;key&#x60; assigned to the project or the &#x60;UUID&#x60; (surrounded by curly-braces (&#x60;{}&#x60;)).  | 
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameProjectsProjectKeyGetRequest struct via the builder pattern


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


## TeamsUsernameProjectsProjectKeyPut

> Project TeamsUsernameProjectsProjectKeyPut(ctx, projectKey, username).Body(body).Execute()





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
    projectKey := "projectKey_example" // string | The project in question. This can either be the actual `key` assigned to the project or the `UUID` (surrounded by curly-braces (`{}`)). 
    username := "username_example" // string | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`. An account is either a team or user. 
    body := *openapiclient.NewProject() // Project | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.TeamsUsernameProjectsProjectKeyPut(context.Background(), projectKey, username).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.TeamsUsernameProjectsProjectKeyPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameProjectsProjectKeyPut`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.TeamsUsernameProjectsProjectKeyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project in question. This can either be the actual &#x60;key&#x60; assigned to the project or the &#x60;UUID&#x60; (surrounded by curly-braces (&#x60;{}&#x60;)).  | 
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameProjectsProjectKeyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Project**](Project.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceProjectsPost

> Project WorkspacesWorkspaceProjectsPost(ctx, workspace).Body(body).Execute()





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
    body := *openapiclient.NewProject() // Project | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.WorkspacesWorkspaceProjectsPost(context.Background(), workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.WorkspacesWorkspaceProjectsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceProjectsPost`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.WorkspacesWorkspaceProjectsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceProjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Project**](Project.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceProjectsProjectKeyDelete

> WorkspacesWorkspaceProjectsProjectKeyDelete(ctx, projectKey, workspace).Execute()





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
    resp, r, err := api_client.ProjectsApi.WorkspacesWorkspaceProjectsProjectKeyDelete(context.Background(), projectKey, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.WorkspacesWorkspaceProjectsProjectKeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project in question. This is the actual &#x60;key&#x60; assigned to the project.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceProjectsProjectKeyDeleteRequest struct via the builder pattern


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
    resp, r, err := api_client.ProjectsApi.WorkspacesWorkspaceProjectsProjectKeyGet(context.Background(), projectKey, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.WorkspacesWorkspaceProjectsProjectKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceProjectsProjectKeyGet`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.WorkspacesWorkspaceProjectsProjectKeyGet`: %v\n", resp)
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


## WorkspacesWorkspaceProjectsProjectKeyPut

> Project WorkspacesWorkspaceProjectsProjectKeyPut(ctx, projectKey, workspace).Body(body).Execute()





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
    body := *openapiclient.NewProject() // Project | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.WorkspacesWorkspaceProjectsProjectKeyPut(context.Background(), projectKey, workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.WorkspacesWorkspaceProjectsProjectKeyPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspacesWorkspaceProjectsProjectKeyPut`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.WorkspacesWorkspaceProjectsProjectKeyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project in question. This is the actual &#x60;key&#x60; assigned to the project.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceProjectsProjectKeyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Project**](Project.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

