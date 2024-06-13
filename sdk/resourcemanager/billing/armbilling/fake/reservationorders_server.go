//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ReservationOrdersServer is a fake server for instances of the armbilling.ReservationOrdersClient type.
type ReservationOrdersServer struct {
	// GetByBillingAccount is the fake for method ReservationOrdersClient.GetByBillingAccount
	// HTTP status codes to indicate success: http.StatusOK
	GetByBillingAccount func(ctx context.Context, billingAccountName string, reservationOrderID string, options *armbilling.ReservationOrdersClientGetByBillingAccountOptions) (resp azfake.Responder[armbilling.ReservationOrdersClientGetByBillingAccountResponse], errResp azfake.ErrorResponder)

	// NewListByBillingAccountPager is the fake for method ReservationOrdersClient.NewListByBillingAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBillingAccountPager func(billingAccountName string, options *armbilling.ReservationOrdersClientListByBillingAccountOptions) (resp azfake.PagerResponder[armbilling.ReservationOrdersClientListByBillingAccountResponse])
}

// NewReservationOrdersServerTransport creates a new instance of ReservationOrdersServerTransport with the provided implementation.
// The returned ReservationOrdersServerTransport instance is connected to an instance of armbilling.ReservationOrdersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReservationOrdersServerTransport(srv *ReservationOrdersServer) *ReservationOrdersServerTransport {
	return &ReservationOrdersServerTransport{
		srv:                          srv,
		newListByBillingAccountPager: newTracker[azfake.PagerResponder[armbilling.ReservationOrdersClientListByBillingAccountResponse]](),
	}
}

// ReservationOrdersServerTransport connects instances of armbilling.ReservationOrdersClient to instances of ReservationOrdersServer.
// Don't use this type directly, use NewReservationOrdersServerTransport instead.
type ReservationOrdersServerTransport struct {
	srv                          *ReservationOrdersServer
	newListByBillingAccountPager *tracker[azfake.PagerResponder[armbilling.ReservationOrdersClientListByBillingAccountResponse]]
}

// Do implements the policy.Transporter interface for ReservationOrdersServerTransport.
func (r *ReservationOrdersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReservationOrdersClient.GetByBillingAccount":
		resp, err = r.dispatchGetByBillingAccount(req)
	case "ReservationOrdersClient.NewListByBillingAccountPager":
		resp, err = r.dispatchNewListByBillingAccountPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReservationOrdersServerTransport) dispatchGetByBillingAccount(req *http.Request) (*http.Response, error) {
	if r.srv.GetByBillingAccount == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByBillingAccount not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reservationOrders/(?P<reservationOrderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
	if err != nil {
		return nil, err
	}
	reservationOrderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reservationOrderId")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armbilling.ReservationOrdersClientGetByBillingAccountOptions
	if expandParam != nil {
		options = &armbilling.ReservationOrdersClientGetByBillingAccountOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := r.srv.GetByBillingAccount(req.Context(), billingAccountNameParam, reservationOrderIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ReservationOrder, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReservationOrdersServerTransport) dispatchNewListByBillingAccountPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByBillingAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBillingAccountPager not implemented")}
	}
	newListByBillingAccountPager := r.newListByBillingAccountPager.get(req)
	if newListByBillingAccountPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reservationOrders`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderByUnescaped, err := url.QueryUnescape(qp.Get("orderBy"))
		if err != nil {
			return nil, err
		}
		orderByParam := getOptional(orderByUnescaped)
		skiptokenUnescaped, err := url.QueryUnescape(qp.Get("skiptoken"))
		if err != nil {
			return nil, err
		}
		skiptokenParam, err := parseOptional(skiptokenUnescaped, func(v string) (float32, error) {
			p, parseErr := strconv.ParseFloat(v, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return float32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armbilling.ReservationOrdersClientListByBillingAccountOptions
		if filterParam != nil || orderByParam != nil || skiptokenParam != nil {
			options = &armbilling.ReservationOrdersClientListByBillingAccountOptions{
				Filter:    filterParam,
				OrderBy:   orderByParam,
				Skiptoken: skiptokenParam,
			}
		}
		resp := r.srv.NewListByBillingAccountPager(billingAccountNameParam, options)
		newListByBillingAccountPager = &resp
		r.newListByBillingAccountPager.add(req, newListByBillingAccountPager)
		server.PagerResponderInjectNextLinks(newListByBillingAccountPager, req, func(page *armbilling.ReservationOrdersClientListByBillingAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBillingAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByBillingAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBillingAccountPager) {
		r.newListByBillingAccountPager.remove(req)
	}
	return resp, nil
}
