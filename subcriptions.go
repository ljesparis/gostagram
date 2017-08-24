package gostagram

import "errors"

var (
	SubscriptionError = errors.New("Subscriptions not yet supported.")
)

func (c Client) CreateSubscription() error {
	return SubscriptionError
}

func (c Client) ListSubscription() error {
	return SubscriptionError
}

func (c Client) DeleteSubscription() error {
	return SubscriptionError
}
