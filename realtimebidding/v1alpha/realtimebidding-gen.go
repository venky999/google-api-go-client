// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package realtimebidding provides access to the Real-time Bidding API.
//
// For product documentation, see: https://developers.google.com/authorized-buyers/apis/realtimebidding/reference/rest/
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/realtimebidding/v1alpha"
//   ...
//   ctx := context.Background()
//   realtimebiddingService, err := realtimebidding.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   realtimebiddingService, err := realtimebidding.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   realtimebiddingService, err := realtimebidding.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package realtimebidding // import "google.golang.org/api/realtimebidding/v1alpha"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "realtimebidding:v1alpha"
const apiName = "realtimebidding"
const apiVersion = "v1alpha"
const basePath = "https://realtimebidding.googleapis.com/"
const mtlsBasePath = "https://realtimebidding.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// See, create, edit, and delete your Authorized Buyers and Open Bidding
	// account entities
	RealtimeBiddingScope = "https://www.googleapis.com/auth/realtime-bidding"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/realtime-bidding",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Bidders = NewBiddersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Bidders *BiddersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewBiddersService(s *Service) *BiddersService {
	rs := &BiddersService{s: s}
	rs.BiddingFunctions = NewBiddersBiddingFunctionsService(s)
	return rs
}

type BiddersService struct {
	s *Service

	BiddingFunctions *BiddersBiddingFunctionsService
}

func NewBiddersBiddingFunctionsService(s *Service) *BiddersBiddingFunctionsService {
	rs := &BiddersBiddingFunctionsService{s: s}
	return rs
}

type BiddersBiddingFunctionsService struct {
	s *Service
}

// BiddingFunction: The bidding function to be executed as part of the
// TURTLEDOVE simulation experiment bidding flow.
type BiddingFunction struct {
	// BiddingFunction: The raw Javascript source code of the bidding
	// function. The function takes in a Javascript object, `inputs`, that
	// contains the following named fields: `openrtbContextualBidRequest` OR
	// `googleContextualBidRequest`, `customContextualSignal`,
	// `interestBasedBidData`, `interestGroupData`, `recentImpressionAges`,
	// and returns the bid price CPM (double). Example: ``` /* Returns a bid
	// price CPM (double). * * @param {Object} inputs an object with the *
	// following named fields: * - openrtbContextualBidRequest * OR
	// googleContextualBidRequest * - customContextualSignal * -
	// interestBasedBidData * - interestGroupData * - recentImpressionAges
	// */ function biddingFunction(inputs) { ... return
	// inputs.interestBasedBidData.cpm *
	// inputs.customContextualSignals.placementMultiplier; } ```
	BiddingFunction string `json:"biddingFunction,omitempty"`

	// Name: The name of the bidding function that must follow the pattern:
	// `bidders/{bidder_account_id}/biddingFunctions/{bidding_function_name}`
	// .
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "BiddingFunction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BiddingFunction") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *BiddingFunction) MarshalJSON() ([]byte, error) {
	type NoMethod BiddingFunction
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListBiddingFunctionsResponse: A response containing a list of a
// bidder's bidding functions.
type ListBiddingFunctionsResponse struct {
	// BiddingFunctions: A list of a bidder's bidding functions.
	BiddingFunctions []*BiddingFunction `json:"biddingFunctions,omitempty"`

	// NextPageToken: A token which can be passed to a subsequent call to
	// the `ListBiddingFunctions` method to retrieve the next page of
	// results in ListBiddingFunctionsRequest.pageToken.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "BiddingFunctions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BiddingFunctions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListBiddingFunctionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListBiddingFunctionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "realtimebidding.bidders.biddingFunctions.create":

type BiddersBiddingFunctionsCreateCall struct {
	s               *Service
	parent          string
	biddingfunction *BiddingFunction
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// Create: Creates a new bidding function.
func (r *BiddersBiddingFunctionsService) Create(parent string, biddingfunction *BiddingFunction) *BiddersBiddingFunctionsCreateCall {
	c := &BiddersBiddingFunctionsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.biddingfunction = biddingfunction
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersBiddingFunctionsCreateCall) Fields(s ...googleapi.Field) *BiddersBiddingFunctionsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersBiddingFunctionsCreateCall) Context(ctx context.Context) *BiddersBiddingFunctionsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersBiddingFunctionsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersBiddingFunctionsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210310")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.biddingfunction)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+parent}/biddingFunctions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "realtimebidding.bidders.biddingFunctions.create" call.
// Exactly one of *BiddingFunction or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *BiddingFunction.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersBiddingFunctionsCreateCall) Do(opts ...googleapi.CallOption) (*BiddingFunction, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BiddingFunction{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new bidding function.",
	//   "flatPath": "v1alpha/bidders/{biddersId}/biddingFunctions",
	//   "httpMethod": "POST",
	//   "id": "realtimebidding.bidders.biddingFunctions.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the bidder for which to create the bidding function. Format: `bidders/{bidderAccountId}`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+parent}/biddingFunctions",
	//   "request": {
	//     "$ref": "BiddingFunction"
	//   },
	//   "response": {
	//     "$ref": "BiddingFunction"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/realtime-bidding"
	//   ]
	// }

}

// method id "realtimebidding.bidders.biddingFunctions.list":

type BiddersBiddingFunctionsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the bidding functions that a bidder currently has
// registered.
func (r *BiddersBiddingFunctionsService) List(parent string) *BiddersBiddingFunctionsListCall {
	c := &BiddersBiddingFunctionsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of bidding functions to return.
func (c *BiddersBiddingFunctionsListCall) PageSize(pageSize int64) *BiddersBiddingFunctionsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return. This value is
// received from a previous `ListBiddingFunctions` call in
// ListBiddingFunctionsResponse.nextPageToken.
func (c *BiddersBiddingFunctionsListCall) PageToken(pageToken string) *BiddersBiddingFunctionsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersBiddingFunctionsListCall) Fields(s ...googleapi.Field) *BiddersBiddingFunctionsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersBiddingFunctionsListCall) IfNoneMatch(entityTag string) *BiddersBiddingFunctionsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersBiddingFunctionsListCall) Context(ctx context.Context) *BiddersBiddingFunctionsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersBiddingFunctionsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersBiddingFunctionsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210310")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+parent}/biddingFunctions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "realtimebidding.bidders.biddingFunctions.list" call.
// Exactly one of *ListBiddingFunctionsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListBiddingFunctionsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersBiddingFunctionsListCall) Do(opts ...googleapi.CallOption) (*ListBiddingFunctionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListBiddingFunctionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the bidding functions that a bidder currently has registered.",
	//   "flatPath": "v1alpha/bidders/{biddersId}/biddingFunctions",
	//   "httpMethod": "GET",
	//   "id": "realtimebidding.bidders.biddingFunctions.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of bidding functions to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return. This value is received from a previous `ListBiddingFunctions` call in ListBiddingFunctionsResponse.nextPageToken.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. Name of the bidder whose bidding functions will be listed. Format: `bidders/{bidder_account_id}`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+parent}/biddingFunctions",
	//   "response": {
	//     "$ref": "ListBiddingFunctionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/realtime-bidding"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersBiddingFunctionsListCall) Pages(ctx context.Context, f func(*ListBiddingFunctionsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
