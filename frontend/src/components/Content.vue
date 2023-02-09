<template>
    <div class="box">
        <div id="leftEditor" class="box-left">

        </div>
        <div style="border-left: 2px solid #ccc">

        </div>
        <div id="rightEditor" class="box-right">

        </div>
        <div class="toStructBtn">
            <va-button size="medium" icon="manage_history" round @click="toStructBtn"/>
        </div>
    </div>
</template>

<script lang="ts" setup>
import * as monaco from 'monaco-editor/esm/vs/editor/editor.api'
//  语法高亮
import 'monaco-editor/esm/vs/editor/editor.main.js';
import {defineProps, onMounted, ref} from "vue";
import {JsonToStruct} from "../../wailsjs/go/main/App";
import {useToast} from "vuestic-ui";
import JsonWorker from 'monaco-editor/esm/vs/language/json/json.worker.js?worker'
import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker.js?worker'
import {getValue} from "../utils/editorValue";

const props = defineProps(['menuItemValue'])

self.MonacoEnvironment = {
    getWorker(_, label) {
        switch (label) {
            case "json":
                return new JsonWorker()
            default:
                return new editorWorker()
        }
    }
}

let leftEditerRight: monaco.editor.IStandaloneCodeEditor
let leftEditerLeft: monaco.editor.IStandaloneCodeEditor
const toast = useToast()
const value = ref('')
const options: monaco.editor.IStandaloneEditorConstructionOptions = {
    minimap: {enabled: false},
    folding: true,
    detectIndentation: true,
    formatOnPaste: true,
    tabSize: 4,
    insertSpaces: false,
    trimAutoWhitespace: false,
    autoIndent: "advanced",
    overviewRulerBorder: false,
    foldingStrategy: 'indentation',
    automaticLayout: true
}

const init = () => {
    const leftEditor: HTMLElement | null = document.getElementById("leftEditor")
    if (leftEditor) {
        options.language = props.menuItemValue
        options.value = getValue(props.menuItemValue)
        leftEditerLeft = monaco.editor.create(leftEditor, options);
        leftEditerLeft.getValue()
    } else {
        console.log('组件加载失败')
    }

}

const initRight = () => {
    const rightEditor: HTMLElement | null = document.getElementById("rightEditor")
    if (rightEditor) {
        options.language = "go"
        options.value = `等待转换`
        leftEditerRight = monaco.editor.create(rightEditor, options);
        leftEditerRight.updateOptions({tabSize: 4})
        leftEditerRight.getValue()
    } else {
        console.log('组件加载失败')
    }

}

onMounted(() => {
    init()
    initRight()
})
const toStructBtn = () => {
    console.log(props.menuItemValue)
    JsonToStruct(props.menuItemValue, leftEditerLeft.getValue()).then((r) => {
        console.log(r)
        toast.init({message: "转换成功", position: 'bottom-right', offsetY: 55})
        leftEditerRight.setValue(r)
    }).catch(e => {
        toast.init({message: e, color: 'danger', position: 'bottom-right'})
    })
    setTimeout(() => {
        toast.closeAll(true)
    }, 1500)
}
</script>

<style scoped>
.box {
    width: 100%;
    height: 100%;
    display: flex;
}

.box-left {
    height: 100%;
    width: 50%;
}

.box-right {
    height: 100%;
    width: 50%;
}

.toStructBtn {
    position: fixed;
    bottom: 10px;
    right: 10px;
}
</style>