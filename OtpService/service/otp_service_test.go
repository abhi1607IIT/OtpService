package service

import (
	"OtpService/dal/models"
	"OtpService/dal/repo/mocks"
	"OtpService/dtos"
	"OtpService/utils"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type OtpServiceTestSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	repo    *mocks.MockOtpRepo
	service OtpService
}

func (suite *OtpServiceTestSuite) BeforeTest(suitName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockOtpRepo(suite.ctrl)
	suite.service = NewOtpService(suite.repo)
}

func (suite *OtpServiceTestSuite) AfterTest(suitName, testName string) {

}

func TestOtpServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OtpServiceTestSuite))
}

func (suite *OtpServiceTestSuite) TestRequireOtpSuccess() {

	userOtp := models.UserOtp{
		UserId:    "user1",
		Otp:       "1234",
		Status:    "CREATED",
		ValidTill: time.Now().Add(2 * time.Minute).Unix(),
	}

	suite.repo.EXPECT().CreateUserOtp(utils.EqUserOtpMatcher(&userOtp)).
		Return(nil)

	req := dtos.GenerateOtpRequest{UserId: "user1"}

	resp, err := suite.service.RequestOtp(&req)

	suite.Equal(err, nil)
	suite.Equal(resp.Otp, "1234")
}

func (suite *OtpServiceTestSuite) TestRequireOtpFailure() {

	userOtp := models.UserOtp{
		UserId:    "user1",
		Otp:       "1234",
		Status:    "CREATED",
		ValidTill: time.Now().Add(2 * time.Minute).Unix(),
	}

	suite.repo.EXPECT().CreateUserOtp(utils.EqUserOtpMatcher(&userOtp)).
		Return(errors.New("some_error"))

	req := dtos.GenerateOtpRequest{UserId: "user1"}

	_, err := suite.service.RequestOtp(&req)

	suite.NotEqual(err, nil)
}

func (suite *OtpServiceTestSuite) TestValidateOtpSuccess() {

	userOtp := models.UserOtp{
		UserId:    "user1",
		Otp:       "1234",
		Status:    "CREATED",
		ValidTill: time.Now().Add(2 * time.Minute).Unix(),
	}

	updatedUserOtp := models.UserOtp{
		UserId:    "user1",
		Otp:       "1234",
		Status:    "VALIDATED",
		ValidTill: time.Now().Add(2 * time.Minute).Unix(),
	}

	suite.repo.EXPECT().GetByUserIdAndOtp("user1", "1234").
		Return(&userOtp, nil)

	suite.repo.EXPECT().SaveUserOtp(utils.EqUserOtpMatcher(&updatedUserOtp)).
		Return(nil)

	req := dtos.ValidateOtpRequest{
		UserId: "user1",
		Otp:    "1234",
	}

	resp, err := suite.service.ValidateOtp(&req)

	suite.Equal(err, nil)
	suite.Equal(resp.Msg, "otp validated successfully")
}

func (suite *OtpServiceTestSuite) TestValidateOtpExpired() {
	userOtp := models.UserOtp{
		UserId:    "user1",
		Otp:       "1234",
		Status:    "CREATED",
		ValidTill: time.Now().Unix() - 200,
	}

	updatedUserOtp := models.UserOtp{
		UserId:    "user1",
		Otp:       "1234",
		Status:    "EXPIRED",
		ValidTill: time.Now().Add(2 * time.Minute).Unix(),
	}

	suite.repo.EXPECT().GetByUserIdAndOtp("user1", "1234").
		Return(&userOtp, nil)

	suite.repo.EXPECT().SaveUserOtp(utils.EqUserOtpMatcher(&updatedUserOtp)).
		Return(nil)

	req := dtos.ValidateOtpRequest{
		UserId: "user1",
		Otp:    "1234",
	}

	resp, err := suite.service.ValidateOtp(&req)

	suite.Equal(err, nil)
	suite.Equal(resp.Error, "otp_expired")
}

func (suite *OtpServiceTestSuite) TestValidateOtpValidated() {

	userOtp := models.UserOtp{
		UserId:    "user1",
		Otp:       "1234",
		Status:    "VALIDATED",
		ValidTill: time.Now().Add(2 * time.Minute).Unix(),
	}

	suite.repo.EXPECT().GetByUserIdAndOtp("user1", "1234").
		Return(&userOtp, nil)

	req := dtos.ValidateOtpRequest{
		UserId: "user1",
		Otp:    "1234",
	}

	resp, err := suite.service.ValidateOtp(&req)

	suite.Equal(err, nil)
	suite.Equal(resp.Error, "otp_already_verified")
}
