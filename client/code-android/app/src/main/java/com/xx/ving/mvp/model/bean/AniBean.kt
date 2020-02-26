package com.hazz.ving.mvp.model.bean

import java.io.Serializable

/**
 * Created by xx on 2020/02/24.
 * desc: 动画
 */
data class AniBean(val id: Long, val name: String, val description: String, val bgPicture: String, val videoNum: Int) : Serializable
