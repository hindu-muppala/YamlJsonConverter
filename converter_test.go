package converter_test
 
import (
	"testing"
	"fmt"
	//"os/exec"
	//"log"
	c "github.com/hindu-muppala/YamlJsonConverter"
)

func TestAdd(t *testing.T){
	x:=c.ConvertJson("./cmd/testfiles/testfiles/test2.yaml","./cmd/testfiles/testfiles/test1.json")
	y:=c.ConvertYaml("./cmd/testfiles/testfiles/test.json","./cmd/testfiles/testfiles/test3.yaml")
	if x!=0 || y!=0 {
		t.Errorf("check the file directory")
	}
}