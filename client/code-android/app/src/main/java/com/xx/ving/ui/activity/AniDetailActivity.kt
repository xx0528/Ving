package com.xx.ving.ui.activity

import android.annotation.TargetApi
import android.content.res.Configuration
import android.os.Build
import android.support.v4.view.ViewCompat
import android.support.v7.widget.LinearLayoutManager
import android.transition.Transition
import android.widget.ImageView
import com.bumptech.glide.load.DecodeFormat
import com.bumptech.glide.load.resource.drawable.DrawableTransitionOptions
import com.xx.ving.base.BaseActivity
import com.xx.ving.glide.GlideApp
import com.xx.ving.mvp.contract.AniDetailContract
import com.xx.ving.mvp.model.bean.AniBean
import com.xx.ving.mvp.presenter.AniDetailPresenter
import com.xx.ving.ui.adapter.AniDetailAdapter
import com.xx.ving.utils.CleanLeakUtils
import com.xx.ving.utils.StatusBarUtil
import com.xx.ving.utils.WatchHistoryUtils
import com.xx.ving.view.VideoListener
import com.scwang.smartrefresh.header.MaterialHeader
import com.shuyu.gsyvideoplayer.utils.OrientationUtils
import java.text.SimpleDateFormat
import java.util.*
import com.orhanobut.logger.Logger
import com.shuyu.gsyvideoplayer.video.StandardGSYVideoPlayer
import com.shuyu.gsyvideoplayer.video.base.GSYVideoPlayer
import com.xx.ving.*
import kotlinx.android.synthetic.main.activity_ani_detail.*
import kotlin.collections.ArrayList

/**
 * Created by xx on 2020/2/25 19:33.
 * desc: AniDetailActivity
 */
class AniDetailActivity : BaseActivity(), AniDetailContract.View {

    companion object {
        const val IMG_TRANSITION = "IMG_TRANSITION"
        const val TRANSITION = "TRANSITION"
    }

    /**
     * 第一次调用的时候初始化
     */
    private val mPresenter by lazy { AniDetailPresenter() }

    private val mAdapter by lazy { AniDetailAdapter(this, itemList) }

    private val mFormat by lazy { SimpleDateFormat("yyyyMMddHHmmss");}

    /**
     * Item 详细数据
     */
    private lateinit var itemData: AniBean.AItem
    private var orientationUtils: OrientationUtils? = null

    private var itemList = ArrayList<AniBean.AItem>()

    private var isPlay: Boolean = false
    private var isPause: Boolean =false

    private var isTransition: Boolean = false

    private var transition: Transition? = null
    private var mMaterialHeader: MaterialHeader? = null

    override fun layoutId(): Int = R.layout.activity_ani_detail

    override fun initData() {
        itemData = intent.getSerializableExtra(Constants.BUNDLE_VIDEO_DATA) as AniBean.AItem
        isTransition = intent.getBooleanExtra(TRANSITION, false)

        saveWatchAniHistoryInfo(itemData)
    }

    /**
     * 保存观看记录
     */
    private fun saveWatchAniHistoryInfo(watchItem: AniBean.AItem) {
        //保存之前要先查询sp中是否有该value的记录，有则删除。这样保证搜索历史记录不会有重复的条目
        var historyMap = WatchHistoryUtils.getAll(Constants.FILE_WATCH_HISTORY_NAME, MyApplication.context) as Map<*, *>
        for ((key, _) in historyMap) {
            if (watchItem == WatchHistoryUtils.getObject(Constants.FILE_WATCH_HISTORY_NAME, MyApplication.context, key as String)) {
                WatchHistoryUtils.remove(Constants.FILE_WATCH_HISTORY_NAME, MyApplication.context, key)
            }
        }
        WatchHistoryUtils.putObject(Constants.FILE_WATCH_HISTORY_NAME, MyApplication.context, watchItem, "" + mFormat.format(Date()))
    }

    /**
     * 加载视频信息
     */
    fun loadAniInfo() {
        mPresenter.loadAniInfo(itemData)
    }

    @TargetApi(Build.VERSION_CODES.LOLLIPOP)
    private fun addTransitionListener() {
        transition = window.sharedElementEnterTransition
        transition?.addListener(object : Transition.TransitionListener {
            override fun onTransitionResume(p0: Transition?) {
            }

            override fun onTransitionPause(p0: Transition?) {
            }

            override fun onTransitionCancel(p0: Transition?) {
            }

            override fun onTransitionStart(p0: Transition?) {
            }

            override fun onTransitionEnd(p0: Transition?) {
                Logger.d("onTransitionEnd()------")

                loadAniInfo()
                transition?.removeListener(this)
            }
        })
    }


    private fun getCurPlay(): GSYVideoPlayer {
        return if (mAniVideoView.fullWindowPlayer != null) {
            mAniVideoView.fullWindowPlayer
        } else mAniVideoView
    }

    private fun initTransition() {
        if (isTransition && Build.VERSION.SDK_INT >= Build.VERSION_CODES.LOLLIPOP) {
            postponeEnterTransition()
            ViewCompat.setTransitionName(mAniVideoView, IMG_TRANSITION)
            addTransitionListener()
            startPostponedEnterTransition()
        } else {
            loadAniInfo()
        }
    }

