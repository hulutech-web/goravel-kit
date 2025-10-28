<template>
  <ThemeProvider :color="{ middle: { 'bg-base': '#fff' }, primary: { DEFAULT: '#1896ff' } }">


    <div class="login-box rounded-sm">

      <a-form :model="form" :wrapperCol="{ span: 24 }" @finish="login"
        class="login-form w-[400px] p-lg xl:w-[440px] xl:p-xl h-fit text-text">
        <a-form-item :required="true" name="account">
          <a-input v-model:value="form.account" autocomplete="new-username" placeholder="请输入用户名或邮箱: admin"
            class="login-input h-[40px]" />
        </a-form-item>
        <a-form-item :required="true" name="password">
          <a-input v-model:value="form.password" autocomplete="new-password" placeholder="请输入登录密码: 888888"
            class="login-input h-[40px]" type="password" />
        </a-form-item>
        <a-button htmlType="submit" class="h-[40px] w-full" type="primary" :loading="loading"> 登录 </a-button>
        <a-divider></a-divider>
        <div class="terms">
          登录即代表您同意我们的
          <span class="font-bold">用户条款 </span>、<span class="font-bold"> 数据使用协议 </span>、以及
          <span class="font-bold">Cookie使用协议</span>。
        </div>
      </a-form>
    </div>
  </ThemeProvider>
</template>
<script lang="ts" setup>

import { reactive, ref, onMounted } from 'vue';
import { useAccountStore } from '@/store';
import { ThemeProvider } from 'stepin';
import { useRouter } from 'vue-router';

import Api from "@/api";
const value1 = ref("")
const loading = ref(false);
const router = useRouter();

const form = reactive<Account.LoginForm>({
  account: '',
  password: 'admin888'
});

const emit = defineEmits<{
  (e: 'success', fields: Account.LoginForm): void;
  (e: 'failure', reason: string, fields: Account.LoginForm): void;
}>();

const accountStore = useAccountStore();
const login = async (params: Account.LoginForm) => {
  loading.value = true;
  await accountStore.login(params)
  loading.value = false
  emit('success', params);
  setTimeout(()=>{
    router.replace({path:'/workplace'});
  },500)
}

</script>
