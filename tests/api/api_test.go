package tests

import (
	"imgo/app/resources"
	"imgo/app/routes"
	"net/http/httptest"

	"github.com/cucumber/godog"
)

// InitializeScenario init scenario
func InitializeScenario(ctx *godog.ScenarioContext) {
	apiFeature := APIFeature{}
	resource := resources.NewIMResource()
	_, err := resource.Config([]string{}, "api.yaml")
	if err != nil {
		panic(err)
	}

	if APISQL == nil {
		mySQL, err := resource.MySQLConn()
		if err != nil {
			panic(err)
		}
		APISQL = mySQL
	}

	if APIRoute == nil {
		APIRoute = routes.SetupRouter(APISQL)
	}

	if APIServer == nil {
		APIServer = httptest.NewServer(APIRoute)
	}

	ctx.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`, apiFeature.iSendrequestTo)
	ctx.Step(`^the response status code should be (\d+)$`, apiFeature.theResponseStatusCodeShouldBe)
	ctx.Step(`^the response should match json:$`, apiFeature.theResponseShouldMatchJSON)
}
