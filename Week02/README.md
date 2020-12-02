作业
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么？应该怎么做请写出代码

举例：
 GetUsers（[]User, error）

1. dao层查询，当遇到sql.ErrNoRows， 应该不报错返回空切片， 其他错误，Wrap后返回
2. 数据库其他增、删、改遇到错误，Wrap后返回

代码实现：

