package endpoint

import (
	"errors"
	"net/http/httptest"
	"testing"

	"restapi/models"
	"restapi/services/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedSVC := mocks.NewMockUserService(ctrl)
	endPoint := &UserEndPoint{svc: mockedSVC}
	w := &httptest.ResponseRecorder{}
	ctx, _ := gin.CreateTestContext(w)

	var testCases = []struct {
		ctx               *gin.Context
		listUserSVCResult []models.User
		listErrSVCErr     error
		statusCode        int
		expStatusCode     int
	}{
		{
			ctx:               ctx,
			listUserSVCResult: []models.User{models.User{ID: 1, Name: "test"}},
			expStatusCode:     200,
		},
		{
			ctx:           ctx,
			listErrSVCErr: errors.New("500 error"),
			expStatusCode: 500,
		},
	}
	for _, tc := range testCases {
		mockedSVC.
			EXPECT().
			ListUser(tc.ctx).
			Return(tc.listUserSVCResult, tc.listErrSVCErr)

		endPoint.ListUser(tc.ctx)
		assert.Equal(t, tc.expStatusCode, tc.ctx.Writer.Status())
	}
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedSVC := mocks.NewMockUserService(ctrl)
	endPoint := &UserEndPoint{svc: mockedSVC}
	w := &httptest.ResponseRecorder{}
	ctx, _ := gin.CreateTestContext(w)

	var testCases = []struct {
		ctx              *gin.Context
		ID               string
		GetUserSVCResult *models.User
		GetErrSVCErr     error
		statusCode       int
		expStatusCode    int
	}{
		{
			ctx:              ctx,
			ID:               "1",
			GetUserSVCResult: &models.User{ID: 1, Name: "test"},
			statusCode:       200,
			expStatusCode:    200,
		},
		{
			ctx:              ctx,
			ID:               "1",
			GetUserSVCResult: &models.User{ID: 1, Name: "test"},
			GetErrSVCErr:     errors.New("500 err"),
			statusCode:       500,
			expStatusCode:    500,
		},
		{
			ctx:              ctx,
			ID:               "1",
			GetUserSVCResult: &models.User{ID: 1, Name: "test"},
			GetErrSVCErr:     errors.New("record not found"),
			statusCode:       404,
			expStatusCode:    404,
		},
	}

	for _, tc := range testCases {
		mockedSVC.
			EXPECT().
			GetUser(tc.ctx, tc.GetUserSVCResult.ID).
			Return(tc.GetUserSVCResult, tc.GetErrSVCErr)

		tc.ctx.Params = []gin.Param{{Key: "id", Value: tc.ID}}
		endPoint.GetUser(tc.ctx)
		assert.Equal(t, tc.expStatusCode, tc.ctx.Writer.Status())
	}
}
