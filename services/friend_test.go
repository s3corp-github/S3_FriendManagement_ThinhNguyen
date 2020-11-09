package services

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFriendService_CreateFriend(t *testing.T) {
	testCases := []struct {
		name          string
		input         *model.FriendsServiceInput
		expectedErr   error
		mockRepoInput *model.FriendsRepoInput
		mockRepoErr   error
	}{
		{
			name: "Create friend connection failed with error",
			input: &model.FriendsServiceInput{
				FirstID:  1,
				SecondID: 2,
			},
			expectedErr: errors.New("create friend connection failed with error"),
			mockRepoInput: &model.FriendsRepoInput{
				FirstID:  1,
				SecondID: 2,
			},
			mockRepoErr: errors.New("create friend connection failed with error"),
		},
		{
			name: "Create friend connection success",
			input: &model.FriendsServiceInput{
				FirstID:  1,
				SecondID: 2,
			},
			expectedErr: errors.New("create friend connection failed with error"),
			mockRepoInput: &model.FriendsRepoInput{
				FirstID:  1,
				SecondID: 2,
			},
			mockRepoErr: errors.New("create friend connection failed with error"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockFriendRepo := new(mockFriendRepo)
			mockFriendRepo.On("CreateFriend", testCase.mockRepoInput).
				Return(testCase.mockRepoErr)
			service := FriendService{
				IFriendRepo: mockFriendRepo,
			}
			//When
			err := service.CreateFriend(testCase.input)

			//Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})

	}
}

func TestFriendService_GetFriendListByID(t *testing.T) {
	type mockGetFriendListByID struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockedListByID struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockingListByID struct {
		input  int
		result []int
		err    error
	}
	type mockGetEmailListByIDs struct {
		input  []int
		result []string
		err    error
	}
	testCases := []struct {
		name                string
		input               int
		expectedResult      []string
		expectedErr         error
		mockGetFriendsList  mockGetFriendListByID
		mockGetBlockedList  mockGetBlockedListByID
		mockGetBlockingList mockGetBlockingListByID
		mockGetEmailList    mockGetEmailListByIDs
	}{
		{
			name:           "Get friends list failed with error",
			input:          1,
			expectedResult: nil,
			expectedErr:    errors.New("get friends list failed with error"),
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: nil,
				err:    errors.New("get friends list failed with error"),
			},
		},
		{
			name:           "get blocking list failed",
			input:          1,
			expectedResult: nil,
			expectedErr:    errors.New("get blocking list failed with error"),
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockingList: mockGetBlockingListByID{
				input:  1,
				result: nil,
				err:    errors.New("get blocking list failed with error"),
			},
		},
		{
			name:           "Get email list by IDs failed with error",
			input:          1,
			expectedResult: nil,
			expectedErr:    errors.New("get email list by userIDs failed with error"),
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: []int{2, 3, 4, 5},
				err:    nil,
			},
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockingList: mockGetBlockingListByID{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailList: mockGetEmailListByIDs{
				input:  []int{2, 5},
				result: nil,
				err:    errors.New("get email list by userIDs failed with error"),
			},
		},
		{
			name:           "Get friend connection list success",
			input:          1,
			expectedResult: []string{"xyz@xyz.com", "xyzk@abc.com"},
			expectedErr:    nil,
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: []int{2, 3, 4, 5},
				err:    nil,
			},
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockingList: mockGetBlockingListByID{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailList: mockGetEmailListByIDs{
				input:  []int{2, 5},
				result: []string{"xyz@xyz.com", "xyzk@abc.com"},
				err:    nil,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriendRepo := new(mockFriendRepo)
			mockUserRepo := new(mockUserRepo)
			mockFriendRepo.On("GetFriendListByID", testCase.mockGetFriendsList.input).
				Return(testCase.mockGetFriendsList.result, testCase.mockGetFriendsList.err)
			mockFriendRepo.On("GetBlockedListByID", testCase.mockGetBlockedList.input).
				Return(testCase.mockGetBlockedList.result, testCase.mockGetBlockedList.err)
			mockFriendRepo.On("GetBlockingListByID", testCase.mockGetBlockingList.input).
				Return(testCase.mockGetBlockingList.result, testCase.mockGetBlockingList.err)
			mockUserRepo.On("GetEmailListByIDs", testCase.mockGetEmailList.input).
				Return(testCase.mockGetEmailList.result, testCase.mockGetEmailList.err)

			service := FriendService{
				IFriendRepo: mockFriendRepo,
				IUserRepo:   mockUserRepo,
			}

			// When
			result, err := service.GetFriendListByID(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.ElementsMatch(t, result, testCase.expectedResult)
			}
		})
	}
}

func TestFriendService_GetCommonFriendListByID(t *testing.T) {
	type mockGetFriendListByID struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockedListByID struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockingListByID struct {
		input  int
		result []int
		err    error
	}
	type mockGetEmailListByIDs struct {
		input  []int
		result []string
		err    error
	}
	testCases := []struct {
		name                          string
		input                         []int
		expectedResult                []string
		expectedErr                   error
		mockGetFriendsListFirstUser   mockGetFriendListByID
		mockGetBlockedListFirstUser   mockGetBlockedListByID
		mockGetBlockingListFirstUser  mockGetBlockingListByID
		mockGetEmailsByFirstUSerList  mockGetEmailListByIDs
		mockGetFriendsListSecondUser  mockGetFriendListByID
		mockGetBlockedListSecondUser  mockGetBlockedListByID
		mockGetBlockingListSecondUser mockGetBlockingListByID
		mockGetEmailsBySecondUSerList mockGetEmailListByIDs
	}{
		{
			name:           "Get first user's friend list failed with error",
			input:          []int{1, 2},
			expectedResult: nil,
			expectedErr:    errors.New("get first user's friend list failed with error"),
			mockGetFriendsListFirstUser: mockGetFriendListByID{
				input:  1,
				result: nil,
				err:    errors.New("get first user's friend list failed with error"),
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriendRepo := new(mockFriendRepo)
			mockUserRepo := new(mockUserRepo)

			mockFriendRepo.On("GetFriendListByID", testCase.mockGetFriendsListFirstUser.input).
				Return(testCase.mockGetFriendsListFirstUser.result, testCase.mockGetFriendsListFirstUser.err)
			mockFriendRepo.On("GetBlockedListByID", testCase.mockGetBlockedListFirstUser.input).
				Return(testCase.mockGetBlockedListFirstUser.result, testCase.mockGetBlockedListFirstUser.err)
			mockFriendRepo.On("GetBlockingListByID", testCase.mockGetBlockingListFirstUser.input).
				Return(testCase.mockGetBlockingListFirstUser.result, testCase.mockGetBlockingListFirstUser.err)
			mockUserRepo.On("GetEmailListByIDs", testCase.mockGetEmailsByFirstUSerList.input).
				Return(testCase.mockGetEmailsByFirstUSerList.result, testCase.mockGetEmailsByFirstUSerList.err)

			mockFriendRepo.On("GetFriendListByID", testCase.mockGetEmailsBySecondUSerList.input).
				Return(testCase.mockGetEmailsBySecondUSerList.result, testCase.mockGetEmailsBySecondUSerList.err)
			mockFriendRepo.On("GetBlockedListByID", testCase.mockGetBlockedListSecondUser.input).
				Return(testCase.mockGetBlockedListSecondUser.result, testCase.mockGetBlockedListSecondUser.err)
			mockFriendRepo.On("GetBlockingListByID", testCase.mockGetBlockingListSecondUser.input).
				Return(testCase.mockGetBlockingListSecondUser.result, testCase.mockGetBlockingListSecondUser.err)
			mockUserRepo.On("GetEmailListByIDs", testCase.mockGetEmailsBySecondUSerList.input).
				Return(testCase.mockGetEmailsBySecondUSerList.result, testCase.mockGetEmailsBySecondUSerList.err)

			services := FriendService{
				IFriendRepo: mockFriendRepo,
				IUserRepo:   mockUserRepo,
			}

			// When
			result, err := services.GetCommonFriendListByID(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.ElementsMatch(t, result, testCase.expectedResult)
			}
		})
	}
}

func TestFriendService_IsBlockedByOtherEmail(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedResult bool
		expectedErr    error
		mockRepoInput  []int
		mockRepoResult bool
		mockRepoError  error
	}{
		{
			name:           "Check is blocked failed with error",
			input:          []int{1, 2},
			expectedResult: false,
			expectedErr:    errors.New("failed with error"),
			mockRepoInput:  []int{1, 2},
			mockRepoResult: false,
			mockRepoError:  errors.New("failed with error"),
		},

		{
			name:           "Check success",
			input:          []int{1, 2},
			expectedResult: true,
			expectedErr:    nil,
			mockRepoInput:  []int{1, 2},
			mockRepoResult: true,
			mockRepoError:  nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriendRepo := new(mockFriendRepo)
			mockFriendRepo.On("IsBlockedByOtherEmail", testCase.mockRepoInput[0], testCase.mockRepoInput[1]).
				Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := FriendService{
				IFriendRepo: mockFriendRepo,
			}

			// When
			result, err := service.IsBlockedByOtherEmail(testCase.input[0], testCase.input[1])

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedResult)
			}
		})
	}
}

