package com.hazz.ving.mvp.model

import com.hazz.ving.mvp.model.bean.AniBean
import com.hazz.ving.net.RetrofitManager
import com.hazz.ving.rx.scheduler.SchedulerUtils
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

    fun loadMoreAniData(num:Int):Observable<AniBean> {
        return RetrofitManager.service.getMoreAnimationData(num)
                .compose(SchedulerUtils.ioToMain())
    }

}