package demo

type FView struct {
}

// Get
// @Summary 获取该主体的当前最新主流程合同
// @Description xxxxxxxaaa
// @Description asdasdop
// @Tags ygb - 合同
// @Param Authorization header string true "身份加密串"
// @Success 200 {object}  dubbo.RespGetContractCurrentByUserKey  "返回成功样式"
// @Router /a [POST]
func (s FView) Get(t string) (x string, err error) {
	return "123123", nil
}

// Get2
// @Summary 获取该主体的当前最新主流程合同2
// @Description xxxxxxxaaa22
// @Description asdasdop2
// @Tags ygb - 合同
// @Param Authorization header string true "身份加密串"
// @Success 200 {object}  dubbo.RespGetContractCurrentByUserKey  "返回成功样式"
// @Router /b [POST]
func (FView) Get2(t string) (x string, err error) {
	return "123123", nil
}
