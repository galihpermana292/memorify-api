package email

// DTOConfirmationEmail dto withdrawal incentive reseller
type DTOConfirmationEmail struct {
	PaymentID    string `json:"payment_id"`
	Username     string `json:"username"`
	UserEmail    string `json:"user_email"`
	PaymentProof string `json:"payment_proof"`
	Date         string `json:"date"`
	Amount       string `json:"amount"`
}
