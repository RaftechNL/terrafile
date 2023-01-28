package providers

type ProviderIface interface {
	DownloadModule() error
}
