package v1alpha1

type PluginComponent struct {
	BaseComponent   `json:",inline"`
	ImportReference `json:",inline"`
	PluginOverrides `json:",inline"`
}
