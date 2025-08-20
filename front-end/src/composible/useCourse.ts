import Api from "@/api"
import {destroy} from "@/api/courseController";
import {VxeGridProps} from "vxe-table";
import {request} from "@/utils/request";
import XEUtils from "xe-utils";

class RowVO {
}

export default () => {
    // 方法
    let teachList = ref<Array<{ value: any, label: string }>>([])
    let typeList = ref<Array<{ value: any, label: string }>>([])
    let tmpList = ref<Array<{ value: any, label: string }>>([])
    let venueList = ref<Array<{ value: any, label: string }>>([])
    const fillOptions = async () => {
        teachList.value = await Api.coachController.option()
        typeList.value = await Api.courseTypeController.option()
        tmpList.value = await Api.couponTemplateController.option()
        venueList.value = await Api.venueController.option()

        let row1 = gridOptions.columns.find(item => item.field === "coach_id")
        if (row1) {
            row1.editRender.options = teachList.value
        }

        let row2 = gridOptions.columns.find(item => item.field === "course_type_id")
        if (row2) {
            row2.editRender.options = typeList.value
        }

        let row3 = gridOptions.columns.find(item => item.field === "coupon_template_id")
        if (row3) {
            row3.editRender.options = tmpList.value
        }

        let row4 = gridOptions.columns.find(item => item.field === "venue_id")
        if (row4) {
            row4.editRender.options = venueList.value
        }
    }

    const changeCoach = async (venue_id) => {
        teachList.value = await Api.coachController.option({venue_id: venue_id})
        let row = gridOptions.columns.find(item => item.field === "coach_id")
        if (row) {
            row.editRender.options = teachList.value
        }
    }
    // 表单配置
    const serveApiUrl = import.meta.env.VITE_API_URL;
    const gridOptions = reactive<VxeGridProps<RowVO>>({
        border: "full",
        size: "small",
        cellConfig: {
            height: 100,
        },
        showHeaderOverflow: true,
        showOverflow: true,
        keepSource: true,
        autoResize: true,

        expandConfig: {
            trigger: "row",
            showIcon: true,
            expandRowKeys: [1, 2],
            iconOpen: "vxe-icon--caret-bottom",
            iconClose: "vxe-icon--caret-top",
        },
        formConfig: {
            titleWidth: 120,
            titleAlign: "right",
            items: [
                {
                    field: "title",
                    title: "课程标题",
                    span: 6,
                    itemRender: {
                        name: "$input",
                        props: {placeholder: "请输入角色名称"},
                    },
                },
                {
                    span: 8,
                    align: "left",
                    collapseNode: true,
                    itemRender: {
                        name: "$buttons",
                        children: [
                            {
                                props: {
                                    type: "submit",
                                    content: "搜索",
                                    status: "primary",
                                },
                            },
                            {props: {type: "reset", content: "重置"}},
                        ],
                    },
                },
            ],
        },
        stripe: true,
        id: "full_edit_9999",
        rowConfig: {
            keyField: "#",
            isHover: true,
            useKey: true,
        },
        columnConfig: {
            resizable: true,
        },
        customConfig: {
            storage: true,
            checkMethod({column}) {
                if (["name", "no"].includes(column.field)) {
                    return false;
                }
                return true;
            },
        },
        printConfig: {
            columns: [{field: "name"}],
        },
        sortConfig: {
            trigger: "cell",
            remote: true,
        },
        filterConfig: {
            remote: true,
        },
        pagerConfig: {
            enabled: true,
            pageSize: 10,
            pageSizes: [5, 10, 15, 20, 50, 100, 200, 500, 1000],
        },

        toolbarConfig: {
            buttons: [
                {code: "insert_edit", name: "新增", status: "primary"},
                {code: "delete", name: "删除", status: "danger"},
                {code: "save", name: "保存", status: "success"}
            ],
            refresh: true, // 显示刷新按钮
            import: false, // 显示导入按钮
            export: true, // 显示导出按钮
            print: true, // 显示打印按钮
            zoom: true, // 显示全屏按钮
            custom: true, // 显示自定义列按钮
        },
        proxyConfig: {
            seq: true, // 启用动态序号代理，每一页的序号会根据当前页数变化
            sort: true, // 启用排序代理，当点击排序时会自动触发 query 行为
            filter: true, // 启用筛选代理，当点击筛选时会自动触发 query 行为
            form: true, // 启用表单代理，当点击表单提交按钮时会自动触发 reload 行为
            props: {
                // 对应响应结果 Promise<{ result: [], page: { total: 100 } }>
                result: "data", // 配置响应结果列表字段
                total: "total", // 配置响应结果总页数字段
            },
            showResponseMsg: false,
            // 只接收Promise，具体实现自由发挥
            ajax: {
                // 当点击工具栏查询按钮或者手动提交指令 query或reload 时会被触发
                query: async ({page, sorts, filters, form}) => {
                    await fillOptions()
                    //通过当前对象为选框添加选项
                    const queryParams: any = Object.assign({}, form);
                    // 处理排序条件
                    const firstSort = sorts[0];
                    if (firstSort) {
                        queryParams.sort = firstSort.field;
                        queryParams.order = firstSort.order;
                    }
                    // 处理筛选条件
                    filters.forEach(({field, values}) => {
                        queryParams[field] = values.join(",");
                    });
                    return request({
                        url: `/api/admin/course?pageSize=${page.pageSize}&currentPage=${page.currentPage
                        }&${XEUtils.serialize(queryParams)}`,
                        method: "GET",
                    });
                },
                delete: ({body}) => {
                    const ids = body.removeRecords.map((item: any) => item.id)
                    return destroy(ids)
                },
                save: ({body}) => {
                    const {insertRecords, updateRecords} = body
                    const promises = []

                    if (insertRecords.length > 0) {
                        insertRecords.forEach(record => {
                            promises.push(Api.courseController.store(record))
                        })
                    }

                    if (updateRecords.length > 0) {
                        updateRecords.forEach(record => {
                            promises.push(Api.courseController.update({id: record.id}, record))
                        })
                    }

                    return Promise.all(promises)
                }
            },
        },
        columns: [
            {field: "id", title: "序号",width:64},
            // 配置日期选择器
            {
                field: "title",
                title: "标题",
                sortable: true,
                editRender: {
                    name: 'AInput',
                },
            },
            {
                field: "course_type_id",
                title: "类型",
                sortable: true,
                editRender: {
                    name: 'ASelect',
                    options: typeList.value
                },
            },
            {
                field: "venue_id",
                title: "场馆",
                sortable: true,
                editRender: {
                    name: 'ASelect',
                    options: venueList.value,
                    events: {
                        change: (event, {row}) => {
                            changeCoach(event.row.venue_id)
                        }
                    }
                },
            },
            {
                field: "coach_id",
                title: "授课",
                sortable: true,
                editRender: {
                    name: 'ASelect',
                    options: teachList.value
                },
            },
            {
                field: "coupon_template_id",
                title: "定价方案",
                width: 200,
                sortable: true,
                editRender: {
                    name: 'ASelect',
                    options: tmpList.value
                },
            },
            {
                field: "total",
                title: "总价",
                sortable: true,
                editRender: {
                    name: 'AInputNumber',
                },
            },
            {
                field: "cover",
                title: "封面",
                sortable: true,
                width: 120,
                slots: {
                    default: "cover"
                }
            },
            {
                field: "description",
                title: "描述",
                width: 200,
                sortable: true,
                editRender: {
                    name: 'VxeTextarea',
                },
            },
            {
                field: "duration",
                title: "时长",
                sortable: true,
                editRender: {
                    name: 'AInputNumber',
                },
            },
            {
                field: "tags",
                title: "标签（逗号隔开）",
                sortable: true,
                width: 200,
                slots: {
                    default: "tags",
                },
            },
            {
                field: "state",
                title: "状态",
                editRender: {
                    name: 'ASelect',
                    options: [
                        {value: 1, label: "启用"},
                        {value: 2, label: "禁用"}
                    ]
                },
                formatter: ({cellValue}) => {
                    return cellValue === 1 ? "启用" : "禁用"
                }
            },

            {
                field: "created_at",
                title: "创建时间",
            },
            {
                width: 140,
                field: "action",
                title: "操作",
                slots: {
                    default: "action",
                },
            },
        ],
        importConfig: {},
        exportConfig: {},
        checkboxConfig: {
            labelField: "ID",
            checkStrictly: true,
        },
        editRules: {
            name: [
                {required: true, message: "必填"},
                {min: 3, max: 50, message: "名称长度在 3 到 50 个字符"},
            ],
            no: [{required: true, message: "必填"}],
        },
        editConfig: {
            trigger: "click",
            mode: "row",
            showStatus: true,
            showUpdateStatus: true,
            showInsertStatus: true,
            autoClear: true,
        },
    });

    return {
        gridOptions,
        typeList,
        tmpList,
        venueList,
    };
};
