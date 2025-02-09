// app.ts
import auth from './utils/auth'

App<IAppOption>({
  globalData: {},
  onLaunch() {
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    auth.checkAndLogin().then(isLoggedIn => {
      if (!isLoggedIn) {
        wx.showToast({
          title: '登录失败',
          icon: 'none'
        })
      }
    })
  },
})