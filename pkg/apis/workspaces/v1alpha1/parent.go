package v1alpha1

type Parent struct {
	ImportReference `json:",inline"`
	ParentOverrides       `json:",inline"`
}
