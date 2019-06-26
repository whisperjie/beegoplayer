var a=0
function getResults(op) {
    if (op=="next") {
    a++
    }else{
        a=1
        $("#fafa tr:not(:first)").empty("");
    }
    var keyword=$("#kaka").val();
    if (keyword==""){
        alert("未输入关键词");
        return
    }
    var link="query?keyword="+keyword+"?page="+a;
    var xhr=new XMLHttpRequest();

    xhr.open("GET",link,true);
    xhr.onreadystatechange=function () {
        if (xhr.readyState==4&&xhr.status==200){
            res=xhr.responseText;
            var o=JSON.parse(res);
            for (var i in o){
                var id=o[i].MUSICRID.substring(6,o[i].MUSICRID.length)
                console.log(id)
                mylink="/play?id="+id
                index=parseInt(i)+10*a-10
                var trstr=`<tr>
                                <td>${index}</td>
                                <td>${o[i].ARTIST}</td>
                                <td>${o[i].SONGNAME}</td>
                                <td><a href='${mylink}'>播放</a></td>
                            </tr>`;
                $("#fafa").append(trstr);
            }
        }
    }
    xhr.send();
}