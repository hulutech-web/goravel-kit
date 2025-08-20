import dayGridPlugin from "@fullcalendar/daygrid";
import interactionPlugin from "@fullcalendar/interaction";
import timeGridPlugin from "@fullcalendar/timegrid";
import momentPlugin from "@fullcalendar/moment";
import listPlugin from '@fullcalendar/list';
import zhLocale from "@fullcalendar/core/locales/zh-cn";
import Api from "@/api"
const colorPalette = [
    '#FF0000', // Vivid Red
    '#FF4500', // Orange Red
    '#FFA500', // Orange
    '#FFD700', // Gold
    '#FFFF00', // Yellow
    '#9ACD32', // Yellow Green
    '#00FF00', // Lime
    '#00FA9A', // Medium Spring Green
    '#00FFFF', // Cyan
    '#1E90FF', // Dodger Blue
    '#0000FF', // Blue
    '#8A2BE2', // Blue Violet
    '#9932CC', // Dark Orchid
    '#FF00FF', // Magenta
    '#FF1493', // Deep Pink
    '#FF69B4', // Hot Pink
    '#FF6347', // Tomato
    '#FF8C00', // Dark Orange
    '#7CFC00', // Lawn Green
    '#40E0D0'  // Turquoise
];

function getRandomColor() {
    return colorPalette[Math.floor(Math.random() * colorPalette.length)];
}

const formatTime = (timeString) => {
    const time = new Date(timeString);
    const hours = time.getHours().toString().padStart(2, '0');
    const minutes = time.getMinutes().toString().padStart(2, '0');
    return `${hours}:${minutes}`;
}
const baseUrl = import.meta.env.VITE_API_URL;
const port = import.meta.env.VITE_API_PORT;


export default (eventClickCallback, dateClickCallback,
                eventDidMountCallback,
                selEventCallback,
                eventAddCallback,
                eventChangeCallback,
                eventContentCallback,
                eventMouseEnterCallback,
                eventMouseLeaveCallback,
                eventDropCallback,coach_id) => {
    return {

        aspectRatio: 20,
        expandRows: false,
        plugins: [
            dayGridPlugin,
            interactionPlugin,
            timeGridPlugin,
            momentPlugin,
            listPlugin
        ],
        themeSystem: 'bootstrap', // 使用主题系统
        //eventBackgroundColor: 'rgb(214,231,254)', // 事件背景色
        eventBorderColor: '#3788d8', // 事件边框色
        // 高度：每一格400px
        height: "1000px",
        slotEventOverlap: false,
        slotMinTime: "09:00:00",
        slotMaxTime: "20:00:00",
        initialView: "timeGridDay",
        // dayHeaderFormat: {
        //     weekday: "short",
        // },
        contentHeight:600,
        slotDuration: "00:10:00",
        slotLabelInterval: "01:00:00",
        selectMirror: false,
        locale: zhLocale,
        buttonText: {
            month: "月",
            week: "周",
            day: "日",
            list:"列表",
        },
        headerToolbar: {
            left: "prev,next today",
            center: "title",
            right: "dayGridMonth,timeGridWeek,timeGridDay,listWeek",
        },
        customButtons: {
            button: {
                text: 'custom',
                click: function () {
                    alert('clicked the custom button!');
                }
            }
        },
        editable: false,
        selectable: false,//禁止按下拖动添加事件
        weekNumbers: true,
        views: {
            timeGridFourDay: {
                type: "timeGrid",
                duration: {
                    days: 7,
                },
            },
        },
        events: async (fetchInfo, successCallback, failureCallback) => {
            console.log("1、fetchInfo",coach_id)
            let data = await Api.scheduleController.list({startStr: fetchInfo.startStr, endStr: fetchInfo.endStr,coach_id:coach_id})
            let req = data.map(item => {
                return {
                    ...item,
                    title: item.course.title,
                    start: item.start_time,
                    end: item.end_time,
                }
            })
            console.log("1、req",req)
            await successCallback(req)
            return false;
        },
        eventClick: (info) => {
            eventClickCallback(info)
        },
        // 点击空白处触发
        dateClick: function (info) {
            dateClickCallback(info)
        },
        eventDidMount: function (info) {
            //每一个事件都会触发
            eventDidMountCallback(info)
        },
        select: function (info) {
            // 转换为中国时间
            let startTime = new Date(info.start.setHours(info.start.getHours() + 8));
            let endTime = new Date(info.end.setHours(info.end.getHours() + 8));
            let newColor = getRandomColor()
            this.setOption('eventColor', newColor);
            // 创建新的事件对象，并设置颜色
            let event = {
                title: '新事件', // 事件标题
                start: startTime, // 开始时间
                end: endTime, // 结束时间
                backgroundColor: newColor, // 设置背景颜色
                borderColor: newColor, // 设置边框颜色，可选
                textColor: 'white', // 设置文字颜色，可选
            };
            //临时事件到页面
            selEventCallback(event);
        },
        eventAdd: function (info) {
            eventAddCallback(info)
        },
        eventChange: function (info) {
            console.log("1、eventChange", info)
            eventChangeCallback(info)
        },
        eventContent: (arg, createElement) => {
            const startTime = arg.event.start;
            const endTime = arg.event.end;
            //获取背景颜色
            const color = arg.backgroundColor;
            // 创建一个容器元素来包裹事件内容
            const containerEl = document.createElement('div');
            containerEl.style.height = '100%';
            containerEl.style.overflow = 'hidden'; // 防止内容溢出
            // 设置containerEl的css属性
            containerEl.innerHTML = `
            <div style='background-color:${color};height:100%;max-width:300px;'>
            <div class='bg-gray-800  text-md font-bold' style="box-sizing:border-box;max-width:300px;">
            ⏰${formatTime(startTime)} - ${formatTime(endTime)}${arg.event.title} #${arg.event.id}
            </div>
                <div style='padding:6px;box-sizing:border-box;'>
                    <div class='text-xl font-bold'>👩🏻‍🏫${arg.event.extendedProps.coach?arg.event.extendedProps.coach.username:'未知'}</div>
                    <div class='flex justify-between items-center overflow-hidden'>
                        <div>${arg.event.title ? arg.event.title : "未知"}</div>
                        <div>${arg.event.extendedProps.course? arg.event.extendedProps.course.venue.name : "未知"}</div>
                    </div>
                    <div class='flex justify-between items-center overflow-hidden'>
                        <div>约${arg.event.extendedProps.schedule ? arg.event.extendedProps.schedule : "未知"}</div>
                        <div>签${arg.event.extendedProps.sign ? arg.event.extendedProps.sign : "未知"}</div>
                    </div>
                </div>
    </div>
    `;
            eventContentCallback(arg, createElement, containerEl)
            return {domNodes: [containerEl]};
        },
        eventMouseEnter: function (info) {
            eventMouseEnterCallback(info)
        },
        eventMouseLeave: function (info) {
            eventMouseLeaveCallback(info)
        },
        eventDrop: function (info) {
            eventDropCallback(info)
        },
        dayMaxEventRows:2,
        // 设置中国时区【坑】
        timeZone: 'local',
        dayHeaderContent: function (args) {
            //设置背景色，设置字体大小，使用h函数
            return h('div', {
                style: {
                    width: '100%',
                    lineHeight: '50px',
                    fontSize: '18px',
                    height: '50px',
                    color:'#1E90FFFF'
                }
            }, args.text);
        }
    }
}