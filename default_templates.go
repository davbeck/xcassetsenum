package main

var defaultSwiftTemplate = `import UIKit

// this class is used to find the bundle the file is in
private class {{.EnumName}}Class: NSObject {}

extension UIImage {
	{{.AccessControl}} enum {{.EnumName}}: String {
		{{- range $key, $value := .Assets }}
		case {{ $key }} = "{{ $value }}"
		{{- end }}
	}
    
	{{.AccessControl}} convenience init!({{.EnumInitName}}: {{.EnumName}} compatibleWithTraitCollection: UITraitCollection? = nil) {
		self.init(named: {{.EnumInitName}}.rawValue, inBundle: NSBundle(forClass: {{.EnumName}}Class.self), compatibleWithTraitCollection: compatibleWithTraitCollection)
	}
}
`
