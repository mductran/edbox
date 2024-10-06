<script setup>
import {EditorState} from "@codemirror/state";
import {EditorView, keymap} from "@codemirror/view";
import {defaultKeymap} from "@codemirror/commands";

let startState = EditorState.create({
  extensions: [keymap.of(defaultKeymap)]
})

let view = new EditorView({
  state: startState,
  parent: document.body
})

</script>

<template>
  <div id="sandbox">
  </div>
</template>

<style scoped>
  #sandbox {
    flex: 1;
    padding: 20px;
    overflow-y: auto;
    width: 100%;
    height: 100%;
  }
</style>