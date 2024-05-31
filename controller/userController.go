package controller

import "github.com/gin-gonic/gin"


//Get Users Fuction
func GetUsers() gin.HandlerFunc {

	return func (c *gin.Context)  {
		
	}

}

//Get User By id
func GetUser() gin.HandlerFunc{

	return func (c *gin.Context)  {
		
	}
}


// Signup function
func Signup() gin.HandlerFunc{
	return func (c *gin.Context)  {
		
	}
}

//Login Function

func Login() gin.HandlerFunc{

	return func (c *gin.Context)  {
		
	}
}

//Hash Password function

func HashPassword(password string) string{

}

//Verify Password

func VerifyPassword(userPassword string, providePassword string) (bool, string){

}