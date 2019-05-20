
#服务绑定
├── bindService
│   └── container.go   //容器

这里面是在初始化实现绑定



#实现服务
── service  //对外服务的定义
    ├── projectFactory.go  //复杂的服务
    ├── projectService.go  //定义服务
    └── projectServiceImpl.go  //实现服务

在projectServiceImpl中
type ProjectServiceImpl struct { //接口声明
	// ProjectService            //实现的接口不用起名字 ?重复的调用了？不写也是能实现继承
	ProjectDAO domain.ProjectDAO //要使用的接口起名字
}

若是在projectServiceImpl里写实现，同时声明了ProjectService会出现重复的调用。报错。
可以将实现放在容器里。参考dbal的在容器中绑定的方法

#使用

在具体的路由中，指定要使用的controller，在容器中找到它，对应到相应的容器上。


#测试
启动项目： bee run aos   //在编译时，已经完成了服务的绑定
访问：http://127.0.0.1:6001/v1/servicetest
简要流程，在路由中找到对应的controller，对应的方法，在使用已经绑定的服务提供的List方法，在具体方法的实现中使用数据库连接进而找到数据，一层层返回结束

