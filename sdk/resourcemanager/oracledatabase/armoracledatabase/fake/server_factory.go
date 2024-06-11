//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armoracledatabase.ClientFactory type.
type ServerFactory struct {
	AutonomousDatabaseBackupsServer               AutonomousDatabaseBackupsServer
	AutonomousDatabaseCharacterSetsServer         AutonomousDatabaseCharacterSetsServer
	AutonomousDatabaseNationalCharacterSetsServer AutonomousDatabaseNationalCharacterSetsServer
	AutonomousDatabaseVersionsServer              AutonomousDatabaseVersionsServer
	AutonomousDatabasesServer                     AutonomousDatabasesServer
	CloudExadataInfrastructuresServer             CloudExadataInfrastructuresServer
	CloudVMClustersServer                         CloudVMClustersServer
	DNSPrivateViewsServer                         DNSPrivateViewsServer
	DNSPrivateZonesServer                         DNSPrivateZonesServer
	DbNodesServer                                 DbNodesServer
	DbServersServer                               DbServersServer
	DbSystemShapesServer                          DbSystemShapesServer
	GiVersionsServer                              GiVersionsServer
	OperationsServer                              OperationsServer
	OracleSubscriptionsServer                     OracleSubscriptionsServer
	SystemVersionsServer                          SystemVersionsServer
	VirtualNetworkAddressesServer                 VirtualNetworkAddressesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armoracledatabase.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armoracledatabase.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                             *ServerFactory
	trMu                                            sync.Mutex
	trAutonomousDatabaseBackupsServer               *AutonomousDatabaseBackupsServerTransport
	trAutonomousDatabaseCharacterSetsServer         *AutonomousDatabaseCharacterSetsServerTransport
	trAutonomousDatabaseNationalCharacterSetsServer *AutonomousDatabaseNationalCharacterSetsServerTransport
	trAutonomousDatabaseVersionsServer              *AutonomousDatabaseVersionsServerTransport
	trAutonomousDatabasesServer                     *AutonomousDatabasesServerTransport
	trCloudExadataInfrastructuresServer             *CloudExadataInfrastructuresServerTransport
	trCloudVMClustersServer                         *CloudVMClustersServerTransport
	trDNSPrivateViewsServer                         *DNSPrivateViewsServerTransport
	trDNSPrivateZonesServer                         *DNSPrivateZonesServerTransport
	trDbNodesServer                                 *DbNodesServerTransport
	trDbServersServer                               *DbServersServerTransport
	trDbSystemShapesServer                          *DbSystemShapesServerTransport
	trGiVersionsServer                              *GiVersionsServerTransport
	trOperationsServer                              *OperationsServerTransport
	trOracleSubscriptionsServer                     *OracleSubscriptionsServerTransport
	trSystemVersionsServer                          *SystemVersionsServerTransport
	trVirtualNetworkAddressesServer                 *VirtualNetworkAddressesServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "AutonomousDatabaseBackupsClient":
		initServer(s, &s.trAutonomousDatabaseBackupsServer, func() *AutonomousDatabaseBackupsServerTransport {
			return NewAutonomousDatabaseBackupsServerTransport(&s.srv.AutonomousDatabaseBackupsServer)
		})
		resp, err = s.trAutonomousDatabaseBackupsServer.Do(req)
	case "AutonomousDatabaseCharacterSetsClient":
		initServer(s, &s.trAutonomousDatabaseCharacterSetsServer, func() *AutonomousDatabaseCharacterSetsServerTransport {
			return NewAutonomousDatabaseCharacterSetsServerTransport(&s.srv.AutonomousDatabaseCharacterSetsServer)
		})
		resp, err = s.trAutonomousDatabaseCharacterSetsServer.Do(req)
	case "AutonomousDatabaseNationalCharacterSetsClient":
		initServer(s, &s.trAutonomousDatabaseNationalCharacterSetsServer, func() *AutonomousDatabaseNationalCharacterSetsServerTransport {
			return NewAutonomousDatabaseNationalCharacterSetsServerTransport(&s.srv.AutonomousDatabaseNationalCharacterSetsServer)
		})
		resp, err = s.trAutonomousDatabaseNationalCharacterSetsServer.Do(req)
	case "AutonomousDatabaseVersionsClient":
		initServer(s, &s.trAutonomousDatabaseVersionsServer, func() *AutonomousDatabaseVersionsServerTransport {
			return NewAutonomousDatabaseVersionsServerTransport(&s.srv.AutonomousDatabaseVersionsServer)
		})
		resp, err = s.trAutonomousDatabaseVersionsServer.Do(req)
	case "AutonomousDatabasesClient":
		initServer(s, &s.trAutonomousDatabasesServer, func() *AutonomousDatabasesServerTransport {
			return NewAutonomousDatabasesServerTransport(&s.srv.AutonomousDatabasesServer)
		})
		resp, err = s.trAutonomousDatabasesServer.Do(req)
	case "CloudExadataInfrastructuresClient":
		initServer(s, &s.trCloudExadataInfrastructuresServer, func() *CloudExadataInfrastructuresServerTransport {
			return NewCloudExadataInfrastructuresServerTransport(&s.srv.CloudExadataInfrastructuresServer)
		})
		resp, err = s.trCloudExadataInfrastructuresServer.Do(req)
	case "CloudVMClustersClient":
		initServer(s, &s.trCloudVMClustersServer, func() *CloudVMClustersServerTransport {
			return NewCloudVMClustersServerTransport(&s.srv.CloudVMClustersServer)
		})
		resp, err = s.trCloudVMClustersServer.Do(req)
	case "DNSPrivateViewsClient":
		initServer(s, &s.trDNSPrivateViewsServer, func() *DNSPrivateViewsServerTransport {
			return NewDNSPrivateViewsServerTransport(&s.srv.DNSPrivateViewsServer)
		})
		resp, err = s.trDNSPrivateViewsServer.Do(req)
	case "DNSPrivateZonesClient":
		initServer(s, &s.trDNSPrivateZonesServer, func() *DNSPrivateZonesServerTransport {
			return NewDNSPrivateZonesServerTransport(&s.srv.DNSPrivateZonesServer)
		})
		resp, err = s.trDNSPrivateZonesServer.Do(req)
	case "DbNodesClient":
		initServer(s, &s.trDbNodesServer, func() *DbNodesServerTransport { return NewDbNodesServerTransport(&s.srv.DbNodesServer) })
		resp, err = s.trDbNodesServer.Do(req)
	case "DbServersClient":
		initServer(s, &s.trDbServersServer, func() *DbServersServerTransport { return NewDbServersServerTransport(&s.srv.DbServersServer) })
		resp, err = s.trDbServersServer.Do(req)
	case "DbSystemShapesClient":
		initServer(s, &s.trDbSystemShapesServer, func() *DbSystemShapesServerTransport {
			return NewDbSystemShapesServerTransport(&s.srv.DbSystemShapesServer)
		})
		resp, err = s.trDbSystemShapesServer.Do(req)
	case "GiVersionsClient":
		initServer(s, &s.trGiVersionsServer, func() *GiVersionsServerTransport { return NewGiVersionsServerTransport(&s.srv.GiVersionsServer) })
		resp, err = s.trGiVersionsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "OracleSubscriptionsClient":
		initServer(s, &s.trOracleSubscriptionsServer, func() *OracleSubscriptionsServerTransport {
			return NewOracleSubscriptionsServerTransport(&s.srv.OracleSubscriptionsServer)
		})
		resp, err = s.trOracleSubscriptionsServer.Do(req)
	case "SystemVersionsClient":
		initServer(s, &s.trSystemVersionsServer, func() *SystemVersionsServerTransport {
			return NewSystemVersionsServerTransport(&s.srv.SystemVersionsServer)
		})
		resp, err = s.trSystemVersionsServer.Do(req)
	case "VirtualNetworkAddressesClient":
		initServer(s, &s.trVirtualNetworkAddressesServer, func() *VirtualNetworkAddressesServerTransport {
			return NewVirtualNetworkAddressesServerTransport(&s.srv.VirtualNetworkAddressesServer)
		})
		resp, err = s.trVirtualNetworkAddressesServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
