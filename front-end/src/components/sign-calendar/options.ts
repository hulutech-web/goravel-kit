import dayGridPlugin from "@fullcalendar/daygrid";
import interactionPlugin from "@fullcalendar/interaction";
import timeGridPlugin from "@fullcalendar/timegrid";
import momentPlugin from "@fullcalendar/moment";
import zhLocale from "@fullcalendar/core/locales/zh-cn";
import dayjs from "dayjs"
// import bootstrap5Plugin from '@fullcalendar/bootstrap5';
//import 'bootstrap/dist/css/bootstrap.css';
//import 'bootstrap-icons/font/bootstrap-icons.css'; // needs additional webpack config!
import Api from "@/api";

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
                eventDropCallback, user_id) => {
    return {
        themeSystem: 'bootstrap', // 使用主题系统

        aspectRatio: 10,
        expandRows: true,
        plugins: [
            dayGridPlugin,
            interactionPlugin,
            timeGridPlugin,
            momentPlugin,
            // bootstrap5Plugin
        ],
        // 高度：每一格400px
        height: "1000px",
        slotMinTime: "09:00:00",
        slotMaxTime: "20:00:00",
        initialView: "timeGridDay",
        // dayHeaderFormat: {
        //     weekday: "short",
        // },
        contentHeight:600,
        slotDuration: "00:10:00",
        slotLabelInterval: "01:00:00",
        selectMirror: true,
        // 关闭拖拽
        locale: zhLocale,
        buttonText: {
            month: "月",
            week: "周",
            day: "日",
        },
        headerToolbar: {
            left: "prev,next today",
            center: "title",
            right: "dayGridMonth,timeGridWeek,timeGridDay",
        },
        customButtons: {
            button: {
                text: 'custom',
                click: function () {
                    alert('clicked the custom button!');
                }
            }
        },
        selectable: false,
        weekNumbers: true,
        slotEventOverlap: false, //相同时间段的多个日程视觉上是否允许重叠，默认true允许
        views: {
            timeGridFourDay: {
                type: "timeGrid",
                duration: {
                    days: 1,
                },
            },
        },
        events: async (fetchInfo, successCallback, failureCallback) => {
            let data = await Api.scheduleController.list({
                startStr: fetchInfo.startStr,
                endStr: fetchInfo.endStr
            })

            let req = data.map(item => {
                return {
                    ...item,
                    title: item.course.title,
                    start: item.start_time,
                    end: item.end_time,
                }
            })
            await successCallback(req)
        },
        eventColor: "#ffffff",
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
            // console.log("1、newColor",newColor)
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
            eventChangeCallback(info)
        },
        eventContent: (arg, createElement) => {
            const startTime = arg.event.start;
            const endTime = arg.event.end;
            //获取背景颜色
            const color = arg.event.backgroundColor;
            // 创建一个容器元素来包裹事件内容
            const containerEl = document.createElement('div');
            containerEl.innerHTML = `
            <div style='background-color:${color};height:100%;'>
            <div class='bg-gray-800  text-md font-bold'>📅${formatTime(startTime)}-${formatTime(endTime)}</div>
                <div style='padding:6px;box-sizing:border-box;'>
                    
                    <div class='flex justify-between items-center'>
                    <div class='text-md font-bold'>${arg.event.extendedProps.coach.username}</div>
                    <div class='text-md font-bold'>${arg.event.title}</div>
                    </div>
                    <div class='flex justify-between items-center'>
                     <div>约(${arg.event.extendedProps.reservations.length}/${arg.event.extendedProps.max_students})名</div>
                        <div>签(${arg.event.extendedProps.signs.length}/${arg.event.extendedProps.max_students})名</div>
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
        editable: false,
        dayMaxEventRows: 1,
        // 设置中国时区【坑】
        timeZone: 'local',
        dayHeaderContent: function (args) {
            //设置背景色，设置字体大小，使用h函数
            return h('div', {
                style: {
                    width: '100%',
                    lineHeight: '50px',
                    fontSize: '18px',
                    height: '50px'
                }
            }, args.text);
        }
    }
}