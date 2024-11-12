package models

import "github.com/OldBigBuddha/gqlgen/integration/server/remote_api"

type Viewer struct {
	User *remote_api.User
}
