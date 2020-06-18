package services

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type mockUserOption struct {
	list    bool
	listErr bool
	get     bool
	getErr  bool
}

func TestListUser(t *testing.T) {
	w := &httptest.ResponseRecorder{}
	ctx, _ := gin.CreateTestContext(w)
	db, mock, _ := sqlmock.New()
	defer db.Close()
	gdb, _ := gorm.Open("mysql", db)
	ctx.Keys = map[string]interface{}{}
	ctx.Keys["db"] = gdb

	var testCases = []struct {
		ctx    *gin.Context
		opt    mockUserOption
		expLen int
		expErr bool
	}{
		{
			ctx:    ctx,
			opt:    mockUserOption{list: true},
			expLen: 1,
		},
		{
			ctx:    ctx,
			opt:    mockUserOption{listErr: true},
			expErr: true,
		},
	}

	svc := NewUserService()
	for _, tc := range testCases {
		mockUserStep(mock, tc.opt)
		users, err := svc.ListUser(tc.ctx)
		if err != nil {
			assert.True(t, tc.expErr)
		} else {
			assert.Equal(t, tc.expLen, len(users))
			assert.False(t, tc.expErr)
		}
	}
}

func TestGetUser(t *testing.T) {
	w := &httptest.ResponseRecorder{}
	ctx, _ := gin.CreateTestContext(w)
	db, mock, _ := sqlmock.New()
	defer db.Close()
	gdb, _ := gorm.Open("mysql", db)
	ctx.Keys = map[string]interface{}{}
	ctx.Keys["db"] = gdb

	var testCases = []struct {
		ctx    *gin.Context
		opt    mockUserOption
		ID     int
		expID  int
		expErr bool
	}{
		{
			ctx:   ctx,
			opt:   mockUserOption{get: true},
			ID:    1,
			expID: 1,
		},
		{
			ctx:    ctx,
			opt:    mockUserOption{getErr: true},
			ID:     1,
			expErr: true,
		},
	}

	svc := NewUserService()
	for _, tc := range testCases {
		mockUserStep(mock, tc.opt)
		user, err := svc.GetUser(tc.ctx, tc.ID)
		if err != nil {
			assert.True(t, tc.expErr)
		} else {
			assert.Equal(t, tc.expID, user.ID)
			assert.False(t, tc.expErr)
		}
	}
}

func mockUserStep(mock sqlmock.Sqlmock, opt mockUserOption) {
	row := []string{"id", "name"}
	if opt.list {
		sqlmock.NewRows(row)
		mock.ExpectQuery("^SELECT").WillReturnRows(sqlmock.NewRows(row).AddRow("1", "test"))
		return
	}
	if opt.listErr {
		mock.ExpectQuery("^SELECT").WillReturnError(errors.New("500"))
		return
	}
	if opt.get {
		sqlmock.NewRows(row)
		mock.ExpectQuery("^SELECT").WillReturnRows(sqlmock.NewRows(row).AddRow("1", "test"))
		return
	}
	if opt.getErr {
		mock.ExpectQuery("^SELECT").WillReturnError(errors.New("500"))
		return
	}
}
