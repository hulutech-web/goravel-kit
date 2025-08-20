<script setup lang="ts">
import Api from "@/api";

let accountState = ref({
  id: null,
  venue_id: null,
  realname: "",
  phone: "",
  password: ""
});
const formRef = ref(null)
const route = useRoute()
const id = ref(route.params.id)
const loadData = async () => {
  accountState.value.venue_id = +(id.value)
  let res = await Api.userController.getAdmin({venue_id: id.value})
  accountState.value.realname = res.realname;
  accountState.value.phone = res.phone;
  accountState.value.password = res.password;
  accountState.value.id = res.id;
}

loadData()

const onUpdate =  () => {
  formRef.value.validateFields().then(async (e) => {
    await Api.userController.setAdmin(accountState.value)
  });
}
</script>

<template>
  <div>
    <a-divider>系统账号</a-divider>
    <a-row>
      <a-col :span="12">

        <a-form ref="formRef" :model="accountState">
          <a-form-item
              label="场馆ID"
              name="venue_id"
              :rules="[{ required: true,message:'场馆ID' }]"
          >
            <a-input v-model:value="accountState.venue_id" disabled></a-input>
          </a-form-item>

          <a-form-item
              label="用户"
              name="id"
              :rules="[{ required: true,message:'用户ID' }]"
          >
            <UserSelect v-model:id="accountState.id" v-model:username="accountState.realname"></UserSelect>
          </a-form-item>

          <a-form-item
              label="管理员姓名"
              name="realname"
              :rules="[{ required: true,message:'必填' }]"
          >
            <a-input v-model:value="accountState.realname"></a-input>
          </a-form-item>
          <a-form-item
              label="登录手机号"
              name="phone"
              :rules="[{ required: true ,message:'必填'}]"
          >
            <a-input v-model:value="accountState.phone"></a-input>
          </a-form-item>
          <a-form-item
              label="登录密码"
              name="password"
              :rules="[{ required: true ,message:'必填'}]"
          >
            <a-input type="password" v-model:value="accountState.password"></a-input>
          </a-form-item>
          <a-form-item>
            <a-space>
              <a-button type="primary" @click="onUpdate">提交</a-button>
            </a-space>
          </a-form-item>
        </a-form>

      </a-col>
    </a-row>
  </div>
</template>

<style scoped lang="less">

</style>