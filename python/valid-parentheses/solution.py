class Solution:

    def isValid(self, s: str) -> bool:
        num_pairs = len(s)

        for i in range(num_pairs):
            s = s.replace("{}", "").replace("()", "").replace("[]", "")

        if not s:
            return True
        else:
            return False
