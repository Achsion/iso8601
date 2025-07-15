module github.com/Achsion/iso8601/v2

go 1.24

require github.com/stretchr/testify v1.10.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v2.0.1 // This version was accidentally released. Please use version 2.0.2 instead.
	v2.0.0 // This version was accidentally released. Please use version 2.0.2 instead.
)
