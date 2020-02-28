package com.xx.ving.ui.activity

import android.annotation.SuppressLint
import android.graphics.Rect
import android.os.Bundle
import android.support.v4.app.FragmentTransaction
import android.view.KeyEvent
import com.flyco.tablayout.listener.CustomTabEntity
import com.flyco.tablayout.listener.OnTabSelectListener
import com.orhanobut.logger.Logger
import com.xx.ving.R
import com.xx.ving.base.BaseActivity
import com.xx.ving.mvp.model.bean.TabEntity
import com.xx.ving.showToast
import com.xx.ving.ui.fragment.*
import kotlinx.android.synthetic.main.activity_main.*
import java.util.*


/**
 * @author Jake.Ho
 * created: 2017/10/25
 * desc:
 */


class MainActivity : BaseActivity() {


//    private val mTitles = arrayOf("动漫", "我的", "每日精选", "发现", "热门")

    private val mTitles = arrayOf("动漫", "我的")

    // 未被选中的图标
    private val mIconUnSelectIds = intArrayOf(R.mipmap.ic_animation_normal, R.mipmap.ic_mine_normal, R.mipmap.ic_home_normal, R.mipmap.ic_discovery_normal, R.mipmap.ic_hot_normal)
    // 被选中的图标
    private val mIconSelectIds = intArrayOf(R.mipmap.ic_animation_selected, R.mipmap.ic_mine_selected, R.mipmap.ic_home_selected, R.mipmap.ic_discovery_selected, R.mipmap.ic_hot_selected)

    private val mTabEntities = ArrayList<CustomTabEntity>()

    private var mHomeFragment: HomeFragment? = null
    private var mDiscoveryFragment: DiscoveryFragment? = null
    private var mHotFragment: HotFragment? = null
    private var mMineFragment: MineFragment? = null
    private var mAniFragment: AniFragment? = null

    //默认为0
    private var mIndex = 0


    override fun onCreate(savedInstanceState: Bundle?) {
        if (savedInstanceState != null) {
            mIndex = savedInstanceState.getInt("currTabIndex")
        }
        super.onCreate(savedInstanceState)
        initTab()
        tab_layout.currentTab = mIndex
        switchFragment(mIndex)

    }

    override fun layoutId(): Int {
        return R.layout.activity_main
    }


    //初始化底部菜单
    private fun initTab() {
        (0 until mTitles.size)
                .mapTo(mTabEntities) { TabEntity(mTitles[it], mIconSelectIds[it], mIconUnSelectIds[it]) }
        //为Tab赋值
        tab_layout.setTabData(mTabEntities)
        tab_layout.setOnTabSelectListener(object : OnTabSelectListener {
            override fun onTabSelect(position: Int) {
                //切换Fragment
                switchFragment(position)
            }

            override fun onTabReselect(position: Int) {

            }
        })
    }

    /**
     * 切换Fragment
     * @param position 下标
     */
    private fun switchFragment(position: Int) {
        val transaction = supportFragmentManager.beginTransaction()
        hideFragments(transaction)
//        var outRect = Rect()
//        tab_layout.getDrawingRect(outRect)
//        Logger.d("高度--- %d", fl_container.height)
//        Logger.d("tab layout --- %s", outRect)
        when (position) {
            0 //动漫
            -> mAniFragment?.let {
                transaction.show(it)
            } ?: AniFragment.getInstance(mTitles[position]).let {
                mAniFragment = it
                transaction.add(R.id.fl_container, it, "animation") }
            1 //我的
            -> mMineFragment?.let {
                transaction.show(it)
            } ?: MineFragment.getInstance(mTitles[position]).let {
                mMineFragment = it
                transaction.add(R.id.fl_container, it, "mine") }
            2 // 首页
            -> mHomeFragment?.let {
                transaction.show(it)
            } ?: HomeFragment.getInstance(mTitles[position]).let {
                mHomeFragment = it
                transaction.add(R.id.fl_container, it, "home")
            }
            3  //发现
            -> mDiscoveryFragment?.let {
                transaction.show(it)
            } ?: DiscoveryFragment.getInstance(mTitles[position]).let {
                mDiscoveryFragment = it
                transaction.add(R.id.fl_container, it, "discovery") }
            4  //热门
            -> mHotFragment?.let {
                transaction.show(it)
            } ?: HotFragment.getInstance(mTitles[position]).let {
                mHotFragment = it
                transaction.add(R.id.fl_container, it, "hot") }

            else -> {

            }
        }
//        tab_layout.getDrawingRect(outRect)
//        Logger.d("高度-  === -- %d", fl_container.height)
//        Logger.d("tab layout rect--- %s", outRect)
        mIndex = position
        tab_layout.currentTab = mIndex
        transaction.commitAllowingStateLoss()
    }


    /**
     * 隐藏所有的Fragment
     * @param transaction transaction
     */
    private fun hideFragments(transaction: FragmentTransaction) {
        mAniFragment?.let { transaction.hide(it) }
        mMineFragment?.let { transaction.hide(it) }
        mHomeFragment?.let { transaction.hide(it) }
        mDiscoveryFragment?.let { transaction.hide(it) }
        mHotFragment?.let { transaction.hide(it) }
    }


    @SuppressLint("MissingSuperCall")
    override fun onSaveInstanceState(outState: Bundle) {
//        showToast("onSaveInstanceState->"+mIndex)
//        super.onSaveInstanceState(outState)
        //记录fragment的位置,防止崩溃 activity被系统回收时，fragment错乱
        if (tab_layout != null) {
            outState.putInt("currTabIndex", mIndex)
        }
    }

    override fun initView() {

    }

    override fun initData() {

    }

    override fun start() {

    }

    private var mExitTime: Long = 0

    override fun onKeyDown(keyCode: Int, event: KeyEvent?): Boolean {
        if (keyCode == KeyEvent.KEYCODE_BACK) {
            if (System.currentTimeMillis().minus(mExitTime) <= 2000) {
                finish()
            } else {
                mExitTime = System.currentTimeMillis()
                showToast("再按一次退出程序")
            }
            return true
        }
        return super.onKeyDown(keyCode, event)
    }


}
