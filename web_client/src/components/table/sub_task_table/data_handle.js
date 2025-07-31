export default {
    subTaskCol: [
        {
            field_key: "title",
            name: "标题",
            desc: "",
            mode: "main_text_com",
            enum_values: null,
            expr: "",
            can_modify: "yes",
            required: "yes",
            level: "level2",
            read_only_rule: "",
            related_tmpl_id: 0
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
            related_tmpl_id: 0
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
            related_tmpl_id: 0
        },
        {
            field_key: "creator",
            name: "创建人",
            desc: "",
            mode: "person_com",
            enum_values: null,
            expr: "single",
            can_modify: "no",
            required: "no",
            level: "level1",
            read_only_rule: "",
            related_tmpl_id: 0
        },
        {
            field_key: "created_at",
            name: "创建时间",
            desc: "",
            mode: "time_com",
            enum_values: null,
            expr: "YmdHis",
            can_modify: "no",
            required: "no",
            level: "level1",
            read_only_rule: "",
            related_tmpl_id: 0
        },
        {
            field_key: "updated_at",
            name: "最后更新时间",
            desc: "",
            mode: "time_com",
            enum_values: null,
            expr: "YmdHis",
            can_modify: "no",
            required: "no",
            level: "level1",
            read_only_rule: "",
            related_tmpl_id: 0
        },
        {
            field_key: "operation",
            name: "操作",
            desc: "",
            mode: "operation",
            enum_values: null,
            expr: "",
            can_modify: "",
            required: "",
            level: "",
            read_only_rule: "",
            related_tmpl_id: 0
        }
    ]
};
