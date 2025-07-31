<template>
    <div>
        <div class="slot-content">
            <!-- :data="tableData"
                :col="tableCol"
                @open-html-col="openHtmlCol"
                @open-detail="openDetail"
                @multiple-selection="multipleSelection"
                @refresh="refreshTable"
                @copy-link="copyLink"
                @delete-row="deleteRow"
                @change-column-width="changeColumnWidth" -->
            <progress-table
                :col="tableCol"
                :data="tableData"
                @open-detail="openDetail"
            ></progress-table>
        </div>
    </div>
</template>

<script>
// import ProgressTable from "@/components/table/p_table/progress_table.vue";
export default {
    components: {
        // ProgressTable
        ProgressTable: () =>
            import("@/components/table/p_table/progress_table.vue")
    },
    data() {
        return {
            tableCol: [
                {
                    field_key: "title",
                    name: "标题",
                    desc: "",
                    mode: "text_com",
                    enum_values: null,
                    expr: "",
                    can_modify: "yes",
                    required: "yes",
                    level: "level2",
                    read_only_rule: "",
                    related_tmpl_id: 0,
                    columnName: "title"
                },
                {
                    field_key: "handler",
                    name: "处理人",
                    desc: "",
                    mode: "person_com",
                    enum_values: null,
                    expr: "single",
                    can_modify: "yes",
                    required: "yes",
                    level: "level2",
                    read_only_rule: "",
                    related_tmpl_id: 0,
                    columnName: "handler"
                },
                {
                    field_key: "status",
                    name: "状态",
                    desc: "",
                    mode: "status_com",
                    enum_values: null,
                    expr: "",
                    can_modify: "yes",
                    required: "yes",
                    level: "level1",
                    read_only_rule: "",
                    related_tmpl_id: 0,
                    columnName: "status"
                },
                {
                    field_key:
                        "custom_field_233_2db3bc56dafd8d25f05e299c7e399b20",
                    name: "riqi",
                    desc: "",
                    mode: "date_com",
                    enum_values: null,
                    expr: "",
                    can_modify: "yes",
                    required: "yes",
                    level: "level3",
                    read_only_rule: "",
                    related_tmpl_id: 0,
                    columnName:
                        "custom_field_233_2db3bc56dafd8d25f05e299c7e399b20"
                },
                {
                    field_key:
                        "custom_field_233_d0527f523658f09d9aa5bf437a41ea68",
                    name: "html",
                    desc: "",
                    mode: "html_text_com",
                    enum_values: null,
                    expr: "",
                    can_modify: "yes",
                    required: "yes",
                    level: "level3",
                    read_only_rule: "",
                    related_tmpl_id: 0,
                    columnName:
                        "custom_field_233_d0527f523658f09d9aa5bf437a41ea68"
                }
            ],
            tableData: [
                {
                    _id: "67d2e6c7bb89c8d261415eff",
                    created_at: "2025-03-13 22:08:07",
                    creator: [
                        {
                            id: 46,
                            username: "liuchenghui",
                            full_name: "刘成辉举头望明月",
                            email: "894659666@qq.com",
                            phone: "18710601494",
                            avatar: "/image/avatar/46/90ece9a3-75c9-4c01-b6d4-996558d034ac.png"
                        }
                    ],
                    custom_field_233_16bc39d30e71c5916c7303346d3400dd: [
                        {
                            id: 46,
                            username: "liuchenghui",
                            full_name: "刘成辉举头望明月",
                            email: "894659666@qq.com",
                            phone: "18710601494",
                            avatar: "/image/avatar/46/90ece9a3-75c9-4c01-b6d4-996558d034ac.png"
                        }
                    ],
                    custom_field_233_2c09e04ae6c959c8f3850e4aec8b4285: [
                        {
                            id: 46,
                            username: "liuchenghui",
                            full_name: "刘成辉举头望明月",
                            email: "894659666@qq.com",
                            phone: "18710601494",
                            avatar: "/image/avatar/46/90ece9a3-75c9-4c01-b6d4-996558d034ac.png"
                        }
                    ],
                    custom_field_233_2db3bc56dafd8d25f05e299c7e399b20:
                        "2025-03-13",
                    custom_field_233_369fcdadf43b066e77c8c3b64420c32d: 1,
                    custom_field_233_530c6b11108e5f3ce723065ce0fcb3a9: {
                        name: "1",
                        url: "1"
                    },
                    custom_field_233_5578ee327d63c0c82b8b87c6ee6770f3: "1",
                    custom_field_233_701e5537ca8c9faf08e9322a5822ac21: "1",
                    custom_field_233_a916d37252a1a89313a578f2cf11a14b: ["a"],
                    custom_field_233_cae55a8e30a5115d953be3567a54c35a: [
                        "a",
                        "b"
                    ],
                    custom_field_233_d0527f523658f09d9aa5bf437a41ea68:
                        "<p>1</p>",
                    custom_field_233_e29449c4492db8bfa3d81b834359cc41: 1,
                    custom_field_233_e36aeaade67111cab2a4bc89bb668339: [
                        {
                            _id: "6795a08b280c7f7c2d4c3ae6",
                            created_at: "2025-01-26 10:40:11",
                            creator: [46],
                            custom_field_210_2323964f7d341b6138b0fd6e815f8276: [
                                41, 44, 43, 45, 42, 46
                            ],
                            custom_field_210_43cde9b906a89b80c52b4bfe526f02c3: [
                                46
                            ],
                            custom_field_210_7e305625b787498b3c1b712e12beb7cf: [
                                "111"
                            ],
                            custom_field_210_be1fc2eadfa7de817e4bcdbaf1d2fe0e: [
                                "1111111111111111",
                                "22222222222222222",
                                "333333333333333333",
                                "444444444444444444444"
                            ],
                            handler: [46],
                            is_member: "yes",
                            issue_url:
                                "http://marktable.cn:8082/#/task/detail/81/210/6795a08b280c7f7c2d4c3ae6",
                            status: 110,
                            title: "测试状态变更",
                            tmpl_id: 210,
                            updated_at: "2025-05-09 23:03:09",
                            ws_id: 81
                        }
                    ],
                    custom_field_233_f7c6a2d80960308c9e51374152f5d909:
                        "2025-03-13 00:00:00",
                    handler: [
                        {
                            id: 46,
                            username: "liuchenghui",
                            full_name: "刘成辉举头望明月",
                            email: "894659666@qq.com",
                            phone: "18710601494",
                            avatar: "/image/avatar/46/90ece9a3-75c9-4c01-b6d4-996558d034ac.png"
                        }
                    ],
                    progress_num: 0,
                    status: {
                        id: 188,
                        ws_id: 81,
                        tmpl_id: 233,
                        name: "未开始",
                        color: "rgb(24, 144, 255)",
                        mode: "first"
                    },
                    title: "1",
                    tmpl_id: 233,
                    updated_at: "2025-05-25 19:02:41",
                    ws_id: 81
                },
                {
                    _id: "67d429ee88791e98e4e27cb9",
                    created_at: "2025-03-14 21:06:54",
                    creator: [
                        {
                            id: 46,
                            username: "liuchenghui",
                            full_name: "刘成辉举头望明月",
                            email: "894659666@qq.com",
                            phone: "18710601494",
                            avatar: "/image/avatar/46/90ece9a3-75c9-4c01-b6d4-996558d034ac.png"
                        }
                    ],
                    custom_field_233_16bc39d30e71c5916c7303346d3400dd: null,
                    custom_field_233_2c09e04ae6c959c8f3850e4aec8b4285: null,
                    custom_field_233_2db3bc56dafd8d25f05e299c7e399b20: "",
                    custom_field_233_369fcdadf43b066e77c8c3b64420c32d: "",
                    custom_field_233_530c6b11108e5f3ce723065ce0fcb3a9: {
                        name: "",
                        url: ""
                    },
                    custom_field_233_5578ee327d63c0c82b8b87c6ee6770f3: "",
                    custom_field_233_701e5537ca8c9faf08e9322a5822ac21: "",
                    custom_field_233_a916d37252a1a89313a578f2cf11a14b: [],
                    custom_field_233_cae55a8e30a5115d953be3567a54c35a: [],
                    custom_field_233_d0527f523658f09d9aa5bf437a41ea68: "",
                    custom_field_233_e36aeaade67111cab2a4bc89bb668339: [],
                    custom_field_233_f7c6a2d80960308c9e51374152f5d909: "",
                    handler: [
                        {
                            id: 46,
                            username: "liuchenghui",
                            full_name: "刘成辉举头望明月",
                            email: "894659666@qq.com",
                            phone: "18710601494",
                            avatar: "/image/avatar/46/90ece9a3-75c9-4c01-b6d4-996558d034ac.png"
                        }
                    ],
                    progress_num: 0,
                    status: {
                        id: 188,
                        ws_id: 81,
                        tmpl_id: 233,
                        name: "未开始",
                        color: "rgb(24, 144, 255)",
                        mode: "first"
                    },
                    title: "123",
                    tmpl_id: 233,
                    updated_at: "2025-05-12 22:30:07",
                    ws_id: 81
                },
                {
                    _id: "67d2e7d4bb89c8d261415f02",
                    created_at: "2025-03-13 22:12:36",
                    creator: [
                        {
                            id: 46,
                            username: "liuchenghui",
                            full_name: "刘成辉举头望明月",
                            email: "894659666@qq.com",
                            phone: "18710601494",
                            avatar: "/image/avatar/46/90ece9a3-75c9-4c01-b6d4-996558d034ac.png"
                        }
                    ],
                    custom_field_233_16bc39d30e71c5916c7303346d3400dd: [
                        {
                            id: 46,
                            username: "liuchenghui",
                            full_name: "刘成辉举头望明月",
                            email: "894659666@qq.com",
                            phone: "18710601494",
                            avatar: "/image/avatar/46/90ece9a3-75c9-4c01-b6d4-996558d034ac.png"
                        }
                    ],
                    custom_field_233_2c09e04ae6c959c8f3850e4aec8b4285: [
                        {
                            id: 46,
                            username: "liuchenghui",
                            full_name: "刘成辉举头望明月",
                            email: "894659666@qq.com",
                            phone: "18710601494",
                            avatar: "/image/avatar/46/90ece9a3-75c9-4c01-b6d4-996558d034ac.png"
                        }
                    ],
                    custom_field_233_2db3bc56dafd8d25f05e299c7e399b20:
                        "2025-03-04",
                    custom_field_233_369fcdadf43b066e77c8c3b64420c32d: 1,
                    custom_field_233_530c6b11108e5f3ce723065ce0fcb3a9: {
                        name: "a",
                        url: "a"
                    },
                    custom_field_233_5578ee327d63c0c82b8b87c6ee6770f3: "1",
                    custom_field_233_701e5537ca8c9faf08e9322a5822ac21: "1",
                    custom_field_233_a916d37252a1a89313a578f2cf11a14b: [],
                    custom_field_233_cae55a8e30a5115d953be3567a54c35a: ["a"],
                    custom_field_233_d0527f523658f09d9aa5bf437a41ea68:
                        "<p>a</p>",
                    custom_field_233_e29449c4492db8bfa3d81b834359cc41: 1,
                    custom_field_233_e36aeaade67111cab2a4bc89bb668339: [
                        {
                            _id: "67cbfb80b0252e64f970c31e",
                            created_at: "2025-03-08 16:10:40",
                            creator: [46],
                            custom_field_210_123ad281a631a7a41483c7c13ec944ac:
                                "",
                            custom_field_210_2323964f7d341b6138b0fd6e815f8276: [
                                46
                            ],
                            custom_field_210_43cde9b906a89b80c52b4bfe526f02c3: [
                                46
                            ],
                            custom_field_210_4bb04aa8870f27f1740f307624ae2166:
                                "",
                            custom_field_210_58708c44578c4df94ccf371a8ecdfbf8: {
                                name: "123",
                                url: "123"
                            },
                            custom_field_210_6e7cd969b5faaa3abf8594f7a0ae6df4:
                                [],
                            custom_field_210_7e305625b787498b3c1b712e12beb7cf: [
                                "111"
                            ],
                            custom_field_210_be1fc2eadfa7de817e4bcdbaf1d2fe0e:
                                [],
                            custom_field_210_d23fbd7ffc9f3a0bd0aa48378f241891:
                                "2025-03-11",
                            custom_field_210_e7c1a974f4edfde8a5bd37c093c100a4:
                                [],
                            handler: [46],
                            is_member: "yes",
                            issue_url:
                                "http://marktable.cn:8082/#/task/detail/81/210/67cbfb80b0252e64f970c31e",
                            status: 110,
                            title: "1111",
                            tmpl_id: 210,
                            updated_at: "2025-03-08 18:00:32",
                            ws_id: 81
                        }
                    ],
                    custom_field_233_f7c6a2d80960308c9e51374152f5d909:
                        "2025-03-13 00:00:00",
                    handler: [
                        {
                            id: 46,
                            username: "liuchenghui",
                            full_name: "刘成辉举头望明月",
                            email: "894659666@qq.com",
                            phone: "18710601494",
                            avatar: "/image/avatar/46/90ece9a3-75c9-4c01-b6d4-996558d034ac.png"
                        }
                    ],
                    progress_num: 0,
                    status: {
                        id: 188,
                        ws_id: 81,
                        tmpl_id: 233,
                        name: "未开始",
                        color: "rgb(24, 144, 255)",
                        mode: "first"
                    },
                    title: "a",
                    tmpl_id: 233,
                    updated_at: "2025-03-14 21:07:26",
                    ws_id: 81
                }
            ]
        };
    },
    watch: {},
    mounted() {},
    methods: {
        openHtmlCol(row, col) {
            this.$emit("open-html-col", row, col);
        },
        openDetail(scope, type) {
            this.$emit("open-detail", scope, type);
        }
        // multipleSelection() {},
        // refreshTable() {}
        // copyLink
        // deleteRow
        // changeColumnWidth
    }
};
</script>

<style lang="scss" scoped>
.slot-content {
    padding: 24px 14px;
    background-color: #f2f2f2;
}
</style>
<style lang="scss"></style>
