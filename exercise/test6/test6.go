package main

/*
1.总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。
	答：1、发送方和接收方规定固定大小的缓冲区，也就是发送和接收都使用固定大小的 byte[] 数组长度，当字符长度不够时使用空字符弥补；
			优缺点分析：从以上代码可以看出，虽然这种方式可以解决粘包和半包的问题，但这种固定缓冲区大小的方式增加了不必要的数据传输，
				因为这种方式当发送的数据比较小时会使用空字符来弥补，所以这种方式就大大的增加了网络传输的负担，所以它也不是最佳的解决方案。
	   2、在 TCP 协议的基础上封装一层数据请求协议，既将数据包封装成数据头（存储数据正文大小）+ 数据正文的形式，这样在服务端就可以知道每个
	      数据包的具体长度了，知道了发送数据的具体边界之后，就可以解决半包和粘包的问题了；
			这种解决方案的实现思路是将请求的数据封装为两部分：数据头+数据正文，在数据头中存储数据正文的大小，当读取的数据小于数据头中的大小时，继续读取数据，直到读取的数据长度等于数据头中的长度时才停止。
			因为这种方式可以拿到数据的边界，所以也不会导致粘包和半包的问题，但这种实现方式的编码成本较大也不够优雅，因此不是最佳的实现方案，因此我们这里就略过，直接来看最终的解决方案吧。
	   3、以特殊的字符结尾，比如以“\n”结尾，这样我们就知道结束字符，从而避免了半包和粘包问题（推荐解决方案）。
			以特殊字符结尾就可以知道流的边界了，因此也可以用来解决粘包和半包的问题，此实现方案是我们推荐最终解决方案。
			这种解决方案的核心是，使用 Java 中自带的 BufferedReader 和 BufferedWriter，也就是带缓冲区的输入字符流和输出字符流，通过写入的时候加上 \n 来结尾，
			读取的时候使用 readLine 按行来读取数据，这样就知道流的边界了，从而解决了粘包和半包的问题。
*/


