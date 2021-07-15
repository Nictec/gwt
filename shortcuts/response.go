package shortcuts

import (
	"encoding/json"
	"github.com/Nictec/gwt/handler"
	"github.com/Nictec/gwt/logger"
	"github.com/flosch/pongo2/v4"
)

func Render(ctx handler.Context, file string, pongoCtx pongo2.Context){
	tpl := pongo2.Must(pongo2.FromFile(file))
	err := tpl.ExecuteWriter(pongoCtx, ctx.W)
	if err != nil{
		logger.Error(err.Error())
	}
}

func JsonResponse(ctx handler.Context, data interface{}) {
	ctx.W.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.W).Encode(data)
}
