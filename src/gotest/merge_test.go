package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		dialog      Dialog
		QcWordStats []*QcWordStat
	}
	tests := []struct {
		name string
		args args
		want []*QcWordStat
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
				dialog:      Dialog{
					QcWord: []*QcWordStat{
						&QcWordStat{0,"one",1},
						&QcWordStat{1,"one",1},
						&QcWordStat{0,"two",1},
					},
				},
				QcWordStats: []*QcWordStat{
					&QcWordStat{1,"one",2},
					&QcWordStat{0,"two",1},
					&QcWordStat{1,"two",1},
				},
			},
			want: []*QcWordStat{
				&QcWordStat{0,"one",1},
				&QcWordStat{1,"one",3},
				&QcWordStat{0,"two",2},
				&QcWordStat{1,"two",1},
			},
		},

		{	//2
			name:"merge_test_2",
			args: args{
				dialog:      Dialog{
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
				dialog: Dialog{
					QcWord:nil,
				},
				QcWordStats: []*QcWordStat{
					&QcWordStat{1,"one",2},
					&QcWordStat{0,"two",1},
					&QcWordStat{1,"two",1},
				},
			},
			want: []*QcWordStat{
				&QcWordStat{1,"one",2},
				&QcWordStat{0,"two",1},
				&QcWordStat{1,"two",1},
			},
		},

		{
			//4
			name: "merge_test_4",
			args: args{
				dialog: Dialog{
					QcWord: []*QcWordStat{
						&QcWordStat{0,"one",1},
						&QcWordStat{1,"one",1},
						&QcWordStat{0,"two",1},
					},
				},
				QcWordStats: nil,
			},
			want: []*QcWordStat{
				&QcWordStat{0,"one",1},
				&QcWordStat{1,"one",1},
				&QcWordStat{0,"two",1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//got1 := Merge(tt.args.dialog, tt.args.QcWordStats)
			Merge(&tt.args.dialog, tt.args.QcWordStats)
			got := tt.args.dialog.QcWord

			b :=reflect.DeepEqual(got, tt.want)

			gotjson,_ :=json.Marshal(got)
			wantjson,_ :=json.Marshal(tt.want)
			if !b{
				fmt.Println(string(gotjson))
				fmt.Println(string(wantjson))
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
/*
package main

import (

	"reflect"
	"testing"
)

//有返回值的2
func TestMerge1(t *testing.T) {
	type args struct {
		dialog      []*QcWordStat
		QcWordStats []*QcWordStat
	}



	tests := []struct {
		name string
		args args
		want []*QcWordStat
	}{
		// TODO: Add test cases.

			//4个测试用例：
			//1、dialog：数据	QcWordStats：数据
			//2、dialog：nil	QcWordStats：nil
			//3、dialog：nil	QcWordStats：数据
			//4、dialog：nil	QcWordStats：nil
			//
		{	//1、dialog和QcWordStats都有值
			name:"merge_test_1",
			args:args{ dialog: []*QcWordStat{
				&QcWordStat{0,"one",1},
				&QcWordStat{1,"one",1},
				&QcWordStat{0,"two",1},
			},
				QcWordStats: []*QcWordStat{
					&QcWordStat{1,"one",2},
					&QcWordStat{0,"two",1},
					&QcWordStat{1,"two",1},
				}},
			want: []*QcWordStat{
				&QcWordStat{0,"one",1},
				&QcWordStat{1,"one",3},
				&QcWordStat{0,"two",2},
				&QcWordStat{1,"two",1},
			}},

			{//2、dialog和QcWordStats都为nil
			name: "merge_test_2",
			args: args{dialog: nil,QcWordStats: nil},
			want:nil },

			{//3、dialog为nil和QcWordStats有值
			name: "merge_test_3",
			args:args{ dialog: nil,
				QcWordStats: []*QcWordStat{
					&QcWordStat{1,"one",2},
					&QcWordStat{0,"two",1},
					&QcWordStat{1,"two",1},
				}},
			want: []*QcWordStat{
				&QcWordStat{1,"one",2},
				&QcWordStat{0,"two",1},
				&QcWordStat{1,"two",1},
			},
			},

			{//4、dialog有值和QcWordStats为nil
			name: "merge_test_4",
			args: args{ dialog: []*QcWordStat{
				&QcWordStat{0,"one",1},
				&QcWordStat{1,"one",1},
				&QcWordStat{0,"two",1},
				},
				QcWordStats: nil,
			},
			want: []*QcWordStat{
				&QcWordStat{0,"one",1},
				&QcWordStat{1,"one",1},
				&QcWordStat{0,"two",1},
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := Merge(tt.args.dialog, tt.args.QcWordStats)

			flag := 0
			b := false

			if  got == nil && tt.want == nil{ //为nil情况
				b = true
			} else if got!=nil && tt.want!=nil && len(got) == len(tt.want)  {//dialog和QcWordStats至少有一方值

				//查重got
				qcWordMergeMap := make(map[string]map[int]*QcWordStat)
				for _,v := range got {
					m := qcWordMergeMap[v.Word]
					if m == nil {
						m = map[int]*QcWordStat{}
						qcWordMergeMap[v.Word] = m
						m[v.Source] = v
						continue
					}

					if m[v.Source] == nil {
						m[v.Source] = v
						continue
					}

					t.Errorf("got repeat ,got = %v", got)
				}

				//查重want
				qcWordMergeMap1 := make(map[string]map[int]*QcWordStat)
				for _,v := range tt.want {
					m := qcWordMergeMap1[v.Word]
					if m == nil {
						m = map[int]*QcWordStat{}
						qcWordMergeMap1[v.Word] = m
						m[v.Source] = v
						continue
					}

					if m[v.Source] == nil {
						m[v.Source] = v
						continue
					}

					t.Errorf("want repeat ,want = %v", tt.want)

				}

				//比较got和want是否相等
				for _,date1 := range got {
					for _,date2 := range tt.want {
						if (*date1) == (*date2) {
							flag++
						}
					}
				}
				if flag == len(tt.want) {
					b = true
				}
			}

			//gotjson,_ :=json.Marshal(got)
			//wantjson,_ :=json.Marshal(tt.want)

			b :=reflect.DeepEqual(got, tt.want)

			if !b{
				//fmt.Println(string(gotjson))
				//fmt.Println(string(wantjson))
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

 //有返回值的1
func TestMerge(t *testing.T) {
	qcWordStat01 := QcWordStat{0,"one",1}
	qcWordStat02 := QcWordStat{1,"one",1}
	qcWordStat03 := QcWordStat{0,"two",1}
	dialog := []*QcWordStat{&qcWordStat01,&qcWordStat02,&qcWordStat03}

	//测试用例[]*QcWordStats
	qcWordStat04 := QcWordStat{1,"one",2}
	qcWordStat05 := QcWordStat{0,"two",1}
	qcWordStat06 := QcWordStat{1,"two",1}
	qcw := []*QcWordStat{&qcWordStat04,&qcWordStat05,&qcWordStat06}

	//结果
	qcWordStat1 := QcWordStat{0,"one",1}
	qcWordStat2 := QcWordStat{1,"one",3}
	qcWordStat3 := QcWordStat{0,"two",2}
	qcWordStat4 := QcWordStat{1,"two",1}
	want := []*QcWordStat{&qcWordStat1,&qcWordStat2,&qcWordStat3,&qcWordStat4}

	got :=Merge(dialog,qcw)

	gotjson,_ :=json.Marshal(got)
	wantjson,_ :=json.Marshal(want)

	b :=reflect.DeepEqual(gotjson, wantjson)

	if !b {
		fmt.Println(string(gotjson))
		fmt.Println(string(wantjson))
		t.Errorf("Merge(dialog,qcw) = false\n got:%v\n want:%v", got, want)
	}
}
*/
