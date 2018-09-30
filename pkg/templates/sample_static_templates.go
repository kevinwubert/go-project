package templates

var staticTemplates = templates{
	"hello-world": template{
		name: "hello-world",
		rootDir: &directory{
			name: "",
			files: []file{
				file{
					name: "Makefile",
					data: []byte("all: build\n\nbuild:\n\tgo build -o bin/main cmd/*.go"),
				},
			},
			dirs: []*directory{
				&directory{
					name: "cmd",
					files: []file{
						file{
							name: "main.go",
							data: []byte("package main\n\nimport fmt"),
						},
					},
				},
			},
		},
	},
}
