<template>
  <div id="content">
    <div class="article_wrap">
      <div class="article_title article_detail_title">{{ this.article.title }}</div>
      <div class="article_info">
        <span class="article_info_date">发表于：{{ this.article.date }}</span>
        <span class="article_info_label">标签：
        <span v-if="this.article.labels == ''">未分类</span>
        <el-tag v-else class="tag_margin" type="primary" v-for="tag in this.article.labels">{{ tag }}</el-tag>
        </span>
      </div>
      <div class="article_detail_content" v-html="compiledMarkdown()"></div>
    </div>
  </div>
</template>

<script>
  import marked from 'marked'
  import highlight from 'highlight.js'
  import '../assets/atom-one-light.css'
  marked.setOptions({
    highlight: function (code) {
      return highlight.highlightAuto(code).value
    }
  })
  export default {
    name: 'articleDetail',
    data() {
      return {
        article: {}
      }
    },
    mounted: function () {
      let _articleId = this.$route.params.articleId
      console.log("开始打印_articleId--->")
      console.log(_articleId)
      this.$axios.post(global.articleDetail,{
        'articleId' : _articleId
      }).then(
        response => {
          this.article = JSON.parse(response.data)["data"]
        },
        response => console.log(response)
      )
    },
    methods: {
      compiledMarkdown: function () {
        return marked(this.article.content || '', {sanitize: true})
      }
    }
  }
</script>

<style>
  .article_detail_title {
    cursor: default;
    margin: 40px 0 0;
  }

  .article_detail_content {
    text-align: left;
    padding: 60px 0;
    font-size: 18px;
  }
</style>