package com.xx.ving.ui.adapter

import android.annotation.SuppressLint
import android.app.Activity
import android.content.Context
import android.content.Intent
import android.support.v4.app.ActivityCompat
import android.support.v4.app.ActivityOptionsCompat
import android.support.v4.util.Pair
import android.view.View
import android.widget.TextView
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
class AniDetailSelectionAdapter (mContext: Context, selectionList: ArrayList<String>, layoutId: Int) :
        CommonAdapter<String>(mContext, selectionList, layoutId) {

    private var mSelections: ArrayList<String>? = null
    private var mCurSelect : Int = 0

    @SuppressLint("ResourceAsColor")
    override fun bindData(holder: ViewHolder, data: String, position: Int) {

        holder.setText(R.id.tv_select_id, data)
        if (position == mCurSelect)
            holder.getView<TextView>(R.id.tv_select_id).setTextColor(R.color.color_black)

        with(holder) {
            setOnItemClickListener(listener = View.OnClickListener {
                goToVideoPlayer(mContext as Activity, holder.getView(R.id.tv_select_id), position)
            })
        }
    }

    /**
     * 跳转到视频详情页面播放
     */
    private fun goToVideoPlayer(activity: Activity, view: View, selectionPos: Int) {
//        Logger.d("AniDetailSelection 进入播放设置数据=========== ")
//        val intent = Intent(activity, AniDetailActivity::class.java)
//        intent.putExtra(Constants.BUNDLE_VIDEO_DATA, selectionPos)
//        intent.putExtra(AniDetailActivity.Companion.TRANSITION, true)
//        if (android.os.Build.VERSION.SDK_INT >= android.os.Build.VERSION_CODES.LOLLIPOP) {
//            val pair = Pair<View, String>(view, AniDetailActivity.IMG_TRANSITION)
//            val activityOptions = ActivityOptionsCompat.makeSceneTransitionAnimation(
//                    activity, pair)
//            ActivityCompat.startActivity(activity, intent, activityOptions.toBundle())
//        }
//        else {
//            activity.startActivity(intent)
//            activity.overridePendingTransition(R.anim.anim_in, R.anim.anim_out)
//        }

        var act = activity as AniDetailActivity
        act.setAniByIdx(selectionPos)
        mCurSelect = selectionPos
        notifyDataSetChanged()
    }

    fun setData(selections: ArrayList<String>, curSelect: Int) {
//        Logger.d("AniDetailSelection 设置数据=========== ")
        mSelections = selections
        mCurSelect = curSelect
        notifyDataSetChanged()
    }

}