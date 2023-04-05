package main
import(
	"fmt"
	"os"
	"flag"
	c "github.com/hindu-muppala/YamlJsonConverter"
)
func exit(x string ) bool{ //checking file exit or not
	_,err:=os.Stat(x)
		
		if os.IsNotExist(err)  {
			fmt.Printf("%v is not exit",x)
			return false
		}else {
			return true
		}
}
func ConvertToYaml(path1 string,path2 string){
	x:=c.ConvertYaml(path1,path2)
	if x!=0 {
		fmt.Print("conversion failed")
	}else{
		fmt.Print("conversion sucessfull")
	}
}
func ConvertToJson(path1 string,path2 string){
	x:=c.ConvertJson(path1,path2)
	if x!=0 {
		fmt.Print("conversion failed")
	}else{
		fmt.Print("conversion sucessfull")
	}
}
func Convert(path1 string,path2 string,j bool,y bool) { //to convert
	
	if exit(path1) {
		if (j && y){
			fmt.Print("Both conversion not possible")
			return
		}
		if (j){
			ConvertToJson(path1 ,path2 )
		}
		if (y) {
			ConvertToYaml(path1 ,path2 )
		}
	}
	return 
}

func main(){
	var y,j bool
	var convert string
	//flag.BoolVar(&t,"t",false,"display in terminal")
	flag.BoolVar(&y,"y",false,"convert to yaml")
	flag.BoolVar(&j,"j",false,"convert to json")
	flag.StringVar(&convert,"convert","","conversion")
	flag.Parse()
	values:=flag.Args()
	if len(values)==0 {
		fmt.Print("No arguments provided")
		os.Exit(1)
	}else if len(values)==2 {
		Convert(values[0],values[1],j,y)
	}else{
		fmt.Print("command takes two arguments")
	}
}




