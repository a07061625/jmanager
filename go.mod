module jmanager

go 1.12

require (
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/gocolly/colly v1.2.0
	github.com/json-iterator/go v1.1.6
	golang.org/x/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65
	golang.org/x/text v0.3.2
)

replace (
	golang.org/x/crypto v0.0.0-20190605123033-f99c8df09eb5 => github.com/golang/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65 => github.com/golang/net v0.0.0-20190606173856-1492cefac77f
	golang.org/x/text v0.3.2 => github.com/golang/text v0.3.2
)
