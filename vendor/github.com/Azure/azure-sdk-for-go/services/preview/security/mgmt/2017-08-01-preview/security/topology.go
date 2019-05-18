package security

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// TopologyClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type TopologyClient struct {
	BaseClient
}

// NewTopologyClient creates an instance of the TopologyClient client.
func NewTopologyClient(subscriptionID string, ascLocation string) TopologyClient {
	return NewTopologyClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewTopologyClientWithBaseURI creates an instance of the TopologyClient client.
func NewTopologyClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) TopologyClient {
	return TopologyClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Get gets a specific topology component.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// topologyResourceName - name of a topology resources collection.
func (client TopologyClient) Get(ctx context.Context, resourceGroupName string, topologyResourceName string) (result TopologyResource, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.TopologyClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, topologyResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TopologyClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.TopologyClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TopologyClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client TopologyClient) GetPreparer(ctx context.Context, resourceGroupName string, topologyResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":          autorest.Encode("path", client.AscLocation),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
		"topologyResourceName": autorest.Encode("path", topologyResourceName),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/topologies/{topologyResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TopologyClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TopologyClient) GetResponder(resp *http.Response) (result TopologyResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list that allows to build a topology view of a subscription.
func (client TopologyClient) List(ctx context.Context) (result TopologyListPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.TopologyClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TopologyClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.tl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.TopologyClient", "List", resp, "Failure sending request")
		return
	}

	result.tl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TopologyClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client TopologyClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/topologies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TopologyClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TopologyClient) ListResponder(resp *http.Response) (result TopologyList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client TopologyClient) listNextResults(lastResults TopologyList) (result TopologyList, err error) {
	req, err := lastResults.topologyListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.TopologyClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.TopologyClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TopologyClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client TopologyClient) ListComplete(ctx context.Context) (result TopologyListIterator, err error) {
	result.page, err = client.List(ctx)
	return
}

// ListByHomeRegion gets a list that allows to build a topology view of a subscription and location.
func (client TopologyClient) ListByHomeRegion(ctx context.Context) (result TopologyListPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.TopologyClient", "ListByHomeRegion", err.Error())
	}

	result.fn = client.listByHomeRegionNextResults
	req, err := client.ListByHomeRegionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TopologyClient", "ListByHomeRegion", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHomeRegionSender(req)
	if err != nil {
		result.tl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.TopologyClient", "ListByHomeRegion", resp, "Failure sending request")
		return
	}

	result.tl, err = client.ListByHomeRegionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TopologyClient", "ListByHomeRegion", resp, "Failure responding to request")
	}

	return
}

// ListByHomeRegionPreparer prepares the ListByHomeRegion request.
func (client TopologyClient) ListByHomeRegionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":    autorest.Encode("path", client.AscLocation),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/topologies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHomeRegionSender sends the ListByHomeRegion request. The method will close the
// http.Response Body if it receives an error.
func (client TopologyClient) ListByHomeRegionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByHomeRegionResponder handles the response to the ListByHomeRegion request. The method always
// closes the http.Response Body.
func (client TopologyClient) ListByHomeRegionResponder(resp *http.Response) (result TopologyList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHomeRegionNextResults retrieves the next set of results, if any.
func (client TopologyClient) listByHomeRegionNextResults(lastResults TopologyList) (result TopologyList, err error) {
	req, err := lastResults.topologyListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.TopologyClient", "listByHomeRegionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHomeRegionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.TopologyClient", "listByHomeRegionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHomeRegionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TopologyClient", "listByHomeRegionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHomeRegionComplete enumerates all values, automatically crossing page boundaries as required.
func (client TopologyClient) ListByHomeRegionComplete(ctx context.Context) (result TopologyListIterator, err error) {
	result.page, err = client.ListByHomeRegion(ctx)
	return
}
