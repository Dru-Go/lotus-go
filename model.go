package lotusgo

import (
	"time"
)

type any = interface{}

type Filters struct {
	PropertyName string `json:"propertyName,omitempty"`
	Value        string `json:"value,omitempty"`
}

// CUSTOMERS
type Customer struct {
	Name             string `json:"customerName,omitempty"`
	Id               string `json:"customerId,omitempty"`
	Email            string `json:"email,omitempty"`
	Payment_provider string `json:"payment_provider,omitempty"`
	Properties       any    `json:"properties,omitempty"`
	Integrations     any    `json:"integrations,omitempty"`
}

type CreateCustomerParams struct {
	CustomerId          string `json:"customerId,omitempty"`
	Email               string `json:"email,omitempty"`
	PaymentProvider     string `json:"paymentProvider,omitempty"`
	PaymentProviderId   string `json:"paymentProviderId,omitempty"`
	CustomerName        string `json:"customerName,omitempty"`
	Properties          string `json:"properties,omitempty"`
	Integrations        string `json:"integrations,omitempty"`
	DefaultCurrencyCode string `json:"default_currency_code,omitempty"`
}

type CreateBatchCustomerParams struct {
	Customers          []CreateCustomerParams `json:"customers,omitempty"`
	BehaviorOnExisting string                 `json:"behaviorOnExisting,omitempty"` // "merge" | "ignore" | "overwrite"
}

type CustomerDetailsParams struct {
	CustomerId string `json:"customerId,omitempty"`
}

type CustomerMetricAccessParams struct {
	CustomerId          string    `json:"customerId,omitempty"`
	EventName           string    `json:"eventName,omitempty"`
	SubscriptionFilters []Filters `json:"subscriptionFilters,omitempty"`
}

type CustomerFeatureAccess struct {
	CustomerId          string    `json:"customerId,omitempty"`
	FeatureName         string    `json:"featureName,omitempty"`
	SubscriptionFilters []Filters `json:"subscription_id,omitempty"`
}

type CustomerFeatureAccessResponse struct {
	Feature              string    `json:"feature,omitempty"`
	Subscription_id      string    `json:"subscription_id,omitempty"`
	Subscription_filters []Filters `json:"subscriptionFilters,omitempty"`
	Access               bool      `json:"access,omitempty"`
}

// SUBSCRIPTIONS
type CreateSubscriptionParams struct {
	CustomerId          string    `json:"customerId"`
	PlanId              string    `json:"planId"`
	StartDate           string    `json:"startDate"`
	EndDate             string    `json:"endDate,omitempty"`
	Status              string    `json:"status,omitempty"`
	AutoRenew           bool      `json:"autoRenew,omitempty"`
	IsNew               bool      `json:"isNew,omitempty"`
	SubscriptionId      string    `json:"subscriptionId,omitempty"`
	SubscriptionFilters []Filters `json:"subscriptionFilters,omitempty"`
}

type ChangeSubscriptionParams struct {
	CustomerId                   string    `json:"customerId"`
	PlanId                       string    `json:"planId,omitempty"`
	SubscriptionFilters          []Filters `json:"subscriptionFilters,omitempty"`
	ReplacePlanId                string    `json:"replacePlanId,omitempty"`
	ReplacePlanInvoicingBehavior string    `json:"replacePlanInvoicingBehavior,omitempty"` // "add_to_next_invoice" | "invoice_now"
	TurnOffAutoRenew             bool      `json:"turnOffAutoRenew,omitempty"`
	EndDate                      string    `json:"endDate,omitempty"`
}

type SubscriptionDetailsParams struct {
	SubscriptionId string `json:"subscriptionId,omitempty"`
}

type ListAllSubscriptionsParams struct {
	CustomerId string `json:"customerId,omitempty"`
	Status     string `json:"status,omitempty"` // "active" | "ended" | "not_started";
}

type CancelSubscriptionParams struct {
	PlanId                    string  `json:"planId"`
	BillUsage                 bool    `json:"billUsage,omitempty"`
	CustomerId                string  `json:"customerId,omitempty"`
	InvoicingBehaviorOnCancel string  `json:"invoicingBehaviorOnCancel,omitempty"` // "add_to_next_invoice" | "invoice_now"
	FlatFeeBehavior           string  `json:"flatFeeBehavior,omitempty"`           // "refund" | "prorate" | "charge_full"
	SubscriptionFilters       Filters `json:"subscriptionFilters,omitempty"`
}

// EVENT

type TrackEventEntity struct {
	EventName   string    `json:"eventName,omitempty"`
	CustomerId  string    `json:"customerId,omitempty"`
	ImpotencyId string    `json:"idempotencyId,omitempty"`
	TimeCreated time.Time `json:"timeCreated,omitempty"`
	Properties  any       `json:"properties,omitempty"`
}

type TrackEvent struct {
	Batch []TrackEventEntity `json:"batch,omitempty"`
}

// PLAN
type PlanDetailsParams struct {
	PlanId string `json:"planId,omitempty"`
}

type Currency struct {
	Code   string `json:"code,omitempty"`
	Name   string `json:"name,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}

type LineItems struct {
	Name                string    `json:"name,omitempty"`
	StartDate           string    `json:"start_date,omitempty"`
	EndDate             string    `json:"end_date,omitempty"`
	Quantity            int32     `json:"quantity,omitempty"`
	Subtotal            int32     `json:"subtotal,omitempty"`
	BillingType         string    `json:"billing_type,omitempty"`
	Metadata            any       `json:"metadata,omitempty"`
	PlanVersionId       string    `json:"plan_version_id,omitempty"`
	PlanName            string    `json:"plan_name,omitempty"`
	SubscriptionFilters []Filters `json:"subscription_filters,omitempty"`
}

type GetInvoicesParams struct {
	CustomerId    string `json:"customerId,omitempty"`
	PaymentStatus string `json:"paymentStatus,omitempty"` //"paid" | "unpaid"
}