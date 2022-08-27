package db

type IdKeyObject struct {
	Id  string `json:"id"`
	Key string `json:"key"`
}

type PartitionResponse struct {
	TotalRows int           `json:"total_rows"`
	Offset    int           `json:"offset"`
	Rows      []IdKeyObject `json:"rows,omitempty"`
}

type EntityKeyObject struct {
	CompositeKey string
	Entity       string
	Id           string
}
