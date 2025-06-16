package handlers

import (
	"net/http"

	"github.com/zemld/Shop/admin-service/db"
	"github.com/zemld/Shop/admin-service/domain/models"
	"github.com/zemld/Shop/admin-service/internal"
)

// @description Checks if user is admin.
// @tag.name Admin operations
// @param code query string true "Secret auth code"
// @produce json
// @success 200 {object} models.IsAdminResponse
// @failure 400 {object} models.StatusResponse
// @failure 403 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/admins/auth [get]
func AuthentificateAdminHandler(w http.ResponseWriter, r *http.Request) {
	secretCode := r.URL.Query().Get("code")

	adminName, err := db.CreateDBConnectionAndCheckAdmin(db.AdminDB, secretCode)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    adminName,
			Message: "Failed to authenticate admin: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.WriteResponse(w, models.IsAdminResponse{
		Name:    adminName,
		IsAdmin: adminName != "",
	}, http.StatusOK)
}
