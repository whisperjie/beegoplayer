package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"player/structs"
)

type PlayController struct {
	beego.Controller
}

func GetSongdetail(id string) structs.Song  {
	str:="https://v1.itooi.cn/kuwo/song?id="+id
	resp,err:=http.Get(str)
	if err!=nil {
		beego.Error("获取失败")
	}
	defer resp.Body.Close()
	body,_:=ioutil.ReadAll(resp.Body)
	mysong :=structs.Song{}
	mysong.Artist,_=jsonparser.GetString(body,"data","artist")
	mysong.SongId,_=jsonparser.GetString(body,"data","id")
	mysong.AlbumId,_=jsonparser.GetString(body,"data","albumId")
	mysong.Name,_=jsonparser.GetString(body,"data","songName")
	mysong.Pic,_=jsonparser.GetString(body,"data","pic")
	mysong.SongTimeMinutes,_=jsonparser.GetString(body,"data","songTimeMinutes")
	return mysong
}

func (c *PlayController) Get() {
	id:=c.Input().Get("id")
	//beego.Info(id)
	str:="https://v1.itooi.cn/kuwo/url?id="+id+"&quality=192&isRedirect=0"
	resp,err:=http.Get(str)
	if err!=nil {
		beego.Error("获取失败")
	}
	defer resp.Body.Close()
	body,_:=ioutil.ReadAll(resp.Body)
	//beego.Info(string(body))
	mymap:=make(map[string]interface{})
	err=json.Unmarshal(body,&mymap)
	if err!=nil {
		beego.Error("地址获取失败")
	}

	song:=GetSongdetail(id)
	//beego.Info(song)
	c.Data["irc"]=GetIrc(id)
	c.Data["song"]=song
	c.Data["url"]=mymap["data"]
	c.TplName = "play.html"
}
