class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if not s or not t:
            return ""
        slow, fast = 0, 0
        need = collections.Counter(t)
        counts = {}
        required, formed = len(need), 0
        ans = float("inf"), None, None
        while fast < len(s):
            ch = s[fast]
            try:
                counts[ch] += 1
            except KeyError:
                counts[ch] = 1
            if ch in need and need[ch] == counts[ch]:
                formed += 1
            while slow <= fast and formed == required:
                c = s[slow]
                if fast - slow + 1 < ans[0]:
                    ans = (fast-slow+1, slow, fast)
                counts[c] -= 1
                if c in need and counts[c] < need[c]:
                    formed -= 1
                slow += 1
            fast += 1
        return "" if ans[0] == float("inf") else s[ans[1]:ans[2]+1]
