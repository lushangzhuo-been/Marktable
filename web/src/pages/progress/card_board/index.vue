<template>
    <div>
        <div class="card-component">
            <!-- 维度枚举 -->
            <group-by-temp
                v-if="groupByFieldInfo && groupByFieldInfo.field_key"
                class="group-by-block"
                :groupByFieldInfo="groupByFieldInfo"
                :groupByList="groupByList"
                :cardFilterDown="cardFilterDown"
                :currentTab="currentTab"
                @check-field-search="checkFieldSearch"
            ></group-by-temp>
            <!-- 列表 -->
            <div class="group-list">
                <div
                    v-for="(groupItem, groupIndex) in cardsGroup"
                    :key="groupIndex"
                    class="group-item"
                    :class="{
                        'could-drop': groupItem.disableDrop === 'no',
                        'could-not-drop': groupItem.disableDrop === 'yes'
                    }"
                >
                    <div class="group-title">
                        <div
                            class="group-title-content"
                            :style="{
                                color: '#fff',
                                backgroundColor: rgbToRgba(
                                    `${groupItem.color || '#fff'}`,
                                    0.7
                                )
                            }"
                        >
                            {{ groupItem.name }}
                        </div>
                        {{ groupItem?.cnt || "" }}
                    </div>
                    <div class="list-item">
                        <div
                            class="list-content"
                            :class="{
                                'could-drop': groupItem.disableDrop === 'no',
                                'could-not-drop':
                                    groupItem.disableDrop === 'yes'
                            }"
                            @scroll="(e) => onScroll(e, groupItem)"
                        >
                            <draggable
                                dragClass="dragClass"
                                ghostClass="ghostClass"
                                chosenClass="chosenClass"
                                v-model="groupItem.list"
                                group="aaa"
                                animation="300"
                                @start="(e) => onStart(e, groupItem)"
                                @end="onEnd"
                                @add="onAdd"
                                :move="onMove"
                                class="drag-draggable"
                            >
                                <transition-group
                                    class="drag-group"
                                    :id="groupItem.id"
                                    :class="groupItem.class"
                                >
                                    <div
                                        class="card-item"
                                        v-for="(
                                            cardItem, cardIndex
                                        ) in groupItem.list"
                                        :key="cardIndex"
                                        :id="cardItem._id"
                                    >
                                        <card-item
                                            class="card-item-temp"
                                            :fieldKeyConfig="fieldKeyConfig"
                                            :cardDetail="cardItem"
                                            @open-detail="openDetail"
                                        ></card-item>
                                    </div>
                                </transition-group>
                            </draggable>
                            <div
                                v-if="groupItem.disableDrop === 'yes'"
                                class="disabled-block"
                            >
                                不能到达此状态
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <node-operation-dialog
            ref="NodeOperationDialog"
            @confirm-next-state="confirmNextState"
            @cancel-submit="refreshChangeCardsList"
            :nodeOperationConfig="nodeOperationConfig"
            :nodeOperationData="nodeOperationData"
        ></node-operation-dialog>
        <detail-drawer ref="DetailDrawer" @refresh="refresh"></detail-drawer>
    </div>
</template>

