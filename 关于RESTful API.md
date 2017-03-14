<b>

restful是一种范式，不是一种协议。restful接口只是对http协议使用更加规范，回归了http协议本身之前的约定。原则上是对资源、集合、服务（URL），get、post、put、delete（操作）的合理使用。

get post delete put等各种Method的区别是语义上的。是为了让你的应用程序能够清晰简化的进行通信。对于HTTP本身来讲没有任何区别。当然安全性也没有任何区别。但有一些客观因素比如公司斥巨资买了一个防火墙，发现只能拦截过滤POST和GET请求，那就没办法了。

REST的目的是“建立十年内不会过时的软件系统架构"，所以它具备三个特点： 

1. 状态无关 —— 确保系统的横向拓展能力 

2. 超文本驱动，Fielding的原话是”hypertext-driven" —— 确保系统的演化能力 

3. 对 resource 相关的模型建立统一的原语，例如：uri、http的method定义等 —— 确保系统能够接纳多样而又标准的客户端 

优点：

REST 的应用可以充分地挖掘 HTTP 协议对缓存支持的能力。

从另外一个角度看，第一条保证服务端演化，第三条保证客户端演化，第二条保证应用本身的演化，这实在是一个极具抽象能力的方案。

关键原则：

一、资源: 为所有“事物”定义ID

在这里我使用了“事物”来代替更正式准确的术语“资源”，每个事物都应该是可标识的，都应该拥有一个明显的ID——在Web中，代表ID的统一概念是：URI。

二、超媒体被当作应用状态引擎: 将所有事物链接在一起

任何可能的情况下，使用链接指引可以被标识的事物（资源）。

三、使用标准方法

GET方法支持非常高效、成熟的缓存，所以在很多情况下，你甚至不需要向服务器发送请求。还可以肯定的是，GET方法具有幂等性[译注：指多个相同请求返回相同的结果]——如果你发送了一个GET请求没有得到结果，你可能不知道原因是请求未能到达目的地，还是响应在反馈的途中丢失了。幂等性保证了你可以简单地再发送一次请求解决问题。幂等性同样适用于PUT（基本的含义是“更新资源数据，如果资源不存在的话，则根据此URI创建一个新的资源”）和DELETE（你完全可以一遍又一遍地操作它，直到得出结论——删除不存在的东西没有任何问题）方法。POST方法，通常表示“创建一个新资源”，也能被用于调用任意过程，因而它既不安全也不具有幂等性。

    /orders
        GET - list all orders
        PUT - unused
        POST - add a new order
        DELETE - unused
    /orders/{id}
        GET - get order details
        PUT - update order
        POST - add item
        DELETE - cancel order
    /customers
        GET - list all customers
        PUT - unused
        POST - add a new customer
        DELETE - unused
    /customers/{id}
        GET - get customer details
        PUT - update customer
        POST - unused
        DELETE - cancel customer
    /customers/{id}/orders
        GET - get all orders for customer
        PUT - unused
        POST - unused
        DELETE - cancel all customer orders

标识一个顾客的URI上的GET方法正好相当于getCustomerDetails操作。

四、资源多重表述

如果客户程序知道如何处理一种特定的数据格式，那就可以与所有提供这种表述格式的资源交互，包括XML、JSON、HTML等。即服务器端需要向外部提供多种格式的资源表述，供不同的客户端使用。比如移动应用可以使用XML或JSON和服务器端通信，而浏览器则能够理解HTML。

五、无状态通信

例如我订阅了一个人的博客，想要获取他发表的所有文章（这里『他发表的所有文章』就是一个资源Resource）。于是我就向他的服务发出请求，说『我要获取你发表的所有文章，最好是atom格式的』，这时候服务器向你返回了atom格式的文章列表第一页（这里『atom格式的文章列表』就是表征Representation）。你看到了第一页的页尾，想要看第二页，这时候有趣的事情就来了。如果服务器记录了应用的状态（stateful），那么你只要向服务询问『我要看下一页』，那么服务器自然就会返回第二页。类似的，如果你当前在第二页，想服务器请求『我要看下一页』，那就会得到第三页。但是REST的服务器恰恰是无状态的（stateless），服务器并没有保持你当前处于第几页，也就无法响应『下一页』这种具有状态性质的请求。因此客户端需要去维护当前应用的状态（application state），也就是『如何获取下一页资源』。当然，『下一页资源』的业务逻辑必然是由服务端来提供。服务器在文章列表的atom表征中加入一个URI超链接（hyper link），指向下一页文章列表对应的资源。客户端就可以使用统一接口（Uniform Interface）的方式，从这个URI中获取到他想要的下一页文章列表资源。上面的『能够进入下一页』就是应用的状态（State）。服务器把『能够进入下一页』这个状态以atom表征形式传输（Transfer）给客户端就是表征状态传输（REpresentational State Transfer）这个概念。

举个具体API的例子：

请求：
    GET /posts HTTP/1.1
    Accept: application/atom+xml

