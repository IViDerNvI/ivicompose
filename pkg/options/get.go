package options

type GetOptions struct {
	DryRun bool `json:"dry_run" yaml:"dry_run"`
	Force  bool `json:"force" yaml:"force"`
}
