package com.hazz.kotlinmvp.mvp.presenter

import com.hazz.kotlinmvp.base.BasePresenter
import com.hazz.kotlinmvp.mvp.contract.AniContract
import com.hazz.kotlinmvp.mvp.model.AniModel
import com.hazz.kotlinmvp.net.exception.ExceptionHandle

/**
 * Created by xx on 2020/2/24 0:17.
 * desc: AniPresenter
 */
class AniPresenter : BasePresenter<AniContract.View>(), AniContract.Presenter {

    var pageIdx = 0

    private val aniModel: AniModel by lazy {
        AniModel()
    }

    override fun getAnimationData() {
        checkViewAttached()
        mRootView?.showLoading()
        val disposable = aniModel.getAniData()
                .subscribe({aniList ->
                    mRootView?.apply {
                        dismissLoading()
                        pageIdx += 1
                        setAniData(aniList)
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
                .subscribe({aniList ->
                    mRootView?.apply {
                        dismissLoading()
                        pageIdx += 1
                        setMoreData(aniList)
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