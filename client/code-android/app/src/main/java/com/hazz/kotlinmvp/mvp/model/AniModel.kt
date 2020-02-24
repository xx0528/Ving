package com.hazz.kotlinmvp.mvp.model

import com.hazz.kotlinmvp.mvp.model.bean.AniBean
import com.hazz.kotlinmvp.net.RetrofitManager
import com.hazz.kotlinmvp.rx.scheduler.SchedulerUtils
import io.reactivex.Observable

/**
 * Created by xx on 2020/2/24 0:20.
 * desc: AniModel
 */
class AniModel {
    /**
     * 获取动画信息
     */
    fun getAniData(): Observable<ArrayList<AniBean>> {
        return RetrofitManager.service.getAnimationData()
                .compose(SchedulerUtils.ioToMain())
    }
}