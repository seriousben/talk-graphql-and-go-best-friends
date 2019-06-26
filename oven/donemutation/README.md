* `vi gqlgen.yml`
* `export GO111MODULE=on`
* `go run github.com/99designs/gqlgen`

```gql
query getChannelsForFrontpage{
  channels {
    name
    createdBy {
      name
    }
    messages {
      text
      createdBy {
        name
      }
      channel {
        name
      }
    }
  }
}
```

```gql
mutation createMsg($message: NewMessage!) {
  createMessage(input: $message) {
    id
    text
    channel {
      name
      createdBy {
        name
      }
    }
  }
}
```

```gql
{
  "message": {
    "channelId": 3,
    "text": "testing"
  }
}
```