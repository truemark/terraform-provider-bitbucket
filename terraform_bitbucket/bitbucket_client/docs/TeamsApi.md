# \TeamsApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsGet**](TeamsApi.md#TeamsGet) | **Get** /teams | 
[**TeamsUsernameFollowersGet**](TeamsApi.md#TeamsUsernameFollowersGet) | **Get** /teams/{username}/followers | 
[**TeamsUsernameFollowingGet**](TeamsApi.md#TeamsUsernameFollowingGet) | **Get** /teams/{username}/following | 
[**TeamsUsernameGet**](TeamsApi.md#TeamsUsernameGet) | **Get** /teams/{username} | 
[**TeamsUsernameHooksGet**](TeamsApi.md#TeamsUsernameHooksGet) | **Get** /teams/{username}/hooks | 
[**TeamsUsernameHooksPost**](TeamsApi.md#TeamsUsernameHooksPost) | **Post** /teams/{username}/hooks | 
[**TeamsUsernameHooksUidDelete**](TeamsApi.md#TeamsUsernameHooksUidDelete) | **Delete** /teams/{username}/hooks/{uid} | 
[**TeamsUsernameHooksUidGet**](TeamsApi.md#TeamsUsernameHooksUidGet) | **Get** /teams/{username}/hooks/{uid} | 
[**TeamsUsernameHooksUidPut**](TeamsApi.md#TeamsUsernameHooksUidPut) | **Put** /teams/{username}/hooks/{uid} | 
[**TeamsUsernameMembersGet**](TeamsApi.md#TeamsUsernameMembersGet) | **Get** /teams/{username}/members | 
[**TeamsUsernamePermissionsGet**](TeamsApi.md#TeamsUsernamePermissionsGet) | **Get** /teams/{username}/permissions | 
[**TeamsUsernamePermissionsRepositoriesGet**](TeamsApi.md#TeamsUsernamePermissionsRepositoriesGet) | **Get** /teams/{username}/permissions/repositories | 
[**TeamsUsernamePermissionsRepositoriesRepoSlugGet**](TeamsApi.md#TeamsUsernamePermissionsRepositoriesRepoSlugGet) | **Get** /teams/{username}/permissions/repositories/{repo_slug} | 
[**TeamsWorkspaceRepositoriesGet**](TeamsApi.md#TeamsWorkspaceRepositoriesGet) | **Get** /teams/{workspace}/repositories | 
[**UserPermissionsTeamsGet**](TeamsApi.md#UserPermissionsTeamsGet) | **Get** /user/permissions/teams | 
[**UsersWorkspaceRepositoriesGet**](TeamsApi.md#UsersWorkspaceRepositoriesGet) | **Get** /users/{workspace}/repositories | 



## TeamsGet

> PaginatedTeams TeamsGet(ctx).Role(role).Execute()





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
    role := "role_example" // string |  Filters the teams based on the authenticated user's role on each team.  * **member**: returns a list of all the teams which the caller is a member of   at least one team group or repository owned by the team * **contributor**: returns a list of teams which the caller has write access   to at least one repository owned by the team * **admin**: returns a list teams which the caller has team administrator access  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.TeamsGet(context.Background()).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGet`: PaginatedTeams
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **string** |  Filters the teams based on the authenticated user&#39;s role on each team.  * **member**: returns a list of all the teams which the caller is a member of   at least one team group or repository owned by the team * **contributor**: returns a list of teams which the caller has write access   to at least one repository owned by the team * **admin**: returns a list teams which the caller has team administrator access  | 

### Return type

[**PaginatedTeams**](PaginatedTeams.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUsernameFollowersGet

> PaginatedUsers TeamsUsernameFollowersGet(ctx, username).Execute()





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
    resp, r, err := api_client.TeamsApi.TeamsUsernameFollowersGet(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUsernameFollowersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameFollowersGet`: PaginatedUsers
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUsernameFollowersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameFollowersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedUsers**](PaginatedUsers.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUsernameFollowingGet

> PaginatedUsers TeamsUsernameFollowingGet(ctx, username).Execute()





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
    resp, r, err := api_client.TeamsApi.TeamsUsernameFollowingGet(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUsernameFollowingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameFollowingGet`: PaginatedUsers
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUsernameFollowingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameFollowingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedUsers**](PaginatedUsers.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUsernameGet

> Team TeamsUsernameGet(ctx, username).Execute()





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
    resp, r, err := api_client.TeamsApi.TeamsUsernameGet(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUsernameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameGet`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUsernameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Team**](Team.md)

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
    resp, r, err := api_client.TeamsApi.TeamsUsernameHooksGet(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUsernameHooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameHooksGet`: PaginatedWebhookSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUsernameHooksGet`: %v\n", resp)
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
    resp, r, err := api_client.TeamsApi.TeamsUsernameHooksPost(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUsernameHooksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameHooksPost`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUsernameHooksPost`: %v\n", resp)
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
    resp, r, err := api_client.TeamsApi.TeamsUsernameHooksUidDelete(context.Background(), uid, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUsernameHooksUidDelete``: %v\n", err)
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
    resp, r, err := api_client.TeamsApi.TeamsUsernameHooksUidGet(context.Background(), uid, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUsernameHooksUidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameHooksUidGet`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUsernameHooksUidGet`: %v\n", resp)
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
    resp, r, err := api_client.TeamsApi.TeamsUsernameHooksUidPut(context.Background(), uid, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUsernameHooksUidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameHooksUidPut`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUsernameHooksUidPut`: %v\n", resp)
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


## TeamsUsernameMembersGet

> User TeamsUsernameMembersGet(ctx, username).Execute()





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
    resp, r, err := api_client.TeamsApi.TeamsUsernameMembersGet(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUsernameMembersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernameMembersGet`: User
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUsernameMembersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernameMembersGetRequest struct via the builder pattern


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


## TeamsUsernamePermissionsGet

> PaginatedTeamPermissions TeamsUsernamePermissionsGet(ctx, username).Q(q).Sort(sort).Execute()





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
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../../meta/filtering). (optional)
    sort := "sort_example" // string |  Name of a response property sort the result by as per [filtering and sorting](../../../meta/filtering#query-sort).  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.TeamsUsernamePermissionsGet(context.Background(), username).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUsernamePermissionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernamePermissionsGet`: PaginatedTeamPermissions
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUsernamePermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernamePermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** |  Query string to narrow down the response as per [filtering and sorting](../../../meta/filtering). | 
 **sort** | **string** |  Name of a response property sort the result by as per [filtering and sorting](../../../meta/filtering#query-sort).  | 

### Return type

[**PaginatedTeamPermissions**](PaginatedTeamPermissions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUsernamePermissionsRepositoriesGet

> PaginatedRepositoryPermissions TeamsUsernamePermissionsRepositoriesGet(ctx, username).Q(q).Sort(sort).Execute()





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
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../../../meta/filtering). (optional)
    sort := "sort_example" // string |  Name of a response property sort the result by as per [filtering and sorting](../../../../meta/filtering#query-sort).  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.TeamsUsernamePermissionsRepositoriesGet(context.Background(), username).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUsernamePermissionsRepositoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernamePermissionsRepositoriesGet`: PaginatedRepositoryPermissions
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUsernamePermissionsRepositoriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernamePermissionsRepositoriesGetRequest struct via the builder pattern


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


## TeamsUsernamePermissionsRepositoriesRepoSlugGet

> PaginatedRepositoryPermissions TeamsUsernamePermissionsRepositoriesRepoSlugGet(ctx, repoSlug, username).Q(q).Sort(sort).Execute()





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
    username := "username_example" // string | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`. An account is either a team or user. 
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../../../meta/filtering). (optional)
    sort := "sort_example" // string |  Name of a response property sort the result by as per [filtering and sorting](../../../../meta/filtering#query-sort).  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.TeamsUsernamePermissionsRepositoriesRepoSlugGet(context.Background(), repoSlug, username).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUsernamePermissionsRepositoriesRepoSlugGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUsernamePermissionsRepositoriesRepoSlugGet`: PaginatedRepositoryPermissions
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUsernamePermissionsRepositoriesRepoSlugGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoSlug** | **string** | This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
**username** | **string** | This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUsernamePermissionsRepositoriesRepoSlugGetRequest struct via the builder pattern


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
    resp, r, err := api_client.TeamsApi.TeamsWorkspaceRepositoriesGet(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsWorkspaceRepositoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsWorkspaceRepositoriesGet`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsWorkspaceRepositoriesGet`: %v\n", resp)
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


## UserPermissionsTeamsGet

> PaginatedTeamPermissions UserPermissionsTeamsGet(ctx).Q(q).Sort(sort).Execute()





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
    q := "q_example" // string |  Query string to narrow down the response as per [filtering and sorting](../../../meta/filtering). (optional)
    sort := "sort_example" // string |  Name of a response property sort the result by as per [filtering and sorting](../../../meta/filtering#query-sort).  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.UserPermissionsTeamsGet(context.Background()).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UserPermissionsTeamsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserPermissionsTeamsGet`: PaginatedTeamPermissions
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UserPermissionsTeamsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserPermissionsTeamsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  Query string to narrow down the response as per [filtering and sorting](../../../meta/filtering). | 
 **sort** | **string** |  Name of a response property sort the result by as per [filtering and sorting](../../../meta/filtering#query-sort).  | 

### Return type

[**PaginatedTeamPermissions**](PaginatedTeamPermissions.md)

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
    resp, r, err := api_client.TeamsApi.UsersWorkspaceRepositoriesGet(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UsersWorkspaceRepositoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersWorkspaceRepositoriesGet`: ModelError
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UsersWorkspaceRepositoriesGet`: %v\n", resp)
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

