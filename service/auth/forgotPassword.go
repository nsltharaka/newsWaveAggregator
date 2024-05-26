package auth

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/utils"
	"gopkg.in/gomail.v2"
)

var (
	chnCaseDeleteSignal = make(chan struct{})
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

	// 6 digit code associated with the case
	code := generateCode(6)

	// if there is an existing password reset request for the user
	_, err = h.db.GetCaseForUser(r.Context(), user.ID)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user cannot send password reset requests more than once in a 10 minutes time frame"))
		return
	}

	// create an entry in the database
	hashedCode, _ := HashPassword(code)
	caseId := uuid.New()
	if _, err = h.db.CreatePasswordResetCase(r.Context(), database.CreatePasswordResetCaseParams{
		CaseNumber: caseId,
		Code:       hashedCode,
		UserID:     user.ID,
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// send the email with the code
	if err := sendEmail(user.Username, payload.Email, code); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("sending 6-digit code failed for : %s", payload.Email))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"msg":    fmt.Sprintf("an email with 6-digit code has been sent to %s", payload.Email),
		"userId": user.ID,
	})

	// reset case get deleted after one minute.
	go func(db *database.Queries, caseId uuid.UUID) {
		timer := time.NewTimer(10 * time.Minute)
		<-timer.C

		db.DeleteCase(context.Background(), caseId)
		log.Printf("DELETED case_number %s closed.\n", caseId)

	}(h.db, caseId)

}

func (h *Handler) handleCase(w http.ResponseWriter, r *http.Request) {

	// get the 6-digit code and case number
	payload := struct {
		Code     string `json:"code"`
		UserId   int    `json:"user_id"`
		Password string `json:"password"`
	}{}

	// payload validation
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("verification failed"))
		return
	}

	// check with db
	resetCase, err := h.db.GetCaseForUser(r.Context(), int32(payload.UserId))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("verification failed"))
		return
	}

	// code validation
	if !VerifyPassword(resetCase.Code, payload.Code) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("verification failed"))
		return
	}

	// password validation
	err = utils.Validate.Var(payload.Password, "required,min=8,max=130")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("password should be at least 8 characters long"))
		return
	}

	// password hashing
	hashedPassword, _ := HashPassword(payload.Password)
	if err := h.db.UpdatePasswordForUser(r.Context(), database.UpdatePasswordForUserParams{
		Password: hashedPassword,
		ID:       int32(payload.UserId),
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("password reset failed"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"msg": "password has been reset",
	})

}

func sendEmail(username, email, code string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "nsltharaka@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset - NewsWave")
	m.SetBody("text/html", generateTemplate(username, code))

	d := gomail.NewDialer("smtp.gmail.com", 587, "nsltharaka@gmail.com", "achkubgvrbveoplr")

	return d.DialAndSend(m)
}

func generateTemplate(username, code string) string {
	return fmt.Sprintf(`
<p>Hey %s,</p>
<br>
<p>It seems like you forgot your password. If this is true, use the following 6 digit code to continue the process.</p>
<br>
<p>please note that this code will only be valid for 10 minutes.</p>
<br>
<p>code: <b>%s</b></p>
<br><br>
If you did not forget your password, please disregard this email.
`, username, code)
}

func generateCode(digitsCount int) string {
	num := ""
	for range digitsCount {
		num += strconv.Itoa(rand.Intn(10))
	}

	return num
}
