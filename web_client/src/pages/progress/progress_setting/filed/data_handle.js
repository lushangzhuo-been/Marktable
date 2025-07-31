import { PRIMARY_COLOR_LIST } from '@/assets/tool/color_list'
export default {
    timePersonalOptions: [{
        label: '1',
        name: '单人',
    },
    {
        label: '2',
        name: '多人',
    },],
    timeRadioOptions: [
        {
            label: '1',
            name: '年-月-日',
        },
        {
            label: '2',
            name: '年-月-日-时-分-秒',
        },
    ],
    numberUnitList: [
        {
            label: '1',
            name: '自定义',
        },
        {
            label: '2',
            name: '无',
        },
        {
            label: '3',
            name: '￥',
        },
        {
            label: '4',
            name: '$',
        },
        {
            label: '5',
            name: '%',
        },
    ],
    typeDrapDownList: [
        {
            item_no: 1,
            label: '选项1',
            value: '',
            color: PRIMARY_COLOR_LIST.PRIMARY_COLOR_RGB,
        },
        {
            item_no: 2,
            label: '选项2',
            value: '',
            color: PRIMARY_COLOR_LIST.PRIMARY_COLOR_RGB,
        },
    ],
    typeLabelList: [
        {
            item_no: 1,
            label: '选项1',
            value: '',
            color: PRIMARY_COLOR_LIST.PRIMARY_COLOR_RGB,
        },
        {
            item_no: 2,
            label: '选项2',
            value: '',
            color: PRIMARY_COLOR_LIST.PRIMARY_COLOR_RGB,
        },
    ],
    typeProgressList: [
        {
            status: '',
            progress: '',
        },
        {
            status: '',
            progress: '',
        },
        {
            status: '',
            progress: '',
        },
    ],
    typeProgressListOptions: [
        {
            label: '待开始',
            value: '待开始',
        },
        {
            label: '进行中',
            value: '进行中',
        },
        {
            label: '已完成',
            value: '已完成',
        },
    ],
    rule: {
        type: [
            {
                required: true,
                message: '请选择类型',
                trigger: 'blur',
            },
        ],
        name: [
            {
                required: true,
                message: '请输入名称',
                trigger: 'blur',
            },
        ],

    },
    typeIconMapping: {
        "text_com": require(`@/assets/image/field_type_icon/text.svg`),
        "textarea_com": require(`@/assets/image/field_type_icon/rich_text.svg`),
        "html_text_com": require(`@/assets/image/field_type_icon/html.svg`),
        "select_com": require(`@/assets/image/field_type_icon/select.svg`),
        'multi_select_com': require(`@/assets/image/field_type_icon/tag.svg`),
        "status_com": require(`@/assets/image/field_type_icon/select.svg`),
        "progress_com": require(`@/assets/image/field_type_icon/progress.svg`),
        "file_com": require(`@/assets/image/field_type_icon/file.svg`),
        "number_com": require(`@/assets/image/field_type_icon/number.svg`),
        "time_com": require(`@/assets/image/field_type_icon/time.svg`),
        "date_com": require(`@/assets/image/field_type_icon/date.svg`),
        "person_com": require(`@/assets/image/field_type_icon/people_multiple.svg`),
        "link_com": require(`@/assets/image/field_type_icon/link.svg`),
        "blank_com_creator": require(`@/assets/image/field_type_icon/people_single.svg`),
        "blank_com_status": require(`@/assets/image/field_type_icon/select.svg`),
        "related_com": require(`@/assets/image/field_type_icon/related.svg`),
    },
    filedOptions: [
        {
            label: '文本框',
            value: '文本框',
            icon: require(`@/assets/image/field_type_icon/text.svg`),
        },
        {
            label: '富文本框',
            value: '富文本框',
            icon: require(`@/assets/image/field_type_icon/rich_text.svg`),
        },
        {
            label: 'HTML编辑器',
            value: 'HTML编辑器',
            icon: require(`@/assets/image/field_type_icon/html.svg`),
        },
        {
            label: '下拉',
            value: '下拉',
            icon: require(`@/assets/image/field_type_icon/select.svg`),
        },
        {
            label: '标签',
            value: '标签',
            icon: require(`@/assets/image/field_type_icon/tag.svg`),
        },
        {
            label: '进展',
            value: '进展',
            icon: require(`@/assets/image/field_type_icon/progress.svg`),
        },
        {
            label: '文件',
            value: '文件',
            icon: require(`@/assets/image/field_type_icon/file.svg`),
        },
        {
            label: '数字',
            value: '数字',
            icon: require(`@/assets/image/field_type_icon/number.svg`),
        },
        {
            label: '时间',
            value: '时间',
            icon: require(`@/assets/image/field_type_icon/time.svg`),
        },
        {
            label: '日期',
            value: '日期',
            icon: require(`@/assets/image/field_type_icon/date.svg`),
        },
        {
            label: '人员',
            value: '人员',
            icon: require(`@/assets/image/field_type_icon/people_multiple.svg`),
        },
        {
            label: 'Link',
            value: 'Link',
            icon: require(`@/assets/image/field_type_icon/link.svg`),
        },
        {
            label: '关联字段',
            value: '关联字段',
            icon: require(`@/assets/image/field_type_icon/related.svg`),
        },
    ],
}