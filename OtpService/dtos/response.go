package dtos

type GenerateOtpResponse struct {
	UserId string `json:"user_id"`
	Otp    string `json:"otp"`
}

type ValidateOtpResponse struct {
	UserId           string `json:"user_id,omitempty"`
	Msg              string `json:"msg,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}
