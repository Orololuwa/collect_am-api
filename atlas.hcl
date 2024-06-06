data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./src/models",
    "--dialect", "postgres",
  ]
}

env "gorm" {
  migration {
    dir = "file://./migrations-atlas"
  }
  src = data.external_schema.gorm.url
  dev = "postgres://orololuwa:@localhost:5432/collect_am_api_clean?sslmode=disable"
}
