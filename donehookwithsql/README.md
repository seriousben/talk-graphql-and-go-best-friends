
* `export GO111MODULE=on`
* `vi server/server.go`
* `vi resolver.go`
* ```gql
{
  channels{
    id
    name
    createdBy { id, name }
  }
}
```