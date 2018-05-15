//+build prod

package constants

const (
	MongoDbName      = "bus-event"
	MongoDbUser      = "bus-event"
	MongoDbPassword  = "password"
	MongoDbPoolLimit = 10
)

var MongoDbHosts = []string{
	"172.16.0.10:27017",
	"172.16.0.11:27017",
	"172.16.0.11:27017",
}
