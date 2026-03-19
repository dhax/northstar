package resources

import (
	"net/http"

	"northstar/config"

	"github.com/templui/templui/utils"
)

const (
	LibsDirectoryPath   = "web/libs"
	StylesDirectoryPath = "web/resources/styles"
	StaticDirectoryPath = "web/resources/static"
)

func TemplUIScriptRoutes() http.Handler {
	mux := http.NewServeMux()

	utils.SetupScriptRoutes(mux, config.Global.Environment == config.Dev)

	return mux
}
