<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="注册">
          <b-form>
            <b-form-group label="姓名">
              <b-form-input
                v-model="$v.user.name.$model"
                type="text"
                placeholder="请输入用户名(选填)"
              ></b-form-input>
            </b-form-group>

            <b-form-group label="手机号">
              <b-form-input
                v-model="$v.user.telephone.$model"
                type="number"
                placeholder="请输入手机号"
                state="validateState('telephone')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('telephone')">
                手机号必须为11位
              </b-form-invalid-feedback>
            </b-form-group>

            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                placeholder="请输入密码"
                state="validateState('password')"
              ></b-form-input>

              <b-form-invalid-feedback :state="validateState('password')">
                密码至少为6位
              </b-form-invalid-feedback>
            </b-form-group>

            <b-form-group>
              <b-button @click="register" variant="outline-primary" block>注册</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength } from 'vuelidate/lib/validators';
import customValidator from '@/helper/validator';

export default {
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: '',
      },
      // shwoTelephoneValidata: false, // 手机号校验结果
    };
  },
  // vuelidate定义了一些model
  // 然后根据以下validations方法生成了一些属性
  // 每个属性有自己的model
  validations: {
    user: {
      name: {},
      telephone: {
        required,
        // minLength: minLength(11),
        // maxLength: maxLength(11),
        telephone: customValidator.telephoneValidate,
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    validateState(name) {
      // 解构赋值
      const { $dirty, $error } = this.$v.user[name];
      // 当与表单交互时$dirty变为true，刚开始初始化时为false
      return $dirty ? !$error : null;
    },
    register() {
      console.log('register');
    },
  },
};
</script>
<style lang="scss" scoped></style>
