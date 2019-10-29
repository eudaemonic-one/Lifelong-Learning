def bubble_sort(A):
    'BUBBLE-SORT(A)'
    for i in range(len(A)):
        for j in range(len(A) - 1, i, -1):
            print(j)
            if A[j] < A[j - 1]:
                A[j], A[j - 1] = A[j - 1], A[j]

if __name__ == '__main__':
    A = [1, 4, 7, 2, 5, 3, 8, 6, 0, 9]
    bubble_sort(A)
    print(A)