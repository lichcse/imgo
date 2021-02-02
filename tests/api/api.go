package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"

	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// APIFeature app context
type APIFeature struct {
	resp *http.Response
}

// APIRoute route
var APIRoute *gin.Engine

// APISQL RDB
var APISQL *gorm.DB

// APIServer server
var APIServer *httptest.Server

func (_this *APIFeature) theResponseStatusCodeShouldBe(code int) error {
	if code != _this.resp.StatusCode {
		return fmt.Errorf("expected response status code to be: %d, but actual is: %d", code, _this.resp.StatusCode)
	}
	return nil
}

func (_this *APIFeature) theResponseShouldMatchJSON(body *godog.DocString) (err error) {
	var expected, actual interface{}
	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return
	}

	bodyRes, err := ioutil.ReadAll(_this.resp.Body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(bodyRes, &actual); err != nil {
		return
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected json response does not match actual, %v vs. %v", expected, actual)
	}
	return nil
}

func (_this *APIFeature) iSendrequestTo(method string, endpoint string) (err error) {
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	switch method {
	case "GET":
		_this.resp, err = http.Get(fmt.Sprintf("%s%s", APIServer.URL, endpoint))
	default:
		err = fmt.Errorf("unknown method: %s", method)
	}
	return
}
