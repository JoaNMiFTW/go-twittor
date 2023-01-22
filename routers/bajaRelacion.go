package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JoaNMiFTW/go-twittor/bd"
	"github.com/JoaNMiFTW/go-twittor/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var id id
	err := decoder.Decode(&id)

	if err != nil {
		panic(err)
	}

	ID := id.ID

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se hja logrado borrar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
