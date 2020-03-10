package com.xx.ving.mvp.model

import com.xx.ving.mvp.model.bean.AniBean
import com.xx.ving.net.RetrofitManager
import com.xx.ving.rx.scheduler.SchedulerUtils
import io.reactivex.Observable

/**
 * Created by xx on 2020/2/24 0:20.
 * desc: AniModel
 */
class AniModel {
    /**
     * 获取动画信息
     */
    fun getAniData(): Observable<AniBean> {
        return RetrofitManager.service.getAnimationData()
                .compose(SchedulerUtils.ioToMain())
    }

    fun loadMoreAniData(p:Int):Observable<AniBean> {
        return RetrofitManager.service.getMoreAnimationData(10, p)
                .compose(SchedulerUtils.ioToMain())
    }

}