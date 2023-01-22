package lotusgo

const (
	GET_CUSTOMERS               = "/app/customers/"
	CREATE_CUSTOMERS            = "/app/customers/"
	CREATE_BATCH_CUSTOMERS      = "/api/batch_create_customers/"
	GET_CUSTOMER_DETAIL         = `/app/customers/` ///app/customers/${customerId}/
	CREATE_SUBSCRIPTION         = "/app/subscriptions/"
	CANCEL_SUBSCRIPTION         = `/app/subscriptions/cancel/`
	CHANGE_SUBSCRIPTION         = `/app/subscriptions/update/`
	GET_ALL_SUBSCRIPTIONS       = "/app/subscriptions/"
	GET_SUBSCRIPTION_DETAILS    = `/app/subscriptions/` ///app/subscriptions/${subscriptionId}/
	GET_ALL_PLANS               = "/api/plans/"
	GET_PLAN_DETAILS            = `/api/plans/` /// api/plans/${planId}/
	GET_CUSTOMER_FEATURE_ACCESS = "/api/customer_feature_access/"
	GET_CUSTOMER_METRIC_ACCESS  = "/api/customer_metric_access/"
	TRACK_EVENT                 = "/api/track/"
	GET_INVOICES                = "/api/invoices/"
)

const (
	CreateCustomer         = "createCustomer"
	Track_Event            = "trackEvent"
	CustomerDetails        = "customerDetails"
	Create_Subscription    = "createSubscription"
	CancelSubscription     = "cancelSubscription"
	ChangeSubscription     = "changeSubscription"
	SubscriptionDetails    = "subscriptionDetails"
	PlanDetails            = "planDetails"
	CustomerMetricAccess   = "customerMetricAccess"
	Customer_FeatureAccess = "customerFeatureAccess"
	CreateCustomersBatch   = "createCustomersBatch"
	GetInvoices            = "getInvoices"
)
