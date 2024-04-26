package repository_auth_server

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	requestmodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/requestModel"
	responsemodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/responseModel"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	email = "vajidhussain77@gmail.com"
)

var desiredTime = time.Date(2024, time.April, 25, 12, 0, 0, 0, time.UTC)

func TestSignup(t *testing.T) {
	testData := map[string]struct {
		arg     requestmodel_auth_server.UserSignup
		stub    func(sqlmock.Sqlmock)
		want    *responsemodel_auth_server.UserSignup
		wantErr error
	}{
		"succesfull signup": {
			arg: requestmodel_auth_server.UserSignup{
				UserName:  "vajid77",
				Name:      "vajid",
				Email:     email,
				Password:  "12387",
				CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			stub: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, user_name, email, password, created_at) VALUES($1, $2, $3, $4, $5) RETURNING *")).
					WillReturnRows(mockSql.NewRows([]string{"id", "user_name", "name", "email", "create_at"}).AddRow("12", "vajid77", "vajid", email, desiredTime))
			},
			want: &responsemodel_auth_server.UserSignup{
				ID:        "12",
				UserName:  "vajid77",
				Name:      "vajid",
				Email:     email,
				CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		"signup failed": {
			arg: requestmodel_auth_server.UserSignup{
				UserName:  "vajid77",
				Name:      "vajid",
				Email:     email,
				Password:  "12387",
				CreatedAt: time.Date(2024, time.April, 25, 12, 0, 0, 0, time.UTC),
			},
			stub: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, user_name, email, password, created_at) VALUES($1, $2, $3, $4, $5) RETURNING *")).
					WillReturnRows(sqlmock.NewRows([]string{}))
			},
			want:    nil,
			wantErr: errors.New("no row affected"),
		},
	}

	for _, test := range testData {
		mockDB, mockSql, _ := sqlmock.New()
		defer mockDB.Close()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockDB,
		}), &gorm.Config{})

		test.stub(mockSql)

		UserRepository := NewUserRepository(DB)

		result, err := UserRepository.Signup(test.arg)

		assert.Equal(t, test.want, result)
		assert.Equal(t, test.wantErr, err)
	}
}

func TestUserNameIsExist(t *testing.T) {
	testData := map[string]struct {
		arg     string
		stub    func(sqlmock.Sqlmock)
		want    int
		wantErr error
	}{
		"success": {
			arg: "vajid",
			stub: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM users WHERE user_name = $1  AND status != 'delete'")).
					WillReturnRows(mockSql.NewRows([]string{}))
			},
			want:    0,
			wantErr: errors.New("no row affected"),
		},
		"user exist": {
			arg: "vajid",
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM users WHERE user_name = $1  AND status != 'delete'")).
					WillReturnRows(s.NewRows([]string{"name"}).AddRow(2))
			},
			want:    2,
			wantErr: nil,
		},
	}

	for _, test := range testData {
		mockdDB, mockSql, _ := sqlmock.New()
		defer mockdDB.Close()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockdDB,
		}), &gorm.Config{})

		test.stub(mockSql)

		UserRepository := NewUserRepository(DB)

		result, err := UserRepository.UserNameIsExist(test.arg)

		assert.Equal(t, test.want, result)
		assert.Equal(t, test.wantErr, err)
	}
}
