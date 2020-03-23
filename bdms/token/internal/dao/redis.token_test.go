package dao

import (
	"context"
	"github.com/smartystreets/goconvey/convey"
	"time"

	"testing"
)

func TestRequestUploadToken(t *testing.T) {
	var (
		c        = context.Background()
		operator = "manager"
		now      = time.Now().Unix()
	)
	convey.Convey("RequestUploadToken", t, func(ctx convey.C) {
		_, _, err := d.RequestUploadToken(c, 1, operator, now, 555)
		//fmt.Println(token)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}
