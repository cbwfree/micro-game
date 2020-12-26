package redis

//  指令列表
const (
	//  键管理
	CmdDel       = "DEL"       // 删除key
	CmdDump      = "DUMP"      // 序列化key, 并返回序列化的值
	CmdReStore   = "RESTORE"   // 反序列化给定的序列化值，并将它和给定的 key 关联
	CmdExists    = "EXISTS"    // 检查key是否存在
	CmdExpire    = "EXPIRE"    // 为 key 设置过期时间, 值为秒
	CmdExpireAt  = "EXPIREAT"  // 为key设置过期时间, 值为秒级unix时间戳
	CmdPExpire   = "PEXPIRE"   // 为key设置过期时间, 值为毫秒
	CmdPExpireAt = "PEXPIREAT" // 为key设置过期时间, 值为毫秒级unix时间戳
	CmdKeys      = "KEYS"      // 查找符合特定模式的key
	CmdMove      = "MOVE"      // 移动当前数据库的key到指定数据库中
	CmdPersist   = "PERSIST"   // 移除key的过期时间
	CmdPTTL      = "PTTL"      // 以毫秒单位返回key的剩余过期时间
	CmdTTL       = "TTL"       // 以秒为单位返回key的剩余过期时间
	CmdRandomKey = "RANDOMKEY" // 从当前数据库中随机返回一个KEY
	CmdRename    = "RENAME"    // 修改key名称
	CmdRenameNX  = "RENAMENX"  // 仅当newkey不存在时, 将key修改为newkey
	CmdType      = "TYPE"      // 返回key所存储的值类型
	CmdSort      = "SORT"      // 返回或保存给定列表、集合、有序集合 key 中经过排序的元素
	CmdScan      = "SCAN"      // 迭代key

	//  字符串
	CmdAppend      = "APPEND"      // 如果 key 已经存在并且是一个字符串， APPEND 命令将指定的 value 追加到该 key 原来值（value）的末尾
	CmdSet         = "SET"         // 设置key的值
	CmdGet         = "GET"         // 获取key的值
	CmdGetRange    = "GETRANGE"    // 返回key中字符串的子字符串
	CmdGetSet      = "GETSET"      // 设置key的值, 并返回旧值
	CmdGetBit      = "GETBIT"      // 对key的值, 获取指定偏移量上的位(bit)
	CmdMGet        = "MGET"        // 获取一个或多个key的值
	CmdSetBit      = "SETBIT"      // 对key的值, 设置或清除指定偏移量上的位(bit)
	CmdSetEX       = "SETEX"       // 设置key的值, 并同时设置过期时间, 单位秒
	CmdSetNX       = "SETNX"       // 只有在key不存在时设置key的值
	CmdSetRange    = "SETRANGE"    // 设置key的值, 从指定偏移量开始
	CmdSTRLen      = "STRLEN"      // 返回key的值长度
	CmdMSet        = "MSET"        // 同时设置一个或多个key的值
	CmdMSetNX      = "MSETNX"      // 当key都不存在时, 设置一个或多个key的值
	CmdPSetEX      = "PSETEX"      // 设置key的值, 并同时设置过期时间, 单位毫秒
	CmdIncr        = "INCR"        // 将key中存储的数值增一
	CmdIncrBy      = "INCRBY"      // 将key中存储的数值加上给定的增量值
	CmdIncrByFloat = "INCRBYFLOAT" // 将key所储存的值加上给定的浮点增量值
	CmdDecr        = "DECR"        // 将 key 中储存的数字值减一
	CmdDecrBy      = "DECRBY"      // 将 key 所储存的值减去给定的减量值

	//  Hash (哈希表)
	CmdHDel         = "HDEL"         // 删除一个或多个hash表字段
	CmdHExists      = "HEXISTS"      // 检查hash表中指定字段是否存在
	CmdHGetAll      = "HGETALL"      // 获取hash表中所有字段和值
	CmdHGet         = "HGET"         // 获取hash表中一个或多个field的值
	CmdHMGet        = "HMGET"        // 获取所有给定字段的值
	CmdHIncrBy      = "HINCRBY"      // 为哈希表 key 中的指定字段的整数值加上增量值
	CmdHIncrByFloat = "HINCRBUFLOAT" // 为哈希表 key 中的指定字段的浮点数值加上增量值
	CmdHKeys        = "HKEYS"        // 获取所有哈希表中的字段
	CmdHVals        = "HVALS"        // 获取哈希表中所有值
	CmdHLen         = "HLEN"         // 获取哈希表中字段的数量
	CmdHMSet        = "HMSET"        // 同时将多个 field-value (域-值)对设置到哈希表 key 中
	CmdHSet         = "HSET"         // 将哈希表 key 中的字段 field 的值设为 value
	CmdHSetNX       = "HSETNX"       // 只有在字段 field 不存在时，设置哈希表字段的值
	CmdHScan        = "HSCAN"        // 迭代哈希表中的键值对。
	CmdHStrLen      = "HSTRLEN"      // 返回哈希表 key 中， 与给定域 field 相关联的值的字符串长度

	//  List (列表)
	CmdBLPop      = "BLPOP"      // 移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
	CmdBRPop      = "BRPOP"      // 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
	CmdBRPopLPush = "BRPOPLPUSH" // 从列表中弹出一个值，将弹出的元素插入到另外一个列表中并返回它； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
	CmdLIndex     = "LINDEX"     // 通过索引获取列表中的元素
	CmdLInsert    = "LINSERT"    // 在列表的元素前或者后插入元素
	CmdLLen       = "LLEN"       // 获取列表长度
	CmdLPop       = "LPOP"       // 移出并获取列表的第一个元素
	CmdLPush      = "LPUSH"      // 将一个或多个值插入到列表头部
	CmdLPushHX    = "LPUSHHX"    // 将一个值插入到已存在的列表头部
	CmdLRange     = "LRANGE"     // 获取列表指定范围内的元素
	CmdLRem       = "LREM"       // 移除列表元素
	CmdLSet       = "LSET"       // 通过索引设置列表元素的值
	CmdLTrim      = "LTRIM"      // 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除
	CmdRPop       = "RPOP"       // 移除并获取列表最后一个元素
	CmdRPopLPush  = "RPOPLPUSH"  // 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
	CmdRPush      = "RPUSH"      // 在列表末尾中添加一个或多个值
	CmdRPushHX    = "RPUSHHX"    // 在已存在的列表末尾添加值

	//  Set (无序集合)
	CmdSAdd        = "SADD"        // 向集合添加一个或多个元素
	CmdSCard       = "SCARD"       // 获取集合的成员数
	CmdSDiff       = "SDIFF"       // 返回给定所有集合的差集
	CmdSDiffStore  = "SDIFFSTORE"  // 返回给定所有集合的差集, 并存储在新集合中
	CmdSInter      = "SINTER"      // 返回给定所有集合的交集
	CmdSInterStore = "SINTERSTORE" // 返回所有给定集合的交集, 并存储在新集合中
	CmdSUnion      = "SUNION"      // 返回所有给定集合的并集
	CmdSUnionStore = "SUNIONSTORE" // 返回所有给定集合的并集, 并保存到新集合中
	CmdSisMember   = "SISMEMBER"   // 判断元素是否是集合的成员
	CmdSMembers    = "SMEMBERS"    // 返回集合中的所有成员
	CmdSMove       = "SMOVE"       // 将元素从key1移动到key2
	CmdSPop        = "SPOP"        // 移除并返回集合中的一个随机元素
	CmdSRandMember = "SRANDMEMBER" // 返回集合中一个或多个随机数
	CmdSRem        = "SREM"        // 移除集合中的一个或多个成员
	CmdSScan       = "SSCAN"       // 迭代集合中的元素

	//  Sorted Set (有序集合)
	CmdZAdd             = "ZADD"             // 添加一个或多个成员, 或者更新已存在成员的分数
	CmdZCard            = "ZCARD"            // 获取有序集合的成员数
	CmdZCount           = "ZCOUNT"           // 计算集合中指定区间分数的成员数
	CmdZIncrBy          = "ZINCRBY"          // 对集合中指定成员的分数加上增量值
	CmdZInterStore      = "ZINTERSTORE"      // 计算所有给定集合的交集, 并将结果存储在新的集合中
	CmdZUnionStore      = "ZUNIONSTORE"      // 计算给定的一个或多个有序集的并集，并存储在新的 key 中
	CmdZLexCount        = "ZLEXCOUNT"        // 在集合中计算指定字典区间成员数量
	CmdZRange           = "ZRANGE"           // 通过索引区间返回集合内指定区间内的成员
	CmdZRangeByLex      = "ZRANGEBYLEX"      // 通过字典区间返回集合的成员
	CmdZRangeByScore    = "ZRANGEBYSCORE"    // 通过分数返回集合指定区间内的成员
	CmdZRank            = "ZRANK"            // 返回集合中指定成员的索引
	CmdZRem             = "ZREM"             // 移除集合中一个或多个成员
	CmdZRemRangeByLex   = "ZREMRANGEBYLEX"   // 移除有序集合中给定的字典区间的所有成员
	CmdZRemRangeByRank  = "ZREMRANGEBYRANK"  // 移除有序集合中给定的排名区间的所有成员
	CmdZRemRangeByScore = "ZREMRANGEBYSCORE" // 移除有序集合中给定的分数区间的所有成员
	CmdZRevRange        = "ZREVRANGE"        // 返回有序集中指定区间内的成员，通过索引，分数从高到底
	CmdZRevRangeByScore = "ZREVRANGEBYSCORE" // 返回有序集中指定分数区间内的成员，分数从高到低排序
	CmdZRevRank         = "ZREVRANK"         // 返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序
	CmdZScore           = "ZSCORE"           // 返回有序集中，成员的分数值
	CmdZScan            = "ZSCAN"            // 迭代有序集合中的元素（包括元素成员和元素分值）

	//  事务
	CmdMulti   = "MULTI"   // 标记一个事务块的开始
	CmdExec    = "EXEC"    // 执行所有事务命令
	CmdDiscard = "DISCARD" // 取消事务，放弃执行事务块内的所有命令
	CmdUnwatch = "UNWATCH" // 取消watch命令对所有key的监听
	CmdWatch   = "WATCH"   // 监听一个或多个key, 如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断

	//  GEO (地理位置)
	CmdGEOAdd            = "GEOADD"            // 将给定的空间元素（纬度、经度、名字）添加到指定的键里面
	CmdGEOPos            = "GEOPOS"            // 从键里面返回所有给定位置元素的位置（经度和纬度）
	CmdGEODist           = "GEODIST"           // 返回两个给定位置之间的距离
	CmdGEORadius         = "GEORADIUS"         // 以给定的经纬度为中心，返回键包含的位置元素当中，与中心的距离不超过给定最大距离的所有位置元素
	CmdGEORadiusByMember = "GEORADIUSBYMEMBER" // 以给定的位置元素为中心, 返回键包含的位置元素当中，与中心的距离不超过给定最大距离的所有位置元素
	CmdGEOHash           = "GEOHASH"           // 返回一个或多个位置元素的 Geohash 表示

	//  Lua 脚本
	CmdEval         = "EVAL"    // 执行lua脚本
	CmdEvalSHA      = "EVALSHA" // 通过sha key执行保存在缓存中的lua脚本
	CmdScript       = "SCRIPT"  // REDIS 脚本操作
	CmdScriptExists = "EXISTS"  // 检查指定的脚本是否已经被保存到缓存中
	CmdScriptFlush  = "FLUSH"   // 清除缓存中所有已保存的lua脚本
	CmdScriptKill   = "KILL"    // 结束正在执行的Lua脚本
	CmdScriptLoad   = "LOAD"    // 将lua脚本加载到缓存中

	//  服务器操作
	CmdAuth          = "AUTH"           // 验证密码是否正确
	CmdPing          = "PING"           // 检查服务是否运行, 成功返回 PONG
	CmdSelect        = "SELECT"         // 切换到指定数据库
	CmdEcho          = "ECHO"           // 打印字符串
	CmdSave          = "SAVE"           // 同步保存数据到硬盘
	CmdBGSave        = "BGSAVE"         // 在后台异步保存当前数据库的数据到磁盘
	CmdLastSave      = "LASTSAVE"       // 返回最近一次redis成功存盘时间, 以unix时间戳表示
	CmdClientGetName = "CLIENT GETNAME" // 获取连接的名称
	CmdClientSetName = "CLIENT SETNAME" // 设置当前连接的名称
	CmdClientPause   = "CLIENT PAUSE"   // 在指定时间内终止运行来自客户端的命令
	CmdClusterSlots  = "CLUSTER SLOTS"  // 获取集群节点的映射数据
	CmdTime          = "TIME"           // 获取当前服务器时间
	CmdConfigSet     = "CONFIG SET"     // 修改redis配置参数
	CmdConfigGet     = "CONFIG GET"     // 获取redis配置参数
	CmdConfigRewrite = "CONFIG REWRITE" // 重写 redis.conf 配置文件
	CmdDBSize        = "DBSIZE"         // 返回当前数据库的key的数量
	CmdFlushAll      = "FLUSHALL"       // 删除所有数据库的所有key
	CmdFlushDB       = "FLUSHDB"        // 删除当前数据的所有key
	CmdInfo          = "INFO"           // 获取redis服务器的各种信息和统计数值
	CmdMonitor       = "MONITOR"        // 实时显示服务器接收到的命令 (调试使用)
)
