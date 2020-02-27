package com.xx.ving.ui.adapter

import android.app.Activity
import android.content.Context
import android.support.v7.widget.LinearLayoutManager
import android.support.v7.widget.RecyclerView
import android.view.View
import android.widget.ImageView
import com.xx.ving.R
import com.xx.ving.durationFormat
import com.xx.ving.glide.GlideApp
import com.xx.ving.mvp.model.bean.AniBean
import com.xx.ving.view.recyclerview.MultipleType
import com.xx.ving.view.recyclerview.ViewHolder
import com.xx.ving.view.recyclerview.adapter.CommonAdapter
import com.orhanobut.logger.Logger

/**
 * Created by xx on 2020/2/25 21:46.
 * desc: AniDetailAdapter 动画详情
 */
class AniDetailAdapter(mContext: Context, aniList: ArrayList<AniBean.AItem>) :
        CommonAdapter<AniBean.AItem>(mContext, aniList, object : MultipleType<AniBean.AItem>{
            override fun getLayoutId(item: AniBean.AItem, position: Int): Int {
                return when {
                    position == 0 ->
                        R.layout.item_ani_detail_info

                    aniList[position].aType == "textCard" ->
                        R.layout.item_ani_text_card

                    aniList[position].aType == "videoSmallCard" ->
                        R.layout.item_ani_small_card
                    else ->
                        throw IllegalAccessException("Api 解析出错了，出现其他类型")
                }
            }
        }) {

    /**
     * 添加动画的详细信息
     */
    fun addData(item: AniBean.AItem) {
        Logger.d("ani detail adapter add data-----")
        mData.clear()
        notifyDataSetChanged()
        mData.add(item)
        notifyItemInserted(0)
    }

    /**
     * 添加相关推荐等数据
     */
    fun addData(item: ArrayList<AniBean.AItem>) {
        mData.addAll(item)
        notifyItemRangeInserted(1, item.size)
    }

    private var mOnItemClickRelatedAni: ((item: AniBean.AItem) -> Unit)? = null

    fun setOnItemDetailClick(mItemRelatedAni: (item: AniBean.AItem) -> Unit) {
        this.mOnItemClickRelatedAni = mItemRelatedAni
    }

    /**
     * 绑定数据
     */
    override fun bindData(holder: ViewHolder, data: AniBean.AItem, position: Int) {
//        Logger.d("bind data ------%d", position)

        when {
            position == 0 -> setAniDetailInfo(data, holder)

            data.aType == "textCard" -> {

            }

            data.aType == "aniSmallCard" -> {

            }
            else -> throw IllegalAccessException("Api 解析出错了，出现其他类型--" + data.aType)
        }
    }

    /**
     * 设置动画详情数据
     */
    private fun setAniDetailInfo(data: AniBean.AItem, holder: ViewHolder) {
        data.data?.title?.let { holder.setText(R.id.tv_title, it) }
        data.data?.title?.let { holder.setText(R.id.expandable_text, it) }
        holder.setText(R.id.tv_tag, "#${data.data?.category} / ${durationFormat(data.data?.duration)}")

//        Logger.d("set detail info data ------")

        if (data.data?.author != null) {
            with(holder) {
                setText(R.id.tv_author_name, data.data.author.name)
                setText(R.id.tv_author_desc, data.data.author.description)
                setText(R.id.tv_ani_num, data.data.aniNum.toString()+"集全")

                setImagePath(R.id.iv_avatar, object : ViewHolder.HolderImageLoader(data.data.author.icon) {
                    override fun loadImage(iv: ImageView, path: String) {
                        //加载头像
                        GlideApp.with(mContext)
                                .load(path)
                                .placeholder(R.mipmap.default_avatar).circleCrop()
                                .into(iv)
                    }
                })
            }
        } else {
            holder.setViewVisibility(R.id.layout_author_view, View.GONE)
        }

        /**
         * 设置嵌套水平的 RecyclerView
         */
        val num = data.data?.aniNum ?: 10
        var selections = ArrayList<Int>()
        for (index in 1..num){
            selections.add(index)
        }

        val adapter = AniDetailSelectionAdapter(mContext, selections, R.layout.item_ani_selection)
        val recyclerView = holder.getView<RecyclerView>(R.id.rv_selections)
        recyclerView.layoutManager = LinearLayoutManager(mContext as Activity,LinearLayoutManager.HORIZONTAL,false)
        recyclerView.adapter = adapter
        adapter.setData(data)

    }
}