package com.xx.ving.ui.fragment

import android.os.Bundle
import android.support.v7.widget.LinearLayoutManager
import com.xx.ving.R
import com.xx.ving.base.BaseFragment
import com.xx.ving.mvp.contract.RankContract
import com.xx.ving.mvp.model.bean.HomeBean
import com.xx.ving.mvp.presenter.RankPresenter
import com.xx.ving.net.exception.ErrorStatus
import com.xx.ving.showToast
import com.xx.ving.ui.adapter.CategoryDetailAdapter
import kotlinx.android.synthetic.main.fragment_rank.*

/**
 * Created by xuhao on 2017/11/30.
 * desc:
 */
class RankFragment : BaseFragment(), RankContract.View {


    private val mPresenter by lazy { RankPresenter() }

    private val mAdapter by lazy { activity?.let { CategoryDetailAdapter(it, itemList, R.layout.item_category_detail) } }

    private var itemList = ArrayList<HomeBean.Issue.Item>()

    private var apiUrl: String? = null

    companion object {
        fun getInstance(apiUrl: String): RankFragment {
            val fragment = RankFragment()
            val bundle = Bundle()
            fragment.arguments = bundle
            fragment.apiUrl = apiUrl
            return fragment
        }
    }

    init {
        mPresenter.attachView(this)
    }


    override fun getLayoutId(): Int = R.layout.fragment_rank

    override fun initView() {
        mRecyclerView.layoutManager = LinearLayoutManager(activity)
        mRecyclerView.adapter = mAdapter

        mLayoutStatusView =multipleStatusView

    }

    override fun lazyLoad() {
        if (!apiUrl.isNullOrEmpty()) {
            mPresenter.requestRankList(apiUrl!!)
        }
    }

    override fun showLoading() {
        multipleStatusView.showLoading()
    }

    override fun dismissLoading() {

    }

    override fun setRankList(itemList: ArrayList<HomeBean.Issue.Item>) {
        multipleStatusView.showContent()

        mAdapter?.addData(itemList)
    }

    override fun showError(errorMsg: String,errorCode:Int) {
        showToast(errorMsg)
        if (errorCode == ErrorStatus.NETWORK_ERROR) {
            multipleStatusView.showNoNetwork()
        } else {
            multipleStatusView.showError()
        }
    }

    override fun onDestroy() {
        super.onDestroy()
        mPresenter.detachView()
    }
}