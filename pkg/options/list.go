package options

type ListOptions struct {
	DryRun bool `json:"dry_run" yaml:"dry_run"`
	Force  bool `json:"force" yaml:"force"`

	Offset int `json:"offset" yaml:"offset"`
	Limit  int `json:"limit" yaml:"limit"`

	SortField string `json:"sort_field" yaml:"sort_field"`
	SortOrder string `json:"sort_order" yaml:"sort_order"`
}
