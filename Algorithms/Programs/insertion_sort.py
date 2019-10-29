def insertion_sort(A):
    'INSERTION-SORT(A)'
    for j in range(1, len(A)):
        key = A[j]
        # Insert A[j] into the sorted sequence A[1..j-1].
        i = j - 1
        while i >= 0 and A[i] > key:
            A[i+1] = A[i]
            i -= 1
        A[i+1] = key
    return A


if __name__ == '__main__':
    A = [1, 4, 7, 2, 5, 3, 8, 6, 0, 9]
    insertion_sort(A)
    print(A)