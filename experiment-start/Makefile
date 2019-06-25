gen:
	export GO111MODULE=on
	go run github.com/99designs/gqlgen
	go run github.com/vektah/dataloaden UserLoader string "*github.com/seriousben/talk-graphql/donedataloader.User"
	go run github.com/vektah/dataloaden ChannelLoader string "*github.com/seriousben/talk-graphql/donedataloader.Channel"

clean:
	rm generated.go models_gen.go resolver.go channelloader_gen.go userloader_gen.go

run:
	go run server/server.go
