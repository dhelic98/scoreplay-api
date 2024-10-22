package enum

var ALLOWED_FILE_EXTENSION map[string]bool = map[string]bool{
	".jpeg": true,
	".jpg":  true,
	".svg":  true,
	".png":  true,
}

func IsAllowed(fileExtension string) bool {
	_, ok := ALLOWED_FILE_EXTENSION[fileExtension]
	return ok
}
