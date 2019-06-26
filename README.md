# GraphQL and Go, Best friends?

REST has served us so well, but it has started to get in the way. With the now ubiquitous presence of Single Page Apps the need for more flexible ways of fetching data from a server is really starting to be felt. The 1+N types of requests (where you fetch an array to do a separate fetch for each element of the array) are starting to be pretty common and have a huge impact on user experience causing long rendering time and wasting bandwidth with data not useful to most API clients. Is it really worth being RESTful? 

There are ways to fix this like using standards like the JSON API (https://jsonapi.org/) or others. But these standards or specifications are often hard to implement or are pushing developers towards more framework types of approach. Often letting developers implement their own flavor. But these days, you might not have to be RESTful anymore to have a simple and flexible API. GraphQL allows for flexibility, simplicity and safety.

In this talk, I will:
* Introduce GraphQL
* Design a simple API using GraphQL (with connections, mutations and queries)
* Explain how API consumers would leverage this GraphQL API
* Introduce https://github.com/99designs/gqlgen
* Explain the important pieces of implementing a GraphQL API in Go with gqlgen

Folders:
* `db`: Our database package used to focus only on GraphQL
* `experiement`: Folder in which live coding will happen
* `talk`: Source of the slides
* `oven`: Contains fallbacks when something really bad happens during live coding

## [Start experimenting](./experiment)

Lab steps:
1. Show Schema
1. Show gqlgen.yml
1. Gen code `make gen`
1. Show models
1. Show graph
1. Show resolver
1. Fix resolver
1. Run server `make run`
1. Support getting list of channels
1. Demo List channels
1. Support list messages
1. Demo messages of channels
1. Copy implemented resolver
1. Show Full query
1. Show number of SQL Queries
1. Do dataloader
1. Show number of queries
