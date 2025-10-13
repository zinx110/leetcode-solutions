# 208. Implement Trie (Prefix Tree)

Implement the `Trie` class:

- `Trie()` Initializes the trie object.
- `void insert(String word)` Inserts the string `word` into the trie.
- `boolean search(String word)` Returns `true` if the string `word` is in the trie (i.e., was previously inserted); otherwise, returns `false`.
- `boolean startsWith(String prefix)` Returns `true` if there is any string in the trie that starts with the string `prefix`; otherwise, returns `false`.

---

## Example 1

Input
```
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
```
Output
```
[null, null, true, false, true, null, true]
```

Explanation
```
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return true
trie.search("app");     // return false
trie.startsWith("app"); // return true
trie.insert("app");
trie.search("app");     // return true
```

---

## Constraints

- `1 <= word.length, prefix.length <= 2000`
- `word` and `prefix` consist only of lowercase English letters.
- At most `3 * 10^4` calls will be made to `insert`, `search`, and `startsWith`.


