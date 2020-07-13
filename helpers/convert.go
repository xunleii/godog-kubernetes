package helpers

import (
	"encoding/json"
	"strings"

	"gopkg.in/yaml.v3"
)

// YamlToJson converts naively YAML string to JSON []byte.
func YamlToJson(in string) ([]byte, error) {
	var x map[string]interface{}
	if err := yaml.Unmarshal([]byte(in), &x); err != nil {
		return nil, err
	}
	return json.Marshal(x)
}

// SanitizeJsonPatch replace all '/' and '~' in the given JsonPath expression.
func SanitizeJsonPatch(expr string) string {
	return strings.ReplaceAll(strings.ReplaceAll(expr, "~", "~0"), "/", "~1")
}
