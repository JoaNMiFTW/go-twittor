package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JoaNMiFTW/go-twittor/bd"
	"github.com/JoaNMiFTW/go-twittor/models"
)

type id struct {
	ID string `json:"id"`
}

func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var id id
	err := decoder.Decode(&id)

	if err != nil {
		panic(err)
	}

	ID := id.ID
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se hja logrado insertar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