<script>
import _ from "lodash";
import draggable from "vuedraggable";
import CardItem from "@/pages/progress/card_board/card_item.vue";
import GroupByTemp from "./group_by_temp";
import NodeOperationDialog from "@/pages/progress/detail/components/node_operation_dialog";
import DataHandle from "@/pages/progress/card_board/data_handle.js";
import { baseMixin } from "@/mixin.js";
import api from "@/common/api/module/progress";
import configApi from "@/common/api/module/progress_setting";
import dashApi from "@/common/api/module/progress_dashboard";
import DetailDrawer from "@/pages/progress/detail/index.vue";
import { rgbToRgba } from "@/assets/tool/func";
export default {
    mixins: [baseMixin],
    components: {
        draggable,
        CardItem,
        DetailDrawer,
        NodeOperationDialog,
        GroupByTemp
    },
    props: {
        currentTab: {
            type: String,
            default: ""
        },
        curTabItems: {
            type: Object,
            default: () => {}
        },
        filterParams: {
            type: Array,
            default: () => []
        },
        sortOrder: {
            type: Array,
            default: () => []
        },
        filterRelation: {
            type: String,
            default: ""
        },
        cardFilterDown: {
            type: [String, Object],
            default: ""
        }
    },
    data() {
        return {
            groupByFieldInfo: {}, //分组字段信息
            groupByList: [], // 分组字段的枚举列表
            statusGroupInfo: [], // 当前流程拥有的所有状态
            groupByFieldEnumInfo: {}, // 分组枚举值详情信息
            cardsGroup: [], // 看板循环信息
            fieldKeyConfig: [], // 字段config
            pageSize: 10,
            reachBottomDistance: 40,
            allStatusConfig: [], // 状态切换全配置
            couldHandleStatus: [], // 抓起时赋值 可以放入的列id
            nodeOperationConfig: [],
            nodeOperationData: {},
            operationCardInfo: {}, //当前正在操作的卡片
            cardFilter: []
        };
    },
    watch: {
        currentTab: {
            handler() {
                // 重置
                this.groupByFieldInfo = {};
            }
        },

        filterParams: {
            handler() {
                if (this.curTabItems.mode !== "card") {
                    return;
                }
                this.confimGroupBy(this.groupByFieldInfo);
                this.getAllStatusCardsList();
            },
            deep: true
        },
        filterRelation: {
            handler() {
                if (this.curTabItems.mode !== "card") {
                    return;
                }
                this.getAllStatusCardsList();
            },
            deep: true
        },
        cardFilterDown: {
            handler() {
                if (this.curTabItems.mode !== "card") {
                    return;
                }
                this.getAllStatusCardsList();
            },
            deep: true
        }
    },
    mounted() {},
    methods: {
        // 新增刷新
        refreshForAdd() {
            // 可能会有新增的枚举值
            this.confimGroupBy(this.groupByFieldInfo); // 刷新枚举->刷全列
        },
        // 刷新  编辑后刷新
        refresh(changeInfo) {
            // 获取起始状态列信息
            if (changeInfo.status.start && changeInfo.status.end) {
                // 刷新前后两个变动列， 可能前后一致
                let fromCardInfo = _.find(this.cardsGroup, {
                    id: changeInfo.status.start
                });
                let toCardInfo = _.find(this.cardsGroup, {
                    id: changeInfo.status.end
                });
                this.checkCurListAll(fromCardInfo);
                this.checkCurListAll(toCardInfo);
            } else if (
                !(changeInfo.status.start && changeInfo.status.end) &&
                changeInfo.other
            ) {
                // 刷新当前列的状态 ,  可能需要重新刷新枚举值
                let curCardInfo = _.find(this.cardsGroup, {
                    id: changeInfo.defaultStatus
                });
                this.confimGroupBy(this.groupByFieldInfo, "by-update"); // 刷新枚举
                this.checkCurListAll(curCardInfo); // 刷新当前状态列
            } else {
                // 不用刷新
            }
        },
        rgbToRgba(color, opacity) {
            return rgbToRgba(color, opacity);
        },
        saveGroupView() {
            let groupCard;
            if (this.groupByFieldInfo.id) {
                groupCard = {
                    group_axis: this.groupByFieldInfo.field_key,
                    group_value:
                        parseInt(this.groupByFieldEnumInfo.user_id) ||
                        parseInt(this.groupByFieldEnumInfo.status_id) ||
                        this.groupByFieldEnumInfo.name ||
                        "" // 维度值  角色取user_id
                };
            } else {
                // 无分组
                groupCard = "";
            }
            // 更新卡片分组
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.currentTab, //视图id
                card_group: JSON.stringify(groupCard)
            };
            api.updateViewCardGroup(params);
        },
        openDetail(cardItem) {
            this.$refs.DetailDrawer.openDrawer("tmpl_detail", cardItem);
        },
        getConfigTypeInfo(paramsColumn) {
            this.getRuleStatusList("isFirst");
            this.getColConfig(paramsColumn);
            this.getAllStatusConfig();
        },
        onStart(e, groupItem) {
            this.couldHandleStatus = [];
            // 取groupItem.id，获取可操作状态列表allStatusConfig[groupItem.id]，更新遍历cardsGroup
            let curStatus = groupItem.id;
            let couldHandleStatusObj = this.allStatusConfig[curStatus] || {};
            this.couldHandleStatus = Object.keys(couldHandleStatusObj) || [];
            this.cardsGroup.forEach((cardList) => {
                if (
                    _.indexOf(this.couldHandleStatus, String(cardList.id)) >
                        -1 ||
                    cardList.id === groupItem.id
                ) {
                    this.$set(cardList, "disableDrop", "no");
                } else {
                    this.$set(cardList, "disableDrop", "yes");
                }
            });
        },
        onEnd(e) {
            this.cardsGroup.forEach((cardList) => {
                this.$set(cardList, "disableDrop", "");
            });
        },
        // 这里用来提交接口转换状态
        onAdd(cardInfo) {
            // 获取from.id  to.id  获取入参用的记录id
            let fromId = cardInfo.from.id;
            let toId = cardInfo.to.id;
            let recordInfo;
            // 哪里来放哪里
            if (fromId !== toId) {
                let fromInfo = this.allStatusConfig[fromId] || {};
                recordInfo = fromInfo[toId] || {};
            }
            // 状态切换，检查字段
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                step_id: recordInfo.id,
                id: cardInfo.clone.id
            };
            api.getBtnScreen(params).then((res) => {
                if (res && res.resultCode === 200) {
                    if (!res.data || (res.data && !res.data.length)) {
                        // 如果没有限制字段，直接放行， 设置状态为选中的，并且刷新表格内容
                        this.switchStepAction(cardInfo, recordInfo);
                    } else {
                        this.operationCardInfo = cardInfo.clone;
                        // res.data为获取的字段
                        this.nodeOperationConfig = res.data;
                        this.$refs.NodeOperationDialog.openDialog(recordInfo);
                        document.body.click(); //关闭poppover
                        // 获取节点配置数据
                        this.getNodeOperationData(recordInfo);
                    }
                }
            });
        },
        // 做阻止放入逻辑
        onMove(e) {
            let moveInfo = e;
            // 如果是自己本身  或者目标列在this.couldHandleStatus中  就可以放入
            if (
                _.indexOf(this.couldHandleStatus, moveInfo.to.id) > -1 ||
                moveInfo.to.id === moveInfo.from.id
            ) {
                return true;
            } else {
                return false;
            }
        },
        // 切换状态赋值
        switchStepAction(cardInfo, recordInfo) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                step_id: recordInfo.id,
                id: cardInfo.clone.id
            };
            api.switchStepAction(params).then((res) => {
                this.refreshChangeCardsList(recordInfo);
            });
        },
        // 获取弹窗节点配置数据
        getNodeOperationData(item) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                step_id: item.id,
                id: this.operationCardInfo.id
            };
            api.getFormData(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.nodeOperationData = res.data;
                } else {
                    this.nodeOperationData = {};
                }
            });
        },
        // 弹窗跳出后填写完毕，提交
        confirmNextState(btn, form) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                step_id: btn.id,
                id: this.operationCardInfo.id,
                // 解构被编辑的参数
                ...form
            };
            api.switchStepAction(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$refs.NodeOperationDialog.closeDialog();
                }
                this.refreshChangeCardsList(btn);
            });
        },
        // 拖拽刷新前后变动状态
        refreshChangeCardsList(btn) {
            let fromCardInfo = _.find(this.cardsGroup, {
                id: btn.start_status_id
            });
            let toCardInfo = _.find(this.cardsGroup, { id: btn.end_status_id });
            this.checkCurListAll(fromCardInfo);
            this.checkCurListAll(toCardInfo);
        },
        confimGroupBy: _.debounce(function (groupInfo, order) {
            // 如果是无分组 需要直接调卡片列表 否则调枚举值列表
            this.groupByFieldInfo = groupInfo;
            if (groupInfo.id) {
                let param = {
                    filter: this.filterParams.length
                        ? JSON.stringify(this.filterParams)
                        : "",
                    lor: "filter_and",
                    mode: "base_histogram",
                    order: "desc",
                    title: "title",
                    tmpl_id: parseInt(this.curProgress),
                    userid: this.userInfo.id,
                    view_id: parseInt(this.currentTab),
                    ws_id: this.curSpace.id,
                    show_empty: "yes", //枚举值中返回空值
                    xMode: groupInfo.mode, // 分组字段类型
                    x_axis: groupInfo.field_key, // 分组字段
                    y_axis: "count" //  写死
                };
                let selectAll = {
                    filed_mode: "all",
                    name: "全部"
                };
                dashApi.getPrviewChart(param).then((res) => {
                    if (res.resultCode === 200) {
                        this.groupByList = res.data.board_statistics;
                        // 拼接全部
                        this.groupByList.unshift(selectAll);
                    } else {
                        this.groupByList = [];
                        this.groupByList.unshift(selectAll);
                    }
                });
            } else {
                // 无分组
                // 枚举字段为空
                if (order !== "by-update") {
                    this.groupByList = [];
                    this.checkFieldSearch(DataHandle.noGroupInfo);
                }
            }
        }, 300),
        // 当前流程拥有的全部状态
        getRuleStatusList(key) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            configApi.getRuleStatusList(params).then((res) => {
                if (res.resultCode === 200) {
                    this.statusGroupInfo = res.data || [];
                    this.cardsGroup = res.data;
                } else {
                    this.statusGroupInfo = [];
                    this.cardsGroup = [];
                }
                // 这里之后才能开始调卡片列表
                if (key === "isFirst") {
                    // 挂载后的第一次请求,判断当前所选枚举值是否为空
                    this.checkFieldSearch(this.groupByFieldEnumInfo);
                }
            });
        },
        getAllStatusConfig() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: parseInt(this.curProgress)
            };
            api.getAllStatusConfig(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.allStatusConfig = res.data || [];
                } else {
                    this.allStatusConfig = [];
                }
            });
        },
        // 切换枚举值第一次请求
        checkFieldSearch(info) {
            this.groupByFieldEnumInfo = info;
            this.getAllStatusCardsList();
        },
        getAllStatusCardsList: _.debounce(function () {
            // 因为使用的v-show  需要判断当前view是否为卡片
            // qes  单选时入参  取name  返回值错误
            this.statusGroupInfo.forEach((status) => {
                let filterDown = {
                    x_axis: "status", // 状态
                    x_value: parseInt(status.id),
                    group_axis:
                        this.groupByFieldEnumInfo.filed_mode === "all"
                            ? ""
                            : this.groupByFieldInfo.field_key || "", // 维度
                    group_value:
                        this.groupByFieldEnumInfo.filed_mode === "all"
                            ? ""
                            : parseInt(this.groupByFieldEnumInfo.user_id) ||
                              parseInt(this.groupByFieldEnumInfo.status_id) ||
                              this.groupByFieldEnumInfo.name ||
                              "" // 维度值  角色取user_id
                };
                let params = {
                    board_id: 0,
                    ws_id: this.curSpace.id,
                    tmpl_id: parseInt(this.curProgress),
                    filter: JSON.stringify(this.filterParams),
                    sort_order: JSON.stringify(this.sortOrder),
                    lor: this.filterRelation,
                    page_num: 1,
                    page_size: this.pageSize,
                    filter_down: JSON.stringify(filterDown)
                };
                api.getProgressListData(params).then((res) => {
                    if (res && res.resultCode === 200 && res.data) {
                        let curCard = _.find(this.cardsGroup, {
                            id: status.id
                        });
                        this.$set(curCard, "list", res.data.list || []);
                        this.$set(curCard, "cnt", res.data.cnt || 0);
                        this.$set(curCard, "pageNum", 1);
                    } else {
                        let curCard = _.find(this.cardsGroup, {
                            id: status.id
                        });
                        this.$set(curCard, "list", []);
                        this.$set(curCard, "cnt", 0);
                        this.$set(curCard, "pageNum", 1);
                    }
                });
            });
        }, 400),
        // 获取字段信息
        getColConfig(paramsColumn) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                module: "list",
                columns: paramsColumn
            };
            api.getProgressColConfig(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    Object.keys(res.data).length
                ) {
                    // 按照res.data全部显示到卡片上
                    this.fieldKeyConfig = res.data || [];
                } else {
                    this.fieldKeyConfig = [];
                }
            });
        },
        onScroll: _.debounce(function (e, cardList) {
            let scrollTop = e.target.scrollTop;
            let scrollHeight = e.target.scrollHeight;
            let offsetHeight = Math.ceil(
                e.target.getBoundingClientRect().height
            );
            let currentHeight =
                scrollTop + offsetHeight + this.reachBottomDistance;
            if (currentHeight >= scrollHeight) {
                if (cardList?.list?.length < cardList.pageNum * this.pageSize) {
                    // 重新请求当前全部数据
                    this.checkCurListAll(cardList);
                } else {
                    // 请求下一页数据
                    this.requestNextPage(cardList);
                }
            }
        }, 400),
        // 重新请求当前列已请求分页全部数据
        checkCurListAll(cardList) {
            if (!cardList) {
                return;
            }
            let filterDown = {
                x_axis: "status", // 状态
                x_value: parseInt(cardList.id),
                group_axis:
                    this.groupByFieldEnumInfo.filed_mode === "all"
                        ? ""
                        : this.groupByFieldInfo.field_key || "", // 维度
                group_value:
                    this.groupByFieldEnumInfo.filed_mode === "all"
                        ? ""
                        : parseInt(this.groupByFieldEnumInfo.user_id) ||
                          parseInt(this.groupByFieldEnumInfo.status_id) ||
                          this.groupByFieldEnumInfo.name ||
                          "" // 维度值  角色取user_id
            };
            let params = {
                board_id: 0,
                ws_id: this.curSpace.id,
                tmpl_id: parseInt(this.curProgress),
                filter: JSON.stringify(this.filterParams),
                sort_order: JSON.stringify(this.sortOrder),
                lor: this.filterRelation,
                page_num: 1,
                page_size: this.pageSize * cardList.pageNum,
                filter_down: JSON.stringify(filterDown)
            };
            let curCard = _.find(this.cardsGroup, {
                id: cardList.id
            });
            api.getProgressListData(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    curCard.list = res.data.list || [];
                    curCard.cnt = res.data.cnt || 0;
                } else {
                    curCard.list = [];
                    curCard.cnt = 0;
                }
            });
        },
        // 触底加载请求下一页
        requestNextPage(cardList) {
            let filterDown = {
                x_axis: "status", // 状态
                x_value: parseInt(cardList.id),
                group_axis:
                    this.groupByFieldEnumInfo.filed_mode === "all"
                        ? ""
                        : this.groupByFieldInfo.field_key || "", // 维度
                group_value:
                    this.groupByFieldEnumInfo.filed_mode === "all"
                        ? ""
                        : parseInt(this.groupByFieldEnumInfo.user_id) ||
                          parseInt(this.groupByFieldEnumInfo.status_id) ||
                          this.groupByFieldEnumInfo.name ||
                          "" // 维度值  角色取user_id
            };
            let params = {
                board_id: 0,
                ws_id: this.curSpace.id,
                tmpl_id: parseInt(this.curProgress),
                filter: JSON.stringify(this.filterParams),
                sort_order: JSON.stringify(this.sortOrder),
                lor: this.filterRelation,
                page_num: cardList.pageNum + 1,
                page_size: this.pageSize,
                filter_down: JSON.stringify(filterDown)
            };
            api.getProgressListData(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    let curCard = _.find(this.cardsGroup, {
                        id: cardList.id
                    });
                    let curList = cardList.list;
                    let nextPageArr = res.data.list;
                    let allList = curList.concat(nextPageArr);
                    curCard.list = allList;
                    curCard.pageNum = cardList.pageNum + 1;
                    curCard.cnt = res.data.cnt;
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.card-component {
    display: flex;
    .group-list {
        display: flex;
        overflow-x: auto;
        .group-title {
            box-sizing: border-box;
            width: 100%;
            height: 52px;
            line-height: 52px;
            padding: 0 16px;
            background-color: #f8f8f8;
            margin-bottom: 4px;
            .group-title-content {
                display: inline-block;
                padding: 0 16px;
                height: 30px;
                line-height: 30px;
                text-align: center;
                border-radius: 88px;
            }
        }
        .group-item {
            box-sizing: border-box;
            min-width: 286px;
            max-width: 286px;
            margin-right: 16px;
            border: 2px solid rgba($color: #000000, $alpha: 0);
            &:last-child {
                margin-right: 0;
            }
            &.could-drop {
                border: 2px dotted #96eda5;
                .card-item {
                    display: none;
                    &.chosenClass {
                        display: block;
                    }
                }
            }

            &.could-not-drop {
                .card-item {
                    display: none;
                }
            }

            .list-item {
                box-sizing: border-box;
                margin-right: 16px;
                width: 100%;
                padding: 16px 0;
                background-color: #f8f8f8;
                &:last-child {
                    margin-right: 0;
                }
                .list-content {
                    box-sizing: border-box;
                    padding: 0 12px;
                    background-color: #f8f8f8;
                    height: calc(100vh - 296px);
                    position: relative;
                    overflow-y: auto;
                    transition: 0.3s;
                    .drag-draggable {
                        display: block;
                        height: 100%;
                    }
                    .drag-group {
                        display: block;
                        height: 100%;
                        .card-item {
                            box-sizing: border-box;
                            margin-bottom: 8px;
                            &:last-child {
                                margin-bottom: 0px;
                            }
                            .card-item-temp {
                                box-sizing: border-box;
                                padding: 16px 16px 8px 16px;
                                min-height: 96px;
                                border-radius: 4px;
                                background-color: #ffffff;
                                border: 1px solid #e6e9f0;
                                display: flex;
                                flex-flow: wrap;
                                cursor: pointer;
                            }
                        }
                    }
                    .disabled-block {
                        position: absolute;
                        z-index: 1100;
                        top: 0;
                        height: 100%;
                        width: 268px;
                        text-align: center;
                        line-height: 600px;
                        background-color: #f8f8f8;
                        opacity: 0.8;
                    }
                }
            }
        }
    }
}
.ghostClass {
    // background-color: blue !important;
}
// 选中
.chosenClass {
    // background-color: red !important;
    opacity: 1 !important;
}
// 拖拽中
:deep(.dragClass) {
    opacity: 1 !important;
    outline: none !important;
    background-image: none !important;
    .card-item-temp {
        transform: rotate(4deg) !important;
        box-shadow: 2px 2px 8px 1px rgba(47, 56, 76, 0.1) !important;
    }
}
</style>
