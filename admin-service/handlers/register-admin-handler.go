package handlers

import (
	"log"
	"net/http"

	"github.com/zemld/Shop/admin-service/db"
	"github.com/zemld/Shop/admin-service/domain/models"
	"github.com/zemld/Shop/admin-service/internal"
)

// @description Registers a new admin.
// @tags Admins
// @param name query string true "Admin which you want to register"
// @produce json
// @success 200 {object} models.Admin
// @failure 400 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/admins/register [post]
func RegisterAdminHandler(w http.ResponseWriter, r *http.Request) {
	adminName := r.URL.Query().Get("name")
	if adminName == "" {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    adminName,
			Message: "Admin name is required",
		}, http.StatusBadRequest)
		return
	}
	secretCode := internal.GetSecretCode(adminName)
	err := db.CreateDBConnectionAndRegisterAdmin(db.AdminDB, adminName, secretCode)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    adminName,
			Message: "Failed to register admin: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.WriteResponse(w, models.Admin{
		Name: adminName,
		Code: secretCode,
	}, http.StatusOK)
	log.Printf("Admin %s registered successfully with code %s", adminName, secretCode)
}
