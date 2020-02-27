package com.xx.ving.ui.adapter

import android.app.Activity
import android.content.Context
import android.content.Intent
import android.support.v4.app.ActivityCompat
import android.support.v4.app.ActivityOptionsCompat
import android.support.v4.util.Pair
import android.view.View
import com.xx.ving.Constants
import com.xx.ving.R
import com.xx.ving.mvp.model.bean.AniBean
import com.xx.ving.ui.activity.AniDetailActivity
import com.xx.ving.view.recyclerview.ViewHolder
import com.xx.ving.view.recyclerview.adapter.CommonAdapter
import com.orhanobut.logger.Logger

/**
 * Created by xx on 2020/2/27 15:02.
 * desc: AniDetailSelectionAdapter 选集的adapter
 */
class AniDetailSelectionAdapter (mContext: Context, selectionList: ArrayList<Int>, layoutId: Int) :
        CommonAdapter<Int>(mContext, selectionList, layoutId) {

    private var mAniData: AniBean.AItem? = null

    override fun bindData(holder: ViewHolder, data: Int, position: Int) {

        holder.setText(R.id.tv_select_id, data.toString())

        with(holder) {
            setOnItemClickListener(listener = View.OnClickListener {
                mAniData?.let { it1 -> goToVideoPlayer(mContext as Activity, holder.getView(R.id.tv_select_id), it1) }
            })
        }
    }

    /**
     * 跳转到视频详情页面播放
     */
    private fun goToVideoPlayer(activity: Activity, view: View, aniData: AniBean.AItem) {
        Logger.d("AniDetailSelection 进入播放设置数据=========== ")
        val intent = Intent(activity, AniDetailActivity::class.java)
        intent.putExtra(Constants.BUNDLE_VIDEO_DATA, aniData)
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

    fun setData(aniData: AniBean.AItem) {
//        Logger.d("AniDetailSelection 设置数据=========== ")
        mAniData = aniData
        notifyDataSetChanged()
    }

}