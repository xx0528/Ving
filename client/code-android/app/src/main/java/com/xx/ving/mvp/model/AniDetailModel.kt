package com.xx.ving.mvp.model

import com.xx.ving.mvp.model.bean.AniBean
import com.xx.ving.net.RetrofitManager
import com.xx.ving.rx.scheduler.SchedulerUtils
import io.reactivex.Observable

/**
 * Created by xx on 2020/2/25 21:32.
 * desc: AniDetailModel
 */
class AniDetailModel {
    fun requestRelatedAniData(id:Long): Observable<AniBean> {
        return RetrofitManager.service.getRelatedAni(id)
                .compose(SchedulerUtils.ioToMain())
    }
}