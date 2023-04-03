package resolve

import (
	"fmt"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestResolveTemplate(t *testing.T) {
	template := "My name is ${person.name} and I work at ${company.location}"
	jsonStr := "{\"person\":{\"name\":\"rick\"},\"company\":{\"location\":\"shanghai\"}}"
	res, _ := ResolveTemplate(template, []byte(jsonStr))
	fmt.Println(res)
	assert.DeepEqual(t, res, "My name is rick and I work at shanghai")
}
