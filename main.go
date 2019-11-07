package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"
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

var (
	// Config All config
	Config = viper.New()
	// Users wechat notify users
	Users string
)

func init() {
	log.SetFlags(log.Llongfile | log.Ldate)
	Config.SetConfigFile("config.yaml")
	Config.SetConfigType("yaml")
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	Users = strings.Join(Config.GetStringSlice("wechat.users"), "|")
	log.Printf("users: %s\n", Users)

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
	sendWechatMSG(Users, fmt.Sprintf("%s", msg))
}
func keyValueToMap(s string) map[string]string {
	// e.g: %7Bport%3D10080%2C+protocol%3Dtcp%2C+userId%3D1854764472318598%2C+instanceId%3Dlb-t4n4zszmjw0o8to3nyedp%7D&expression=%24Maximum%3C10000}
	// src: {port=10080, protocol=tcp, userId=1854764472318598, instanceId=lb-t4n4zszmjw0o8to3nyedp}
	// dst: port=10080,protocol=tcp,userId=1854764472318598,instanceId=lb-t4n4zszmjw0o8to3nyedp
	s1 := strings.Replace(s, " ", "", -1)
	s2 := strings.Replace(s1, "{", "", -1)
	s3 := strings.Replace(s2, "}", "", -1)
	s4 := strings.Split(s3, ",")
	m := make(map[string]string)
	for _, v := range s4 {
		s5 := strings.Split(v, "=")
		m[s5[0]] = s5[1]
	}
	return m
}
