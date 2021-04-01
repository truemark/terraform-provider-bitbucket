# \PropertiesApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCommitHostedPropertyValue**](PropertiesApi.md#DeleteCommitHostedPropertyValue) | **Delete** /repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name} | 
[**DeletePullRequestHostedPropertyValue**](PropertiesApi.md#DeletePullRequestHostedPropertyValue) | **Delete** /repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name} | 
[**DeleteRepositoryHostedPropertyValue**](PropertiesApi.md#DeleteRepositoryHostedPropertyValue) | **Delete** /repositories/{workspace}/{repo_slug}/properties/{app_key}/{property_name} | 
[**DeleteUserHostedPropertyValue**](PropertiesApi.md#DeleteUserHostedPropertyValue) | **Delete** /users/{selected_user}/properties/{app_key}/{property_name} | 
[**GetCommitHostedPropertyValue**](PropertiesApi.md#GetCommitHostedPropertyValue) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name} | 
[**GetPullRequestHostedPropertyValue**](PropertiesApi.md#GetPullRequestHostedPropertyValue) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name} | 
[**GetRepositoryHostedPropertyValue**](PropertiesApi.md#GetRepositoryHostedPropertyValue) | **Get** /repositories/{workspace}/{repo_slug}/properties/{app_key}/{property_name} | 
[**RetrieveUserHostedPropertyValue**](PropertiesApi.md#RetrieveUserHostedPropertyValue) | **Get** /users/{selected_user}/properties/{app_key}/{property_name} | 
[**UpdateCommitHostedPropertyValue**](PropertiesApi.md#UpdateCommitHostedPropertyValue) | **Put** /repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name} | 
[**UpdatePullRequestHostedPropertyValue**](PropertiesApi.md#UpdatePullRequestHostedPropertyValue) | **Put** /repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name} | 
[**UpdateRepositoryHostedPropertyValue**](PropertiesApi.md#UpdateRepositoryHostedPropertyValue) | **Put** /repositories/{workspace}/{repo_slug}/properties/{app_key}/{property_name} | 
[**UpdateUserHostedPropertyValue**](PropertiesApi.md#UpdateUserHostedPropertyValue) | **Put** /users/{selected_user}/properties/{app_key}/{property_name} | 



## DeleteCommitHostedPropertyValue

> DeleteCommitHostedPropertyValue(ctx, workspace, repoSlug, commit, appKey, propertyName).Execute()





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
    workspace := "workspace_example" // string | The repository container; either the workspace slug or the UUID in curly braces.
    repoSlug := "repoSlug_example" // string | The repository.
    commit := "commit_example" // string | The commit.
    appKey := "appKey_example" // string | The key of the Connect app.
    propertyName := "propertyName_example" // string | The name of the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PropertiesApi.DeleteCommitHostedPropertyValue(context.Background(), workspace, repoSlug, commit, appKey, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesApi.DeleteCommitHostedPropertyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The repository container; either the workspace slug or the UUID in curly braces. | 
**repoSlug** | **string** | The repository. | 
**commit** | **string** | The commit. | 
**appKey** | **string** | The key of the Connect app. | 
**propertyName** | **string** | The name of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommitHostedPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePullRequestHostedPropertyValue

> DeletePullRequestHostedPropertyValue(ctx, workspace, repoSlug, pullrequestId, appKey, propertyName).Execute()





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
    workspace := "workspace_example" // string | The repository container; either the workspace slug or the UUID in curly braces.
    repoSlug := "repoSlug_example" // string | The repository.
    pullrequestId := "pullrequestId_example" // string | The pull request ID.
    appKey := "appKey_example" // string | The key of the Connect app.
    propertyName := "propertyName_example" // string | The name of the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PropertiesApi.DeletePullRequestHostedPropertyValue(context.Background(), workspace, repoSlug, pullrequestId, appKey, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesApi.DeletePullRequestHostedPropertyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The repository container; either the workspace slug or the UUID in curly braces. | 
**repoSlug** | **string** | The repository. | 
**pullrequestId** | **string** | The pull request ID. | 
**appKey** | **string** | The key of the Connect app. | 
**propertyName** | **string** | The name of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePullRequestHostedPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRepositoryHostedPropertyValue

> DeleteRepositoryHostedPropertyValue(ctx, workspace, repoSlug, appKey, propertyName).Execute()





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
    workspace := "workspace_example" // string | The repository container; either the workspace slug or the UUID in curly braces.
    repoSlug := "repoSlug_example" // string | The repository.
    appKey := "appKey_example" // string | The key of the Connect app.
    propertyName := "propertyName_example" // string | The name of the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PropertiesApi.DeleteRepositoryHostedPropertyValue(context.Background(), workspace, repoSlug, appKey, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesApi.DeleteRepositoryHostedPropertyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The repository container; either the workspace slug or the UUID in curly braces. | 
**repoSlug** | **string** | The repository. | 
**appKey** | **string** | The key of the Connect app. | 
**propertyName** | **string** | The name of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryHostedPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserHostedPropertyValue

> DeleteUserHostedPropertyValue(ctx, selectedUser, appKey, propertyName).Execute()





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
    selectedUser := "selectedUser_example" // string | Either the UUID of the account surrounded by curly-braces, for example `{account UUID}`, OR an Atlassian Account ID.
    appKey := "appKey_example" // string | The key of the Connect app.
    propertyName := "propertyName_example" // string | The name of the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PropertiesApi.DeleteUserHostedPropertyValue(context.Background(), selectedUser, appKey, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesApi.DeleteUserHostedPropertyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
**appKey** | **string** | The key of the Connect app. | 
**propertyName** | **string** | The name of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserHostedPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommitHostedPropertyValue

> GetCommitHostedPropertyValue(ctx, workspace, repoSlug, commit, appKey, propertyName).Execute()





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
    workspace := "workspace_example" // string | The repository container; either the workspace slug or the UUID in curly braces.
    repoSlug := "repoSlug_example" // string | The repository.
    commit := "commit_example" // string | The commit.
    appKey := "appKey_example" // string | The key of the Connect app.
    propertyName := "propertyName_example" // string | The name of the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PropertiesApi.GetCommitHostedPropertyValue(context.Background(), workspace, repoSlug, commit, appKey, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesApi.GetCommitHostedPropertyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The repository container; either the workspace slug or the UUID in curly braces. | 
**repoSlug** | **string** | The repository. | 
**commit** | **string** | The commit. | 
**appKey** | **string** | The key of the Connect app. | 
**propertyName** | **string** | The name of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitHostedPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPullRequestHostedPropertyValue

> GetPullRequestHostedPropertyValue(ctx, workspace, repoSlug, pullrequestId, appKey, propertyName).Execute()





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
    workspace := "workspace_example" // string | The repository container; either the workspace slug or the UUID in curly braces.
    repoSlug := "repoSlug_example" // string | The repository.
    pullrequestId := "pullrequestId_example" // string | The pull request ID.
    appKey := "appKey_example" // string | The key of the Connect app.
    propertyName := "propertyName_example" // string | The name of the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PropertiesApi.GetPullRequestHostedPropertyValue(context.Background(), workspace, repoSlug, pullrequestId, appKey, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesApi.GetPullRequestHostedPropertyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The repository container; either the workspace slug or the UUID in curly braces. | 
**repoSlug** | **string** | The repository. | 
**pullrequestId** | **string** | The pull request ID. | 
**appKey** | **string** | The key of the Connect app. | 
**propertyName** | **string** | The name of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequestHostedPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryHostedPropertyValue

> GetRepositoryHostedPropertyValue(ctx, workspace, repoSlug, appKey, propertyName).Execute()





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
    workspace := "workspace_example" // string | The repository container; either the workspace slug or the UUID in curly braces.
    repoSlug := "repoSlug_example" // string | The repository.
    appKey := "appKey_example" // string | The key of the Connect app.
    propertyName := "propertyName_example" // string | The name of the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PropertiesApi.GetRepositoryHostedPropertyValue(context.Background(), workspace, repoSlug, appKey, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesApi.GetRepositoryHostedPropertyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The repository container; either the workspace slug or the UUID in curly braces. | 
**repoSlug** | **string** | The repository. | 
**appKey** | **string** | The key of the Connect app. | 
**propertyName** | **string** | The name of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryHostedPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveUserHostedPropertyValue

> RetrieveUserHostedPropertyValue(ctx, selectedUser, appKey, propertyName).Execute()





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
    selectedUser := "selectedUser_example" // string | Either the UUID of the account surrounded by curly-braces, for example `{account UUID}`, OR an Atlassian Account ID.
    appKey := "appKey_example" // string | The key of the Connect app.
    propertyName := "propertyName_example" // string | The name of the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PropertiesApi.RetrieveUserHostedPropertyValue(context.Background(), selectedUser, appKey, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesApi.RetrieveUserHostedPropertyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
**appKey** | **string** | The key of the Connect app. | 
**propertyName** | **string** | The name of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUserHostedPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCommitHostedPropertyValue

> UpdateCommitHostedPropertyValue(ctx, workspace, repoSlug, commit, appKey, propertyName).Execute()





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
    workspace := "workspace_example" // string | The repository container; either the workspace slug or the UUID in curly braces.
    repoSlug := "repoSlug_example" // string | The repository.
    commit := "commit_example" // string | The commit.
    appKey := "appKey_example" // string | The key of the Connect app.
    propertyName := "propertyName_example" // string | The name of the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PropertiesApi.UpdateCommitHostedPropertyValue(context.Background(), workspace, repoSlug, commit, appKey, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesApi.UpdateCommitHostedPropertyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The repository container; either the workspace slug or the UUID in curly braces. | 
**repoSlug** | **string** | The repository. | 
**commit** | **string** | The commit. | 
**appKey** | **string** | The key of the Connect app. | 
**propertyName** | **string** | The name of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommitHostedPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePullRequestHostedPropertyValue

> UpdatePullRequestHostedPropertyValue(ctx, workspace, repoSlug, pullrequestId, appKey, propertyName).Execute()





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
    workspace := "workspace_example" // string | The repository container; either the workspace slug or the UUID in curly braces.
    repoSlug := "repoSlug_example" // string | The repository.
    pullrequestId := "pullrequestId_example" // string | The pull request ID.
    appKey := "appKey_example" // string | The key of the Connect app.
    propertyName := "propertyName_example" // string | The name of the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PropertiesApi.UpdatePullRequestHostedPropertyValue(context.Background(), workspace, repoSlug, pullrequestId, appKey, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesApi.UpdatePullRequestHostedPropertyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The repository container; either the workspace slug or the UUID in curly braces. | 
**repoSlug** | **string** | The repository. | 
**pullrequestId** | **string** | The pull request ID. | 
**appKey** | **string** | The key of the Connect app. | 
**propertyName** | **string** | The name of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePullRequestHostedPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepositoryHostedPropertyValue

> UpdateRepositoryHostedPropertyValue(ctx, workspace, repoSlug, appKey, propertyName).Execute()





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
    workspace := "workspace_example" // string | The repository container; either the workspace slug or the UUID in curly braces.
    repoSlug := "repoSlug_example" // string | The repository.
    appKey := "appKey_example" // string | The key of the Connect app.
    propertyName := "propertyName_example" // string | The name of the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PropertiesApi.UpdateRepositoryHostedPropertyValue(context.Background(), workspace, repoSlug, appKey, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesApi.UpdateRepositoryHostedPropertyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The repository container; either the workspace slug or the UUID in curly braces. | 
**repoSlug** | **string** | The repository. | 
**appKey** | **string** | The key of the Connect app. | 
**propertyName** | **string** | The name of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryHostedPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserHostedPropertyValue

> UpdateUserHostedPropertyValue(ctx, selectedUser, appKey, propertyName).Execute()





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
    selectedUser := "selectedUser_example" // string | Either the UUID of the account surrounded by curly-braces, for example `{account UUID}`, OR an Atlassian Account ID.
    appKey := "appKey_example" // string | The key of the Connect app.
    propertyName := "propertyName_example" // string | The name of the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PropertiesApi.UpdateUserHostedPropertyValue(context.Background(), selectedUser, appKey, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesApi.UpdateUserHostedPropertyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
**appKey** | **string** | The key of the Connect app. | 
**propertyName** | **string** | The name of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserHostedPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

