export default {
    ruleListCol: [
        {
            prop: "name",
            name: "规则名称",
            minWidth: 160
        },
        {
            prop: "description",
            name: "描述",
            minWidth: 220
        },
        {
            prop: "creator",
            name: "创建人",
            type: "slot",
            minWidth: 120
        },
        {
            prop: "created_at",
            name: "创建时间",
            minWidth: 160
        },
        {
            prop: "is_enable",
            name: "是否启用",
            type: "slot",
            minWidth: 120
        },
        {
            prop: "operation",
            name: "操作",
            type: "slot",
            minWidth: 220
        }
    ],
    logListCol: [
        {
            prop: "rule_name",
            name: "规则名称",
            type: "interactive",
            minWidth: 160
        },
        {
            prop: "rule_type",
            name: "触发类型",
            type: "slot",
            minWidth: 160
        },
        {
            prop: "data_name",
            name: "任务名称",
            minWidth: 160
        },
        {
            prop: "execute_at",
            name: "执行时间",
            minWidth: 160
        }
        // {
        //     prop: "status",
        //     name: "执行结果",
        //     type: "slot",
        //     minWidth: 160
        // }
        // {
        //     prop: "name",
        //     name: "流程",
        //     minWidth: 160
        // // },
        // {
        //     prop: "name",
        //     name: "任务",
        //     minWidth: 260
        // },
        // {
        //     prop: "enable_status",
        //     name: "是否启用",
        //     type: "slot",
        //     width: 120
        // }
    ]
};
