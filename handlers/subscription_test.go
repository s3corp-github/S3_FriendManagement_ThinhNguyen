package handlers

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubscriptionHandler_CreateSubscription(t *testing.T) {
	type mockGetUserIDFromEmail struct {
		input  string
		result int
		err    error
	}
	type mockIsExistedSubscription struct {
		input  []int
		result bool
		err    error
	}
	type mockIsBlocked struct {
		input  []int
		result bool
		err    error
	}
	type mockCreateSubscription struct {
		input *model.SubscriptionServiceInput
		err   error
	}
	testCases := []struct {
		name                      string
		requestBody               map[string]interface{}
		expectedResponseBody      string
		expectedStatus            int
		mockGetRequestorUserID    mockGetUserIDFromEmail
		mockGetTargetUserID       mockGetUserIDFromEmail
		mockIsExistedSubscription mockIsExistedSubscription
		mockIsBlocked             mockIsBlocked
		mockCreateSubscription    mockCreateSubscription
	}{
		{
			name: "Body no data",
			requestBody: map[string]interface{}{
				"": "",
			},
			expectedResponseBody: "\"requestor\" is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockUserService := new(mockUserService)
			mockSubscriptionService := new(mockSubscriptionService)

			handlers := SubscriptionHandler{
				IUserService:         mockUserService,
				ISubscriptionService: mockSubscriptionService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}
			//When
			req, err := http.NewRequest(http.MethodPost, "/subscription", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}

			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.CreateSubscription)
			handler.ServeHTTP(responseRecorder, req)

			// Then
			require.Equal(t, testCase.expectedStatus, responseRecorder.Code)
			require.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}
