package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Buffalo is working!"}))
}
// func UnAuthorized(c buffalo.Context) error {
//   // Use the Error function on the context.
//   // This will result in a status code of 401.
//   return c.Error(401, r.JSON(map[string]string{"error" : "401 unauthorized." }))
// }

func GetPersonInfo(c buffalo.Context) error {
  var name = c.Param("name")
  return c.Render(200, r.JSON(map[string]string{"name" : name }))
}