<script lang="ts" setup></script>
<template>
  <div class="page-footer">
    <div class="links">
      <span class="link" >goravel脚手架</span>
      ×
      <span class="link">葫芦科技</span>
    </div>
    <div class="copyright">
      Copyright
      <CopyrightOutlined class="icon-copyright" /> 2025 葫芦科技  · 出品
    </div>
  </div>
</template>

<style scoped lang="less">
  .page-footer {
    text-align: center;
    @apply text-gray-400;

    .links {
      display: flex;
      justify-content: center;

      .link {
        @apply hover:text-gray-400 pl-4 pr-4;
      }
    }

    .copyright {
      margin-top: 8px;

      .icon-copyright {
        margin: 0;
      }
    }
  }
</style>
