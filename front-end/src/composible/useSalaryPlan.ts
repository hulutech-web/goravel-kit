import Api from "@/api"
import {destroy} from "@/api/courseController";
import {VxeGridProps} from "vxe-table";
import {request} from "@/utils/request";
import XEUtils from "xe-utils";

class RowVO {
}

export default () => {
    // 方法

    // 表单配置
    const serveApiUrl = import.meta.env.VITE_API_URL;
    const gridOptions = reactive<VxeGridProps<RowVO>>({
        border: "full",
        size: "small",
        showHeaderOverflow: true,
        showOverflow: true,
        keepSource: true,
        autoResize: true,
        cellConfig: {
            height: 80,
        },

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
                    field: "plan_name",
                    title: "方案名",
                    span: 6,
                    itemRender: {
                        name: "$input",
                        props: {placeholder: "请输入队名"},
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
                        url: `/api/admin/salary_plan?pageSize=${page.pageSize}&currentPage=${page.currentPage
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
                            promises.push(Api.salaryPlanController.store(record))
                        })
                    }

                    if (updateRecords.length > 0) {
                        updateRecords.forEach(record => {
                            promises.push(Api.salaryPlanController.update({id: record.id}, record))
                        })
                    }

                    return Promise.all(promises)
                }
            },
        },
        columns: [
            {field: "id", title: "序号", width: 66},
            // 配置日期选择器
            {
                field: "plan_name",
                title: "方案名",
                editRender: {
                    name: 'AInput',
                    placeholder: '请输入方案名',
                },
            },
            {
                field: "base_salary",
                title: "底薪",
                width: 300,
                editRender: {
                    defaultValue: 0,
                    name: 'AInputNumber',
                    placeholder: '请输入',
                },
            },
            {
                field: "hourly_rate",
                title: "课时费（节）",
                sortable: true,
                editRender: {
                    defaultValue: 0,
                    name: 'AInputNumber',
                    placeholder: '请输入',
                },
            },
            {
                field: "student_rate",
                title: "提成（每人）",
                sortable: true,
                editRender: {
                    defaultValue: 0,
                    name: 'AInputNumber',
                    placeholder: '请输入',
                },
            },
            {
                field: "bonus_rate",
                title: "提成比例",
                sortable: true,
                editRender: {
                    defaultValue: 0,
                    name: 'AInputNumber',
                    placeholder: '请输入',
                },
            },
            {
                field: "tax_rate",
                title: "税率（0.6%）",
                sortable: true,
                editRender: {
                    defaultValue: 6,
                    name: 'AInputNumber',
                    placeholder: '请输入',
                },
            },
            {
                field: "effective_date",
                title: "生效时间",
                editRender: {
                    name: 'VxeDatePicker',
                    props: {
                        type: 'datetime'
                    },
                    placeholder: '请输入',
                },
            },
            {
                field: "expiry_date",
                title: "失效时间",
                editRender: {
                    name: 'VxeDatePicker',
                    props: {
                        type: 'datetime'
                    },
                    placeholder: '请输入',
                },
            },
            {
                field: "is_active",
                title: "状态",
                editRender: {
                    name: 'ASelect',
                    defaultValue: 1,
                    options: [
                        {value: 1, label: "生效"},
                        {value: 2, label: "失效"}
                    ]
                },
                formatter: ({cellValue}) => {
                    return cellValue === 1 ? "已生效√" : "已失效×"
                }
            },
            {
                field: "created_at",
                title: "创建时间",
            },
            {
                field: "action",
                title: "操作",
                sortable: true,
                width: 140,
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
    };
};
