// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks_repository

import (
	context "context"

	domain "github.com/paw1a/ecommerce-api/internal/domain"
	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	testing "testing"
)

// Reviews is an autogenerated mock type for the Reviews type
type Reviews struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, review
func (_m *Reviews) Create(ctx context.Context, review domain.Review) (domain.Review, error) {
	ret := _m.Called(ctx, review)

	var r0 domain.Review
	if rf, ok := ret.Get(0).(func(context.Context, domain.Review) domain.Review); ok {
		r0 = rf(ctx, review)
	} else {
		r0 = ret.Get(0).(domain.Review)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Review) error); ok {
		r1 = rf(ctx, review)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, reviewID
func (_m *Reviews) Delete(ctx context.Context, reviewID primitive.ObjectID) error {
	ret := _m.Called(ctx, reviewID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(ctx, reviewID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByProductID provides a mock function with given fields: ctx, productID
func (_m *Reviews) DeleteByProductID(ctx context.Context, productID primitive.ObjectID) error {
	ret := _m.Called(ctx, productID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(ctx, productID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx
func (_m *Reviews) FindAll(ctx context.Context) ([]domain.Review, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Review
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Review); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Review)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, reviewID
func (_m *Reviews) FindByID(ctx context.Context, reviewID primitive.ObjectID) (domain.Review, error) {
	ret := _m.Called(ctx, reviewID)

	var r0 domain.Review
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) domain.Review); ok {
		r0 = rf(ctx, reviewID)
	} else {
		r0 = ret.Get(0).(domain.Review)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, reviewID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByProductID provides a mock function with given fields: ctx, productID
func (_m *Reviews) FindByProductID(ctx context.Context, productID primitive.ObjectID) ([]domain.Review, error) {
	ret := _m.Called(ctx, productID)

	var r0 []domain.Review
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) []domain.Review); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Review)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserID provides a mock function with given fields: ctx, userID
func (_m *Reviews) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]domain.Review, error) {
	ret := _m.Called(ctx, userID)

	var r0 []domain.Review
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) []domain.Review); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Review)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewReviews creates a new instance of Reviews. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewReviews(t testing.TB) *Reviews {
	mock := &Reviews{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}