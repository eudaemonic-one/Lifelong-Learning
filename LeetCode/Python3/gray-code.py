class Solution:
    def grayCode(self, n: int) -> List[int]:
        ans = [0]
        for i in range(n):
            for x in reversed(ans):
                ans.append(x | 1 << i)
        return ans
