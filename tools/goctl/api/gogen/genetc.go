package gogen

import (
	"fmt"
	"strconv"

	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
	"github.com/tal-tech/go-zero/tools/goctl/api/util"
	"github.com/tal-tech/go-zero/tools/goctl/config"
	"github.com/tal-tech/go-zero/tools/goctl/util/format"
)

const (
	defaultPort = 8888
	etcDir      = "etc"
	etcTemplate = `Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}
`
)

func genEtc(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, api.Service.Name)
	if err != nil {
		return err
	}

	service := api.Service
	host, ok := util.GetAnnotationValue(service.Groups[0].Annotations, "server", "host")
	if !ok {
		host = "0.0.0.0"
	}
	port, ok := util.GetAnnotationValue(service.Groups[0].Annotations, "server", "port")
	if !ok {
		port = strconv.Itoa(defaultPort)
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          etcDir,
		filename:        fmt.Sprintf("%s.yaml", filename),
		templateName:    "etcTemplate",
		category:        category,
		templateFile:    etcTemplateFile,
		builtinTemplate: etcTemplate,
		data: map[string]string{
			"serviceName": service.Name,
			"host":        host,
			"port":        port,
		},
	})
}
