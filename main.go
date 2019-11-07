package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type recallMsg struct {
	// InstanceName  string `form:"instanceName"`
	// UserID        string `form:"userId"`
	AlterName  string `form:"alertName"`
	Timestamp  string `form:"timestamp"`
	AlertState string `form:"alertState"`
	Dimensions string `form:"dimensions"`
	Expression string `form:"expression"`
	CurValue   string `form:"curValue"`
	MetricName string `form:"metricName"`
	// MetricProject string `form:"metricProject"`
}

func (r *recallMsg) String() string {
	return fmt.Sprintf("告警规则: [ %s ]\n告警状态: [ %s ]\n主机名称: [ %s ]\n告警条件: [ %v ]\n告警指标: [ %s ]\n当前值: [ %s ]\n",
		// r.UserID,
		r.AlterName,
		alertState[r.AlertState],
		// 根据instanceID 获取主机名
		instance[keyValueToMap(r.Dimensions)["instanceId"]],
		r.Expression,
		metric[r.MetricName],
		r.CurValue,
	)
}
func init() {
	log.SetFlags(log.Llongfile | log.Ldate)
}
func main() {
	router := gin.Default()
	router.POST("/recall", recall)
	router.Run(":8080")

}

func recall(c *gin.Context) {
	log.Println(c.Request.Header)
	msg := &recallMsg{}
	err := c.MustBindWith(msg, binding.Form)
	if err != nil {
		log.Println(err)
		c.Abort()
	}
	notifyUser := "yangjiangdong"
	sendWechatMSG(notifyUser, fmt.Sprintf("%s", msg))
}
func keyValueToMap(s string) map[string]string {
	s1 := strings.Replace(s, "=", ":", 100)
	s2 := strings.Replace(s1, " ", "", 100)
	s3 := strings.Replace(s2, "{", "", 1)
	s4 := strings.Replace(s3, "}", "", 1)
	s5 := strings.Split(s4, ",")
	m := make(map[string]string)
	for _, v := range s5 {
		s5 := strings.Split(v, ":")
		m[s5[0]] = s5[1]
	}
	return m
}
