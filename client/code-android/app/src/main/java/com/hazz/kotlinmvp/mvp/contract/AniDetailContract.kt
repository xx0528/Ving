package com.hazz.kotlinmvp.mvp.contract

import com.hazz.kotlinmvp.base.IBaseView
import com.hazz.kotlinmvp.base.IPresenter
import com.hazz.kotlinmvp.mvp.model.bean.AniBean

/**
 * Created by xx on 2020/2/25 19:34.
 * desc: AniDetailContract
 */
interface AniDetailContract {
    interface View : IBaseView {
        /**
         * 设置动画播放源
         */
        fun setAni(url: String)

        /**
         * 设置动画信息
         */
        fun setAniInfo(itemInfo: AniBean)

        /**
         * 设置背景
         */
        fun setBackground(url: String)

        /**
         * 设置最新相关视频
         */
        fun setRecentRelatedAni(itemList: ArrayList<AniBean>)

        /**
         * 设置错误信息
         */
        fun setErrorMsg(errorMsg: String)
    }

    interface Presenter : IPresenter<View> {
        /**
         * 加载动画信息
         */
        fun loadAniInfo(itemInfo: AniBean)

        /**
         * 请求相关的动画数据
         */
        fun requestRelatedAni(id: Long)
    }
}