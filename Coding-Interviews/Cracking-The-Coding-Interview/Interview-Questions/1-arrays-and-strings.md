# Arrays and Strings

## Hash Tables

* A hash table is a data structure that maps keys to values for highly efficient lookup
* First, compute the key's hash code, which will usually be an `int` or `long`
  * Note that two different keys could have the same hash code, as there may be an infinite number of keys and a finite number of `int`s
* Then, map the hash code to an index in the array
  * This could be done with something like `hash(key) % array_length`
  * Two different hash codes could, of course, map to the same index
* At this index, there is a linked list of keys and values
  * Store the key and value in this index
  * We must use a linked list because of collisions: you could have two different keys with the same hash code, or two different hash codes that map to the same index
* To retrieve the value pair by its key, you repeat this process

## ArrayList & Resizable Arrays

* When you need an array-like data structure that offers dynamic resizing, you would usually use an Arraylist
* An Arraylist is an array that resizes itself as needed while still providing $O(1)$ access
* A typical implementa­tion is that when the array is full, the array doubles in size
* Each doubling takes $O(n)$ time, but happens so rarely that its amortized insertion time is still $O(1)$

### Amortized Insertion Runtime O(1)

``` text
final capacity increase: n/2 elements to copy
previous capacity increase: n/4 elements to copy
previous capacity increase: n/8 elements to copy
previous capacity increase: n/16 elements to copy
...
second capacity increase: 2 elements to copy
first capacity increase: 1 element to copy
```

* Therefore, inserting $N$ elements takes $O(N)$ work total
* Each insertion is $O(1)$ on average, even though some insertions take $O(N)$ time in the worst case

## StringBuilder

* `StringBuilder` simply creates a resizable array of all the strings, copying them back to a string only when necessary

## Interview Questions

* **1.1 Is Unique:**
  * Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
* **1.2 Check Permutation:**
  * Given two strings,write a method to decide if one is a permutation of the other.
* **1.3 URLify:**
  * Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at the end to hold the additional characters,and that you are given the "true" length of the string. (Note: If implementing in Java,please use a character array so that you can perform this operation in place.)
  * EXAMPLE
    * Input: "Mr John Smith ", 13
    * Output: "Mr%20John%20Smith"
* **1.4 Palindrome Permutation:**
  * Given a string, write a function to check if it is a permutation of a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
  * Example
    * Input: Tact Coa
    * Output: True (permutations: "taco cat", "atco eta", etc.)
* **1.5 One Away:**
  * There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character. Given two strings, write a function to check if they are one edit (or zero edits) away.
  * EXAMPLE
    * pale, ple -> true
    * pales, pale -> true
    * pale, bale -> true
    * pale, bake -> false
* **1.6 String Compression:**
  * Implement a method to perform basic string compression using the counts of repeated characters. For example, the string aabcccccaaa would become a2b1c5a3. If the "compressed" string would not become smaller than the original string, your method should return the original string. You can assume the string has only uppercase and lowercase letters (a - z).
* **1.7 Rotate Matrix:**
  * Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
* **1.8 Zero Matrix:**
  * Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0.
* **1.9 String Rotation:**
  * Assume you have a method isSubstring which checks if one word is a substring of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one call to isSubstring (e.g., "waterbottle" is a rotation of "erbottlewat").
