package handlers

import (
	"log"
	"net/http"

	"github.com/zemld/Shop/admin-service/db"
	"github.com/zemld/Shop/admin-service/domain/dto"
	"github.com/zemld/Shop/admin-service/domain/models"
	"github.com/zemld/Shop/admin-service/internal"
)

// @description Registers a new admin.
// @tag.name Users operations
// @param name query string true "Admin which you want to register"
// @produce json
// @success 200 {object} models.Admin
// @failure 400 {object} dto.StatusResponse
// @failure 500 {object} dto.StatusResponse
// @router /v1/admins/register [post]
func RegisterAdminHandler(w http.ResponseWriter, r *http.Request) {
	// получаю секретный код администратора и отправляю запрос на сохранение в бд.
	// возвращаю ответ с кодом 200, если всё успешно, или 500, если произошла ошибка.
	adminName := r.URL.Query().Get("name")
	if adminName == "" {
		internal.WriteResponse(w, dto.StatusResponse{
			User:    adminName,
			Message: "Admin name is required",
		}, http.StatusBadRequest)
		return
	}
	secretCode := internal.GetSecretCode(adminName)
	err := db.CreateDBConnectionAndRegisterAdmin(db.AdminDB, adminName, secretCode)
	if err != nil {
		internal.WriteResponse(w, dto.StatusResponse{
			User:    adminName,
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
