# Word Search II

## Description:

Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cells, where "adjacent" cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

### Constraints:

* The board and words consist only of lowercase and uppercase English letters.
* The board's dimension is `m x n`, where 1 <= m, n <= 12.
* The length of words[i] is `l`, where 1 <= l <= 10.
* The list of words is `words`, where 1 <= words.length <= 3000.

### Example:

```
Input: 
board = [
['o','a','a','n'],
['e','t','a','e'],
['i','h','k','r'],
['i','f','l','v']
]
words = ["oath","pea","eat","rain"]

Output: ["eat","oath"]
```

In the given example, the words "eat" and "oath" can be found in the board, but "pea" and "rain" can't be found, so the output is ["eat","oath"].
