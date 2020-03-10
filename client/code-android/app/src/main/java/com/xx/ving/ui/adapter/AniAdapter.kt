package com.xx.ving.ui.adapter

import android.annotation.SuppressLint
import android.app.Activity
import android.content.Context
import android.content.Intent
import android.support.v4.app.ActivityCompat
import android.support.v4.app.ActivityOptionsCompat
import android.support.v4.util.Pair
import android.view.View
import android.widget.ImageView
import android.widget.TextView
import com.bumptech.glide.load.resource.drawable.DrawableTransitionOptions
import com.xx.ving.Constants
import com.xx.ving.R
import com.xx.ving.glide.GlideApp
import com.xx.ving.mvp.model.bean.AniBean
import com.xx.ving.ui.activity.AniDetailActivity
import com.xx.ving.view.recyclerview.ViewHolder
import com.xx.ving.view.recyclerview.adapter.CommonAdapter
import com.orhanobut.logger.Logger

/**
 * Created by xx on 2020/2/24 0:31.
 * desc: AniAdapter
 */
class AniAdapter(mContext: Context, aniList: ArrayList<AniBean.AItem>, layoutId: Int) :
        CommonAdapter<AniBean.AItem>(mContext, aniList, layoutId){

    fun setData(aniList: ArrayList<AniBean.AItem>){
        mData.clear()
        mData = aniList
        notifyDataSetChanged()
    }

    fun addItemData(aniList: ArrayList<AniBean.AItem>) {
        this.mData.addAll(aniList)
        notifyDataSetChanged()
    }

    override fun getItemCount(): Int {
//        Logger.d("size --- %d", this.mData.size)
        return this.mData.size
    }

    @SuppressLint("SetTextI18n")
    override fun bindData(holder: ViewHolder, data: AniBean.AItem, position: Int) {
        data.data?.cover?.homepage?.let {
            holder.setImagePath(R.id.iv_ani_img, object : ViewHolder.HolderImageLoader(it) {
            override fun loadImage(iv: ImageView, path: String) {
                GlideApp.with(mContext)
                        .load(path)
                        .placeholder(R.color.color_darker_gray)
                        .transition(DrawableTransitionOptions().crossFade())
                        .thumbnail(0.5f)
                        .into(iv)
            }
        })
        }

        holder.getView<TextView>(R.id.tv_ani_name).text = data.data?.name ?: "默认名字"
        holder.getView<TextView>(R.id.tv_ani_desc).text = data.data?.desc ?: "默认描述"

        var num = data.data?.aniNum
        if (num == null || num == 0)
        {
            holder.getView<TextView>(R.id.tv_ani_num).text = ""
        }
        else
        {
            holder.getView<TextView>(R.id.tv_ani_num).text = num.toString() + "集全"
        }

        holder.setOnItemClickListener(listener = View.OnClickListener {
            goToAnimationPlayer(mContext as Activity, holder.getView(R.id.iv_ani_img), data)
        })
    }

    private fun goToAnimationPlayer(activity: Activity, view: View, itemData: AniBean.AItem) {
        val intent = Intent(activity, AniDetailActivity::class.java)
        intent.putExtra(Constants.BUNDLE_VIDEO_DATA, itemData)
        intent.putExtra(AniDetailActivity.TRANSITION, true)
        if (android.os.Build.VERSION.SDK_INT >= android.os.Build.VERSION_CODES.LOLLIPOP) {
            val pair = Pair(view, AniDetailActivity.IMG_TRANSITION)
            val activityOptions = ActivityOptionsCompat.makeSceneTransitionAnimation(
                    activity, pair)
            ActivityCompat.startActivity(activity, intent, activityOptions.toBundle())
        } else {
            activity.startActivity(intent)
            activity.overridePendingTransition(R.anim.anim_in, R.anim.anim_out)
        }
    }
}