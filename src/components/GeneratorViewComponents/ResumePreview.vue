<script setup lang="ts">
import { NAvatar, NDataTable, NDescriptions, NDescriptionsItem, NDivider } from 'naive-ui'
import { generateEductionColumns, generateJobColumns } from './index.data'
const educationInfo = [
  {
    time: '2020-12~2023-1',
    school: '大学',
    master: '学士'
  },
  {
    time: '2011-12~2021-1',
    school: '高中',
    master: '-'
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
      <div class="baseInfo block">
        <h2>基础信息</h2>
        <div class="item">
          <n-descriptions label-placement="left" :column="1">
            <n-descriptions-item label="年龄"> 23 </n-descriptions-item>
            <n-descriptions-item label="工作年限"> 2年 </n-descriptions-item>
            <n-descriptions-item label="联系电话"> 13213556789 </n-descriptions-item>
            <n-descriptions-item label="邮箱"> earha@gmail.com </n-descriptions-item>
          </n-descriptions>
        </div>
      </div>
      <n-divider />
      <div class="skills block">
        <h2>专业技能</h2>
        <div class="item">
          <n-descriptions label-placement="left" :column="1">
            <n-descriptions-item label="JavaScript"> 精通 </n-descriptions-item>
            <n-descriptions-item label="GoLang"> 了解 </n-descriptions-item>
            <n-descriptions-item label="UE"> 熟悉 </n-descriptions-item>
            <n-descriptions-item label="React"> 了解 </n-descriptions-item>
          </n-descriptions>
        </div>
      </div>
      <n-divider />
      <div class="jobRequire block">
        <h2>求职意向</h2>
        <div class="item">
          <n-descriptions label-placement="left" :column="1">
            <n-descriptions-item label="意向岗位"> 前端 </n-descriptions-item>
            <n-descriptions-item label="意向城市"> 上海 </n-descriptions-item>
            <n-descriptions-item label="薪资要求"> 100-200 </n-descriptions-item>
            <n-descriptions-item label="求职状态"> 在职 </n-descriptions-item>
          </n-descriptions>
        </div>
      </div>
    </div>
    <div class="split"></div>
    <div class="right-side">
      <h1 class="name">张三</h1>
      <div class="block breief">
        <p class="para">2017年毕业 ，专业软件工程</p>
      </div>
      <div class="block education">
        <h2>教育背景</h2>
        <n-data-table :columns="generateEductionColumns()" :data="educationInfo" />
      </div>
      <div class="block job">
        <h2>工作经历</h2>
        <n-data-table :columns="generateJobColumns()" :data="jobInfo" />
      </div>
      <div class="block project">
        <h2>项目经验</h2>
        <n-descriptions label-placement="top"  :column="1">
          <n-descriptions-item label="2020~2021">
            Lorem ipsum dolor sit, amet consectetur adipisicing elit. Obcaecati architecto mollitia
            exercitationem eum, atque vero nulla dolores delectus! Dolore quo minima itaque sed
            maxime nulla, reprehenderit labore dicta! Laudantium, iste?
          </n-descriptions-item>
          <n-descriptions-item label="2022~2023">
            Lorem ipsum dolor sit, amet consectetur adipisicing elit. Obcaecati architecto mollitia
            exercitationem eum, atque vero nulla dolores delectus
          </n-descriptions-item>
        </n-descriptions>
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
