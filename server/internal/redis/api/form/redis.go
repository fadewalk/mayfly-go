package form

type Redis struct {
	Id                 uint64
	Name               string `json:"name"`
	Host               string `binding:"required" json:"host"`
	Password           string `json:"password"`
	Mode               string `json:"mode"`
	Db                 string `json:"db"`
	SshTunnelMachineId int    `json:"sshTunnelMachineId"` // ssh隧道机器id
	TagId              uint64 `binding:"required" json:"tagId"`
	TagPath            string `json:"tagPath"`
	Remark             string `json:"remark"`
}

type Rename struct {
	Key    string `binding:"required" json:"key"`
	NewKey string `binding:"required" json:"newKey"`
}

type Expire struct {
	Key     string `binding:"required" json:"key"`
	Seconds int64  `binding:"required" json:"seconds"`
}

type KeyInfo struct {
	Key   string `binding:"required" json:"key"`
	Timed int64
}

type StringValue struct {
	KeyInfo
	Value interface{} `binding:"required" json:"value"`
}

type HashValue struct {
	KeyInfo
	Value []map[string]interface{} `binding:"required" json:"value"`
}

type SetValue struct {
	KeyInfo
	Value []interface{} `binding:"required" json:"value"`
}

type ListValue struct {
	KeyInfo
	Value []interface{} `binding:"required" json:"value"`
}

// list lset命令参数入参
type ListSetValue struct {
	Key   string `binding:"required" json:"key"`
	Index int64
	Value interface{} `binding:"required" json:"value"`
}

type RedisScanForm struct {
	Cursor map[string]uint64 `json:"cursor"`
	Match  string            `json:"match"`
	Count  int64             `json:"count"`
}

type ScanForm struct {
	Key    string `json:"key"`
	Cursor uint64 `json:"cursor"`
	Match  string `json:"match"`
	Count  int64  `json:"count"`
}

type SmemberOption struct {
	Key    string `json:"key"`
	Member any    `json:"member"`
}

type LRemOption struct {
	Key    string `json:"key"`
	Count  int64  `json:"count"`
	Member any    `json:"member"`
}

type ZAddOption struct {
	Key    string  `json:"key"`
	Score  float64 `json:"score"`
	Member any     `json:"member"`
}
