// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package proton provides the client and types for making API
// requests to AWS Proton.
//
// This is the AWS Proton Service API Reference. It provides descriptions, syntax
// and usage examples for each of the actions (https://docs.aws.amazon.com/proton/latest/APIReference/API_Operations.html)
// and data types (https://docs.aws.amazon.com/proton/latest/APIReference/API_Types.html)
// for the AWS Proton service.
//
// The documentation for each action shows the Query API request parameters
// and the XML response.
//
// Alternatively, you can use the AWS CLI to access an API. For more information,
// see the AWS Command Line Interface User Guide (https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html).
//
// The AWS Proton service is a two-pronged automation framework. Administrators
// create service templates to provide standardized infrastructure and deployment
// tooling for serverless and container based applications. Developers, in turn,
// select from the available service templates to automate their application
// or service deployments.
//
// Because administrators define the infrastructure and tooling that AWS Proton
// deploys and manages, they need permissions to use all of the listed API operations.
//
// When developers select a specific infrastructure and tooling set, AWS Proton
// deploys their applications. To monitor their applications that are running
// on AWS Proton, developers need permissions to the service create, list, update
// and delete API operations and the service instance list and update API operations.
//
// To learn more about AWS Proton administration, see the AWS Proton Administrator
// Guide (https://docs.aws.amazon.com/proton/latest/adminguide/Welcome.html).
//
// To learn more about deploying serverless and containerized applications on
// AWS Proton, see the AWS Proton User Guide (https://docs.aws.amazon.com/proton/latest/userguide/Welcome.html).
//
// Ensuring Idempotency
//
// When you make a mutating API request, the request typically returns a result
// before the asynchronous workflows of the operation are complete. Operations
// might also time out or encounter other server issues before they're complete,
// even if the request already returned a result. This might make it difficult
// to determine whether the request succeeded. Moreover, you might need to retry
// the request multiple times to ensure that the operation completes successfully.
// However, if the original request and the subsequent retries are successful,
// the operation occurs multiple times. This means that you might create more
// resources than you intended.
//
// Idempotency ensures that an API request action completes no more than one
// time. With an idempotent request, if the original request action completes
// successfully, any subsequent retries complete successfully without performing
// any further actions. However, the result might contain updated information,
// such as the current creation status.
//
// The following lists of APIs are grouped according to methods that ensure
// idempotency.
//
// Idempotent create APIs with a client token
//
// The API actions in this list support idempotency with the use of a client
// token. The corresponding AWS CLI commands also support idempotency using
// a client token. A client token is a unique, case-sensitive string of up to
// 64 ASCII characters. To make an idempotent API request using one of these
// actions, specify a client token in the request. We recommend that you don't
// reuse the same client token for other API requests. If you don’t provide
// a client token for these APIs, a default client token is automatically provided
// by SDKs.
//
// Given a request action that has succeeded:
//
// If you retry the request using the same client token and the same parameters,
// the retry succeeds without performing any further actions other than returning
// the original resource detail data in the response.
//
// If you retry the request using the same client token, but one or more of
// the parameters are different, the retry throws a ValidationException with
// an IdempotentParameterMismatch error.
//
// Client tokens expire eight hours after a request is made. If you retry the
// request with the expired token, a new resource is created.
//
// If the original resource is deleted and you retry the request, a new resource
// is created.
//
// Idempotent create APIs with a client token:
//
//    * CreateEnvironmentTemplateVersion
//
//    * CreateServiceTemplateVersion
//
//    * CreateEnvironmentAccountConnection
//
// Idempotent delete APIs
//
// Given a request action that has succeeded:
//
// When you retry the request with an API from this group and the resource was
// deleted, its metadata is returned in the response.
//
// If you retry and the resource doesn't exist, the response is empty.
//
// In both cases, the retry succeeds.
//
// Idempotent delete APIs:
//
//    * DeleteEnvironmentTemplate
//
//    * DeleteEnvironmentTemplateVersion
//
//    * DeleteServiceTemplate
//
//    * DeleteServiceTemplateVersion
//
//    * DeleteEnvironmentAccountConnection
//
// Asynchronous idempotent delete APIs
//
// Given a request action that has succeeded:
//
// If you retry the request with an API from this group, if the original request
// delete operation status is DELETE_IN_PROGRESS, the retry returns the resource
// detail data in the response without performing any further actions.
//
// If the original request delete operation is complete, a retry returns an
// empty response.
//
// Asynchronous idempotent delete APIs:
//
//    * DeleteEnvironment
//
//    * DeleteService
//
// See https://docs.aws.amazon.com/goto/WebAPI/proton-2020-07-20 for more information on this service.
//
// See proton package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/proton/
//
// Using the Client
//
// To contact AWS Proton with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Proton client Proton for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/proton/#New
package proton
