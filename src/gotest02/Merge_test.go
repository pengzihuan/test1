package main

import (
	"awesomeProject/pkg/model"
	"reflect"
	"testing"
)


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


func TestMerge(t *testing.T) {
	type args struct {
		dialog      model.Dialog
		QcWordStats []*model.QcWordStat
	}
	tests := []struct {
		name string
		args args
		want []*model.QcWordStat
	}{
		// TODO: Add test cases.
		/*
			4个测试用例：
			1、dialog：数据	QcWordStats：数据
			2、dialog：nil	QcWordStats：nil
			3、dialog：nil	QcWordStats：数据
			4、dialog：数据	QcWordStats：nil
		*/
		{	//1
			name:"merge_test_1",
			args: args{
				dialog:      model.Dialog{
					QcWord: []*model.QcWordStat{
						&model.QcWordStat{0,"one",1},
						&model.QcWordStat{1,"one",1},
						&model.QcWordStat{0,"two",1},
					},
				},
				QcWordStats: []*model.QcWordStat{
					&model.QcWordStat{1,"one",2},
					&model.QcWordStat{0,"two",1},
					&model.QcWordStat{1,"two",1},
				},
			},
			want: []*model.QcWordStat{
				&model.QcWordStat{0,"one",1},
				&model.QcWordStat{1,"one",3},
				&model.QcWordStat{0,"two",2},
				&model.QcWordStat{1,"two",1},
			},
		},

		{	//2
			name:"merge_test_2",
			args: args{
				dialog:      model.Dialog{
					QcWord:nil,
				},
				QcWordStats: nil,
			},
			want: nil,
		},

		{
			//3
			name: "merge_test_3",
			args: args{
				dialog: model.Dialog{
					QcWord:nil,
				},
				QcWordStats: []*model.QcWordStat{
					&model.QcWordStat{1,"one",2},
					&model.QcWordStat{0,"two",1},
					&model.QcWordStat{1,"two",1},
				},
			},
			want: []*model.QcWordStat{
				&model.QcWordStat{1,"one",2},
				&model.QcWordStat{0,"two",1},
				&model.QcWordStat{1,"two",1},
			},
		},

		{
			//4
			name: "merge_test_4",
			args: args{
				dialog: model.Dialog{
					QcWord: []*model.QcWordStat{
						&model.QcWordStat{0,"one",1},
						&model.QcWordStat{1,"one",1},
						&model.QcWordStat{0,"two",1},
					},
				},
				QcWordStats: nil,
			},
			want: []*model.QcWordStat{
				&model.QcWordStat{0,"one",1},
				&model.QcWordStat{1,"one",1},
				&model.QcWordStat{0,"two",1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

				Merge(&(tt.args.dialog), tt.args.QcWordStats)
				got := tt.args.dialog.QcWord
				b :=reflect.DeepEqual(got, tt.want)

				if !b{
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
