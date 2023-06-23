package pages

import (
	"net/http"

	"github.com/MetalbolicX/vanilla-go-webserver/pkg/render"
	"github.com/MetalbolicX/vanilla-go-webserver/pkg/types"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home-page.html", &types.TemplateData{})
}
