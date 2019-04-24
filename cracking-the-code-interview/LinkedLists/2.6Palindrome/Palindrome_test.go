package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	head := LLFromArray(input)
	shouldBeTrue := IsPalindrome(head)
	if !shouldBeTrue {
		t.Errorf("Expected true for %+v as a palindrome - got false", input)

	}

	secondInput := []int{1, 2, 3, 4, 5, 6, 7}
	secondHead := LLFromArray(secondInput)
	shouldBeFalse := IsPalindrome(secondHead)
	if shouldBeFalse {
		t.Errorf("Expected false for %+v as a palindrome - got true", secondInput)
	}

	// 1 2 3 4 5 4 3 2 1
	// 5 1 2 3 4 1 2 3 4

}

func TestLLFromArray(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	head := LLFromArray(input)
	for i := range input {
		if input[i] != head.value {
			t.Errorf("expected %d got %d", input[i], head.value)
		}
		head = head.next
	}
}
