package com.xx.ving.mvp.contract

import com.xx.ving.base.IBaseView
import com.xx.ving.base.IPresenter
import com.xx.ving.mvp.model.bean.AniBean

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
        fun setAniInfo(itemInfo: AniBean.AItem)

        /**
         * 设置背景
         */
        fun setBackground(url: String)

        /**
         * 设置最新相关视频
         */
        fun setRecentRelatedAni(itemList: ArrayList<AniBean.AItem>)

        /**
         * 设置错误信息
         */
        fun setErrorMsg(errorMsg: String)
    }

    interface Presenter : IPresenter<View> {
        /**
         * 加载动画信息
         */
        fun loadAniInfo(itemInfo: AniBean.AItem)

        /**
         * 请求相关的动画数据
         */
        fun requestRelatedAni(id: Long)
    }
}