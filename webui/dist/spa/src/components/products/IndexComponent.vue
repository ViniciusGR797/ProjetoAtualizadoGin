<template>
  <div class="q-pa-sm row items-start q-gutter-sm">

    <div class="q-pa-sm row bg-grey-3 rounded-borders" style="width: 100%">
      <div class="text-h4">Lista de Produtos</div>
      <q-space />
      <q-btn flat icon="add" color="positive" :to="{name: 'CreateProduct'}" >Adicionar</q-btn>
    </div>

    <q-card v-for="item, idx in list_products" :key="idx" class="my-card" flat bordered>

      <q-card-section class="q-pt-xs">
        <q-btn flat round icon="category" />
        <div class="text-h6 q-mt-sm q-mb-xs">{{item.name}}</div>
        Código {{ item.code }} <br>
        Preço R$ {{ item.price }} <br>
      </q-card-section>

      <q-separator />

      <q-card-actions align="right">
        <q-btn flat icon="search" color="primary" :to="{name: 'UpdateProduct', params: {id: item.id, obj: item }}" >Acessar</q-btn>
        <q-btn flat icon="delete_forever" color="negative" @click="deleteProduct(item)" >Remover</q-btn>
      </q-card-actions>

    </q-card>

  </div>
</template>

<script>
import {
  mapState,
  mapActions
} from 'vuex'

export default {

  computed: {
    ...mapState('Products', [
      'list_products'
    ])
  },

  created () {
    this.setListProductsAction()
  },

  methods: {

    ...mapActions('Products', [
      'iniciarProductAction',
      'setProductAction',
      'setListProductsAction',
      'deleteProductAction'
    ]),

    async deleteProduct (product) {
      await this.setProductAction({ product })

      if (confirm('Deseja realmente deletar este produto?') === true) {
        await this.deleteProductAction()
        await this.setListProductsAction()
      } else {
        this.iniciarProductAction()
      }
    }

  }
}
</script>

<style lang="sass" scoped>
.my-card
  width: 100%
  max-width: 350px
</style>
