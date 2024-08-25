package db

var Db map[string]any

func init() {
	Db = make(map[string]any)
}
