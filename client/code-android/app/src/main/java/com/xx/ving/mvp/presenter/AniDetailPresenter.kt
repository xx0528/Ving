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
        var playUrl : String? = ""
        Logger.d("from -------%s", itemInfo.data?.playFrom)
        var playData = itemInfo.data?.let { getPlayData(it.playUrl, it.playNote, it.playFrom, it.playServer) }
        if (playData != null) {
            Logger.d("play data -------")
            for (urls in playData.values) {
                Logger.d("play data ------- 2")
                for (url in urls ) {
                    Logger.d("url ---- %s", url)
                    playUrl = url
                }
            }
        }
        playUrl = "https://cdn-yong.bejingyongjiu.com/20200214/3423_e7c45068/index.m3u8"
//        itemInfo.data?.playUrl?.let { mRootView?.setAni(it) }
        Logger.d("playurl --- %s", playUrl)
        if (playUrl != null) {
            mRootView?.setAni(playUrl)
        }
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

    private fun getPlayData(playUrl: String, playNote: String, playFrom: String, playServer: String) : MutableMap<String, ArrayList<String>> {
        var fromList= emptyList<String>()
//        var noteList= emptyList<String>()
        var urlList= emptyList<String>()
//        var serverList= emptyList<String>()
        Logger.d("from str === %s", playFrom)
        if (!playFrom.isNullOrEmpty()) {
            fromList = playFrom.split("$$$")
        }
//        if (!playNote.isNullOrEmpty()) {
//            noteList = playNote.split("$$$")
//        }
        if (!playUrl.isNullOrEmpty()) {
            urlList = playUrl.split("$$$")
        }
//        if (!playServer.isNullOrEmpty()) {
//            serverList = playServer.split("$$$")
//        }

        Logger.d("fromlist num -- %d", fromList.count())
        val data = mutableMapOf<String,ArrayList<String>>()
        for ((idx, from) in fromList.withIndex()){
            val dataUrlList = ArrayList<String>()
            Logger.d("urlList[idx]---- %s", urlList[idx])
            var urls = urlList[idx].split("#")
            for(url in urls) {
                Logger.d("url---- %s", url)
                var urlInfo =  url.split("$")
                dataUrlList += urlInfo[1]
            }
            data[from] = dataUrlList
        }
        return data;
    }
}