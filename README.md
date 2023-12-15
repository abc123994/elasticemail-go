# Go API client for ElasticEmail

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.

Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.

The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.

To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://app.elasticemail.com/marketing/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.

Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 4.0.0
- Package version: 4.0.22
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import ElasticEmail "github.com/elasticemail/elasticemail-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), ElasticEmail.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), ElasticEmail.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), ElasticEmail.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), ElasticEmail.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.elasticemail.com/v4*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CampaignsApi* | [**CampaignsByNameDelete**](docs/CampaignsApi.md#campaignsbynamedelete) | **Delete** /campaigns/{name} | Delete Campaign
*CampaignsApi* | [**CampaignsByNameGet**](docs/CampaignsApi.md#campaignsbynameget) | **Get** /campaigns/{name} | Load Campaign
*CampaignsApi* | [**CampaignsByNamePut**](docs/CampaignsApi.md#campaignsbynameput) | **Put** /campaigns/{name} | Update Campaign
*CampaignsApi* | [**CampaignsGet**](docs/CampaignsApi.md#campaignsget) | **Get** /campaigns | Load Campaigns
*CampaignsApi* | [**CampaignsPost**](docs/CampaignsApi.md#campaignspost) | **Post** /campaigns | Add Campaign
*ContactsApi* | [**ContactsByEmailDelete**](docs/ContactsApi.md#contactsbyemaildelete) | **Delete** /contacts/{email} | Delete Contact
*ContactsApi* | [**ContactsByEmailGet**](docs/ContactsApi.md#contactsbyemailget) | **Get** /contacts/{email} | Load Contact
*ContactsApi* | [**ContactsByEmailPut**](docs/ContactsApi.md#contactsbyemailput) | **Put** /contacts/{email} | Update Contact
*ContactsApi* | [**ContactsDeletePost**](docs/ContactsApi.md#contactsdeletepost) | **Post** /contacts/delete | Delete Contacts Bulk
*ContactsApi* | [**ContactsExportByIdStatusGet**](docs/ContactsApi.md#contactsexportbyidstatusget) | **Get** /contacts/export/{id}/status | Check Export Status
*ContactsApi* | [**ContactsExportPost**](docs/ContactsApi.md#contactsexportpost) | **Post** /contacts/export | Export Contacts
*ContactsApi* | [**ContactsGet**](docs/ContactsApi.md#contactsget) | **Get** /contacts | Load Contacts
*ContactsApi* | [**ContactsImportPost**](docs/ContactsApi.md#contactsimportpost) | **Post** /contacts/import | Upload Contacts
*ContactsApi* | [**ContactsPost**](docs/ContactsApi.md#contactspost) | **Post** /contacts | Add Contact
*EmailsApi* | [**EmailsByMsgidViewGet**](docs/EmailsApi.md#emailsbymsgidviewget) | **Get** /emails/{msgid}/view | View Email
*EmailsApi* | [**EmailsMergefilePost**](docs/EmailsApi.md#emailsmergefilepost) | **Post** /emails/mergefile | Send Bulk Emails CSV
*EmailsApi* | [**EmailsPost**](docs/EmailsApi.md#emailspost) | **Post** /emails | Send Bulk Emails
*EmailsApi* | [**EmailsTransactionalPost**](docs/EmailsApi.md#emailstransactionalpost) | **Post** /emails/transactional | Send Transactional Email
*EventsApi* | [**EventsByTransactionidGet**](docs/EventsApi.md#eventsbytransactionidget) | **Get** /events/{transactionid} | Load Email Events
*EventsApi* | [**EventsChannelsByNameExportPost**](docs/EventsApi.md#eventschannelsbynameexportpost) | **Post** /events/channels/{name}/export | Export Channel Events
*EventsApi* | [**EventsChannelsByNameGet**](docs/EventsApi.md#eventschannelsbynameget) | **Get** /events/channels/{name} | Load Channel Events
*EventsApi* | [**EventsChannelsExportByIdStatusGet**](docs/EventsApi.md#eventschannelsexportbyidstatusget) | **Get** /events/channels/export/{id}/status | Check Channel Export Status
*EventsApi* | [**EventsExportByIdStatusGet**](docs/EventsApi.md#eventsexportbyidstatusget) | **Get** /events/export/{id}/status | Check Export Status
*EventsApi* | [**EventsExportPost**](docs/EventsApi.md#eventsexportpost) | **Post** /events/export | Export Events
*EventsApi* | [**EventsGet**](docs/EventsApi.md#eventsget) | **Get** /events | Load Events
*FilesApi* | [**FilesByNameDelete**](docs/FilesApi.md#filesbynamedelete) | **Delete** /files/{name} | Delete File
*FilesApi* | [**FilesByNameGet**](docs/FilesApi.md#filesbynameget) | **Get** /files/{name} | Download File
*FilesApi* | [**FilesByNameInfoGet**](docs/FilesApi.md#filesbynameinfoget) | **Get** /files/{name}/info | Load File Details
*FilesApi* | [**FilesGet**](docs/FilesApi.md#filesget) | **Get** /files | List Files
*FilesApi* | [**FilesPost**](docs/FilesApi.md#filespost) | **Post** /files | Upload File
*InboundRouteApi* | [**InboundrouteByIdDelete**](docs/InboundRouteApi.md#inboundroutebyiddelete) | **Delete** /inboundroute/{id} | Delete Route
*InboundRouteApi* | [**InboundrouteByIdGet**](docs/InboundRouteApi.md#inboundroutebyidget) | **Get** /inboundroute/{id} | Get Route
*InboundRouteApi* | [**InboundrouteByIdPut**](docs/InboundRouteApi.md#inboundroutebyidput) | **Put** /inboundroute/{id} | Update Route
*InboundRouteApi* | [**InboundrouteGet**](docs/InboundRouteApi.md#inboundrouteget) | **Get** /inboundroute | Get Routes
*InboundRouteApi* | [**InboundrouteOrderPut**](docs/InboundRouteApi.md#inboundrouteorderput) | **Put** /inboundroute/order | Update Sorting
*InboundRouteApi* | [**InboundroutePost**](docs/InboundRouteApi.md#inboundroutepost) | **Post** /inboundroute | Create Route
*ListsApi* | [**ListsByNameContactsPost**](docs/ListsApi.md#listsbynamecontactspost) | **Post** /lists/{name}/contacts | Add Contacts to List
*ListsApi* | [**ListsByNameContactsRemovePost**](docs/ListsApi.md#listsbynamecontactsremovepost) | **Post** /lists/{name}/contacts/remove | Remove Contacts from List
*ListsApi* | [**ListsByNameDelete**](docs/ListsApi.md#listsbynamedelete) | **Delete** /lists/{name} | Delete List
*ListsApi* | [**ListsByNameGet**](docs/ListsApi.md#listsbynameget) | **Get** /lists/{name} | Load List
*ListsApi* | [**ListsByNamePut**](docs/ListsApi.md#listsbynameput) | **Put** /lists/{name} | Update List
*ListsApi* | [**ListsGet**](docs/ListsApi.md#listsget) | **Get** /lists | Load Lists
*ListsApi* | [**ListsPost**](docs/ListsApi.md#listspost) | **Post** /lists | Add List
*SecurityApi* | [**SecurityApikeysByNameDelete**](docs/SecurityApi.md#securityapikeysbynamedelete) | **Delete** /security/apikeys/{name} | Delete ApiKey
*SecurityApi* | [**SecurityApikeysByNameGet**](docs/SecurityApi.md#securityapikeysbynameget) | **Get** /security/apikeys/{name} | Load ApiKey
*SecurityApi* | [**SecurityApikeysByNamePut**](docs/SecurityApi.md#securityapikeysbynameput) | **Put** /security/apikeys/{name} | Update ApiKey
*SecurityApi* | [**SecurityApikeysGet**](docs/SecurityApi.md#securityapikeysget) | **Get** /security/apikeys | List ApiKeys
*SecurityApi* | [**SecurityApikeysPost**](docs/SecurityApi.md#securityapikeyspost) | **Post** /security/apikeys | Add ApiKey
*SecurityApi* | [**SecuritySmtpByNameDelete**](docs/SecurityApi.md#securitysmtpbynamedelete) | **Delete** /security/smtp/{name} | Delete SMTP Credential
*SecurityApi* | [**SecuritySmtpByNameGet**](docs/SecurityApi.md#securitysmtpbynameget) | **Get** /security/smtp/{name} | Load SMTP Credential
*SecurityApi* | [**SecuritySmtpByNamePut**](docs/SecurityApi.md#securitysmtpbynameput) | **Put** /security/smtp/{name} | Update SMTP Credential
*SecurityApi* | [**SecuritySmtpGet**](docs/SecurityApi.md#securitysmtpget) | **Get** /security/smtp | List SMTP Credentials
*SecurityApi* | [**SecuritySmtpPost**](docs/SecurityApi.md#securitysmtppost) | **Post** /security/smtp | Add SMTP Credential
*SegmentsApi* | [**SegmentsByNameDelete**](docs/SegmentsApi.md#segmentsbynamedelete) | **Delete** /segments/{name} | Delete Segment
*SegmentsApi* | [**SegmentsByNameGet**](docs/SegmentsApi.md#segmentsbynameget) | **Get** /segments/{name} | Load Segment
*SegmentsApi* | [**SegmentsByNamePut**](docs/SegmentsApi.md#segmentsbynameput) | **Put** /segments/{name} | Update Segment
*SegmentsApi* | [**SegmentsGet**](docs/SegmentsApi.md#segmentsget) | **Get** /segments | Load Segments
*SegmentsApi* | [**SegmentsPost**](docs/SegmentsApi.md#segmentspost) | **Post** /segments | Add Segment
*StatisticsApi* | [**StatisticsCampaignsByNameGet**](docs/StatisticsApi.md#statisticscampaignsbynameget) | **Get** /statistics/campaigns/{name} | Load Campaign Stats
*StatisticsApi* | [**StatisticsCampaignsGet**](docs/StatisticsApi.md#statisticscampaignsget) | **Get** /statistics/campaigns | Load Campaigns Stats
*StatisticsApi* | [**StatisticsChannelsByNameGet**](docs/StatisticsApi.md#statisticschannelsbynameget) | **Get** /statistics/channels/{name} | Load Channel Stats
*StatisticsApi* | [**StatisticsChannelsGet**](docs/StatisticsApi.md#statisticschannelsget) | **Get** /statistics/channels | Load Channels Stats
*StatisticsApi* | [**StatisticsGet**](docs/StatisticsApi.md#statisticsget) | **Get** /statistics | Load Statistics
*SubAccountsApi* | [**SubaccountsByEmailCreditsPatch**](docs/SubAccountsApi.md#subaccountsbyemailcreditspatch) | **Patch** /subaccounts/{email}/credits | Add, Subtract Email Credits
*SubAccountsApi* | [**SubaccountsByEmailDelete**](docs/SubAccountsApi.md#subaccountsbyemaildelete) | **Delete** /subaccounts/{email} | Delete SubAccount
*SubAccountsApi* | [**SubaccountsByEmailGet**](docs/SubAccountsApi.md#subaccountsbyemailget) | **Get** /subaccounts/{email} | Load SubAccount
*SubAccountsApi* | [**SubaccountsByEmailSettingsEmailPut**](docs/SubAccountsApi.md#subaccountsbyemailsettingsemailput) | **Put** /subaccounts/{email}/settings/email | Update SubAccount Email Settings
*SubAccountsApi* | [**SubaccountsGet**](docs/SubAccountsApi.md#subaccountsget) | **Get** /subaccounts | Load SubAccounts
*SubAccountsApi* | [**SubaccountsPost**](docs/SubAccountsApi.md#subaccountspost) | **Post** /subaccounts | Add SubAccount
*SuppressionsApi* | [**SuppressionsBouncesGet**](docs/SuppressionsApi.md#suppressionsbouncesget) | **Get** /suppressions/bounces | Get Bounce List
*SuppressionsApi* | [**SuppressionsBouncesImportPost**](docs/SuppressionsApi.md#suppressionsbouncesimportpost) | **Post** /suppressions/bounces/import | Add Bounces Async
*SuppressionsApi* | [**SuppressionsBouncesPost**](docs/SuppressionsApi.md#suppressionsbouncespost) | **Post** /suppressions/bounces | Add Bounces
*SuppressionsApi* | [**SuppressionsByEmailDelete**](docs/SuppressionsApi.md#suppressionsbyemaildelete) | **Delete** /suppressions/{email} | Delete Suppression
*SuppressionsApi* | [**SuppressionsByEmailGet**](docs/SuppressionsApi.md#suppressionsbyemailget) | **Get** /suppressions/{email} | Get Suppression
*SuppressionsApi* | [**SuppressionsComplaintsGet**](docs/SuppressionsApi.md#suppressionscomplaintsget) | **Get** /suppressions/complaints | Get Complaints List
*SuppressionsApi* | [**SuppressionsComplaintsImportPost**](docs/SuppressionsApi.md#suppressionscomplaintsimportpost) | **Post** /suppressions/complaints/import | Add Complaints Async
*SuppressionsApi* | [**SuppressionsComplaintsPost**](docs/SuppressionsApi.md#suppressionscomplaintspost) | **Post** /suppressions/complaints | Add Complaints
*SuppressionsApi* | [**SuppressionsGet**](docs/SuppressionsApi.md#suppressionsget) | **Get** /suppressions | Get Suppressions
*SuppressionsApi* | [**SuppressionsUnsubscribesGet**](docs/SuppressionsApi.md#suppressionsunsubscribesget) | **Get** /suppressions/unsubscribes | Get Unsubscribes List
*SuppressionsApi* | [**SuppressionsUnsubscribesImportPost**](docs/SuppressionsApi.md#suppressionsunsubscribesimportpost) | **Post** /suppressions/unsubscribes/import | Add Unsubscribes Async
*SuppressionsApi* | [**SuppressionsUnsubscribesPost**](docs/SuppressionsApi.md#suppressionsunsubscribespost) | **Post** /suppressions/unsubscribes | Add Unsubscribes
*TemplatesApi* | [**TemplatesByNameDelete**](docs/TemplatesApi.md#templatesbynamedelete) | **Delete** /templates/{name} | Delete Template
*TemplatesApi* | [**TemplatesByNameGet**](docs/TemplatesApi.md#templatesbynameget) | **Get** /templates/{name} | Load Template
*TemplatesApi* | [**TemplatesByNamePut**](docs/TemplatesApi.md#templatesbynameput) | **Put** /templates/{name} | Update Template
*TemplatesApi* | [**TemplatesGet**](docs/TemplatesApi.md#templatesget) | **Get** /templates | Load Templates
*TemplatesApi* | [**TemplatesPost**](docs/TemplatesApi.md#templatespost) | **Post** /templates | Add Template
*VerificationsApi* | [**VerificationsByEmailDelete**](docs/VerificationsApi.md#verificationsbyemaildelete) | **Delete** /verifications/{email} | Delete Email Verification Result
*VerificationsApi* | [**VerificationsByEmailGet**](docs/VerificationsApi.md#verificationsbyemailget) | **Get** /verifications/{email} | Get Email Verification Result
*VerificationsApi* | [**VerificationsByEmailPost**](docs/VerificationsApi.md#verificationsbyemailpost) | **Post** /verifications/{email} | Verify Email
*VerificationsApi* | [**VerificationsFilesByIdDelete**](docs/VerificationsApi.md#verificationsfilesbyiddelete) | **Delete** /verifications/files/{id} | Delete File Verification Result
*VerificationsApi* | [**VerificationsFilesByIdResultDownloadGet**](docs/VerificationsApi.md#verificationsfilesbyidresultdownloadget) | **Get** /verifications/files/{id}/result/download | Download File Verification Result
*VerificationsApi* | [**VerificationsFilesByIdResultGet**](docs/VerificationsApi.md#verificationsfilesbyidresultget) | **Get** /verifications/files/{id}/result | Get Detailed File Verification Result
*VerificationsApi* | [**VerificationsFilesByIdVerificationPost**](docs/VerificationsApi.md#verificationsfilesbyidverificationpost) | **Post** /verifications/files/{id}/verification | Start verification
*VerificationsApi* | [**VerificationsFilesPost**](docs/VerificationsApi.md#verificationsfilespost) | **Post** /verifications/files | Upload File with Emails
*VerificationsApi* | [**VerificationsFilesResultGet**](docs/VerificationsApi.md#verificationsfilesresultget) | **Get** /verifications/files/result | Get Files Verification Results
*VerificationsApi* | [**VerificationsGet**](docs/VerificationsApi.md#verificationsget) | **Get** /verifications | Get Emails Verification Results


## Documentation For Models

 - [AccessLevel](docs/AccessLevel.md)
 - [AccountStatusEnum](docs/AccountStatusEnum.md)
 - [ApiKey](docs/ApiKey.md)
 - [ApiKeyNew](docs/ApiKeyNew.md)
 - [ApiKeyPayload](docs/ApiKeyPayload.md)
 - [BodyContentType](docs/BodyContentType.md)
 - [BodyPart](docs/BodyPart.md)
 - [Campaign](docs/Campaign.md)
 - [CampaignOptions](docs/CampaignOptions.md)
 - [CampaignRecipient](docs/CampaignRecipient.md)
 - [CampaignStatus](docs/CampaignStatus.md)
 - [CampaignTemplate](docs/CampaignTemplate.md)
 - [ChannelLogStatusSummary](docs/ChannelLogStatusSummary.md)
 - [CompressionFormat](docs/CompressionFormat.md)
 - [ConsentData](docs/ConsentData.md)
 - [ConsentTracking](docs/ConsentTracking.md)
 - [Contact](docs/Contact.md)
 - [ContactActivity](docs/ContactActivity.md)
 - [ContactPayload](docs/ContactPayload.md)
 - [ContactSource](docs/ContactSource.md)
 - [ContactStatus](docs/ContactStatus.md)
 - [ContactUpdatePayload](docs/ContactUpdatePayload.md)
 - [ContactsList](docs/ContactsList.md)
 - [DeliveryOptimizationType](docs/DeliveryOptimizationType.md)
 - [EmailContent](docs/EmailContent.md)
 - [EmailData](docs/EmailData.md)
 - [EmailMessageData](docs/EmailMessageData.md)
 - [EmailRecipient](docs/EmailRecipient.md)
 - [EmailSend](docs/EmailSend.md)
 - [EmailStatus](docs/EmailStatus.md)
 - [EmailTransactionalMessageData](docs/EmailTransactionalMessageData.md)
 - [EmailValidationResult](docs/EmailValidationResult.md)
 - [EmailValidationStatus](docs/EmailValidationStatus.md)
 - [EmailView](docs/EmailView.md)
 - [EmailsPayload](docs/EmailsPayload.md)
 - [EncodingType](docs/EncodingType.md)
 - [EventType](docs/EventType.md)
 - [EventsOrderBy](docs/EventsOrderBy.md)
 - [ExportFileFormats](docs/ExportFileFormats.md)
 - [ExportLink](docs/ExportLink.md)
 - [ExportStatus](docs/ExportStatus.md)
 - [FileInfo](docs/FileInfo.md)
 - [FilePayload](docs/FilePayload.md)
 - [FileUploadResult](docs/FileUploadResult.md)
 - [InboundPayload](docs/InboundPayload.md)
 - [InboundRoute](docs/InboundRoute.md)
 - [InboundRouteActionType](docs/InboundRouteActionType.md)
 - [InboundRouteFilterType](docs/InboundRouteFilterType.md)
 - [ListPayload](docs/ListPayload.md)
 - [ListUpdatePayload](docs/ListUpdatePayload.md)
 - [LogJobStatus](docs/LogJobStatus.md)
 - [LogStatusSummary](docs/LogStatusSummary.md)
 - [MergeEmailPayload](docs/MergeEmailPayload.md)
 - [MessageAttachment](docs/MessageAttachment.md)
 - [MessageCategory](docs/MessageCategory.md)
 - [Options](docs/Options.md)
 - [RecipientEvent](docs/RecipientEvent.md)
 - [Segment](docs/Segment.md)
 - [SegmentPayload](docs/SegmentPayload.md)
 - [SmtpCredentials](docs/SmtpCredentials.md)
 - [SmtpCredentialsNew](docs/SmtpCredentialsNew.md)
 - [SmtpCredentialsPayload](docs/SmtpCredentialsPayload.md)
 - [SortOrderItem](docs/SortOrderItem.md)
 - [SplitOptimizationType](docs/SplitOptimizationType.md)
 - [SplitOptions](docs/SplitOptions.md)
 - [SubAccountInfo](docs/SubAccountInfo.md)
 - [SubaccountEmailCreditsPayload](docs/SubaccountEmailCreditsPayload.md)
 - [SubaccountEmailSettings](docs/SubaccountEmailSettings.md)
 - [SubaccountEmailSettingsPayload](docs/SubaccountEmailSettingsPayload.md)
 - [SubaccountPayload](docs/SubaccountPayload.md)
 - [SubaccountSettingsInfo](docs/SubaccountSettingsInfo.md)
 - [SubaccountSettingsInfoPayload](docs/SubaccountSettingsInfoPayload.md)
 - [Suppression](docs/Suppression.md)
 - [Template](docs/Template.md)
 - [TemplatePayload](docs/TemplatePayload.md)
 - [TemplateScope](docs/TemplateScope.md)
 - [TemplateType](docs/TemplateType.md)
 - [TransactionalRecipient](docs/TransactionalRecipient.md)
 - [Utm](docs/Utm.md)
 - [VerificationFileResult](docs/VerificationFileResult.md)
 - [VerificationFileResultDetails](docs/VerificationFileResultDetails.md)
 - [VerificationStatus](docs/VerificationStatus.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### apikey

- **Type**: API key
- **API key parameter name**: X-ElasticEmail-ApiKey
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-ElasticEmail-ApiKey and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@elasticemail.com

