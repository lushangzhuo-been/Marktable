<template>
    <div v-if="curFile.real_relative_path">
        <!-- docx预览 -->
        <vue-office-docx
            v-if="curFile.fileType === 'docx' || curFile.fileType === 'doc'"
            style="width: 100%"
            :src="curFile.real_relative_path"
            @rendered="rendered"
        />
        <!-- <div
            v-if="curFile.fileType === 'docx' || curFile.fileType === 'doc'"
            style="width: 100%; height: 100%"
            id="bodyContainer"
        ></div> -->
        <!-- xlsx预览 -->
        <!-- <div
            v-else-if="
                curFile.fileType === 'xlsx' || curFile.fileType === 'xls'
            "
            class="table_box"
            v-html="curFile.real_relative_path"
        ></div> -->
        <vue-office-excel
            v-else-if="
                curFile.fileType === 'xlsx' || curFile.fileType === 'xls'
            "
            style="width: 100%"
            :src="curFile.real_relative_path"
            :options="{
                xls: false //预览xlsx文件设为false；预览xls文件设为true
                // minColLength: 0, // excel最少渲染多少列，如果想实现xlsx文件内容有几列，就渲染几列，可以将此值设置为0.
                // minRowLength: 0, // excel最少渲染多少行，如果想实现根据xlsx实际函数渲染，可以将此值设置为0.
                // widthOffset: 10, //如果渲染出来的结果感觉单元格宽度不够，可以在默认渲染的列表宽度上再加 Npx宽
                // heightOffset: 10, //在默认渲染的列表高度上再加 Npx高
                // beforeTransformData: (workbookData) => {
                //     return workbookData;
                // }, //底层通过exceljs获取excel文件内容，通过该钩子函数，可以对获取的excel文件内容进行修改，比如某个单元格的数据显示不正确，可以在此自行修改每个单元格的value值。
                // transformData: (workbookData) => {
                //     return workbookData;
                // } //将获取到的excel数据进行处理之后且渲染到页面之前，可通过transformData对即将渲染的数据及样式进行修改，此时每个单元格的text值就是即将渲染到页面上的内容
            }"
            @rendered="rendered"
            @error="error"
        />
        <!-- xls预览 -->
        <!-- <vue-office-excel
            v-else-if="curFile.fileType === 'xls'"
            style="height: 100%; width: 100%"
            :src="curFile.real_relative_path"
            :options="{
                xls: true, //预览xlsx文件设为false；预览xls文件设为true
                // minColLength: 0, // excel最少渲染多少列，如果想实现xlsx文件内容有几列，就渲染几列，可以将此值设置为0.
                // minRowLength: 0, // excel最少渲染多少行，如果想实现根据xlsx实际函数渲染，可以将此值设置为0.
                // widthOffset: 10, //如果渲染出来的结果感觉单元格宽度不够，可以在默认渲染的列表宽度上再加 Npx宽
                // heightOffset: 10, //在默认渲染的列表高度上再加 Npx高
                beforeTransformData: (workbookData) => {
                    return workbookData;
                }, //底层通过exceljs获取excel文件内容，通过该钩子函数，可以对获取的excel文件内容进行修改，比如某个单元格的数据显示不正确，可以在此自行修改每个单元格的value值。
                transformData: (workbookData) => {
                    return workbookData;
                } //将获取到的excel数据进行处理之后且渲染到页面之前，可通过transformData对即将渲染的数据及样式进行修改，此时每个单元格的text值就是即将渲染到页面上的内容
            }"
            @rendered="rendered"
            @error="error"
        /> -->
        <!-- pdf预览 -->
        <!-- <vue-office-pdf
            v-else-if="getPdfType"
            style="width: 100%"
            :class="curFile.fileType === 'pdf' ? 'pdf' : 'ppt'"
            :src="curFile.real_relative_path"
            @rendered="rendered"
            @error="error"
        /> -->
        <iframe
            v-else-if="getPdfType"
            :src="`${curFile.real_relative_path}#toolbar=0`"
            frameborder="0"
            style="width: 100%; height: calc(100% - 20px)"
        ></iframe>
        <!-- 图片预览 -->
        <div class="preview-container" v-else-if="getImgType">
            <img :src="curFile.real_relative_path" alt="" />
        </div>
        <!-- txt， json， xml， csv预览 -->
        <template v-else-if="getTextType">
            <JsonViewer
                v-if="curFile.fileType === 'json' && curFile.real_relative_path"
                :data="JSON.parse(curFile.real_relative_path)"
            />
            <pre v-else-if="curFile.fileType === 'xml'" class="text">
                <code class="language-xml">{{curFile.real_relative_path}}</code>
            </pre>
            <!-- csv和txt -->
            <template v-else>
                <span
                    v-show="curFile.real_relative_path"
                    style="white-space: pre-wrap; word-wrap: break-word"
                    v-html="curFile.real_relative_path"
                >
                </span>
            </template>
        </template>
        <!-- pptx -->
        <!-- <div v-show="curFile.fileType === 'pptx'" ref="pptx" id="pptx"></div> -->
        <!-- <iframe
            v-if="curFile.fileType === 'pptx'"
            src="/PPTXjs-1.21.1/index.html?file=http://47.109.66.52:8081//file/tmpl/148/dd1ea530-cd82-4729-8024-4f5f85c92c76.pptx"
            frameborder="0"
            width="100%"
            height="1000"
        ></iframe> -->
    </div>
    <no-data
        v-else
        :isTextShow="false"
        :loading="loading"
        :imgShow="imgShow"
        :text="showMessage"
        :loadingMessage="message"
    ></no-data>
