# \PipelinesApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentVariable**](PipelinesApi.md#CreateDeploymentVariable) | **Post** /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables | 
[**CreatePipelineForRepository**](PipelinesApi.md#CreatePipelineForRepository) | **Post** /repositories/{workspace}/{repo_slug}/pipelines/ | 
[**CreatePipelineVariableForTeam**](PipelinesApi.md#CreatePipelineVariableForTeam) | **Post** /teams/{username}/pipelines_config/variables/ | 
[**CreatePipelineVariableForUser**](PipelinesApi.md#CreatePipelineVariableForUser) | **Post** /users/{selected_user}/pipelines_config/variables/ | 
[**CreatePipelineVariableForWorkspace**](PipelinesApi.md#CreatePipelineVariableForWorkspace) | **Post** /workspaces/{workspace}/pipelines-config/variables | 
[**CreateRepositoryPipelineKnownHost**](PipelinesApi.md#CreateRepositoryPipelineKnownHost) | **Post** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/known_hosts/ | 
[**CreateRepositoryPipelineSchedule**](PipelinesApi.md#CreateRepositoryPipelineSchedule) | **Post** /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/ | 
[**CreateRepositoryPipelineVariable**](PipelinesApi.md#CreateRepositoryPipelineVariable) | **Post** /repositories/{workspace}/{repo_slug}/pipelines_config/variables/ | 
[**DeleteDeploymentVariable**](PipelinesApi.md#DeleteDeploymentVariable) | **Delete** /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables/{variable_uuid} | 
[**DeletePipelineVariableForTeam**](PipelinesApi.md#DeletePipelineVariableForTeam) | **Delete** /teams/{username}/pipelines_config/variables/{variable_uuid} | 
[**DeletePipelineVariableForUser**](PipelinesApi.md#DeletePipelineVariableForUser) | **Delete** /users/{selected_user}/pipelines_config/variables/{variable_uuid} | 
[**DeletePipelineVariableForWorkspace**](PipelinesApi.md#DeletePipelineVariableForWorkspace) | **Delete** /workspaces/{workspace}/pipelines-config/variables/{variable_uuid} | 
[**DeleteRepositoryPipelineCache**](PipelinesApi.md#DeleteRepositoryPipelineCache) | **Delete** /repositories/{workspace}/{repo_slug}/pipelines-config/caches/{cache_uuid} | 
[**DeleteRepositoryPipelineKeyPair**](PipelinesApi.md#DeleteRepositoryPipelineKeyPair) | **Delete** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/key_pair | 
[**DeleteRepositoryPipelineKnownHost**](PipelinesApi.md#DeleteRepositoryPipelineKnownHost) | **Delete** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/known_hosts/{known_host_uuid} | 
[**DeleteRepositoryPipelineSchedule**](PipelinesApi.md#DeleteRepositoryPipelineSchedule) | **Delete** /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/{schedule_uuid} | 
[**DeleteRepositoryPipelineVariable**](PipelinesApi.md#DeleteRepositoryPipelineVariable) | **Delete** /repositories/{workspace}/{repo_slug}/pipelines_config/variables/{variable_uuid} | 
[**GetDeploymentVariables**](PipelinesApi.md#GetDeploymentVariables) | **Get** /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables | 
[**GetPipelineForRepository**](PipelinesApi.md#GetPipelineForRepository) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid} | 
[**GetPipelineStepForRepository**](PipelinesApi.md#GetPipelineStepForRepository) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid} | 
[**GetPipelineStepLogForRepository**](PipelinesApi.md#GetPipelineStepLogForRepository) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/log | 
[**GetPipelineStepsForRepository**](PipelinesApi.md#GetPipelineStepsForRepository) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/ | 
[**GetPipelineVariableForTeam**](PipelinesApi.md#GetPipelineVariableForTeam) | **Get** /teams/{username}/pipelines_config/variables/{variable_uuid} | 
[**GetPipelineVariableForUser**](PipelinesApi.md#GetPipelineVariableForUser) | **Get** /users/{selected_user}/pipelines_config/variables/{variable_uuid} | 
[**GetPipelineVariableForWorkspace**](PipelinesApi.md#GetPipelineVariableForWorkspace) | **Get** /workspaces/{workspace}/pipelines-config/variables/{variable_uuid} | 
[**GetPipelineVariablesForTeam**](PipelinesApi.md#GetPipelineVariablesForTeam) | **Get** /teams/{username}/pipelines_config/variables/ | 
[**GetPipelineVariablesForUser**](PipelinesApi.md#GetPipelineVariablesForUser) | **Get** /users/{selected_user}/pipelines_config/variables/ | 
[**GetPipelineVariablesForWorkspace**](PipelinesApi.md#GetPipelineVariablesForWorkspace) | **Get** /workspaces/{workspace}/pipelines-config/variables | 
[**GetPipelinesForRepository**](PipelinesApi.md#GetPipelinesForRepository) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/ | 
[**GetRepositoryPipelineCacheContentURI**](PipelinesApi.md#GetRepositoryPipelineCacheContentURI) | **Get** /repositories/{workspace}/{repo_slug}/pipelines-config/caches/{cache_uuid}/content-uri | 
[**GetRepositoryPipelineCaches**](PipelinesApi.md#GetRepositoryPipelineCaches) | **Get** /repositories/{workspace}/{repo_slug}/pipelines-config/caches/ | 
[**GetRepositoryPipelineConfig**](PipelinesApi.md#GetRepositoryPipelineConfig) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config | 
[**GetRepositoryPipelineKnownHost**](PipelinesApi.md#GetRepositoryPipelineKnownHost) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/known_hosts/{known_host_uuid} | 
[**GetRepositoryPipelineKnownHosts**](PipelinesApi.md#GetRepositoryPipelineKnownHosts) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/known_hosts/ | 
[**GetRepositoryPipelineSchedule**](PipelinesApi.md#GetRepositoryPipelineSchedule) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/{schedule_uuid} | 
[**GetRepositoryPipelineScheduleExecutions**](PipelinesApi.md#GetRepositoryPipelineScheduleExecutions) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/{schedule_uuid}/executions/ | 
[**GetRepositoryPipelineSchedules**](PipelinesApi.md#GetRepositoryPipelineSchedules) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/ | 
[**GetRepositoryPipelineSshKeyPair**](PipelinesApi.md#GetRepositoryPipelineSshKeyPair) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/key_pair | 
[**GetRepositoryPipelineVariable**](PipelinesApi.md#GetRepositoryPipelineVariable) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/variables/{variable_uuid} | 
[**GetRepositoryPipelineVariables**](PipelinesApi.md#GetRepositoryPipelineVariables) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/variables/ | 
[**StopPipeline**](PipelinesApi.md#StopPipeline) | **Post** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/stopPipeline | 
[**UpdateDeploymentVariable**](PipelinesApi.md#UpdateDeploymentVariable) | **Put** /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables/{variable_uuid} | 
[**UpdatePipelineVariableForTeam**](PipelinesApi.md#UpdatePipelineVariableForTeam) | **Put** /teams/{username}/pipelines_config/variables/{variable_uuid} | 
[**UpdatePipelineVariableForUser**](PipelinesApi.md#UpdatePipelineVariableForUser) | **Put** /users/{selected_user}/pipelines_config/variables/{variable_uuid} | 
[**UpdatePipelineVariableForWorkspace**](PipelinesApi.md#UpdatePipelineVariableForWorkspace) | **Put** /workspaces/{workspace}/pipelines-config/variables/{variable_uuid} | 
[**UpdateRepositoryBuildNumber**](PipelinesApi.md#UpdateRepositoryBuildNumber) | **Put** /repositories/{workspace}/{repo_slug}/pipelines_config/build_number | 
[**UpdateRepositoryPipelineConfig**](PipelinesApi.md#UpdateRepositoryPipelineConfig) | **Put** /repositories/{workspace}/{repo_slug}/pipelines_config | 
[**UpdateRepositoryPipelineKeyPair**](PipelinesApi.md#UpdateRepositoryPipelineKeyPair) | **Put** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/key_pair | 
[**UpdateRepositoryPipelineKnownHost**](PipelinesApi.md#UpdateRepositoryPipelineKnownHost) | **Put** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/known_hosts/{known_host_uuid} | 
[**UpdateRepositoryPipelineSchedule**](PipelinesApi.md#UpdateRepositoryPipelineSchedule) | **Put** /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/{schedule_uuid} | 
[**UpdateRepositoryPipelineVariable**](PipelinesApi.md#UpdateRepositoryPipelineVariable) | **Put** /repositories/{workspace}/{repo_slug}/pipelines_config/variables/{variable_uuid} | 



## CreateDeploymentVariable

> DeploymentVariable CreateDeploymentVariable(ctx, workspace, repoSlug, environmentUuid, variableUuid).Body(body).Execute()





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
    environmentUuid := "environmentUuid_example" // string | The environment.
    variableUuid := "variableUuid_example" // string | The UUID of the variable to update.
    body := *openapiclient.NewDeploymentVariable("Type_example") // DeploymentVariable | The variable to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.CreateDeploymentVariable(context.Background(), workspace, repoSlug, environmentUuid, variableUuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.CreateDeploymentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeploymentVariable`: DeploymentVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.CreateDeploymentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**environmentUuid** | **string** | The environment. | 
**variableUuid** | **string** | The UUID of the variable to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**DeploymentVariable**](DeploymentVariable.md) | The variable to create | 

### Return type

[**DeploymentVariable**](DeploymentVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePipelineForRepository

> Pipeline CreatePipelineForRepository(ctx, workspace, repoSlug).Body(body).Execute()





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
    body := *openapiclient.NewPipeline("Type_example") // Pipeline | The pipeline to initiate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.CreatePipelineForRepository(context.Background(), workspace, repoSlug).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.CreatePipelineForRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePipelineForRepository`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.CreatePipelineForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePipelineForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Pipeline**](Pipeline.md) | The pipeline to initiate. | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePipelineVariableForTeam

> PipelineVariable CreatePipelineVariableForTeam(ctx, workspace).Body(body).Execute()





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
    body := *openapiclient.NewPipelineVariable("Type_example") // PipelineVariable | The variable to create. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.CreatePipelineVariableForTeam(context.Background(), workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.CreatePipelineVariableForTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePipelineVariableForTeam`: PipelineVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.CreatePipelineVariableForTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePipelineVariableForTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PipelineVariable**](PipelineVariable.md) | The variable to create. | 

### Return type

[**PipelineVariable**](PipelineVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePipelineVariableForUser

> PipelineVariable CreatePipelineVariableForUser(ctx, selectedUser).Body(body).Execute()





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
    body := *openapiclient.NewPipelineVariable("Type_example") // PipelineVariable | The variable to create. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.CreatePipelineVariableForUser(context.Background(), selectedUser).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.CreatePipelineVariableForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePipelineVariableForUser`: PipelineVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.CreatePipelineVariableForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePipelineVariableForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PipelineVariable**](PipelineVariable.md) | The variable to create. | 

### Return type

[**PipelineVariable**](PipelineVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePipelineVariableForWorkspace

> PipelineVariable CreatePipelineVariableForWorkspace(ctx, workspace).Body(body).Execute()





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
    body := *openapiclient.NewPipelineVariable("Type_example") // PipelineVariable | The variable to create. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.CreatePipelineVariableForWorkspace(context.Background(), workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.CreatePipelineVariableForWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePipelineVariableForWorkspace`: PipelineVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.CreatePipelineVariableForWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePipelineVariableForWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PipelineVariable**](PipelineVariable.md) | The variable to create. | 

### Return type

[**PipelineVariable**](PipelineVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRepositoryPipelineKnownHost

> PipelineKnownHost CreateRepositoryPipelineKnownHost(ctx, workspace, repoSlug).Body(body).Execute()





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
    body := *openapiclient.NewPipelineKnownHost("Type_example") // PipelineKnownHost | The known host to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.CreateRepositoryPipelineKnownHost(context.Background(), workspace, repoSlug).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.CreateRepositoryPipelineKnownHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRepositoryPipelineKnownHost`: PipelineKnownHost
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.CreateRepositoryPipelineKnownHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryPipelineKnownHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PipelineKnownHost**](PipelineKnownHost.md) | The known host to create. | 

### Return type

[**PipelineKnownHost**](PipelineKnownHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRepositoryPipelineSchedule

> PipelineSchedule CreateRepositoryPipelineSchedule(ctx, workspace, repoSlug).Body(body).Execute()





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
    body := *openapiclient.NewPipelineSchedule("Type_example") // PipelineSchedule | The schedule to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.CreateRepositoryPipelineSchedule(context.Background(), workspace, repoSlug).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.CreateRepositoryPipelineSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRepositoryPipelineSchedule`: PipelineSchedule
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.CreateRepositoryPipelineSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryPipelineScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PipelineSchedule**](PipelineSchedule.md) | The schedule to create. | 

### Return type

[**PipelineSchedule**](PipelineSchedule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRepositoryPipelineVariable

> PipelineVariable CreateRepositoryPipelineVariable(ctx, workspace, repoSlug).Body(body).Execute()





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
    body := *openapiclient.NewPipelineVariable("Type_example") // PipelineVariable | The variable to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.CreateRepositoryPipelineVariable(context.Background(), workspace, repoSlug).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.CreateRepositoryPipelineVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRepositoryPipelineVariable`: PipelineVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.CreateRepositoryPipelineVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryPipelineVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PipelineVariable**](PipelineVariable.md) | The variable to create. | 

### Return type

[**PipelineVariable**](PipelineVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeploymentVariable

> DeleteDeploymentVariable(ctx, workspace, repoSlug, environmentUuid, variableUuid).Execute()





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
    environmentUuid := "environmentUuid_example" // string | The environment.
    variableUuid := "variableUuid_example" // string | The UUID of the variable to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.DeleteDeploymentVariable(context.Background(), workspace, repoSlug, environmentUuid, variableUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.DeleteDeploymentVariable``: %v\n", err)
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
**environmentUuid** | **string** | The environment. | 
**variableUuid** | **string** | The UUID of the variable to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentVariableRequest struct via the builder pattern


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


## DeletePipelineVariableForTeam

> DeletePipelineVariableForTeam(ctx, username, variableUuid).Execute()





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
    username := "username_example" // string | The account.
    variableUuid := "variableUuid_example" // string | The UUID of the variable to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.DeletePipelineVariableForTeam(context.Background(), username, variableUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.DeletePipelineVariableForTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The account. | 
**variableUuid** | **string** | The UUID of the variable to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineVariableForTeamRequest struct via the builder pattern


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


## DeletePipelineVariableForUser

> DeletePipelineVariableForUser(ctx, selectedUser, variableUuid).Execute()





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
    variableUuid := "variableUuid_example" // string | The UUID of the variable to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.DeletePipelineVariableForUser(context.Background(), selectedUser, variableUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.DeletePipelineVariableForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
**variableUuid** | **string** | The UUID of the variable to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineVariableForUserRequest struct via the builder pattern


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


## DeletePipelineVariableForWorkspace

> DeletePipelineVariableForWorkspace(ctx, workspace, variableUuid).Execute()





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
    variableUuid := "variableUuid_example" // string | The UUID of the variable to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.DeletePipelineVariableForWorkspace(context.Background(), workspace, variableUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.DeletePipelineVariableForWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**variableUuid** | **string** | The UUID of the variable to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineVariableForWorkspaceRequest struct via the builder pattern


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


## DeleteRepositoryPipelineCache

> DeleteRepositoryPipelineCache(ctx, workspace, repoSlug, cacheUuid).Execute()





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
    workspace := "workspace_example" // string | The account.
    repoSlug := "repoSlug_example" // string | The repository.
    cacheUuid := "cacheUuid_example" // string | The UUID of the cache to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.DeleteRepositoryPipelineCache(context.Background(), workspace, repoSlug, cacheUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.DeleteRepositoryPipelineCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The account. | 
**repoSlug** | **string** | The repository. | 
**cacheUuid** | **string** | The UUID of the cache to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryPipelineCacheRequest struct via the builder pattern


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


## DeleteRepositoryPipelineKeyPair

> DeleteRepositoryPipelineKeyPair(ctx, workspace, repoSlug).Execute()





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
    resp, r, err := api_client.PipelinesApi.DeleteRepositoryPipelineKeyPair(context.Background(), workspace, repoSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.DeleteRepositoryPipelineKeyPair``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryPipelineKeyPairRequest struct via the builder pattern


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


## DeleteRepositoryPipelineKnownHost

> DeleteRepositoryPipelineKnownHost(ctx, workspace, repoSlug, knownHostUuid).Execute()





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
    knownHostUuid := "knownHostUuid_example" // string | The UUID of the known host to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.DeleteRepositoryPipelineKnownHost(context.Background(), workspace, repoSlug, knownHostUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.DeleteRepositoryPipelineKnownHost``: %v\n", err)
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
**knownHostUuid** | **string** | The UUID of the known host to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryPipelineKnownHostRequest struct via the builder pattern


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


## DeleteRepositoryPipelineSchedule

> DeleteRepositoryPipelineSchedule(ctx, workspace, repoSlug, scheduleUuid).Execute()





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
    scheduleUuid := "scheduleUuid_example" // string | The uuid of the schedule.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.DeleteRepositoryPipelineSchedule(context.Background(), workspace, repoSlug, scheduleUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.DeleteRepositoryPipelineSchedule``: %v\n", err)
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
**scheduleUuid** | **string** | The uuid of the schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryPipelineScheduleRequest struct via the builder pattern


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


## DeleteRepositoryPipelineVariable

> DeleteRepositoryPipelineVariable(ctx, workspace, repoSlug, variableUuid).Execute()





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
    variableUuid := "variableUuid_example" // string | The UUID of the variable to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.DeleteRepositoryPipelineVariable(context.Background(), workspace, repoSlug, variableUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.DeleteRepositoryPipelineVariable``: %v\n", err)
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
**variableUuid** | **string** | The UUID of the variable to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryPipelineVariableRequest struct via the builder pattern


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


## GetDeploymentVariables

> PaginatedDeploymentVariable GetDeploymentVariables(ctx, workspace, repoSlug, environmentUuid).Execute()





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
    environmentUuid := "environmentUuid_example" // string | The environment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetDeploymentVariables(context.Background(), workspace, repoSlug, environmentUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetDeploymentVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentVariables`: PaginatedDeploymentVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetDeploymentVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**environmentUuid** | **string** | The environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PaginatedDeploymentVariable**](PaginatedDeploymentVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineForRepository

> Pipeline GetPipelineForRepository(ctx, workspace, repoSlug, pipelineUuid).Execute()





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
    pipelineUuid := "pipelineUuid_example" // string | The pipeline UUID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetPipelineForRepository(context.Background(), workspace, repoSlug, pipelineUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetPipelineForRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineForRepository`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetPipelineForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**pipelineUuid** | **string** | The pipeline UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineStepForRepository

> PipelineStep GetPipelineStepForRepository(ctx, workspace, repoSlug, pipelineUuid, stepUuid).Execute()





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
    pipelineUuid := "pipelineUuid_example" // string | The UUID of the pipeline.
    stepUuid := "stepUuid_example" // string | The UUID of the step.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetPipelineStepForRepository(context.Background(), workspace, repoSlug, pipelineUuid, stepUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetPipelineStepForRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineStepForRepository`: PipelineStep
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetPipelineStepForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**pipelineUuid** | **string** | The UUID of the pipeline. | 
**stepUuid** | **string** | The UUID of the step. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineStepForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PipelineStep**](PipelineStep.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineStepLogForRepository

> GetPipelineStepLogForRepository(ctx, workspace, repoSlug, pipelineUuid, stepUuid).Execute()





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
    pipelineUuid := "pipelineUuid_example" // string | The UUID of the pipeline.
    stepUuid := "stepUuid_example" // string | The UUID of the step.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetPipelineStepLogForRepository(context.Background(), workspace, repoSlug, pipelineUuid, stepUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetPipelineStepLogForRepository``: %v\n", err)
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
**pipelineUuid** | **string** | The UUID of the pipeline. | 
**stepUuid** | **string** | The UUID of the step. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineStepLogForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineStepsForRepository

> PaginatedPipelineSteps GetPipelineStepsForRepository(ctx, workspace, repoSlug, pipelineUuid).Execute()





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
    pipelineUuid := "pipelineUuid_example" // string | The UUID of the pipeline.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetPipelineStepsForRepository(context.Background(), workspace, repoSlug, pipelineUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetPipelineStepsForRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineStepsForRepository`: PaginatedPipelineSteps
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetPipelineStepsForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**pipelineUuid** | **string** | The UUID of the pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineStepsForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PaginatedPipelineSteps**](PaginatedPipelineSteps.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineVariableForTeam

> PipelineVariable GetPipelineVariableForTeam(ctx, username, variableUuid).Execute()





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
    username := "username_example" // string | The account.
    variableUuid := "variableUuid_example" // string | The UUID of the variable to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetPipelineVariableForTeam(context.Background(), username, variableUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetPipelineVariableForTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineVariableForTeam`: PipelineVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetPipelineVariableForTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The account. | 
**variableUuid** | **string** | The UUID of the variable to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineVariableForTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PipelineVariable**](PipelineVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineVariableForUser

> PipelineVariable GetPipelineVariableForUser(ctx, selectedUser, variableUuid).Execute()





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
    variableUuid := "variableUuid_example" // string | The UUID of the variable to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetPipelineVariableForUser(context.Background(), selectedUser, variableUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetPipelineVariableForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineVariableForUser`: PipelineVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetPipelineVariableForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
**variableUuid** | **string** | The UUID of the variable to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineVariableForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PipelineVariable**](PipelineVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineVariableForWorkspace

> PipelineVariable GetPipelineVariableForWorkspace(ctx, workspace, variableUuid).Execute()





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
    variableUuid := "variableUuid_example" // string | The UUID of the variable to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetPipelineVariableForWorkspace(context.Background(), workspace, variableUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetPipelineVariableForWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineVariableForWorkspace`: PipelineVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetPipelineVariableForWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**variableUuid** | **string** | The UUID of the variable to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineVariableForWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PipelineVariable**](PipelineVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineVariablesForTeam

> PaginatedPipelineVariables GetPipelineVariablesForTeam(ctx, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetPipelineVariablesForTeam(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetPipelineVariablesForTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineVariablesForTeam`: PaginatedPipelineVariables
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetPipelineVariablesForTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineVariablesForTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedPipelineVariables**](PaginatedPipelineVariables.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineVariablesForUser

> PaginatedPipelineVariables GetPipelineVariablesForUser(ctx, selectedUser).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetPipelineVariablesForUser(context.Background(), selectedUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetPipelineVariablesForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineVariablesForUser`: PaginatedPipelineVariables
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetPipelineVariablesForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineVariablesForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedPipelineVariables**](PaginatedPipelineVariables.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineVariablesForWorkspace

> PaginatedPipelineVariables GetPipelineVariablesForWorkspace(ctx, workspace).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetPipelineVariablesForWorkspace(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetPipelineVariablesForWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineVariablesForWorkspace`: PaginatedPipelineVariables
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetPipelineVariablesForWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineVariablesForWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedPipelineVariables**](PaginatedPipelineVariables.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelinesForRepository

> PaginatedPipelines GetPipelinesForRepository(ctx, workspace, repoSlug).Execute()





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
    resp, r, err := api_client.PipelinesApi.GetPipelinesForRepository(context.Background(), workspace, repoSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetPipelinesForRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelinesForRepository`: PaginatedPipelines
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetPipelinesForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelinesForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedPipelines**](PaginatedPipelines.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryPipelineCacheContentURI

> PipelineCacheContentUri GetRepositoryPipelineCacheContentURI(ctx, workspace, repoSlug, cacheUuid).Execute()





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
    workspace := "workspace_example" // string | The account.
    repoSlug := "repoSlug_example" // string | The repository.
    cacheUuid := "cacheUuid_example" // string | The UUID of the cache.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetRepositoryPipelineCacheContentURI(context.Background(), workspace, repoSlug, cacheUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetRepositoryPipelineCacheContentURI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryPipelineCacheContentURI`: PipelineCacheContentUri
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetRepositoryPipelineCacheContentURI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The account. | 
**repoSlug** | **string** | The repository. | 
**cacheUuid** | **string** | The UUID of the cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryPipelineCacheContentURIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PipelineCacheContentUri**](PipelineCacheContentUri.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryPipelineCaches

> PaginatedPipelineCaches GetRepositoryPipelineCaches(ctx, workspace, repoSlug).Execute()





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
    workspace := "workspace_example" // string | The account.
    repoSlug := "repoSlug_example" // string | The repository.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetRepositoryPipelineCaches(context.Background(), workspace, repoSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetRepositoryPipelineCaches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryPipelineCaches`: PaginatedPipelineCaches
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetRepositoryPipelineCaches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The account. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryPipelineCachesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedPipelineCaches**](PaginatedPipelineCaches.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryPipelineConfig

> PipelinesConfig GetRepositoryPipelineConfig(ctx, workspace, repoSlug).Execute()





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
    workspace := "workspace_example" // string | The account.
    repoSlug := "repoSlug_example" // string | The repository.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetRepositoryPipelineConfig(context.Background(), workspace, repoSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetRepositoryPipelineConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryPipelineConfig`: PipelinesConfig
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetRepositoryPipelineConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The account. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryPipelineConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PipelinesConfig**](PipelinesConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryPipelineKnownHost

> PipelineKnownHost GetRepositoryPipelineKnownHost(ctx, workspace, repoSlug, knownHostUuid).Execute()





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
    knownHostUuid := "knownHostUuid_example" // string | The UUID of the known host to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetRepositoryPipelineKnownHost(context.Background(), workspace, repoSlug, knownHostUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetRepositoryPipelineKnownHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryPipelineKnownHost`: PipelineKnownHost
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetRepositoryPipelineKnownHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**knownHostUuid** | **string** | The UUID of the known host to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryPipelineKnownHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PipelineKnownHost**](PipelineKnownHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryPipelineKnownHosts

> PaginatedPipelineKnownHosts GetRepositoryPipelineKnownHosts(ctx, workspace, repoSlug).Execute()





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
    resp, r, err := api_client.PipelinesApi.GetRepositoryPipelineKnownHosts(context.Background(), workspace, repoSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetRepositoryPipelineKnownHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryPipelineKnownHosts`: PaginatedPipelineKnownHosts
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetRepositoryPipelineKnownHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryPipelineKnownHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedPipelineKnownHosts**](PaginatedPipelineKnownHosts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryPipelineSchedule

> PipelineSchedule GetRepositoryPipelineSchedule(ctx, workspace, repoSlug, scheduleUuid).Execute()





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
    scheduleUuid := "scheduleUuid_example" // string | The uuid of the schedule.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetRepositoryPipelineSchedule(context.Background(), workspace, repoSlug, scheduleUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetRepositoryPipelineSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryPipelineSchedule`: PipelineSchedule
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetRepositoryPipelineSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**scheduleUuid** | **string** | The uuid of the schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryPipelineScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PipelineSchedule**](PipelineSchedule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryPipelineScheduleExecutions

> PaginatedPipelineScheduleExecutions GetRepositoryPipelineScheduleExecutions(ctx, workspace, repoSlug).Execute()





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
    resp, r, err := api_client.PipelinesApi.GetRepositoryPipelineScheduleExecutions(context.Background(), workspace, repoSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetRepositoryPipelineScheduleExecutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryPipelineScheduleExecutions`: PaginatedPipelineScheduleExecutions
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetRepositoryPipelineScheduleExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryPipelineScheduleExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedPipelineScheduleExecutions**](PaginatedPipelineScheduleExecutions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryPipelineSchedules

> PaginatedPipelineSchedules GetRepositoryPipelineSchedules(ctx, workspace, repoSlug).Execute()





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
    resp, r, err := api_client.PipelinesApi.GetRepositoryPipelineSchedules(context.Background(), workspace, repoSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetRepositoryPipelineSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryPipelineSchedules`: PaginatedPipelineSchedules
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetRepositoryPipelineSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryPipelineSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedPipelineSchedules**](PaginatedPipelineSchedules.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryPipelineSshKeyPair

> PipelineSshKeyPair GetRepositoryPipelineSshKeyPair(ctx, workspace, repoSlug).Execute()





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
    resp, r, err := api_client.PipelinesApi.GetRepositoryPipelineSshKeyPair(context.Background(), workspace, repoSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetRepositoryPipelineSshKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryPipelineSshKeyPair`: PipelineSshKeyPair
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetRepositoryPipelineSshKeyPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryPipelineSshKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PipelineSshKeyPair**](PipelineSshKeyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryPipelineVariable

> PipelineVariable GetRepositoryPipelineVariable(ctx, workspace, repoSlug, variableUuid).Execute()





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
    variableUuid := "variableUuid_example" // string | The UUID of the variable to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.GetRepositoryPipelineVariable(context.Background(), workspace, repoSlug, variableUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetRepositoryPipelineVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryPipelineVariable`: PipelineVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetRepositoryPipelineVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**variableUuid** | **string** | The UUID of the variable to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryPipelineVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PipelineVariable**](PipelineVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryPipelineVariables

> PaginatedPipelineVariables GetRepositoryPipelineVariables(ctx, workspace, repoSlug).Execute()





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
    resp, r, err := api_client.PipelinesApi.GetRepositoryPipelineVariables(context.Background(), workspace, repoSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.GetRepositoryPipelineVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryPipelineVariables`: PaginatedPipelineVariables
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.GetRepositoryPipelineVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryPipelineVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedPipelineVariables**](PaginatedPipelineVariables.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopPipeline

> StopPipeline(ctx, workspace, repoSlug, pipelineUuid).Execute()





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
    pipelineUuid := "pipelineUuid_example" // string | The UUID of the pipeline.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.StopPipeline(context.Background(), workspace, repoSlug, pipelineUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.StopPipeline``: %v\n", err)
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
**pipelineUuid** | **string** | The UUID of the pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopPipelineRequest struct via the builder pattern


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


## UpdateDeploymentVariable

> DeploymentVariable UpdateDeploymentVariable(ctx, workspace, repoSlug, environmentUuid, variableUuid).Body(body).Execute()





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
    environmentUuid := "environmentUuid_example" // string | The environment.
    variableUuid := "variableUuid_example" // string | The UUID of the variable to update.
    body := *openapiclient.NewDeploymentVariable("Type_example") // DeploymentVariable | The updated deployment variable.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.UpdateDeploymentVariable(context.Background(), workspace, repoSlug, environmentUuid, variableUuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.UpdateDeploymentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeploymentVariable`: DeploymentVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.UpdateDeploymentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**environmentUuid** | **string** | The environment. | 
**variableUuid** | **string** | The UUID of the variable to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**DeploymentVariable**](DeploymentVariable.md) | The updated deployment variable. | 

### Return type

[**DeploymentVariable**](DeploymentVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePipelineVariableForTeam

> PipelineVariable UpdatePipelineVariableForTeam(ctx, username, variableUuid).Body(body).Execute()





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
    username := "username_example" // string | The account.
    variableUuid := "variableUuid_example" // string | The UUID of the variable.
    body := *openapiclient.NewPipelineVariable("Type_example") // PipelineVariable | The updated variable.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.UpdatePipelineVariableForTeam(context.Background(), username, variableUuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.UpdatePipelineVariableForTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePipelineVariableForTeam`: PipelineVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.UpdatePipelineVariableForTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The account. | 
**variableUuid** | **string** | The UUID of the variable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePipelineVariableForTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PipelineVariable**](PipelineVariable.md) | The updated variable. | 

### Return type

[**PipelineVariable**](PipelineVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePipelineVariableForUser

> PipelineVariable UpdatePipelineVariableForUser(ctx, selectedUser, variableUuid).Body(body).Execute()





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
    variableUuid := "variableUuid_example" // string | The UUID of the variable.
    body := *openapiclient.NewPipelineVariable("Type_example") // PipelineVariable | The updated variable.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.UpdatePipelineVariableForUser(context.Background(), selectedUser, variableUuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.UpdatePipelineVariableForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePipelineVariableForUser`: PipelineVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.UpdatePipelineVariableForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selectedUser** | **string** | Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
**variableUuid** | **string** | The UUID of the variable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePipelineVariableForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PipelineVariable**](PipelineVariable.md) | The updated variable. | 

### Return type

[**PipelineVariable**](PipelineVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePipelineVariableForWorkspace

> PipelineVariable UpdatePipelineVariableForWorkspace(ctx, workspace, variableUuid).Body(body).Execute()





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
    variableUuid := "variableUuid_example" // string | The UUID of the variable.
    body := *openapiclient.NewPipelineVariable("Type_example") // PipelineVariable | The updated variable.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.UpdatePipelineVariableForWorkspace(context.Background(), workspace, variableUuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.UpdatePipelineVariableForWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePipelineVariableForWorkspace`: PipelineVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.UpdatePipelineVariableForWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**variableUuid** | **string** | The UUID of the variable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePipelineVariableForWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PipelineVariable**](PipelineVariable.md) | The updated variable. | 

### Return type

[**PipelineVariable**](PipelineVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepositoryBuildNumber

> PipelineBuildNumber UpdateRepositoryBuildNumber(ctx, workspace, repoSlug).Body(body).Execute()





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
    body := *openapiclient.NewPipelineBuildNumber("Type_example") // PipelineBuildNumber | The build number to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.UpdateRepositoryBuildNumber(context.Background(), workspace, repoSlug).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.UpdateRepositoryBuildNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRepositoryBuildNumber`: PipelineBuildNumber
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.UpdateRepositoryBuildNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryBuildNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PipelineBuildNumber**](PipelineBuildNumber.md) | The build number to update. | 

### Return type

[**PipelineBuildNumber**](PipelineBuildNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepositoryPipelineConfig

> PipelinesConfig UpdateRepositoryPipelineConfig(ctx, workspace, repoSlug).Body(body).Execute()





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
    body := *openapiclient.NewPipelinesConfig("Type_example") // PipelinesConfig | The updated repository pipelines configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.UpdateRepositoryPipelineConfig(context.Background(), workspace, repoSlug).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.UpdateRepositoryPipelineConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRepositoryPipelineConfig`: PipelinesConfig
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.UpdateRepositoryPipelineConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryPipelineConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PipelinesConfig**](PipelinesConfig.md) | The updated repository pipelines configuration. | 

### Return type

[**PipelinesConfig**](PipelinesConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepositoryPipelineKeyPair

> PipelineSshKeyPair UpdateRepositoryPipelineKeyPair(ctx, workspace, repoSlug).Body(body).Execute()





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
    body := *openapiclient.NewPipelineSshKeyPair("Type_example") // PipelineSshKeyPair | The created or updated SSH key pair.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.UpdateRepositoryPipelineKeyPair(context.Background(), workspace, repoSlug).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.UpdateRepositoryPipelineKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRepositoryPipelineKeyPair`: PipelineSshKeyPair
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.UpdateRepositoryPipelineKeyPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryPipelineKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PipelineSshKeyPair**](PipelineSshKeyPair.md) | The created or updated SSH key pair. | 

### Return type

[**PipelineSshKeyPair**](PipelineSshKeyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepositoryPipelineKnownHost

> PipelineKnownHost UpdateRepositoryPipelineKnownHost(ctx, workspace, repoSlug, knownHostUuid).Body(body).Execute()





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
    knownHostUuid := "knownHostUuid_example" // string | The UUID of the known host to update.
    body := *openapiclient.NewPipelineKnownHost("Type_example") // PipelineKnownHost | The updated known host.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.UpdateRepositoryPipelineKnownHost(context.Background(), workspace, repoSlug, knownHostUuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.UpdateRepositoryPipelineKnownHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRepositoryPipelineKnownHost`: PipelineKnownHost
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.UpdateRepositoryPipelineKnownHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**knownHostUuid** | **string** | The UUID of the known host to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryPipelineKnownHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**PipelineKnownHost**](PipelineKnownHost.md) | The updated known host. | 

### Return type

[**PipelineKnownHost**](PipelineKnownHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepositoryPipelineSchedule

> PipelineSchedule UpdateRepositoryPipelineSchedule(ctx, workspace, repoSlug, scheduleUuid).Body(body).Execute()





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
    scheduleUuid := "scheduleUuid_example" // string | The uuid of the schedule.
    body := *openapiclient.NewPipelineSchedule("Type_example") // PipelineSchedule | The schedule to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.UpdateRepositoryPipelineSchedule(context.Background(), workspace, repoSlug, scheduleUuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.UpdateRepositoryPipelineSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRepositoryPipelineSchedule`: PipelineSchedule
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.UpdateRepositoryPipelineSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**scheduleUuid** | **string** | The uuid of the schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryPipelineScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**PipelineSchedule**](PipelineSchedule.md) | The schedule to update. | 

### Return type

[**PipelineSchedule**](PipelineSchedule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepositoryPipelineVariable

> PipelineVariable UpdateRepositoryPipelineVariable(ctx, workspace, repoSlug, variableUuid).Body(body).Execute()





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
    variableUuid := "variableUuid_example" // string | The UUID of the variable to update.
    body := *openapiclient.NewPipelineVariable("Type_example") // PipelineVariable | The updated variable

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PipelinesApi.UpdateRepositoryPipelineVariable(context.Background(), workspace, repoSlug, variableUuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.UpdateRepositoryPipelineVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRepositoryPipelineVariable`: PipelineVariable
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.UpdateRepositoryPipelineVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
**repoSlug** | **string** | The repository. | 
**variableUuid** | **string** | The UUID of the variable to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryPipelineVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**PipelineVariable**](PipelineVariable.md) | The updated variable | 

### Return type

[**PipelineVariable**](PipelineVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

