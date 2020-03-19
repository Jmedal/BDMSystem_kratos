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
		bucket   = "exp"
		operator = "manager"
		now     = time.Now().Unix()
	)
	convey.Convey("RequestUploadToken", t, func(ctx convey.C) {
		_, err := d.RequestUploadToken(c, bucket, operator, now)
		//fmt.Println(token)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}
