package controllers

import (
	"github.com/astaxie/beego"
	"net/url"
	"player/structs"
	//"github.com/buger/jsonparser"
	"encoding/json"
	//"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type QueryController struct {
	beego.Controller
}

func GetIrc(id string)string{

url:="https://v1.itooi.cn/kuwo/lrc?id="+id
resp,_:=http.Get(url)
defer resp.Body.Close()
body,_:=ioutil.ReadAll(resp.Body)
return string(body)
}
func GetQueryResult(keyword string,page string)[]byte{
	word:=url.QueryEscape(keyword)
	queryword:="https://v1.itooi.cn/kuwo/search?keyword="+word+"&type=song&pageSize=20&page="+page
	resp,err:=http.Get(queryword)
	if err!=nil {
		beego.Info("获取失败")
	}
	defer resp.Body.Close()
	body,_:=ioutil.ReadAll(resp.Body)
	return body
}
func (c *QueryController)Get(){
	keyword:=c.GetString("keyword")
	page:=c.GetString("page")
	//beego.Info(keyword)
	body:=GetQueryResult(keyword,page)
	var result structs.QueryResult
	err:=json.Unmarshal(body,&result)
	if err!=nil {
		beego.Error(err)
	}
	c.Data["json"]=result.Data
	c.ServeJSON()
	//c.TplName="index.html"
}