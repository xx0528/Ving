package com.xx.ving.mvp.model.bean

import java.io.Serializable

/**
 * Created by xx on 2020/02/24.
 * desc: 动画数据
 */
data class AniBean(val aType: String, val releaseTime: Long, val date: Long, val total: Int, val publishTime: Long, val itemList: ArrayList<AItem>, var count: Int, val nextPage: String) {
    data class AItem(val aType: String, val data: AData?, val tag: String) : Serializable {
        data class AData(val id: Long,
                         val name: String,
                         val title: String,
                         val text: String,
                         val aniTitle: String,
                         val dataType: String,
                         val desc: String,
                         val author: Author,
                         val cover: Cover,
                         val category: String,
                         val consumption: Consumption,
                         val likeCount: Int,
                         val score: Float,
                         val playUrl: String,
                         val duration: Long,
                         val createTime: Long,
                         val tags: ArrayList<Tag>,
                         val aType: String,
                         val date: Long,
                         val idx: Int,
                         val aniNum: Int,
                         val curPlayNum: Int,
                         val label: Any,
                         val promotion: Any,
                         val played: Boolean,
                         val subtitles: Any,
                         val lastViewTime: Any
                        ) : Serializable {
                            data class Tag(val id: Int, val name: String, val bgPicture: String, val headerImage: String, val actionUrl: String, val adTract: Any) : Serializable

                            data class Author(val id: Int, val videoNum: Int, val icon: String, val name: String, val description: String) : Serializable

                            data class Cover(val feed: String, val detail: String, val blurred: String, val sharing: String, val homepage: String) : Serializable

                            data class Consumption(val collectionCount: Int, val shareCount: Int, val replyCount: Int) : Serializable
                        }
    }
}
































