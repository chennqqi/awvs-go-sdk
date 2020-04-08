# \DefaultApi

All URIs are relative to *https://localhost:10443//api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTarget**](DefaultApi.md#AddTarget) | **Post** /targets | Add a new target to the scan list
[**DeleteScan**](DefaultApi.md#DeleteScan) | **Delete** /scans/{scanid} | delete scan by scanid
[**GetInfo**](DefaultApi.md#GetInfo) | **Get** /info | get awvs info
[**GetMe**](DefaultApi.md#GetMe) | **Get** /me | get user info
[**GetScanProfiles**](DefaultApi.md#GetScanProfiles) | **Get** /scanning_profiles | get scan profile
[**GetScanReports**](DefaultApi.md#GetScanReports) | **Get** /reports/{scanid} | get scan reports by scanid
[**GetScanStatus**](DefaultApi.md#GetScanStatus) | **Get** /scans/{scanid} | get scan status by scanid
[**GetScans**](DefaultApi.md#GetScans) | **Get** /scans | get scan list
[**GetStat**](DefaultApi.md#GetStat) | **Get** /scans/{scanid}/results/{sessionid}/statistics | get stat by scanid,sessionid
[**GetTarget**](DefaultApi.md#GetTarget) | **Get** /targets/{targetid} | get target by id
[**GetTargets**](DefaultApi.md#GetTargets) | **Get** /targets | get all targets
[**GetVuln**](DefaultApi.md#GetVuln) | **Get** /scans/{scanid}/results/{sessionid}/vulnerabilities | get results by scanid,sessionid
[**Login**](DefaultApi.md#Login) | **Post** /me/login | login
[**StartScan**](DefaultApi.md#StartScan) | **Post** /scans | start scan by scanid
[**StopScan**](DefaultApi.md#StopScan) | **Post** /scans/{scanid}/abort | stop scan by scanid



## AddTarget

> TargetResp AddTarget(ctx, target)

Add a new target to the scan list

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | [**Target**](Target.md)| target parameter | 

### Return type

[**TargetResp**](TargetResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScan

> DeleteScan(ctx, scanid)

delete scan by scanid

https://github.com/bit4woo/Ashe/blob/master/lib/WVS11.py

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanid** | [**string**](.md)| scan task id | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfo

> Info GetInfo(ctx, )

get awvs info

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Info**](Info.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMe

> Me GetMe(ctx, )

get user info

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Me**](Me.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScanProfiles

> Profiles GetScanProfiles(ctx, )

get scan profile

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Profiles**](Profiles.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScanReports

> GetScanReports(ctx, scanid)

get scan reports by scanid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanid** | [**string**](.md)| scan task id | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScanStatus

> ScanDetail GetScanStatus(ctx, scanid)

get scan status by scanid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanid** | [**string**](.md)| scan task id | 

### Return type

[**ScanDetail**](ScanDetail.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScans

> Scans GetScans(ctx, )

get scan list

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Scans**](Scans.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStat

> ScanStat GetStat(ctx, scanid, sessionid)

get stat by scanid,sessionid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanid** | [**string**](.md)| scan task id | 
**sessionid** | **string**| scan session id | 

### Return type

[**ScanStat**](ScanStat.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTarget

> TargetDetail GetTarget(ctx, targetid)

get target by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetid** | [**string**](.md)| target id | 

### Return type

[**TargetDetail**](TargetDetail.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTargets

> TargetsResp GetTargets(ctx, )

get all targets

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**TargetsResp**](TargetsResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVuln

> GetVuln(ctx, scanid, sessionid)

get results by scanid,sessionid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanid** | [**string**](.md)| scan task id | 
**sessionid** | **string**| scan session id | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> Login(ctx, loginReq)

login

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loginReq** | [**LoginReq**](LoginReq.md)| login parameter | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartScan

> StartScan(ctx, scan)

start scan by scanid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scan** | [**Scan**](Scan.md)| target response | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopScan

> StopScan(ctx, scanid)

stop scan by scanid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanid** | [**string**](.md)| scan task id | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

