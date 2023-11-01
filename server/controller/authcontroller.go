package controller

import (
	"context"
	"log"
	"net/http"
	"server/graphql"
	"server/types"
	jwt_modules "server/utils/jwt"
	"server/utils/smpt"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var signupRequest = types.SignupRequest{}

	if err := ctx.ShouldBindJSON(&signupRequest); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request",
			"success": false,
		})
		return
	}

	// log.Println(signupRequest.Input.Arg)

	// Verifay the email not exists
	client := graphql.Client()
	variables := map[string]interface{}{
		"email": signupRequest.Input.Arg.Email,
	}
	if err := client.Query(context.Background(), &graphql.GetUserByEmail, variables); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": err,
		})
		return
	}
	log.Println("sdc", &graphql.GetUserByEmail)
	if len(graphql.GetUserByEmail.User) > 0 {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "Email already exists",
			"success": false,
		})
		return
	}

	type users_insert_input struct {
		Email     string `graphql:"email" json:"email"`
		FirstName string `graphql:"first_name" json:"first_name"`
		LastName  string `graphql:"last_name" json:"last_name"`
		Password  string `graphql:"password" json:"password"`
	}
	variables = map[string]interface{}{

		"userdata": users_insert_input{
			Email:     signupRequest.Input.Arg.Email,
			FirstName: signupRequest.Input.Arg.FirstName,
			LastName:  signupRequest.Input.Arg.LastName,
			Password:  signupRequest.Input.Arg.Password,
		},
	}
	// log.Println(&graphql.INSERT_ONE_MUTATION.InsertUserOne.Returing.ID)
	if err := client.Mutate(context.Background(), &graphql.INSERT_ONE_MUTATION, variables); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"success": false,
		})

		return
	}

	token, err := jwt_modules.GenerateToken(string(graphql.INSERT_ONE_MUTATION.InsertUserOne.ID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"success": false,
		})
		return
	}

	smpt.SendEmail(signupRequest.Input.Arg.Email, "Verifay Your Email", token)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Verifay Your Email",
		"success": true,
	})

}

func Signin(ctx *gin.Context) {
	var signinRequest = types.SigninRequest{}
	if err := ctx.ShouldBindJSON(&signinRequest); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	client := graphql.Client()
	// Verifay the email not exists
	variables := map[string]interface{}{
		"email": signinRequest.Email,
	}

	if err := client.Query(context.Background(), &graphql.GetUserByEmail, variables); err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if len(graphql.GetUserByEmail.User) == 0 {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "Wrong Credentials",
		})
		return
	}

	if !(graphql.GetUserByEmail.User[0].ISVerifird) {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "Verifay your email",
		})

		return
	}

	if graphql.GetUserByEmail.User[0].Password != signinRequest.Password {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "Wrong Credentials",
		})
		return
	}

	token, err := jwt_modules.GenerateToken(string(graphql.INSERT_ONE_MUTATION.InsertUserOne.ID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully",
		"data": gin.H{
			"id":              graphql.GetUserByEmail.User[0].ID,
			"first_name":      graphql.GetUserByEmail.User[0].FirstName,
			"last_name":       graphql.GetUserByEmail.User[0].LastName,
			"email":           graphql.GetUserByEmail.User[0].Email,
			"profile_picture": graphql.GetUserByEmail.User[0].ProfilePic,
			"accessToken":     token,
		},
	})
}

func VerifyEmail(ctx *gin.Context) {
	var verifyEmailRequest = types.VerifyEmailRequest{}

	if err := ctx.ShouldBind(&verifyEmailRequest); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	client := graphql.Client()

	// extract id
	claims, err := jwt_modules.VerifyToken(verifyEmailRequest.Input.Arg.Token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	log.Println("claims", claims.Payload.UserID)

	type uuid string

	u := uuid(claims.Payload.UserID)

	log.Println("uuid", u)

	variables := map[string]interface{}{
		"id": u,
	}

	log.Println(variables)

	if err := client.Mutate(context.Background(), &graphql.VerifyEmailMutation, variables); err != nil {
		log.Println("graphql.Verifay_Email_Mutation", err)
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": err,
		})
		return
	}

	log.Println("graphql.Verifay_Email_Mutation", graphql.VerifyEmailMutation.UpdateUserByID.ID)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Email verified successfully",
	})
}
