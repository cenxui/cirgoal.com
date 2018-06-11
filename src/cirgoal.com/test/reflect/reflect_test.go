package reflect_test

import (
	"testing"
	"reflect"
	"github.com/aws/aws-sdk-go/service/cloud9"
)

func Test(t *testing.T) {
	t.Log(reflect.TypeOf("S"))
	t.Log(reflect.TypeOf(10))
	t.Log(reflect.TypeOf(10.0))
	t.Log(reflect.TypeOf(070))
	t.Log(reflect.TypeOf(0xff))
	t.Log(reflect.TypeOf(cloud9.Environment{}))


}


func TestAddress(t *testing.T) {
	t.Log(reflect.TypeOf("S").Method(0).Name)
}