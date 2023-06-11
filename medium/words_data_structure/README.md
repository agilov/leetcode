# Problem 211: Design Add and Search Words Data Structure

## Problem Description

Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the `WordDictionary` class:

- `WordDictionary()` Initializes the object.
- `void addWord(word)` Adds `word` to the data structure, it can be matched later.
- `bool search(word)` Returns true if there is any string in the data structure that matches `word` or false otherwise. A `word` could contain the dot character '.' to represent any one letter.

## Constraints

- The number of calls to `addWord` and `search` is at most 30000.
- You may assume that `addWord` always follows `search`.
- The maximum length of `word` is 500.
- `word` is lowercase and is composed of only lowercase English letters and the dot character '.'.
- For a call to `search`, `word` will not start or end with the dot character '.'.

## Examples

### Example 1:

```plaintext
Input:
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output: [null,null,null,null,false,true,true,true]

Explanation:
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True
