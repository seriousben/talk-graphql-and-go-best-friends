package donedataloader

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/seriousben/talk-graphql/db"
)

const userLoaderKey = "userloader"
const channelLoaderKey = "channelloader"

func DataloaderMiddleware(db *db.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     10 * time.Millisecond,
			fetch: func(keys []string) ([]*User, []error) {
				dbIds := make([]int, len(keys))
				for i := 0; i < len(keys); i++ {
					id, err := strconv.Atoi(keys[i])
					if err != nil {
						panic(err)
					}

					dbIds[i] = id
				}

				us, err := db.ListUsers(r.Context(), dbIds)
				if err != nil {
					panic(err)
				}

				objByKey := map[string]*User{}
				for _, u := range us {
					gu := User{
						ID:   strconv.Itoa(u.ID),
						Name: u.Name,
					}
					objByKey[gu.ID] = &gu
				}
				gus := make([]*User, len(us))
				for i := 0; i < len(keys); i++ {
					gus[i] = objByKey[keys[i]]
				}

				return gus, nil
			},
		}

		channelloader := ChannelLoader{
			maxBatch: 100,
			wait:     10 * time.Millisecond,
			fetch: func(keys []string) ([]*Channel, []error) {
				dbIds := make([]int, len(keys))
				for i := 0; i < len(keys); i++ {
					id, err := strconv.Atoi(keys[i])
					if err != nil {
						panic(err)
					}

					dbIds[i] = id
				}

				cs, err := db.ListChannelsByID(r.Context(), dbIds)
				if err != nil {
					panic(err)
				}

				objByKey := map[string]*Channel{}
				for _, c := range cs {
					gc := Channel{
						ID:          strconv.Itoa(c.ID),
						Name:        c.Name,
						CreatedByID: strconv.Itoa(c.CreatedByID),
					}
					objByKey[gc.ID] = &gc
				}
				gcs := make([]*Channel, len(cs))
				for i := 0; i < len(keys); i++ {
					gcs[i] = objByKey[keys[i]]
				}

				return gcs, nil
			},
		}

		ctx := context.WithValue(r.Context(), userLoaderKey, &userloader)
		ctx = context.WithValue(ctx, channelLoaderKey, &channelloader)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
