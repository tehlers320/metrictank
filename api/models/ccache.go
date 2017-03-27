package models

type CCacheDelete struct {
	Patterns  []string `json:"patterns" form:"patterns" binding:"Required"`
	OrgId     int      `json:"orgId" form:"orgId" binding:"Required"`
	Propagate bool     `json:"propagate" form:"propagate"`
}

//go:generate msgp
type CCacheDeleteResp struct {
	PeerErrors      int `json:"peerErrors"`
	DeletedSeries   int `json:"deletedSeries"`
	DeletedArchives int `json:"deletedArchives"`
}
