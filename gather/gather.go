package gather

import (
	"github.com/ukitazume/ey-berks/config"
)

type Gather struct {
	Berks Berks
}

func NewGather(berksFilePath) Gather {
	bersk := config.Parse(berksFilePath)
	return &Gather{Berks: bersk}
}

func (g *Gather) Gather() error {
	err := ok.Mkdir("~/.ey-berks")
}
