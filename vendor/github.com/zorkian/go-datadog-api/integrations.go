/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2018 by authors and contributors.
 */

package datadog

import "net/url"

/*
	PagerDuty Integration
*/

type servicePD struct {
	ServiceName *string `json:"service_name"`
	ServiceKey  *string `json:"service_key"`
}

type integrationPD struct {
	Services  []servicePD `json:"services"`
	Subdomain *string     `json:"subdomain"`
	Schedules []string    `json:"schedules"`
	APIToken  *string     `json:"api_token"`
}

// ServicePDRequest defines the Services struct that is part of the IntegrationPDRequest.
type ServicePDRequest struct {
	ServiceName *string `json:"service_name"`
	ServiceKey  *string `json:"service_key"`
}

// IntegrationPDRequest defines the request payload for
// creating & updating Datadog-PagerDuty integration.
type IntegrationPDRequest struct {
	Services  []ServicePDRequest `json:"services,omitempty"`
	Subdomain *string            `json:"subdomain,omitempty"`
	Schedules []string           `json:"schedules,omitempty"`
	APIToken  *string            `json:"api_token,omitempty"`
	RunCheck  *bool              `json:"run_check,omitempty"`
}

// CreateIntegrationPD creates new PagerDuty Integrations.
// Use this if you want to setup the integration for the first time
// or to add more services/schedules.
func (client *Client) CreateIntegrationPD(pdIntegration *IntegrationPDRequest) error {
	return client.doJsonRequest("PUT", "/v1/integration/pagerduty", pdIntegration, nil)
}

// UpdateIntegrationPD updates the PagerDuty Integration.
// This will replace the existing values with the new values.
func (client *Client) UpdateIntegrationPD(pdIntegration *IntegrationPDRequest) error {
	return client.doJsonRequest("PUT", "/v1/integration/pagerduty", pdIntegration, nil)
}

