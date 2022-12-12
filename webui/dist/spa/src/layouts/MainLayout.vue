<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>

      <q-toolbar>

        <q-toolbar-title shrink>Stoq UI</q-toolbar-title>

        <q-separator dark vertical inset />

        <q-btn stretch flat icon="home" label="Home" :to="{name: 'Index'}"/>
        <q-btn stretch flat icon="category" label="Produtos" :to="{name: 'ListProducts'}"/>
        <q-btn stretch flat icon="info" label="Sobre" :to="{name: 'ListProducts'}"/>

        <q-space />

        <div class="q-pa-md">
          <q-btn color="primary" icon="settings" label="Settings">
            <q-menu auto-close>
              <q-list style="min-width: 100px">

                <q-item clickable>
                  <q-item-section>Settings</q-item-section>
                </q-item>

                <q-separator />

                <q-item clickable>
                  <q-item-section><q-btn @click="$router.push({name: 'Login'})" color="negative" icon="logout" stretch flat label="Sair" /></q-item-section>
                </q-item>

              </q-list>
            </q-menu>
          </q-btn>
        </div>

      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>

import { defineComponent } from 'vue'

export default defineComponent({
  name: 'MainLayout',
  setup () {
    return {}
  }
})
</script>
