


### day02  context

#### 设计context 
* 必要性
    * 对于Web服务来说，无非是根据请求 http.Request ,构造响应 http.ResponseWriter.
    但这个两个对象所提供的接口粒度太细，如我们呢要构造一个完整的响应，需要考虑Header、body,而header包含 来状态码（StatusCode）、消息类型(ContentType)等信息.
    因此，如果不能进行有效的封装，那么框架的用户将需要写大量的繁杂代码，不具备可重用性。
    * 针对使用场景，封装*http.Request和http.ResponseWriter的方法，简化相关接口的调用，只是设计 Context 的原因之一。
    对于框架来说，还需要支撑额外的功能。例如，将来解析动态路由/hello/:name，参数:name的值放在哪呢？
    再比如，框架需要支持中间件，那中间件产生的信息放在哪呢？Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载。因此，设计 Context 结构，扩展性和复杂性留在了内部，而对外简化了接口。
    路由的处理函数，以及将要实现的中间件，参数都统一使用 Context 实例， Context 就像一次会话的百宝箱，可以找到任何东西。
    
    
#### day03 前缀树路由

* 实现动态路由最常用的数据结构，被称为前缀树(Trie树)。看到名字你大概也能知道前缀树长啥样了：每一个节点的所有的子节点都拥有相同的前缀。这种结构非常适用于路由匹配，比如我们定义了如下路由规则：
  
    * /:lang/doc
    * /:lang/tutorial
    * /:lang/intro
    * /about
    * /p/blog
    * /p/related
    
    我们用前缀树来表示，是这样的。