def merge(A, p, q, r):
    'MERGE(A, p, q, r)'
    n1 = q - p + 1
    n2 = r - q
    L, R = [None] * (n1 + 1), [None] * (n2 + 1)
    for i in range(n1):
        L[i] = A[p + i]
    for j in range(n2):
        R[j] = A[q + j + 1]
    import sys
    L[n1] = sys.maxsize
    R[n2] = sys.maxsize
    i, j = 0, 0
    for k in range(p, r+1):
        if L[i] <= R[j]:
            A[k] = L[i]
            i += 1
        else:
            A[k] = R[j]
            j += 1


def merge_sort(A, p, r):
    'MERGE-SORT(A, p, r)'
    if p < r:
        q = (p + r) // 2
        merge_sort(A, p, q)
        merge_sort(A, q + 1, r)
        merge(A, p, q, r)


if __name__ == '__main__':
    A = [1, 4, 7, 2, 5, 3, 8, 6, 0, 9]
    merge_sort(A, 0, 9)
    print(A)