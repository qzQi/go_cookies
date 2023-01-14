go的http应该是仅支持restful API？get里面的body是解析不到的，不是说真的没有。

而是说，post方法可以把formValue放在body里面。

通过get提交的formValue必须通过queryString的方法。

bug：之前以为formValue是通过body传输的，现在知道原来传递的时候就要说明是form类型。
body是json的话需要自己解析。


读一下FormValue的doc。form形式的可以直接使用；但是json这种就是普通的字符串（序列化与反序列化），经过我们自己的解析才能够使用。

其实不算是错误，在使用的时候最后的那个/区分不区分差别不大。