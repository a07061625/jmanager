module jmanager

go 1.12

require (
    github.com/json-iterator/go v1.1.6
    github.com/axgle/mahonia
    github.com/gocolly/colly v1.2.0
    github.com/PuerkitoBio/goquery v1.5.0
)

replace (
    golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
    golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
    golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)