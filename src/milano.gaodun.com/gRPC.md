# gRPC 是什么？
```
在 gRPC 里客户端应用可以像调用本地对象一样直接调用另一台不同的机器上服务端应用的方法，使得您能够更容易地创建分布式应用和服务。与许多 RPC 系统类似，gRPC 也是基于以下理念：定义一个服务，指定其能够被远程调用的方法（包含参数和返回类型）。在服务端实现这个接口，并运行一个 gRPC 服务器来处理客户端调用。在客户端拥有一个存根能够像服务端一样的方法。
使用 protocol buffers
gRPC 默认使用 protocol buffers，这是 Google 开源的一套成熟的结构数据序列化机制（当然也可以使用其他数据格式如 JSON）。
```

# gRPC 概念
```
本文档通过对于 gRPC 的架构和 RPC 生命周期的概览来介绍 gRPC 的主要概念。本文是在假设你已经读过文档部分的前提下展开的。针对具体语言细节请查看对应语言的快速开始、教程和参考文档（很快就会有完整的文档）。

概览
服务定义
正如其他 RPC 系统，gRPC 基于如下思想：定义一个服务， 指定其可以被远程调用的方法及其参数和返回类型。gRPC 默认使用 protocol buffers 作为接口定义语言，来描述服务接口和有效载荷消息结构。如果有需要的话，可以使用其他替代方案。
```

# grpc 使用
```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

// protoc --go_out=plugins=grpc:. helloworld.proto

go get google.golang.org/grpc 翻墙 🤢

在 github 可以找到源码，下载后复制到对应目录即可的： 
google.golang.org/grpc 对应的代码地址在： https://github.com/grpc/grpc-go 
google.golang.org/cloud/compute/metadata 对应的代码地址在： https://github.com/GoogleCloudPlatform/gcloud-golang 
golang.org/x/oauth2 对应的代码地址在： https://github.com/golang/oauth2 
golang.org/x/net/context 对应的代码地址在： https://github.com/golang/net 
这些包的源码也可以通过 http://gopm.io/ 或者 http://golangtc.com/download/package 进行下载.

ERROR:package google.golang.org/genproto/googleapis/rpc/status: unrecognized import path "google.golang.org/genproto/googleapis/rpc/status" (https fetch: Get https://google.golang.org/genproto/googleapis/rpc/status?go-get=1: unexpected EOF)
解决方案：
1、cd $GOPATH/src/google.golang.org
2、git clone https://github.com/google/go-genproto
3、mv -f go-genproto  genproto
```

# 服务定义
```GO
service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  required string greeting = 1;
}

message HelloResponse {
  required string reply = 1;
}
```

```
gRPC 允许你定义四类服务方法：
1、单项 RPC，即客户端发送一个请求给服务端，从服务端获取一个应答，就像一次普通的函数调用。
rpc SayHello(HelloRequest) returns (HelloResponse){
}

2、服务端流式 RPC，即客户端发送一个请求给服务端，可获取一个数据流用来读取一系列消息。客户端从返回的数据流里一直读取直到没有更多消息为止。
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){
}

3、客户端流式 RPC，即客户端用提供的一个数据流写入并发送一系列消息给服务端。一旦客户端完成消息写入，就等待服务端读取这些消息并返回应答。
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse) {
}

4、双向流式 RPC，即两边都可以分别通过一个读写数据流来发送一系列消息。这两个数据流操作是相互独立的，所以客户端和服务端能按其希望的任意顺序读写，例如：服务端可以在写应答前等待所有的客户端消息，或者它可以先读一个消息再写一个消息，或者是读写相结合的其他方式。每个数据流里消息的顺序会被保持。
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse){
}
```

