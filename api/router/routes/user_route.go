package routes

import (
	"net/http"

	"github.com/yadavsushil07/user-management/api/controllers"
)

var userRoutes = []Route{

	Route{
		Url:          "/user",
		Method:       http.MethodGet,
		Handler:      controllers.UserProfile,
		AuthRequired: true,
	},
	Route{
		Url:          "/user/{id}",
		Method:       http.MethodPut,
		Handler:      controllers.UpdateUser,
		AuthRequired: true,
	},

	Route{
		Url:          "/profilepic",
		Method:       http.MethodPut,
		Handler:      controllers.UploadImage,
		AuthRequired: true,
	},
	Route{
		Url:          "/changepassword",
		Method:       http.MethodPut,
		Handler:      controllers.ChangePassword,
		AuthRequired: true,
	},
}
