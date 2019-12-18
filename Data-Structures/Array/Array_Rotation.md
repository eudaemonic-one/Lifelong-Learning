# Array Rotation

Write a function rotate(ar[], d, n) that rotates arr[] of size n by d elements.

## METHOD 1 (Using temp array)

    Store d elements in a temp array
    Shift rest of the arr[]
    Store back the d elements

## METHOD 2 (Rotate one by one)

    leftRotate(arr[], d, n)
    start
    For i = 0 to i < d
        Left rotate all elements of arr[] by one
    end

## METHOD 3 (A Juggling Algorithm)

Instead of moving one by one, divide the array in different sets where number of sets is equal to GCD of n and d and move the elements within sets. For eacn i in the range of GCD of n and d, we just start with temp = arr[i] and keep moving arr[I+d] to arr[I] and finally store temp at the right place.

## Method 4 (The Reversal Algorithm)

    rotate(arr[], d, n)
        reverse(arr[], 1, d) ;
        reverse(arr[], d + 1, n);
        reverse(arr[], 1, n);
