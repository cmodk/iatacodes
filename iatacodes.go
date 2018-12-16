package iatacodes

import (
	"github.com/cmodk/go-simplehttp"
	"github.com/sirupsen/logrus"
)

type IATACodes struct {
	lg    *logrus.Logger
	sh    simplehttp.SimpleHttp
	key   string
	debug bool
}

func New(k string, logger *logrus.Logger) *IATACodes {
	iatacodes := IATACodes{
		lg:  logger,
		sh:  simplehttp.New("https://iatacodes.org/api/", logger),
		key: k,
	}

	iatacodes.sh.AddHeader("Accept", "application/json")

	return &iatacodes

}

func (iatacodes *IATACodes) SetDebug(d bool) {
	iatacodes.debug = d
	iatacodes.sh.SetDebug(d)
}
