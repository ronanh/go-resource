package go-resource

func check_resource_dir() {
	if resource_dir == "" {
		log.Println("RESOURCE_DIR not set")
		os.Exit(100)
	}
	_, err := os.Stat(resource_dir)
	if nil != err {
		log.Println("RESOURCE_DIR Directory not found: " + resource_dir)
		log.Println(err)
		os.Exit(101)
	}
}
