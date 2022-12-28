package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JoaNMiFTW/go-twittor/bd"
	"github.com/JoaNMiFTW/go-twittor/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	fmt.Println(IDUsuario)

	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
