1. 多个defer执行顺序
   defer先进后出，新出现的defer语句最后执行
2. defer和return那个先执行
   return之后的语句先执行，defer后的语句后执行
3. 返回值参数有名，遇见defer可能被修改
4. defer遇见panic
   panic之前的 defer 按照出现顺序，最晚的现在执行。panic出现之后的defer不执行
5. defer中带有panic
   panic仅有最后一个可以被revover捕获。
   触发panic("panic")后defer顺序出栈执行，
   第一个被执行的defer中 会有panic("defer panic")异常语句， 
   这个异常将会覆盖掉main中的异常panic("panic")，最后这个异常被第二个执行的defer捕获到