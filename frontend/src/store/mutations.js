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


}
