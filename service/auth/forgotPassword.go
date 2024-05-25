package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/utils"
	"gopkg.in/gomail.v2"
)

func (h *Handler) handleForgotPassword(w http.ResponseWriter, r *http.Request) {
	// get user email from the request
	// example request body
	payload := struct {
		Email string `json:"email"`
	}{}

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// check there is a user with the given email
	user, err := h.db.GetUserByEmail(r.Context(), payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("cannot find a matching account for the given email"))
		return
	}

	// check if there is an existing password reset request for the user
	_, err = h.db.GetCaseForUser(r.Context(), user.ID)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("there is an existing password reset request for this user. try again in 1 hour"))
		return
	}

	// create case for the user
	caseNumber := uuid.New()
	h.db.CreatePasswordResetCase(r.Context(), database.CreatePasswordResetCaseParams{
		CaseNumber: caseNumber,
		UserID:     user.ID,
		Opened:     false,
	})

	// reset link
	baseUrl := os.Getenv("API_BASE_URL")
	resetLink := fmt.Sprintf("%s/auth/forgot-password?case-number=%s", baseUrl, caseNumber)

	// send email
	// app password : achk ubgv rbve oplr
	if err := sendEmail(user.Username, payload.Email, resetLink); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("couldn't send the email. please try again"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"msg": fmt.Sprintf("an email has been sent to %s.", payload.Email),
	})

}

func (h *Handler) handleCase(w http.ResponseWriter, r *http.Request) {

	// get the case number
	caseNumber := r.URL.Query().Get("case-number")
	w.Write([]byte(caseNumber))

}

func sendEmail(username, email, resetLink string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "nsltharaka@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset - NewsWave")
	m.SetBody("text/html", generateTemplate(username, resetLink))

	d := gomail.NewDialer("smtp.gmail.com", 587, "nsltharaka@gmail.com", "achkubgvrbveoplr")

	return d.DialAndSend(m)
}

func generateTemplate(username, resetLink string) string {
	return fmt.Sprintf(`
<p>Hey %s,</p>
<br>
<p>It seems like you forgot your password. If this is true, click the link below to reset your password.</p>
<br>
<a href="%s" target="_blank">Reset my password</a>
<br><br>
If you did not forget your password, please disregard this email.
`, username, resetLink)
}
