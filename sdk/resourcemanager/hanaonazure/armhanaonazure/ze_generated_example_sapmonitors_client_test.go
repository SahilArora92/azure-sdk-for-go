//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhanaonazure_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_List.json
func ExampleSapMonitorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhanaonazure.NewSapMonitorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_Get.json
func ExampleSapMonitorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhanaonazure.NewSapMonitorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<sap-monitor-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_Create.json
func ExampleSapMonitorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhanaonazure.NewSapMonitorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<sap-monitor-name>",
		armhanaonazure.SapMonitor{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
			Properties: &armhanaonazure.SapMonitorProperties{
				EnableCustomerAnalytics:        to.Ptr(true),
				LogAnalyticsWorkspaceArmID:     to.Ptr("<log-analytics-workspace-arm-id>"),
				LogAnalyticsWorkspaceID:        to.Ptr("<log-analytics-workspace-id>"),
				LogAnalyticsWorkspaceSharedKey: to.Ptr("<log-analytics-workspace-shared-key>"),
				MonitorSubnet:                  to.Ptr("<monitor-subnet>"),
			},
		},
		&armhanaonazure.SapMonitorsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_Delete.json
func ExampleSapMonitorsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhanaonazure.NewSapMonitorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<sap-monitor-name>",
		&armhanaonazure.SapMonitorsClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_PatchTags_Delete.json
func ExampleSapMonitorsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhanaonazure.NewSapMonitorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<sap-monitor-name>",
		armhanaonazure.Tags{
			Tags: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
