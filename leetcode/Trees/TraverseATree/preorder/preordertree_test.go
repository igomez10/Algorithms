package main

import (
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	ans := preorderTraversal(root)
	expected := []int{3, 1, 2}
	if len(ans) != len(expected) {
		t.Fatal("unexpected length ", len(ans), len(expected))
	}

	for i := range ans {
		if ans[i] != expected[i] {
			t.Errorf("unexpected value in index %d, expected %d got %d", i, expected[i], ans[i])
		}
	}

}
