// 通过 NPM dependencies 成功安装 NPM 包后此处可引入使用
// 如安装 linq 包后就可以引入并使用这个包
// const linq = require("linq");

/**
 * @param {Params}  params     自定义参数
 * @param {Context} context    上下文参数，可通过此参数下钻获取上下文变量信息等
 * @param {Logger}  logger     日志记录器
 *
 * @return 函数的返回数据
 */
module.exports = async function (params, context, logger) {
  // 日志功能
  // logger.info(`${new Date()} 函数开始执行`);

  // 在这里补充业务代码
  // 在对象中创建一条记录
  const ss = await application.data.object("object_imya").create({
    "email": "test@bytedance.com",
    "phone": new application.constants.type.Phone("18610595108"),
  })
  logger.info("函数日志1：", ss)
	return {
    output1: {
			_id: ss._id
		}
	}
}