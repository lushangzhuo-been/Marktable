import Vue from "vue";
import VueRouter from "vue-router";
import store from "../common/store";
import { Form, Message } from "element-ui";
Vue.use(VueRouter);
// 解决路由重复等问题
const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
    return originalPush.call(this, location).catch((err) => err);
};
// 无权限页面
const noPermission = () => import("@/pages/common/no_permissions_progress.vue");
// vue-flow
const vueFlow = () => import("@/pages/common/vue_flow");
// 登陆页面
const login = () => import("@/pages/common/login/index");
// 注册页面
const register = () => import("@/pages/common/register/index_new.vue");
// 重置密码
const resetPwd = () => import("@/pages/common/reset_pwd/index");
// 退出登录
const logout = () => import("@/pages/common/logout/index");
// 菜单
const home = () => import("@/pages/common/menu/index");
// 流程
const progress = () => import("@/pages/progress/index");
// 详情单路由
const TaskDetail = () => import("@/pages/progress/detail/single_route/index");
// 文件夹信息管理
const progressGroupInfo = () =>
    import("@/pages/progress/progress_group_info/index");
// 流程基本信息修改 progressSetting
const progressInfo = () => import("@/pages/progress/progress_info/index");
// 自定义规则
const customerRule = () => import("@/pages/progress/customer_rule/index");
const progressSetting = () => import("@/pages/progress/progress_setting/index");
// 私有流程-角色&成员
const progressMember = () =>
    import("@/pages/progress/progress_private_member_role/index");
// 我的首页
const myHomePage = () => import("@/pages/my_homepage/index");
// 我的待办
const myTodo = () => import("@/pages/my_todo/index");
// 空间详情
const spaceDetail = () => import("@/pages/space/space_detail/index");
// 空间首页
const worckspacePage = () => import("@/pages/space/workspace_page/index");
// 空间成员
const spaceMember = () => import("@/pages/space/space_member/index");
// 我的通知
const messageNotice = () => import("@/pages/message_notice/index");
// 个人中心-基本信息
const basicInfo = () => import("@/pages/personal_center/basic_info");
// 个人中心-修改密码
const modifyPwd = () => import("@/pages/personal_center/modify_pwd");
// 个人中心-更换邮箱
const updateEmail = () => import("@/pages/personal_center/update_email");
// 流程-仪表盘
const dashboard = () => import("@/pages/progress/dashboard/index");
// 节点
// demo
// const node = () => import("@/pages/demo/node/index");
// const quill = () => import("@/pages/demo/quill");
// 卡片看板
const cardBoard = () => import("@/pages/progress/card_board/index");
// 全局弹窗
import { JumpDialogBox } from "@/assets/tool/login";

