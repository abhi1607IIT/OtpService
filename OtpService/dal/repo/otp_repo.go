package repo

import (
	"OtpService/dal/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

type otpRepo struct {
	db *gorm.DB
}

func NewOtpRepo(db *gorm.DB) OtpRepo {
	return &otpRepo{db: db}
}

const OtpTableName = "user_otp"

var RecordNotFoundError = fmt.Errorf("record not found")

//go:generate mockgen -destination=./mocks/otp_repo.go -package=mocks -source=./otp_repo.go
type OtpRepo interface {
	CreateUserOtp(userOtp *models.UserOtp) error
	GetByUserIdAndOtp(userId, otp string) (*models.UserOtp, error)
	SaveUserOtp(userOtp *models.UserOtp) error
}

func (o *otpRepo) CreateUserOtp(userOtp *models.UserOtp) error {
	return o.db.Table(OtpTableName).Create(userOtp).Error
}

func (o *otpRepo) GetByUserIdAndOtp(userId, otp string) (*models.UserOtp, error) {
	var userOtp models.UserOtp
	res := o.db.Table(OtpTableName).
		Where("user_id=?", userId).
		Where("otp=?", otp).
		First(&userOtp)

	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, RecordNotFoundError
		}

		return nil, res.Error
	}

	return &userOtp, nil
}

func (o *otpRepo) SaveUserOtp(userOtp *models.UserOtp) error {
	return o.db.Table(OtpTableName).Save(userOtp).Error
}
