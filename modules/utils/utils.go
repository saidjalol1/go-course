package utils



func ExportedFunc() string {
	return "This is an exported function from utils package"
}

func unexportedFunc() string {
	return "This is an unexported function from utils package"
}