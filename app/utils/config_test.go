package utils

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestIMConfig_Load(t *testing.T) {
	imConfig := NewIMConfig()

	err := imConfig.Load("")
	NotEqual(t, nil, err)

	err = imConfig.Load("api.yaml")
	Equal(t, nil, err)
}

func TestIMConfig_GetPort(t *testing.T) {
	imConfig := NewIMConfig()

	err := imConfig.Load("api.yaml")
	Equal(t, nil, err)

	port := imConfig.GetPort()
	Equal(t, ":8080", port)
}

func TestIMConfig_Mongo(t *testing.T) {
	imConfig := NewIMConfig()

	err := imConfig.Load("api.yaml")
	Equal(t, nil, err)

	mongoConfig := imConfig.Mongo()
	Equal(t, 1, len(mongoConfig))
}

func TestIMConfig_MongoItem(t *testing.T) {
	imConfig := NewIMConfig()

	err := imConfig.Load("api.yaml")
	Equal(t, nil, err)

	mongoConfig := imConfig.MongoItem("im")
	Equal(t, "im", mongoConfig.Database)
}

func TestIMConfig_MySQL(t *testing.T) {
	imConfig := NewIMConfig()

	err := imConfig.Load("api.yaml")
	Equal(t, nil, err)

	mysqlConfig := imConfig.MySQL()
	Equal(t, 1, len(mysqlConfig))
}

func TestIMConfig_MySQLItem(t *testing.T) {
	imConfig := NewIMConfig()

	err := imConfig.Load("api.yaml")
	Equal(t, nil, err)

	mysqlConfig := imConfig.MySQLItem("im")
	NotEqual(t, "", mysqlConfig.Database)

	mysqlConfig = imConfig.MySQLItem("im_not")
	Equal(t, "", mysqlConfig.Database)
}
