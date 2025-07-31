import { Message } from 'element-ui'
// 私有属性，只在当前文件可用
const showMessage = Symbol('showMessage')
export default class domMessage {
    success(options, single = true) {
        this[showMessage]('success', options, single)
    }
    warning(options, single = true) {
        this[showMessage]('warning', options, single)
    }
    info(options, single = true) {
        this[showMessage]('info', options, single)
    }
    error(options, single = true) {
        this[showMessage]('error', options, single)
    }
    [showMessage](type, options, single) {
        if (single) {
            // 判断当前页是否有el-message标签，如果没有则执行弹窗操作
            if (document.getElementsByClassName('el-message').length === 0) {
                Message[type](options)
            }
        } else {
            Message[type](options)
        }
    }
}