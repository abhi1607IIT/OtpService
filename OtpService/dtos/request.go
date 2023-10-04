package dtos

type GenerateOtpRequest struct {
	UserId string `json:"user_id"`
}

type ValidateOtpRequest struct {
	UserId string `json:"user_id"`
	Otp    string `json:"otp"`
}
