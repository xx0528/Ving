package com.xx.ving.mvp.presenter

import com.orhanobut.logger.Logger
import com.xx.ving.MyApplication
import com.xx.ving.base.BasePresenter
import com.xx.ving.mvp.contract.AniDetailContract
import com.xx.ving.mvp.model.AniDetailModel
import com.xx.ving.mvp.model.bean.AniBean
import com.xx.ving.net.exception.ExceptionHandle
import com.xx.ving.utils.DisplayManager
import com.xx.ving.utils.NetworkUtil
import java.util.*
import kotlin.collections.ArrayList

/**
 * Created by xx on 2020/2/25 21:29.
 * desc: AniDetailPresenter
 */
class AniDetailPresenter : BasePresenter<AniDetailContract.View>(), AniDetailContract.Presenter {

    private val aniDetailModel: AniDetailModel by lazy {
        AniDetailModel()
    }

    override fun loadAniInfo(itemInfo: AniBean.AItem) {

//        val netType = NetworkUtil.isWifi(MyApplication.context)
        // 检测是否绑定 View
        checkViewAttached()
        val playUrl = itemInfo.data?.playUrl?:""
        var urlList= emptyList<String>()
        urlList = playUrl.split("#")
        var url = urlList[0].split("$")[1]
//        itemInfo.data?.playUrl?.let { mRootView?.setAni(it) }
//        Logger.d("url -- %s", url)
        mRootView?.setAni(url)
        //设置背景
        val backgroundUrl = (itemInfo.data?.cover?.blurred ?: "") + "/thumbnail/${DisplayManager.getScreenHeight()!! - DisplayManager.dip2px(250f)!!}x${DisplayManager.getScreenWidth()}"
        backgroundUrl.let { mRootView?.setBackground(it) }

        mRootView?.setAniInfo(itemInfo)
    }

    override fun requestRelatedAni(id: Long) {
        mRootView?.showLoading()
        val disposable = aniDetailModel.requestRelatedAniData(id)
                .subscribe({aniData ->
                    mRootView?.apply {
                        dismissLoading()
                        setRecentRelatedAni(aniData.itemList)
                    }
                }, { t ->
                    mRootView?.apply {
                        dismissLoading()
                        setErrorMsg(ExceptionHandle.handleException(t))
                    }
                })

        addSubscription(disposable)
    }
}