package main

import (
	"strconv"
)

type QcWordStat struct {
	//质检词来源 0 顾客 1 客服
	Source int    `bson:"source" json:"source"`
	Word   string `bson:"word" json:"word"`
	Count  int    `bson:"count" json:"count"`
}

type Dialog struct {
	QcWord []*QcWordStat
}

// 按质检词，消息来源 合并 count
func Merge(dialog *Dialog , QcWordStats []*QcWordStat) {
	
	qcWordMergeMap := make(map[string]*QcWordStat)
	var str string	//存放拼接字符串

	for _, v := range dialog.QcWord {
		 str  = v.Word + " " + strconv.Itoa(v.Source) //拼接字符串，以空格隔开

		if qcWordMergeMap[str] == nil {
			qcWordMergeMap[str] = v
		} else {
			qcWordMergeMap[str].Count += v.Count//有重复，map中计算count
		}
	}

	for _,v := range QcWordStats {
		str =  v.Word + " " + strconv.Itoa(v.Source)

		if qcWordMergeMap[str] == nil {
			qcWordMergeMap[str] = v

			dialog.QcWord = append(dialog.QcWord, v)//添加到dialog

		} else {
			qcWordMergeMap[str].Count += v.Count //有重复，map中计算count
		}
	}


}

	/*
	//原代码
	func Merge(dialog []*QcWordStat, QcWordStats []*QcWordStat) []*QcWordStat {

	qcWordMergeMap := make(map[string]map[int]*QcWordStat)
	for _, v := range dialog {
		m := qcWordMergeMap[v.Word]

		if m == nil {
			m = map[int]*QcWordStat{}
			qcWordMergeMap[v.Word] = m
		}
		m[v.Source] = v
	}

	for _, v := range QcWordStats {
		m := qcWordMergeMap[v.Word]


		if m == nil {
			qcWordMergeMap[v.Word] = map[int]*QcWordStat{v.Source: v}
			continue
		}
		qcw := m[v.Source]
		if qcw == nil {
			m[v.Source] = v
			continue
		}
		qcw.Count += v.Count


	}

	dialog = dialog[0:0]
	for _, m := range qcWordMergeMap {
		for _, v := range m {
			dialog = append(dialog, v)
		}
	}
	return dialog
}
	*/
