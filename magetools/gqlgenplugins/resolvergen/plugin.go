package resolvergen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"
)

func New() plugin.Plugin {
	return &Plugin{}
}

type Plugin struct{}

var _ plugin.CodeGenerator = &Plugin{}

func (m *Plugin) Name() string {
	return "improved-resolvergen"
}

func (m *Plugin) GenerateCode(data *codegen.Data) error {
	if !data.Config.Resolver.IsDefined() {
		return nil
	}

	resolverBuild := &ResolverBuild{
		Data:         data,
		PackageName:  data.Config.Resolver.Package,
		ResolverType: data.Config.Resolver.Type,
	}

	err := templates.Render(templates.Options{
		PackageName: data.Config.Resolver.Package,
		Filename:    data.Config.Resolver.Filename,
		Data:        resolverBuild,
	})
	if err != nil {
		return err
	}

	for _, o := range data.Objects {
		if !o.HasResolvers() {
			continue
		}
		objectFilename := filepath.Join(filepath.Dir(data.Config.Resolver.Filename), strings.ToLower(o.Name)+".go")

		if _, err := os.Stat(objectFilename); err == nil {
			fmt.Printf("Skip generation of %s, file already exist\n", objectFilename)
			continue
		}

		_, callerFile, _, _ := runtime.Caller(0)
		tplPath := filepath.Join(filepath.Dir(callerFile), "objectresolver.tpl")

		tmpl, err := ioutil.ReadFile(tplPath)
		if err != nil {
			return err
		}

		err = templates.Render(templates.Options{
			PackageName: data.Config.Resolver.Package,
			Filename:    objectFilename,
			Template:    string(tmpl),
			Data: &ObjectResolverBuild{
				ResolverBuild: resolverBuild,
				Object:        o,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

type ResolverBuild struct {
	*codegen.Data

	PackageName  string
	ResolverType string
}

type ObjectResolverBuild struct {
	*ResolverBuild

	Object *codegen.Object
}
