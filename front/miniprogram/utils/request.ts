import auth from './auth'

interface RequestOptions extends WechatMiniprogram.RequestOption {
  noAuth?: boolean // 是否需要鉴权
  retry?: boolean // 是否允许重试
}

const BASE_URL = 'http://localhost:3000'
const MAX_RETRY_TIMES = 3

// 创建请求实例
const request = async (options: RequestOptions) => {
  // 处理 URL
  if (!options.url.startsWith('http')) {
    options.url = BASE_URL + options.url
  }

  // 非鉴权接口跳过 token 检查
  if (!options.noAuth) {
    // 检查登录状态
    const isLoggedIn = await auth.checkAndLogin()
    if (!isLoggedIn) {
      throw new Error('登录失败')
    }

    // 添加 token
    options.header = {
      ...options.header,
      'Authorization': `Bearer ${auth.getToken()}`
    }
  }

  let retryTimes = 0
  const maxRetryTimes = options.retry === false ? 0 : MAX_RETRY_TIMES

  while (retryTimes <= maxRetryTimes) {
    try {
      const res = await wx.request(options)
      
      // 处理业务错误
      if (res.statusCode === 401) {
        // token 过期,清除登录信息
        auth.clearLoginInfo()
        // 重新登录
        const isLoggedIn = await auth.login()
        if (isLoggedIn) {
          // 重试请求
          return request({
            ...options,
            retry: false // 防止无限重试
          })
        } else {
          throw new Error('重新登录失败')
        }
      }
      
      if (res.statusCode !== 200) {
        throw new Error(`请求失败: ${res.statusCode}`)
      }

      return res.data

    } catch (err) {
      retryTimes++
      
      // 达到最大重试次数,抛出错误
      if (retryTimes > maxRetryTimes) {
        // 显示错误提示
        wx.showToast({
          title: err.message || '请求失败',
          icon: 'none'
        })
        throw err
      }

      // 延迟重试
      await new Promise(resolve => setTimeout(resolve, 1000 * retryTimes))
    }
  }
}

export default request 