package static_front

import (
	_ "embed"

	"github.com/darklab8/fl-darkcore/darkcore/core_types"
)

//go:embed custom/custom.js
var CustomJSContent string

var CustomJS core_types.StaticFile = core_types.StaticFile{
	Content:  CustomJSContent,
	Filename: "custom.js",
	Kind:     core_types.StaticFileJS,
}

//go:embed custom/table_resizer.js
var CustomResizerJSContent string

var CustomJSResizer core_types.StaticFile = core_types.StaticFile{
	Content:  CustomResizerJSContent,
	Filename: "table_resizer.js",
	Kind:     core_types.StaticFileJS,
}

//go:embed custom/filtering.js
var CustomFilteringJS string

var CustomJSFiltering core_types.StaticFile = core_types.StaticFile{
	Content:  CustomFilteringJS,
	Filename: "filtering.js",
	Kind:     core_types.StaticFileJS,
}

//go:embed common.css
var CommonCSSContent string

var CommonCSS core_types.StaticFile = core_types.StaticFile{
	Content:  CommonCSSContent,
	Filename: "common.css",
	Kind:     core_types.StaticFileCSS,
}

//go:embed custom.css
var CustomCSSContent string

var CustomCSS core_types.StaticFile = core_types.StaticFile{
	Content:  CustomCSSContent,
	Filename: "custom.css",
	Kind:     core_types.StaticFileCSS,
}
