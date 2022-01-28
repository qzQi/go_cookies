the difference between a byte, a character, and a rune

https://go.dev/blog/strings

字符串string是由bytes组成的，所以使用index获得的是byte，但是一个byte不一定是一个character。

“When I index a Go string at position n, why don’t I get the nth character?”