class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        def backtrack(addr, comb, path, idx):
            if idx == 4:
                if not addr:
                    comb.append(path[:-1])
                return
            for i in range(1, 4):
                if i <= len(addr):
                    if int(addr[:i]) <= 255:
                        backtrack(addr[i:], comb, path+addr[:i]+".", idx+1)
                    if addr[0] == "0":
                        break
        comb = []
        backtrack(s, comb, "", 0)
        return comb
