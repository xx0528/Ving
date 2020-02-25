package com.hazz.kotlinmvp.mvp.model

import com.hazz.kotlinmvp.mvp.model.bean.AniBean
import com.hazz.kotlinmvp.net.RetrofitManager
import com.hazz.kotlinmvp.rx.scheduler.SchedulerUtils
import io.reactivex.Observable

/**
 * Created by xx on 2020/2/25 21:32.
 * desc: AniDetailModel
 */
class AniDetailModel {
    fun requestRelatedData(id:Long): Observable<ArrayList<AniBean>> {
        return RetrofitManager.service.getRelatedAni(id)
                .compose(SchedulerUtils.ioToMain())
    }
}