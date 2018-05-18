package token_collector

import (
	"time"

	"github.com/IhorBondartsov/OLX_Parser/userMS/storage"
	"github.com/IhorBondartsov/OLX_Parser/userms/entities"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// NewCollectorForTokenStor - make new collector
func NewCollectorForTokenStor(cfg CollectorCfg) *collector {
	return &collector{
		CheckInterval: cfg.CheckInterval,
		ExpiredTime:   cfg.ExpiredTime,
		RefreshStor:   cfg.RefreshStor,
	}
}

// CollectorCfg - config for collector
type CollectorCfg struct {
	CheckInterval int //UNIX time
	RefreshStor   storage.RefreshToken
	ExpiredTime   int64
}

// collector - delete all token which was expired from database
type collector struct {
	CheckInterval int //UNIX time
	RefreshStor   storage.RefreshToken
	ExpiredTime   int64
}

// Start -
func (c *collector) Start() {
	t := time.Duration(c.CheckInterval) * time.Second
	ticker := time.NewTicker(t)
	for {
		select {
		case <-ticker.C:
			now := time.Now().Unix()
			from := now - c.ExpiredTime
			models, err := c.RefreshStor.GetTokenByRange(from, now)

			if err != nil {
				log.Errorf("[Collector][Start] Error with db %v", err)
			}
			c.deleteTokens(models)
		}
	}
}

func (c *collector) deleteTokens(t []entities.Token) {
	for _, v := range t {
		if err := c.RefreshStor.DeleteToken(v); err != nil {
			log.Errorf("[Collector][DeleteToken] Error with db %v", err)
		}
	}
}
