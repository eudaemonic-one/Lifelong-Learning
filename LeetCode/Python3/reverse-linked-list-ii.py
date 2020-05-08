# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        if not head:
            return None
        curr, prev = head, None
        while m > 1:
            prev = curr
            curr = curr.next
            m -= 1
            n -= 1
        left_last, mid = prev, curr
        while n:
            tmp = curr.next
            curr.next = prev
            prev = curr
            curr = tmp
            n -= 1
        if left_last:
            left_last.next = prev
        else:
            head = prev
        mid.next = curr
        return head
