package converter_test
 
import (
	"testing"
	"fmt"
	//"os/exec"
	//"log"
	converter "github.com/hindu-muppala/YamlJsonConverter"
)

func TestAdd(t *testing.T){
	x:=c.convertJson("./cmd/testfiles/test2.yaml","./cmd/testfiles/test1.json")
	y:=c.convertYaml("./cmd/testfiles/test.json","./cmd/testfiles/test3.yaml")
	if x!=0 || y!=0 {
		t.Errorf("eroor")
	}else{
		fmt.Print("conversion successful")
	}

}