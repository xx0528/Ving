package com.xx.ving.mvp.presenter

import com.xx.ving.base.BasePresenter
import com.xx.ving.mvp.contract.AniContract
import com.xx.ving.mvp.model.AniModel
import com.xx.ving.net.exception.ExceptionHandle

/**
 * Created by xx on 2020/2/24 0:17.
 * desc: AniPresenter
 */
class AniPresenter : BasePresenter<AniContract.View>(), AniContract.Presenter {

    var pageIdx = 1

    private val aniModel: AniModel by lazy {
        AniModel()
    }

    override fun getAnimationData() {
        checkViewAttached()
        mRootView?.showLoading()
        val disposable = aniModel.getAniData()
                .subscribe({aniData ->
                    mRootView?.apply {
                        dismissLoading()
                        pageIdx += 1
                        setAniData(aniData.itemList)
                    }
                }, {t ->
                    mRootView?.apply {
                        showError(ExceptionHandle.handleException(t), ExceptionHandle.errorCode)
                    }
                })

        addSubscription(disposable)
    }

    /**
     * 加载更多
     */

    override fun loadMoreData() {
        val disposable = aniModel.loadMoreAniData(pageIdx)
                .subscribe({aniData ->
                    mRootView?.apply {
                        dismissLoading()
                        pageIdx += 1
                        setMoreData(aniData.itemList)
                    }
                }, {t ->
                    mRootView?.apply {
                        showError(ExceptionHandle.handleException(t), ExceptionHandle.errorCode)
                    }
                })

        if (disposable != null)
            addSubscription(disposable)
    }

}