// GetIntegrationPD gets all the PagerDuty Integrations from the system.
func (client *Client) GetIntegrationPD() (*integrationPD, error) {
	var out integrationPD
	if err := client.doJsonRequest("GET", "/v1/integration/pagerduty", nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// DeleteIntegrationPD removes the PagerDuty Integration from the system.
func (client *Client) DeleteIntegrationPD() error {
	return client.doJsonRequest("DELETE", "/v1/integration/pagerduty", nil, nil)
}

// CreateIntegrationPDService creates a single service object in the PagerDuty integration
// Note that creating a service object requires the integration to be activated
func (client *Client) CreateIntegrationPDService(serviceObject *ServicePDRequest) error {
	return client.doJsonRequest("POST", "/v1/integration/pagerduty/configuration/services", serviceObject, nil)
}

// UpdateIntegrationPDService updates a single service object in the PagerDuty integration
func (client *Client) UpdateIntegrationPDService(serviceObject *ServicePDRequest) error {
	// we can only post the ServiceKey, not ServiceName
	toPost := struct {
		ServiceKey *string `json:"service_key,omitempty"`
	}{
		serviceObject.ServiceKey,
	}
	uri := "/v1/integration/pagerduty/configuration/services/" + *serviceObject.ServiceName
	return client.doJsonRequest("PUT", uri, toPost, nil)
}

// GetIntegrationPDService gets a single service object in the PagerDuty integration
// NOTE: the service key is never returned by the API, so it won't be set
func (client *Client) GetIntegrationPDService(serviceName string) (*ServicePDRequest, error) {
	uri := "/v1/integration/pagerduty/configuration/services/" + serviceName
	var out ServicePDRequest
	if err := client.doJsonRequest("GET", uri, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteIntegrationPDService deletes a single service object in the PagerDuty integration
func (client *Client) DeleteIntegrationPDService(serviceName string) error {
	uri := "/v1/integration/pagerduty/configuration/services/" + serviceName
	return client.doJsonRequest("DELETE", uri, nil, nil)
}

/*
	Slack Integration
*/

// ServiceHookSlackRequest defines the ServiceHooks struct that is part of the IntegrationSlackRequest.
type ServiceHookSlackRequest struct {
	Account *string `json:"account"`
	Url     *string `json:"url"`
}

// ChannelSlackRequest defines the Channels struct that is part of the IntegrationSlackRequest.
type ChannelSlackRequest struct {
	ChannelName             *string `json:"channel_name"`
	TransferAllUserComments *bool   `json:"transfer_all_user_comments,omitempty,string"`
	Account                 *string `json:"account"`
}

// IntegrationSlackRequest defines the request payload for
// creating & updating Datadog-Slack integration.
type IntegrationSlackRequest struct {
	ServiceHooks []ServiceHookSlackRequest `json:"service_hooks,omitempty"`
	Channels     []ChannelSlackRequest     `json:"channels,omitempty"`
	RunCheck     *bool                     `json:"run_check,omitempty,string"`
}

// CreateIntegrationSlack creates new Slack Integrations.
// Use this if you want to setup the integration for the first time
// or to add more channels.
func (client *Client) CreateIntegrationSlack(slackIntegration *IntegrationSlackRequest) error {
	return client.doJsonRequest("POST", "/v1/integration/slack", slackIntegration, nil)
}

// UpdateIntegrationSlack updates the Slack Integration.
// This will replace the existing values with the new values.
func (client *Client) UpdateIntegrationSlack(slackIntegration *IntegrationSlackRequest) error {
	return client.doJsonRequest("PUT", "/v1/integration/slack", slackIntegration, nil)
}

// GetIntegrationSlack gets all the Slack Integrations from the system.
func (client *Client) GetIntegrationSlack() (*IntegrationSlackRequest, error) {
	var out IntegrationSlackRequest
	if err := client.doJsonRequest("GET", "/v1/integration/slack", nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// DeleteIntegrationSlack removes the Slack Integration from the system.
func (client *Client) DeleteIntegrationSlack() error {
	return client.doJsonRequest("DELETE", "/v1/integration/slack", nil, nil)
}

/*
	AWS Integration
*/

// IntegrationAWSAccount defines the request payload for
// creating & updating Datadog-AWS integration.
type IntegrationAWSAccount struct {
	AccountID                     *string         `json:"account_id"`
	RoleName                      *string         `json:"role_name"`
	FilterTags                    []string        `json:"filter_tags"`
	HostTags                      []string        `json:"host_tags"`
	AccountSpecificNamespaceRules map[string]bool `json:"account_specific_namespace_rules"`
}

// IntegrationAWSAccountCreateResponse defines the response payload for
// creating & updating Datadog-AWS integration.
type IntegrationAWSAccountCreateResponse struct {
	ExternalID string `json:"external_id"`
}

type IntegrationAWSAccountGetResponse struct {
	Accounts []IntegrationAWSAccount `json:"accounts"`
}

type IntegrationAWSAccountDeleteRequest struct {
	AccountID *string `json:"account_id"`
	RoleName  *string `json:"role_name"`
}

type IntegrationAWSLambdaARNRequest struct {
	AccountID *string `json:"account_id"`
	LambdaARN *string `json:"lambda_arn"`
}

// IntegrationAWSLambdaARN is only defined to properly parse the AWS logs GET response
type IntegrationAWSLambdaARN struct {
	LambdaARN *string `json:"arn"`
}

type IntegrationAWSServicesLogCollection struct {
	AccountID *string  `json:"account_id"`
	Services  []string `json:"services"`
}

// IntegrationAWSLogs is only defined to properly parse the AWS logs GET response
type IntegrationAWSLogCollection struct {
	AccountID  *string                   `json:"account_id"`
	LambdaARNs []IntegrationAWSLambdaARN `json:"lambdas"`
	Services   []string                  `json:"services"`
}

// CreateIntegrationAWS adds a new AWS Account in the AWS Integrations.
// Use this if you want to setup the integration for the first time
// or to add more accounts.
func (client *Client) CreateIntegrationAWS(awsAccount *IntegrationAWSAccount) (*IntegrationAWSAccountCreateResponse, error) {
	var out IntegrationAWSAccountCreateResponse
	if err := client.doJsonRequest("POST", "/v1/integration/aws", awsAccount, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// UpdateIntegrationAWS updates an already existing AWS Account in the AWS Integration
func (client *Client) UpdateIntegrationAWS(awsAccount *IntegrationAWSAccount) error {
	additionalParameters := url.Values{}
	additionalParameters.Set("account_id", *awsAccount.AccountID)
	additionalParameters.Add("role_name", *awsAccount.RoleName)
	return client.doJsonRequest("PUT", "/v1/integration/aws?"+additionalParameters.Encode(), awsAccount, nil)
}

// GetIntegrationAWS gets all the AWS Accounts in the AWS Integrations from Datadog.
func (client *Client) GetIntegrationAWS() (*[]IntegrationAWSAccount, error) {
	var response IntegrationAWSAccountGetResponse
	if err := client.doJsonRequest("GET", "/v1/integration/aws", nil, &response); err != nil {
		return nil, err
	}

	return &response.Accounts, nil
}

// DeleteIntegrationAWS removes a specific AWS Account from the AWS Integration.
func (client *Client) DeleteIntegrationAWS(awsAccount *IntegrationAWSAccountDeleteRequest) error {
	return client.doJsonRequest("DELETE", "/v1/integration/aws", awsAccount, nil)
}

// AttachLambdaARNIntegrationAWS attach a lambda ARN to an AWS account ID to enable log collection
func (client *Client) AttachLambdaARNIntegrationAWS(lambdaARN *IntegrationAWSLambdaARNRequest) error {
	return client.doJsonRequest("POST", "/v1/integration/aws/logs", lambdaARN, nil)
}

// EnableLogCollectionAWSServices enables the log collection for the given AWS services
func (client *Client) EnableLogCollectionAWSServices(services *IntegrationAWSServicesLogCollection) error {
	return client.doJsonRequest("POST", "/v1/integration/aws/logs/services", services, nil)
}

// GetIntegrationAWSLogCollection gets all the configuration for the AWS log collection
func (client *Client) GetIntegrationAWSLogCollection() (*[]IntegrationAWSLogCollection, error) {
	var response []IntegrationAWSLogCollection
	if err := client.doJsonRequest("GET", "/v1/integration/aws/logs", nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteAWSLogCollection removes the log collection configuration for a given ARN and AWS account
func (client *Client) DeleteAWSLogCollection(lambdaARN *IntegrationAWSLambdaARNRequest) error {
	return client.doJsonRequest("DELETE", "/v1/integration/aws/logs", lambdaARN, nil)
}

/*
	Google Cloud Platform Integration
*/

// IntegrationGCP defines the response for listing Datadog-Google CloudPlatform integration.
type IntegrationGCP struct {
	ProjectID   *string `json:"project_id"`
	ClientEmail *string `json:"client_email"`
	HostFilters *string `json:"host_filters"`
}

// IntegrationGCPCreateRequest defines the request payload for creating Datadog-Google CloudPlatform integration.
type IntegrationGCPCreateRequest struct {
	Type                    *string `json:"type"` // Should be service_account
	ProjectID               *string `json:"project_id"`
	PrivateKeyID            *string `json:"private_key_id"`
	PrivateKey              *string `json:"private_key"`
	ClientEmail             *string `json:"client_email"`
	ClientID                *string `json:"client_id"`
	AuthURI                 *string `json:"auth_uri"`                    // Should be https://accounts.google.com/o/oauth2/auth
	TokenURI                *string `json:"token_uri"`                   // Should be https://accounts.google.com/o/oauth2/token
	AuthProviderX509CertURL *string `json:"auth_provider_x509_cert_url"` // Should be https://www.googleapis.com/oauth2/v1/certs
	ClientX509CertURL       *string `json:"client_x509_cert_url"`        // https://www.googleapis.com/robot/v1/metadata/x509/<CLIENT_EMAIL>
	HostFilters             *string `json:"host_filters,omitempty"`
}

// IntegrationGCPUpdateRequest defines the request payload for updating Datadog-Google CloudPlatform integration.
type IntegrationGCPUpdateRequest struct {
	ProjectID   *string `json:"project_id"`
	ClientEmail *string `json:"client_email"`
	HostFilters *string `json:"host_filters,omitempty"`
}

// IntegrationGCPDeleteRequest defines the request payload for deleting Datadog-Google CloudPlatform integration.
type IntegrationGCPDeleteRequest struct {
	ProjectID   *string `json:"project_id"`
	ClientEmail *string `json:"client_email"`
}

// ListIntegrationGCP gets all Google Cloud Platform Integrations.
func (client *Client) ListIntegrationGCP() ([]*IntegrationGCP, error) {
	var list []*IntegrationGCP
	if err := client.doJsonRequest("GET", "/v1/integration/gcp", nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

// CreateIntegrationGCP creates a new Google Cloud Platform Integration.
func (client *Client) CreateIntegrationGCP(cir *IntegrationGCPCreateRequest) error {
	return client.doJsonRequest("POST", "/v1/integration/gcp", cir, nil)
}

// UpdateIntegrationGCP updates a Google Cloud Platform Integration.
func (client *Client) UpdateIntegrationGCP(cir *IntegrationGCPUpdateRequest) error {
	return client.doJsonRequest("POST", "/v1/integration/gcp/host_filters", cir, nil)
}

// DeleteIntegrationGCP deletes a Google Cloud Platform Integration.
func (client *Client) DeleteIntegrationGCP(cir *IntegrationGCPDeleteRequest) error {
	return client.doJsonRequest("DELETE", "/v1/integration/gcp", cir, nil)
}
