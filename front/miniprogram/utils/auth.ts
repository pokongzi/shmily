interface LoginResult {
  token: string
  expires: number 
}

class Auth {
  private static instance: Auth
  private token: string = ''
  private expires: number = 0
  private retryTimes: number = 0 // 登录重试次数
  private maxRetryTimes: number = 3 // 最大重试次数
  
  private constructor() {
    // 从存储中恢复登录态
    const token = wx.getStorageSync('token')
    const expires = wx.getStorageSync('expires') 
    if(token && expires) {
      this.token = token
      this.expires = expires
    }
  }

  public static getInstance(): Auth {
    if (!Auth.instance) {
      Auth.instance = new Auth()
    }
    return Auth.instance
  }

  // 检查是否已登录
  public isLoggedIn(): boolean {
    return !!this.token && this.expires > Date.now()
  }

  // 获取 token
  public getToken(): string {
    return this.token
  }

  // 保存登录信息
  private saveLoginInfo(result: LoginResult) {
    this.token = result.token
    this.expires = result.expires
    wx.setStorageSync('token', result.token)
    wx.setStorageSync('expires', result.expires)
    // 重置重试次数
    this.retryTimes = 0
  }

  // 清除登录信息
  public clearLoginInfo() {
    this.token = ''
    this.expires = 0
    wx.removeStorageSync('token')
    wx.removeStorageSync('expires') 
  }

  // 登录
  public async login(): Promise<boolean> {
    try {
      if (this.retryTimes >= this.maxRetryTimes) {
        throw new Error('登录重试次数超限')
      }

      // 获取登录code
      const {code} = await wx.login()
      
      // 请求后端登录
      const res = await wx.request({
        url: 'http://localhost:3000/login',
        method: 'POST',
        data: {
          code
        }
      })

      if (res.statusCode !== 200) {
        throw new Error('登录请求失败')
      }

      // 保存登录信息
      this.saveLoginInfo(res.data as LoginResult)
      return true

    } catch (err) {
      console.error('登录失败:', err)
      // 登录失败重试
      this.retryTimes++
      if (this.retryTimes < this.maxRetryTimes) {
        return this.login()
      }
      return false
    }
  }

  // 检查并自动登录
  public async checkAndLogin(): Promise<boolean> {
    if(this.isLoggedIn()) {
      return true
    }
    return this.login()
  }
}

export default Auth.getInstance() 