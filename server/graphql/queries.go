package graphql

import "github.com/hasura/go-graphql-client"

var GetUserByEmail struct {
	User []struct {
		ID         graphql.ID `graphql:"id" json:"id"`
		Password   string     `graphql:"password" json:"password"`
		ISVerifird bool       `graphql:"is_verified" json:"is_verified"`
		FirstName  string     `graphql:"first_name" json:"first_name"`
		LastName   string     `graphql:"last_name" json:"last_name"`
		Email      string     `graphql:"email" json:"email"`
		ProfilePic string     `graphql:"profile_picture" json:"profile_picture"`
	} `graphql:" users(where: {email: {_eq: $email}})"`
}

var INSERT_ONE_MUTATION struct {
	InsertUserOne struct {
		ID graphql.ID `graphql:"id" json:"id"`
	} `graphql:"insert_users_one(object: $userdata)"`
}

var VerifyEmailMutation struct {
	UpdateUserByID struct {
		ID graphql.ID `graphql:"id"`
	} `graphql:"update_users_by_pk(pk_columns: { id: $id }, _set: { is_verified: true })"`
}
