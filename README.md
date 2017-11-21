go-dry
======

-DRY (don't repeat yourself) package for Go
  		  
-Documentation: http://godoc.org/github.com/ungerik/go-dry


## Desc

此项目作为 go 的工具库，提供了很多常用操作的的封装，每个文件都是对该文件名域下的封装。


## Notes
* Crc64: https://en.wikipedia.org/wiki/Cyclic_redundancy_check
    * hash func 的一种，主要用于检测数据在传输的过程中是否损坏, 应该不适合像md5那种检测文件是否改动
* AES encription, 

## Files

```
├── LICENSE
├── README.md
├── bytes.go
├── compression.go                 # ！sync.Pool 对象需要注意下，使用方式，提供了压缩的方法。
├── debug.go                       # 
├── doc.go                         # package 文档
├── encryption.go                  # AES encription
├── endian.go                      #
├── errors.go                      # 封装了些error处理的函数
├── file.go                        # read file from url, md5, crc64, flate, Gzip, File size|exist|readlastline, format|scan string to|from file, list dir, copy file, copy dir
├── http.go                        # 封装了 http 的一些工具，可以检测 compression type 的request，并提供 Gzip or flate 的压缩返回, 还提供了常用的 http.Get|Post...方法
├── io.go                          # CountingReader, CountingWriter, CountingReadWriter 读取和写入的时候的标记会包含在定义的对象中
├── net.go                         # get current machine's ip addr
├── os.go                          # helper methods that deals with environments
├── rand.go                        # 用于生成随机的16 base的字符串
├── reflect.go                     # 这个文件需要特别注意下，毕竟reflect不熟，但基本的核心思想就是调用 Refect.TypeOf | ValueOf 方法将目标转换到 Reflect 世界中的type类型做处理， newReflectSortable 方法需要注意下， 这个文件也没有特别看懂，后期再处理
├── reflect_test.go
├── shortcuts.go
├── string.go                      # create url params, various hash, json , string list contain or not, html replace/remove tags, joins, with slice, format methods, sort map keys/values, split, StringSet, StringGroupedNumberPostfixSorter & sort.Sort 的组合使用需要注意下 
├── string_test.go
├── stringbuilder.go               # StringBuilder, like java.
└── sync.go                        # SyncBool, SyncInt(atomic包应该有int), SyncString, Sync[Map|StringMap|PoolMap] 

```