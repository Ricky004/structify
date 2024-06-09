package cmd

var projectStructures = map[string][]string{
    "go": {
        "cmd",
        "pkg",
        "internal",
        "api",
        "web",
        "scripts",
        "build",
        "deploy",
    },
    "backend": {
        "controllers",
		"constants",
        "models",
        "routes",
        "middlewares",
        "utils",
        "config",
        "tests",
    },
    "cli": {
        "cmd",
        "internal",
        "pkg",
        "docs",
        "tests",
    },
}