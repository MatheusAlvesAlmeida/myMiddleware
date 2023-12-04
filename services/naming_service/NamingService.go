package naming

import (
	clientproxy "github.com/MatheusAlvesAlmeida/myMiddleware/distribution/client_proxy"
)

type NamingService struct {
	Repository map[string]clientproxy.PercentageProxy
}

func (namingService *NamingService) Register(name string, proxy clientproxy.PercentageProxy) bool {
	response := false

	if namingService.Repository == nil {
		namingService.Repository = make(map[string]clientproxy.PercentageProxy)
	}

	if _, ok := namingService.Repository[name]; !ok {
		namingService.Repository[name] = proxy
		response = true
	}

	return response
}

func (namingService *NamingService) Lookup(name string) clientproxy.PercentageProxy {
	return namingService.Repository[name]
}

func (namingService *NamingService) List() []string {
	var names []string
	for name := range namingService.Repository {
		names = append(names, name)
	}
	return names
}
