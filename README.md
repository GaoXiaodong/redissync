背景：
    部署了redis cluster，采用分片方式，没办法支持redis官方提出的redlock算法
    本身自己业务没有并发量

在官方redsync的基础上，去掉了redlock算法的代码，相当于是单机应用，有可能出现当槽点丢失的情况下获取锁冲突的情况
