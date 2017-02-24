package models

type CCacheDelete struct {
	Patterns []string `json:"patterns" form:"patterns" binding:"Required"`
	OrgId    int      `json:"orgId" form:"orgId" binding:"Required"`
}

//go:generate msgp
type CCacheDeleteResp struct {
	Nodes []string
}

func NewCCacheDeleteResp() *CCacheDeleteResp {
	return &CCacheDeleteResp{
		Nodes: make([]string, 0),
	}
}
