# \DeploymentsApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironment**](DeploymentsApi.md#CreateEnvironment) | **Post** /repositories/{workspace}/{repo_slug}/environments/ | 
[**DeleteEnvironmentForRepository**](DeploymentsApi.md#DeleteEnvironmentForRepository) | **Delete** /repositories/{workspace}/{repo_slug}/environments/{environment_uuid} | 
[**GetDeploymentForRepository**](DeploymentsApi.md#GetDeploymentForRepository) | **Get** /repositories/{workspace}/{repo_slug}/deployments/{deployment_uuid} | 
[**GetDeploymentsForRepository**](DeploymentsApi.md#GetDeploymentsForRepository) | **Get** /repositories/{workspace}/{repo_slug}/deployments/ | 
[**GetEnvironmentForRepository**](DeploymentsApi.md#GetEnvironmentForRepository) | **Get** /repositories/{workspace}/{repo_slug}/environments/{environment_uuid} | 
[**GetEnvironmentsForRepository**](DeploymentsApi.md#GetEnvironmentsForRepository) | **Get** /repositories/{workspace}/{repo_slug}/environments/ | 
[**RepositoriesWorkspaceRepoSlugDeployKeysGet**](DeploymentsApi.md#RepositoriesWorkspaceRepoSlugDeployKeysGet) | **Get** /repositories/{workspace}/{repo_slug}/deploy-keys | 
[**RepositoriesWorkspaceRepoSlugDeployKeysKeyIdDelete**](DeploymentsApi.md#RepositoriesWorkspaceRepoSlugDeployKeysKeyIdDelete) | **Delete** /repositories/{workspace}/{repo_slug}/deploy-keys/{key_id} | 
[**RepositoriesWorkspaceRepoSlugDeployKeysKeyIdGet**](DeploymentsApi.md#RepositoriesWorkspaceRepoSlugDeployKeysKeyIdGet) | **Get** /repositories/{workspace}/{repo_slug}/deploy-keys/{key_id} | 
[**RepositoriesWorkspaceRepoSlugDeployKeysKeyIdPut**](DeploymentsApi.md#RepositoriesWorkspaceRepoSlugDeployKeysKeyIdPut) | **Put** /repositories/{workspace}/{repo_slug}/deploy-keys/{key_id} | 
[**RepositoriesWorkspaceRepoSlugDeployKeysPost**](DeploymentsApi.md#RepositoriesWorkspaceRepoSlugDeployKeysPost) | **Post** /repositories/{workspace}/{repo_slug}/deploy-keys | 
[**UpdateEnvironmentForRepository**](DeploymentsApi.md#UpdateEnvironmentForRepository) | **Post** /repositories/{workspace}/{repo_slug}/environments/{environment_uuid}/changes/ | 



## CreateEnvironment

> DeploymentEnvironment CreateEnvironment(ctx, workspace, repoSlug).Body(body).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    body := *openapiclient.NewDeploymentEnvironment("Type_example") // DeploymentEnvironment | The environment to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.CreateEnvironment(context.Background(), workspace, repoSlug).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.CreateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironment`: DeploymentEnvironment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.CreateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DeploymentEnvironment**](DeploymentEnvironment.md) | The environment to create. | 

### Return type

[**DeploymentEnvironment**](DeploymentEnvironment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentForRepository

> DeleteEnvironmentForRepository(ctx, workspace, repoSlug, environmentUuid).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    environmentUuid := "environmentUuid_example" // string | The environment UUID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.DeleteEnvironmentForRepository(context.Background(), workspace, repoSlug, environmentUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.DeleteEnvironmentForRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**environmentUuid** | **string** | The environment UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentForRepository

> Deployment GetDeploymentForRepository(ctx, workspace, repoSlug, deploymentUuid).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    deploymentUuid := "deploymentUuid_example" // string | The deployment UUID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetDeploymentForRepository(context.Background(), workspace, repoSlug, deploymentUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeploymentForRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentForRepository`: Deployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeploymentForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**deploymentUuid** | **string** | The deployment UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentsForRepository

> PaginatedDeployments GetDeploymentsForRepository(ctx, workspace, repoSlug).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetDeploymentsForRepository(context.Background(), workspace, repoSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeploymentsForRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentsForRepository`: PaginatedDeployments
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeploymentsForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentsForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedDeployments**](PaginatedDeployments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentForRepository

> DeploymentEnvironment GetEnvironmentForRepository(ctx, workspace, repoSlug, environmentUuid).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    environmentUuid := "environmentUuid_example" // string | The environment UUID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetEnvironmentForRepository(context.Background(), workspace, repoSlug, environmentUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetEnvironmentForRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentForRepository`: DeploymentEnvironment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetEnvironmentForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**environmentUuid** | **string** | The environment UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeploymentEnvironment**](DeploymentEnvironment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentsForRepository

> PaginatedEnvironments GetEnvironmentsForRepository(ctx, workspace, repoSlug).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetEnvironmentsForRepository(context.Background(), workspace, repoSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetEnvironmentsForRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentsForRepository`: PaginatedEnvironments
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetEnvironmentsForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentsForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedEnvironments**](PaginatedEnvironments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugDeployKeysGet

> PaginatedDeployKeys RepositoriesWorkspaceRepoSlugDeployKeysGet(ctx, repoSlug, workspace).Execute()





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
    resp, r, err := api_client.DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysGet(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugDeployKeysGet`: PaginatedDeployKeys
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDeployKeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedDeployKeys**](PaginatedDeployKeys.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugDeployKeysKeyIdDelete

> RepositoriesWorkspaceRepoSlugDeployKeysKeyIdDelete(ctx, keyId, repoSlug, workspace).Execute()





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
    keyId := "keyId_example" // string | The key ID matching the deploy key.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysKeyIdDelete(context.Background(), keyId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysKeyIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The key ID matching the deploy key. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDeployKeysKeyIdDeleteRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugDeployKeysKeyIdGet

> DeployKey RepositoriesWorkspaceRepoSlugDeployKeysKeyIdGet(ctx, keyId, repoSlug, workspace).Execute()





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
    keyId := "keyId_example" // string | The key ID matching the deploy key.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysKeyIdGet(context.Background(), keyId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysKeyIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugDeployKeysKeyIdGet`: DeployKey
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysKeyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The key ID matching the deploy key. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDeployKeysKeyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeployKey**](DeployKey.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugDeployKeysKeyIdPut

> DeployKey RepositoriesWorkspaceRepoSlugDeployKeysKeyIdPut(ctx, keyId, repoSlug, workspace).Execute()





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
    keyId := "keyId_example" // string | The key ID matching the deploy key.
    repoSlug := "repoSlug_example" // string | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysKeyIdPut(context.Background(), keyId, repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysKeyIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugDeployKeysKeyIdPut`: DeployKey
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysKeyIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The key ID matching the deploy key. | 
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDeployKeysKeyIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeployKey**](DeployKey.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesWorkspaceRepoSlugDeployKeysPost

> DeployKey RepositoriesWorkspaceRepoSlugDeployKeysPost(ctx, repoSlug, workspace).Execute()





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
    resp, r, err := api_client.DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysPost(context.Background(), repoSlug, workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesWorkspaceRepoSlugDeployKeysPost`: DeployKey
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.RepositoriesWorkspaceRepoSlugDeployKeysPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugDeployKeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeployKey**](DeployKey.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironmentForRepository

> UpdateEnvironmentForRepository(ctx, workspace, repoSlug, environmentUuid).Execute()





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
    workspace := "workspace_example" // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.
    repoSlug := "repoSlug_example" // string | The repository.
    environmentUuid := "environmentUuid_example" // string | The environment UUID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.UpdateEnvironmentForRepository(context.Background(), workspace, repoSlug, environmentUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.UpdateEnvironmentForRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**environmentUuid** | **string** | The environment UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

