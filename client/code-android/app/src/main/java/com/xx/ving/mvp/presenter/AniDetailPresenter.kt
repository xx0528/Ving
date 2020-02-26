package com.hazz.ving.mvp.presenter

import com.hazz.ving.MyApplication
import com.hazz.ving.base.BasePresenter
import com.hazz.ving.mvp.contract.AniDetailContract
import com.hazz.ving.mvp.model.AniDetailModel
import com.hazz.ving.mvp.model.bean.AniBean
import com.hazz.ving.utils.NetworkUtil

/**
 * Created by xx on 2020/2/25 21:29.
 * desc: AniDetailPresenter
 */
class AniDetailPresenter : BasePresenter<AniDetailContract.View>(), AniDetailContract.Presenter {

    private val aniDetailModel: AniDetailModel by lazy {
        AniDetailModel()
    }

    override fun loadAniInfo(itemInfo: AniBean.AItem) {

        val netType = NetworkUtil.isWifi(MyApplication.context)
        // 检测是否绑定 View
        checkViewAttached()

        itemInfo.data?.playUrl?.let { mRootView?.setAni(it) }

        //设置背景
//        val backgroundUrl = itemInfo.data.cover.blurred + "/thumbnail/${DisplayManager.getScreenHeight()!! - DisplayManager.dip2px(250f)!!}x${DisplayManager.getScreenWidth()}"
//        backgroundUrl.let { mRootView?.setBackground(it) }

        mRootView?.setAniInfo(itemInfo)
    }

    override fun requestRelatedAni(id: Long) {

    }
}