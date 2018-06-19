package main

import (
        "testing"
)

func TestGuess(t *testing.T) {
        cases := []struct{
                secret int
                deck []int
                want int
        }{
    {0, []int{}, -1},
    {0, []int{2}, 2},
    {1, []int{2}, 2},
    {1, []int{2,3,4}, 4},
    {2, []int{2,3,4}, 3},
    {2, []int{2,3,4,5,6}, 6},
    {1, []int{0, 1, 2, 3, 4, 5}, 5},
    {1, []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
    {1, []int{11, 1, 2, 3, 4, 5}, 5},
    {1, []int{12, 1, 2, 3, 4, 5}, 5},
    {1, []int{12, 1, 2, 3, 4, 5, 13, 14, 15, 16, 17}, 17},
        }
        for _, c := range cases {
                got := guess(c.secret, c.deck)
                if got != c.want {
                        t.Errorf("guess(%v, %v)=%v; want %v", c.secret, c.deck, got, c.want)
                }
        }
}
