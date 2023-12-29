package getFieldGogo

import (
    "context"
    "time"

    cUtils "github.com/byted-apaas/server-common-go/utils"
    "github.com/byted-apaas/server-sdk-go/application"
)

type Params struct{}

type Result struct {
    SubObjectField   interface{} `json:"subObjectField"`
    BigintField      interface{} `json:"bigintField"`
    DecimalField     interface{} `json:"decimalField"`
    RegionField      interface{} `json:"regionField"`
    RollupCountField interface{} `json:"rollupCountField"`
}

func Handler(ctx context.Context, params *Params) (*Result, error) {
    // 日志功能
    application.GetLogger(ctx).Infof("%s 函数开始执行", time.Now().Format("2006-01-02 15:04:05.999"))

    // 在这里补充业务代码
    var subObjectField interface{}
    err := application.Metadata.Object("a").GetField(ctx, "subObject", &subObjectField)
    if err != nil {
        panic(err)
    }
    application.GetLogger(ctx).Infof("subObjectField: %s", cUtils.ToString(subObjectField))

    var bigintField interface{}
    err = application.Metadata.Object("a").GetField(ctx, "bigint", &bigintField)
    if err != nil {
        panic(err)
    }
    application.GetLogger(ctx).Infof("bigintField: %s", cUtils.ToString(bigintField))

    var decimalField interface{}
    err = application.Metadata.Object("a").GetField(ctx, "decimal", &decimalField)
    if err != nil {
        panic(err)
    }
    application.GetLogger(ctx).Infof("decimalField: %s", cUtils.ToString(decimalField))

    var regionField interface{}
    err = application.Metadata.Object("a").GetField(ctx, "region", &regionField)
    if err != nil {
        panic(err)
    }
    application.GetLogger(ctx).Infof("a: %s", cUtils.ToString(regionField))

    return &Result{
        SubObjectField:   subObjectField,
        BigintField:      bigintField,
        DecimalField:     decimalField,
        RegionField:      regionField,
    }, nil
}