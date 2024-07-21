package builder

import (
	"bytes"
	"context"

	"github.com/darklab8/fl-darkstat/darkstat/common/types"

	"github.com/a-h/templ"
	"github.com/darklab8/go-utils/utils/utils_filepath"
	"github.com/darklab8/go-utils/utils/utils_types"
)

type Component struct {
	pagepath   utils_types.FilePath
	templ_comp templ.Component
}

func NewComponent(
	pagepath utils_types.FilePath,
	templ_comp templ.Component,
) *Component {
	return &Component{
		pagepath:   pagepath,
		templ_comp: templ_comp,
	}
}

type WriteResult struct {
	realpath utils_types.FilePath
	bytes    []byte
}

func (h *Component) Write(gp Params) WriteResult {
	buf := bytes.NewBuffer([]byte{})

	// gp.Pagepath = string(h.pagepath)

	realpath := utils_filepath.Join(gp.GetBuildPath(), h.pagepath)

	h.templ_comp.Render(context.WithValue(context.Background(), types.GlobalParamsCtxKey, gp), buf)

	// Usage of gohtml is not obligatory, but nice touch simplifying debugging view.
	return WriteResult{
		realpath: realpath,
		bytes:    buf.Bytes(),
	}
}
