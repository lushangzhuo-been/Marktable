export default {
    defaultProgressInfo: [
        {
            label: "名称",
            value: ""
        },
        {
            label: "描述",
            value: ""
        },
        // {
        //     label: "类型",
        //     value: ""
        // },
        {
            label: "创建人",
            value: ""
        },
        {
            label: "创建时间",
            value: ""
        }
    ],
    defaultProgressInfoForm: (res) => {
        return [
            {
                label: "名称",
                value: res.name
            },
            {
                label: "描述",
                value: res.desc
            },
            // {
            //     label: "类型",
            //     value: res.mode
            // },
            {
                label: "创建人",
                value: res.created_full_name
            },
            {
                label: "创建时间",
                value: res.created_at
            }
        ];
    },
    defaultfolderInfo: [
        {
            label: "名称",
            value: ""
        },
        {
            label: "描述",
            value: ""
        },
        {
            label: "创建人",
            value: ""
        },
        {
            label: "创建时间",
            value: ""
        }
    ],
    defaultfolderInfoForm: (res) => {
        return [
            {
                label: "名称",
                value: res.name
            },
            {
                label: "描述",
                value: res.desc
            },
            {
                label: "创建人",
                value: res.created_full_name
            },
            {
                label: "创建时间",
                value: res.created_at
            }
        ];
    },
    progressBasicUnshiftCol: [
        {
            prop: "subTask",
            label: "",
            align: "center",
            mode: "subTask",
            fixed: "left",
            columnName: "subTask"
        }
    ],
    progressBasicCol: [
        // {
        //     mode: "selection",
        // },
        {
            prop: "operation",
            label: "操作",
            fixed: "right",
            align: "center",
            mode: "operation",
            columnName: "operation"
        }
    ]
};
