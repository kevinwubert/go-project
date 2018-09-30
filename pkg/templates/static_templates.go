package templates

var staticsTemplates = templates{
	"hello-world": template{
		name: "hello-world",
		rootDir: &directory{
			name: "",
			files: []file{
				file{
					name: "Makefile",
					data: []byte{97, 108, 108, 58, 32, 98, 117, 105, 108, 100, 10, 10, 98, 117, 105, 108, 100, 58, 10, 9, 103, 111, 32, 98, 117, 105, 108, 100, 32, 45, 111, 32, 98, 105, 110, 47, 109, 97, 105, 110, 32, 99, 109, 100, 47, 42, 46, 103, 111},
				},
			},
			dirs: []*directory{
				&directory{
					name: "cmd",
					files: []file{
						file{
							name: "main.go",
							data: []byte{112, 97, 99, 107, 97, 103, 101, 32, 109, 97, 105, 110, 10, 10, 105, 109, 112, 111, 114, 116, 32, 102, 109, 116},
						},
					},
					dirs: []*directory{},
				},
			},
		},
	},
}
