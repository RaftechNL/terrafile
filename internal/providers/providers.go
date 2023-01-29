package providers

// ModuleSpec is used to deliver details about the module to be downloaded into the provider
type ModuleSpec struct {
	Source  string `yaml:"source"`
	Version string `yaml:"version"`
}

type ProviderIface interface {
	DownloadModule(moduleSpec ModuleSpec, outputPath string) error
}