</template>
<script>
import _ from "lodash";
import NoData from "@/pages/common/no_data_preview.vue";
// docx插件
import VueOfficeDocx from "@vue-office/docx";
import "@vue-office/docx/lib/index.css";
// xlsx插件
import VueOfficeExcel from "@vue-office/excel";
import "@vue-office/excel/lib/index.css";
import XLSX from "xlsx";
// pdf插件
// import VueOfficePdf from "@vue-office/pdf";
// json格式预览插件
import JsonViewer from "vue-json-views";
// pptx预览插件，有问题
// import { init } from "pptx-preview"; // 预览ppt，有问题

import { readBuffer, readText } from "./methods.js"; // 预览txt，json等

// xml高亮等处理 插件
import hljs from "highlight.js";
import "highlight.js/styles/a11y-light.css";
import xml from "highlight.js/lib/languages/xml";
hljs.registerLanguage("xml", xml);

// doc文件
import { renderAsync } from "docx-preview";
import api from "@/common/api/module/progress";
import { imgHost } from "@/assets/tool/const";
export default {
    props: {
        curPreviewFile: {
            type: Object,
            default: () => {}
        },
        detailId: {
            type: String,
            default: ""
        }
    },
    components: {
        JsonViewer,
        VueOfficeDocx,
        VueOfficeExcel,
        // VueOfficePdf,
        NoData
    },
    data() {
        return {
            showMessage: "暂无数据",
            message: "加载中",
            imgShow: false,
            loading: true,
            curFile: {},
            docxOptions: {
                className: "doc", // string：默认和文档样式类的类名/前缀
                inWrapper: true, // boolean：启用围绕文档内容的包装器渲染
                ignoreWidth: false, // boolean：禁用页面的渲染宽度
                ignoreHeight: false, // boolean：禁止渲染页面高度
                ignoreFonts: false, // boolean：禁用字体渲染
                breakPages: true, // boolean：在分页符上启用分页
                ignoreLastRenderedPageBreak: true, // boolean：在lastRenderedPageBreak 元素上禁用分页
                experimental: false, // boolean：启用实验功能（制表符停止计算）
                trimXmlDeclaration: true, // boolean：如果为true，解析前会从​​ xmlTemplate 文档中移除 xmlTemplate 声明
                // useBase64URL: false, // boolean：如果为true，图片、字体等会转为base 64 URL，否则使用URL.createObjectURL
                // useMathMLPolyfill: false, // boolean：包括用于 chrome、edge 等的 MathML polyfill。
                // showChanges: false, // boolean：启用文档更改的实验性渲染（插入/删除）
                debug: false // boolean：启用额外的日志记录
            }
        };
    },
    computed: {
        getImgType() {
            return (
                this.curFile.fileType === "png" ||
                this.curFile.fileType === "jpg" ||
                this.curFile.fileType === "jpeg" ||
                this.curFile.fileType === "gif"
            );
        },
        getTextType() {
            return (
                this.curFile.fileType === "txt" ||
                this.curFile.fileType === "json" ||
                this.curFile.fileType === "xml" ||
                this.curFile.fileType === "csv"
            );
        },
        getPdfType() {
            return (
                this.curFile.fileType === "pdf" ||
                this.curFile.fileType === "ppt" ||
                this.curFile.fileType === "pptx"
            );
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
    },
    watch: {
        curPreviewFile: {
            handler(obj) {
                this.showMessage = "暂无数据";
                this.message = "加载中";
                this.imgShow = false;
                this.loading = true;
                if (obj && Object.keys(obj).length) {
                    this.curFile = _.cloneDeep(obj);
                    this.curFile.fileType = this.splitAndGetLast(
                        this.curFile.file_name,
                        "."
                    );
                    if (this.getImgType) {
                        this.$set(
                            this.curFile,
                            "real_relative_path",
                            `${imgHost}${this.curFile.relative_path}`
                        );
                    } else {
                        this.getFile();
                    }
                } else {
                    this.curFile = {};
                }
            },
            deep: true
        }
    },
    methods: {
        getFile() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                id: this.curPreviewFile.id,
                download_file_type: "transformed_original_name"
            };
            api.getFileStatus(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    let status = res.data.transformed_status;
                    if (status === "succeed" || status === "no-transform") {
                        // 成功 或者不需要转义直接预览
                        this.preview();
                    } else if (status === "transforming") {
                        this.message = "文件正在转化中，请稍等...";
                        setTimeout(() => {
                            this.getFile();
                        }, 500);
                    } else {
                        this.showMessage = "加载失败，请重新上传";
                        this.loading = false;
                        this.imgShow = true;
                    }
                } else {
                    this.loading = false;
                    this.imgShow = true;
                }
            });
        },
        rendered() {},
        error(err) {},
        splitAndGetLast(str, separator) {
            // 使用split方法按separator分割字符串，然后取最后一个元素
            return str.split(separator).slice(-1)[0].toLowerCase();
        },
        async preview() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                id: this.curPreviewFile.id,
                download_file_type: "transformed_original_name"
            };
            api.previewFile(params).then((res) => {
                let blobFile = new Blob([res]);
                if (this.getTextType) {
                    let mapping = {
                        // xml: "text/xml",
                        xml: "text/pain",
                        txt: "text/pain",
                        csv: "text/csv",
                        json: "application/json"
                    };
                    readBuffer(blobFile).then((buffer) => {
                        readText(
                            buffer,
                            mapping[this.curFile.fileType] || "text/pain"
                        ).then((res) => {
                            this.$set(this.curFile, "real_relative_path", res);
                            this.$nextTick(() => {
                                // 注意数据显示后再去渲染高亮
                                document
                                    .querySelectorAll("pre code")
                                    .forEach((el) => {
                                        hljs.highlightElement(el);
                                    });
                            });
                        });
                    });
                } else if (this.getPdfType) {
                    let blob = new Blob([res], { type: "application/pdf" });
                    let fileUrl = window.URL.createObjectURL(blob);
                    this.$set(this.curFile, "real_relative_path", fileUrl);
                } else {
                    let fileReader = new FileReader();
                    fileReader.readAsArrayBuffer(res);
                    fileReader.onload = () => {
                        // if (
                        //     this.curFile.fileType === "xls" ||
                        //     this.curFile.fileType === "xlsx"
                        // ) {
                        //     // 通过event.target.result获取转换后的ArrayBuffer结果
                        //     const arrayBuffer = res.arrayBuffer();
                        //     // 处理arrayBuffer...
                        //     const data = new Uint8Array(arrayBuffer);
                        //     const workbook = XLSX.read(data, {
                        //         type: "array"
                        //     });
                        //     // 删除掉没有数据的sheet
                        //     Object.values(workbook.Sheets).forEach(
                        //         (sheet, index) => {
                        //             if (
                        //                 Object.keys(sheet).indexOf(
                        //                     "!ref"
                        //                 ) === -1
                        //             ) {
                        //                 delete workbook.Sheets[
                        //                     workbook.SheetNames[index]
                        //                 ];
                        //             }
                        //         }
                        //     );
                        //     this.tableToHtml(workbook);
                        // } else {
                        this.$set(
                            this.curFile,
                            "real_relative_path",
                            fileReader.result
                        );
                        // }
                    };
                    // }
                }
                if (this.curFile.real_relative_path) {
                    this.loading = false;
                    setTimeout(() => {
                        this.imgShow = true;
                    }, 50);
                } else {
                    this.loading = true;
                    this.imgShow = false;
                }
            });
            // }
        },
        tableToHtml(workbook) {
            const sheetList = workbook.SheetNames.filter(
                (v) => v.indexOf("数据源") === -1
            );
            sheetList.forEach((sheet) => {
                const worksheet = workbook.Sheets[sheet];
                if (worksheet) {
                    this.tableContext = XLSX.utils.sheet_to_html(worksheet);
                    this.$set(
                        this.curFile,
                        "real_relative_path",
                        XLSX.utils.sheet_to_html(worksheet)
                    );
                }
            });
        }
    }
};
</script>
<style lang="scss" scoped>
.table_box {
    width: 100%;
    max-height: 500px;
    background-color: #ffffff;
    border: 1px solid #cccccc;
    border-bottom: none;
    overflow: auto;
}

