// 🔌 Fiber is an Express.js inspired web framework build on 🚀 Fasthttp.
// 📌 Please open an issue if you got suggestions or found a bug!
// 🖥 https://github.com/gofiber/fiber

// 🦸 Not all heroes wear capes, thank you to some amazing people
// 💖 @valyala, @dgrr, @erikdubbelboer, @savsgio, @julienschmidt

package fiber

import (
	"sync"

	"github.com/valyala/fasthttp"
)

// Ctx : struct
type Ctx struct {
	route    *Route
	next     bool
	params   *[]string
	values   []string
	Fasthttp *fasthttp.RequestCtx
}

// Cookie : struct
type Cookie struct {
	Expire int // time.Unix(1578981376, 0)
	MaxAge int
	Domain string
	Path   string

	HTTPOnly bool
	Secure   bool
	SameSite string
}

// Ctx pool
var poolCtx = sync.Pool{
	New: func() interface{} {
		return new(Ctx)
	},
}

// Get new Ctx from pool
func acquireCtx(fctx *fasthttp.RequestCtx) *Ctx {
	ctx := poolCtx.Get().(*Ctx)
	ctx.Fasthttp = fctx
	return ctx
}

// Return Context to pool
func releaseCtx(ctx *Ctx) {
	ctx.route = nil
	ctx.next = false
	ctx.params = nil
	ctx.values = nil
	ctx.Fasthttp = nil
	poolCtx.Put(ctx)
}
