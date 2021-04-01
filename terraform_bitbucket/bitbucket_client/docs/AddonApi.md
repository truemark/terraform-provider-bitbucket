# \AddonApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddonDelete**](AddonApi.md#AddonDelete) | **Delete** /addon | 
[**AddonLinkersGet**](AddonApi.md#AddonLinkersGet) | **Get** /addon/linkers | 
[**AddonLinkersLinkerKeyGet**](AddonApi.md#AddonLinkersLinkerKeyGet) | **Get** /addon/linkers/{linker_key} | 
[**AddonLinkersLinkerKeyValuesDelete**](AddonApi.md#AddonLinkersLinkerKeyValuesDelete) | **Delete** /addon/linkers/{linker_key}/values | 
[**AddonLinkersLinkerKeyValuesGet**](AddonApi.md#AddonLinkersLinkerKeyValuesGet) | **Get** /addon/linkers/{linker_key}/values | 
[**AddonLinkersLinkerKeyValuesPost**](AddonApi.md#AddonLinkersLinkerKeyValuesPost) | **Post** /addon/linkers/{linker_key}/values | 
[**AddonLinkersLinkerKeyValuesPut**](AddonApi.md#AddonLinkersLinkerKeyValuesPut) | **Put** /addon/linkers/{linker_key}/values | 
[**AddonLinkersLinkerKeyValuesValueIdDelete**](AddonApi.md#AddonLinkersLinkerKeyValuesValueIdDelete) | **Delete** /addon/linkers/{linker_key}/values/{value_id} | 
[**AddonLinkersLinkerKeyValuesValueIdGet**](AddonApi.md#AddonLinkersLinkerKeyValuesValueIdGet) | **Get** /addon/linkers/{linker_key}/values/{value_id} | 
[**AddonPut**](AddonApi.md#AddonPut) | **Put** /addon | 



## AddonDelete

> AddonDelete(ctx).Execute()





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
    resp, r, err := api_client.AddonApi.AddonDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonApi.AddonDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAddonDeleteRequest struct via the builder pattern


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


## AddonLinkersGet

> AddonLinkersGet(ctx).Execute()





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
    resp, r, err := api_client.AddonApi.AddonLinkersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonApi.AddonLinkersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAddonLinkersGetRequest struct via the builder pattern


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


## AddonLinkersLinkerKeyGet

> AddonLinkersLinkerKeyGet(ctx, linkerKey).Execute()





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
    linkerKey := "linkerKey_example" // string | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddonApi.AddonLinkersLinkerKeyGet(context.Background(), linkerKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonApi.AddonLinkersLinkerKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkerKey** | **string** | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonLinkersLinkerKeyGetRequest struct via the builder pattern


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


## AddonLinkersLinkerKeyValuesDelete

> AddonLinkersLinkerKeyValuesDelete(ctx, linkerKey).Execute()





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
    linkerKey := "linkerKey_example" // string | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddonApi.AddonLinkersLinkerKeyValuesDelete(context.Background(), linkerKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonApi.AddonLinkersLinkerKeyValuesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkerKey** | **string** | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonLinkersLinkerKeyValuesDeleteRequest struct via the builder pattern


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


## AddonLinkersLinkerKeyValuesGet

> AddonLinkersLinkerKeyValuesGet(ctx, linkerKey).Execute()





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
    linkerKey := "linkerKey_example" // string | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddonApi.AddonLinkersLinkerKeyValuesGet(context.Background(), linkerKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonApi.AddonLinkersLinkerKeyValuesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkerKey** | **string** | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonLinkersLinkerKeyValuesGetRequest struct via the builder pattern


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


## AddonLinkersLinkerKeyValuesPost

> AddonLinkersLinkerKeyValuesPost(ctx, linkerKey).Execute()





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
    linkerKey := "linkerKey_example" // string | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddonApi.AddonLinkersLinkerKeyValuesPost(context.Background(), linkerKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonApi.AddonLinkersLinkerKeyValuesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkerKey** | **string** | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonLinkersLinkerKeyValuesPostRequest struct via the builder pattern


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


## AddonLinkersLinkerKeyValuesPut

> AddonLinkersLinkerKeyValuesPut(ctx, linkerKey).Execute()





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
    linkerKey := "linkerKey_example" // string | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddonApi.AddonLinkersLinkerKeyValuesPut(context.Background(), linkerKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonApi.AddonLinkersLinkerKeyValuesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkerKey** | **string** | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonLinkersLinkerKeyValuesPutRequest struct via the builder pattern


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


## AddonLinkersLinkerKeyValuesValueIdDelete

> AddonLinkersLinkerKeyValuesValueIdDelete(ctx, linkerKey, valueId).Execute()





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
    linkerKey := "linkerKey_example" // string | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor.
    valueId := int32(56) // int32 | The numeric ID of the linker value.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddonApi.AddonLinkersLinkerKeyValuesValueIdDelete(context.Background(), linkerKey, valueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonApi.AddonLinkersLinkerKeyValuesValueIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkerKey** | **string** | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 
**valueId** | **int32** | The numeric ID of the linker value. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonLinkersLinkerKeyValuesValueIdDeleteRequest struct via the builder pattern


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


## AddonLinkersLinkerKeyValuesValueIdGet

> AddonLinkersLinkerKeyValuesValueIdGet(ctx, linkerKey, valueId).Execute()





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
    linkerKey := "linkerKey_example" // string | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor.
    valueId := int32(56) // int32 | The numeric ID of the linker value.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddonApi.AddonLinkersLinkerKeyValuesValueIdGet(context.Background(), linkerKey, valueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonApi.AddonLinkersLinkerKeyValuesValueIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkerKey** | **string** | The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 
**valueId** | **int32** | The numeric ID of the linker value. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonLinkersLinkerKeyValuesValueIdGetRequest struct via the builder pattern


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


## AddonPut

> AddonPut(ctx).Execute()





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
    resp, r, err := api_client.AddonApi.AddonPut(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonApi.AddonPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAddonPutRequest struct via the builder pattern


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