响应：
    HTTP/1.1 200 OK
    Content-Type: application/atom+xml

    <?xml version="1.0" encoding="utf-8"?>
    …
    <link href="http://example.org/posts" rel="self" />
    <link href="http://example.org/posts?pn=2" rel="next" />
    …
        <link href="http://example.org/post-xxx" />
    …
    </feed>

注意上面atom格式中的多个<link>元素，它们分别定义了当前状态下合法的状态转移。

例如，这是一个指向自己的链接，其中rel属性指定了状态转移的关系为自身。

    <link href="http://example.org/posts" rel="self" />

这是下一页的链接，

    <link href="http://example.org/posts?pn=2" rel="next" />

如果当前不是第一页的话，就会有类似如下的链接来表示上一页，

    <link href="http://example.org/posts?pn=2" rel="prev" />

而这个是某一篇文章的链接，

    <link href="http://example.org/post-xxx" />

总结一下，就是：

1、服务器生成包含状态转移的表征数据，用来响应客户端对于一个资源的请求；

2、客户端借助这份表征数据，记录了当前的应用状态以及对应可转移状态的方式。

当然，为了要实现这一系列的功能，一个不可或缺的东西就是超文本（hypertext）或者说超媒体类型（hypermedia type）。这绝对不是一个简简单单的媒体类型（例如，JSON属性列表）可以做到的。

大部分号称REST的API实际上属于其第二层HTTP Verbs，并没有达到Richardson成熟度模型的第三个级别：Hypermedia Controls。 而REST的发明者Roy Fielding博士更是直言“Hypermedia作为应用引擎”是REST的前提， 这不是一个可选项，如果没有Hypermedia，那就不是REST。相比第二层，第三层的Web服务具有一个很明显的优势，客户端与服务端交互解耦。服务端可以仅仅提供单一的入口，客户端只要依次“遍历”超链接，就可以完成对应的合法业务逻辑。当资源的转换规则发送变化时（如某一页由于历史文章被删除了而没有下一页，又或者某篇文章转储在了其他网站上），客户端不需要作额外的更新升级，只需要升级服务端返回的超链接即可。

    {
      "links": {
        "self": "http://example.com/articles",
        "next": "http://example.com/articles?page[offset]=2",  # 上一页下一页
        "last": "http://example.com/articles?page[offset]=10"
        or
        "self": "http://localhost:57900/orders/123456",  # get请求
        "cancel": "http://localhost:57900/orders/123456",  # delete请求
        "payment": "http://localhost:57900/orders/123456/payments"
      },
      "data": {
        ...
      },
      "msg" : "done", // 请求状态描述，调试用
      "code": 1001, // 业务自定义状态码
      "extra" : { // 全局附加数据，字段、内容不定
            "type": 1,
            "desc": "签到成功！"
      }
    }

自定义状态码

    // 授权相关
    1001: 无权限访问
    1002: access_token过期
    1003: unique_token无效
    // 用户相关
    2001: 未登录
    2002: 用户信息错误
    2003: 用户不存在
    // 业务1
    3001: 业务1XXX
    3002: 业务1XXX

Headers

很多REST API犯的比较大的一个问题是：不怎么理会request headers。对于REST API，有一些HTTP headers很重要：

Accept：服务器需要返回什么样的content。如果客户端要求返回"application/xml"，服务器端只能返回"application/json"，那么最好返回status code 406 not acceptable（RFC2616），当然，返回application/json也并不违背RFC的定义。一个合格的REST API需要根据Accept头来灵活返回合适的数据。

If-Modified-Since/If-None-Match：如果客户端提供某个条件，那么当这条件满足时，才返回数据，否则返回304 not modified。比如客户端已经缓存了某个数据，它只是想看看有没有新的数据时，会用这两个header之一，服务器如果不理不睬，依旧做足全套功课，返回200 ok，那就既不专业，也不高效了。

If-Match：在对某个资源做PUT/PATCH/DELETE操作时，服务器应该要求客户端提供If-Match头，只有客户端提供的Etag与服务器对应资源的Etag一致，才进行操作，否则返回412 precondition failed。这个头非常重要，下文详解。 


安全性

前面说过，REST API承前启后，是系统暴露给外界的接口，所以，其安全性非常重要。安全并单单不意味着加密解密，而是一致性（integrity），机密性（confidentiality）和可用性（availibility）。

一、请求数据验证

我们从数据流入REST API的第一步 —— 请求数据的验证 —— 来保证安全性。你可以把请求数据验证看成一个巨大的漏斗，把不必要的访问统统过滤在第一线：

1.Request headers是否合法：如果出现了某些不该有的头，或者某些必须包含的头没有出现或者内容不合法，根据其错误类型一律返回4xx。比如说你的API需要某个特殊的私有头（e.g. X-Request-ID），那么凡是没有这个头的请求一律拒绝。这可以防止各类漫无目的的webot或crawler的请求，节省服务器的开销。 

