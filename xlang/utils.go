package xlang

import "github.com/subchen/go-xmldom"

// func printNode(node *xmldom.Node) {
// 	fmt.Printf("%s, %T, %s\n", node.Name, node.Attributes, node.Text)

//	for _, xmlattr := range node.Attributes {
//		fmt.Printf("attr[%s]=%s\n", xmlattr.Name, xmlattr.Value)
//	}
//
//		for _, child := range node.Children {
//			printNode(child)
//		}
//	}

func xmlAttrToAttributes(xmlAttributes []*xmldom.Attribute) map[string]string {
	attributesMap := map[string]string{}
	for _, xmlattr := range xmlAttributes {
		attributesMap[xmlattr.Name] = xmlattr.Value
	}
	return attributesMap
}