const routes = [
    {
        path: "/vueFlow",
        component: vueFlow,
        name: "vueFlow"
    },
    {
        path: "/login",
        component: login,
        name: "login"
    },
    {
        path: "/register",
        component: register,
        name: "register"
    },
    {
        path: "/reset/pwd",
        component: resetPwd,
        name: "resetPwd"
    },
    {
        path: "/logout",
        component: logout,
        name: "logout"
    },
    {
        path: "/",
        component: home,
        name: "home",
        redirect: "/home",
        children: [
            {
                // 无权限页码
                path: "no_permissions_progress",
                component: noPermission,
                name: "noPermission",
                meta: {
                    requireAuth: false
                }
            },
            {
                // 我的首页
                path: "home",
                component: myHomePage,
                name: "myHomePage",
                meta: {
                    requireAuth: true
                }
            },
            {
                // 我的代办
                path: "todo",
                component: myTodo,
                name: "myTodo",
                meta: {
                    requireAuth: true
                }
            },
            {
                // 我的通知
                path: "notice",
                component: messageNotice,
                name: "messageNotice",
                meta: {
                    requireAuth: true
                }
            },
            {
                // 个人中心-基本信息
                path: "personal/center/basic/info",
                component: basicInfo,
                name: "basicInfo",
                meta: {
                    requireAuth: true
                }
            },
            {
                // 个人中心-修改密码
                path: "personal/center/modify/pwd",
                component: modifyPwd,
                name: "modifyPwd",
                meta: {
                    requireAuth: true
                }
            },
            {
                // 个人中心-更新邮箱
                path: "personal/center/update/email",
                component: updateEmail,
                name: "updateEmail",
                meta: {
                    requireAuth: true,
                    module: "updateEmail"
                }
            },
            {
                // 空间更多-空间详情
                path: "space/detail",
                component: spaceDetail,
                name: "spaceDetail",
                meta: {
                    requireAuth: true,
                    module: "spaceDetail"
                }
            },
            {
                // 空间更多-空间成员
                path: "space/member",
                component: spaceMember,
                name: "spaceMember",
                meta: {
                    requireAuth: true,
                    module: "spaceMember"
                }
            },
            {
                // 空间首页-空间信息
                path: "workspace/:id",
                component: worckspacePage,
                name: "worckspacePage",
                meta: {
                    requireAuth: true,
                    module: "worckspacePage"
                }
            },
            {
                // 点击流程
                path: "progress/:wsId/:id",
                component: progress,
                name: "progress",
                meta: {
                    requireAuth: true,
                    module: "progress"
                }
            },
            {
                // 流程更多-基本信息修改
                path: "progress/:wsId/:id/basicinfo/update",
                component: progressInfo,
                name: "progress",
                meta: {
                    requireAuth: true,
                    module: "progressInfo"
                }
            },
            {
                // 流程更多-流程设置
                path: "progress/:wsId/:id/setting",
                component: progressSetting,
                name: "progress",
                meta: {
                    requireAuth: true,
                    module: "progressSetting"
                }
            },
            {
                // 流程更多-流程设置
                path: "progress/:wsId/:id/customer_rule",
                component: customerRule,
                name: "progress",
                meta: {
                    requireAuth: true,
                    module: "customerRule"
                }
            },
            {
                // 私有流程更多-角色&成员
                path: "progress/:wsId/:id/member",
                component: progressMember,
                name: "progress",
                meta: {
                    requireAuth: true,
                    module: "progressMember"
                }
            },
            {
                // 文件夹更多-基本信息修改
                path: "progress/:wsId/:id/group/basicinfo/update",
                component: progressGroupInfo,
                name: "progress",
                meta: {
                    requireAuth: true,
                    module: "progressGroupInfo"
                }
            },
            {
                path: "/dashboard",
                component: dashboard
            }
        ]
    },
    // 详情单路由
    {
        path: "/task/detail/:wsid/:id/:taskid",
        component: TaskDetail,
        name: "TaskDetail"
    },
    // demo
    // cardBoard
    {
        path: "/card",
        component: cardBoard,
        name: "cardBoard"
    }
    // {
    //     path: "/node",
    //     component: node
    // },
    // {
    //     path: '/quill',
    //     component: quill,
    //     name: 'quill',
    // },
];

var router = new VueRouter({
    routes
});

router.beforeEach((to, from, next) => {
    if (!to.name) {
        JumpDialogBox.install({
            code: 403,
            msg: "路由不存在"
        });
    }
    const userInfo = store.getters.userInfo;
    store.commit("clearToken"); // 取消请求
    if (to.meta.requireAuth) {
        // 进入该路由需要登录
        if (localStorage.getItem("token")) {
            // 检查token是否存在，不存在切到登录页重新登录
            if (userInfo && Object.keys(userInfo).length) {
                next();
            } else {
                store.dispatch("fetchUserInfo").then(() => {
                    next();
                });
            }
        } else {
            next({
                path: "/login",
                query: {
                    redirect: to.fullPath
                }
            });
        }
    } else {
        next();
    }
});
export default router;
