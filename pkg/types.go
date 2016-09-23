package k8s-webhook

type k8sAuthzRequestBody struct {
	APIVersion string `json:"apiVersion"`
	Kind string `json:"kind"`
	Spec struct {
		ResourceAttributes struct {
			Namespace string `json:"namespace"`
			Verb string `json:"verb"`
			Group string `json:"group"`
			Resource string `json:"resource"`
		} `json:"resourceAttributes"`
		User string `json:"user"`
		Group []string `json:"group"`
	} `json:"spec"`
}