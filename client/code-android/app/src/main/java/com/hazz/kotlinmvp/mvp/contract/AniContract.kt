package com.hazz.kotlinmvp.mvp.contract

import com.hazz.kotlinmvp.base.IBaseView
import com.hazz.kotlinmvp.base.IPresenter
import com.hazz.kotlinmvp.mvp.model.bean.AniBean

/**
 * Create by xx on 2020/02/23.
 * desc: 分类动画
 */
interface AniContract {
    interface View : IBaseView {
        /**
         * 显示动画信息
         */
        fun showAnimation(aniList: ArrayList<AniBean>)

        /**
         * 显示错误信息
         */
        fun showError(errorMsg:String, errorCode:Int)
    }

    interface Presenter:IPresenter<View> {
        /**
         * 获取动画信息
         */
        fun getAnimationData()
    }
}