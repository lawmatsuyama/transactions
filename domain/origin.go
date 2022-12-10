package domain

// OriginChannel represents origin channel
type OriginChannel string

const (
	DesktopWeb    OriginChannel = "desktop-web"
	MobileAndroid OriginChannel = "mobile-android"
	MobileIos     OriginChannel = "mobile-ios"
)

// IsValid check if origin channel is valid
func (origin OriginChannel) IsValid() bool {
	switch origin {
	case DesktopWeb, MobileAndroid, MobileIos:
		return true
	default:
		return false
	}
}
