package main

func isP0wer22(n int) bool {
  return n > 0 && n &(n-1) == 0
}