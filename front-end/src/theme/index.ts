import { useThemeStore } from 'stepin/es/theme-provider';
import { useSettingStore } from '@/store';



export function configTheme(key: string) {
  const { setBgSeriesColors } = useThemeStore();
  const { setNavigation } = useSettingStore();
  switch (key) {
    case 'night':
      setBgSeriesColors({ 'bg-base': '#1D1D1F' });
      break;
    case 'header-dark':
      setNavigation('head');
      break;
    default:
      setNavigation('side');
  }
}

export const themeList: Theme.ThemeConfig[] = [
  {
    title: '亮色模式',
    key: 'light',
    config: { color: { middle: { 'bg-base': '#fff' } } },
  },
  {
    title: '侧边暗色菜单',
    key: 'side-dark',
    config: { color: { middle: { 'bg-base': '#fff', 'bg-side': '#001129' } }, size: { 'width-side': '220px' } },
  },
  {
    title: '顶部暗色菜单',
    key: 'header-dark',
    config: { color: { middle: { 'bg-base': '#fff', 'bg-header': '#001129' } } },
  },
  {
    title: 'VSCode风',
    key: 'vscode',
    config: {
      color: { middle: { 'bg-base': '#23272E' } },
    },
  },
  {
    title: 'IDEA风',
    key: 'idea',
    config: {
      color: { middle: { 'bg-base': '#2B2B2B' } },
    },
  },
  {
    title: '墨绿风',
    key: 'green',
    config: {
      color: { middle: { 'bg-base': '#013a54' } },
    },
  },
  {
    title: '芭比粉',
    key: 'pink',
    config: {
      color: { middle: { 'bg-base': '#B6266D' } },
    },
  },
  {
    title: '暗夜紫',
    key: 'purple',
    config: {
      color: { middle: { 'bg-base': '#361F68' } },
    },
  },
  {
    title: '中国红',
    key: 'china',
    config: {
      color: { middle: { 'bg-base': 'rgb(230, 0, 0)' } },
    },
  },
  {
    title: '活力橙',
    key: 'orange',
    config: {
      color: { middle: { 'bg-base': '#B1740D' } },
    },
  },
];
