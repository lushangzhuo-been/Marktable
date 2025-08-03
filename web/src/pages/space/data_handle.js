export default {
    defaultSpaceDetail: [{
        label: '名称',
        value: ''
    }, {
        label: '描述',
        value: ''
    }, {
        label: '创建人',
        value: ''
    }, {
        label: '创建时间',
        value: ''
    }, {
        lael: '管理员',
        value: ''
    }],
    spaceDetailForm: (res) => {
        return [{
            label: '名称',
            value: res.name
        }, {
            label: '描述',
            value: res.desc
        }, {
            label: '创建人',
            value: res.creator_name
        }, {
            label: '创建时间',
            value: res.created_at
        }, {
            label: '管理员',
            value: res.admin
        }]
    }
}