package validator

import (
	"encoding/json"
	"mux-training/dao"
	"mux-training/model"
	"mux-training/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"github.com/asaskevich/govalidator"
)

// Todo create
func TodoPayload() util.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			// Bind and parse to struct
			var payload model.TodoPayload
			err := json.NewDecoder(r.Body).Decode(&payload)

			if err != nil {
				util.Response400(w, err.Error())
				return
			}

			// Validate struct
			_, err = govalidator.ValidateStruct(payload)

			if err != nil {
				util.Response400(w, err.Error())
				return
			}

			// Set body and move to next process
			context.Set(r, "payload", payload)
			next.ServeHTTP(w, r)

		}
	}
}

// Todo validate
func TodoCheckExistance() util.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			// Get ID
			var idStr = mux.Vars(r)["id"]

			// Check valid driver ID
			id, err := strconv.Atoi(strings.Trim(idStr, " "))
			if err != nil {
				util.Response400(w, err.Error())
				return
			}

			// Validate todo
			_, err = dao.TodoFindByID(id)
			if err != nil {
				util.Response404(w, err.Error())
				return
			}

			// Set body and move to next process
			context.Set(r, "id", id)
			next.ServeHTTP(w, r)
		}
	}
}
