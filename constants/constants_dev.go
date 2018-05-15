//+build dev

package constants

const (
	MongoDbName      = "go_app"
	MongoDbUser      = "go_app"
	MongoDbPassword  = "GoAppPass123"
	MongoDbPoolLimit = 10

)

var MongoDbHosts = []string{
	"192.168.5.237:27017",
}