func TestFriendService_IsExistedFriend(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedResult bool
		expectedErr    error
		mockRepoInput  []int
		mockRepoResult bool
		mockRepoError  error
	}{
		{
			name:           "check is existed friend failed",
			input:          []int{1, 2},
			expectedResult: false,
			expectedErr:    errors.New("query database failed"),
			mockRepoInput:  []int{1, 2},
			mockRepoResult: false,
			mockRepoError:  errors.New("query database failed"),
		},
		{
			name:           "check is existed friend successfully",
			input:          []int{1, 2},
			expectedResult: true,
			expectedErr:    nil,
			mockRepoInput:  []int{1, 2},
			mockRepoResult: true,
			mockRepoError:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriendRepo := new(mockFriendRepo)
			mockFriendRepo.On("IsExistedFriend", testCase.mockRepoInput[0], testCase.mockRepoInput[1]).
				Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := FriendService{
				IFriendRepo: mockFriendRepo,
			}

			// When
			result, err := service.IsExistedFriend(testCase.input[0], testCase.input[1])

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedResult)
			}
		})
	}
}

func TestFriendService_GetEmailsReceiveUpdate(t *testing.T) {
	type mockGetFriendListByID struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockedListByID struct {
		input  int
		result []int
		err    error
	}
	type mockGetSubscribers struct {
		input  int
		result []int
		err    error
	}
	type mockGetEmailListByIDs struct {
		input  []int
		result []string
		err    error
	}
	testCases := []struct {
		name                        string
		sender                      int
		mentionedEmailList          []string
		expectedResult              []string
		expectedErr                 error
		mockGetBlockedList          mockGetBlockedListByID
		mockGetEmailListBlockedUser mockGetEmailListByIDs
		mockGetFriendsList          mockGetFriendListByID
		mockGetEmailListFriendUser  mockGetEmailListByIDs
		mockGetSubscribers          mockGetSubscribers
		mockGetEmailListSubsUser    mockGetEmailListByIDs
	}{
		{
			name:           "Get blocked list failed with error",
			sender:         1,
			expectedResult: nil,
			expectedErr:    errors.New("failed with error"),
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: nil,
				err:    errors.New("failed with error"),
			},
		},
		{
			name:           "Get blocked email list failed with error",
			sender:         1,
			expectedResult: nil,
			expectedErr:    errors.New("failed with error"),
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetEmailListBlockedUser: mockGetEmailListByIDs{
				input:  []int{2},
				result: nil,
				err:    errors.New("failed with error"),
			},
		},
		{
			name:               "Get friend connection userID list failed with error",
			sender:             1,
			mentionedEmailList: nil,
			expectedResult:     nil,
			expectedErr:        errors.New("failed with error"),
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetEmailListBlockedUser: mockGetEmailListByIDs{
				input:  []int{2},
				result: []string{"abc@xyz.com"},
				err:    nil,
			},
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: nil,
				err:    errors.New("failed with error"),
			},
		},
		{
			name:               "Get friend connection email list failed with error",
			sender:             1,
			mentionedEmailList: nil,
			expectedResult:     nil,
			expectedErr:        errors.New("failed with error"),
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetEmailListBlockedUser: mockGetEmailListByIDs{
				input:  []int{2},
				result: []string{"abc@xyz.com"},
				err:    nil,
			},
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetEmailListFriendUser: mockGetEmailListByIDs{
				input:  []int{3},
				result: nil,
				err:    errors.New("failed with error"),
			},
		},
		{
			name:               "Get subscriber list userID failed with error",
			sender:             1,
			mentionedEmailList: nil,
			expectedResult:     nil,
			expectedErr:        errors.New("failed with error"),
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetEmailListBlockedUser: mockGetEmailListByIDs{
				input:  []int{2},
				result: []string{"abc@xyz.com"},
				err:    nil,
			},
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetEmailListFriendUser: mockGetEmailListByIDs{
				input:  []int{3},
				result: []string{"xyz@abc.com"},
				err:    nil,
			},
			mockGetSubscribers: mockGetSubscribers{
				input:  1,
				result: nil,
				err:    errors.New("failed with error"),
			},
		},
		{
			name:               "Get subscriber email list failed with error",
			sender:             1,
			mentionedEmailList: nil,
			expectedResult:     nil,
			expectedErr:        errors.New("failed with error"),
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetEmailListBlockedUser: mockGetEmailListByIDs{
				input:  []int{2},
				result: []string{"abc@xyz.com"},
				err:    nil,
			},
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetEmailListFriendUser: mockGetEmailListByIDs{
				input:  []int{3},
				result: []string{"xyz@abc.com"},
				err:    nil,
			},
			mockGetSubscribers: mockGetSubscribers{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailListSubsUser: mockGetEmailListByIDs{
				input:  []int{4},
				result: nil,
				err:    errors.New("failed with error"),
			},
		},
		{
			name:               "Get email list receiving updates success",
			sender:             1,
			mentionedEmailList: []string{"mentioned@gmail.com"},
			expectedResult:     []string{"xyzk@gmail.com", "mentioned@gmail.com"},
			expectedErr:        nil,
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{2, 3},
				err:    nil,
			},
			mockGetEmailListBlockedUser: mockGetEmailListByIDs{
				input:  []int{2, 3},
				result: []string{"abc@xyz.com", "xyz@abc.com"},
				err:    nil,
			},
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetEmailListFriendUser: mockGetEmailListByIDs{
				input:  []int{3},
				result: []string{"xyz@abc.com"},
				err:    nil,
			},
			mockGetSubscribers: mockGetSubscribers{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailListSubsUser: mockGetEmailListByIDs{
				input:  []int{4},
				result: []string{"xyzk@gmail.com"},
				err:    nil,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriendRepo := new(mockFriendRepo)
			mockUserRepo := new(mockUserRepo)
			mockFriendRepo.On("GetFriendListByID", testCase.mockGetFriendsList.input).
				Return(testCase.mockGetFriendsList.result, testCase.mockGetFriendsList.err)
			mockFriendRepo.On("GetBlockedListByID", testCase.mockGetBlockedList.input).
				Return(testCase.mockGetBlockedList.result, testCase.mockGetBlockedList.err)
			mockFriendRepo.On("GetSubscriberList", testCase.mockGetSubscribers.input).
				Return(testCase.mockGetSubscribers.result, testCase.mockGetSubscribers.err)
			mockUserRepo.On("GetEmailListByIDs", testCase.mockGetEmailListBlockedUser.input).
				Return(testCase.mockGetEmailListBlockedUser.result, testCase.mockGetEmailListBlockedUser.err)
			mockUserRepo.On("GetEmailListByIDs", testCase.mockGetEmailListFriendUser.input).
				Return(testCase.mockGetEmailListFriendUser.result, testCase.mockGetEmailListFriendUser.err)
			mockUserRepo.On("GetEmailListByIDs", testCase.mockGetEmailListSubsUser.input).
				Return(testCase.mockGetEmailListSubsUser.result, testCase.mockGetEmailListSubsUser.err)

			service := FriendService{
				IFriendRepo: mockFriendRepo,
				IUserRepo:   mockUserRepo,
			}

			// When
			result, err := service.GetEmailsReceiveUpdate(testCase.sender, testCase.mentionedEmailList)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedResult, result)
			}
		})
	}
}
