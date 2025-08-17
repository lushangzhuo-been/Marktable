<template>
    <div id="fullsm">
        <div class="mark-table-wang-editor">
            <Toolbar
                style="border-bottom: 1px solid #ccc"
                :editor="editor"
                :defaultConfig="toolbarConfig"
                :mode="mode"
            />
            <Editor
                :style="{ height: `${height}px`, overflowY: 'hidden' }"
                v-model="html"
                :defaultConfig="editorConfig"
                :mode="mode"
                @onChange="onChange"
                @onCreated="onCreated"
                @customPaste="customPaste"
            />
        </div>
    </div>
</template>
<script>
import { Editor, Toolbar } from "@wangeditor/editor-for-vue";
import { imgHost } from "@/assets/tool/const";
import api from "@/common/api";
export default {
    components: { Editor, Toolbar },
    props: {
        content: {
            type: String,
            default: ""
        },
        height: {
            type: [Number, String],
            default: 300
        },
        rowObj: {
            type: Object,
            default: () => {}
        },
        auth: {
            type: Boolean,
            default: false
        }
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
    },
    data() {
        return {
            editor: null,
            html: "",
            toolbarConfig: {},
            editorConfig: {
                placeholder: "请输入内容...",
                MENU_CONF: {
                    //   配置上传图片
                    uploadImage: {
                        // customUpload: this.uploadHtmlImg()
                        customUpload: async (file, insertFn) => {
                            await this.uploadHtmlImg(file, file.name, insertFn);
                        }
                    }
                }
            },
            mode: "default" // 'default' or 'simple'
        };
    },
    watch: {
        auth(val) {
            if (val) {
                this.editorConfig.readOnly = true;
            } else {
                this.editorConfig.readOnly = false;
            }
        }
    },
    methods: {
        onChange() {
            this.$emit("changeData", this.html);
        },
        onCreated(editor) {
            this.editor = Object.seal(editor); // 一定要用 Object.seal() ，否则会报错
            this.toolbarConfig = {
                // toolbarKeys: [
                //     //自定义菜单选项
                //     // 菜单 key
                //     'headerSelect',
                //     // // 分割线
                //     '|',
                //     // 菜单 key
                //     'bold',
                //     'italic',
                //     // // 菜单组，包含多个菜单
                //     {
                //         key: 'group-more-style', // 必填，要以 group 开头
                //         title: '更多样式', // 必填
                //         iconSvg: '<svg>123</svg>', // 可选
                //         menuKeys: ['through', 'code', 'clearStyle'] // 下级菜单 key ，必填
                //     },
                //     'fontSize',
                //     'fontFamily',
                //     'underline',
                //     'justifyLeft',
                //     'justifyRight',
                //     'justifyCenter',
                //     'color',
                //     'bgColor',
                //     'fullScreen'
                // ],
                excludeKeys: [
                    "group-video",
                    "codeBlock",
                    // 'fullScreen',
                    // 'bold',
                    // 'underline',
                    // 'italic',
                    // 'through',
                    // 'code',
                    // 'sub',
                    // 'sup',
                    // 'clearStyle',
                    // 'color',
                    // 'bgColor',
                    // 'fontSize',
                    // 'fontFamily',
                    // 'indent',
                    // 'delIndent',
                    // 'justifyLeft',
                    // 'justifyRight',
                    // 'justifyCenter',
                    // 'justifyJustify',
                    // 'lineHeight',
                    // 'insertImage',
                    // 'deleteImage',
                    // 'editImage',
                    // 'viewImageLink',
                    // 'imageWidth30',
                    // 'imageWidth50',
                    // 'imageWidth100',
                    // 'divider',
                    // 'emotion',
                    // 'insertLink',
                    // 'editLink',
                    // 'unLink',
                    // 'viewLink',
                    // 'codeBlock',
                    // 'blockquote',
                    // 'headerSelect',
                    // 'header1',
                    // 'header2',
                    // 'header3',
                    // 'header4',
                    // 'header5',
                    // 'todo',
                    // 'redo',
                    // 'undo',
                    // 'fullScreen',
                    // 'enter',
                    // 'bulletedList',
                    // 'numberedList',
                    // 'insertTable',
                    // 'deleteTable',
                    // 'insertTableRow',
                    // 'deleteTableRow',
                    // 'insertTableCol',
                    // 'deleteTableCol',
                    // 'tableHeader',
                    // 'tableFullWidth',
                    // 'insertVideo',
                    // 'uploadVideo',
                    // 'editVideoSize',
                    // 'uploadImage',
                    // 'codeSelectLang',
                    // 'group-more-style',
                    "sub",
                    "sup"
                ]
            };
            this.editor.on("fullScreen", () => {
                const fullsm = document.getElementById("fullsm");
                fullsm.style.zIndex = 20001;
                fullsm.style.position = "relative";
                this.$emit("fullScreen");
            });
            this.editor.on("unFullScreen", () => {
                const fullsm = document.getElementById("fullsm");
                fullsm.style = "";
                this.$emit("unFullScreen");
            });
        },
        customPaste(pasteStr, b, c) {
            // // 处理粘贴的图片，传递的参数是粘贴来的数据
            // let _this = this;
            // return new Promise(function (resolve) {
            //     // 用于计数图片数量
            //     let imgNum = 0;
            //     //匹配图片
            //     var imgReg = /<img.*?(?:>|\/>)/gi;
            //     //匹配src属性
            //     // eslint-disable-next-line no-useless-escape
            //     var srcReg = /src=[\'\"]?([^\'\"]*)[\'\"]?/i;
            //     // 提取图片连接
            //     pasteStr.replace(imgReg, function (txt) {
            //         return txt.replace(srcReg, function (src) {
            //             var img_src = src.match(srcReg)[1];
            //             //正则把?x-oss-process后面的都去掉
            //             img_src = img_src.replace(/\?.*/i, '');
            //             // 查找到一个图片地址就讲图片数量加1
            //             imgNum++;
            //             // 将图片转成图片文件对象并上传得到图片地址，传递的参数是图片地址
            //             api.uploadHtmlImg(img_src).then((res) => {
            //                 /**
            //                  * 得到图片地址进行替换并将替换之后的文本返回渲染
            //                  */
            //                 // 图片地址进行替换
            //                 pasteStr = pasteStr.replace(img_src, res);
            //                 // 替换之后将图片数量减1
            //                 imgNum--;
            //                 // 只有图片数量变成0的时候才会进行返回
            //                 if (imgNum == 0) {
            //                     resolve(pasteStr);
            //                 }
            //             });
            //         });
            //     });
            // });
        },
        uploadHtmlImg(file, name, insertFn) {
            this.$nextTick(() => {
                const form = new FormData();
                form.append("tmpl_html_image", file);
                form.append("ws_id", this.curSpace.id);
                form.append("tmpl_id", this.curProgress);
                form.append("issue_id", this.rowObj._id);
                api.uploadHtmlImg(form)
                    .then((res) => {
                        if (res && res.resultCode === 200 && res.data) {
                            insertFn(`${imgHost}${res.data.substr(1)}`);
                        }
                    })
                    .catch((error) => {
                        this.$message("上传失败,请重新上传");
                    });
            });
        }
    },
    created() {
        this.html = this.content;
    },
    mounted() {},
    beforeDestroy() {
        const editor = this.editor;
        if (editor == null) return;
        editor.destroy(); // 组件销毁时，及时销毁编辑器
    }
};
</script>
<style src="@wangeditor/editor/dist/css/style.css"></style>
<style lang="scss">
.mark-table-wang-editor {
    border: 1px solid #ccc;
    .w-e-text-container .w-e-scroll p {
        margin: 21px 0;
    }
}
</style>
