package standalone

import (
	"github.com/Zenika/marcel/api"
	"github.com/Zenika/marcel/backoffice"
	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/frontend"
	"github.com/Zenika/marcel/httputil"
	"github.com/Zenika/marcel/module"
)

func Module() module.Module {
	var base = httputil.NormalizeBase(config.Default().HTTP().BasePath())

	return module.Module{
		Name: "Standalone",
		SubModules: []module.Module{
			api.Module(),
			backoffice.Module(),
			frontend.Module(),
		},
		HTTP: module.HTTP{
			BasePath: httputil.TrimTrailingSlash(base),
		},
	}
}
