package converter
import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"encoding/json"	
)
func ConvertJson(path string,path2 string) int{
	b, not_found:=ioutil.ReadFile(path)
	if not_found!=nil {
		fmt.Print("file is not there",path)
		return 1
	}
	h:=make(map[interface{}]interface{})
	e:=yaml.Unmarshal(b,h)
	if e!=nil {
		fmt.Print(e)
		return 2
	}
 	h1:= make(map[string]interface{})
 	for key,val:= range h{
	compare:= fmt.Sprintf("%T",val)
	if "map[interface {}]interface {}"== compare {
	h1[key.(string)]=converting(val)
	
}else{
	d:=key.(string)
	h1[d]=val
	}

 }
 u,err:= json.Marshal(h1)
 if err!=nil {
	fmt.Print(err)
 }
er:=ioutil.WriteFile(path2,u,0)// automatic it creates the file
if er!=nil {
 fmt.Print(er)
 return 2
}
return 0
}
func converting (val interface{}) interface{}{
	compare:= fmt.Sprintf("%T",val)
	if "map[interface {}]interface {}"!= compare {
		return val
	}
		val2:=val.(map[interface {}]interface {})
		g:=make(map[string]interface{},len(val2))
	for x,y :=range val2{
		x:=fmt.Sprintf("%v",x)
		g[x]=converting(y)
	}
	return g
}
func ConvertYaml(path1 string,path2 string) int{
	b,err:=ioutil.ReadFile(path1)
	if err!=nil{
		fmt.Print(err)
		return 1
	}
	h:=make(map[string]interface{})
	err2:=json.Unmarshal(b,&h)
	if err2!=nil{
		fmt.Print(err2)
		return 2
	}
	b, err3:=  yaml.Marshal(&h)
	if err3 != nil {
		fmt.Print(err3)
		return 2
	}
	err4:=ioutil.WriteFile(path2,b,0)
	if err4 != nil{
		fmt.Print(err4)
		return 2
	}
	return 0
}