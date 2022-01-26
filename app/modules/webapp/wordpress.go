package webapp

type wordpressApp struct {
	Webapp
}

// func NewWordpressApp() Webapp {
// 	return &wordpressApp{}
// }

func (wordpressApp *wordpressApp) Install() bool {
	return true
}
func (wordpressApp *wordpressApp) Backup() bool {
	return true
}
