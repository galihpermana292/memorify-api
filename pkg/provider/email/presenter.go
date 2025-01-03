package email

// DTOConfirmationEmail dto withdrawal incentive reseller
type DTOConfirmationEmail struct {
	PaymentID string `json:"payment_id"`
	Username  string `json:"username"`
}
