class Solution:
    def climbStairs(self, n: int) -> int:
        fib1, fib2 = 0, 1

        for i in range(n):
            res = fib1 + fib2
            fib1 = fib2
            fib2 = res
        return res