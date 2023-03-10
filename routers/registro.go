package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JoaNMiFTW/go-twittor/bd"
	"github.com/JoaNMiFTW/go-twittor/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña no puede ser menor a 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, error := bd.InsertoRegistro(t)

	if error != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro de usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
