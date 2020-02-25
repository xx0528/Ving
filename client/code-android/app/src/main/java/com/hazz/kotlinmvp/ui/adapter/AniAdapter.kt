package com.hazz.kotlinmvp.ui.adapter

import android.content.Context
import android.widget.ImageView
import android.widget.TextView
import com.bumptech.glide.load.resource.drawable.DrawableTransitionOptions
import com.hazz.kotlinmvp.R
import com.hazz.kotlinmvp.glide.GlideApp
import com.hazz.kotlinmvp.mvp.model.bean.AniBean
import com.hazz.kotlinmvp.view.recyclerview.ViewHolder
import com.hazz.kotlinmvp.view.recyclerview.adapter.CommonAdapter
import com.orhanobut.logger.Logger
import kotlin.math.log

/**
 * Created by xx on 2020/2/24 0:31.
 * desc: AniAdapter
 */
class AniAdapter(mContext: Context, aniList: ArrayList<AniBean>, layoutId: Int) :
        CommonAdapter<AniBean>(mContext, aniList, layoutId){

    fun setData(aniList: ArrayList<AniBean>){
        mData.clear()
        mData = aniList
        notifyDataSetChanged()
    }

    fun addItemData(aniList: ArrayList<AniBean>) {
        this.mData.addAll(aniList)
        notifyDataSetChanged()
    }

    override fun getItemCount(): Int {
//        Logger.d("size --- %d", this.mData.size)
        return this.mData.size
    }

    override fun bindData(holder: ViewHolder, data: AniBean, position: Int) {
        holder.setImagePath(R.id.iv_ani_img, object : ViewHolder.HolderImageLoader(data.bgPicture) {
            override fun loadImage(iv: ImageView, path: String) {
                GlideApp.with(mContext)
                        .load(path)
                        .placeholder(R.color.color_darker_gray)
                        .transition(DrawableTransitionOptions().crossFade())
                        .thumbnail(0.5f)
                        .into(iv)
            }
        })

        holder.getView<TextView>(R.id.tv_ani_name).text = data.name
        holder.getView<TextView>(R.id.tv_ani_desc).text = data.description
        holder.getView<TextView>(R.id.tv_ani_num).text = data.videoNum.toString() + "集全"

    }
}