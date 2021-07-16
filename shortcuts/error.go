package shortcuts

import (
	"github.com/Nictec/gwt/handler"
	"github.com/flosch/pongo2/v4"
)

func Error(ctx handler.Context, code int, message string) {
	Render(ctx, "templates/error.html", pongo2.Context{"code": code, "message": message})
}
