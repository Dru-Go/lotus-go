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
	Name             string `json:"customer_name,omitempty"`
	Id               string `json:"customer_id,omitempty"`
	Email            string `json:"email,omitempty"`
	Payment_provider string `json:"payment_provider,omitempty"`
	Properties       any    `json:"properties,omitempty"`
	Integrations     any    `json:"integrations,omitempty"`
}

type CreateCustomerParams struct {
	CustomerId          string `json:"customer_id,omitempty" url:"customer_id"`
	Email               string `json:"email,omitempty" url:"email"`
	PaymentProvider     string `json:"payment_provider,omitempty" url:"payment_provider,omitempty"`
	PaymentProviderId   string `json:"payment_provider_id,omitempty" url:"payment_provider_id,omitempty"`
	CustomerName        string `json:"customer_name,omitempty" url:"customer_name"`
	Properties          string `json:"properties,omitempty" url:"properties,omitempty"`
	Integrations        string `json:"integrations,omitempty" url:"integrations,omitempty"`
	DefaultCurrencyCode string `json:"default_currency_code,omitempty" url:"default_currency_code,omitempty"`
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

type Integrations struct {
	property1 any
	property2 any
}

type DefaultCurrency struct {
	code   string
	name   string
	symbol string
}

type CustomerResponse struct {
	CustomerId       string               `json:"customer_id,omitempty"`
	Email            string               `json:"email,omitempty"`
	CustomerName     string               `json:"customer_name,omitempty"`
	Invoices         []Invoice            `json:"invoices,omitempty"`
	TotalAmountDue   float32              `json:"total_amount_due,omitempty"`
	Subscriptions    []CreateSubscription `json:"subscriptions,omitempty"`
	Integrations     Integrations         `json:"integrations,omitempty"`
	DefaultCurrency  DefaultCurrency      `json:"default_currency,omitempty"`
	HasPaymentMethod bool                 `json:"has_payment_method,omitempty"`
	PaymentProvider  string               `json:"payment_provider,omitempty"`
	Address          any                  `json:"address,omitempty"`
	TaxRate          int16                `json:"tax_rate,omitempty"`
}

// SUBSCRIPTIONS

type LightPlan struct {
	plan_name string
	plan_id   string
	version   string
}

type LightCustomer struct {
	customer_name string
	email         string
	customer_id   string
}

type CreateSubscription struct {
	customer             LightCustomer
	subscription_filters []SubscriptionFilter
	start_date           string
	end_date             string
	fully_billed         bool
	is_new               bool
	auto_renew           bool
	billing_plan         LightPlan
}

type SubscriptionFilter struct {
	value         string
	property_name string
}

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

// INVOICE

type Invoice struct {
	external_payment_obj_type string
	invoice_number            string
	currency                  Currency
	external_payment_obj_id   string
	due_date                  time.Time
	payment_status            string
	issue_date                time.Time
	cost_due                  int32
}

type GetInvoicesParams struct {
	CustomerId    string `json:"customerId,omitempty"`
	PaymentStatus string `json:"paymentStatus,omitempty"` //"paid" | "unpaid"
}
