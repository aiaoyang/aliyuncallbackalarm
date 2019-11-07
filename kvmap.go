package main

var (
	instance   = make(map[string]string)
	alertState = make(map[string]string)
	metric     = make(map[string]string)
)

func init() {
	metric["Host.mem.used"] = "量用内存"
	metric["Host.disk.writebytes"] = "磁盘每秒写入的字节数"
	metric["Host.cpu.system"] = "当前内核空间占用CPU百分比"
	metric["Host.diskussage.total"] = "磁盘存储总量"
	metric["Host.netout.rate"] = "网卡每秒发送的比特数"
	metric["Host.tcpconnection"] = "各种状态下的TCP连接数包括"
	metric["Host.load15"] = "过去15分钟的系统平均负载"
	metric["Host.diskusage.used"] = "磁盘的已用存储空间"
	metric["Host.load5"] = "过去5分钟的系统平均负载"
	metric["Host.disk.writeiops"] = "磁盘每秒的写请求数量"
	metric["Host.fs.inode"] = "inode使用率"
	metric["Host.netin.errorpackage"] = "设备驱动器检测到的接收错误包的数量"
	metric["Host.cpu.idle"] = "当前空闲CPU百分比"
	metric["Host.cpu.other"] = "其他占用CUP百分比"
	metric["Host.disk.readbytes"] = "磁盘每秒读取的字节数"
	metric["Host.disk.readiops"] = "磁盘每秒的读请求数量"
	metric["Host.mem.actualused"] = "用户实际使用的内存"
	metric["Host.mem.free"] = "剩余内存量"
	metric["Host.netout.packages"] = "网卡每秒发送的数据包数"
	metric["Host.netout.errorpackages"] = "设备驱动器检测到的发送错误包的数量"
	metric["Host.mem.total"] = "内存总量"
	metric["Host.mem.freeutilization"] = "剩余内存百分比"
	metric["Host.disk.utilization"] = "磁盘使用率"
	metric["Host.diskusage.free"] = "磁盘的剩余存储空间"
	metric["Host.load1"] = "过去1分钟的系统平均负载"
	metric["Host.netin.rate"] = "网卡每秒接收的比特数"
	metric["Host.netin.packages"] = "网卡每秒接收的数据包数"
	metric["Host.cpu.user"] = "当前用户空间占用CPU百分比"
	metric["Host.cpu.iowait"] = "当前等待IO操作的CPU百分比"
	metric["Host.cpu.total"] = "当前消耗的总CPU百分比"
	metric["Host.mem.usedutilization"] = "内存使用率"

	alertState["OK"] = "恢复正常"
	alertState["ALERT"] = "发生告警"
	alertState["INSUFFICIENT_DATA "] = "不重要"

}
