# golang_valid_parenthesis_string

Given a string `s` containing only three types of characters: `'('`, `')'` and `'*'`, return `true` *if* `s` *is **valid***.

The following rules define a **valid** string:

- Any left parenthesis `'('` must have a corresponding right parenthesis `')'`.
- Any right parenthesis `')'` must have a corresponding left parenthesis `'('`.
- Left parenthesis `'('` must go before the corresponding right parenthesis `')'`.
- `'*'` could be treated as a single right parenthesis `')'` or a single left parenthesis `'('` or an empty string `""`.

## Examples

**Example 1:**

```
Input: s = "()"
Output: true

```

**Example 2:**

```
Input: s = "(*)"
Output: true

```

**Example 3:**

```
Input: s = "(*))"
Output: true

```

**Constraints:**

- `1 <= s.length <= 100`
- `s[i]` is `'('`, `')'` or `'*'`.

## 解析

給定一個字串 s, 其中字串只能由 3 種字元所組成 ‘(’,  ‘*’,  ‘)’

如果字串 s 是 valid 必須符合以下3個條件：

1. 每個 ‘(’ 必須有一個對應的 ‘)’ 在其右方
2. 每個 ‘)’ 必須有一個對應的 ‘(’ 在其左方
3. ‘*’ 可以替換成 ‘(’, ‘)’ 或是空字元

要求寫一個演算法判斷給定的字串 s 是否為 valid

觀察首先在不考慮有出現 ‘*’ 的情況下

由左至右觀察

為了符合前兩個 valid 條件

對每個位置的字元去紀錄 ‘(’, ‘)’ 的個數

可以發現  ‘(’ 個數必須 ≥ ‘)’ 個數 才有機會 成為 valid 字串

當要考慮 ‘*’ 的情況會需要額外處理

使用動態規劃會是如下

定義 dp[i][j] 代表 s[i]…s[j] 是否為 valid

dp[i][j] 要是 true 有以下兩種可能

1. s[i] 是 ‘*’ , 且 s[i+1]s[i+2]…s[j] 是 valid
2. s[i] 是 ‘(’, 且存在 k , i+1 ≤ k ≤ j , s[k] = ‘)’ or ‘*’

     且 s[i+1]..s[k-1] 與 s[k+1]..s[j] 是 valid

透過以上關係式

可以初始化 dp[i][i] = true if s[i] = ‘*’

dp[i][i+1] = true , iff s[i]= ‘*’ | ‘(’  and s[i+1] = ‘*’ | ‘)’

然後 從 size = 2 遞增到 size = n

從 i =0 到 n

開使設定 dp[i][size] = true iff  s[i] == ‘*’ 且 dp[i+1][i+size] == true

或是 dp[i][size] = true iff s[i] == ‘(’ | ‘*’ 且

從 k = i+1 到 i+size 如果 dp[i+1][k-1] && dp[k+1][i+size]

![](https://i.imgur.com/y6sTxH9.png)

這樣每次需要 根據不同 size 還有每個起始位置 i 找尋對應的 k

所以時間複雜度是 O($n^3$)

因為要儲存 每個起始到結束位置的可能性，空間複雜度是 O($n^2$)

要做優化需要觀察一下

因為 對於每個 valid 的 s, 在每個位置 ‘(’ 的個數都要 ≥ ‘)’ 個數

所以其實只需要關注當下 ‘(’ 可能狀況

思考一下 ‘*’ 的轉化狀況

假設把 ‘*’ 都轉成 ‘(’ 紀錄 ，這種情況做紀錄 ‘(’ 個數會是最多’(’的情況

假設把 ‘*’ 都轉成 ‘)’ 紀錄 ，這種情況做紀錄 ‘(’ 個數會是最小’(’的情況

假設遇到非 ‘)’ 都做 +1, 遇到 ‘)’ 做 -1 

則會產生 ‘(’ 個數最多 比較情況 maxLeft , 用 maxLeft ≥ 0 來檢驗 ‘(’ 是否能夠配完

假設遇到非 ‘(’ 都做 -1, 遇到 ‘(’ 做 +1 

則會產生 ‘(’ 個數最小 比較情況 minLeft，用 minLeft == 0來檢驗 ‘)’ 是否能夠配完


![](https://i.imgur.com/NNhrNAL.png)

當遇到 maxLeft < 0 代表不可能 valid 回傳 false

當所有 maxLeft ≥ 0 則檢驗 minLeft 是否等於 0

如果等於 0 代表可以配完 ‘(’ 回傳 true 否則回傳 false

因為只需要 loop 整個字串所有字元

所以時間複雜度是 O(n)

空間複雜度是 O(1)

## 程式碼
```go
package sol

func checkValidString(s string) bool {
	sLen := len(s)
	maxLeft, minLeft := 0, 0
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for pos := 0; pos < sLen; pos++ {
		ch := s[pos]
		if ch != ')' {
			maxLeft += 1
		} else {
			maxLeft -= 1
		}
		if ch != '(' {
			minLeft -= 1
		} else {
			minLeft += 1
		}
		minLeft = max(0, minLeft)
		if maxLeft < 0 {
			return false
		}
	}
	return minLeft == 0
}
```
## 困難點

1. 需要看出 ‘(’ 與 ‘)’ 個數關係

## Solve Point

- [x]  初始化 maxLeft =0, minLeft =0
- [x]  逐步檢驗每個字元 s[i]
- [x]  如果 s[i] ≠ ‘)’ ， 則 maxLeft += 1 否則 maxLeft -= 1
- [x]  如果 s[i] ≠ ‘(’ ， 則 minLeft -= 1 否則 minLeft += 1
- [x]  更新 minLeft = max(0, minLeft)
- [x]  如果 maxLeft < 0 回傳 false
- [x]  當遍歷完所有字元，maxLeft ≥ 0 , 檢驗 minLeft == 0, 如果是 回傳 true, 否則回傳 false