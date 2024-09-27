package lib

import (
	"log"
	"os"
)

type Writer interface {
	writeFile(*os.File) error
}

type ExtractClass struct {
	id          string
	class       string
	classValues []string
	states      []func(string) string
	style       func(*os.File, []string) error
}

func (e *ExtractClass) writeFile(file *os.File) error {
	for _, state := range e.states {
		e.class = state(e.class)
	}

	_, err := file.WriteString(e.class + "{")

	if err != nil {
		log.Fatal(err)
	}

	err = e.style(file, e.classValues)

	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("}")

	if err != nil {
		log.Fatal(err)
	}

	return err
}

type MediaTree struct {
	classes     []*ExtractClass
	mediaValues []string
	media       func(*os.File, []string) error
	parent      *MediaTree
	child       *MediaTree
}

func (m *MediaTree) writeFile(file *os.File) error {
	err := m.media(file, m.mediaValues)

	if err != nil {
		log.Fatal(err)
	}

	for _, class := range m.classes {
		err = class.writeFile(file)
	}

	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("}")

	if err != nil {
		log.Fatal(err)
	}

	return err
}
