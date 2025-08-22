import {useAuthStore} from "@/plugins";



function setupDirective(app){
    app.directive('role', {
        mounted(el, binding) {
            const {hasRole} = useAuthStore();
            const { arg, value } = binding;
            const roles = arg || value;

            if (roles && !hasRole(roles)) {
                el.style.display = 'none';
                // 或者完全移除元素: el.parentNode?.removeChild(el);
            }
        },
        updated(el, binding) {
            // 处理权限动态变化
            const {hasRole} = useAuthStore();
            const { arg, value } = binding;
            const roles = arg || value;

            if (roles && !hasRole(roles)) {
                el.style.display = 'none';
            } else {
                el.style.display = '';
            }
        }
    });
}

export {setupDirective}