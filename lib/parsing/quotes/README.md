# how to convert json to Go struct
cat file.json | grep -v '#' | gojson -forcefloats -name StructName -pkg packagename > packagename.go
