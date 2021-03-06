package billing

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PoliciesClient is the billing client provides access to billing resources for Azure subscriptions.
type PoliciesClient struct {
	BaseClient
}

// NewPoliciesClient creates an instance of the PoliciesClient client.
func NewPoliciesClient(subscriptionID string) PoliciesClient {
	return NewPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPoliciesClientWithBaseURI creates an instance of the PoliciesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPoliciesClientWithBaseURI(baseURI string, subscriptionID string) PoliciesClient {
	return PoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetByBillingProfile lists the policies for a billing profile. This operation is supported only for billing accounts
// with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
func (client PoliciesClient) GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result Policy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PoliciesClient.GetByBillingProfile")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByBillingProfilePreparer(ctx, billingAccountName, billingProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PoliciesClient", "GetByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PoliciesClient", "GetByBillingProfile", resp, "Failure sending request")
		return
	}

	result, err = client.GetByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PoliciesClient", "GetByBillingProfile", resp, "Failure responding to request")
	}

	return
}

// GetByBillingProfilePreparer prepares the GetByBillingProfile request.
func (client PoliciesClient) GetByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/policies/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByBillingProfileSender sends the GetByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client PoliciesClient) GetByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByBillingProfileResponder handles the response to the GetByBillingProfile request. The method always
// closes the http.Response Body.
func (client PoliciesClient) GetByBillingProfileResponder(resp *http.Response) (result Policy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByCustomer lists the policies for a customer. This operation is supported only for billing accounts with
// agreement type Microsoft Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// customerName - the ID that uniquely identifies a customer.
func (client PoliciesClient) GetByCustomer(ctx context.Context, billingAccountName string, customerName string) (result CustomerPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PoliciesClient.GetByCustomer")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByCustomerPreparer(ctx, billingAccountName, customerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PoliciesClient", "GetByCustomer", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByCustomerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PoliciesClient", "GetByCustomer", resp, "Failure sending request")
		return
	}

	result, err = client.GetByCustomerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PoliciesClient", "GetByCustomer", resp, "Failure responding to request")
	}

	return
}

// GetByCustomerPreparer prepares the GetByCustomer request.
func (client PoliciesClient) GetByCustomerPreparer(ctx context.Context, billingAccountName string, customerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"customerName":       autorest.Encode("path", customerName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/policies/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByCustomerSender sends the GetByCustomer request. The method will close the
// http.Response Body if it receives an error.
func (client PoliciesClient) GetByCustomerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByCustomerResponder handles the response to the GetByCustomer request. The method always
// closes the http.Response Body.
func (client PoliciesClient) GetByCustomerResponder(resp *http.Response) (result CustomerPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates the policies for a billing profile. This operation is supported only for billing accounts with
// agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// parameters - request parameters that are provided to the update policies operation.
func (client PoliciesClient) Update(ctx context.Context, billingAccountName string, billingProfileName string, parameters Policy) (result Policy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PoliciesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, billingAccountName, billingProfileName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PoliciesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PoliciesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PoliciesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PoliciesClient) UpdatePreparer(ctx context.Context, billingAccountName string, billingProfileName string, parameters Policy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/policies/default", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PoliciesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PoliciesClient) UpdateResponder(resp *http.Response) (result Policy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateCustomer updates the policies for a customer. This operation is supported only for billing accounts with
// agreement type Microsoft Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// customerName - the ID that uniquely identifies a customer.
// parameters - request parameters that are provided to the update policies operation.
func (client PoliciesClient) UpdateCustomer(ctx context.Context, billingAccountName string, customerName string, parameters CustomerPolicy) (result CustomerPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PoliciesClient.UpdateCustomer")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateCustomerPreparer(ctx, billingAccountName, customerName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PoliciesClient", "UpdateCustomer", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateCustomerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PoliciesClient", "UpdateCustomer", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateCustomerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PoliciesClient", "UpdateCustomer", resp, "Failure responding to request")
	}

	return
}

// UpdateCustomerPreparer prepares the UpdateCustomer request.
func (client PoliciesClient) UpdateCustomerPreparer(ctx context.Context, billingAccountName string, customerName string, parameters CustomerPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"customerName":       autorest.Encode("path", customerName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/policies/default", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateCustomerSender sends the UpdateCustomer request. The method will close the
// http.Response Body if it receives an error.
func (client PoliciesClient) UpdateCustomerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateCustomerResponder handles the response to the UpdateCustomer request. The method always
// closes the http.Response Body.
func (client PoliciesClient) UpdateCustomerResponder(resp *http.Response) (result CustomerPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
