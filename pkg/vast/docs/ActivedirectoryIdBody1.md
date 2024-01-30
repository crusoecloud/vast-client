# ActivedirectoryIdBody1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminUsername** | **string** | An Active Directory admin user with permission to join the Active Directory server. | [default to null]
**AdminPasswd** | **string** | The password for the specified Active Directory admin user. | [default to null]
**Enabled** | **bool** | Set to true to join Active Directory. Set to false to leave Active Directory. | [default to null]
**SmbAllowed** | **bool** | Indicates if AD is allowed for SMB. | [optional] [default to false]
**NtlmEnabled** | **bool** | Manages support of NTLM authentication method for SMB protocol. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

