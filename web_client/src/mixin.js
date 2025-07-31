import _ from "lodash";
import { FieldColumnWidth } from "@/assets/tool/const";
export const baseMixin = {
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        },
        userInfo() {
            return this.$store.getters.userInfo;
        }
    },
    methods: {
        curProgressInfo(id) {
            let allProgress = this.$store.state.progressTree;
            let progressList = [];
            progressList = this.getAllprogress(allProgress, progressList);
            if (_.find(progressList, { id: id })) {
                return _.find(progressList, { id: id });
            } else {
                return {};
            }
        },
        getAllprogress(progressTreeArr, treeArr) {
            if (progressTreeArr && progressTreeArr.length) {
                progressTreeArr.forEach((item) => {
                    if (item.impl_table_name === "ws_file") {
                        this.getAllprogress(item.children, treeArr);
                    } else {
                        treeArr.push(item);
                    }
                });
                return treeArr;
            } else {
                return treeArr;
            }
        },
        // 获取 流程列表宽度
        getWidth(type, field, exper) {
            let fieldType;
            if (type === "person_com") {
                fieldType = type + "_" + exper;
            } else if (field === "title") {
                fieldType = field + "_" + type;
            } else {
                fieldType = type;
            }
            // 查询缓存中是否有此列宽
            let configArrJSON = localStorage.getItem(
                `progress-${this.curProgress}`
            );
            if (configArrJSON) {
                let configArr = JSON.parse(configArrJSON);
                let targetConfig = _.find(configArr, { field: field });
                if (targetConfig) {
                    return targetConfig.width;
                }
            }
            if (fieldType) {
                return FieldColumnWidth[fieldType];
            }
            return 220;
        }
    }
};
