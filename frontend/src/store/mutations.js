import Vue from 'vue'

import types from './mutation_types'


export default {
  [types.INDEXBLOCKLIST] (state, {indexBlockList}) {
    state.indexBlockList = indexBlockList
  },

  [types.MARKETINFO] (state, {marketInfo}) {
    state.marketInfo = marketInfo
  },

  [types.INDEXTXNLIST] (state, {indexTxnList}) {
    state.indexTxnList = indexTxnList
  },

  [types.BLOCKINFO] (state, {blockInfo}) {
    state.blockInfo = blockInfo
  },

  [types.BLOCKDETAIL] (state, {blockDetail}) {
    state.blockDetail = blockDetail
  },

  [types.TXNINFO] (state, {txnInfo}) {
    state.txnInfo = txnInfo
  },

  [types.TXNDETAIL] (state, {txnDetail}) {
    state.txnDetail = txnDetail
  },

  [types.ACCOUNTINFO] (state, {accountInfo}) {
    state.accountInfo = accountInfo
  },

  [types.ACCOUNTDETAIL] (state, {accountDetail}) {
    state.accountDetail = accountDetail
  },

  [types.ACCOUNTTXNINFO] (state, {accountTxnInfo}) {
    state.accountTxnInfo = accountTxnInfo
  },

  // [DECREMENT_FOOD_COUNT](state, {food}) {
  //   /*
  //   * 因为删除的按钮，是有动画效果的，
  //   *   可能会存在删除按钮没有消失时，连续的点击
  //   * 做判断，就是为了避免。>0才能减
  //   * */
  //   if (food.count) {
  //     food.count--
  //     //也要从购物车列表中删除
  //     if (food.count === 0) {
  //       state.shopCartFoods.splice(state.shopCartFoods.indexOf(food), 1)
  //     }
  //   }
  // },

}
