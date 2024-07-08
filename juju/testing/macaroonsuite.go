// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package testing

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/go-macaroon-bakery/macaroon-bakery/v3/bakery/checkers"
	"github.com/go-macaroon-bakery/macaroon-bakery/v3/bakerytest"
	"github.com/go-macaroon-bakery/macaroon-bakery/v3/httpbakery"
	"github.com/juju/errors"
	"github.com/juju/names/v5"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
	"gopkg.in/macaroon.v2"

	"github.com/juju/juju/api"
	"github.com/juju/juju/controller"
	"github.com/juju/juju/core/permission"
	accesserrors "github.com/juju/juju/domain/access/errors"
	"github.com/juju/juju/domain/access/service"
)

// MacaroonSuite wraps a ApiServerSuite with macaroon authentication
// enabled.
type MacaroonSuite struct {
	ApiServerSuite

	// discharger holds the third-party discharger used
	// for authentication.
	discharger *bakerytest.Discharger

	// DischargerLogin is called by the discharger when an
	// API macaroon is discharged. It should either return
	// the chosen username or an empty string, in which case
	// the discharge is denied.
	// If this is nil, func() {return ""} is implied.
	DischargerLogin func() string
}

func (s *MacaroonSuite) SetUpTest(c *gc.C) {
	s.DischargerLogin = nil
	s.discharger = bakerytest.NewDischarger(nil)
	s.discharger.CheckerP = httpbakery.ThirdPartyCaveatCheckerPFunc(func(ctx context.Context, p httpbakery.ThirdPartyCaveatCheckerParams) ([]checkers.Caveat, error) {
		cond, _, err := checkers.ParseCaveat(string(p.Caveat.Condition))
		if err != nil {
			return nil, errors.Trace(err)
		}
		if cond != "is-authenticated-user" {
			return nil, errors.New("unknown caveat")
		}
		var username string
		if s.DischargerLogin != nil {
			username = s.DischargerLogin()
		}
		if username == "" {
			return nil, errors.New("login denied by discharger")
		}
		return []checkers.Caveat{checkers.DeclaredCaveat("username", username)}, nil
	})
	s.ApiServerSuite.ControllerConfigAttrs = map[string]interface{}{
		controller.IdentityURL: s.discharger.Location(),
	}
	s.ApiServerSuite.SetUpTest(c)
}

func (s *MacaroonSuite) TearDownTest(c *gc.C) {
	s.discharger.Close()
	s.ApiServerSuite.TearDownTest(c)
}

// DischargerLocation returns the URL of the third party caveat
// discharger.
func (s *MacaroonSuite) DischargerLocation() string {
	return s.discharger.Location()
}

