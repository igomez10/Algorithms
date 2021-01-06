package main

import (
	"strconv"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	cases := map[*TreeNode][]int{
		{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		}: {3, 1, 2},
		{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		}: {1, 2, 3},
	}
	i := 0
	for k, expected := range cases {
		i++
		t.Run("t"+strconv.Itoa(i), func(t *testing.T) {
			ans := preorderTraversal(k)
			if len(ans) != len(expected) {
				t.Fatal("unexpected length ", len(ans), len(expected))
			}

			for i := range ans {
				if ans[i] != expected[i] {
					t.Errorf("unexpected value in index %d, expected %d got %d", i, expected[i], ans[i])
				}
			}
		})
	}

}
