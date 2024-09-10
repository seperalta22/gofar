package gofar

const version = "0.1.0"

type Gofar struct {
	AppName string
	IsDebug bool // working in dev or prod?
	Version string
}

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
