package com.xx.ving.mvp.contract

import com.xx.ving.base.IBaseView
import com.xx.ving.base.IPresenter
import com.xx.ving.mvp.model.bean.HomeBean

/**
 * Created by xuhao on 2017/11/30.
 * desc: 契约类
 */
interface RankContract {

    interface View:IBaseView{
        /**
         * 设置排行榜的数据
         */
        fun setRankList(itemList: ArrayList<HomeBean.Issue.Item>)

        fun showError(errorMsg:String,errorCode:Int)
    }


    interface Presenter:IPresenter<View>{
        /**
         * 获取 TabInfo
         */
        fun requestRankList(apiUrl:String)
    }
}