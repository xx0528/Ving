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
import com.xx.ving.Constants
import com.xx.ving.R
import com.xx.ving.mvp.model.bean.AniBean
import com.xx.ving.ui.activity.AniDetailActivity
import com.xx.ving.view.recyclerview.ViewHolder
import com.xx.ving.view.recyclerview.adapter.CommonAdapter
import com.orhanobut.logger.Logger
import com.xx.ving.durationFormat
import com.xx.ving.glide.GlideApp
import com.xx.ving.glide.GlideRoundTransform

/**
 * Created by xx on 2020/3/18 23:14.
 * desc: AniRelatedAdapter
 */

class AniRelatedAdapter (mContext: Context, relatedList: ArrayList<AniBean.AItem>, layoutId: Int) :
        CommonAdapter<AniBean.AItem>(mContext, relatedList, layoutId) {

    private var mRelateds: ArrayList<AniBean.AItem>? = null
    private var mCurSelect : Int = 0

    @SuppressLint("ResourceAsColor")
    override fun bindData(holder: ViewHolder, data: AniBean.AItem, position: Int) {

        with(holder) {
            setText(R.id.tv_title, data.data?.name!!)
//            setText(R.id.tv_tag, "#${data.data.area}/${data.data.lang}/${data.data.year}/${durationFormat(data.data.duration)}")
            setImagePath(R.id.iv_ani_small_card, object : ViewHolder.HolderImageLoader(data.data.cover.detail) {
                override fun loadImage(iv: ImageView, path: String) {
                    GlideApp.with(mContext)
                            .load(path)
                            .optionalTransform(GlideRoundTransform())
                            .placeholder(R.drawable.placeholder_banner)
                            .into(iv)
                }
            })
        }

        with(holder) {
            setOnItemClickListener(listener = View.OnClickListener {
                goToVideoPlayer(mContext as Activity, holder.getView(R.id.iv_ani_small_card), data)
            })
        }
    }

    /**
     * 跳转到视频详情页面播放
     */
    private fun goToVideoPlayer(activity: Activity, view: View, itemData: AniBean.AItem) {
        Logger.d("AniDetailSelection 进入播放设置数据=========== ")
        val intent = Intent(activity, AniDetailActivity::class.java)
        intent.putExtra(Constants.BUNDLE_VIDEO_DATA, itemData)
        intent.putExtra(AniDetailActivity.Companion.TRANSITION, true)
        if (android.os.Build.VERSION.SDK_INT >= android.os.Build.VERSION_CODES.LOLLIPOP) {
            val pair = Pair<View, String>(view, AniDetailActivity.IMG_TRANSITION)
            val activityOptions = ActivityOptionsCompat.makeSceneTransitionAnimation(
                    activity, pair)
            ActivityCompat.startActivity(activity, intent, activityOptions.toBundle())
        }
        else {
            activity.startActivity(intent)
            activity.overridePendingTransition(R.anim.anim_in, R.anim.anim_out)
        }

    }


}