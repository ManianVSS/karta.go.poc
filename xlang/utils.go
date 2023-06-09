package xlang

import "github.com/subchen/go-xmldom"

func xmlAttrToAttributes(xmlAttributes []*xmldom.Attribute) map[string]string {
	attributesMap := map[string]string{}
	for _, xmlattr := range xmlAttributes {
		attributesMap[xmlattr.Name] = xmlattr.Value
	}
	return attributesMap
}
