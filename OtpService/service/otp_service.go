package service

import (
	"OtpService/dal/models"
	"OtpService/dal/repo"
	"OtpService/dtos"
	"OtpService/utils"
	"time"
)

const (
	OtpCreatedStatus   = "CREATED"
	OtpValidatedStatus = "VALIDATED"
	OTPExpiredStatus   = "EXPIRED"
)

type otpService struct {
	otpRepo repo.OtpRepo
}

func NewOtpService(otpRepo repo.OtpRepo) OtpService {
	return &otpService{
		otpRepo,
	}
}

type OtpService interface {
	RequestOtp(req *dtos.GenerateOtpRequest) (*dtos.GenerateOtpResponse, error)
	ValidateOtp(req *dtos.ValidateOtpRequest) (*dtos.ValidateOtpResponse, error)
}

func (o *otpService) RequestOtp(req *dtos.GenerateOtpRequest) (*dtos.GenerateOtpResponse, error) {

	userOtp := models.UserOtp{
		UserId:    req.UserId,
		Otp:       utils.GenerateOtp(),
		Status:    OtpCreatedStatus,
		ValidTill: time.Now().Add(2 * time.Minute).Unix(),
	}
	err := o.otpRepo.CreateUserOtp(&userOtp)

	if err != nil {
		return nil, err
	}

	result := dtos.GenerateOtpResponse{
		UserId: req.UserId,
		Otp:    userOtp.Otp,
	}

	return &result, nil
}

func (o *otpService) ValidateOtp(req *dtos.ValidateOtpRequest) (*dtos.ValidateOtpResponse, error) {

	userOtp, err := o.otpRepo.GetByUserIdAndOtp(req.UserId, req.Otp)

	if err != nil {

		if err == repo.RecordNotFoundError {
			return &dtos.ValidateOtpResponse{
				Error:            "otp_not_found",
				ErrorDescription: "OTP not found",
			}, nil
		}

		return nil, err

	}

	if userOtp.Status == OtpValidatedStatus {
		return &dtos.ValidateOtpResponse{
			Error:            "otp_already_verified",
			ErrorDescription: "OTP already verified",
		}, nil
	}

	if time.Now().Unix() <= userOtp.ValidTill {

		userOtp.Status = OtpValidatedStatus
		err := o.otpRepo.SaveUserOtp(userOtp)

		if err != nil {
			return nil, err
		}

		return &dtos.ValidateOtpResponse{
			UserId: req.UserId,
			Msg:    "otp validated successfully",
		}, nil

	} else {
		userOtp.Status = OTPExpiredStatus
		err := o.otpRepo.SaveUserOtp(userOtp)

		if err != nil {
			return nil, err
		}

		return &dtos.ValidateOtpResponse{
			Error:            "otp_expired",
			ErrorDescription: "OTP expired",
		}, nil
	}
}
