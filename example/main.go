package main

import (
	"github.com/qovery/qovery-client-go"
	"golang.org/x/net/context"
)

func main() {
	auth := context.WithValue(context.Background(), qovery.ContextAccessToken, "ACCESS_TOKEN")
	client := qovery.NewAPIClient(qovery.NewConfiguration())
	app, res, err := client.ApplicationsApi.CreateApplication(auth, "ENV_ID").ApplicationRequest(qovery.ApplicationRequest{}).Execute()

	if err != nil {
		panic(err)
	}
	if res.StatusCode >= 400 {
		panic("Error Response")
	}

	println(app.Id)
	println(app.Environment)
	println(app.CreatedAt.String())
}
