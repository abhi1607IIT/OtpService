package models

type UserOtp struct {
	Id        int
	UserId    string
	Otp       string
	Status    string
	ValidTill int64
}
