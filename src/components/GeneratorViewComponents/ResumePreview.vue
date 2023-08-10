<script setup lang="ts">

import { NAvatar,NDivider } from 'naive-ui'
</script>
<template>
  <div class="resume">
    <div class="left-side">
      <div class="profile">
        <n-avatar :size="100" src="/avator.jpg" />
      </div>
      <div class="contact">
        <div class="heading-text">联系方式</div>
        <p class="para">
          家庭住址：72-2 <br>
          +8612344577 <br>
          +1233 <br>
          example@gmail.com <br>
          www.baidu.com <br>
        </p>
      </div>
      <div class="expert">
        <div class="heading-text">
          专业领域
        </div>
        <p>
          JS <br>
          语文 <br>
          数学 <br>
          英语 <br>
          政治 <br>
        </p>
      </div>
      <div class="skill">
        <div class="heading-text">
          我的技能
        </div>
        <p>
          技能1 <br>
          语文1 <br>
          数学1 <br>
          英语1 <br>
          政治1 <br>
        </p>
      </div>
      <div class="hobby">
        <div class="heading-text">
          我的爱好
        </div>
        <p>
          爱好1 <br>
          爱好2 <br>
          爱好3 <br>
          爱好4 <br>
          政治1 <br>
        </p>
      </div>
    </div>
    <n-divider vertical></n-divider>
    <div class="right-side">
      <h2 class="name">
        姓名
      </h2>
      <div class="contact">
        <div class="heading-text">关于我</div>
        <p class="para">
          2017年毕业 ，专业软件工程
        </p>
      </div>
      <div class="personal">
        <div class="heading-text">个人信息</div>
        <table class="per-info">
          <tbody>
            <tr class="border">
              <td>父亲名称</td>
              <td>父亲</td>
            </tr>
            <tr class="border">
              <td>母亲名称</td>
              <td>母亲</td>
            </tr>
            <tr class="border">
              <td>出生日期</td>
              <td>2020-12-13</td>
            </tr>
            <tr class="border">
              <td>地址</td>
              <td>华盛顿</td>
            </tr>
          </tbody>

        </table>
      </div>
      <div class="education">
        <div class="heading-text">教育信息</div>
        <table class="per-info">
          <tbody>
            <tr class="border">
              <td>2016-2020</td>
              <td>本科</td>
            </tr>
            <tr class="border">
              <td>2016-2020</td>
              <td>高中</td>
            </tr>
          </tbody>

        </table>
      </div>
      <div class="skills">
        <div class="heading-text">关键技能</div>
        <ul>
          <li>JS</li>
          <li>html</li>
        </ul>
      </div>
    </div>
  </div>
</template>
<style lang="less">
.resume{
  background-color: cyan;
  display: flex;
  .left-side{
    flex:1;
  }
  .right-side{
    flex:1;
  }
}
</style>