.table_box ::v-deep table {
    width: 100%;
    text-align: center;
}

.table_box ::v-deep td {
    padding: 5px 20px;
    border-right: 1px solid #cccccc;
    border-bottom: 1px solid #cccccc;
}

// 图片预览
.preview-container {
    height: 100%;
    text-align: center;
    // overflow-y: auto;
}
.preview-container img {
    // position: relative;
    // top: 50%;
    // transform: translateY(-50%);
    max-width: 90%;
    // max-height: 90%;
    // border: 1px solid #ddd;
    // border-radius: 4px;
    // box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
::v-deep .vue-office-pdf {
    // overflow-y: clip !important;
    // .vue-office-pdf-wrapper {
    //     width: auto;
    //     min-width: 756px;
    // }
}
// 文本样式优化
pre.text ::v-deep code.hljs {
    overflow-x: inherit;
}
// pdf样式优化
::v-deep .vue-office-pdf {
    overflow-y: clip !important;
    &.pdf {
        .vue-office-pdf-wrapper {
            min-width: 902px !important;
            width: none;
            canvas {
                width: 862px !important;
            }
        }
    }
    &.ppt {
        .vue-office-pdf-wrapper {
            min-width: 1163px !important;
            width: none;
            canvas {
                width: 1123px !important;
            }
        }
    }
}
// doc文件样式
::v-deep .vue-office-docx {
    height: auto;
    overflow-y: clip !important;
    .docx-wrapper {
        min-width: 840px;
        > section.docx {
            width: 800px !important;
        }
    }
}
</style>
<style lang="scss">
BODY {
    scrollbar-face-color: rgb(10, 236, 209); //滚动条凸出部分的颜色
    scrollbar-highlight-color: rgb(23, 255, 155); //滚动条空白部分的颜色
    scrollbar-shadow-color: rgb(255, 116, 23); //立体滚动条阴影的颜色
    scrollbar-3dlight-color: rgb(66, 93, 127); //滚动条亮边的颜色
    scrollbar-arrow-color: rgb(93, 232, 255); //上下按钮上三角箭头的颜色
    scrollbar-track-color: rgb(255, 70, 130); //滚动条的背景颜色
    scrollbar-darkshadow-color: rgb(10, 0, 209); //滚动条强阴影的颜
    scrollbar-base-color: rgb(66, 93, 128) / 滚动条的基本颜色;
}
</style>
