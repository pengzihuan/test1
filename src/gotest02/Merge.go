package main

import "awesomeProject/pkg/model"

//type QcWordStat struct {
//	//质检词来源 0 顾客 1 客服
//	Source int    `bson:"source" json:"source"`
//	Word   string `bson:"word" json:"word"`
//	Count  int    `bson:"count" json:"count"`
//}
//
//type Dialog struct {
//	QcWord []*QcWordStat
//}


//按质检词，消息来源 合并 count
func Merge(dialog *model.Dialog, QcWordStats []*model.QcWordStat) {
	type Key struct {
		Source int    `bson:"source" json:"source"`
		Word   string `bson:"word" json:"word"`
	}
	mp := make(map[Key]*model.QcWordStat)
	for _, v := range dialog.QcWord {
		k := Key{Source: v.Source, Word: v.Word}
		if ptr, ok := mp[k]; ok {
			ptr.Count += v.Count
		} else {
			mp[k] = v
		}
	}

	for _, v := range QcWordStats {
		k := Key{Source: v.Source, Word: v.Word}
		if ptr, ok := mp[k]; ok {
			ptr.Count += v.Count
		} else {
			mp[k] = v
			dialog.QcWord = append(dialog.QcWord, v)
		}
	}
}


//func Merge(dialog *Dialog, QcWordStats []*QcWordStat) {
//	type Key struct {
//		Source int    `bson:"source" json:"source"`
//		Word   string `bson:"word" json:"word"`
//	}
//	mp := make(map[Key]*QcWordStat)
//	for _, v := range dialog.QcWord {
//		k := Key{Source: v.Source, Word: v.Word}
//		if ptr, ok := mp[k]; ok {
//			ptr.Count += v.Count
//		} else {
//			mp[k] = v
//		}
//	}
//
//	for _, v := range QcWordStats {
//		k := Key{Source: v.Source, Word: v.Word}
//		if ptr, ok := mp[k]; ok {
//			ptr.Count += v.Count
//		} else {
//			mp[k] = v
//			dialog.QcWord = append(dialog.QcWord, v)
//		}
//	}
//}

