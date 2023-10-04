package utils

import (
	"OtpService/dal/models"
	"github.com/golang/mock/gomock"
	"reflect"
)

type UserOtpMatcher struct {
	Expected *models.UserOtp
}

func (m *UserOtpMatcher) String() string {
	return "UserOtpMatcher"
}

func (m *UserOtpMatcher) Matches(x interface{}) bool {
	actual, ok := x.(*models.UserOtp)

	if ok {
		actual.ValidTill = m.Expected.ValidTill
		actual.Otp = m.Expected.Otp
	}

	return reflect.DeepEqual(actual, m.Expected)
}

func EqUserOtpMatcher(otp *models.UserOtp) gomock.Matcher {
	return &UserOtpMatcher{otp}
}
