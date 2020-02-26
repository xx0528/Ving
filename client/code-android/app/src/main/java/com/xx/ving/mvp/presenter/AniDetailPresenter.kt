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

    override fun loadAniInfo(itemInfo: AniBean) {

        val netType = NetworkUtil.isWifi(MyApplication.context)
        // 检测是否绑定 View
        checkViewAttached()

        mRootView?.setAni(itemInfo.bgPicture)

        //设置背景
//        val backgroundUrl = itemInfo.data.cover.blurred + "/thumbnail/${DisplayManager.getScreenHeight()!! - DisplayManager.dip2px(250f)!!}x${DisplayManager.getScreenWidth()}"
//        backgroundUrl.let { mRootView?.setBackground(it) }

        mRootView?.setAniInfo(itemInfo)
    }

    override fun requestRelatedAni(id: Long) {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}