import VMdPreview from '@kangc/v-md-editor/lib/preview';
import '@kangc/v-md-editor/lib/style/preview.css';
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import '@kangc/v-md-editor/lib/theme/style/vuepress.css';

import createCopyCodePlugin from '@kangc/v-md-editor/lib/plugins/copy-code/index';
import '@kangc/v-md-editor/lib/plugins/copy-code/copy-code.css';

import createLineNumbertPlugin from '@kangc/v-md-editor/lib/plugins/line-number/index';

import createEmojiPlugin from '@kangc/v-md-editor/lib/plugins/emoji/index';
import '@kangc/v-md-editor/lib/plugins/emoji/emoji.css';

// Prism
import Prism from 'prismjs';
// highlight code
import 'prismjs/components/prism-json';

VMdPreview.use(vuepressTheme, {
    Prism,
});

// markdown支持显示代码行数
VMdPreview.use(createLineNumbertPlugin())
// markdown支持代码快速复制
VMdPreview.use(createCopyCodePlugin());
// markdown支持emoji
VMdPreview.use(createEmojiPlugin());

export const mdInstall = (app: any) => {
    app.use(VMdPreview)
}