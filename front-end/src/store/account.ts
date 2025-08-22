import {defineStore} from 'pinia';
import {request} from '@/utils/request';
import {useMenuStore} from './menu';
import {useAuthStore} from '@/plugins';
import {useLoadingStore} from './loading';
import Api from "@/api"

export interface Profile {
    user: Account;
    permissions: string[];
    roles: string[];
}

export interface Account {
    username: string;
    avatar: string;
}

export type TokenResult = {
    token: string;
    expires: number;
};
export const useAccountStore = defineStore('account', {
    state() {
        return {
            account: {} as Account,
            permissions: [] as string[],
            roles: [] as string[],
            logged: true,
            token: "" as string,
        };
    },
    actions: {
        async login(obj: Account.LoginForm) {

            const res = await Api.authController.login(obj)
            if (res) {
                this.logged = true;
                // @ts-ignore
                this.token = res.token;
                // @ts-ignore
                localStorage.setItem("token", res.token)
                // @ts-ignore
                localStorage.setItem("user", JSON.stringify(res.user))
                await useMenuStore().getMenuList();
                return res;
            }
        },
        async logout() {
            return new Promise<boolean>((resolve) => {
                localStorage.removeItem('token');
                localStorage.removeItem('user');
                this.logged = false;
                resolve(true);
            });
        },
        async profile() {
            const {setAuthLoading} = useLoadingStore();
            setAuthLoading(true);
            let res = await Api.userController.own()

            const {setAuthorities} = useAuthStore();
            const {setRoles} = useAuthStore();

            // @ts-ignore
            const {user, permissions, roles} = res;

            this.account = user;
            this.permissions = permissions;
            this.roles = roles;
            setRoles(roles)
            setAuthorities(permissions);
            setAuthLoading(false)
            return res;
        },
        setLogged(logged: boolean) {
            this.logged = logged;
        },
    },
});
