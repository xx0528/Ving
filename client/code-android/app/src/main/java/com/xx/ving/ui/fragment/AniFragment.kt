package com.hazz.ving.ui.fragment

import android.graphics.Rect
import android.os.Bundle
import android.support.v7.widget.GridLayoutManager
import android.support.v7.widget.RecyclerView
import android.view.View
import com.hazz.ving.R
import com.hazz.ving.base.BaseFragment
import com.hazz.ving.mvp.contract.AniContract
import com.hazz.ving.mvp.model.bean.AniBean
import com.hazz.ving.mvp.presenter.AniPresenter
import com.hazz.ving.net.exception.ErrorStatus
import com.hazz.ving.showToast
import com.hazz.ving.ui.adapter.AniAdapter
import com.hazz.ving.utils.DisplayManager
import kotlinx.android.synthetic.main.fragment_ani.*

/**
 * Created by xx on 2020/02/21.
 * 动漫
 */
class AniFragment: BaseFragment(), AniContract.View{


    private val mPresenter by lazy { AniPresenter() }

    private val mAdapter by lazy {activity?.let{ AniAdapter(it, mAniList, R.layout.item_ani) }}

    private var mTitle:String? =null

    private var loadingMore = false

    private var mAniList = ArrayList<AniBean>()

    companion object {
        fun getInstance(title:String): AniFragment {
            val fragment = AniFragment()
            val bundle = Bundle()
            fragment.arguments = bundle
            fragment.mTitle = title
            return fragment
        }
    }


    override fun getLayoutId(): Int= R.layout.fragment_ani

    override fun initView() {
        mPresenter.attachView(this)

        mLayoutStatusView = multipleStatusView

        mRecyclerView.adapter = mAdapter
        mRecyclerView.layoutManager = GridLayoutManager(activity, 2)
        mRecyclerView.addItemDecoration(object : RecyclerView.ItemDecoration(){
            override fun getItemOffsets(outRect: Rect?, view: View?, parent: RecyclerView?, state: RecyclerView.State?) {
                super.getItemOffsets(outRect, view, parent, state)
                val position = parent?.getChildAdapterPosition(view)?:0
                val offset = DisplayManager.dip2px(2f)!!

                outRect?.set(if (position % 2 == 0) 0 else offset, offset,
                        if (position % 2 == 0) offset else 0, offset)
            }
        })

        //实现自动加载
        mRecyclerView.addOnScrollListener(object : RecyclerView.OnScrollListener() {
            override fun onScrollStateChanged(recyclerView: RecyclerView?, newState: Int) {
                super.onScrollStateChanged(recyclerView, newState)
                val itemCount = mRecyclerView.layoutManager.itemCount
                val lastVisibleItem = (mRecyclerView.layoutManager as GridLayoutManager).findLastVisibleItemPosition()
                if (!loadingMore && lastVisibleItem == (itemCount - 1)) {
                    loadingMore = true
                    mPresenter.loadMoreData()
                }
            }
        })
    }

    override fun lazyLoad() {
        mPresenter.getAnimationData()
    }

    override fun setAniData(aniList: ArrayList<AniBean>) {
//        Logger.d("setAniData    ----- ")
        mAniList = aniList
        mAdapter?.setData(mAniList)
    }

    override fun setMoreData(aniList: ArrayList<AniBean>) {
//        Logger.d("setMoreData ----- ")
        loadingMore = false
        mAdapter?.addItemData(aniList)
    }
    override fun showError(errorMsg: String, errorCode: Int) {
        showToast(errorMsg)
        if (errorCode == ErrorStatus.NETWORK_ERROR) {
            multipleStatusView?.showNoNetwork()
        } else {
            multipleStatusView?.showError()
        }
    }

    override fun showLoading() {
        multipleStatusView?.showLoading()
    }

    override fun dismissLoading() {
        multipleStatusView?.showContent()
    }

    override fun onDestroy() {
        super.onDestroy()
        mPresenter.detachView()
    }

}