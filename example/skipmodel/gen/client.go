// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package gen

import (
	"context"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/Yamashou/gqlgenc/example/skipmodel/model"
)

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli clientv2.HttpClient, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) *Client {
	return &Client{Client: clientv2.NewClient(cli, baseURL, options, interceptors...)}
}

type UserFragment struct {
	Profile *UserFragment_Profile "json:\"profile\" graphql:\"profile\""
}

func (t *UserFragment) GetProfile() *UserFragment_Profile {
	if t == nil {
		t = &UserFragment{}
	}
	return t.Profile
}

type ProfileFragment struct {
	Name string "json:\"name\" graphql:\"name\""
}

func (t *ProfileFragment) GetName() string {
	if t == nil {
		t = &ProfileFragment{}
	}
	return t.Name
}

type UserFragment_Profile struct {
	Name string "json:\"name\" graphql:\"name\""
}

func (t *UserFragment_Profile) GetName() string {
	if t == nil {
		t = &UserFragment_Profile{}
	}
	return t.Name
}

type A_User_Profile struct {
	ProfileFragment *ProfileFragment
	ID              string "json:\"id\" graphql:\"id\""
	Name            string "json:\"name\" graphql:\"name\""
}

func (t *A_User_Profile) GetProfileFragment() *ProfileFragment {
	if t == nil {
		t = &A_User_Profile{}
	}
	return t.ProfileFragment
}
func (t *A_User_Profile) GetID() string {
	if t == nil {
		t = &A_User_Profile{}
	}
	return t.ID
}
func (t *A_User_Profile) GetName() string {
	if t == nil {
		t = &A_User_Profile{}
	}
	return t.Name
}

type A_User struct {
	ID      string          "json:\"id\" graphql:\"id\""
	Profile *A_User_Profile "json:\"profile\" graphql:\"profile\""
}

func (t *A_User) GetID() string {
	if t == nil {
		t = &A_User{}
	}
	return t.ID
}
func (t *A_User) GetProfile() *A_User_Profile {
	if t == nil {
		t = &A_User{}
	}
	return t.Profile
}

type A struct {
	User A_User "json:\"user\" graphql:\"user\""
}

func (t *A) GetUser() *A_User {
	if t == nil {
		t = &A{}
	}
	return &t.User
}

const ADocument = `mutation A ($input: UserInput!) {
	user(input: $input) {
		id
		profile {
			id
			... ProfileFragment
		}
		... UserFragment
	}
}
fragment ProfileFragment on Profile {
	name
}
fragment UserFragment on User {
	profile {
		name
	}
}
`

func (c *Client) A(ctx context.Context, input model.UserInput, interceptors ...clientv2.RequestInterceptor) (*A, error) {
	vars := map[string]any{
		"input": input,
	}

	var res A
	if err := c.Client.Post(ctx, "A", ADocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

var DocumentOperationNames = map[string]string{
	ADocument: "A",
}
