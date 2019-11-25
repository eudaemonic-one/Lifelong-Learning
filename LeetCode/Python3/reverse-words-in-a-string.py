class Solution:
    def reverseWords(self, s: str) -> str:
        word = ""
        stack = []
        for c in s:
            if c != ' ':
                word += c
            else:
                if word != "":
                    stack.append(word)
                word = ""
        if word != "":
            stack.append(word)
        return " ".join(reversed(stack))
