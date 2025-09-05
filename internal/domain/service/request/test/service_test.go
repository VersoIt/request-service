package test

import (
	"RequestService/internal/domain/model"
	requestservice "RequestService/internal/domain/service/request"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestService_CreateRequest(t *testing.T) {
	t.Parallel()

	controller := gomock.NewController(t)
	defer controller.Finish()

	tests := []struct {
		name          string
		mockBehaviour func(rr *MockrequestRepository)
		request       model.Request
		userID        int64
		newRequestID  int64
		hasError      bool
	}{
		{
			name: "Success",
			mockBehaviour: func(rr *MockrequestRepository) {
				rr.EXPECT().CreateRequest(gomock.Any(), model.Request{}, int64(1)).Return(int64(2), nil)
				rr.EXPECT().GetHandlingRequests(gomock.Any(), int64(1)).Return(model.Requests{}, nil)
			},
			request:      model.Request{},
			userID:       1,
			newRequestID: 2,
			hasError:     false,
		},
		{
			name: "Get handling requests error",
			mockBehaviour: func(rr *MockrequestRepository) {
				rr.EXPECT().GetHandlingRequests(gomock.Any(), int64(1)).Return(model.Requests{}, errors.New("get error"))
			},
			request:      model.Request{},
			userID:       1,
			newRequestID: 0,
			hasError:     true,
		},
		{
			name: "Create request error",
			mockBehaviour: func(rr *MockrequestRepository) {
				rr.EXPECT().GetHandlingRequests(gomock.Any(), int64(1)).Return(model.Requests{}, nil)
				rr.EXPECT().CreateRequest(gomock.Any(), model.Request{}, int64(1)).Return(int64(0), errors.New("create request error"))
			},
			request:      model.Request{},
			userID:       1,
			newRequestID: 0,
			hasError:     true,
		},
		{
			name: "handling requests > 1",
			mockBehaviour: func(rr *MockrequestRepository) {
				rr.EXPECT().GetHandlingRequests(gomock.Any(), int64(1)).Return(model.Requests{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
				}, nil)
			},
			request:      model.Request{},
			userID:       1,
			newRequestID: 0,
			hasError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRequestRepo := NewMockrequestRepository(controller)

			tt.mockBehaviour(mockRequestRepo)

			service := requestservice.New(mockRequestRepo)

			id, err := service.CreateRequest(context.Background(), tt.request, tt.userID)
			if (err != nil) != tt.hasError {
				t.Fatalf("unexpected error: %v", err)
			}

			if id != tt.newRequestID {
				t.Fatalf("got: %v, want: %v", id, tt.newRequestID)
			}
		})
	}
}
