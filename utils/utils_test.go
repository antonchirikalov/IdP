package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestViperConfig(t *testing.T) {
	Convey("Given yaml config ", t, func() {
		Convey("When config is read", func() {
			v := GetConfig()
			Convey("Values must be extracted", func() {
				So(v.GetString("appname"), ShouldEqual, "IdP")
				So(v.GetString("httpaddr"), ShouldEqual, "127.0.0.1")
			})
		})
	})
}

func TestOSEnvVariables(t *testing.T) {
	Convey("Given OS variable", t, func() {
		os.Setenv("IDP_HYDRAADMIN", "test")
		os.Setenv("IDP_POSTGRES_PORT", "666")
		Convey("When config is read", func() {
			Convey("Values must be extracted", func() {
				So(GetConfig().GetString("hydraAdmin"), ShouldEqual, "test")
				So(GetConfig().GetString("postgres.port"), ShouldEqual, "666")
				os.Clearenv()
				So(GetConfig().GetString("hydraAdmin"), ShouldNotEqual, "test")
				So(GetConfig().GetString("postgres.port"), ShouldNotEqual, "666")
			})
		})
	})
}
