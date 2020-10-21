package model

type QcWordStat struct {
	//质检词来源 0 顾客 1 客服
	Source int    `bson:"source" json:"source"`
	Word   string `bson:"word" json:"word"`
	Count  int    `bson:"count" json:"count"`
}

type Dialog struct {
	QcWord []*QcWordStat
}
