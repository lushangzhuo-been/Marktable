export default {
    defaultBasicInfoList: [{
        label: '头像',
        value: ''
    }, {
        label: '昵称',
        value: ''
    }, {
        label: '邮箱',
        value: ''
    }, {
        label: '电话',
        value: ''
    }],
    basicInfoList: res => {
        return [{
            label: '头像',
            value: res.avatar
        }, {
            label: '昵称',
            value: res.username
        }, {
            label: '邮箱',
            value: res.email
        }, {
            label: '电话',
            value: res.phone
        }]
    }
}