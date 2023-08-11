<script setup lang="ts">
import { NAvatar, NDataTable } from 'naive-ui'
import { generateEductionColumns, generatePersonalColumns, generateJobColumns } from './index.data'
const educationInfo = [
  {
    time: '2020-12~2023-1',
    school: '大学'
  },
  {
    time: '2011-12~2021-1',
    school: '高中'
  }
]
const personalInfo = [
  {
    birth: '2012-123',
    address: '上海'
  }
]
const jobInfo = [
  {
    time: '2020-12~2023-1',
    company: '甲方',
    position: '前端'
  },
  {
    time: '2020-12~2023-1',
    company: '乙方',
    position: '后端'
  }
]
</script>
<template>
  <div class="resume">
    <div class="left-side">
      <div class="profile">
        <n-avatar :size="120" round src="/avatar.jpg" />
      </div>
      <div class="contact">
        <h2>联系方式</h2>
        <p class="para">
          家庭住址：72-2 <br />
          +8612344577 <br />
          +1233 <br />
          example@gmail.com <br />
          www.baidu.com <br />
        </p>
      </div>
      <div class="expert">
        <h2>专业领域</h2>
        <p>
          JS <br />
          语文 <br />
          数学 <br />
          英语 <br />
          政治 <br />
        </p>
      </div>
      <div class="skill">
        <h2>我的技能</h2>
        <p>
          技能1 <br />
          语文1 <br />
          数学1 <br />
          英语1 <br />
          政治1 <br />
        </p>
      </div>
      <div class="hobby">
        <h2>我的爱好</h2>
        <p>
          爱好1 <br />
          爱好2 <br />
          爱好3 <br />
          爱好4 <br />
          政治1 <br />
        </p>
      </div>
    </div>
    <div class="split"></div>
    <div class="right-side">
      <h1>张三</h1>
      <div class="block about">
        <h2>关于我</h2>
        <p class="para">2017年毕业 ，专业软件工程</p>
      </div>
      <div class="block personal">
        <h2>个人信息</h2>
        <n-data-table :columns="generatePersonalColumns()" :data="personalInfo" />
      </div>
      <div class="block education">
        <h2>教育信息</h2>
        <n-data-table :columns="generateEductionColumns()" :data="educationInfo" />
      </div>
      <div class="block job">
        <h2>工作经历</h2>
        <n-data-table :columns="generateJobColumns()" :data="jobInfo" />
      </div>
      <div class="block skills">
        <h2>专业技能</h2>
        <ul>
          <li>JS</li>
          <li>html</li>
          <li>goLang</li>
        </ul>
      </div>
    </div>
  </div>
</template>
<style lang="less" scoped>
h1 {
  font-weight: bold;
  font-size: 2em;
}
.resume {
  box-sizing: border-box;
  display: flex;
  background-color: #b0c4de;
  border: 1px solid #000;
  padding: 40px;
  .left-side {
    flex-grow: 0.3;
  }
  .right-side {
    flex-grow: 1;
    margin-top: 10px;
  }
  .split {
    width: 1px;
    background-color: #fff;
    margin: -40px 20px;
  }
  .block {
    margin-bottom: 20px;
    > h2 {
      margin-bottom: 10px;
    }
  }
}
</style>
