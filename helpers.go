package gofar

import "os"

func (g *Gofar) CreateDirIfNotExists(path string) error {
	const mode = 0755 // read, write, execute for owner, read and execute for group and others
	if _, err := os.Stat(path); os.IsNotExist(err) {
		//if err := os.MkdirAll(path, mode); err != nil {
		//	return err
		//}
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gofar) CreateFileIfNotExists(path string) error {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	return nil
}
