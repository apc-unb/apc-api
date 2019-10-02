package config

import (
	"context"
	"log"
	"os"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/togatoga/goforces"
)

const (
	mongoHost        = "mongo-host"
	mongoPort        = "mongo-port"
	port             = "port"
	logLevel         = "log-level"
	jwtKey           = "jwt-key"
	codeforcesKey    = "codeforces-key"
	codeforcesSecret = "codeforces-secret"
)

// Flags define the fields that will be passed via cmd
type Flags struct {
	Port             string
	MongoHost        string
	MongoPort        string
	LogLevel         string
	JwtSecret        string
	CodeforcesKey    string
	CodeforcesSecret string
}

// WebBuilder defines the parametric information of a whisper server instance
type WebBuilder struct {
	*Flags
	DataBase *mongo.Client
	GoForces *goforces.Client
}

// AddFlags adds flags for Builder.
func AddFlags(flags *pflag.FlagSet) {
	flags.StringP(port, "p", "8080", "[optional] Custom port for accessing Dragon T's services. Defaults to 8080")
	flags.StringP(mongoHost, "m", "localhost", "Custom host for accessing Mongo DB services. Defaults to localhost")
	flags.StringP(mongoPort, "t", "27017", "Custom port for accessing Mongo DB services. Defaults to 27017")
	flags.StringP(logLevel, "l", "info", "[optional] Sets the Log Level to one of seven (trace, debug, info, warn, error, fatal, panic). Defaults to info")
	flags.StringP(jwtKey, "k", "", "Sets the secret key used to hash the JWT")
	flags.StringP(codeforcesKey, "c", "", "Sets the secret key of Codeforces API")
	flags.StringP(codeforcesSecret, "s", "", "Sets the secret of Codeforces API")
}

// InitFromViper initializes the web server builder with properties retrieved from Viper.
func (b *WebBuilder) InitFromViper(v *viper.Viper) *WebBuilder {
	flags := new(Flags)
	flags.Port = v.GetString(port)
	flags.MongoHost = v.GetString(mongoHost)
	flags.MongoPort = v.GetString(mongoPort)
	flags.LogLevel = v.GetString(logLevel)
	flags.JwtSecret = v.GetString(jwtKey)
	flags.CodeforcesKey = v.GetString(codeforcesKey)
	flags.CodeforcesSecret = v.GetString(codeforcesSecret)

	flags.check()

	b.Flags = flags
	b.DataBase = b.getMongoDB(flags.MongoHost, flags.MongoPort)
	b.GoForces = b.getGoForces(flags.CodeforcesKey, flags.CodeforcesSecret)

	return b
}

func (flags *Flags) check() {
	logrus.Infof("Flags: '%v'", flags)
	if flags.JwtSecret == "" || flags.CodeforcesSecret == "" || flags.CodeforcesKey == "" {
		panic("secret-key, codeforces-key and codeforces-secret cannot be empty")
	}
}

func (b *WebBuilder) getMongoDB(host, port string) *mongo.Client {

	db, err := mongo.Connect(context.TODO(), "mongodb://"+host+":"+port)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	logrus.Infof("Connected to MongoDB!")

	return db
}

func (b *WebBuilder) getGoForces(codeforcesKey, codeforcesSecret string) *goforces.Client {

	var goForces *goforces.Client
	var err error

	goForces, err = goforces.NewClient(log.New(os.Stderr, "*** ", log.LstdFlags))

	if err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Infof("Connected to Codeforces API!")
	}

	goForces.SetAPIKey(codeforcesKey)
	goForces.SetAPISecret(codeforcesSecret)

	return goForces
}
