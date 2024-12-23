// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package generated

import (
	"context"

	"github.com/Yamashou/gqlgenc/clientv2"
)

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli clientv2.HttpClient, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) *Client {
	return &Client{Client: clientv2.NewClient(cli, baseURL, options, interceptors...)}
}

type UserFragment struct {
	ID      string               "json:\"id\" graphql:\"id\""
	Profile UserFragment_Profile "json:\"profile\" graphql:\"profile\""
}

func (t *UserFragment) GetID() string {
	if t == nil {
		t = &UserFragment{}
	}
	return t.ID
}
func (t *UserFragment) GetProfile() *UserFragment_Profile {
	if t == nil {
		t = &UserFragment{}
	}
	return &t.Profile
}

type UserFragment_Profile_Detail struct {
	BirthDate string "json:\"birthDate\" graphql:\"birthDate\""
}

func (t *UserFragment_Profile_Detail) GetBirthDate() string {
	if t == nil {
		t = &UserFragment_Profile_Detail{}
	}
	return t.BirthDate
}

type UserFragment_Profile struct {
	Detail UserFragment_Profile_Detail "json:\"detail\" graphql:\"detail\""
	ID     string                      "json:\"id\" graphql:\"id\""
}

func (t *UserFragment_Profile) GetDetail() *UserFragment_Profile_Detail {
	if t == nil {
		t = &UserFragment_Profile{}
	}
	return &t.Detail
}
func (t *UserFragment_Profile) GetID() string {
	if t == nil {
		t = &UserFragment_Profile{}
	}
	return t.ID
}

type UserDetail_User_Profile_Detail struct {
	BirthDate string "json:\"birthDate\" graphql:\"birthDate\""
	ID        string "json:\"id\" graphql:\"id\""
}

func (t *UserDetail_User_Profile_Detail) GetBirthDate() string {
	if t == nil {
		t = &UserDetail_User_Profile_Detail{}
	}
	return t.BirthDate
}
func (t *UserDetail_User_Profile_Detail) GetID() string {
	if t == nil {
		t = &UserDetail_User_Profile_Detail{}
	}
	return t.ID
}

type UserDetail_User_Profile struct {
	Company string                         "json:\"company\" graphql:\"company\""
	Detail  UserDetail_User_Profile_Detail "json:\"detail\" graphql:\"detail\""
	ID      string                         "json:\"id\" graphql:\"id\""
	Name    string                         "json:\"name\" graphql:\"name\""
}

func (t *UserDetail_User_Profile) GetCompany() string {
	if t == nil {
		t = &UserDetail_User_Profile{}
	}
	return t.Company
}
func (t *UserDetail_User_Profile) GetDetail() *UserDetail_User_Profile_Detail {
	if t == nil {
		t = &UserDetail_User_Profile{}
	}
	return &t.Detail
}
func (t *UserDetail_User_Profile) GetID() string {
	if t == nil {
		t = &UserDetail_User_Profile{}
	}
	return t.ID
}
func (t *UserDetail_User_Profile) GetName() string {
	if t == nil {
		t = &UserDetail_User_Profile{}
	}
	return t.Name
}

type UserDetail_User struct {
	ID      string                  "json:\"id\" graphql:\"id\""
	Profile UserDetail_User_Profile "json:\"profile\" graphql:\"profile\""
}

func (t *UserDetail_User) GetID() string {
	if t == nil {
		t = &UserDetail_User{}
	}
	return t.ID
}
func (t *UserDetail_User) GetProfile() *UserDetail_User_Profile {
	if t == nil {
		t = &UserDetail_User{}
	}
	return &t.Profile
}

type UserDetail struct {
	User UserDetail_User "json:\"user\" graphql:\"user\""
}

func (t *UserDetail) GetUser() *UserDetail_User {
	if t == nil {
		t = &UserDetail{}
	}
	return &t.User
}

const UserDetailDocument = `query UserDetail {
	user {
		... UserFragment
		id
		profile {
			name
			company
			detail {
				id
			}
		}
	}
}
fragment UserFragment on User {
	id
	profile {
		id
		detail {
			birthDate
		}
	}
}
`

func (c *Client) UserDetail(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*UserDetail, error) {
	vars := map[string]any{}

	var res UserDetail
	if err := c.Client.Post(ctx, "UserDetail", UserDetailDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

var DocumentOperationNames = map[string]string{
	UserDetailDocument: "UserDetail",
}
