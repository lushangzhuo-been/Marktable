<template>
    <div>
        <div class="html-icon">
            <img
                :src="getHtmlImg(formData)"
                :width="20"
                :height="20"
                @click="openHtmlDialog(formData)"
            />
        </div>
        <!-- html富文本编辑器 -->
        <html-text-dialog
            ref="htmlColDialog"
            @edit-form-item="editFormItem"
        ></html-text-dialog>
        <div v-show="validateFailed" class="validate-desc">此项为必填</div>
    </div>
</template>

<script>
import HtmlTextDialog from './html_col_dialog.vue';
import _ from 'lodash';
export default {
    components: {
        HtmlTextDialog
    },
    props: {
        formItem: {
            type: Object,
            default: () => {}
        },
        formData: {
            type: Object,
            default: () => {}
        },
        validateOrder: {
            handler() {
                this.doValidate();
            }
        }
    },
    computed: {
        getHtmlImg() {
            return (row) => {
                if (
                    row &&
                    this.formItem &&
                    Object.keys(this.formItem).length &&
                    row[this.formItem.field_key]
                ) {
                    return require(`@/assets/image/progress_column/col_html_content.png`);
                }
                return require(`@/assets/image/progress_column/col_html_empty.png`);
            };
        }
    },
    data() {
        return {
            validateFailed: false,
            isEditing: false,
            popoverWidth: 220,
            text: '',
            curFormItem: {},
            curFormData: {}
        };
    },
    watch: {
        validateOrder: {
            handler() {
                this.doValidate();
            }
        }
    },
    mounted() {},
    methods: {
        doValidate() {
            this.curFormItem = this.formItem;
            if (
                this.curFormItem.required === 'yes' &&
                !this.formData[this.formItem.field_key]
            ) {
                this.validateFailed = true;
                this.curFormItem.validateFailed = true;
            } else {
                this.validateFailed = false;
                this.curFormItem.validateFailed = false;
            }
        },
        openHtmlDialog(row) {
            this.$refs['htmlColDialog'].openDialog(row, this.formItem);
        },
        editFormItem(content) {
            this.$emit(
                'on-change',
                content,
                this.formItem.field_key,
                this.formData._id,
                this.formItem.mode
            );
            this.doValidate();
        }
    }
};
</script>

<style lang="scss" scoped>
@import '../../style.scss';
.html-icon {
    img {
        vertical-align: middle;
        cursor: pointer;
    }
}
.validate-desc {
    color: #f56c6c;
    font-size: 12px;
    height: 12px;
    line-height: 12px;
    position: absolute;
    bottom: -16px;
}
</style>
