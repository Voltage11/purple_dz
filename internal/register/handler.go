package register

import (
	"purple1/pkg/tadapter"
	"purple1/pkg/validator"
	"purple1/views/components"
	"purple1/views/pages"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type RegisterHandler struct {
	router       fiber.Router
	customLogger zerolog.Logger
}

func NewHadnler(router fiber.Router) {
	h := &RegisterHandler{
		router:       router,
	}

	h.router.Get("/register", h.getRegisterUser)
	h.router.Post("/api/register", h.postRegisterUser)
}

func (h *RegisterHandler) getRegisterUser(c *fiber.Ctx) error {
	component := pages.Register()
	return tadapter.Render(c, component)
}

func (h *RegisterHandler) postRegisterUser(c *fiber.Ctx) error {
	form := RegisterCreateForm{
		Name: c.FormValue("name"),
		Email: c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	errors := validate.Validate(
		&validators.StringIsPresent{Name: "name", Field: form.Name, Message: "Укажите имя"},
		&validators.EmailIsPresent{Name: "email", Field: form.Email, Message: "Укажите email"},
		&validators.StringIsPresent{Name: "password", Field: form.Password, Message: "Укажите пароль"},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component)
	}
	component = components.Notification("Пользователь зарегистрирован", components.NotificationSuccsess)
	return tadapter.Render(c, component)

}