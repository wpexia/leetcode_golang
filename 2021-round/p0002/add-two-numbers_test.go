package p0002

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "1", args: args{l1:buildListNode(342),l2:buildListNode(465)},want: buildListNode(807)},
		{name: "2", args: args{l1:buildListNode(0),l2:buildListNode(0)},want: buildListNode(0)},
		{name: "3", args: args{l1:buildListNode(9999999),l2:buildListNode(9999)},want: buildListNode(10009998)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildListNode(value int) *ListNode{
	var pre,root *ListNode
	for value!=0{
		node := &ListNode{Val: value % 10}
		if pre != nil{
			pre.Next = node
		} else {
			root = node
		}
		pre = node
		value = value/10
	}
	if root == nil{
		root = &ListNode{Val:0}
	}
	return root
}
