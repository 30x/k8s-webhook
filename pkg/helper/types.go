package helper

//AuthzRequestBody is the go struct for a k8s authz request
type AuthzRequestBody struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Spec       struct {
		ResourceAttributes struct {
			Namespace string `json:"namespace"`
			Verb      string `json:"verb"`
			Group     string `json:"group"`
			Resource  string `json:"resource"`
		} `json:"resourceAttributes"`
		User  string   `json:"user"`
		Group []string `json:"group"`
	} `json:"spec"`
}

//AuthzResponseBody is the go struct for a k8s authz response
type AuthzResponseBody struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Status     struct {
		Allowed bool   `json:"allowed"`
		Reason  string `json:"reason"`
	} `json:"status"`
}

//Status is the status of a k8s authz response
type Status struct {
	Allowed bool   `json:"allowed"`
	Reason  string `json:"reason"`
}
