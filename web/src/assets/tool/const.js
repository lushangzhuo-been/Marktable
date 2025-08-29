export const imgHost =
    process.env.NODE_ENV === "out"
        ? window.WEBCONFIG.VUE_APP_DOWNLOAD_FILE_SERVER_PATH
        : process.env.VUE_APP_DOWNLOAD_FILE_SERVER_PATH;
export const FieldType = {
    text_com: require("@/assets/image/field_type_icon/text.svg"),
    select_com: require("@/assets/image/field_type_icon/select.svg"),
    person_com: require("@/assets/image/field_type_icon/people_single.svg"),
    person_com_single: require("@/assets/image/field_type_icon/people_single.svg"),
    person_com_multiple: require("@/assets/image/field_type_icon/people_multiple.svg"),
    time_com: require("@/assets/image/field_type_icon/time.svg"),
    date_com: require("@/assets/image/field_type_icon/date.svg"),
    multi_select_com: require("@/assets/image/field_type_icon/tag.svg"),
    textarea_com: require("@/assets/image/field_type_icon/rich_text.svg"),
    html_text_com: require("@/assets/image/field_type_icon/html.svg"),
    file_com: require("@/assets/image/field_type_icon/file.svg"),
    link_com: require("@/assets/image/field_type_icon/link.svg"),
    number_com: require("@/assets/image/field_type_icon/number.svg"),
    progress_com: require("@/assets/image/field_type_icon/progress.svg"),
    status_com: require("@/assets/image/field_type_icon/select.svg"),
    related_com: require("@/assets/image/field_type_icon/related.svg")
};

export const FieldColumnWidth = {
    title_text_com: 340,
    text_com: 240,
    select_com: 140,
    person_com_single: 180,
    person_com_multiple: 180,
    time_com: 180,
    // created_at: 180,
    // updated_at: 180,
    date_com: 160,
    multi_select_com: 140,
    textarea_com: 220,
    html_text_com: 140,
    file_com: 220,
    link_com: 140,
    number_com: 160,
    progress_com: 180,
    status_com: 120, // 状态列
    related_com: 220, // 关联流程列
    operation: 80
};
