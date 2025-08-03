<template>
    <div>
        <div
            class="html-icon column-block form-width"
            :class="{
                'editing-effect': formItem.can_modify === 'yes',
                'col-required': formItem.required === 'yes'
            }"
        >
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
import HtmlTextDialog from "./html_col_dialog.vue";
import _ from "lodash";
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
            text: ""
        };
    },
    mounted() {},
    methods: {
        openHtmlDialog(row) {
            this.$refs["htmlColDialog"].openDialog(row, this.formItem);
        },
        editFormItem(content) {
            this.$emit(
                "edit-form-item",
                content,
                this.formItem.field_key,
                this.formItem.mode
            );
        }
    }
};
</script>

<style lang="scss" scoped>
@import "../../style.scss";
.html-icon {
    img {
        position: relative;
        top: -2px;
        margin-left: 10px;
        vertical-align: middle;
        cursor: pointer;
    }
}
.column-block {
    &:hover {
        background-color: #e9f0f8;
    }
}
div.column-block.editing-effect.col-required:hover:before {
    transition-duration: 0.3s;
    content: "";
    display: block;
    position: absolute;
    right: 0;
    top: 0;
    width: 12px;
    height: 12px;
    background: linear-gradient(42deg, #fd5050 30%, rgba(0, 0, 0, 0) 40%);
}
</style>