2.Request URI和Request body是否合法：如果请求带有了不该有的数据，或者某些必须包含的数据没有出现或内容不合法，一律返回4xx。比如说，API只允许querystring中含有query，那么"?sort=desc"这样的请求需要直接被拒绝。有不少攻击会在querystring和request body里做文章，最好的对应策略是，过滤所有含有不该出现的数据的请求。 

二、数据完整性验证

REST API往往需要对backend的数据进行修改。修改是个很可怕的操作，我们既要保证正常的服务请求能够正确处理，还需要防止各种潜在的攻击，如replay。数据完整性验证的底线是：保证要修改的数据和服务器里的数据是一致的 —— 这是通过Etag来完成。

Etag可以认为是某个资源的一个唯一的版本号。当客户端请求某个资源时，该资源的Etag一同被返回，而当客户端需要修改该资源时，需要通过"If-Match"头来提供这个Etag。服务器检查客户端提供的Etag是否和服务器同一资源的Etag相同，如果相同，才进行修改，否则返回412 precondition failed。

使用Etag可以防止错误更新。比如A拿到了Resource X的Etag X1，B也拿到了Resource X的Etag X1。B对X做了修改，修改后系统生成的新的Etag是X2。这时A也想更新X，由于A持有旧的Etag，服务器拒绝更新，直至A重新获取了X后才能正常更新。

Etag类似一把锁，是数据完整性的最重要的一道保障。Etag能把绝大多数integrity的问题扼杀在摇篮中，当然，race condition还是存在的：如果B的修改还未进入数据库，而A的修改请求正好通过了Etag的验证时，依然存在一致性问题。这就需要在数据库写入时做一致性写入的前置检查。

三、访问控制

REST API需要清晰定义哪些操作能够公开访问，哪些操作需要授权访问。一般而言，如果对REST API的安全性要求比较高，那么，所有的API的所有操作均需得到授权。

在HTTP协议之上处理授权有很多方法，如HTTP BASIC Auth，OAuth，HMAC Auth等，其核心思想都是验证某个请求是由一个合法的请求者发起。Basic Auth会把用户的密码暴露在网络之中，并非最安全的解决方案，OAuth的核心部分与HMAC Auth差不多，只不过多了很多与token分发相关的内容。这里我们主要讲讲HMAC Auth的思想。

回到Security的三个属性：一致性，机密性，和可用性。HMAC Auth保证一致性：请求的数据在传输过程中未被修改，因此可以安全地用于验证请求的合法性。

HMAC主要在请求头中使用两个字段：Authorization和Date（或X-Auth-Timestamp）。Authorization字段的内容由":"分隔成两部分，":"前是access-key，":"后是HTTP请求的HMAC值。在API授权的时候一般会为调用者生成access-key和access-secret，前者可以暴露在网络中，后者必须安全保存。当客户端调用API时，用自己的access-secret按照要求对request的headers/body计算HMAC，然后把自己的access-key和HMAC填入Authorization头中。服务器拿到这个头，从数据库（或者缓存）中取出access-key对应的secret，按照相同的方式计算HMAC，如果其与Authorization header中的一致，则请求是合法的，且未被修改过的；否则不合法。

在做HMAC的时候，request headers中的request method，request URI，Date/X-Auth-Timestamp等header会被计算在HMAC中。将时间戳计算在HMAC中的好处是可以防止replay攻击。客户端和服务器之间的UTC时间正常来说偏差很小，那么，一个请求携带的时间戳，和该请求到达服务器时服务器的时间戳，中间差别太大，超过某个阈值（比如说120s），那么可以认为是replay，服务器主动丢弃该请求。

使用HMAC可以很大程度上防止DOS攻击 —— 无效的请求在验证HMAC阶段就被丢弃，最大程度保护服务器的计算资源。

四、HTTPS

HMAC Auth尽管在保证请求的一致性上非常安全，可以用于鉴别请求是否由合法的请求者发起，但请求的数据和服务器返回的响应都是明文传输，对某些要求比较高的API来说，安全级别还不够。这时候，需要部署HTTPS。在其之上再加一层屏障。

五、其他

做到了接口一致性（符合RFC）和安全性，REST API可以算得上是合格了。当然，一个实现良好的REST API还应该有如下功能：

* rate limiting：访问频率限制。 
* metrics：服务器应该收集每个请求的访问时间，到达时间，处理时间，latency，便于了解API的性能和客户端的访问分布，以便更好地优化性能和应对突发请求。 
* docs：丰富的接口文档 - API的调用者需要详尽的文档来正确调用API，可以用swagger来实现。 
* hooks/event propogation：其他系统能够比较方便地与该API集成。比如说添加了某资源后，通过kafka或者rabbitMQ向外界暴露某个消息，相应的subscribers可以进行必要的处理。不过要注意的是，hooks/event propogation可能会破坏REST API的幂等性，需要小心使用。

各个社区里面比较成熟的REST API framework/library：

* Python: django-rest-framework（django），eve（flask）。各有千秋。可惜python没有好的类似webmachine的实现。 
* Erlang/Elixir: webmachine/ewebmachine。 
* Ruby: webmachine-ruby。 
* Clojure：liberator。