// AddModelUser is a convenience function that adds an external
// user to the current model.
// It will panic if the user is local.
func (s *MacaroonSuite) AddModelUser(c *gc.C, username string) {
	if names.NewUserTag(username).IsLocal() {
		panic("cannot use MacaroonSuite.AddModelUser to add a local name")
	}

	accessService := s.ControllerServiceFactory(c).Access()
	_, _, err := accessService.AddUser(context.Background(), service.AddUserArg{
		Name:        username,
		DisplayName: "Remote User",
		External:    true,
		CreatorUUID: s.AdminUserUUID,
		Permission: permission.AccessSpec{
			Target: permission.ID{
				ObjectType: permission.Model,
				Key:        s.ControllerModelUUID(),
			},
			Access: permission.WriteAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
}

// AddControllerUser is a convenience function that adds
// a controller user with the specified access.
func (s *MacaroonSuite) AddControllerUser(c *gc.C, username string, access permission.Access) {
	accessService := s.ControllerServiceFactory(c).Access()
	perm := permission.AccessSpec{
		Access: access,
		Target: permission.ID{
			ObjectType: permission.Controller,
			Key:        s.ControllerUUID,
		},
	}

	_, _, err := accessService.AddUser(context.Background(), service.AddUserArg{
		Name:        username,
		DisplayName: "User Name",
		External:    true,
		CreatorUUID: s.AdminUserUUID,
		Permission:  perm,
	})

	// If we created this user prior, just add the controller permission.
	if errors.Is(err, accesserrors.UserAlreadyExists) {
		_, err = accessService.CreatePermission(context.Background(), permission.UserAccessSpec{
			AccessSpec: perm,
			User:       username,
		})
	}

	c.Assert(err, jc.ErrorIsNil)
}

// OpenAPI opens a connection to the API using the given information.
// and empty DialOpts. If info is nil, s.APIInfo(c) is used.
// If jar is non-nil, it will be used as the store for the cookies created
// as a result of API interaction.
func (s *MacaroonSuite) OpenAPI(c *gc.C, info *api.Info, jar http.CookieJar) api.Connection {
	if info == nil {
		info = s.APIInfo(c)
	}
	bakeryClient := httpbakery.NewClient()
	if jar != nil {
		bakeryClient.Client.Jar = jar
	}
	conn, err := api.Open(info, api.DialOpts{
		BakeryClient: bakeryClient,
	})
	c.Assert(err, gc.IsNil)
	return conn
}

// APIInfo returns API connection info suitable for
// connecting to the API using macaroon authentication.
func (s *MacaroonSuite) APIInfo(c *gc.C) *api.Info {
	info := s.ApiServerSuite.ControllerModelApiInfo()
	info.Tag = nil
	info.Password = ""
	// Fill in any old macaroon to ensure we don't attempt
	// an anonymous login.
	mac, err := NewMacaroon("test")
	c.Assert(err, jc.ErrorIsNil)
	info.Macaroons = []macaroon.Slice{{mac}}
	return info
}

// NewClearableCookieJar returns a new ClearableCookieJar.
func NewClearableCookieJar() *ClearableCookieJar {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	return &ClearableCookieJar{
		jar: jar,
	}
}

// ClearableCookieJar implements a cookie jar
// that can be cleared of all cookies for testing purposes.
type ClearableCookieJar struct {
	jar http.CookieJar
}

// Clear clears all the cookies in the jar.
// It is not OK to call Clear concurrently
// with the other methods.
func (jar *ClearableCookieJar) Clear() {
	newJar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	jar.jar = newJar
}

// Cookies implements http.CookieJar.Cookies.
func (jar *ClearableCookieJar) Cookies(u *url.URL) []*http.Cookie {
	return jar.jar.Cookies(u)
}

// SetCookies implements http.CookieJar.SetCookies.
func (jar *ClearableCookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.jar.SetCookies(u, cookies)
}

func MacaroonsEqual(c *gc.C, ms1, ms2 []macaroon.Slice) error {
	if len(ms1) != len(ms2) {
		return errors.Errorf("length mismatch, %d vs %d", len(ms1), len(ms2))
	}

	for i := 0; i < len(ms1); i++ {
		m1 := ms1[i]
		m2 := ms2[i]
		if len(m1) != len(m2) {
			return errors.Errorf("length mismatch, %d vs %d", len(m1), len(m2))
		}
		for i := 0; i < len(m1); i++ {
			MacaroonEquals(c, m1[i], m2[i])
		}
	}
	return nil
}

func MacaroonEquals(c *gc.C, m1, m2 *macaroon.Macaroon) {
	c.Assert(m1.Id(), jc.DeepEquals, m2.Id())
	c.Assert(m1.Signature(), jc.DeepEquals, m2.Signature())
	c.Assert(m1.Location(), gc.Equals, m2.Location())
}

func NewMacaroon(id string) (*macaroon.Macaroon, error) {
	return macaroon.New(nil, []byte(id), "", macaroon.LatestVersion)
}

func MustNewMacaroon(id string) *macaroon.Macaroon {
	mac, err := NewMacaroon(id)
	if err != nil {
		panic(err)
	}
	return mac
}
