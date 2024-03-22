class Solution:

    def romanToInt(self, s: str) -> int:
        ans = 0
        roman_value = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000
        }

        for i, j in zip(s, s[1:]):
            if roman_value[i] < roman_value[j]:
                ans -= roman_value[i]
            else:
                ans += roman_value[i]
        return ans + roman_value[s[-1]]
