package gofar

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const version = "0.1.0"

type Gofar struct {
	AppName  string
	IsDebug  bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	config   config
}

type config struct {
	port     string
	renderer string
}

// New reads the .env file, creates our application config, populates the Gofar type with settings
// based on .env values, and creates necessary folders and files if they don't exist
func (g *Gofar) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"handlers",
			"migrations",
			"views",
			"data",
			"public",
			"tmp",
			"logs",
			"middleware",
		},
	}
	err := g.Init(pathConfig)
	if err != nil {
		return err
	}

	// check if .env file exists
	err = g.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	//read .env file
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	// create loggers
	infoLog, errorLog := g.startLoggers()

	g.InfoLog = infoLog
	g.ErrorLog = errorLog
	g.IsDebug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	g.Version = version
	g.RootPath = rootPath

	g.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	return nil
}

func (g *Gofar) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// create folder if it doesn't exist
		err := g.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gofar) checkDotEnv(path string) error {
	err := g.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}

func (g *Gofar) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog, errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
