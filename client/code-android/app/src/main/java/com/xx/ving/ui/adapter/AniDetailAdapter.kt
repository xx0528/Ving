package com.hazz.ving.ui.adapter

import android.content.Context
import com.hazz.ving.R
import com.hazz.ving.mvp.model.bean.AniBean
import com.hazz.ving.view.recyclerview.MultipleType
import com.hazz.ving.view.recyclerview.ViewHolder
import com.hazz.ving.view.recyclerview.adapter.CommonAdapter

/**
 * Created by xx on 2020/2/25 21:46.
 * desc: AniDetailAdapter 动画详情
 */
class AniDetailAdapter(mContext: Context, aniList: ArrayList<AniBean>) :
        CommonAdapter<AniBean>(mContext, aniList, object : MultipleType<AniBean>{
            override fun getLayoutId(item: AniBean, position: Int): Int {
                return when {
                    position == 0 ->
                        R.layout.item_video_detail_info

                    aniList[position].name == "textCard" ->
                        R.layout.item_video_text_card

                    aniList[position].name == "videoSmallCard" ->
                        R.layout.item_video_small_card
                    else ->
                        throw IllegalAccessException("Api 解析出错了，出现其他类型")
                }
            }
        }) {

    /**
     * 添加动画的详细信息
     */
    fun addData(item: AniBean) {
        mData.clear()
        notifyDataSetChanged()
        mData.add(item)
        notifyItemInserted(0)
    }

    /**
     * 添加相关推荐等数据
     */
    fun addData(item: ArrayList<AniBean>) {
        mData.addAll(item)
        notifyItemRangeInserted(1, item.size)
    }

    private var mOnItemClickRelatedAni: ((item: AniBean) -> Unit)? = null

    fun setOnItemDetailClick(mItemRelatedAni: (item: AniBean) -> Unit) {
        this.mOnItemClickRelatedAni = mItemRelatedAni
    }

    /**
     * 绑定数据
     */
    override fun bindData(holder: ViewHolder, data: AniBean, position: Int) {

    }

    /**
     * 设置动画详情数据
     */
    private fun setAniDetailInfo(data: AniBean, holder: ViewHolder) {

    }
}