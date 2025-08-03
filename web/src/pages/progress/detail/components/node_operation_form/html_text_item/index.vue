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
            isEditing: false,
            popoverWidth: 220,
            text: '',
            curFormData: {}
        };
    },
    watch: {
        formData: {
            handler(obj) {
                this.curFormData = obj;
            }
        }
    },
    methods: {
        openHtmlDialog(row) {
            this.$refs['htmlColDialog'].openDialog(row, this.formItem);
        },
        editFormItem(content) {
            this.curFormData[this.formItem.field_key] = content;
            this.$emit(
                'edit-form-item',
                content,
                this.formItem.field_key,
                this.formItem.mode
            );
        }
    }
};
</script>

<style lang="scss" scoped>
@import '../style.scss';
.html-icon {
    img {
        vertical-align: middle;
        cursor: pointer;
    }
}
</style>
