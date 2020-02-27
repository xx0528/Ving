package com.xx.ving.mvp.model

import com.xx.ving.mvp.model.bean.TabInfoBean
import com.xx.ving.net.RetrofitManager
import com.xx.ving.rx.scheduler.SchedulerUtils
import io.reactivex.Observable

/**
 * Created by xuhao on 2017/11/30.
 * desc: 热门 Model
 */
class HotTabModel {

    /**
     * 获取 TabInfo
     */
    fun getTabInfo(): Observable<TabInfoBean> {

        return RetrofitManager.service.getRankList()
                .compose(SchedulerUtils.ioToMain())
    }

}
