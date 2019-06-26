{{ reserveImport "context"  }}
{{ reserveImport "fmt"  }}
{{ reserveImport "io"  }}
{{ reserveImport "strconv"  }}
{{ reserveImport "time"  }}
{{ reserveImport "sync"  }}
{{ reserveImport "errors"  }}
{{ reserveImport "bytes"  }}

{{ reserveImport "github.com/99designs/gqlgen/handler" }}
{{ reserveImport "github.com/vektah/gqlparser" }}
{{ reserveImport "github.com/vektah/gqlparser/ast" }}
{{ reserveImport "github.com/99designs/gqlgen/graphql" }}
{{ reserveImport "github.com/99designs/gqlgen/graphql/introspection" }}

type {{$.Object.Name}}Resolver struct { *{{$.ResolverType}} }

{{ range $field := $.Object.Fields -}}
  {{- if $field.IsResolver -}}
  func (r *{{$.Object.Name}}Resolver) {{$field.GoFieldName}}{{ $field.ShortResolverDeclaration }} {
    panic("not implemented")
  }
  {{ end -}}
{{ end -}}
