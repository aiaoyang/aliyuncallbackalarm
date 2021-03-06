package main

import (
	"fmt"
	"log"
	"regexp"
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

// Dimensions Dimensions
// 以后阿里更新接口返回数据结构用
// type Dimensions struct {
// }

func (r *recallMsg) String() string {
	// 正则匹配
	// dimensions 正则
	pat1 := "=r-|=rm-|=i-|=dbs"
	// expression 正则
	pat2 := `\w+`

	reg1 := regexp.MustCompile(pat1)
	reg2 := regexp.MustCompile(pat2)
	insType := reg1.FindString(r.Dimensions)
	expr := reg2.FindString(r.Expression)

	alarmInstance := ""
	expression := ""
	metric := ""

	if insType == "" {
		log.Printf("Dimensions: %s\n", r.Dimensions)
		alarmInstance = r.Dimensions
	} else {
		alarmInstance = Instance[keyValueToMap(r.Dimensions)["instanceId"]]
	}
	if Metric[r.MetricName] == "" {
		log.Printf("MetricName: %s\n", r.MetricName)
		metric = r.MetricName
	} else {
		metric = Metric[r.MetricName]
	}
	if expr == "" {
		log.Printf("Expression: %s\n", r.Expression)
		expression = r.Expression
	} else {
		log.Println(expr)
		expression = strings.Replace(r.Expression, "$"+expr, Expression[expr], -1)
	}

	return fmt.Sprintf("告警规则: [ %s ]\n告警状态: [ %s ]\n告警实例: [ %s ]\n实例类型: [ %s ]\n告警条件: [ %v ]\n告警指标: [ %s ]\n当前值: [ %s ]\n",
		r.AlterName,
		AlertState[r.AlertState],
		alarmInstance,
		InstanceType[insType],
		expression,
		metric,
		r.CurValue,
	)
}

var (
	// Config init config
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
	// body := c.Request.Body
	// bodyLine, err := ioutil.ReadAll(body)
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Printf("request's body is: %s\n", string(bodyLine))
	// c.Request.Body = body
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
