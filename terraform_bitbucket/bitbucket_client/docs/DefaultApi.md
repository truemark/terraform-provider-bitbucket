# \DefaultApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidLogsLogUuidGet**](DefaultApi.md#RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidLogsLogUuidGet) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/logs/{log_uuid} | 
[**RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsGet**](DefaultApi.md#RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsGet) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/test_reports | 
[**RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesGet**](DefaultApi.md#RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesGet) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/test_reports/test_cases | 
[**RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesTestCaseUuidTestCaseReasonsGet**](DefaultApi.md#RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesTestCaseUuidTestCaseReasonsGet) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/test_reports/test_cases/{test_case_uuid}/test_case_reasons | 
[**WorkspacesWorkspacePipelinesConfigIdentityOidcKeysJsonGet**](DefaultApi.md#WorkspacesWorkspacePipelinesConfigIdentityOidcKeysJsonGet) | **Get** /workspaces/{workspace}/pipelines-config/identity/oidc/keys.json | 
[**WorkspacesWorkspacePipelinesConfigIdentityOidcWellKnownOpenidConfigurationGet**](DefaultApi.md#WorkspacesWorkspacePipelinesConfigIdentityOidcWellKnownOpenidConfigurationGet) | **Get** /workspaces/{workspace}/pipelines-config/identity/oidc/.well-known/openid-configuration | 



## RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidLogsLogUuidGet

> RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidLogsLogUuidGet(ctx).Execute()



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
    resp, r, err := api_client.DefaultApi.RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidLogsLogUuidGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidLogsLogUuidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidLogsLogUuidGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsGet

> RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsGet(ctx).Execute()



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
    resp, r, err := api_client.DefaultApi.RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesGet

> RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesGet(ctx).Execute()



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
    resp, r, err := api_client.DefaultApi.RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesGetRequest struct via the builder pattern


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


## RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesTestCaseUuidTestCaseReasonsGet

> RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesTestCaseUuidTestCaseReasonsGet(ctx).Execute()



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
    resp, r, err := api_client.DefaultApi.RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesTestCaseUuidTestCaseReasonsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesTestCaseUuidTestCaseReasonsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesWorkspaceRepoSlugPipelinesPipelineUuidStepsStepUuidTestReportsTestCasesTestCaseUuidTestCaseReasonsGetRequest struct via the builder pattern


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


## WorkspacesWorkspacePipelinesConfigIdentityOidcKeysJsonGet

> WorkspacesWorkspacePipelinesConfigIdentityOidcKeysJsonGet(ctx).Execute()



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
    resp, r, err := api_client.DefaultApi.WorkspacesWorkspacePipelinesConfigIdentityOidcKeysJsonGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.WorkspacesWorkspacePipelinesConfigIdentityOidcKeysJsonGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspacePipelinesConfigIdentityOidcKeysJsonGetRequest struct via the builder pattern


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


## WorkspacesWorkspacePipelinesConfigIdentityOidcWellKnownOpenidConfigurationGet

> WorkspacesWorkspacePipelinesConfigIdentityOidcWellKnownOpenidConfigurationGet(ctx).Execute()



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
    resp, r, err := api_client.DefaultApi.WorkspacesWorkspacePipelinesConfigIdentityOidcWellKnownOpenidConfigurationGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.WorkspacesWorkspacePipelinesConfigIdentityOidcWellKnownOpenidConfigurationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspacePipelinesConfigIdentityOidcWellKnownOpenidConfigurationGetRequest struct via the builder pattern


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