    /**
     * 初始化 AniView 的配置
     */
    private fun initAniViewConfig() {
        //设置旋转
        orientationUtils = OrientationUtils(this, mAniVideoView)
        //是否旋转
        mAniVideoView.isRotateViewAuto = false
        //是否可以滑动调整
        mAniVideoView.setIsTouchWiget(true)

        //增加封面
        val imageView = ImageView(this)
        imageView.scaleType = ImageView.ScaleType.CENTER_CROP
        GlideApp.with(this)
                .load(itemData.data?.cover?.homepage)
                .centerCrop()
                .into(imageView)
        mAniVideoView.thumbImageView = imageView

        mAniVideoView.setStandardVideoAllCallBack(object : VideoListener {

            override fun onPrepared(url: String, vararg objects: Any) {
                super.onPrepared(url, *objects)
                //开始播放了才能旋转和全屏
                orientationUtils?.isEnable = true
                isPlay = true
            }

            override fun onAutoComplete(url: String, vararg objects: Any) {
                super.onAutoComplete(url, *objects)
                Logger.d("***** onAutoPlayComplete **** ")
            }

            override fun onPlayError(url: String, vararg objects: Any) {
                super.onPlayError(url, *objects)
                showToast("播放失败")
            }

            override fun onEnterFullscreen(url: String, vararg objects: Any) {
                super.onEnterFullscreen(url, *objects)
                Logger.d("***** onEnterFullscreen **** ")
            }

            override fun onQuitFullscreen(url: String, vararg objects: Any) {
                super.onQuitFullscreen(url, *objects)
                Logger.d("***** onQuitFullscreen **** ")
                //列表返回的样式判断
                orientationUtils?.backToProtVideo()
            }
        })
        //设置返回键功能
        mAniVideoView.backButton.setOnClickListener { onBackPressed() }
        //设置全屏按键功能
        mAniVideoView.fullscreenButton.setOnClickListener {
            //直接横屏
            orientationUtils?.resolveByClick()
            //第一个ture是否需要隐藏actionbar，第二个true是否需要隐藏statusbar
            mAniVideoView.startWindowFullscreen(this, true, true)
        }
        //锁屏事件
        mAniVideoView.setLockClickListener { _, lock ->
            //配合下方的onConfigurationChanged
            orientationUtils?.isEnable = !lock
        }
    }

    override fun onConfigurationChanged(newConfig: Configuration?) {
        super.onConfigurationChanged(newConfig)
        if (isPlay && !isPause) {
            mAniVideoView.onConfigurationChanged(this, newConfig, orientationUtils)
        }
    }

    override fun initView() {

        mPresenter.attachView(this)
        //过渡动画
        initTransition()
        initAniViewConfig()

        mRecyclerView.layoutManager = LinearLayoutManager(this)
        mRecyclerView.adapter = mAdapter

        //设置相关视频Item的点击事件
        mAdapter.setOnItemDetailClick { mPresenter.loadAniInfo(it) }

        //状态栏透明和间距处理
        StatusBarUtil.immersive(this)
        StatusBarUtil.setPaddingSmart(this, mAniVideoView)

        /***
         * 下拉刷新
         */
        //内容跟随偏移
        mRefreshLayout.setEnableHeaderTranslationContent(true)
        mRefreshLayout.setOnRefreshListener {
            loadAniInfo()
        }
        mMaterialHeader = mRefreshLayout.refreshHeader as MaterialHeader?
        //打开下拉刷新区域背景
        mMaterialHeader?.setShowBezierWave(true)
        //设置下拉主题颜色
        mRefreshLayout.setPrimaryColorsId(R.color.color_lighter_gray, R.color.color_title_bg)

    }

    override fun start() {
    }

    override fun setAni(url: String) {
        Logger.d("playUrl:$url")
        mAniVideoView.setUp(url, false, "")
        //开始自动播放
//        mAniVideoView.startPlayLogic()
    }

    override fun setAniInfo(itemInfo: AniBean.AItem) {
        itemData = itemInfo
        mAdapter.addData(itemInfo)

        tv_action_favorites.text = itemData.data?.consumption?.collectionCount.toString()
        tv_action_share.text = itemData.data?.consumption?.shareCount.toString()
        tv_action_reply.text = itemData.data?.consumption?.replyCount.toString()

        // 请求相关最新视频
        mPresenter.requestRelatedAni(itemInfo.data?.id?:0)
    }

    override fun setBackground(url: String) {
        GlideApp.with(this)
                .load(url)
                .centerCrop()
                .format(DecodeFormat.PREFER_ARGB_8888)
                .transition(DrawableTransitionOptions().crossFade())
                .into(mAniBackground)
    }

    override fun setRecentRelatedAni(itemList: ArrayList<AniBean.AItem>) {
        mAdapter.addData(itemList)
        this.itemList = itemList
    }

    override fun setErrorMsg(errorMsg: String) {
    }

    override fun showLoading() {
    }

    override fun dismissLoading() {
        mRefreshLayout.finishRefresh()
    }

    override fun onResume() {
        super.onResume()
        getCurPlay().onVideoResume()
        isPause = false
    }

    override fun onPause() {
        super.onPause()
        getCurPlay().onVideoPause()
        isPause = true
    }

    override fun onDestroy() {
        CleanLeakUtils.fixInputMethodManagerLeak(this)
        super.onDestroy()
        GSYVideoPlayer.releaseAllVideos()
        orientationUtils?.releaseListener()
        mPresenter.detachView()
    }


    /**
     * 监听返回键
     */
    override fun onBackPressed() {
        orientationUtils?.backToProtVideo()
        if (StandardGSYVideoPlayer.backFromWindowFull(this))
            return
        //释放所有
        mAniVideoView.setStandardVideoAllCallBack(null)
        GSYVideoPlayer.releaseAllVideos()
        if (isTransition && Build.VERSION.SDK_INT >= Build.VERSION_CODES.LOLLIPOP) run {
            super.onBackPressed()
        } else {
            finish()
            overridePendingTransition(R.anim.anim_out, R.anim.anim_in)
        }
    }
}