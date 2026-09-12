package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	_ "unsafe"
)
//go:linkname F_add_values_to_range github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_values_to_range
func F_add_values_to_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_inclusion_get_strategy_procinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_inclusion_get_strategy_procinfo
func F_inclusion_get_strategy_procinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_brin_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_brin_form_tuple
func F_brin_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_brin_memtuple_initialize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_memtuple_initialize
func F_brin_memtuple_initialize(m *base.Module, l0 int32, l1 int32)
//go:linkname F_make_attrmap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_attrmap
func F_make_attrmap(m *base.Module, l0 int32) int32
//go:linkname F_free_attrmap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_free_attrmap
func F_free_attrmap(m *base.Module, l0 int32)
//go:linkname F_build_attrmap_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_attrmap_by_name
func F_build_attrmap_by_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_attrmap_by_name_if_req github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_build_attrmap_by_name_if_req
func F_build_attrmap_by_name_if_req(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mask_unused_space github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mask_unused_space
func F_mask_unused_space(m *base.Module, l0 int32)
//go:linkname F_detoast_attr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_detoast_attr
func F_detoast_attr(m *base.Module, l0 int32) int32
//go:linkname F_toast_raw_datum_size github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_toast_raw_datum_size
func F_toast_raw_datum_size(m *base.Module, l0 int32) int32
//go:linkname F_getmissingattr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getmissingattr
func F_getmissingattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_attisnull github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_attisnull
func F_heap_attisnull(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_nocachegetattr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_nocachegetattr
func F_nocachegetattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_copytuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_copytuple
func F_heap_copytuple(m *base.Module, l0 int32) int32
//go:linkname F_heap_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_form_tuple
func F_heap_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_modify_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_modify_tuple
func F_heap_modify_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_heap_deform_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_deform_tuple
func F_heap_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_heap_modify_tuple_by_cols github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_modify_tuple_by_cols
func F_heap_modify_tuple_by_cols(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_heap_form_minimal_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_form_minimal_tuple
func F_heap_form_minimal_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_minimal_tuple_from_heap_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_minimal_tuple_from_heap_tuple
func F_minimal_tuple_from_heap_tuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_form_tuple
func F_index_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_nocache_index_getattr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_nocache_index_getattr
func F_nocache_index_getattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_index_deform_tuple_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_deform_tuple_internal
func F_index_deform_tuple_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_CopyIndexTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CopyIndexTuple
func F_CopyIndexTuple(m *base.Module, l0 int32) int32
//go:linkname F_SendRowDescriptionMessage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SendRowDescriptionMessage
func F_SendRowDescriptionMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_try_relation_open github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_try_relation_open
func F_try_relation_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_relation_close github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_relation_close
func F_relation_close(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_transformRelOptions
func F_transformRelOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_untransformRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_untransformRelOptions
func F_untransformRelOptions(m *base.Module, l0 int32) int32
//go:linkname F_build_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_reloptions
func F_build_reloptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_heap_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_reloptions
func F_heap_reloptions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_partitioned_table_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_partitioned_table_reloptions
func F_partitioned_table_reloptions(m *base.Module, l0 int32)
//go:linkname F_tablespace_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tablespace_reloptions
func F_tablespace_reloptions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ScanKeyEntryInitialize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ScanKeyEntryInitialize
func F_ScanKeyEntryInitialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_ScanKeyInit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ScanKeyInit
func F_ScanKeyInit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ss_get_location github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ss_get_location
func F_ss_get_location(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetCompressionMethodName github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetCompressionMethodName
func F_GetCompressionMethodName(m *base.Module, l0 int32) int32
//go:linkname F_toast_compress_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_toast_compress_datum
func F_toast_compress_datum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_execute_attr_map_slot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_execute_attr_map_slot
func F_execute_attr_map_slot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_populate_compact_attribute github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_populate_compact_attribute
func F_populate_compact_attribute(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateTemplateTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateTemplateTupleDesc
func F_CreateTemplateTupleDesc(m *base.Module, l0 int32) int32
//go:linkname F_CreateTupleDescCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateTupleDescCopy
func F_CreateTupleDescCopy(m *base.Module, l0 int32) int32
//go:linkname F_CreateTupleDescCopyConstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateTupleDescCopyConstr
func F_CreateTupleDescCopyConstr(m *base.Module, l0 int32) int32
//go:linkname F_TupleDescCopyEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TupleDescCopyEntry
func F_TupleDescCopyEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_FreeTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeTupleDesc
func F_FreeTupleDesc(m *base.Module, l0 int32)
//go:linkname F_TupleDescInitEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TupleDescInitEntry
func F_TupleDescInitEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_TupleDescInitBuiltinEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TupleDescInitBuiltinEntry
func F_TupleDescInitBuiltinEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_BuildDescFromLists github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BuildDescFromLists
func F_BuildDescFromLists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_TupleDescGetDefault github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TupleDescGetDefault
func F_TupleDescGetDefault(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ginFindLeafPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ginFindLeafPage
func F_ginFindLeafPage(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_freeGinBtreeStack github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_freeGinBtreeStack
func F_freeGinBtreeStack(m *base.Module, l0 int32)
//go:linkname F_ginInsertValue github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ginInsertValue
func F_ginInsertValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ginInitBA github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ginInitBA
func F_ginInitBA(m *base.Module, l0 int32)
//go:linkname F_ginInsertBAEntries github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ginInsertBAEntries
func F_ginInsertBAEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_disassembleLeaf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_disassembleLeaf
func F_disassembleLeaf(m *base.Module, l0 int32) int32
//go:linkname F_computeLeafRecompressWALData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_computeLeafRecompressWALData
func F_computeLeafRecompressWALData(m *base.Module, l0 int32)
//go:linkname F_createPostingTree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_createPostingTree
func F_createPostingTree(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ginInsertItemPointers github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ginInsertItemPointers
func F_ginInsertItemPointers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_GinFormTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GinFormTuple
func F_GinFormTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_ginReadTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ginReadTuple
func F_ginReadTuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ginCompressPostingList github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ginCompressPostingList
func F_ginCompressPostingList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ginPostingListDecode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ginPostingListDecode
func F_ginPostingListDecode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ginPostingListDecodeAllSegments github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ginPostingListDecodeAllSegments
func F_ginPostingListDecodeAllSegments(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ginMergeItemPointers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ginMergeItemPointers
func F_ginMergeItemPointers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_initGinState github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initGinState
func F_initGinState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_index_getattr_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_index_getattr_1
func F_index_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gintuple_get_key github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gintuple_get_key
func F_gintuple_get_key(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GinNewBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GinNewBuffer
func F_GinNewBuffer(m *base.Module, l0 int32) int32
//go:linkname F_ginCompareAttEntries github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ginCompareAttEntries
func F_ginCompareAttEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ginUpdateStats github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ginUpdateStats
func F_ginUpdateStats(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ginScanToDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ginScanToDelete
func F_ginScanToDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_gistinserttuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gistinserttuple
func F_gistinserttuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_gistFindCorrectParent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gistFindCorrectParent
func F_gistFindCorrectParent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_gistfinishsplit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistfinishsplit
func F_gistfinishsplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_index_getattr_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_getattr_2
func F_index_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_box_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_box_penalty
func F_box_penalty(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_g_box_consider_split github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_g_box_consider_split
func F_g_box_consider_split(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 float64, l5 int32)
//go:linkname F_gistMakeUnionItVec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gistMakeUnionItVec
func F_gistMakeUnionItVec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_gistdentryinit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistdentryinit
func F_gistdentryinit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_gistMakeUnionKey github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistMakeUnionKey
func F_gistMakeUnionKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_gistDeCompressAtt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistDeCompressAtt
func F_gistDeCompressAtt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_gistgetadjusted github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistgetadjusted
func F_gistgetadjusted(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gistchoose github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gistchoose
func F_gistchoose(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gistpenalty github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gistpenalty
func F_gistpenalty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float32
//go:linkname F_gistcheckpage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistcheckpage
func F_gistcheckpage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hashbucketcleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hashbucketcleanup
func F_hashbucketcleanup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32)
//go:linkname F__hash_doinsert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__hash_doinsert
func F__hash_doinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__hash_getbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_getbuf
func F__hash_getbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__hash_pageinit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__hash_pageinit
func F__hash_pageinit(m *base.Module, l0 int32)
//go:linkname F__hash_init github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__hash_init
func F__hash_init(m *base.Module, l0 int32, l1 float64, l2 int32) int32
//go:linkname F__hash_getcachedmetap github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__hash_getcachedmetap
func F__hash_getcachedmetap(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__hash_next github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_next
func F__hash_next(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__hash_first github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__hash_first
func F__hash_first(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__hash_get_totalbuckets github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__hash_get_totalbuckets
func F__hash_get_totalbuckets(m *base.Module, l0 int32) int32
//go:linkname F__hash_checkpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_checkpage
func F__hash_checkpage(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F__hash_convert_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__hash_convert_tuple
func F__hash_convert_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_HeapCheckForSerializableConflictOut github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HeapCheckForSerializableConflictOut
func F_HeapCheckForSerializableConflictOut(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_heap_getnext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_getnext
func F_heap_getnext(m *base.Module, l0 int32) int32
//go:linkname F_heap_getattr_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_getattr_1
func F_heap_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GetBulkInsertState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetBulkInsertState
func F_GetBulkInsertState(m *base.Module) int32
//go:linkname F_FreeBulkInsertState github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FreeBulkInsertState
func F_FreeBulkInsertState(m *base.Module, l0 int32)
//go:linkname F_heap_update github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_heap_update
func F_heap_update(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_heap_tuple_should_freeze github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_tuple_should_freeze
func F_heap_tuple_should_freeze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_HeapTupleHeaderIsOnlyLocked github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_HeapTupleHeaderIsOnlyLocked
func F_HeapTupleHeaderIsOnlyLocked(m *base.Module, l0 int32) int32
//go:linkname F_HeapTupleSatisfiesVisibility github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_HeapTupleSatisfiesVisibility
func F_HeapTupleSatisfiesVisibility(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_toast_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_heap_toast_delete
func F_heap_toast_delete(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_heap_page_prune_opt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_page_prune_opt
func F_heap_page_prune_opt(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetIndexAmRoutineByAmId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetIndexAmRoutineByAmId
func F_GetIndexAmRoutineByAmId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_IndexAmTranslateStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_IndexAmTranslateStrategy
func F_IndexAmTranslateStrategy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IndexAmTranslateCompareType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IndexAmTranslateCompareType
func F_IndexAmTranslateCompareType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_check_amproc_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_amproc_signature
func F_check_amproc_signature(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_check_amoptsproc_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_amoptsproc_signature
func F_check_amoptsproc_signature(m *base.Module, l0 int32) int32
//go:linkname F_opfamily_can_sort_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_opfamily_can_sort_type
func F_opfamily_can_sort_type(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BuildIndexValueDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BuildIndexValueDescription
func F_BuildIndexValueDescription(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_systable_beginscan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_beginscan
func F_systable_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_systable_getnext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_systable_getnext
func F_systable_getnext(m *base.Module, l0 int32) int32
//go:linkname F_systable_endscan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_systable_endscan
func F_systable_endscan(m *base.Module, l0 int32)
//go:linkname F_systable_beginscan_ordered github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_systable_beginscan_ordered
func F_systable_beginscan_ordered(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_systable_getnext_ordered github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_systable_getnext_ordered
func F_systable_getnext_ordered(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_systable_endscan_ordered github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_endscan_ordered
func F_systable_endscan_ordered(m *base.Module, l0 int32)
//go:linkname F_systable_inplace_update_begin github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_inplace_update_begin
func F_systable_inplace_update_begin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_systable_inplace_update_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_inplace_update_finish
func F_systable_inplace_update_finish(m *base.Module, l0 int32, l1 int32)
//go:linkname F_index_open github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_open
func F_index_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_beginscan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_beginscan
func F_index_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_index_beginscan_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_beginscan_internal
func F_index_beginscan_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_index_rescan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_rescan
func F_index_rescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_index_endscan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_endscan
func F_index_endscan(m *base.Module, l0 int32)
//go:linkname F_index_getnext_tid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_index_getnext_tid
func F_index_getnext_tid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_fetch_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_index_fetch_heap
func F_index_fetch_heap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_vacuum_cleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_vacuum_cleanup
func F_index_vacuum_cleanup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_getprocinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_getprocinfo
func F_index_getprocinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_index_opclass_options github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_opclass_options
func F_index_opclass_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_dedup_finish_pending github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_dedup_finish_pending
func F__bt_dedup_finish_pending(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_swap_posting github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_swap_posting
func F__bt_swap_posting(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_stepright github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_stepright
func F__bt_stepright(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__bt_insert_parent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_insert_parent
func F__bt_insert_parent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F__bt_checkpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_checkpage
func F__bt_checkpage(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_relbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_relbuf
func F__bt_relbuf(m *base.Module, l0 int32)
//go:linkname F__bt_getmeta github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_getmeta
func F__bt_getmeta(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_allocbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_allocbuf
func F__bt_allocbuf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_delitems_delete_check github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_delitems_delete_check
func F__bt_delitems_delete_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__bt_parallel_done github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_parallel_done
func F__bt_parallel_done(m *base.Module, l0 int32)
//go:linkname F__bt_search github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_search
func F__bt_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F__bt_binsrch_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__bt_binsrch_insert
func F__bt_binsrch_insert(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_readnextpage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_readnextpage
func F__bt_readnextpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F__bt_buildadd github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_buildadd
func F__bt_buildadd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__bt_mkscankey github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_mkscankey
func F__bt_mkscankey(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_freestack github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_freestack
func F__bt_freestack(m *base.Module, l0 int32)
//go:linkname F__bt_killitems github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_killitems
func F__bt_killitems(m *base.Module, l0 int32)
//go:linkname F__bt_truncate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_truncate
func F__bt_truncate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_check_third_page github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_check_third_page
func F__bt_check_third_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_spgPageIndexMultiDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_spgPageIndexMultiDelete
func F_spgPageIndexMultiDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_spgdoinsert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_spgdoinsert
func F_spgdoinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_initSpGistState github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_initSpGistState
func F_initSpGistState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SpGistUpdateMetaPage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SpGistUpdateMetaPage
func F_SpGistUpdateMetaPage(m *base.Module, l0 int32)
//go:linkname F_spgDeformLeafTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_spgDeformLeafTuple
func F_spgDeformLeafTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_addOrReplaceTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_addOrReplaceTuple
func F_addOrReplaceTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_sequence_open github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sequence_open
func F_sequence_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sequence_close github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sequence_close
func F_sequence_close(m *base.Module, l0 int32, l1 int32)
//go:linkname F_table_open github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_table_open
func F_table_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_slot_callbacks github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_table_slot_callbacks
func F_table_slot_callbacks(m *base.Module, l0 int32) int32
//go:linkname F_table_slot_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_table_slot_create
func F_table_slot_create(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_beginscan_catalog github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_table_beginscan_catalog
func F_table_beginscan_catalog(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_table_index_fetch_tuple_check github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_table_index_fetch_tuple_check
func F_table_index_fetch_tuple_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GetTsmRoutine github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetTsmRoutine
func F_GetTsmRoutine(m *base.Module, l0 int32) int32
//go:linkname F_TransactionIdSetTreeStatus github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TransactionIdSetTreeStatus
func F_TransactionIdSetTreeStatus(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64)
//go:linkname F_TransactionIdGetStatus github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TransactionIdGetStatus
func F_TransactionIdGetStatus(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_error_commit_ts_disabled github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_error_commit_ts_disabled
func F_error_commit_ts_disabled(m *base.Module)
//go:linkname F_GenericXLogStart github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GenericXLogStart
func F_GenericXLogStart(m *base.Module, l0 int32) int32
//go:linkname F_GenericXLogRegisterBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GenericXLogRegisterBuffer
func F_GenericXLogRegisterBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GenericXLogFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GenericXLogFinish
func F_GenericXLogFinish(m *base.Module, l0 int32)
//go:linkname F_MultiXactIdCreateFromMembers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MultiXactIdCreateFromMembers
func F_MultiXactIdCreateFromMembers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetMultiXactIdMembers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetMultiXactIdMembers
func F_GetMultiXactIdMembers(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MultiXactIdIsRunning github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MultiXactIdIsRunning
func F_MultiXactIdIsRunning(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReadNextMultiXactId github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReadNextMultiXactId
func F_ReadNextMultiXactId(m *base.Module) int32
//go:linkname F_SimpleLruZeroPage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SimpleLruZeroPage
func F_SimpleLruZeroPage(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_SlruInternalWritePage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SlruInternalWritePage
func F_SlruInternalWritePage(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_SimpleLruReadPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SimpleLruReadPage
func F_SimpleLruReadPage(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_SlruReportIOError github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SlruReportIOError
func F_SlruReportIOError(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_SlruInternalDeleteSegment github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SlruInternalDeleteSegment
func F_SlruInternalDeleteSegment(m *base.Module, l0 int32, l1 int64)
//go:linkname F_SlruSyncFileTag github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SlruSyncFileTag
func F_SlruSyncFileTag(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SubTransGetParent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SubTransGetParent
func F_SubTransGetParent(m *base.Module, l0 int32) int32
//go:linkname F_TruncateSUBTRANS github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TruncateSUBTRANS
func F_TruncateSUBTRANS(m *base.Module, l0 int32)
//go:linkname F_tliOfPointInHistory github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tliOfPointInHistory
func F_tliOfPointInHistory(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_tliSwitchPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tliSwitchPoint
func F_tliSwitchPoint(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_TransactionIdPrecedes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TransactionIdPrecedes
func F_TransactionIdPrecedes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TransactionIdDidAbort github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TransactionIdDidAbort
func F_TransactionIdDidAbort(m *base.Module, l0 int32) int32
//go:linkname F_ReadTwoPhaseFile github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadTwoPhaseFile
func F_ReadTwoPhaseFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FinishPreparedTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FinishPreparedTransaction
func F_FinishPreparedTransaction(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XlogReadTwoPhaseData github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XlogReadTwoPhaseData
func F_XlogReadTwoPhaseData(m *base.Module, l0 int64, l1 int32, l2 int32)
//go:linkname F_ReadNextFullTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadNextFullTransactionId
func F_ReadNextFullTransactionId(m *base.Module) int64
//go:linkname F_GetCurrentTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetCurrentTransactionId
func F_GetCurrentTransactionId(m *base.Module) int32
//go:linkname F_CommandCounterIncrement github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CommandCounterIncrement
func F_CommandCounterIncrement(m *base.Module)
//go:linkname F_StartTransactionCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_StartTransactionCommand
func F_StartTransactionCommand(m *base.Module)
//go:linkname F_CommitTransactionCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CommitTransactionCommand
func F_CommitTransactionCommand(m *base.Module)
//go:linkname F_DefineSavepoint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DefineSavepoint
func F_DefineSavepoint(m *base.Module, l0 int32)
//go:linkname F_AbortCurrentTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AbortCurrentTransaction
func F_AbortCurrentTransaction(m *base.Module)
//go:linkname F_PreventInTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreventInTransactionBlock
func F_PreventInTransactionBlock(m *base.Module, l0 int32, l1 int32)
//go:linkname F_WarnNoTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_WarnNoTransactionBlock
func F_WarnNoTransactionBlock(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RequireTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RequireTransactionBlock
func F_RequireTransactionBlock(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BeginTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BeginTransactionBlock
func F_BeginTransactionBlock(m *base.Module)
//go:linkname F_PrepareTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PrepareTransactionBlock
func F_PrepareTransactionBlock(m *base.Module, l0 int32) int32
//go:linkname F_EndTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EndTransactionBlock
func F_EndTransactionBlock(m *base.Module, l0 int32) int32
//go:linkname F_RollbackToSavepoint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RollbackToSavepoint
func F_RollbackToSavepoint(m *base.Module, l0 int32)
//go:linkname F_BeginInternalSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BeginInternalSubTransaction
func F_BeginInternalSubTransaction(m *base.Module, l0 int32)
//go:linkname F_RollbackAndReleaseCurrentSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RollbackAndReleaseCurrentSubTransaction
func F_RollbackAndReleaseCurrentSubTransaction(m *base.Module)
//go:linkname F_AbortOutOfAnyTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AbortOutOfAnyTransaction
func F_AbortOutOfAnyTransaction(m *base.Module)
//go:linkname F_WALInsertLockRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WALInsertLockRelease
func F_WALInsertLockRelease(m *base.Module)
//go:linkname F_WALInsertLockAcquireExclusive github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WALInsertLockAcquireExclusive
func F_WALInsertLockAcquireExclusive(m *base.Module)
//go:linkname F_XLogFlush github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogFlush
func F_XLogFlush(m *base.Module, l0 int64)
//go:linkname F_UpdateMinRecoveryPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_UpdateMinRecoveryPoint
func F_UpdateMinRecoveryPoint(m *base.Module, l0 int64, l1 int32)
//go:linkname F_WaitXLogInsertionsToFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WaitXLogInsertionsToFinish
func F_WaitXLogInsertionsToFinish(m *base.Module, l0 int64) int64
//go:linkname F_XLogWrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogWrite
func F_XLogWrite(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogSetAsyncXactLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogSetAsyncXactLSN
func F_XLogSetAsyncXactLSN(m *base.Module, l0 int64)
//go:linkname F_XLogBackgroundFlush github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_XLogBackgroundFlush
func F_XLogBackgroundFlush(m *base.Module) int32
//go:linkname F_XLogFileInitInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogFileInitInternal
func F_XLogFileInitInternal(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CheckXLogRemoved github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckXLogRemoved
func F_CheckXLogRemoved(m *base.Module, l0 int64, l1 int32)
//go:linkname F_StartupXLOG github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StartupXLOG
func F_StartupXLOG(m *base.Module)
//go:linkname F_GetLastImportantRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetLastImportantRecPtr
func F_GetLastImportantRecPtr(m *base.Module) int64
//go:linkname F_LogCheckpointStart github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LogCheckpointStart
func F_LogCheckpointStart(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CheckPointGuts github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CheckPointGuts
func F_CheckPointGuts(m *base.Module, l0 int64, l1 int32)
//go:linkname F_KeepLogSeg github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_KeepLogSeg
func F_KeepLogSeg(m *base.Module, l0 int64, l1 int32)
//go:linkname F_RemoveOldXlogFiles github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RemoveOldXlogFiles
func F_RemoveOldXlogFiles(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)
//go:linkname F_LogCheckpointEnd github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LogCheckpointEnd
func F_LogCheckpointEnd(m *base.Module, l0 int32)
//go:linkname F_RequestXLogSwitch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RequestXLogSwitch
func F_RequestXLogSwitch(m *base.Module, l0 int32) int64
//go:linkname F_RestoreArchivedFile github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RestoreArchivedFile
func F_RestoreArchivedFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_ExecuteRecoveryCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecuteRecoveryCommand
func F_ExecuteRecoveryCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_KeepFileRestoredFromArchive github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_KeepFileRestoredFromArchive
func F_KeepFileRestoredFromArchive(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogArchiveForceDone github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogArchiveForceDone
func F_XLogArchiveForceDone(m *base.Module, l0 int32)
//go:linkname F_XLogArchiveNotify github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogArchiveNotify
func F_XLogArchiveNotify(m *base.Module, l0 int32)
//go:linkname F_XLogBeginInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_XLogBeginInsert
func F_XLogBeginInsert(m *base.Module)
//go:linkname F_XLogEnsureRecordSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogEnsureRecordSpace
func F_XLogEnsureRecordSpace(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogRegisterBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogRegisterBuffer
func F_XLogRegisterBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogRegisterBufData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogRegisterBufData
func F_XLogRegisterBufData(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogInsert
func F_XLogInsert(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_XLogReaderAllocate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_XLogReaderAllocate
func F_XLogReaderAllocate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_XLogReaderFree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogReaderFree
func F_XLogReaderFree(m *base.Module, l0 int32)
//go:linkname F_XLogBeginRead github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogBeginRead
func F_XLogBeginRead(m *base.Module, l0 int32, l1 int64)
//go:linkname F_XLogReadRecord github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogReadRecord
func F_XLogReadRecord(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_WALRead github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WALRead
func F_WALRead(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_CheckForStandbyTrigger github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckForStandbyTrigger
func F_CheckForStandbyTrigger(m *base.Module) int32
//go:linkname F_GetLatestXTime github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetLatestXTime
func F_GetLatestXTime(m *base.Module) int64
//go:linkname F_WakeupRecovery github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WakeupRecovery
func F_WakeupRecovery(m *base.Module)
//go:linkname F_GetXLogReplayRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetXLogReplayRecPtr
func F_GetXLogReplayRecPtr(m *base.Module, l0 int32) int64
//go:linkname F_error_multiple_recovery_targets github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_error_multiple_recovery_targets
func F_error_multiple_recovery_targets(m *base.Module)
//go:linkname F_XLogReadBufferForRedo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogReadBufferForRedo
func F_XLogReadBufferForRedo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_XLogReadBufferForRedoExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_XLogReadBufferForRedoExtended
func F_XLogReadBufferForRedoExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_XLogInitBufferForRedo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogInitBufferForRedo
func F_XLogInitBufferForRedo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_XLogReadDetermineTimeline github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogReadDetermineTimeline
func F_XLogReadDetermineTimeline(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)
//go:linkname F_WALReadRaiseError github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WALReadRaiseError
func F_WALReadRaiseError(m *base.Module, l0 int32)
//go:linkname F_perform_base_backup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_perform_base_backup
func F_perform_base_backup(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_bbsink_forward_begin_backup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bbsink_forward_begin_backup
func F_bbsink_forward_begin_backup(m *base.Module, l0 int32)
//go:linkname F_GetWalSummaries github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetWalSummaries
func F_GetWalSummaries(m *base.Module, l0 int32, l1 int64, l2 int64) int32
//go:linkname F_yy_fatal_error_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_yy_fatal_error_1
func F_yy_fatal_error_1(m *base.Module, l0 int32)
//go:linkname F_boot_get_type_io_data github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_boot_get_type_io_data
func F_boot_get_type_io_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_ExecuteGrantStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecuteGrantStmt
func F_ExecuteGrantStmt(m *base.Module, l0 int32)
//go:linkname F_heap_getattr_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_getattr_2
func F_heap_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_aclcheck_error_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_aclcheck_error_type
func F_aclcheck_error_type(m *base.Module, l0 int32, l1 int32)
//go:linkname F_object_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_object_aclcheck
func F_object_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_object_aclcheck_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_object_aclcheck_ext
func F_object_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_pg_attribute_aclmask_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_attribute_aclmask_ext
func F_pg_attribute_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int64
//go:linkname F_pg_attribute_aclcheck_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_attribute_aclcheck_ext
func F_pg_attribute_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_pg_class_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_class_aclcheck
func F_pg_class_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pg_parameter_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_parameter_aclcheck
func F_pg_parameter_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pg_largeobject_aclcheck_snapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_largeobject_aclcheck_snapshot
func F_pg_largeobject_aclcheck_snapshot(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_object_ownercheck github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_object_ownercheck
func F_object_ownercheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_has_createrole_privilege github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_has_createrole_privilege
func F_has_createrole_privilege(m *base.Module, l0 int32) int32
//go:linkname F_has_bypassrls_privilege github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_has_bypassrls_privilege
func F_has_bypassrls_privilege(m *base.Module, l0 int32) int32
//go:linkname F_GetNewOidWithIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetNewOidWithIndex
func F_GetNewOidWithIndex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetNewRelFileNumber github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetNewRelFileNumber
func F_GetNewRelFileNumber(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_new_object_addresses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_new_object_addresses
func F_new_object_addresses(m *base.Module) int32
//go:linkname F_free_object_addresses github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_free_object_addresses
func F_free_object_addresses(m *base.Module, l0 int32)
//go:linkname F_add_object_address github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_add_object_address
func F_add_object_address(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_recordDependencyOnSingleRelExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_recordDependencyOnSingleRelExpr
func F_recordDependencyOnSingleRelExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_add_exact_object_address github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_exact_object_address
func F_add_exact_object_address(m *base.Module, l0 int32, l1 int32)
//go:linkname F_record_object_address_dependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_record_object_address_dependencies
func F_record_object_address_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_heap_create github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_create
func F_heap_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32) int32
//go:linkname F_CheckAttributeType github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckAttributeType
func F_CheckAttributeType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_InsertPgAttributeTuples github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InsertPgAttributeTuples
func F_InsertPgAttributeTuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_InsertPgClassTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InsertPgClassTuple
func F_InsertPgClassTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_heap_create_with_catalog github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_create_with_catalog
func F_heap_create_with_catalog(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32) int32
//go:linkname F_DeleteRelationTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_DeleteRelationTuple
func F_DeleteRelationTuple(m *base.Module, l0 int32)
//go:linkname F_DeleteAttributeTuples github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DeleteAttributeTuples
func F_DeleteAttributeTuples(m *base.Module, l0 int32)
//go:linkname F_AddRelationNewConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AddRelationNewConstraints
func F_AddRelationNewConstraints(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_index_update_stats github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_update_stats
func F_index_update_stats(m *base.Module, l0 int32, l1 int32, l2 float64)
//go:linkname F_FormIndexDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FormIndexDatum
func F_FormIndexDatum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_index_set_state_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_set_state_flags
func F_index_set_state_flags(m *base.Module, l0 int32, l1 int32)
//go:linkname F_IndexGetRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_IndexGetRelation
func F_IndexGetRelation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CatalogOpenIndexes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CatalogOpenIndexes
func F_CatalogOpenIndexes(m *base.Module, l0 int32) int32
//go:linkname F_CatalogTupleInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CatalogTupleInsert
func F_CatalogTupleInsert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CatalogTupleUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CatalogTupleUpdate
func F_CatalogTupleUpdate(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CatalogTupleDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CatalogTupleDelete
func F_CatalogTupleDelete(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RangeVarGetRelidExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RangeVarGetRelidExtended
func F_RangeVarGetRelidExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_recomputeNamespacePath github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recomputeNamespacePath
func F_recomputeNamespacePath(m *base.Module)
//go:linkname F_spcache_init github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_spcache_init
func F_spcache_init(m *base.Module)
//go:linkname F_spcache_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_spcache_insert
func F_spcache_insert(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RangeVarGetAndCheckCreationNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RangeVarGetAndCheckCreationNamespace
func F_RangeVarGetAndCheckCreationNamespace(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DeconstructQualifiedName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DeconstructQualifiedName
func F_DeconstructQualifiedName(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_NameListToString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_NameListToString
func F_NameListToString(m *base.Module, l0 int32) int32
//go:linkname F_OpernameGetOprid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_OpernameGetOprid
func F_OpernameGetOprid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_OpclassIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_OpclassIsVisibleExt
func F_OpclassIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_OpfamilyIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_OpfamilyIsVisibleExt
func F_OpfamilyIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CollationGetCollid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CollationGetCollid
func F_CollationGetCollid(m *base.Module, l0 int32) int32
//go:linkname F_TSParserIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TSParserIsVisibleExt
func F_TSParserIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TSDictionaryIsVisible github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TSDictionaryIsVisible
func F_TSDictionaryIsVisible(m *base.Module, l0 int32) int32
//go:linkname F_TSTemplateIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TSTemplateIsVisibleExt
func F_TSTemplateIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_ts_config_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_ts_config_oid
func F_get_ts_config_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TSConfigIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TSConfigIsVisibleExt
func F_TSConfigIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_QualifiedNameGetCreationNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_QualifiedNameGetCreationNamespace
func F_QualifiedNameGetCreationNamespace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeRangeVarFromNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeRangeVarFromNameList
func F_makeRangeVarFromNameList(m *base.Module, l0 int32) int32
//go:linkname F_isTempToastNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_isTempToastNamespace
func F_isTempToastNamespace(m *base.Module, l0 int32) int32
//go:linkname F_GetTempNamespaceProcNumber github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetTempNamespaceProcNumber
func F_GetTempNamespaceProcNumber(m *base.Module, l0 int32) int32
//go:linkname F_FindDefaultConversionProc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FindDefaultConversionProc
func F_FindDefaultConversionProc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ResetTempTableNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ResetTempTableNamespace
func F_ResetTempTableNamespace(m *base.Module)
//go:linkname F_fetch_search_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fetch_search_path
func F_fetch_search_path(m *base.Module, l0 int32) int32
//go:linkname F_RunObjectPostCreateHook github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RunObjectPostCreateHook
func F_RunObjectPostCreateHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RunObjectDropHook github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RunObjectDropHook
func F_RunObjectDropHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RunObjectPostAlterHook github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RunObjectPostAlterHook
func F_RunObjectPostAlterHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_RunObjectPostAlterHookStr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RunObjectPostAlterHookStr
func F_RunObjectPostAlterHookStr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_object_address github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_object_address
func F_get_object_address(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_check_object_ownership github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_object_ownership
func F_check_object_ownership(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_get_object_class_descr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_object_class_descr
func F_get_object_class_descr(m *base.Module, l0 int32) int32
//go:linkname F_get_object_oid_index github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_object_oid_index
func F_get_object_oid_index(m *base.Module, l0 int32) int32
//go:linkname F_get_object_catcache_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_object_catcache_oid
func F_get_object_catcache_oid(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_object_attnum_oid
func F_get_object_attnum_oid(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_namespace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_object_attnum_namespace
func F_get_object_attnum_namespace(m *base.Module, l0 int32) int32
//go:linkname F_get_catalog_object_by_oid_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_catalog_object_by_oid_extended
func F_get_catalog_object_by_oid_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getRelationDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_getRelationDescription
func F_getRelationDescription(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_getPublicationSchemaInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getPublicationSchemaInfo
func F_getPublicationSchemaInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getObjectIdentityParts github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getObjectIdentityParts
func F_getObjectIdentityParts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_strlist_to_textarray github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_strlist_to_textarray
func F_strlist_to_textarray(m *base.Module, l0 int32) int32
//go:linkname F_get_partition_parent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_partition_parent
func F_get_partition_parent(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_partition_ancestors github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_partition_ancestors
func F_get_partition_ancestors(m *base.Module, l0 int32) int32
//go:linkname F_map_partition_varattnos github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_map_partition_varattnos
func F_map_partition_varattnos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_has_partition_attrs github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_has_partition_attrs
func F_has_partition_attrs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_default_partition_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_default_partition_oid
func F_get_default_partition_oid(m *base.Module, l0 int32) int32
//go:linkname F_update_default_partition_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_update_default_partition_oid
func F_update_default_partition_oid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_proposed_default_constraint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_proposed_default_constraint
func F_get_proposed_default_constraint(m *base.Module, l0 int32) int32
//go:linkname F_GetAttrDefaultColumnAddress github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetAttrDefaultColumnAddress
func F_GetAttrDefaultColumnAddress(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errdetail_relkind_not_supported github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errdetail_relkind_not_supported
func F_errdetail_relkind_not_supported(m *base.Module, l0 int32)
//go:linkname F_CreateConstraintEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateConstraintEntry
func F_CreateConstraintEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32, l21 int32, l22 int32, l23 int32, l24 int32, l25 int32, l26 int32, l27 int32, l28 int32, l29 int32, l30 int32, l31 int32, l32 int32) int32
//go:linkname F_ConstraintNameIsUsed github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ConstraintNameIsUsed
func F_ConstraintNameIsUsed(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ChooseConstraintName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ChooseConstraintName
func F_ChooseConstraintName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_findNotNullConstraintAttnum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_findNotNullConstraintAttnum
func F_findNotNullConstraintAttnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_extractNotNullColumn github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_extractNotNullColumn
func F_extractNotNullColumn(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetNotNullConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationGetNotNullConstraints
func F_RelationGetNotNullConstraints(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DeconstructFkConstraintRow github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DeconstructFkConstraintRow
func F_DeconstructFkConstraintRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_AlterSetting github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AlterSetting
func F_AlterSetting(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ApplySetting github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ApplySetting
func F_ApplySetting(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_recordMultipleDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recordMultipleDependencies
func F_recordMultipleDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_getExtensionOfObject github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getExtensionOfObject
func F_getExtensionOfObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_checkMembershipInCurrentExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_checkMembershipInCurrentExtension
func F_checkMembershipInCurrentExtension(m *base.Module, l0 int32)
//go:linkname F_deleteDependencyRecordsForClass github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_deleteDependencyRecordsForClass
func F_deleteDependencyRecordsForClass(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_deleteDependencyRecordsForSpecific github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_deleteDependencyRecordsForSpecific
func F_deleteDependencyRecordsForSpecific(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_sequenceIsOwned github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sequenceIsOwned
func F_sequenceIsOwned(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getOwnedSequences_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getOwnedSequences_internal
func F_getOwnedSequences_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getIdentitySequence github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getIdentitySequence
func F_getIdentitySequence(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_index_constraint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_index_constraint
func F_get_index_constraint(m *base.Module, l0 int32) int32
//go:linkname F_find_inheritance_children github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_inheritance_children
func F_find_inheritance_children(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_inheritance_children_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_inheritance_children_extended
func F_find_inheritance_children_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_LargeObjectExists github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LargeObjectExists
func F_LargeObjectExists(m *base.Module, l0 int32) int32
//go:linkname F_LargeObjectExistsWithSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LargeObjectExistsWithSnapshot
func F_LargeObjectExistsWithSnapshot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetSchemaPublicationRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetSchemaPublicationRelations
func F_GetSchemaPublicationRelations(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_shdepLockAndCheckObject github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shdepLockAndCheckObject
func F_shdepLockAndCheckObject(m *base.Module, l0 int32, l1 int32)
//go:linkname F_recordDependencyOnOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_recordDependencyOnOwner
func F_recordDependencyOnOwner(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_shdepChangeDep github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shdepChangeDep
func F_shdepChangeDep(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_shdepDropDependency github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shdepDropDependency
func F_shdepDropDependency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_updateAclDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_updateAclDependencies
func F_updateAclDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_updateInitAclDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_updateInitAclDependencies
func F_updateInitAclDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_checkSharedDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_checkSharedDependencies
func F_checkSharedDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_deleteSharedDependencyRecordsFor github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_deleteSharedDependencyRecordsFor
func F_deleteSharedDependencyRecordsFor(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetSubscription github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetSubscription
func F_GetSubscription(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_UpdateSubscriptionRelState github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_UpdateSubscriptionRelState
func F_UpdateSubscriptionRelState(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32)
//go:linkname F_GetSubscriptionRelState github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetSubscriptionRelState
func F_GetSubscriptionRelState(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RemoveSubscriptionRel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RemoveSubscriptionRel
func F_RemoveSubscriptionRel(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RenameTypeInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RenameTypeInternal
func F_RenameTypeInternal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_makeArrayTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeArrayTypeName
func F_makeArrayTypeName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_log_smgrcreate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_log_smgrcreate
func F_log_smgrcreate(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationDropStorage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationDropStorage
func F_RelationDropStorage(m *base.Module, l0 int32)
//go:linkname F_RelationTruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationTruncate
func F_RelationTruncate(m *base.Module, l0 int32, l1 int32)
//go:linkname F_parse_analyze_fixedparams github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_analyze_fixedparams
func F_parse_analyze_fixedparams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_stmt_requires_parse_analysis github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_stmt_requires_parse_analysis
func F_stmt_requires_parse_analysis(m *base.Module, l0 int32) int32
//go:linkname F_query_requires_rewrite_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_query_requires_rewrite_plan
func F_query_requires_rewrite_plan(m *base.Module, l0 int32) int32
//go:linkname F_applyLockingClause github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_applyLockingClause
func F_applyLockingClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_SystemFuncName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SystemFuncName
func F_SystemFuncName(m *base.Module, l0 int32) int32
//go:linkname F_expand_grouping_sets github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_expand_grouping_sets
func F_expand_grouping_sets(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_finalize_grouping_exprs_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_finalize_grouping_exprs_walker
func F_finalize_grouping_exprs_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_substitute_grouped_columns_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_substitute_grouped_columns_mutator
func F_substitute_grouped_columns_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_build_aggregate_transfn_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_aggregate_transfn_expr
func F_build_aggregate_transfn_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_build_aggregate_finalfn_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_aggregate_finalfn_expr
func F_build_aggregate_finalfn_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_findTargetlistEntrySQL99 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_findTargetlistEntrySQL99
func F_findTargetlistEntrySQL99(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_findTargetlistEntrySQL92 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_findTargetlistEntrySQL92
func F_findTargetlistEntrySQL92(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_coerce_to_target_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_coerce_to_target_type
func F_coerce_to_target_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_can_coerce_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_can_coerce_type
func F_can_coerce_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_coerce_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_coerce_type
func F_coerce_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_coerce_to_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_coerce_to_domain
func F_coerce_to_domain(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_coerce_to_boolean github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_coerce_to_boolean
func F_coerce_to_boolean(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_coerce_to_specific_type_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_coerce_to_specific_type_typmod
func F_coerce_to_specific_type_typmod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_coerce_null_to_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_coerce_null_to_domain
func F_coerce_null_to_domain(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_enforce_generic_type_consistency github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_enforce_generic_type_consistency
func F_enforce_generic_type_consistency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_TypeCategory github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TypeCategory
func F_TypeCategory(m *base.Module, l0 int32) int32
//go:linkname F_IsBinaryCoercible github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IsBinaryCoercible
func F_IsBinaryCoercible(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assign_expr_collations github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_assign_expr_collations
func F_assign_expr_collations(m *base.Module, l0 int32, l1 int32)
//go:linkname F_makeDependencyGraphWalker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeDependencyGraphWalker
func F_makeDependencyGraphWalker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_transformExprRecurse github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformExprRecurse
func F_transformExprRecurse(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getJsonEncodingConst github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_getJsonEncodingConst
func F_getJsonEncodingConst(m *base.Module, l0 int32) int32
//go:linkname F_func_select_candidate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_func_select_candidate
func F_func_select_candidate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_funcname_signature_string github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_funcname_signature_string
func F_funcname_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_fn_arguments github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_make_fn_arguments
func F_make_fn_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_LookupFuncName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LookupFuncName
func F_LookupFuncName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_LookupFuncNameInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LookupFuncNameInternal
func F_LookupFuncNameInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_parser_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parser_errposition
func F_parser_errposition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_op_signature_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_op_signature_string
func F_op_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_sort_group_operators github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_sort_group_operators
func F_get_sort_group_operators(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_oper github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_oper
func F_oper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_compatible_oper_opid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_compatible_oper_opid
func F_compatible_oper_opid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_op github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_op
func F_make_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_refnameNamespaceItem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_refnameNamespaceItem
func F_refnameNamespaceItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_check_lateral_ref_ok github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_lateral_ref_ok
func F_check_lateral_ref_ok(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_scanNSItemForColumn github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_scanNSItemForColumn
func F_scanNSItemForColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_getRTEPermissionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getRTEPermissionInfo
func F_getRTEPermissionInfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addRTEPermissionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addRTEPermissionInfo
func F_addRTEPermissionInfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addRangeTableEntryForRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addRangeTableEntryForRelation
func F_addRangeTableEntryForRelation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_addNSItemToQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_addNSItemToQuery
func F_addNSItemToQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_expandTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expandTupleDesc
func F_expandTupleDesc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32)
//go:linkname F_transformExpressionList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformExpressionList
func F_transformExpressionList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_transformAssignmentSubscripts github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformAssignmentSubscripts
func F_transformAssignmentSubscripts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32
//go:linkname F_LookupTypeNameExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LookupTypeNameExtended
func F_LookupTypeNameExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_TypeNameToString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TypeNameToString
func F_TypeNameToString(m *base.Module, l0 int32) int32
//go:linkname F_typenameType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_typenameType
func F_typenameType(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_typenameTypeId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_typenameTypeId
func F_typenameTypeId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetColumnDefCollation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetColumnDefCollation
func F_GetColumnDefCollation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_typeTypeId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_typeTypeId
func F_typeTypeId(m *base.Module, l0 int32) int32
//go:linkname F_typeOrDomainTypeRelid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_typeOrDomainTypeRelid
func F_typeOrDomainTypeRelid(m *base.Module, l0 int32) int32
//go:linkname F_typeStringToTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_typeStringToTypeName
func F_typeStringToTypeName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parseTypeString github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_parseTypeString
func F_parseTypeString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_generateClonedIndexStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generateClonedIndexStmt
func F_generateClonedIndexStmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_transformPartitionBound github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_transformPartitionBound
func F_transformPartitionBound(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_raw_parser github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_raw_parser
func F_raw_parser(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_core_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_core_yyensure_buffer_stack
func F_core_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_core_yy_create_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_core_yy_create_buffer
func F_core_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_scanner_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scanner_yyerror
func F_scanner_yyerror(m *base.Module, l0 int32, l1 int32)
//go:linkname F_check_escape_warning github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_check_escape_warning
func F_check_escape_warning(m *base.Module, l0 int32)
//go:linkname F_addunicode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_addunicode
func F_addunicode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_scanner_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scanner_errposition
func F_scanner_errposition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_downcase_truncate_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_downcase_truncate_identifier
func F_downcase_truncate_identifier(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_truncate_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_truncate_identifier
func F_truncate_identifier(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecRenameStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecRenameStmt
func F_ExecRenameStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecAlterObjectSchemaStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecAlterObjectSchemaStmt
func F_ExecAlterObjectSchemaStmt(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecAlterOwnerStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecAlterOwnerStmt
func F_ExecAlterOwnerStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AlterObjectOwner_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AlterObjectOwner_internal
func F_AlterObjectOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_table_am_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_table_am_oid
func F_get_table_am_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_Async_Notify github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Async_Notify
func F_Async_Notify(m *base.Module, l0 int32, l1 int32)
//go:linkname F_queue_listen github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_queue_listen
func F_queue_listen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_Async_UnlistenAll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Async_UnlistenAll
func F_Async_UnlistenAll(m *base.Module)
//go:linkname F_asyncQueueReadAllNotifications github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_asyncQueueReadAllNotifications
func F_asyncQueueReadAllNotifications(m *base.Module)
//go:linkname F_cluster_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cluster_rel
func F_cluster_rel(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_check_index_is_clusterable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_index_is_clusterable
func F_check_index_is_clusterable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CommentObject github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CommentObject
func F_CommentObject(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateComments github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateComments
func F_CreateComments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_DeleteSharedComments github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DeleteSharedComments
func F_DeleteSharedComments(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetComment github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetComment
func F_GetComment(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CopyMultiInsertInfoFlush github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CopyMultiInsertInfoFlush
func F_CopyMultiInsertInfoFlush(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_BeginCopyFrom github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BeginCopyFrom
func F_BeginCopyFrom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_BeginCopyTo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BeginCopyTo
func F_BeginCopyTo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_ClosePipeToProgram github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ClosePipeToProgram
func F_ClosePipeToProgram(m *base.Module, l0 int32)
//go:linkname F_CopySendEndOfRow github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CopySendEndOfRow
func F_CopySendEndOfRow(m *base.Module, l0 int32)
//go:linkname F_create_ctas_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_ctas_internal
func F_create_ctas_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CreateIntoRelDestReceiver github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CreateIntoRelDestReceiver
func F_CreateIntoRelDestReceiver(m *base.Module, l0 int32) int32
//go:linkname F_createdb github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_createdb
func F_createdb(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_db_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_db_info
func F_get_db_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32) int32
//go:linkname F_get_database_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_database_oid
func F_get_database_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_database_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_database_name
func F_get_database_name(m *base.Module, l0 int32) int32
//go:linkname F_remove_dbtablespaces github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_remove_dbtablespaces
func F_remove_dbtablespaces(m *base.Module, l0 int32)
//go:linkname F_movedb github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_movedb
func F_movedb(m *base.Module, l0 int32, l1 int32)
//go:linkname F_heap_getattr_6 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_getattr_6
func F_heap_getattr_6(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_defGetString github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_defGetString
func F_defGetString(m *base.Module, l0 int32) int32
//go:linkname F_defGetBoolean github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_defGetBoolean
func F_defGetBoolean(m *base.Module, l0 int32) int32
//go:linkname F_defGetInt32 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_defGetInt32
func F_defGetInt32(m *base.Module, l0 int32) int32
//go:linkname F_defGetInt64 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_defGetInt64
func F_defGetInt64(m *base.Module, l0 int32) int64
//go:linkname F_errorConflictingDefElem github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errorConflictingDefElem
func F_errorConflictingDefElem(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SetDatabaseHasLoginEventTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SetDatabaseHasLoginEventTriggers
func F_SetDatabaseHasLoginEventTriggers(m *base.Module)
//go:linkname F_ExplainOneUtility github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExplainOneUtility
func F_ExplainOneUtility(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_standard_ExplainOneQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_standard_ExplainOneQuery
func F_standard_ExplainOneQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_CreateExplainSerializeDestReceiver github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CreateExplainSerializeDestReceiver
func F_CreateExplainSerializeDestReceiver(m *base.Module, l0 int32) int32
//go:linkname F_ExplainIndentText github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExplainIndentText
func F_ExplainIndentText(m *base.Module, l0 int32)
//go:linkname F_ExplainPropertyInteger github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainPropertyInteger
func F_ExplainPropertyInteger(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_ExplainPropertyFloat github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExplainPropertyFloat
func F_ExplainPropertyFloat(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32)
//go:linkname F_ExplainPropertyBool github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExplainPropertyBool
func F_ExplainPropertyBool(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainOpenGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainOpenGroup
func F_ExplainOpenGroup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExplainCloseGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExplainCloseGroup
func F_ExplainCloseGroup(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainSeparatePlans github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExplainSeparatePlans
func F_ExplainSeparatePlans(m *base.Module, l0 int32)
//go:linkname F_get_extension_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_extension_name
func F_get_extension_name(m *base.Module, l0 int32) int32
//go:linkname F_extension_file_exists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_extension_file_exists
func F_extension_file_exists(m *base.Module, l0 int32) int32
//go:linkname F_AlterForeignDataWrapperOwner_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AlterForeignDataWrapperOwner_internal
func F_AlterForeignDataWrapperOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AlterForeignServerOwner_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AlterForeignServerOwner_internal
func F_AlterForeignServerOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ResolveOpClass github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResolveOpClass
func F_ResolveOpClass(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GetOperatorFromCompareType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetOperatorFromCompareType
func F_GetOperatorFromCompareType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_DefineIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_DefineIndex
func F_DefineIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)
//go:linkname F_GetDefaultOpClass github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetDefaultOpClass
func F_GetDefaultOpClass(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeObjectName github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeObjectName
func F_makeObjectName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LockViewRecurse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LockViewRecurse
func F_LockViewRecurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_LockTableRecurse github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LockTableRecurse
func F_LockTableRecurse(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_SetMatViewPopulatedState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SetMatViewPopulatedState
func F_SetMatViewPopulatedState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationBuildRowSecurity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationBuildRowSecurity
func F_RelationBuildRowSecurity(m *base.Module, l0 int32)
//go:linkname F_PersistHoldablePortal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PersistHoldablePortal
func F_PersistHoldablePortal(m *base.Module, l0 int32)
//go:linkname F_EvaluateParams github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EvaluateParams
func F_EvaluateParams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_FetchPreparedStatement github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FetchPreparedStatement
func F_FetchPreparedStatement(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DropPreparedStatement github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DropPreparedStatement
func F_DropPreparedStatement(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_language_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_language_oid
func F_get_language_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AlterPublicationOwner_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AlterPublicationOwner_internal
func F_AlterPublicationOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AlterSchemaOwner_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AlterSchemaOwner_internal
func F_AlterSchemaOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecSecLabelStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecSecLabelStmt
func F_ExecSecLabelStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F_DeleteSharedSecurityLabel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DeleteSharedSecurityLabel
func F_DeleteSharedSecurityLabel(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fill_seq_fork_with_data github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fill_seq_fork_with_data
func F_fill_seq_fork_with_data(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_nextval_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_nextval_internal
func F_nextval_internal(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_do_setval github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_do_setval
func F_do_setval(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_parse_subscription_options github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_parse_subscription_options
func F_parse_subscription_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_publicationListToArray github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_publicationListToArray
func F_publicationListToArray(m *base.Module, l0 int32) int32
//go:linkname F_AlterSubscription_refresh github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AlterSubscription_refresh
func F_AlterSubscription_refresh(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReplicationSlotDropAtPubNode github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReplicationSlotDropAtPubNode
func F_ReplicationSlotDropAtPubNode(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AlterSubscriptionOwner_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AlterSubscriptionOwner_internal
func F_AlterSubscriptionOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CheckTableNotInUse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CheckTableNotInUse
func F_CheckTableNotInUse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BuildDescForRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BuildDescForRelation
func F_BuildDescForRelation(m *base.Module, l0 int32) int32
//go:linkname F_SetRelationHasSubclass github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SetRelationHasSubclass
func F_SetRelationHasSubclass(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CloneRowTriggersToPartition github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CloneRowTriggersToPartition
func F_CloneRowTriggersToPartition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CloneForeignKeyConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CloneForeignKeyConstraints
func F_CloneForeignKeyConstraints(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_set_attnotnull github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_attnotnull
func F_set_attnotnull(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ConstraintImpliedByRelConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ConstraintImpliedByRelConstraint
func F_ConstraintImpliedByRelConstraint(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecuteTruncateGuts github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecuteTruncateGuts
func F_ExecuteTruncateGuts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_check_for_column_name_collision github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_check_for_column_name_collision
func F_check_for_column_name_collision(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ATController github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ATController
func F_ATController(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_AlterTableGetLockLevel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AlterTableGetLockLevel
func F_AlterTableGetLockLevel(m *base.Module, l0 int32) int32
//go:linkname F_ATSimplePermissions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ATSimplePermissions
func F_ATSimplePermissions(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ATParseTransformCmd github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ATParseTransformCmd
func F_ATParseTransformCmd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_AlterConstrUpdateConstraintEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AlterConstrUpdateConstraintEntry
func F_AlterConstrUpdateConstraintEntry(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_tablespace_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_tablespace_oid
func F_get_tablespace_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_remove_tablespace_symlink github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_remove_tablespace_symlink
func F_remove_tablespace_symlink(m *base.Module, l0 int32)
//go:linkname F_PrepareTempTablespaces github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PrepareTempTablespaces
func F_PrepareTempTablespaces(m *base.Module)
//go:linkname F_get_tablespace_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_tablespace_name
func F_get_tablespace_name(m *base.Module, l0 int32) int32
//go:linkname F_CreateTrigger github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateTrigger
func F_CreateTrigger(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_RelationBuildTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationBuildTriggers
func F_RelationBuildTriggers(m *base.Module, l0 int32)
//go:linkname F_ExecBSInsertTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBSInsertTriggers
func F_ExecBSInsertTriggers(m *base.Module, l0 int32, l1 int32)
//go:linkname F_TriggerEnabled github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TriggerEnabled
func F_TriggerEnabled(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ExecASInsertTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecASInsertTriggers
func F_ExecASInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_TransitionTableAddTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TransitionTableAddTuple
func F_TransitionTableAddTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_cancel_prior_stmt_triggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cancel_prior_stmt_triggers
func F_cancel_prior_stmt_triggers(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_afterTriggerAddEvent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_afterTriggerAddEvent
func F_afterTriggerAddEvent(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecBRInsertTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecBRInsertTriggers
func F_ExecBRInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecARInsertTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecARInsertTriggers
func F_ExecARInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ExecIRInsertTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecIRInsertTriggers
func F_ExecIRInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecBSDeleteTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecBSDeleteTriggers
func F_ExecBSDeleteTriggers(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecBRDeleteTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecBRDeleteTriggers
func F_ExecBRDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_ExecARDeleteTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecARDeleteTriggers
func F_ExecARDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ExecIRDeleteTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecIRDeleteTriggers
func F_ExecIRDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecBSUpdateTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBSUpdateTriggers
func F_ExecBSUpdateTriggers(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecASUpdateTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecASUpdateTriggers
func F_ExecASUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecBRUpdateTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecBRUpdateTriggers
func F_ExecBRUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_ExecARUpdateTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecARUpdateTriggers
func F_ExecARUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_ExecIRUpdateTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecIRUpdateTriggers
func F_ExecIRUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_MakeTransitionCaptureState github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MakeTransitionCaptureState
func F_MakeTransitionCaptureState(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_AfterTriggerEndQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AfterTriggerEndQuery
func F_AfterTriggerEndQuery(m *base.Module, l0 int32)
//go:linkname F_afterTriggerMarkEvents github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_afterTriggerMarkEvents
func F_afterTriggerMarkEvents(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_AfterTriggerFreeQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AfterTriggerFreeQuery
func F_AfterTriggerFreeQuery(m *base.Module, l0 int32)
//go:linkname F_buildDefItem github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_buildDefItem
func F_buildDefItem(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rels_with_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_rels_with_domain
func F_get_rels_with_domain(m *base.Module, l0 int32) int32
//go:linkname F_AlterTypeOwner_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AlterTypeOwner_oid
func F_AlterTypeOwner_oid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AlterTypeOwnerInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AlterTypeOwnerInternal
func F_AlterTypeOwnerInternal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_check_role_membership_authorization github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_role_membership_authorization
func F_check_role_membership_authorization(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_check_role_grantor github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_check_role_grantor
func F_check_role_grantor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_plan_recursive_revoke github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_plan_recursive_revoke
func F_plan_recursive_revoke(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_DelRoleMems github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DelRoleMems
func F_DelRoleMems(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_vacuum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_vacuum
func F_vacuum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_vacuum_delay_point github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vacuum_delay_point
func F_vacuum_delay_point(m *base.Module, l0 int32)
//go:linkname F_ExecReScan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecReScan
func F_ExecReScan(m *base.Module, l0 int32)
//go:linkname F_ExecMarkPos github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecMarkPos
func F_ExecMarkPos(m *base.Module, l0 int32)
//go:linkname F_ExecRestrPos github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecRestrPos
func F_ExecRestrPos(m *base.Module, l0 int32)
//go:linkname F_ExecSupportsBackwardScan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecSupportsBackwardScan
func F_ExecSupportsBackwardScan(m *base.Module, l0 int32) int32
//go:linkname F_ExecInitExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitExpr
func F_ExecInitExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expr_setup_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_expr_setup_walker
func F_expr_setup_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecPushExprSetupSteps github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecPushExprSetupSteps
func F_ExecPushExprSetupSteps(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecInitExprRec github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitExprRec
func F_ExecInitExprRec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecInitQual github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecInitQual
func F_ExecInitQual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecInitExprList github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitExprList
func F_ExecInitExprList(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecBuildProjectionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecBuildProjectionInfo
func F_ExecBuildProjectionInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ExecPrepareExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecPrepareExpr
func F_ExecPrepareExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecPrepareQual github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecPrepareQual
func F_ExecPrepareQual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecCheck
func F_ExecCheck(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecBuildAggTrans github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBuildAggTrans
func F_ExecBuildAggTrans(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ExecBuildHash32Expr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecBuildHash32Expr
func F_ExecBuildHash32Expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ExecReadyInterpretedExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecReadyInterpretedExpr
func F_ExecReadyInterpretedExpr(m *base.Module, l0 int32)
//go:linkname F_execTuplesMatchPrepare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_execTuplesMatchPrepare
func F_execTuplesMatchPrepare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_execTuplesHashPrepare github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_execTuplesHashPrepare
func F_execTuplesHashPrepare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_BuildTupleHashTable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BuildTupleHashTable
func F_BuildTupleHashTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) int32
//go:linkname F_ExecOpenIndices github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecOpenIndices
func F_ExecOpenIndices(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecCloseIndices github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecCloseIndices
func F_ExecCloseIndices(m *base.Module, l0 int32)
//go:linkname F_ExecInsertIndexTuples github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecInsertIndexTuples
func F_ExecInsertIndexTuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_check_exclusion_constraint github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_exclusion_constraint
func F_check_exclusion_constraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_ExecInitJunkFilter github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecInitJunkFilter
func F_ExecInitJunkFilter(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecCheckPermissions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecCheckPermissions
func F_ExecCheckPermissions(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecutorRun github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecutorRun
func F_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_ExecutorFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecutorFinish
func F_ExecutorFinish(m *base.Module, l0 int32)
//go:linkname F_ExecutorEnd github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecutorEnd
func F_ExecutorEnd(m *base.Module, l0 int32)
//go:linkname F_ExecCheckOneRelPerms github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecCheckOneRelPerms
func F_ExecCheckOneRelPerms(m *base.Module, l0 int32) int32
//go:linkname F_CheckValidResultRel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckValidResultRel
func F_CheckValidResultRel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_InitResultRelInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InitResultRelInfo
func F_InitResultRelInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ExecGetTriggerResultRel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecGetTriggerResultRel
func F_ExecGetTriggerResultRel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecPartitionCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecPartitionCheck
func F_ExecPartitionCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecBuildSlotValueDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBuildSlotValueDescription
func F_ExecBuildSlotValueDescription(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecConstraints
func F_ExecConstraints(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecUpdateLockMode github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecUpdateLockMode
func F_ExecUpdateLockMode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecFindRowMark github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecFindRowMark
func F_ExecFindRowMark(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecBuildAuxRowMark github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBuildAuxRowMark
func F_ExecBuildAuxRowMark(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_EvalPlanQual github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_EvalPlanQual
func F_EvalPlanQual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_EvalPlanQualSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EvalPlanQualSlot
func F_EvalPlanQualSlot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EvalPlanQualInit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EvalPlanQualInit
func F_EvalPlanQualInit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_EvalPlanQualEnd github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EvalPlanQualEnd
func F_EvalPlanQualEnd(m *base.Module, l0 int32)
//go:linkname F_ExecSetupPartitionTupleRouting github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecSetupPartitionTupleRouting
func F_ExecSetupPartitionTupleRouting(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecFindPartition github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecFindPartition
func F_ExecFindPartition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_InitPartitionPruneContext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InitPartitionPruneContext
func F_InitPartitionPruneContext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_find_matching_subplans_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_matching_subplans_recurse
func F_find_matching_subplans_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ExecSetTupleBound github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecSetTupleBound
func F_ExecSetTupleBound(m *base.Module, l0 int64, l1 int32)
//go:linkname F_RelationFindReplTupleByIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationFindReplTupleByIndex
func F_RelationFindReplTupleByIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_RelationFindReplTupleSeq github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationFindReplTupleSeq
func F_RelationFindReplTupleSeq(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecSimpleRelationInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecSimpleRelationInsert
func F_ExecSimpleRelationInsert(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecSimpleRelationUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecSimpleRelationUpdate
func F_ExecSimpleRelationUpdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_init_sexpr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_init_sexpr
func F_init_sexpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_ExecMakeTableFunctionResult github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecMakeTableFunctionResult
func F_ExecMakeTableFunctionResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_tupledesc_match github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tupledesc_match
func F_tupledesc_match(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecAssignScanProjectionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecAssignScanProjectionInfo
func F_ExecAssignScanProjectionInfo(m *base.Module, l0 int32)
//go:linkname F_ExecAssignScanProjectionInfoWithVarno github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecAssignScanProjectionInfoWithVarno
func F_ExecAssignScanProjectionInfoWithVarno(m *base.Module, l0 int32)
//go:linkname F_ExecStoreMinimalTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecStoreMinimalTuple
func F_ExecStoreMinimalTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MakeTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MakeTupleTableSlot
func F_MakeTupleTableSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecAllocTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecAllocTableSlot
func F_ExecAllocTableSlot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecResetTupleTable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecResetTupleTable
func F_ExecResetTupleTable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_MakeSingleTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MakeSingleTupleTableSlot
func F_MakeSingleTupleTableSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecDropSingleTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecDropSingleTupleTableSlot
func F_ExecDropSingleTupleTableSlot(m *base.Module, l0 int32)
//go:linkname F_ExecStoreHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecStoreHeapTuple
func F_ExecStoreHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecStorePinnedBufferHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecStorePinnedBufferHeapTuple
func F_ExecStorePinnedBufferHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecForceStoreHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecForceStoreHeapTuple
func F_ExecForceStoreHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecStoreAllNullTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecStoreAllNullTuple
func F_ExecStoreAllNullTuple(m *base.Module, l0 int32)
//go:linkname F_ExecFetchSlotHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecFetchSlotHeapTuple
func F_ExecFetchSlotHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecFetchSlotHeapTupleDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecFetchSlotHeapTupleDatum
func F_ExecFetchSlotHeapTupleDatum(m *base.Module, l0 int32) int32
//go:linkname F_ExecInitResultTypeTL github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitResultTypeTL
func F_ExecInitResultTypeTL(m *base.Module, l0 int32)
//go:linkname F_ExecTypeFromTL github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecTypeFromTL
func F_ExecTypeFromTL(m *base.Module, l0 int32) int32
//go:linkname F_ExecInitResultTupleSlotTL github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitResultTupleSlotTL
func F_ExecInitResultTupleSlotTL(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecInitScanTupleSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecInitScanTupleSlot
func F_ExecInitScanTupleSlot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecInitNullTupleSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecInitNullTupleSlot
func F_ExecInitNullTupleSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_slot_getsomeattrs_int github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_slot_getsomeattrs_int
func F_slot_getsomeattrs_int(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecTypeFromExprList github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecTypeFromExprList
func F_ExecTypeFromExprList(m *base.Module, l0 int32) int32
//go:linkname F_BlessTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BlessTupleDesc
func F_BlessTupleDesc(m *base.Module, l0 int32) int32
//go:linkname F_HeapTupleHeaderGetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HeapTupleHeaderGetDatum
func F_HeapTupleHeaderGetDatum(m *base.Module, l0 int32) int32
//go:linkname F_begin_tup_output_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_begin_tup_output_tupdesc
func F_begin_tup_output_tupdesc(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_do_tup_output github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_do_tup_output
func F_do_tup_output(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CreateExecutorState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateExecutorState
func F_CreateExecutorState(m *base.Module) int32
//go:linkname F_FreeExecutorState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FreeExecutorState
func F_FreeExecutorState(m *base.Module, l0 int32)
//go:linkname F_CreateExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateExprContext
func F_CreateExprContext(m *base.Module, l0 int32) int32
//go:linkname F_ReScanExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReScanExprContext
func F_ReScanExprContext(m *base.Module, l0 int32)
//go:linkname F_MakePerTupleExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MakePerTupleExprContext
func F_MakePerTupleExprContext(m *base.Module, l0 int32) int32
//go:linkname F_ExecAssignExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecAssignExprContext
func F_ExecAssignExprContext(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecGetResultSlotOps github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetResultSlotOps
func F_ExecGetResultSlotOps(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecAssignProjectionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecAssignProjectionInfo
func F_ExecAssignProjectionInfo(m *base.Module, l0 int32)
//go:linkname F_ExecConditionalAssignProjectionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecConditionalAssignProjectionInfo
func F_ExecConditionalAssignProjectionInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecCreateScanSlotFromOuterPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecCreateScanSlotFromOuterPlan
func F_ExecCreateScanSlotFromOuterPlan(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecOpenScanRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecOpenScanRelation
func F_ExecOpenScanRelation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecGetRangeTableRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecGetRangeTableRelation
func F_ExecGetRangeTableRelation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecInitRangeTable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitRangeTable
func F_ExecInitRangeTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecInitResultRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecInitResultRelation
func F_ExecInitResultRelation(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_executor_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_executor_errposition
func F_executor_errposition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RegisterExprContextCallback github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RegisterExprContextCallback
func F_RegisterExprContextCallback(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecGetTriggerOldSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetTriggerOldSlot
func F_ExecGetTriggerOldSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetTriggerNewSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecGetTriggerNewSlot
func F_ExecGetTriggerNewSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetChildToRootMap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecGetChildToRootMap
func F_ExecGetChildToRootMap(m *base.Module, l0 int32) int32
//go:linkname F_ExecGetRootToChildMap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetRootToChildMap
func F_ExecGetRootToChildMap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetInsertedCols github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecGetInsertedCols
func F_ExecGetInsertedCols(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetUpdatedCols github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecGetUpdatedCols
func F_ExecGetUpdatedCols(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_coerce_fn_result_column github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_coerce_fn_result_column
func F_coerce_fn_result_column(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_InstrAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_InstrAlloc
func F_InstrAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_InstrStartNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_InstrStartNode
func F_InstrStartNode(m *base.Module, l0 int32)
//go:linkname F_InstrStopNode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InstrStopNode
func F_InstrStopNode(m *base.Module, l0 int32, l1 float64)
//go:linkname F_find_cols_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_cols_walker
func F_find_cols_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_initialize_phase github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_initialize_phase
func F_initialize_phase(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecParallelHashIncreaseNumBatches github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecParallelHashIncreaseNumBatches
func F_ExecParallelHashIncreaseNumBatches(m *base.Module, l0 int32)
//go:linkname F_ExecParallelHashIncreaseNumBuckets github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecParallelHashIncreaseNumBuckets
func F_ExecParallelHashIncreaseNumBuckets(m *base.Module, l0 int32)
//go:linkname F_dense_alloc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dense_alloc
func F_dense_alloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_hash_memory_limit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_hash_memory_limit
func F_get_hash_memory_limit(m *base.Module) int32
//go:linkname F_ExecHashJoinSaveTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecHashJoinSaveTuple
func F_ExecHashJoinSaveTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecIndexBuildScanKeys github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecIndexBuildScanKeys
func F_ExecIndexBuildScanKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_MJEvalOuterValues github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MJEvalOuterValues
func F_MJEvalOuterValues(m *base.Module, l0 int32) int32
//go:linkname F_MJFillOuter github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MJFillOuter
func F_MJFillOuter(m *base.Module, l0 int32) int32
//go:linkname F_MJFillInner github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MJFillInner
func F_MJFillInner(m *base.Module, l0 int32) int32
//go:linkname F_ExecComputeStoredGenerated github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecComputeStoredGenerated
func F_ExecComputeStoredGenerated(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecInitMergeTupleSlots github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecInitMergeTupleSlots
func F_ExecInitMergeTupleSlots(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecLookupResultRelByOid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecLookupResultRelByOid
func F_ExecLookupResultRelByOid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecCheckPlanOutput github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecCheckPlanOutput
func F_ExecCheckPlanOutput(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecMergeNotMatched github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecMergeNotMatched
func F_ExecMergeNotMatched(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecProcessReturning github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecProcessReturning
func F_ExecProcessReturning(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ExecInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecInsert
func F_ExecInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ExecInitUpdateProjection github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecInitUpdateProjection
func F_ExecInitUpdateProjection(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecUpdate
func F_ExecUpdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ExecPendingInserts github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecPendingInserts
func F_ExecPendingInserts(m *base.Module, l0 int32)
//go:linkname F_ExecUpdateAct github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecUpdateAct
func F_ExecUpdateAct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ExecUpdateEpilogue github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecUpdateEpilogue
func F_ExecUpdateEpilogue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ExecDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecDelete
func F_ExecDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_SeqNext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SeqNext
func F_SeqNext(m *base.Module, l0 int32) int32
//go:linkname F_ExecInitSubPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecInitSubPlan
func F_ExecInitSubPlan(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecSetParamPlanMulti github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecSetParamPlanMulti
func F_ExecSetParamPlanMulti(m *base.Module, l0 int32, l1 int32)
//go:linkname F_spool_tuples github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_spool_tuples
func F_spool_tuples(m *base.Module, l0 int32, l1 int64)
//go:linkname F_WinSetMarkPosition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WinSetMarkPosition
func F_WinSetMarkPosition(m *base.Module, l0 int32, l1 int64)
//go:linkname F_SPI_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_finish
func F_SPI_finish(m *base.Module) int32
//go:linkname F__SPI_commit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__SPI_commit
func F__SPI_commit(m *base.Module, l0 int32)
//go:linkname F__SPI_rollback github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__SPI_rollback
func F__SPI_rollback(m *base.Module, l0 int32)
//go:linkname F_SPI_freetuptable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_freetuptable
func F_SPI_freetuptable(m *base.Module, l0 int32)
//go:linkname F_SPI_execute_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SPI_execute_extended
func F_SPI_execute_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SPI_execute_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_execute_plan
func F_SPI_execute_plan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_execute_plan_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SPI_execute_plan_extended
func F_SPI_execute_plan_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SPI_prepare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_prepare
func F_SPI_prepare(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__SPI_prepare_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__SPI_prepare_plan
func F__SPI_prepare_plan(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SPI_freeplan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_freeplan
func F_SPI_freeplan(m *base.Module, l0 int32)
//go:linkname F_SPI_getvalue github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_getvalue
func F_SPI_getvalue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_getbinval github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SPI_getbinval
func F_SPI_getbinval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SPI_gettypeid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SPI_gettypeid
func F_SPI_gettypeid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SPI_palloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_palloc
func F_SPI_palloc(m *base.Module, l0 int32) int32
//go:linkname F_SPI_cursor_open github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_cursor_open
func F_SPI_cursor_open(m *base.Module, l0 int32) int32
//go:linkname F_SPI_cursor_open_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_cursor_open_internal
func F_SPI_cursor_open_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SPI_cursor_open_with_paramlist github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SPI_cursor_open_with_paramlist
func F_SPI_cursor_open_with_paramlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SPI_cursor_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SPI_cursor_fetch
func F_SPI_cursor_fetch(m *base.Module, l0 int32, l1 int32)
//go:linkname F__SPI_cursor_operation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__SPI_cursor_operation
func F__SPI_cursor_operation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_SPI_cursor_close github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_cursor_close
func F_SPI_cursor_close(m *base.Module, l0 int32)
//go:linkname F_SPI_result_code_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_result_code_string
func F_SPI_result_code_string(m *base.Module, l0 int32) int32
//go:linkname F_SPI_plan_get_cached_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_plan_get_cached_plan
func F_SPI_plan_get_cached_plan(m *base.Module, l0 int32) int32
//go:linkname F_CreateTupleQueueDestReceiver github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateTupleQueueDestReceiver
func F_CreateTupleQueueDestReceiver(m *base.Module, l0 int32) int32
//go:linkname F_TupleQueueReaderNext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TupleQueueReaderNext
func F_TupleQueueReaderNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetForeignDataWrapperExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetForeignDataWrapperExtended
func F_GetForeignDataWrapperExtended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetForeignServer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetForeignServer
func F_GetForeignServer(m *base.Module, l0 int32) int32
//go:linkname F_GetForeignServerExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetForeignServerExtended
func F_GetForeignServerExtended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetForeignServerIdByRelId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetForeignServerIdByRelId
func F_GetForeignServerIdByRelId(m *base.Module, l0 int32) int32
//go:linkname F_GetFdwRoutineByServerId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetFdwRoutineByServerId
func F_GetFdwRoutineByServerId(m *base.Module, l0 int32) int32
//go:linkname F_GetFdwRoutineByRelId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetFdwRoutineByRelId
func F_GetFdwRoutineByRelId(m *base.Module, l0 int32) int32
//go:linkname F_GetFdwRoutineForRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetFdwRoutineForRelation
func F_GetFdwRoutineForRelation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hk_depth_search github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hk_depth_search
func F_hk_depth_search(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_initHyperLogLog github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_initHyperLogLog
func F_initHyperLogLog(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pairingheap_allocate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pairingheap_allocate
func F_pairingheap_allocate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pairingheap_add github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pairingheap_add
func F_pairingheap_add(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pairingheap_remove_first github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pairingheap_remove_first
func F_pairingheap_remove_first(m *base.Module, l0 int32) int32
//go:linkname F_CheckSASLAuth github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CheckSASLAuth
func F_CheckSASLAuth(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_parse_scram_secret github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_scram_secret
func F_parse_scram_secret(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_sanitize_char_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_sanitize_char_2
func F_sanitize_char_2(m *base.Module, l0 int32)
//go:linkname F_set_authn_id github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_set_authn_id
func F_set_authn_id(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lo_import_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lo_import_internal
func F_lo_import_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_role_password github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_role_password
func F_get_role_password(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plain_crypt_verify github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_plain_crypt_verify
func F_plain_crypt_verify(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_free_auth_file github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_free_auth_file
func F_free_auth_file(m *base.Module, l0 int32)
//go:linkname F_open_auth_file github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_open_auth_file
func F_open_auth_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_tokenize_auth_file github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tokenize_auth_file
func F_tokenize_auth_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_regcomp_auth_token github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_regcomp_auth_token
func F_regcomp_auth_token(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_load_hba github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_load_hba
func F_load_hba(m *base.Module) int32
//go:linkname F_check_usermap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_usermap
func F_check_usermap(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_role_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_role_2
func F_check_role_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_range_sockaddr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_range_sockaddr
func F_pg_range_sockaddr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_sockaddr_cidr_mask github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_sockaddr_cidr_mask
func F_pg_sockaddr_cidr_mask(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pq_init github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_init
func F_pq_init(m *base.Module, l0 int32) int32
//go:linkname F_pq_getkeepalivesinterval github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_getkeepalivesinterval
func F_pq_getkeepalivesinterval(m *base.Module, l0 int32) int32
//go:linkname F_pq_getbyte github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_getbyte
func F_pq_getbyte(m *base.Module) int32
//go:linkname F_pq_recvbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_recvbuf
func F_pq_recvbuf(m *base.Module) int32
//go:linkname F_internal_putbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_internal_putbytes
func F_internal_putbytes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_beginmessage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_beginmessage
func F_pq_beginmessage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pq_sendbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_sendbytes
func F_pq_sendbytes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_sendcountedtext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_sendcountedtext
func F_pq_sendcountedtext(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_sendtext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_sendtext
func F_pq_sendtext(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_send_ascii_string github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_send_ascii_string
func F_pq_send_ascii_string(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pq_sendfloat8 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_sendfloat8
func F_pq_sendfloat8(m *base.Module, l0 int32, l1 float64)
//go:linkname F_pq_endmessage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_endmessage
func F_pq_endmessage(m *base.Module, l0 int32)
//go:linkname F_pq_begintypsend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_begintypsend
func F_pq_begintypsend(m *base.Module, l0 int32)
//go:linkname F_pq_getmsgbyte github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_getmsgbyte
func F_pq_getmsgbyte(m *base.Module, l0 int32) int32
//go:linkname F_pq_getmsgint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pq_getmsgint
func F_pq_getmsgint(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_getmsgint64 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pq_getmsgint64
func F_pq_getmsgint64(m *base.Module, l0 int32) int64
//go:linkname F_pq_getmsgfloat8 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pq_getmsgfloat8
func F_pq_getmsgfloat8(m *base.Module, l0 int32) float64
//go:linkname F_pq_getmsgbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_getmsgbytes
func F_pq_getmsgbytes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_getmsgtext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_getmsgtext
func F_pq_getmsgtext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pq_getmsgstring github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_getmsgstring
func F_pq_getmsgstring(m *base.Module, l0 int32) int32
//go:linkname F_pq_getmsgrawstring github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_getmsgrawstring
func F_pq_getmsgrawstring(m *base.Module, l0 int32) int32
//go:linkname F_pq_getmsgend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_getmsgend
func F_pq_getmsgend(m *base.Module, l0 int32)
//go:linkname F_pq_parse_errornotice github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_parse_errornotice
func F_pq_parse_errornotice(m *base.Module, l0 int32, l1 int32)
//go:linkname F_bms_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_copy
func F_bms_copy(m *base.Module, l0 int32) int32
//go:linkname F_bms_make_singleton github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bms_make_singleton
func F_bms_make_singleton(m *base.Module, l0 int32) int32
//go:linkname F_bms_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_free
func F_bms_free(m *base.Module, l0 int32)
//go:linkname F_bms_union github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_union
func F_bms_union(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_intersect github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_intersect
func F_bms_intersect(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_difference github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_difference
func F_bms_difference(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_is_subset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_is_subset
func F_bms_is_subset(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_is_member github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_is_member
func F_bms_is_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_overlap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bms_overlap
func F_bms_overlap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_overlap_list github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bms_overlap_list
func F_bms_overlap_list(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_singleton_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_singleton_member
func F_bms_singleton_member(m *base.Module, l0 int32) int32
//go:linkname F_bms_add_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_add_member
func F_bms_add_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_del_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_del_member
func F_bms_del_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_add_members github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bms_add_members
func F_bms_add_members(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_add_range github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_add_range
func F_bms_add_range(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_bms_int_members github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_int_members
func F_bms_int_members(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_del_members github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bms_del_members
func F_bms_del_members(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_join github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_join
func F_bms_join(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_copyObjectImpl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_copyObjectImpl
func F_copyObjectImpl(m *base.Module, l0 int32) int32
//go:linkname F_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_equal
func F_equal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_make2_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_list_make2_impl
func F_list_make2_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_make3_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_make3_impl
func F_list_make3_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lappend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lappend
func F_lappend(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend_int github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_lappend_int
func F_lappend_int(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_lappend_oid
func F_lappend_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_insert_nth github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_insert_nth
func F_list_insert_nth(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lcons github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lcons
func F_lcons(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lcons_int github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lcons_int
func F_lcons_int(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_concat
func F_list_concat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_list_copy
func F_list_copy(m *base.Module, l0 int32) int32
//go:linkname F_list_member_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_member_ptr
func F_list_member_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_nth_cell github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_delete_nth_cell
func F_list_delete_nth_cell(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_free
func F_list_free(m *base.Module, l0 int32)
//go:linkname F_list_delete_cell github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_delete_cell
func F_list_delete_cell(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_delete_ptr
func F_list_delete_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_last github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_delete_last
func F_list_delete_last(m *base.Module, l0 int32) int32
//go:linkname F_list_delete_first_n github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_list_delete_first_n
func F_list_delete_first_n(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_append_unique github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_append_unique
func F_list_append_unique(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_append_unique_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_append_unique_ptr
func F_list_append_unique_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_append_unique_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_list_append_unique_oid
func F_list_append_unique_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_concat_unique github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_concat_unique
func F_list_concat_unique(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_free_deep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_free_deep
func F_list_free_deep(m *base.Module, l0 int32)
//go:linkname F_list_copy_head github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_copy_head
func F_list_copy_head(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_copy_tail github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_copy_tail
func F_list_copy_tail(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_sort github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_sort
func F_list_sort(m *base.Module, l0 int32, l1 int32)
//go:linkname F_makeVarFromTargetEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeVarFromTargetEntry
func F_makeVarFromTargetEntry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeWholeRowVar github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeWholeRowVar
func F_makeWholeRowVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_makeTargetEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeTargetEntry
func F_makeTargetEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_makeNullConst github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeNullConst
func F_makeNullConst(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeBoolConst github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeBoolConst
func F_makeBoolConst(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeBoolExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeBoolExpr
func F_makeBoolExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeAlias github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeAlias
func F_makeAlias(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeRelabelType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeRelabelType
func F_makeRelabelType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeRangeVar github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeRangeVar
func F_makeRangeVar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeTypeNameFromNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeTypeNameFromNameList
func F_makeTypeNameFromNameList(m *base.Module, l0 int32) int32
//go:linkname F_makeFuncExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeFuncExpr
func F_makeFuncExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeStringConst github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeStringConst
func F_makeStringConst(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeDefElem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeDefElem
func F_makeDefElem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_opclause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_opclause
func F_make_opclause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_andclause github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_andclause
func F_make_andclause(m *base.Module, l0 int32) int32
//go:linkname F_make_orclause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_orclause
func F_make_orclause(m *base.Module, l0 int32) int32
//go:linkname F_make_and_qual github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_and_qual
func F_make_and_qual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_ands_explicit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_ands_explicit
func F_make_ands_explicit(m *base.Module, l0 int32) int32
//go:linkname F_makeJsonValueExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeJsonValueExpr
func F_makeJsonValueExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_exprTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprTypmod
func F_exprTypmod(m *base.Module, l0 int32) int32
//go:linkname F_exprCollation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprCollation
func F_exprCollation(m *base.Module, l0 int32) int32
//go:linkname F_relabel_to_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_relabel_to_typmod
func F_relabel_to_typmod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expression_tree_walker_impl_x2especialized_x2e2 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_expression_tree_walker_impl_x2especialized_x2e2
func F_expression_tree_walker_impl_x2especialized_x2e2(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_exprLocation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprLocation
func F_exprLocation(m *base.Module, l0 int32) int32
//go:linkname F_fix_opfuncids_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fix_opfuncids_walker
func F_fix_opfuncids_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_opfuncid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_opfuncid
func F_set_opfuncid(m *base.Module, l0 int32)
//go:linkname F_expression_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expression_tree_walker_impl
func F_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_query_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_query_tree_walker_impl
func F_query_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_expression_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_expression_tree_mutator_impl
func F_expression_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_query_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_query_tree_mutator_impl
func F_query_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_query_or_expression_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_query_or_expression_tree_walker_impl
func F_query_or_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_query_or_expression_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_query_or_expression_tree_mutator_impl
func F_query_or_expression_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_raw_expression_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_raw_expression_tree_walker_impl
func F_raw_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_nodeToString github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_nodeToString
func F_nodeToString(m *base.Module, l0 int32) int32
//go:linkname F_bmsToString github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bmsToString
func F_bmsToString(m *base.Module, l0 int32) int32
//go:linkname F_copyParamList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_copyParamList
func F_copyParamList(m *base.Module, l0 int32) int32
//go:linkname F_BuildParamLogString github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BuildParamLogString
func F_BuildParamLogString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__jumbleNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__jumbleNode
func F__jumbleNode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AppendJumble32 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AppendJumble32
func F_AppendJumble32(m *base.Module, l0 int32, l1 int32)
//go:linkname F_stringToNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_stringToNode
func F_stringToNode(m *base.Module, l0 int32) int32
//go:linkname F_tbm_free github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tbm_free
func F_tbm_free(m *base.Module, l0 int32)
//go:linkname F_tbm_get_pageentry github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tbm_get_pageentry
func F_tbm_get_pageentry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tbm_mark_page_lossy github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tbm_mark_page_lossy
func F_tbm_mark_page_lossy(m *base.Module, l0 int32, l1 int32)
//go:linkname F_makeInteger github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeInteger
func F_makeInteger(m *base.Module, l0 int32) int32
//go:linkname F_makeBoolean github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeBoolean
func F_makeBoolean(m *base.Module, l0 int32) int32
//go:linkname F_generate_useful_gather_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_generate_useful_gather_paths
func F_generate_useful_gather_paths(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_clauselist_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clauselist_selectivity
func F_clauselist_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_clause_selectivity_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_clause_selectivity_ext
func F_clause_selectivity_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64
//go:linkname F_clamp_row_est github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clamp_row_est
func F_clamp_row_est(m *base.Module, l0 float64) float64
//go:linkname F_cost_qual_eval github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cost_qual_eval
func F_cost_qual_eval(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_cost_qual_eval_node github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cost_qual_eval_node
func F_cost_qual_eval_node(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_cost_tuplesort github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cost_tuplesort
func F_cost_tuplesort(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 float64, l5 int32, l6 float64)
//go:linkname F_cost_append github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_cost_append
func F_cost_append(m *base.Module, l0 int32)
//go:linkname F_cost_agg github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cost_agg
func F_cost_agg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64, l6 int32, l7 int32, l8 float64, l9 float64, l10 float64, l11 float64)
//go:linkname F_cost_subplan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cost_subplan
func F_cost_subplan(m *base.Module, l0 int32, l1 int32)
//go:linkname F_calc_joinrel_size_estimate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_calc_joinrel_size_estimate
func F_calc_joinrel_size_estimate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 float64, l6 int32, l7 int32) float64
//go:linkname F_set_pathtarget_cost_width github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_pathtarget_cost_width
func F_set_pathtarget_cost_width(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_process_equivalence github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_process_equivalence
func F_process_equivalence(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ec_add_clause_to_derives_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ec_add_clause_to_derives_hash
func F_ec_add_clause_to_derives_hash(m *base.Module, l0 int32, l1 int32)
//go:linkname F_generate_join_implied_equalities github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_join_implied_equalities
func F_generate_join_implied_equalities(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_generate_join_implied_equalities_normal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_join_implied_equalities_normal
func F_generate_join_implied_equalities_normal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_generate_join_implied_equalities_broken github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_join_implied_equalities_broken
func F_generate_join_implied_equalities_broken(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_relation_has_unique_index_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_relation_has_unique_index_ext
func F_relation_has_unique_index_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_append_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_append_pathkeys
func F_append_pathkeys(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_pathkey_from_sortinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_pathkey_from_sortinfo
func F_make_pathkey_from_sortinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_build_partition_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_build_partition_pathkeys
func F_build_partition_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_pathkeys_for_sortclauses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_pathkeys_for_sortclauses
func F_make_pathkeys_for_sortclauses(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_initialize_mergeclause_eclasses github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_initialize_mergeclause_eclasses
func F_initialize_mergeclause_eclasses(m *base.Module, l0 int32, l1 int32)
//go:linkname F_query_is_distinct_for github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_query_is_distinct_for
func F_query_is_distinct_for(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_create_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_plan
func F_create_plan(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replace_nestloop_params_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_replace_nestloop_params_mutator
func F_replace_nestloop_params_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_add_vars_to_targetlist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_add_vars_to_targetlist
func F_add_vars_to_targetlist(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_add_vars_to_attr_needed github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_vars_to_attr_needed
func F_add_vars_to_attr_needed(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_restriction_is_always_true github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_restriction_is_always_true
func F_restriction_is_always_true(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_memoizable github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_memoizable
func F_check_memoizable(m *base.Module, l0 int32)
//go:linkname F_build_minmax_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_minmax_path
func F_build_minmax_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_query_planner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_query_planner
func F_query_planner(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_preprocess_qual_conditions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_preprocess_qual_conditions
func F_preprocess_qual_conditions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_preprocess_groupclause github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_preprocess_groupclause
func F_preprocess_groupclause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_apply_scanjoin_target_to_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_apply_scanjoin_target_to_paths
func F_apply_scanjoin_target_to_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_create_ordinary_grouping_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_ordinary_grouping_paths
func F_create_ordinary_grouping_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_adjust_paths_for_srfs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_adjust_paths_for_srfs
func F_adjust_paths_for_srfs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_useful_pathkeys_for_distinct github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_useful_pathkeys_for_distinct
func F_get_useful_pathkeys_for_distinct(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_expression_planner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expression_planner
func F_expression_planner(m *base.Module, l0 int32) int32
//go:linkname F_set_plan_references github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_set_plan_references
func F_set_plan_references(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fix_scan_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fix_scan_expr
func F_fix_scan_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32
//go:linkname F_fix_join_expr_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fix_join_expr_mutator
func F_fix_join_expr_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fix_upper_expr_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fix_upper_expr_mutator
func F_fix_upper_expr_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_upper_references github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_set_upper_references
func F_set_upper_references(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_set_dummy_tlist_references github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_dummy_tlist_references
func F_set_dummy_tlist_references(m *base.Module, l0 int32, l1 int32)
//go:linkname F_build_tlist_index github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_build_tlist_index
func F_build_tlist_index(m *base.Module, l0 int32) int32
//go:linkname F_convert_combining_aggrefs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_combining_aggrefs
func F_convert_combining_aggrefs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_trivial_subqueryscan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_trivial_subqueryscan
func F_trivial_subqueryscan(m *base.Module, l0 int32) int32
//go:linkname F_fix_scan_expr_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fix_scan_expr_mutator
func F_fix_scan_expr_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fix_scan_expr_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fix_scan_expr_walker
func F_fix_scan_expr_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fix_windowagg_condition_expr_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fix_windowagg_condition_expr_mutator
func F_fix_windowagg_condition_expr_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_register_partpruneinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_register_partpruneinfo
func F_register_partpruneinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_inline_cte_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_inline_cte_walker
func F_inline_cte_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replace_correlation_vars_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_replace_correlation_vars_mutator
func F_replace_correlation_vars_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SS_process_sublinks github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SS_process_sublinks
func F_SS_process_sublinks(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SS_identify_outer_params github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SS_identify_outer_params
func F_SS_identify_outer_params(m *base.Module, l0 int32)
//go:linkname F_finalize_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_finalize_plan
func F_finalize_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_preprocess_aggrefs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_preprocess_aggrefs
func F_preprocess_aggrefs(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_agg_clause_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_agg_clause_costs
func F_get_agg_clause_costs(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pull_up_sublinks_jointree_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pull_up_sublinks_jointree_recurse
func F_pull_up_sublinks_jointree_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_expand_virtual_generated_columns github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_expand_virtual_generated_columns
func F_expand_virtual_generated_columns(m *base.Module, l0 int32) int32
//go:linkname F_pull_up_subqueries_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pull_up_subqueries_recurse
func F_pull_up_subqueries_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_is_simple_union_all_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_is_simple_union_all_recurse
func F_is_simple_union_all_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_relids_in_jointree github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_relids_in_jointree
func F_get_relids_in_jointree(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_reduce_outer_joins_pass1 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_reduce_outer_joins_pass1
func F_reduce_outer_joins_pass1(m *base.Module, l0 int32) int32
//go:linkname F_reduce_outer_joins_pass2 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_reduce_outer_joins_pass2
func F_reduce_outer_joins_pass2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_remove_useless_results_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_remove_useless_results_recurse
func F_remove_useless_results_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_negate_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_negate_clause
func F_negate_clause(m *base.Module, l0 int32) int32
//go:linkname F_canonicalize_qual github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_canonicalize_qual
func F_canonicalize_qual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expand_insert_targetlist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_expand_insert_targetlist
func F_expand_insert_targetlist(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_recurse_set_operations github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_recurse_set_operations
func F_recurse_set_operations(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_build_setop_child_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_build_setop_child_paths
func F_build_setop_child_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_adjust_appendrel_attrs_multilevel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_adjust_appendrel_attrs_multilevel
func F_adjust_appendrel_attrs_multilevel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_adjust_child_relids_multilevel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_adjust_child_relids_multilevel
func F_adjust_child_relids_multilevel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_adjust_inherited_attnums_multilevel github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_adjust_inherited_attnums_multilevel
func F_adjust_inherited_attnums_multilevel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_add_row_identity_columns github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_add_row_identity_columns
func F_add_row_identity_columns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_contain_agg_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_contain_agg_clause
func F_contain_agg_clause(m *base.Module, l0 int32) int32
//go:linkname F_contain_mutable_functions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_contain_mutable_functions
func F_contain_mutable_functions(m *base.Module, l0 int32) int32
//go:linkname F_contain_mutable_functions_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_contain_mutable_functions_walker
func F_contain_mutable_functions_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_contain_mutable_functions_after_planning github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_contain_mutable_functions_after_planning
func F_contain_mutable_functions_after_planning(m *base.Module, l0 int32) int32
//go:linkname F_contain_volatile_functions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_contain_volatile_functions
func F_contain_volatile_functions(m *base.Module, l0 int32) int32
//go:linkname F_is_parallel_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_is_parallel_safe
func F_is_parallel_safe(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_nonnullable_rels_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_nonnullable_rels_walker
func F_find_nonnullable_rels_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_is_pseudo_constant_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_is_pseudo_constant_clause
func F_is_pseudo_constant_clause(m *base.Module, l0 int32) int32
//go:linkname F_eval_const_expressions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_eval_const_expressions
func F_eval_const_expressions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expand_function_arguments github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_expand_function_arguments
func F_expand_function_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_simplify_function github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_simplify_function
func F_simplify_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_evaluate_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_evaluate_expr
func F_evaluate_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_rowtype_field_matches github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_rowtype_field_matches
func F_rowtype_field_matches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_convert_saop_to_hashed_saop github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_convert_saop_to_hashed_saop
func F_convert_saop_to_hashed_saop(m *base.Module, l0 int32)
//go:linkname F_estimate_expression_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_estimate_expression_value
func F_estimate_expression_value(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_inline_set_returning_function github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inline_set_returning_function
func F_inline_set_returning_function(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expand_single_inheritance_child github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expand_single_inheritance_child
func F_expand_single_inheritance_child(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_generate_new_exec_param github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_new_exec_param
func F_generate_new_exec_param(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_assign_special_exec_param github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_assign_special_exec_param
func F_assign_special_exec_param(m *base.Module, l0 int32) int32
//go:linkname F_compare_path_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_compare_path_costs
func F_compare_path_costs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_compare_fractional_path_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_compare_fractional_path_costs
func F_compare_fractional_path_costs(m *base.Module, l0 int32, l1 int32, l2 float64) int32
//go:linkname F_set_cheapest github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_cheapest
func F_set_cheapest(m *base.Module, l0 int32)
//go:linkname F_add_partial_path github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_partial_path
func F_add_partial_path(m *base.Module, l0 int32, l1 int32)
//go:linkname F_create_merge_append_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_merge_append_path
func F_create_merge_append_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_group_result_path github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_group_result_path
func F_create_group_result_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_unique_path github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_unique_path
func F_create_unique_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_apply_projection_to_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_apply_projection_to_path
func F_apply_projection_to_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_incremental_sort_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_incremental_sort_path
func F_create_incremental_sort_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) int32
//go:linkname F_create_upper_unique_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_create_upper_unique_path
func F_create_upper_unique_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32
//go:linkname F_create_agg_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_agg_path
func F_create_agg_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 float64) int32
//go:linkname F_create_limit_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_limit_path
func F_create_limit_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int64) int32
//go:linkname F_reparameterize_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_reparameterize_path
func F_reparameterize_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32
//go:linkname F_estimate_rel_size github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_estimate_rel_size
func F_estimate_rel_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_add_function_cost github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_function_cost
func F_add_function_cost(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_predicate_implied_by github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_predicate_implied_by
func F_predicate_implied_by(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_setup_simple_rel_arrays github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_setup_simple_rel_arrays
func F_setup_simple_rel_arrays(m *base.Module, l0 int32)
//go:linkname F_expand_planner_arrays github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_expand_planner_arrays
func F_expand_planner_arrays(m *base.Module, l0 int32, l1 int32)
//go:linkname F_build_simple_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_simple_rel
func F_build_simple_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_base_rel_ignore_join github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_find_base_rel_ignore_join
func F_find_base_rel_ignore_join(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fetch_upper_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fetch_upper_rel
func F_fetch_upper_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_baserel_parampathinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_baserel_parampathinfo
func F_get_baserel_parampathinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_param_path_clause_serials github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_param_path_clause_serials
func F_get_param_path_clause_serials(m *base.Module, l0 int32) int32
//go:linkname F_make_plain_restrictinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_plain_restrictinfo
func F_make_plain_restrictinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32
//go:linkname F_tlist_member github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tlist_member
func F_tlist_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_sortgroupclause_tle github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_sortgroupclause_tle
func F_get_sortgroupclause_tle(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_sortgroupclause_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_sortgroupclause_expr
func F_get_sortgroupclause_expr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_sortgroupref_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_sortgroupref_clause
func F_get_sortgroupref_clause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_pathtarget_from_tlist github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_pathtarget_from_tlist
func F_make_pathtarget_from_tlist(m *base.Module, l0 int32) int32
//go:linkname F_copy_pathtarget github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_copy_pathtarget
func F_copy_pathtarget(m *base.Module, l0 int32) int32
//go:linkname F_add_new_columns_to_pathtarget github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_new_columns_to_pathtarget
func F_add_new_columns_to_pathtarget(m *base.Module, l0 int32, l1 int32)
//go:linkname F_split_pathtarget_at_srfs github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_split_pathtarget_at_srfs
func F_split_pathtarget_at_srfs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_split_pathtarget_at_srfs_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_split_pathtarget_at_srfs_extended
func F_split_pathtarget_at_srfs_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_pull_varnos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pull_varnos
func F_pull_varnos(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pull_varnos_of_level github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pull_varnos_of_level
func F_pull_varnos_of_level(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pull_varattnos github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pull_varattnos
func F_pull_varattnos(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pull_var_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pull_var_clause
func F_pull_var_clause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_flatten_join_alias_vars github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_flatten_join_alias_vars
func F_flatten_join_alias_vars(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_flatten_group_exprs github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_flatten_group_exprs
func F_flatten_group_exprs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_qual_for_list github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_qual_for_list
func F_get_qual_for_list(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_qual_for_range github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_qual_for_range
func F_get_qual_for_range(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_new_partition_bound github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_new_partition_bound
func F_check_new_partition_bound(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RelationGetPartitionDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetPartitionDesc
func F_RelationGetPartitionDesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_PartitionDirectoryLookup github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PartitionDirectoryLookup
func F_PartitionDirectoryLookup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gen_partprune_steps_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gen_partprune_steps_internal
func F_gen_partprune_steps_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_matching_partitions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_matching_partitions
func F_get_matching_partitions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AuxiliaryProcessMainCommon github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AuxiliaryProcessMainCommon
func F_AuxiliaryProcessMainCommon(m *base.Module)
//go:linkname F_BackgroundWorkerInitializeConnectionByOid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BackgroundWorkerInitializeConnectionByOid
func F_BackgroundWorkerInitializeConnectionByOid(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetBackgroundWorkerPid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetBackgroundWorkerPid
func F_GetBackgroundWorkerPid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RequestCheckpoint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RequestCheckpoint
func F_RequestCheckpoint(m *base.Module, l0 int32)
//go:linkname F_ForwardSyncRequest github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ForwardSyncRequest
func F_ForwardSyncRequest(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MaxLivePostmasterChildren github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MaxLivePostmasterChildren
func F_MaxLivePostmasterChildren(m *base.Module) int32
//go:linkname F_AssignPostmasterChildSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AssignPostmasterChildSlot
func F_AssignPostmasterChildSlot(m *base.Module, l0 int32) int32
//go:linkname F_disable_startup_progress_timeout github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_disable_startup_progress_timeout
func F_disable_startup_progress_timeout(m *base.Module)
//go:linkname F_begin_startup_progress_phase github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_begin_startup_progress_phase
func F_begin_startup_progress_phase(m *base.Module)
//go:linkname F_logfile_getname github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_logfile_getname
func F_logfile_getname(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_logfile_open github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_logfile_open
func F_logfile_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_write_syslogger_file github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_write_syslogger_file
func F_write_syslogger_file(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_regcomp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_regcomp
func F_pg_regcomp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_set_regex_collation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_set_regex_collation
func F_pg_set_regex_collation(m *base.Module, l0 int32)
//go:linkname F_freesubre github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_freesubre
func F_freesubre(m *base.Module, l0 int32, l1 int32)
//go:linkname F_newstate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_newstate
func F_newstate(m *base.Module, l0 int32) int32
//go:linkname F_rainbow github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_rainbow
func F_rainbow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_createarc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_createarc
func F_createarc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_next github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_next
func F_next(m *base.Module, l0 int32) int32
//go:linkname F_newcolor github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_newcolor
func F_newcolor(m *base.Module, l0 int32) int32
//go:linkname F_parsebranch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parsebranch
func F_parsebranch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_dupnfa github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_dupnfa
func F_dupnfa(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_cleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cleanup
func F_cleanup(m *base.Module, l0 int32)
//go:linkname F_moveouts github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_moveouts
func F_moveouts(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_emptyreachable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_emptyreachable
func F_emptyreachable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_sortins github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sortins
func F_sortins(m *base.Module, l0 int32, l1 int32)
//go:linkname F_findconstraintloop github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_findconstraintloop
func F_findconstraintloop(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cparc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cparc
func F_cparc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_checkmatchall_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_checkmatchall_recurse
func F_checkmatchall_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_newarc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_newarc
func F_newarc(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_colorcomplement github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_colorcomplement
func F_colorcomplement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_deltraverse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_deltraverse
func F_deltraverse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_regerror github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_regerror
func F_pg_regerror(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_newdfa github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_newdfa
func F_newdfa(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_shortest github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shortest
func F_shortest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_longest github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_longest
func F_longest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_cdissect github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cdissect
func F_cdissect(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_regfree github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_regfree
func F_pg_regfree(m *base.Module, l0 int32)
//go:linkname F_pa_free_worker_info github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pa_free_worker_info
func F_pa_free_worker_info(m *base.Module, l0 int32)
//go:linkname F_pa_find_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pa_find_worker
func F_pa_find_worker(m *base.Module, l0 int32) int32
//go:linkname F_pa_send_data github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pa_send_data
func F_pa_send_data(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pa_lock_stream github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pa_lock_stream
func F_pa_lock_stream(m *base.Module, l0 int32)
//go:linkname F_pa_set_xact_state github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pa_set_xact_state
func F_pa_set_xact_state(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pa_unlock_stream github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pa_unlock_stream
func F_pa_unlock_stream(m *base.Module, l0 int32)
//go:linkname F_pa_decr_and_wait_stream_block github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pa_decr_and_wait_stream_block
func F_pa_decr_and_wait_stream_block(m *base.Module)
//go:linkname F_pa_xact_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pa_xact_finish
func F_pa_xact_finish(m *base.Module, l0 int32, l1 int64)
//go:linkname F_GetTupleTransactionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetTupleTransactionInfo
func F_GetTupleTransactionInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ReportApplyConflict github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReportApplyConflict
func F_ReportApplyConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_InitConflictIndexes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_InitConflictIndexes
func F_InitConflictIndexes(m *base.Module, l0 int32)
//go:linkname F_LogicalDecodingProcessRecord github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LogicalDecodingProcessRecord
func F_LogicalDecodingProcessRecord(m *base.Module, l0 int32, l1 int32)
//go:linkname F_logicalrep_workers_find github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_logicalrep_workers_find
func F_logicalrep_workers_find(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_logicalrep_worker_launch github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_logicalrep_worker_launch
func F_logicalrep_worker_launch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_logicalrep_worker_stop_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_logicalrep_worker_stop_internal
func F_logicalrep_worker_stop_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_logicalrep_worker_wakeup github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_logicalrep_worker_wakeup
func F_logicalrep_worker_wakeup(m *base.Module, l0 int32)
//go:linkname F_logicalrep_worker_wakeup_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_logicalrep_worker_wakeup_ptr
func F_logicalrep_worker_wakeup_ptr(m *base.Module, l0 int32)
//go:linkname F_ApplyLauncherForgetWorkerStartTime github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ApplyLauncherForgetWorkerStartTime
func F_ApplyLauncherForgetWorkerStartTime(m *base.Module, l0 int32)
//go:linkname F_CheckLogicalDecodingRequirements github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckLogicalDecodingRequirements
func F_CheckLogicalDecodingRequirements(m *base.Module)
//go:linkname F_CreateInitDecodingContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateInitDecodingContext
func F_CreateInitDecodingContext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_CreateDecodingContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateDecodingContext
func F_CreateDecodingContext(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_DecodingContextFindStartpoint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_DecodingContextFindStartpoint
func F_DecodingContextFindStartpoint(m *base.Module, l0 int32)
//go:linkname F_FreeDecodingContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FreeDecodingContext
func F_FreeDecodingContext(m *base.Module, l0 int32)
//go:linkname F_filter_prepare_cb_wrapper github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_filter_prepare_cb_wrapper
func F_filter_prepare_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_filter_by_origin_cb_wrapper github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_filter_by_origin_cb_wrapper
func F_filter_by_origin_cb_wrapper(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LogicalConfirmReceivedLocation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LogicalConfirmReceivedLocation
func F_LogicalConfirmReceivedLocation(m *base.Module, l0 int64)
//go:linkname F_UpdateDecodingStats github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_UpdateDecodingStats
func F_UpdateDecodingStats(m *base.Module, l0 int32)
//go:linkname F_replorigin_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_replorigin_by_name
func F_replorigin_by_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replorigin_create github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replorigin_create
func F_replorigin_create(m *base.Module, l0 int32) int32
//go:linkname F_replorigin_advance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_replorigin_advance
func F_replorigin_advance(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32)
//go:linkname F_replorigin_get_progress github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_replorigin_get_progress
func F_replorigin_get_progress(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_replorigin_session_setup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replorigin_session_setup
func F_replorigin_session_setup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_replorigin_session_reset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replorigin_session_reset
func F_replorigin_session_reset(m *base.Module)
//go:linkname F_logicalrep_read_prepare_common github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_logicalrep_read_prepare_common
func F_logicalrep_read_prepare_common(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_logicalrep_read_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_logicalrep_read_tuple
func F_logicalrep_read_tuple(m *base.Module, l0 int32, l1 int32)
//go:linkname F_logicalrep_message_type github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_logicalrep_message_type
func F_logicalrep_message_type(m *base.Module, l0 int32) int32
//go:linkname F_logicalrep_relmap_update github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_logicalrep_relmap_update
func F_logicalrep_relmap_update(m *base.Module, l0 int32)
//go:linkname F_logicalrep_relmap_free_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_logicalrep_relmap_free_entry
func F_logicalrep_relmap_free_entry(m *base.Module, l0 int32)
//go:linkname F_logicalrep_rel_open github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_logicalrep_rel_open
func F_logicalrep_rel_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_logicalrep_rel_close github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_logicalrep_rel_close
func F_logicalrep_rel_close(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferCleanupSerializedTXNs github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReorderBufferCleanupSerializedTXNs
func F_ReorderBufferCleanupSerializedTXNs(m *base.Module, l0 int32)
//go:linkname F_ReorderBufferFreeChange github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReorderBufferFreeChange
func F_ReorderBufferFreeChange(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReorderBufferQueueChange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReorderBufferQueueChange
func F_ReorderBufferQueueChange(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32)
//go:linkname F_ReorderBufferCommitChild github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReorderBufferCommitChild
func F_ReorderBufferCommitChild(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64)
//go:linkname F_ReorderBufferReplay github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferReplay
func F_ReorderBufferReplay(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int32, l6 int64)
//go:linkname F_ReorderBufferTruncateTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReorderBufferTruncateTXN
func F_ReorderBufferTruncateTXN(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReorderBufferCleanupTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReorderBufferCleanupTXN
func F_ReorderBufferCleanupTXN(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferSkipPrepare github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReorderBufferSkipPrepare
func F_ReorderBufferSkipPrepare(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferToastReset github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReorderBufferToastReset
func F_ReorderBufferToastReset(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferForget github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferForget
func F_ReorderBufferForget(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_ReorderBufferProcessXid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReorderBufferProcessXid
func F_ReorderBufferProcessXid(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_ReorderBufferSetBaseSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferSetBaseSnapshot
func F_ReorderBufferSetBaseSnapshot(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_ReorderBufferXidSetCatalogChanges github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferXidSetCatalogChanges
func F_ReorderBufferXidSetCatalogChanges(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_ReorderBufferXidHasCatalogChanges github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferXidHasCatalogChanges
func F_ReorderBufferXidHasCatalogChanges(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReorderBufferXidHasBaseSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReorderBufferXidHasBaseSnapshot
func F_ReorderBufferXidHasBaseSnapshot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SnapBuildSnapDecRefcount github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SnapBuildSnapDecRefcount
func F_SnapBuildSnapDecRefcount(m *base.Module, l0 int32)
//go:linkname F_SnapBuildInitialSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SnapBuildInitialSnapshot
func F_SnapBuildInitialSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_ReplicationSlotNameForTablesync github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotNameForTablesync
func F_ReplicationSlotNameForTablesync(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_finish_sync_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_finish_sync_worker
func F_finish_sync_worker(m *base.Module)
//go:linkname F_FetchTableStates github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FetchTableStates
func F_FetchTableStates(m *base.Module, l0 int32) int32
//go:linkname F_AllTablesyncsReady github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AllTablesyncsReady
func F_AllTablesyncsReady(m *base.Module) int32
//go:linkname F_UpdateTwoPhaseState github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UpdateTwoPhaseState
func F_UpdateTwoPhaseState(m *base.Module, l0 int32)
//go:linkname F_ReplicationOriginNameForLogicalRep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReplicationOriginNameForLogicalRep
func F_ReplicationOriginNameForLogicalRep(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_maybe_reread_subscription github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_maybe_reread_subscription
func F_maybe_reread_subscription(m *base.Module)
//go:linkname F_subxact_info_read github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_subxact_info_read
func F_subxact_info_read(m *base.Module, l0 int32, l1 int32)
//go:linkname F_subxact_info_write github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_subxact_info_write
func F_subxact_info_write(m *base.Module, l0 int32, l1 int32)
//go:linkname F_apply_spooled_messages github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_apply_spooled_messages
func F_apply_spooled_messages(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_apply_handle_commit_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_apply_handle_commit_internal
func F_apply_handle_commit_internal(m *base.Module, l0 int32)
//go:linkname F_handle_streamed_transaction github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_handle_streamed_transaction
func F_handle_streamed_transaction(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_should_apply_changes_for_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_should_apply_changes_for_rel
func F_should_apply_changes_for_rel(m *base.Module, l0 int32) int32
//go:linkname F_create_edata_for_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_create_edata_for_relation
func F_create_edata_for_relation(m *base.Module, l0 int32) int32
//go:linkname F_apply_handle_tuple_routing github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_apply_handle_tuple_routing
func F_apply_handle_tuple_routing(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_check_relation_updatable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_relation_updatable
func F_check_relation_updatable(m *base.Module, l0 int32)
//go:linkname F_slot_modify_data github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_slot_modify_data
func F_slot_modify_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_apply_handle_delete_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_apply_handle_delete_internal
func F_apply_handle_delete_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_stream_open_and_write_change github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_stream_open_and_write_change
func F_stream_open_and_write_change(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_clear_subscription_skip_lsn github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clear_subscription_skip_lsn
func F_clear_subscription_skip_lsn(m *base.Module, l0 int64)
//go:linkname F_apply_handle_prepare_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_apply_handle_prepare_internal
func F_apply_handle_prepare_internal(m *base.Module, l0 int32)
//go:linkname F_store_flush_position github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_store_flush_position
func F_store_flush_position(m *base.Module, l0 int64, l1 int64)
//go:linkname F_start_apply github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_start_apply
func F_start_apply(m *base.Module, l0 int64)
//go:linkname F_SetupApplyOrSyncWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SetupApplyOrSyncWorker
func F_SetupApplyOrSyncWorker(m *base.Module, l0 int32)
//go:linkname F_LogicalRepWorkersWakeupAtCommit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LogicalRepWorkersWakeupAtCommit
func F_LogicalRepWorkersWakeupAtCommit(m *base.Module, l0 int32)
//go:linkname F_replication_yylex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_replication_yylex
func F_replication_yylex(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replication_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_replication_yyensure_buffer_stack
func F_replication_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_replication_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_replication_yyerror
func F_replication_yyerror(m *base.Module, l0 int32)
//go:linkname F_yy_fatal_error_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_yy_fatal_error_3
func F_yy_fatal_error_3(m *base.Module, l0 int32)
//go:linkname F_replication_scanner_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_replication_scanner_finish
func F_replication_scanner_finish(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReplicationSlotRelease
func F_ReplicationSlotRelease(m *base.Module)
//go:linkname F_ReplicationSlotValidateNameInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplicationSlotValidateNameInternal
func F_ReplicationSlotValidateNameInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ReplicationSlotCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReplicationSlotCreate
func F_ReplicationSlotCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_SaveSlotToPath github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SaveSlotToPath
func F_SaveSlotToPath(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReplicationSlotAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplicationSlotAcquire
func F_ReplicationSlotAcquire(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReplicationSlotsComputeRequiredXmin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotsComputeRequiredXmin
func F_ReplicationSlotsComputeRequiredXmin(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotsComputeRequiredLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReplicationSlotsComputeRequiredLSN
func F_ReplicationSlotsComputeRequiredLSN(m *base.Module)
//go:linkname F_ReplicationSlotDrop github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotDrop
func F_ReplicationSlotDrop(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReplicationSlotSave github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotSave
func F_ReplicationSlotSave(m *base.Module)
//go:linkname F_ReplicationSlotPersist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplicationSlotPersist
func F_ReplicationSlotPersist(m *base.Module)
//go:linkname F_ReplicationSlotsDropDBSlots github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotsDropDBSlots
func F_ReplicationSlotsDropDBSlots(m *base.Module, l0 int32)
//go:linkname F_CheckSlotRequirements github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckSlotRequirements
func F_CheckSlotRequirements(m *base.Module)
//go:linkname F_CheckSlotPermissions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckSlotPermissions
func F_CheckSlotPermissions(m *base.Module)
//go:linkname F_ReplicationSlotReserveWal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReplicationSlotReserveWal
func F_ReplicationSlotReserveWal(m *base.Module)
//go:linkname F_InvalidateObsoleteReplicationSlots github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InvalidateObsoleteReplicationSlots
func F_InvalidateObsoleteReplicationSlots(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_StandbySlotsHaveCaughtup github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StandbySlotsHaveCaughtup
func F_StandbySlotsHaveCaughtup(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_WaitForStandbyConfirmation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WaitForStandbyConfirmation
func F_WaitForStandbyConfirmation(m *base.Module, l0 int64)
//go:linkname F_copy_replication_slot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_copy_replication_slot
func F_copy_replication_slot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SyncRepInitConfig github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SyncRepInitConfig
func F_SyncRepInitConfig(m *base.Module)
//go:linkname F_yy_fatal_error_4 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_yy_fatal_error_4
func F_yy_fatal_error_4(m *base.Module, l0 int32)
//go:linkname F_XLogWalRcvFlush github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogWalRcvFlush
func F_XLogWalRcvFlush(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetWalRcvFlushRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetWalRcvFlushRecPtr
func F_GetWalRcvFlushRecPtr(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_WalSndLoop github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WalSndLoop
func F_WalSndLoop(m *base.Module, l0 int32)
//go:linkname F_ProcessRepliesIfAny github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessRepliesIfAny
func F_ProcessRepliesIfAny(m *base.Module)
//go:linkname F_WalSndKeepalive github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WalSndKeepalive
func F_WalSndKeepalive(m *base.Module, l0 int32, l1 int64)
//go:linkname F_WalSndWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WalSndWait
func F_WalSndWait(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_WalSndShutdown github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WalSndShutdown
func F_WalSndShutdown(m *base.Module)
//go:linkname F_setRuleCheckAsUser_Query github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_setRuleCheckAsUser_Query
func F_setRuleCheckAsUser_Query(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AcquireRewriteLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AcquireRewriteLocks
func F_AcquireRewriteLocks(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_expand_generated_columns_in_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_expand_generated_columns_in_expr
func F_expand_generated_columns_in_expr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_generation_expression github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_generation_expression
func F_build_generation_expression(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_QueryRewrite github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_QueryRewrite
func F_QueryRewrite(m *base.Module, l0 int32) int32
//go:linkname F_locate_agg_of_level github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_locate_agg_of_level
func F_locate_agg_of_level(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ChangeVarNodes_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ChangeVarNodes_walker
func F_ChangeVarNodes_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ChangeVarNodes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ChangeVarNodes
func F_ChangeVarNodes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ChangeVarNodesWalkExpression github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ChangeVarNodesWalkExpression
func F_ChangeVarNodesWalkExpression(m *base.Module, l0 int32, l1 int32)
//go:linkname F_remove_nulling_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_remove_nulling_relids
func F_remove_nulling_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_map_variable_attnos github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_map_variable_attnos
func F_map_variable_attnos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_text_to_stavalues github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_text_to_stavalues
func F_text_to_stavalues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_set_stats_slot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_set_stats_slot
func F_set_stats_slot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_statext_mcv_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_statext_mcv_deserialize
func F_statext_mcv_deserialize(m *base.Module, l0 int32) int32
//go:linkname F_statext_ndistinct_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_statext_ndistinct_deserialize
func F_statext_ndistinct_deserialize(m *base.Module, l0 int32) int32
//go:linkname F_stats_check_required_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_stats_check_required_arg
func F_stats_check_required_arg(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_stats_check_arg_array github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_stats_check_arg_array
func F_stats_check_arg_array(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_stats_check_arg_pair github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_stats_check_arg_pair
func F_stats_check_arg_pair(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_stats_fill_fcinfo_from_arg_pairs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_stats_fill_fcinfo_from_arg_pairs
func F_stats_fill_fcinfo_from_arg_pairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgaio_io_update_state github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgaio_io_update_state
func F_pgaio_io_update_state(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgaio_closing_fd github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgaio_closing_fd
func F_pgaio_closing_fd(m *base.Module, l0 int32)
//go:linkname F_read_stream_begin_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_read_stream_begin_relation
func F_read_stream_begin_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_read_stream_next_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_read_stream_next_buffer
func F_read_stream_next_buffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_read_stream_end github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_stream_end
func F_read_stream_end(m *base.Module, l0 int32)
//go:linkname F_BufTableHashCode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BufTableHashCode
func F_BufTableHashCode(m *base.Module, l0 int32) int32
//go:linkname F_BufTableLookup github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BufTableLookup
func F_BufTableLookup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BufTableInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufTableInsert
func F_BufTableInsert(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_UnpinBufferNoOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnpinBufferNoOwner
func F_UnpinBufferNoOwner(m *base.Module, l0 int32)
//go:linkname F_PinBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PinBuffer
func F_PinBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LockBufHdr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LockBufHdr
func F_LockBufHdr(m *base.Module, l0 int32) int32
//go:linkname F_ReadBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReadBuffer
func F_ReadBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReadBufferExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadBufferExtended
func F_ReadBufferExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ExtendBufferedRel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExtendBufferedRel
func F_ExtendBufferedRel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GetVictimBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetVictimBuffer
func F_GetVictimBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_UnpinBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_UnpinBuffer
func F_UnpinBuffer(m *base.Module, l0 int32)
//go:linkname F_ZeroAndLockBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ZeroAndLockBuffer
func F_ZeroAndLockBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_StartReadBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_StartReadBuffer
func F_StartReadBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_WaitReadBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WaitReadBuffers
func F_WaitReadBuffers(m *base.Module, l0 int32)
//go:linkname F_TerminateBufferIO github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TerminateBufferIO
func F_TerminateBufferIO(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_AsyncReadBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AsyncReadBuffers
func F_AsyncReadBuffers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReleaseAndReadBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReleaseAndReadBuffer
func F_ReleaseAndReadBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FlushBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FlushBuffer
func F_FlushBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RelationGetNumberOfBlocksInFork github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetNumberOfBlocksInFork
func F_RelationGetNumberOfBlocksInFork(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BufferGetLSNAtomic github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufferGetLSNAtomic
func F_BufferGetLSNAtomic(m *base.Module, l0 int32) int64
//go:linkname F_InvalidateBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_InvalidateBuffer
func F_InvalidateBuffer(m *base.Module, l0 int32)
//go:linkname F_FlushRelationBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FlushRelationBuffers
func F_FlushRelationBuffers(m *base.Module, l0 int32)
//go:linkname F_FlushOneBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FlushOneBuffer
func F_FlushOneBuffer(m *base.Module, l0 int32)
//go:linkname F_LockBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockBuffer
func F_LockBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_IncrBufferRefCount github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_IncrBufferRefCount
func F_IncrBufferRefCount(m *base.Module, l0 int32)
//go:linkname F_MarkBufferDirtyHint github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MarkBufferDirtyHint
func F_MarkBufferDirtyHint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LockBufferForCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockBufferForCleanup
func F_LockBufferForCleanup(m *base.Module, l0 int32)
//go:linkname F_HoldingBufferPinThatDelaysRecovery github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HoldingBufferPinThatDelaysRecovery
func F_HoldingBufferPinThatDelaysRecovery(m *base.Module) int32
//go:linkname F_sort_pending_writebacks github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sort_pending_writebacks
func F_sort_pending_writebacks(m *base.Module, l0 int32, l1 int32)
//go:linkname F_StrategyFreeBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StrategyFreeBuffer
func F_StrategyFreeBuffer(m *base.Module, l0 int32)
//go:linkname F_GetAccessStrategyWithSize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetAccessStrategyWithSize
func F_GetAccessStrategyWithSize(m *base.Module, l0 int32) int32
//go:linkname F_IOContextForStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IOContextForStrategy
func F_IOContextForStrategy(m *base.Module, l0 int32) int32
//go:linkname F_LocalBufferAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LocalBufferAlloc
func F_LocalBufferAlloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_FlushLocalBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FlushLocalBuffer
func F_FlushLocalBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_InvalidateLocalBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InvalidateLocalBuffer
func F_InvalidateLocalBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BufFileCreateFileSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BufFileCreateFileSet
func F_BufFileCreateFileSet(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BufFileDeleteFileSet github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BufFileDeleteFileSet
func F_BufFileDeleteFileSet(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_BufFileClose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BufFileClose
func F_BufFileClose(m *base.Module, l0 int32)
//go:linkname F_BufFileReadExact github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BufFileReadExact
func F_BufFileReadExact(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_BufFileWrite github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufFileWrite
func F_BufFileWrite(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_BufFileSeek github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BufFileSeek
func F_BufFileSeek(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_BufFileSeekBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufFileSeekBlock
func F_BufFileSeekBlock(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_fsync_fname_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fsync_fname_ext
func F_fsync_fname_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CloseTransientFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CloseTransientFile
func F_CloseTransientFile(m *base.Module, l0 int32) int32
//go:linkname F_OpenTransientFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_OpenTransientFile
func F_OpenTransientFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FileClose github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileClose
func F_FileClose(m *base.Module, l0 int32)
//go:linkname F_BasicOpenFile github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BasicOpenFile
func F_BasicOpenFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LruDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LruDelete
func F_LruDelete(m *base.Module, l0 int32)
//go:linkname F_PathNameOpenFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PathNameOpenFile
func F_PathNameOpenFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_PathNameOpenFilePerm github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PathNameOpenFilePerm
func F_PathNameOpenFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_AllocateDir github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AllocateDir
func F_AllocateDir(m *base.Module, l0 int32) int32
//go:linkname F_FreeDir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeDir
func F_FreeDir(m *base.Module, l0 int32)
//go:linkname F_PathNameDeleteTemporaryFile github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PathNameDeleteTemporaryFile
func F_PathNameDeleteTemporaryFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_OpenTemporaryFile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_OpenTemporaryFile
func F_OpenTemporaryFile(m *base.Module, l0 int32) int32
//go:linkname F_TempTablespacePath github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TempTablespacePath
func F_TempTablespacePath(m *base.Module, l0 int32, l1 int32)
//go:linkname F_FileAccess github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileAccess
func F_FileAccess(m *base.Module, l0 int32) int32
//go:linkname F_FileSize github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FileSize
func F_FileSize(m *base.Module, l0 int32) int64
//go:linkname F_FileTruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileTruncate
func F_FileTruncate(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_AllocateFile github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AllocateFile
func F_AllocateFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FreeFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FreeFile
func F_FreeFile(m *base.Module, l0 int32) int32
//go:linkname F_ReadDir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadDir
func F_ReadDir(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ClosePipeStream github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ClosePipeStream
func F_ClosePipeStream(m *base.Module, l0 int32) int32
//go:linkname F_FileSetInit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FileSetInit
func F_FileSetInit(m *base.Module, l0 int32)
//go:linkname F_FileSetDeleteAll github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileSetDeleteAll
func F_FileSetDeleteAll(m *base.Module, l0 int32)
//go:linkname F_ResetUnloggedRelationsInTablespaceDir github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResetUnloggedRelationsInTablespaceDir
func F_ResetUnloggedRelationsInTablespaceDir(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetPageWithFreeSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetPageWithFreeSpace
func F_GetPageWithFreeSpace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fsm_search github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fsm_search
func F_fsm_search(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fsm_readbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fsm_readbuf
func F_fsm_readbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RecordPageWithFreeSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RecordPageWithFreeSpace
func F_RecordPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_fsm_vacuum_page github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fsm_vacuum_page
func F_fsm_vacuum_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_RecordFreeIndexPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RecordFreeIndexPage
func F_RecordFreeIndexPage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BarrierArriveAndWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BarrierArriveAndWait
func F_BarrierArriveAndWait(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BarrierAttach github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BarrierAttach
func F_BarrierAttach(m *base.Module, l0 int32) int32
//go:linkname F_BarrierDetach github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BarrierDetach
func F_BarrierDetach(m *base.Module, l0 int32)
//go:linkname F_dsm_create github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dsm_create
func F_dsm_create(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsm_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsm_attach
func F_dsm_attach(m *base.Module, l0 int32) int32
//go:linkname F_dsm_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsm_detach
func F_dsm_detach(m *base.Module, l0 int32)
//go:linkname F_on_dsm_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_on_dsm_detach
func F_on_dsm_detach(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dsm_impl_op github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dsm_impl_op
func F_dsm_impl_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_proc_exit_prepare github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_proc_exit_prepare
func F_proc_exit_prepare(m *base.Module, l0 int32)
//go:linkname F_before_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_before_shmem_exit
func F_before_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_on_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_on_shmem_exit
func F_on_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_cancel_before_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cancel_before_shmem_exit
func F_cancel_before_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_WaitLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_WaitLatch
func F_WaitLatch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SetLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SetLatch
func F_SetLatch(m *base.Module, l0 int32)
//go:linkname F_SendPostmasterSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SendPostmasterSignal
func F_SendPostmasterSignal(m *base.Module, l0 int32)
//go:linkname F_ProcArrayAdd github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ProcArrayAdd
func F_ProcArrayAdd(m *base.Module, l0 int32)
//go:linkname F_KnownAssignedXidsRemove github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_KnownAssignedXidsRemove
func F_KnownAssignedXidsRemove(m *base.Module, l0 int32)
//go:linkname F_TransactionIdIsInProgress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TransactionIdIsInProgress
func F_TransactionIdIsInProgress(m *base.Module, l0 int32) int32
//go:linkname F_GlobalVisHorizonKindForRel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GlobalVisHorizonKindForRel
func F_GlobalVisHorizonKindForRel(m *base.Module, l0 int32) int32
//go:linkname F_GetOldestTransactionIdConsideredRunning github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetOldestTransactionIdConsideredRunning
func F_GetOldestTransactionIdConsideredRunning(m *base.Module) int32
//go:linkname F_GetSnapshotData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetSnapshotData
func F_GetSnapshotData(m *base.Module, l0 int32) int32
//go:linkname F_CountOtherDBBackends github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CountOtherDBBackends
func F_CountOtherDBBackends(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GlobalVisTestIsRemovableXid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GlobalVisTestIsRemovableXid
func F_GlobalVisTestIsRemovableXid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GlobalVisCheckRemovableXid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GlobalVisCheckRemovableXid
func F_GlobalVisCheckRemovableXid(m *base.Module, l0 int32) int32
//go:linkname F_ProcSignalInit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcSignalInit
func F_ProcSignalInit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SendProcSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SendProcSignal
func F_SendProcSignal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EmitProcSignalBarrier github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EmitProcSignalBarrier
func F_EmitProcSignalBarrier(m *base.Module) int64
//go:linkname F_ProcessProcSignalBarrier github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcessProcSignalBarrier
func F_ProcessProcSignalBarrier(m *base.Module)
//go:linkname F_shm_mq_set_receiver github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shm_mq_set_receiver
func F_shm_mq_set_receiver(m *base.Module, l0 int32, l1 int32)
//go:linkname F_shm_mq_set_sender github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_shm_mq_set_sender
func F_shm_mq_set_sender(m *base.Module, l0 int32, l1 int32)
//go:linkname F_shm_mq_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_mq_attach
func F_shm_mq_attach(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_shm_mq_wait_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_mq_wait_internal
func F_shm_mq_wait_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_shm_mq_receive_bytes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_mq_receive_bytes
func F_shm_mq_receive_bytes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_shm_mq_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_mq_detach
func F_shm_mq_detach(m *base.Module, l0 int32)
//go:linkname F_shm_toc_allocate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_shm_toc_allocate
func F_shm_toc_allocate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_shm_toc_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_toc_insert
func F_shm_toc_insert(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_shm_toc_lookup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_shm_toc_lookup
func F_shm_toc_lookup(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_shm_toc_estimate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_shm_toc_estimate
func F_shm_toc_estimate(m *base.Module, l0 int32) int32
//go:linkname F_add_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_size
func F_add_size(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_mul_size github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_mul_size
func F_mul_size(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SendSharedInvalidMessages github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SendSharedInvalidMessages
func F_SendSharedInvalidMessages(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ProcessCatchupInterrupt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessCatchupInterrupt
func F_ProcessCatchupInterrupt(m *base.Module)
//go:linkname F_SharedInvalBackendInit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SharedInvalBackendInit
func F_SharedInvalBackendInit(m *base.Module, l0 int32)
//go:linkname F_SICleanupQueue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SICleanupQueue
func F_SICleanupQueue(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ShutdownRecoveryTransactionEnvironment github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ShutdownRecoveryTransactionEnvironment
func F_ShutdownRecoveryTransactionEnvironment(m *base.Module)
//go:linkname F_StandbyReleaseXidEntryLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StandbyReleaseXidEntryLocks
func F_StandbyReleaseXidEntryLocks(m *base.Module, l0 int32)
//go:linkname F_ResolveRecoveryConflictWithSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResolveRecoveryConflictWithSnapshot
func F_ResolveRecoveryConflictWithSnapshot(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ResolveRecoveryConflictWithVirtualXIDs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResolveRecoveryConflictWithVirtualXIDs
func F_ResolveRecoveryConflictWithVirtualXIDs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_WaitEventSetWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WaitEventSetWait
func F_WaitEventSetWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_inv_open github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inv_open
func F_inv_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_inv_seek github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_inv_seek
func F_inv_seek(m *base.Module, l0 int32, l1 int64, l2 int32) int64
//go:linkname F_inv_read github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inv_read
func F_inv_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_inv_write github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_inv_write
func F_inv_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ConditionVariableInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ConditionVariableInit
func F_ConditionVariableInit(m *base.Module, l0 int32)
//go:linkname F_ConditionVariablePrepareToSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ConditionVariablePrepareToSleep
func F_ConditionVariablePrepareToSleep(m *base.Module, l0 int32)
//go:linkname F_ConditionVariableCancelSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ConditionVariableCancelSleep
func F_ConditionVariableCancelSleep(m *base.Module)
//go:linkname F_ConditionVariableSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ConditionVariableSleep
func F_ConditionVariableSleep(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ConditionVariableTimedSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ConditionVariableTimedSleep
func F_ConditionVariableTimedSleep(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ConditionVariableBroadcast github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ConditionVariableBroadcast
func F_ConditionVariableBroadcast(m *base.Module, l0 int32)
//go:linkname F_DeadLockCheckRecurse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DeadLockCheckRecurse
func F_DeadLockCheckRecurse(m *base.Module, l0 int32) int32
//go:linkname F_UnlockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockRelationOid
func F_UnlockRelationOid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LockRelationIdForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LockRelationIdForSession
func F_LockRelationIdForSession(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnlockRelationIdForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockRelationIdForSession
func F_UnlockRelationIdForSession(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LockRelationForExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LockRelationForExtension
func F_LockRelationForExtension(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnlockRelationForExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockRelationForExtension
func F_UnlockRelationForExtension(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnlockPage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UnlockPage
func F_UnlockPage(m *base.Module, l0 int32)
//go:linkname F_LockTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockTuple
func F_LockTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_UnlockTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockTuple
func F_UnlockTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XactLockTableWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XactLockTableWait
func F_XactLockTableWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_SpeculativeInsertionWait github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SpeculativeInsertionWait
func F_SpeculativeInsertionWait(m *base.Module, l0 int32, l1 int32)
//go:linkname F_WaitForLockers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WaitForLockers
func F_WaitForLockers(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnlockSharedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_UnlockSharedObject
func F_UnlockSharedObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockApplyTransactionForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockApplyTransactionForSession
func F_LockApplyTransactionForSession(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_UnlockApplyTransactionForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_UnlockApplyTransactionForSession
func F_UnlockApplyTransactionForSession(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_DescribeLockTag github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DescribeLockTag
func F_DescribeLockTag(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockAcquire
func F_LockAcquire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_LockAcquireExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockAcquireExtended
func F_LockAcquireExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_AbortStrongLockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AbortStrongLockAcquire
func F_AbortStrongLockAcquire(m *base.Module)
//go:linkname F_RemoveFromWaitQueue github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RemoveFromWaitQueue
func F_RemoveFromWaitQueue(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CleanUpLock github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CleanUpLock
func F_CleanUpLock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_LockReleaseAll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LockReleaseAll
func F_LockReleaseAll(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetLockConflicts github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetLockConflicts
func F_GetLockConflicts(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_VirtualXactLock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_VirtualXactLock
func F_VirtualXactLock(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LWLockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockAcquire
func F_LWLockAcquire(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LWLockConditionalAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockConditionalAcquire
func F_LWLockConditionalAcquire(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LWLockDisownInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LWLockDisownInternal
func F_LWLockDisownInternal(m *base.Module, l0 int32) int32
//go:linkname F_LWLockReleaseInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockReleaseInternal
func F_LWLockReleaseInternal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LWLockReleaseAll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockReleaseAll
func F_LWLockReleaseAll(m *base.Module)
//go:linkname F_GetSafeSnapshotBlockingPids github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetSafeSnapshotBlockingPids
func F_GetSafeSnapshotBlockingPids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReleasePredicateLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReleasePredicateLocks
func F_ReleasePredicateLocks(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PredicateLockRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PredicateLockRelation
func F_PredicateLockRelation(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PredicateLockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PredicateLockAcquire
func F_PredicateLockAcquire(m *base.Module, l0 int32)
//go:linkname F_CreatePredicateLock github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreatePredicateLock
func F_CreatePredicateLock(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_TransferPredicateLocksToNewTarget github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TransferPredicateLocksToNewTarget
func F_TransferPredicateLocksToNewTarget(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CheckTargetForConflictsIn github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckTargetForConflictsIn
func F_CheckTargetForConflictsIn(m *base.Module, l0 int32)
//go:linkname F_CheckTableForSerializableConflictIn github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CheckTableForSerializableConflictIn
func F_CheckTableForSerializableConflictIn(m *base.Module, l0 int32)
//go:linkname F_PGProcShmemSize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PGProcShmemSize
func F_PGProcShmemSize(m *base.Module) int32
//go:linkname F_ProcLockWakeup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcLockWakeup
func F_ProcLockWakeup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ProcWaitForSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcWaitForSignal
func F_ProcWaitForSignal(m *base.Module, l0 int32)
//go:linkname F_s_lock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_s_lock
func F_s_lock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_perform_spin_delay github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_perform_spin_delay
func F_perform_spin_delay(m *base.Module, l0 int32)
//go:linkname F_PageInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PageInit
func F_PageInit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PageIsVerified github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PageIsVerified
func F_PageIsVerified(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_PageAddItemExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PageAddItemExtended
func F_PageAddItemExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_PageGetTempPage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageGetTempPage
func F_PageGetTempPage(m *base.Module, l0 int32) int32
//go:linkname F_PageGetTempPageCopySpecial github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageGetTempPageCopySpecial
func F_PageGetTempPageCopySpecial(m *base.Module, l0 int32) int32
//go:linkname F_PageRestoreTempPage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PageRestoreTempPage
func F_PageRestoreTempPage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PageIndexTupleDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageIndexTupleDelete
func F_PageIndexTupleDelete(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PageIndexMultiDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PageIndexMultiDelete
func F_PageIndexMultiDelete(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_smgr_bulk_start_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgr_bulk_start_rel
func F_smgr_bulk_start_rel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgr_bulk_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_smgr_bulk_finish
func F_smgr_bulk_finish(m *base.Module, l0 int32)
//go:linkname F_smgr_bulk_write github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_smgr_bulk_write
func F_smgr_bulk_write(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_smgr_bulk_get_buf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgr_bulk_get_buf
func F_smgr_bulk_get_buf(m *base.Module, l0 int32) int32
//go:linkname F_register_dirty_segment github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_register_dirty_segment
func F_register_dirty_segment(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F__mdfd_getseg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__mdfd_getseg
func F__mdfd_getseg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F__mdfd_openseg github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__mdfd_openseg
func F__mdfd_openseg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_smgrclose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrclose
func F_smgrclose(m *base.Module, l0 int32)
//go:linkname F_smgrreleaseall github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrreleaseall
func F_smgrreleaseall(m *base.Module)
//go:linkname F_smgrexists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrexists
func F_smgrexists(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgrdounlinkall github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_smgrdounlinkall
func F_smgrdounlinkall(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_smgrprefetch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_smgrprefetch
func F_smgrprefetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_EndCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EndCommand
func F_EndCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_EndReplicationCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EndReplicationCommand
func F_EndReplicationCommand(m *base.Module, l0 int32)
//go:linkname F_NullCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_NullCommand
func F_NullCommand(m *base.Module, l0 int32)
//go:linkname F_ReadyForQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReadyForQuery
func F_ReadyForQuery(m *base.Module, l0 int32)
//go:linkname F_errdetail_recovery_conflict github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errdetail_recovery_conflict
func F_errdetail_recovery_conflict(m *base.Module, l0 int32)
//go:linkname F_ShowUsage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ShowUsage
func F_ShowUsage(m *base.Module, l0 int32)
//go:linkname F_pg_rewrite_query github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_rewrite_query
func F_pg_rewrite_query(m *base.Module, l0 int32) int32
//go:linkname F_pg_analyze_and_rewrite_varparams github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_analyze_and_rewrite_varparams
func F_pg_analyze_and_rewrite_varparams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_plan_query github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_plan_query
func F_pg_plan_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_plan_queries github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_plan_queries
func F_pg_plan_queries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_check_log_duration github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_log_duration
func F_check_log_duration(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_start_xact_command github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_start_xact_command
func F_start_xact_command(m *base.Module)
//go:linkname F_errdetail_abort github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_errdetail_abort
func F_errdetail_abort(m *base.Module)
//go:linkname F_FetchStatementTargetList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FetchStatementTargetList
func F_FetchStatementTargetList(m *base.Module, l0 int32) int32
//go:linkname F_PortalStart github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PortalStart
func F_PortalStart(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_PortalSetResultFormat github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PortalSetResultFormat
func F_PortalSetResultFormat(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PortalRun github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PortalRun
func F_PortalRun(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_FillPortalStore github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FillPortalStore
func F_FillPortalStore(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PortalRunSelect github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PortalRunSelect
func F_PortalRunSelect(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64
//go:linkname F_RunFromStore github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RunFromStore
func F_RunFromStore(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int64
//go:linkname F_DoPortalRewind github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DoPortalRewind
func F_DoPortalRewind(m *base.Module, l0 int32)
//go:linkname F_EnsurePortalSnapshotExists github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EnsurePortalSnapshotExists
func F_EnsurePortalSnapshotExists(m *base.Module)
//go:linkname F_PreventCommandIfReadOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreventCommandIfReadOnly
func F_PreventCommandIfReadOnly(m *base.Module, l0 int32)
//go:linkname F_PreventCommandIfParallelMode github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreventCommandIfParallelMode
func F_PreventCommandIfParallelMode(m *base.Module, l0 int32)
//go:linkname F_CreateCommandTag github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateCommandTag
func F_CreateCommandTag(m *base.Module, l0 int32) int32
//go:linkname F_CheckRestrictedOperation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckRestrictedOperation
func F_CheckRestrictedOperation(m *base.Module, l0 int32)
//go:linkname F_ExecDropStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecDropStmt
func F_ExecDropStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ProcessUtilitySlow github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessUtilitySlow
func F_ProcessUtilitySlow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_makeCompoundFlags github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeCompoundFlags
func F_makeCompoundFlags(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_NormalizeSubWord github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_NormalizeSubWord
func F_NormalizeSubWord(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_tsvector github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_tsvector
func F_make_tsvector(m *base.Module, l0 int32) int32
//go:linkname F_t_isalnum_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_t_isalnum_with_len
func F_t_isalnum_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_t_isalnum_cstr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_t_isalnum_cstr
func F_t_isalnum_cstr(m *base.Module, l0 int32) int32
//go:linkname F_tsearch_readline_end github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tsearch_readline_end
func F_tsearch_readline_end(m *base.Module, l0 int32)
//go:linkname F_parsetext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parsetext
func F_parsetext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_hlparsetext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hlparsetext
func F_hlparsetext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_get_tsearch_config_filename github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_tsearch_config_filename
func F_get_tsearch_config_filename(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tt_setup_firstcall github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tt_setup_firstcall
func F_tt_setup_firstcall(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tt_process_call github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tt_process_call
func F_tt_process_call(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_beinit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_beinit
func F_pgstat_beinit(m *base.Module)
//go:linkname F_pgstat_bestart_initial github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_bestart_initial
func F_pgstat_bestart_initial(m *base.Module)
//go:linkname F_pgstat_bestart_final github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_bestart_final
func F_pgstat_bestart_final(m *base.Module)
//go:linkname F_pgstat_report_activity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_report_activity
func F_pgstat_report_activity(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_get_beentry_by_proc_number github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_get_beentry_by_proc_number
func F_pgstat_get_beentry_by_proc_number(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_report_stat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_report_stat
func F_pgstat_report_stat(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_reset github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_reset
func F_pgstat_reset(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_pgstat_reset_of_kind github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_reset_of_kind
func F_pgstat_reset_of_kind(m *base.Module, l0 int32)
//go:linkname F_pgstat_fetch_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_fetch_entry
func F_pgstat_fetch_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pgstat_prep_pending_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_prep_pending_entry
func F_pgstat_prep_pending_entry(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pgstat_count_backend_io_op github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_count_backend_io_op
func F_pgstat_count_backend_io_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64)
//go:linkname F_pgstat_fetch_stat_backend_by_pid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_fetch_stat_backend_by_pid
func F_pgstat_fetch_stat_backend_by_pid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_flush_backend github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_flush_backend
func F_pgstat_flush_backend(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_fetch_stat_bgwriter github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_fetch_stat_bgwriter
func F_pgstat_fetch_stat_bgwriter(m *base.Module) int32
//go:linkname F_pgstat_fetch_stat_checkpointer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_fetch_stat_checkpointer
func F_pgstat_fetch_stat_checkpointer(m *base.Module) int32
//go:linkname F_pgstat_fetch_stat_dbentry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_fetch_stat_dbentry
func F_pgstat_fetch_stat_dbentry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_init_function_usage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_init_function_usage
func F_pgstat_init_function_usage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_fetch_stat_funcentry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_fetch_stat_funcentry
func F_pgstat_fetch_stat_funcentry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_io_flush_cb github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_io_flush_cb
func F_pgstat_io_flush_cb(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_assoc_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_assoc_relation
func F_pgstat_assoc_relation(m *base.Module, l0 int32)
//go:linkname F_pgstat_drop_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_drop_relation
func F_pgstat_drop_relation(m *base.Module, l0 int32)
//go:linkname F_pgstat_count_heap_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_count_heap_delete
func F_pgstat_count_heap_delete(m *base.Module, l0 int32)
//go:linkname F_pgstat_fetch_stat_tabentry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_fetch_stat_tabentry
func F_pgstat_fetch_stat_tabentry(m *base.Module, l0 int32) int32
//go:linkname F_find_tabstat_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_tabstat_entry
func F_find_tabstat_entry(m *base.Module, l0 int32) int32
//go:linkname F_StatsShmemSize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_StatsShmemSize
func F_StatsShmemSize(m *base.Module) int32
//go:linkname F_pgstat_get_entry_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_get_entry_ref
func F_pgstat_get_entry_ref(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32
//go:linkname F_pgstat_lock_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_lock_entry
func F_pgstat_lock_entry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_unlock_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_unlock_entry
func F_pgstat_unlock_entry(m *base.Module, l0 int32)
//go:linkname F_pgstat_drop_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_drop_entry
func F_pgstat_drop_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pgstat_wal_flush_cb github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_wal_flush_cb
func F_pgstat_wal_flush_cb(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_drop_transactional github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_drop_transactional
func F_pgstat_drop_transactional(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_aclmask github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_aclmask
func F_aclmask(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int64
//go:linkname F_acldefault github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_acldefault
func F_acldefault(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_aclnewowner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_aclnewowner
func F_aclnewowner(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_has_privs_of_role github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_has_privs_of_role
func F_has_privs_of_role(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_aclmembers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_aclmembers
func F_aclmembers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_convert_any_priv_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_any_priv_string
func F_convert_any_priv_string(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_get_role_oid_or_public github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_role_oid_or_public
func F_get_role_oid_or_public(m *base.Module, l0 int32) int32
//go:linkname F_pg_role_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_role_aclcheck
func F_pg_role_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_get_role_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_role_oid
func F_get_role_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_initialize_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initialize_acl
func F_initialize_acl(m *base.Module)
//go:linkname F_roles_list_append github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_roles_list_append
func F_roles_list_append(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_is_member_of_role_nosuper github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_is_member_of_role_nosuper
func F_is_member_of_role_nosuper(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_is_admin_of_role github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_is_admin_of_role
func F_is_admin_of_role(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rolespec_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_rolespec_tuple
func F_get_rolespec_tuple(m *base.Module, l0 int32) int32
//go:linkname F_get_rolespec_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rolespec_name
func F_get_rolespec_name(m *base.Module, l0 int32) int32
//go:linkname F_check_rolespec_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_rolespec_name
func F_check_rolespec_name(m *base.Module, l0 int32)
//go:linkname F_indexam_property github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_indexam_property
func F_indexam_property(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_expand_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_expand_array
func F_expand_array(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_deconstruct_expanded_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_expanded_array
func F_deconstruct_expanded_array(m *base.Module, l0 int32)
//go:linkname F_CopyArrayEls github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CopyArrayEls
func F_CopyArrayEls(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_ArrayCastAndSet github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ArrayCastAndSet
func F_ArrayCastAndSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_array_seek github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_seek
func F_array_seek(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_array_get_slice github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_array_get_slice
func F_array_get_slice(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_deconstruct_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_array
func F_deconstruct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_construct_array github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_construct_array
func F_construct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_deconstruct_array_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_array_builtin
func F_deconstruct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_array_contains_nulls github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_contains_nulls
func F_array_contains_nulls(m *base.Module, l0 int32) int32
//go:linkname F_array_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_cmp
func F_array_cmp(m *base.Module, l0 int32) int32
//go:linkname F_array_iterate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_iterate
func F_array_iterate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_initArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_initArrayResult
func F_initArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_accumArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_accumArrayResult
func F_accumArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeArrayResult
func F_makeArrayResult(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeMdArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeMdArrayResult
func F_makeMdArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_initArrayResultArr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initArrayResultArr
func F_initArrayResultArr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_accumArrayResultArr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_accumArrayResultArr
func F_accumArrayResultArr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_generate_subscripts github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generate_subscripts
func F_generate_subscripts(m *base.Module, l0 int32) int32
//go:linkname F_ArrayGetNItems github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ArrayGetNItems
func F_ArrayGetNItems(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ArrayCheckBounds github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ArrayCheckBounds
func F_ArrayCheckBounds(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_encode_to_ascii github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_encode_to_ascii
func F_encode_to_ascii(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_bool_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_bool_with_len
func F_parse_bool_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_cash_mul_float8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cash_mul_float8
func F_cash_mul_float8(m *base.Module, l0 int64, l1 float64) int64
//go:linkname F_append_num_word github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_append_num_word
func F_append_num_word(m *base.Module, l0 int32, l1 int64)
//go:linkname F_cryptohash_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cryptohash_internal
func F_cryptohash_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_anytime_typmod_check github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_anytime_typmod_check
func F_anytime_typmod_check(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_EncodeSpecialDate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_EncodeSpecialDate
func F_EncodeSpecialDate(m *base.Module, l0 int32, l1 int32)
//go:linkname F_time_part_common github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_time_part_common
func F_time_part_common(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_timetz_part_common github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_timetz_part_common
func F_timetz_part_common(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_date2j github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_date2j
func F_date2j(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_j2date github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_j2date
func F_j2date(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetCurrentDateTime github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetCurrentDateTime
func F_GetCurrentDateTime(m *base.Module, l0 int32)
//go:linkname F_GetCurrentTimeUsec github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetCurrentTimeUsec
func F_GetCurrentTimeUsec(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ParseDateTime github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ParseDateTime
func F_ParseDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_DecodeDate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DecodeDate
func F_DecodeDate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_DecodeTimeCommon github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DecodeTimeCommon
func F_DecodeTimeCommon(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_DecodeNumber github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DecodeNumber
func F_DecodeNumber(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_DecodeTimezoneAbbrev github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DecodeTimezoneAbbrev
func F_DecodeTimezoneAbbrev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_DetermineTimeZoneOffsetInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DetermineTimeZoneOffsetInternal
func F_DetermineTimeZoneOffsetInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DecodeTimezoneName github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DecodeTimezoneName
func F_DecodeTimezoneName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DateTimeParseError github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DateTimeParseError
func F_DateTimeParseError(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ParseISO8601Number github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ParseISO8601Number
func F_ParseISO8601Number(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_AppendSeconds github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AppendSeconds
func F_AppendSeconds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_EncodeDateTime github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_EncodeDateTime
func F_EncodeDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_datumCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_datumCopy
func F_datumCopy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datum_image_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_datum_image_eq
func F_datum_image_eq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_datumSerialize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_datumSerialize
func F_datumSerialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_calculate_tablespace_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_calculate_tablespace_size
func F_calculate_tablespace_size(m *base.Module, l0 int32) int64
//go:linkname F_calculate_relation_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_calculate_relation_size
func F_calculate_relation_size(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_domain_state_setup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_domain_state_setup
func F_domain_state_setup(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_domain_check_input github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_domain_check_input
func F_domain_check_input(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_check_safe_enum_use github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_safe_enum_use
func F_check_safe_enum_use(m *base.Module, l0 int32)
//go:linkname F_enum_cmp_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_enum_cmp_internal
func F_enum_cmp_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EOH_get_flat_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EOH_get_flat_size
func F_EOH_get_flat_size(m *base.Module, l0 int32) int32
//go:linkname F_EOH_flatten_into github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EOH_flatten_into
func F_EOH_flatten_into(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_DeleteExpandedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DeleteExpandedObject
func F_DeleteExpandedObject(m *base.Module, l0 int32)
//go:linkname F_expanded_record_get_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expanded_record_get_tuple
func F_expanded_record_get_tuple(m *base.Module, l0 int32) int32
//go:linkname F_deconstruct_expanded_record github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_deconstruct_expanded_record
func F_deconstruct_expanded_record(m *base.Module, l0 int32)
//go:linkname F_expanded_record_fetch_field github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expanded_record_fetch_field
func F_expanded_record_fetch_field(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_float_overflow_error github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float_overflow_error
func F_float_overflow_error(m *base.Module)
//go:linkname F_float_underflow_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_float_underflow_error
func F_float_underflow_error(m *base.Module)
//go:linkname F_float_zero_divide_error github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float_zero_divide_error
func F_float_zero_divide_error(m *base.Module)
//go:linkname F_init_degree_constants github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_init_degree_constants
func F_init_degree_constants(m *base.Module)
//go:linkname F_format_type_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_format_type_extended
func F_format_type_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_format_type_be github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_format_type_be
func F_format_type_be(m *base.Module, l0 int32) int32
//go:linkname F_format_type_with_typemod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_format_type_with_typemod
func F_format_type_with_typemod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_str_tolower github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_str_tolower
func F_str_tolower(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DCH_to_char github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DCH_to_char
func F_DCH_to_char(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_parse_format github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_format
func F_parse_format(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_do_to_timestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_do_to_timestamp
func F_do_to_timestamp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_NUM_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_NUM_cache
func F_NUM_cache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_NUM_processor github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_NUM_processor
func F_NUM_processor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_int_to_roman github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_int_to_roman
func F_int_to_roman(m *base.Module, l0 int32) int32
//go:linkname F_convert_and_check_filename github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_and_check_filename
func F_convert_and_check_filename(m *base.Module, l0 int32) int32
//go:linkname F_read_binary_file github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_binary_file
func F_read_binary_file(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32
//go:linkname F_pg_ls_dir_files github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_ls_dir_files
func F_pg_ls_dir_files(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_ls_tmpdir github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_ls_tmpdir
func F_pg_ls_tmpdir(m *base.Module, l0 int32, l1 int32)
//go:linkname F_path_decode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_path_decode
func F_path_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_box_ar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_box_ar
func F_box_ar(m *base.Module, l0 int32) float64
//go:linkname F_point_dt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_point_dt
func F_point_dt(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_float8_div github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float8_div
func F_float8_div(m *base.Module, l0 float64, l1 float64) float64
//go:linkname F_lseg_interpt_lseg github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lseg_interpt_lseg
func F_lseg_interpt_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lseg_interpt_line github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_lseg_interpt_line
func F_lseg_interpt_line(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lseg_closept_lseg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lseg_closept_lseg
func F_lseg_closept_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_point_invsl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_point_invsl
func F_point_invsl(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_line_closept_point github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_line_closept_point
func F_line_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_box_closept_point github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_box_closept_point
func F_box_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_point_inside github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_point_inside
func F_point_inside(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_point_mul_point github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_point_mul_point
func F_point_mul_point(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_int2send github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_int2send
func F_int2send(m *base.Module, l0 int32) int32
//go:linkname F_buildint2vector github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_buildint2vector
func F_buildint2vector(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generate_series_step_int4 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_series_step_int4
func F_generate_series_step_int4(m *base.Module, l0 int32) int32
//go:linkname F_generate_series_step_int8 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generate_series_step_int8
func F_generate_series_step_int8(m *base.Module, l0 int32) int32
//go:linkname F_array_to_json_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_array_to_json_internal
func F_array_to_json_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_composite_to_json github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_composite_to_json
func F_composite_to_json(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_escape_json github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_escape_json
func F_escape_json(m *base.Module, l0 int32, l1 int32)
//go:linkname F_datum_to_json_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_datum_to_json_internal
func F_datum_to_json_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_json_agg_transfn_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_json_agg_transfn_worker
func F_json_agg_transfn_worker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_json_object_agg_transfn_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_json_object_agg_transfn_worker
func F_json_object_agg_transfn_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_json_build_array_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_json_build_array_worker
func F_json_build_array_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_JsonbToCStringWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_JsonbToCStringWorker
func F_JsonbToCStringWorker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_JsonbToCString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonbToCString
func F_JsonbToCString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_JsonbTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_JsonbTypeName
func F_JsonbTypeName(m *base.Module, l0 int32) int32
//go:linkname F_datum_to_jsonb_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_datum_to_jsonb_internal
func F_datum_to_jsonb_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_jsonb_build_object_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_jsonb_build_object_worker
func F_jsonb_build_object_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_add_jsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_jsonb
func F_add_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_jsonb_agg_transfn_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_jsonb_agg_transfn_worker
func F_jsonb_agg_transfn_worker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_scalar_key github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_scalar_key
func F_make_scalar_key(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonbValueToJsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_JsonbValueToJsonb
func F_JsonbValueToJsonb(m *base.Module, l0 int32) int32
//go:linkname F_pushJsonbValue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pushJsonbValue
func F_pushJsonbValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_JsonbIteratorNext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonbIteratorNext
func F_JsonbIteratorNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_JsonbIteratorInit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_JsonbIteratorInit
func F_JsonbIteratorInit(m *base.Module, l0 int32) int32
//go:linkname F_findJsonbValueFromContainer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_findJsonbValueFromContainer
func F_findJsonbValueFromContainer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_parse_json_or_errsave github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_parse_json_or_errsave
func F_pg_parse_json_or_errsave(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_json_errsave_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_json_errsave_error
func F_json_errsave_error(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_makeJsonLexContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeJsonLexContext
func F_makeJsonLexContext(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_worker
func F_get_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_JsonbValueAsText github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_JsonbValueAsText
func F_JsonbValueAsText(m *base.Module, l0 int32) int32
//go:linkname F_jsonb_get_element github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_jsonb_get_element
func F_jsonb_get_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_get_record_type_from_argument github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_record_type_from_argument
func F_get_record_type_from_argument(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_record_type_from_query github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_record_type_from_query
func F_get_record_type_from_query(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_populate_composite github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_populate_composite
func F_populate_composite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_populate_array_dim_jsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_populate_array_dim_jsonb
func F_populate_array_dim_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_populate_recordset_record github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_populate_recordset_record
func F_populate_recordset_record(m *base.Module, l0 int32, l1 int32)
//go:linkname F_jsonb_set github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_jsonb_set
func F_jsonb_set(m *base.Module, l0 int32) int32
//go:linkname F_jsonb_delete_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_jsonb_delete_path
func F_jsonb_delete_path(m *base.Module, l0 int32) int32
//go:linkname F_parse_jsonb_index_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_jsonb_index_flags
func F_parse_jsonb_index_flags(m *base.Module, l0 int32) int32
//go:linkname F_iterate_jsonb_values github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_iterate_jsonb_values
func F_iterate_jsonb_values(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_iterate_json_values github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_iterate_json_values
func F_iterate_json_values(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_populate_array_report_expected_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_populate_array_report_expected_array
func F_populate_array_report_expected_array(m *base.Module, l0 int32, l1 int32)
//go:linkname F_jsonPathFromCstring github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jsonPathFromCstring
func F_jsonPathFromCstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_jspInitByBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_jspInitByBuffer
func F_jspInitByBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_jspOperationName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jspOperationName
func F_jspOperationName(m *base.Module, l0 int32) int32
//go:linkname F_jspGetNext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_jspGetNext
func F_jspGetNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetJsonTableExecContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetJsonTableExecContext
func F_GetJsonTableExecContext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonTableResetRowPattern github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonTableResetRowPattern
func F_JsonTableResetRowPattern(m *base.Module, l0 int32, l1 int32)
//go:linkname F_executeJsonPath github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_executeJsonPath
func F_executeJsonPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_jsonb_path_match_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_jsonb_path_match_internal
func F_jsonb_path_match_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_jsonb_path_query_array_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jsonb_path_query_array_internal
func F_jsonb_path_query_array_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_executeBoolItem github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_executeBoolItem
func F_executeBoolItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_executeNextItem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_executeNextItem
func F_executeNextItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_executeBinaryArithmExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_executeBinaryArithmExpr
func F_executeBinaryArithmExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_executeUnaryArithmExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_executeUnaryArithmExpr
func F_executeUnaryArithmExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_JsonbType github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_JsonbType
func F_JsonbType(m *base.Module, l0 int32) int32
//go:linkname F_executeItemUnwrapTargetArray github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_executeItemUnwrapTargetArray
func F_executeItemUnwrapTargetArray(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_executeAnyItem github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_executeAnyItem
func F_executeAnyItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_getArrayIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getArrayIndex
func F_getArrayIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_checkTimezoneIsUsedForCast github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_checkTimezoneIsUsedForCast
func F_checkTimezoneIsUsedForCast(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_executeNumericItemMethod github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_executeNumericItemMethod
func F_executeNumericItemMethod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_patternsel_common github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_patternsel_common
func F_patternsel_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64
//go:linkname F_string_to_const github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_string_to_const
func F_string_to_const(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_multirange_get_bounds github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_multirange_get_bounds
func F_multirange_get_bounds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_range_overlaps_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_overlaps_multirange_internal
func F_range_overlaps_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_calc_hist_selectivity_scalar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_calc_hist_selectivity_scalar
func F_calc_hist_selectivity_scalar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_calc_hist_selectivity_contains github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_calc_hist_selectivity_contains
func F_calc_hist_selectivity_contains(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) float64
//go:linkname F_calc_hist_selectivity_contained github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_calc_hist_selectivity_contained
func F_calc_hist_selectivity_contained(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) float64
//go:linkname F_namestrcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_namestrcmp
func F_namestrcmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_network_send github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_network_send
func F_network_send(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cidr_set_masklen_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cidr_set_masklen_internal
func F_cidr_set_masklen_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_convert_network_to_scalar github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_convert_network_to_scalar
func F_convert_network_to_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_inet_hist_value_sel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inet_hist_value_sel
func F_inet_hist_value_sel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64
//go:linkname F_mul_var github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_mul_var
func F_mul_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_add_var github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_var
func F_add_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_apply_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_apply_typmod
func F_apply_typmod(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_set_var_from_num github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_var_from_num
func F_set_var_from_num(m *base.Module, l0 int32, l1 int32)
//go:linkname F_cmp_abs_common github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cmp_abs_common
func F_cmp_abs_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_sub_var github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sub_var
func F_sub_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_generate_series_step_numeric github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_series_step_numeric
func F_generate_series_step_numeric(m *base.Module, l0 int32) int32
//go:linkname F_numericvar_to_double_no_overflow github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_numericvar_to_double_no_overflow
func F_numericvar_to_double_no_overflow(m *base.Module, l0 int32) float64
//go:linkname F_div_var_int github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_div_var_int
func F_div_var_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_numericvar_to_int64 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_numericvar_to_int64
func F_numericvar_to_int64(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_numeric_add_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_numeric_add_opt_error
func F_numeric_add_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_exp_var github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_exp_var
func F_exp_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_estimate_ln_dweight github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_estimate_ln_dweight
func F_estimate_ln_dweight(m *base.Module, l0 int32) int32
//go:linkname F_ln_var github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ln_var
func F_ln_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_int64_to_numeric github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_int64_to_numeric
func F_int64_to_numeric(m *base.Module, l0 int64) int32
//go:linkname F_int64_div_fast_to_numeric github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_int64_div_fast_to_numeric
func F_int64_div_fast_to_numeric(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_numeric_int4_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_numeric_int4_opt_error
func F_numeric_int4_opt_error(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_numeric_int8_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_numeric_int8_opt_error
func F_numeric_int8_opt_error(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_numericvar_serialize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_numericvar_serialize
func F_numericvar_serialize(m *base.Module, l0 int32, l1 int32)
//go:linkname F_do_numeric_discard github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_do_numeric_discard
func F_do_numeric_discard(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_strtoint32 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_strtoint32
func F_pg_strtoint32(m *base.Module, l0 int32) int32
//go:linkname F_pg_strtoint32_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_strtoint32_safe
func F_pg_strtoint32_safe(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_uint32in_subr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_uint32in_subr
func F_uint32in_subr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_ultoa_n github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_ultoa_n
func F_pg_ultoa_n(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_ulltoa_n github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_ulltoa_n
func F_pg_ulltoa_n(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_pg_ultostr_zeropad github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_ultostr_zeropad
func F_pg_ultostr_zeropad(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_buildoidvector github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_buildoidvector
func F_buildoidvector(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_setup_pct_info github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_setup_pct_info
func F_setup_pct_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_pg_perm_setlocale github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_perm_setlocale
func F_pg_perm_setlocale(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_locale
func F_check_locale(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_PGLC_localeconv github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PGLC_localeconv
func F_PGLC_localeconv(m *base.Module) int32
//go:linkname F_pg_newlocale_from_collation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_newlocale_from_collation
func F_pg_newlocale_from_collation(m *base.Module, l0 int32) int32
//go:linkname F_get_collation_actual_version github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_collation_actual_version
func F_get_collation_actual_version(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_strfold github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_strfold
func F_pg_strfold(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_strncoll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_strncoll
func F_pg_strncoll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_strxfrm github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_strxfrm
func F_pg_strxfrm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_strnxfrm github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_strnxfrm
func F_pg_strnxfrm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_create_pg_locale_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_create_pg_locale_builtin
func F_create_pg_locale_builtin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_pg_locale_icu github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_create_pg_locale_icu
func F_create_pg_locale_icu(m *base.Module) int32
//go:linkname F_create_pg_locale_libc github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_pg_locale_libc
func F_create_pg_locale_libc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_quote_literal_cstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_quote_literal_cstr
func F_quote_literal_cstr(m *base.Module, l0 int32) int32
//go:linkname F_get_range_io_data github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_range_io_data
func F_get_range_io_data(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_parse_bound github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_parse_bound
func F_range_parse_bound(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_make_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_range
func F_make_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_range_serialize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_serialize
func F_range_serialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_range_get_typcache github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_get_typcache
func F_range_get_typcache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_range_contains_elem_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_range_contains_elem_internal
func F_range_contains_elem_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_eq_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_eq_internal
func F_range_eq_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_cmp_bounds github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_cmp_bounds
func F_range_cmp_bounds(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_bounds_adjacent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bounds_adjacent
func F_bounds_adjacent(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_cmp
func F_range_cmp(m *base.Module, l0 int32) int32
//go:linkname F_range_gist_single_sorting_split github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_gist_single_sorting_split
func F_range_gist_single_sorting_split(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_range_gist_fallback_split github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_range_gist_fallback_split
func F_range_gist_fallback_split(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_parse_re_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parse_re_flags
func F_parse_re_flags(m *base.Module, l0 int32, l1 int32)
//go:linkname F_similar_escape_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_similar_escape_internal
func F_similar_escape_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_setup_regexp_matches github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_setup_regexp_matches
func F_setup_regexp_matches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_build_regexp_match_result github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_build_regexp_match_result
func F_build_regexp_match_result(m *base.Module, l0 int32) int32
//go:linkname F_stringToQualifiedNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_stringToQualifiedNameList
func F_stringToQualifiedNameList(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_format_procedure github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_format_procedure
func F_format_procedure(m *base.Module, l0 int32) int32
//go:linkname F_format_procedure_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_format_procedure_extended
func F_format_procedure_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_format_operator_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_format_operator_extended
func F_format_operator_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_format_operator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_format_operator
func F_format_operator(m *base.Module, l0 int32) int32
//go:linkname F_ri_CheckTrigger github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ri_CheckTrigger
func F_ri_CheckTrigger(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ri_FetchConstraintInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ri_FetchConstraintInfo
func F_ri_FetchConstraintInfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ri_set github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ri_set
func F_ri_set(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ri_KeysEqual github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ri_KeysEqual
func F_ri_KeysEqual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_record_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_record_cmp
func F_record_cmp(m *base.Module, l0 int32) int32
//go:linkname F_record_image_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_record_image_eq
func F_record_image_eq(m *base.Module, l0 int32) int32
//go:linkname F_record_image_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_record_image_cmp
func F_record_image_cmp(m *base.Module, l0 int32) int32
//go:linkname F_quote_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_quote_identifier
func F_quote_identifier(m *base.Module, l0 int32) int32
//go:linkname F_generate_relation_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generate_relation_name
func F_generate_relation_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generate_qualified_relation_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_generate_qualified_relation_name
func F_generate_qualified_relation_name(m *base.Module, l0 int32) int32
//go:linkname F_set_deparse_for_query github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_deparse_for_query
func F_set_deparse_for_query(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_rule_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_rule_expr
func F_get_rule_expr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_get_viewdef_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_get_viewdef_worker
func F_pg_get_viewdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fastgetattr_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fastgetattr_3
func F_fastgetattr_3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_set_rtable_names github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_rtable_names
func F_set_rtable_names(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_set_simple_column_names github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_simple_column_names
func F_set_simple_column_names(m *base.Module, l0 int32)
//go:linkname F_generate_collation_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_generate_collation_name
func F_generate_collation_name(m *base.Module, l0 int32) int32
//go:linkname F_get_opclass_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_opclass_name
func F_get_opclass_name(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_with_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_with_clause
func F_get_with_clause(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_setop_query github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_setop_query
func F_get_setop_query(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_rule_sortgroupclause github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rule_sortgroupclause
func F_get_rule_sortgroupclause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_target_list github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_target_list
func F_get_target_list(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_from_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_from_clause
func F_get_from_clause(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_appendContextKeyword github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_appendContextKeyword
func F_appendContextKeyword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_get_rule_groupingset github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_rule_groupingset
func F_get_rule_groupingset(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_rule_windowspec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_rule_windowspec
func F_get_rule_windowspec(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_returning_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_returning_clause
func F_get_returning_clause(m *base.Module, l0 int32, l1 int32)
//go:linkname F_processIndirection github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_processIndirection
func F_processIndirection(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rule_list_toplevel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rule_list_toplevel
func F_get_rule_list_toplevel(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_rule_orderby github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_rule_orderby
func F_get_rule_orderby(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pg_get_constraintdef_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_get_constraintdef_worker
func F_pg_get_constraintdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_print_function_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_print_function_rettype
func F_print_function_rettype(m *base.Module, l0 int32, l1 int32)
//go:linkname F_quote_qualified_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_quote_qualified_identifier
func F_quote_qualified_identifier(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_deparse_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_set_deparse_plan
func F_set_deparse_plan(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_restriction_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_restriction_variable
func F_get_restriction_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_mcv_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_mcv_selectivity
func F_mcv_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64
//go:linkname F_examine_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_examine_variable
func F_examine_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ineq_histogram_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ineq_histogram_selectivity
func F_ineq_histogram_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) float64
//go:linkname F_estimate_array_length github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_estimate_array_length
func F_estimate_array_length(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_estimate_num_groups github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_estimate_num_groups
func F_estimate_num_groups(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32) float64
//go:linkname F_genericcostestimate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_genericcostestimate
func F_genericcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32)
//go:linkname F_examine_indexcol_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_examine_indexcol_variable
func F_examine_indexcol_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_anytimestamp_typmod_check github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_anytimestamp_typmod_check
func F_anytimestamp_typmod_check(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AdjustTimestampForTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AdjustTimestampForTypmod
func F_AdjustTimestampForTypmod(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_EncodeSpecialTimestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EncodeSpecialTimestamp
func F_EncodeSpecialTimestamp(m *base.Module, l0 int64, l1 int32)
//go:linkname F_make_timestamp_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_timestamp_internal
func F_make_timestamp_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) int64
//go:linkname F_timestamp2timestamptz_opt_overflow github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamp2timestamptz_opt_overflow
func F_timestamp2timestamptz_opt_overflow(m *base.Module, l0 int64, l1 int32) int64
//go:linkname F_AdjustIntervalForTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AdjustIntervalForTypmod
func F_AdjustIntervalForTypmod(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_intervaltypmodleastfield github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_intervaltypmodleastfield
func F_intervaltypmodleastfield(m *base.Module, l0 int32) int32
//go:linkname F_GetCurrentTimestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetCurrentTimestamp
func F_GetCurrentTimestamp(m *base.Module) int64
//go:linkname F_timestamptz2timestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_timestamptz2timestamp
func F_timestamptz2timestamp(m *base.Module, l0 int64) int64
//go:linkname F_timestamptz_to_str github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_timestamptz_to_str
func F_timestamptz_to_str(m *base.Module, l0 int64) int32
//go:linkname F_timestamp_cmp_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamp_cmp_internal
func F_timestamp_cmp_internal(m *base.Module, l0 int64, l1 int64) int32
//go:linkname F_interval_um_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_interval_um_internal
func F_interval_um_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_timestamptz_pl_interval_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamptz_pl_interval_internal
func F_timestamptz_pl_interval_internal(m *base.Module, l0 int64, l1 int32, l2 int32) int64
//go:linkname F_gin_extract_tsvector github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gin_extract_tsvector
func F_gin_extract_tsvector(m *base.Module, l0 int32) int32
//go:linkname F_makepol_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makepol_1
func F_makepol_1(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_findoprnd_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_findoprnd_recurse
func F_findoprnd_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_maketree github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_maketree
func F_maketree(m *base.Module, l0 int32) int32
//go:linkname F_plainnode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_plainnode
func F_plainnode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clean_stopword_intree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_clean_stopword_intree
func F_clean_stopword_intree(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CompareTSQ github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CompareTSQ
func F_CompareTSQ(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dofindsubquery github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dofindsubquery
func F_dofindsubquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_QT2QTN github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_QT2QTN
func F_QT2QTN(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_QTNFree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_QTNFree
func F_QTNFree(m *base.Module, l0 int32)
//go:linkname F_QTNodeCompare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_QTNodeCompare
func F_QTNodeCompare(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_QTNSort github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_QTNSort
func F_QTNSort(m *base.Module, l0 int32)
//go:linkname F_QTNTernary github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_QTNTernary
func F_QTNTernary(m *base.Module, l0 int32)
//go:linkname F_QTNBinary github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_QTNBinary
func F_QTNBinary(m *base.Module, l0 int32)
//go:linkname F_QTN2QT github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_QTN2QT
func F_QTN2QT(m *base.Module, l0 int32) int32
//go:linkname F_calc_rank_cd github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_calc_rank_cd
func F_calc_rank_cd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32
//go:linkname F_TS_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TS_execute
func F_TS_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_TS_execute_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TS_execute_recurse
func F_TS_execute_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_checkclass_str github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_checkclass_str
func F_checkclass_str(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_init_tsvector_parser github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_init_tsvector_parser
func F_init_tsvector_parser(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_close_tsvector_parser github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_close_tsvector_parser
func F_close_tsvector_parser(m *base.Module, l0 int32)
//go:linkname F_gettoken_tsvector github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gettoken_tsvector
func F_gettoken_tsvector(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_bit_catenate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bit_catenate
func F_bit_catenate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bit_overlay github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bit_overlay
func F_bit_overlay(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_cstring_to_text github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cstring_to_text
func F_cstring_to_text(m *base.Module, l0 int32) int32
//go:linkname F_text_to_cstring github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_text_to_cstring
func F_text_to_cstring(m *base.Module, l0 int32) int32
//go:linkname F_text_to_cstring_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_text_to_cstring_buffer
func F_text_to_cstring_buffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_text_substring github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text_substring
func F_text_substring(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_varstr_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_varstr_cmp
func F_varstr_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_varstr_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_varstr_sortsupport
func F_varstr_sortsupport(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_bytea_overlay github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bytea_overlay
func F_bytea_overlay(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_textToQualifiedNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_textToQualifiedNameList
func F_textToQualifiedNameList(m *base.Module, l0 int32) int32
//go:linkname F_SplitIdentifierString github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SplitIdentifierString
func F_SplitIdentifierString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_concat_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_concat_internal
func F_concat_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_text_format_parse_digits github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_text_format_parse_digits
func F_text_format_parse_digits(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_text_format_append_string github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_text_format_append_string
func F_text_format_append_string(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_varstr_levenshtein github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_varstr_levenshtein
func F_varstr_levenshtein(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_varstr_levenshtein_less_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_varstr_levenshtein_less_equal
func F_varstr_levenshtein_less_equal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_xmlconcat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_xmlconcat
func F_xmlconcat(m *base.Module) int32
//go:linkname F_xmlparse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_xmlparse
func F_xmlparse(m *base.Module) int32
//go:linkname F_map_sql_identifier_to_xml_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_map_sql_identifier_to_xml_name
func F_map_sql_identifier_to_xml_name(m *base.Module) int32
//go:linkname F_query_to_xml_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_query_to_xml_internal
func F_query_to_xml_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_map_sql_type_to_xml_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_map_sql_type_to_xml_name
func F_map_sql_type_to_xml_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CatCacheRemoveCTup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CatCacheRemoveCTup
func F_CatCacheRemoveCTup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateCacheMemoryContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateCacheMemoryContext
func F_CreateCacheMemoryContext(m *base.Module)
//go:linkname F_SearchCatCache github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SearchCatCache
func F_SearchCatCache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SearchCatCacheInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SearchCatCacheInternal
func F_SearchCatCacheInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_SearchCatCache2 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SearchCatCache2
func F_SearchCatCache2(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchCatCacheList github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SearchCatCacheList
func F_SearchCatCacheList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ReleaseCatCacheList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReleaseCatCacheList
func F_ReleaseCatCacheList(m *base.Module, l0 int32)
//go:linkname F_InvalidateSystemCachesExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InvalidateSystemCachesExtended
func F_InvalidateSystemCachesExtended(m *base.Module)
//go:linkname F_LocalExecuteInvalidationMessage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LocalExecuteInvalidationMessage
func F_LocalExecuteInvalidationMessage(m *base.Module, l0 int32)
//go:linkname F_InvalidateSystemCaches github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InvalidateSystemCaches
func F_InvalidateSystemCaches(m *base.Module)
//go:linkname F_PrepareInvalidationState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PrepareInvalidationState
func F_PrepareInvalidationState(m *base.Module) int32
//go:linkname F_RegisterRelcacheInvalidation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RegisterRelcacheInvalidation
func F_RegisterRelcacheInvalidation(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CacheInvalidateRelcache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CacheInvalidateRelcache
func F_CacheInvalidateRelcache(m *base.Module, l0 int32)
//go:linkname F_CacheInvalidateRelcacheByRelid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CacheInvalidateRelcacheByRelid
func F_CacheInvalidateRelcacheByRelid(m *base.Module, l0 int32)
//go:linkname F_CacheRegisterSyscacheCallback github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CacheRegisterSyscacheCallback
func F_CacheRegisterSyscacheCallback(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CacheRegisterRelcacheCallback github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CacheRegisterRelcacheCallback
func F_CacheRegisterRelcacheCallback(m *base.Module, l0 int32)
//go:linkname F_get_op_opfamily_strategy github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_op_opfamily_strategy
func F_get_op_opfamily_strategy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_op_opfamily_properties github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_op_opfamily_properties
func F_get_op_opfamily_properties(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_opfamily_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_opfamily_member
func F_get_opfamily_member(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_opfamily_member_for_cmptype github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_opfamily_member_for_cmptype
func F_get_opfamily_member_for_cmptype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_opfamily_method github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_opfamily_method
func F_get_opfamily_method(m *base.Module, l0 int32) int32
//go:linkname F_get_ordering_op_properties github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_ordering_op_properties
func F_get_ordering_op_properties(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_equality_op_for_ordering_op github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_equality_op_for_ordering_op
func F_get_equality_op_for_ordering_op(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_mergejoin_opfamilies github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_mergejoin_opfamilies
func F_get_mergejoin_opfamilies(m *base.Module, l0 int32) int32
//go:linkname F_get_op_hash_functions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_op_hash_functions
func F_get_op_hash_functions(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_op_index_interpretation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_op_index_interpretation
func F_get_op_index_interpretation(m *base.Module, l0 int32) int32
//go:linkname F_get_negator github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_negator
func F_get_negator(m *base.Module, l0 int32) int32
//go:linkname F_get_attnum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_attnum
func F_get_attnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_attgenerated github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_attgenerated
func F_get_attgenerated(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_atttypetypmodcoll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_atttypetypmodcoll
func F_get_atttypetypmodcoll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_get_collation_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_collation_name
func F_get_collation_name(m *base.Module, l0 int32) int32
//go:linkname F_get_collation_isdeterministic github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_collation_isdeterministic
func F_get_collation_isdeterministic(m *base.Module, l0 int32) int32
//go:linkname F_get_constraint_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_constraint_name
func F_get_constraint_name(m *base.Module, l0 int32) int32
//go:linkname F_get_language_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_language_name
func F_get_language_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_opclass_opfamily_and_input_type github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_opclass_opfamily_and_input_type
func F_get_opclass_opfamily_and_input_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_opfamily_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_opfamily_name
func F_get_opfamily_name(m *base.Module, l0 int32) int32
//go:linkname F_get_opname github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_opname
func F_get_opname(m *base.Module, l0 int32) int32
//go:linkname F_get_op_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_op_rettype
func F_get_op_rettype(m *base.Module, l0 int32) int32
//go:linkname F_op_input_types github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_op_input_types
func F_op_input_types(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_op_mergejoinable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_op_mergejoinable
func F_op_mergejoinable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_op_hashjoinable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_op_hashjoinable
func F_op_hashjoinable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_op_strict github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_op_strict
func F_op_strict(m *base.Module, l0 int32) int32
//go:linkname F_func_volatile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_func_volatile
func F_func_volatile(m *base.Module, l0 int32) int32
//go:linkname F_get_commutator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_commutator
func F_get_commutator(m *base.Module, l0 int32) int32
//go:linkname F_get_oprrest github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_oprrest
func F_get_oprrest(m *base.Module, l0 int32) int32
//go:linkname F_get_func_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_func_signature
func F_get_func_signature(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_func_support github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_func_support
func F_get_func_support(m *base.Module, l0 int32) int32
//go:linkname F_get_relname_relid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_relname_relid
func F_get_relname_relid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rel_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rel_name
func F_get_rel_name(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_namespace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_rel_namespace
func F_get_rel_namespace(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_type_id github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_rel_type_id
func F_get_rel_type_id(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_relkind github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_rel_relkind
func F_get_rel_relkind(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_relispartition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_rel_relispartition
func F_get_rel_relispartition(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_tablespace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_rel_tablespace
func F_get_rel_tablespace(m *base.Module, l0 int32) int32
//go:linkname F_get_typisdefined github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_typisdefined
func F_get_typisdefined(m *base.Module, l0 int32) int32
//go:linkname F_get_typlenbyval github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_typlenbyval
func F_get_typlenbyval(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_typlenbyvalalign github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typlenbyvalalign
func F_get_typlenbyvalalign(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_getBaseType github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getBaseType
func F_getBaseType(m *base.Module, l0 int32) int32
//go:linkname F_getBaseTypeAndTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getBaseTypeAndTypmod
func F_getBaseTypeAndTypmod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_typavgwidth github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_typavgwidth
func F_get_typavgwidth(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_typtype github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typtype
func F_get_typtype(m *base.Module, l0 int32) int32
//go:linkname F_type_is_rowtype github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_type_is_rowtype
func F_type_is_rowtype(m *base.Module, l0 int32) int32
//go:linkname F_type_is_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_type_is_multirange
func F_type_is_multirange(m *base.Module, l0 int32) int32
//go:linkname F_get_type_category_preferred github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_type_category_preferred
func F_get_type_category_preferred(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_typ_typrelid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typ_typrelid
func F_get_typ_typrelid(m *base.Module, l0 int32) int32
//go:linkname F_get_element_type github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_element_type
func F_get_element_type(m *base.Module, l0 int32) int32
//go:linkname F_get_promoted_array_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_promoted_array_type
func F_get_promoted_array_type(m *base.Module, l0 int32) int32
//go:linkname F_get_base_element_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_base_element_type
func F_get_base_element_type(m *base.Module, l0 int32) int32
//go:linkname F_getTypeInputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getTypeInputInfo
func F_getTypeInputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_getTypeBinaryInputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getTypeBinaryInputInfo
func F_getTypeBinaryInputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_getTypeBinaryOutputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getTypeBinaryOutputInfo
func F_getTypeBinaryOutputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_type_is_collatable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_type_is_collatable
func F_type_is_collatable(m *base.Module, l0 int32) int32
//go:linkname F_getSubscriptingRoutines github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_getSubscriptingRoutines
func F_getSubscriptingRoutines(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_attstatsslot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_attstatsslot
func F_get_attstatsslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_free_attstatsslot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_free_attstatsslot
func F_free_attstatsslot(m *base.Module, l0 int32)
//go:linkname F_get_namespace_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_namespace_name
func F_get_namespace_name(m *base.Module, l0 int32) int32
//go:linkname F_get_namespace_name_or_temp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_namespace_name_or_temp
func F_get_namespace_name_or_temp(m *base.Module, l0 int32) int32
//go:linkname F_get_multirange_range github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_multirange_range
func F_get_multirange_range(m *base.Module, l0 int32) int32
//go:linkname F_get_index_column_opclass github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_index_column_opclass
func F_get_index_column_opclass(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_index_isclustered github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_index_isclustered
func F_get_index_isclustered(m *base.Module, l0 int32) int32
//go:linkname F_get_publication_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_publication_name
func F_get_publication_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_subscription_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_subscription_name
func F_get_subscription_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationGetPartitionKey github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetPartitionKey
func F_RelationGetPartitionKey(m *base.Module, l0 int32) int32
//go:linkname F_CreateCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateCachedPlan
func F_CreateCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CompleteCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CompleteCachedPlan
func F_CompleteCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_SaveCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SaveCachedPlan
func F_SaveCachedPlan(m *base.Module, l0 int32)
//go:linkname F_DropCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DropCachedPlan
func F_DropCachedPlan(m *base.Module, l0 int32)
//go:linkname F_GetCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetCachedPlan
func F_GetCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_RevalidateCachedQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RevalidateCachedQuery
func F_RevalidateCachedQuery(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ScanQueryForLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ScanQueryForLocks
func F_ScanQueryForLocks(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReleaseCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReleaseCachedPlan
func F_ReleaseCachedPlan(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CachedPlanGetTargetList github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CachedPlanGetTargetList
func F_CachedPlanGetTargetList(m *base.Module, l0 int32) int32
//go:linkname F_RelationInitIndexAccessInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RelationInitIndexAccessInfo
func F_RelationInitIndexAccessInfo(m *base.Module, l0 int32)
//go:linkname F_RelationGetIndexAttOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetIndexAttOptions
func F_RelationGetIndexAttOptions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationInitTableAccessMethod github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationInitTableAccessMethod
func F_RelationInitTableAccessMethod(m *base.Module, l0 int32)
//go:linkname F_RelationIdGetRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationIdGetRelation
func F_RelationIdGetRelation(m *base.Module, l0 int32) int32
//go:linkname F_RelationRebuildRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationRebuildRelation
func F_RelationRebuildRelation(m *base.Module, l0 int32)
//go:linkname F_RelationInitPhysicalAddr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationInitPhysicalAddr
func F_RelationInitPhysicalAddr(m *base.Module, l0 int32)
//go:linkname F_ScanPgRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ScanPgRelation
func F_ScanPgRelation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RelationParseRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationParseRelOptions
func F_RelationParseRelOptions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationDestroyRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationDestroyRelation
func F_RelationDestroyRelation(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationBuildRuleLock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationBuildRuleLock
func F_RelationBuildRuleLock(m *base.Module, l0 int32)
//go:linkname F_RelationDecrementReferenceCount github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationDecrementReferenceCount
func F_RelationDecrementReferenceCount(m *base.Module, l0 int32)
//go:linkname F_RelationClearRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationClearRelation
func F_RelationClearRelation(m *base.Module, l0 int32)
//go:linkname F_load_relcache_init_file github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_load_relcache_init_file
func F_load_relcache_init_file(m *base.Module, l0 int32) int32
//go:linkname F_formrdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_formrdesc
func F_formrdesc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_RelationCacheInitializePhase3 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationCacheInitializePhase3
func F_RelationCacheInitializePhase3(m *base.Module)
//go:linkname F_RelationGetIndexList github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationGetIndexList
func F_RelationGetIndexList(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexExpressions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetIndexExpressions
func F_RelationGetIndexExpressions(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexPredicate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetIndexPredicate
func F_RelationGetIndexPredicate(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexAttrBitmap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RelationGetIndexAttrBitmap
func F_RelationGetIndexAttrBitmap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_errtable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_errtable
func F_errtable(m *base.Module, l0 int32)
//go:linkname F_errtablecol github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_errtablecol
func F_errtablecol(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errtableconstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errtableconstraint
func F_errtableconstraint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationCacheInitFilePreInvalidate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RelationCacheInitFilePreInvalidate
func F_RelationCacheInitFilePreInvalidate(m *base.Module)
//go:linkname F_read_relmap_file github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_read_relmap_file
func F_read_relmap_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_tablespace_page_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_tablespace_page_costs
func F_get_tablespace_page_costs(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_SearchSysCache1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchSysCache1
func F_SearchSysCache1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCache2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCache2
func F_SearchSysCache2(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchSysCache3 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SearchSysCache3
func F_SearchSysCache3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SearchSysCache4 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCache4
func F_SearchSysCache4(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SearchSysCacheCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchSysCacheCopy
func F_SearchSysCacheCopy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchSysCacheLockedCopy1 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCacheLockedCopy1
func F_SearchSysCacheLockedCopy1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheExists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCacheExists
func F_SearchSysCacheExists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SearchSysCacheAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SearchSysCacheAttName
func F_SearchSysCacheAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheCopyAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SearchSysCacheCopyAttName
func F_SearchSysCacheCopyAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheExistsAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SearchSysCacheExistsAttName
func F_SearchSysCacheExistsAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SysCacheGetAttr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SysCacheGetAttr
func F_SysCacheGetAttr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SysCacheGetAttrNotNull github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SysCacheGetAttrNotNull
func F_SysCacheGetAttrNotNull(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetSysCacheHashValue github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetSysCacheHashValue
func F_GetSysCacheHashValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lookup_ts_dictionary_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lookup_ts_dictionary_cache
func F_lookup_ts_dictionary_cache(m *base.Module, l0 int32) int32
//go:linkname F_getTSCurrentConfig github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_getTSCurrentConfig
func F_getTSCurrentConfig(m *base.Module) int32
//go:linkname F_lookup_type_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lookup_type_cache
func F_lookup_type_cache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DomainHasConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DomainHasConstraints
func F_DomainHasConstraints(m *base.Module, l0 int32) int32
//go:linkname F_lookup_rowtype_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lookup_rowtype_tupdesc
func F_lookup_rowtype_tupdesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookup_rowtype_tupdesc_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lookup_rowtype_tupdesc_internal
func F_lookup_rowtype_tupdesc_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assign_record_type_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_assign_record_type_identifier
func F_assign_record_type_identifier(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_errstart_cold github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errstart_cold
func F_errstart_cold(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errstart github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_errstart
func F_errstart(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_errmsg_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errmsg_internal
func F_errmsg_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errfinish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errfinish
func F_errfinish(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_re_throw github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_re_throw
func F_pg_re_throw(m *base.Module)
//go:linkname F_errsave_start github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errsave_start
func F_errsave_start(m *base.Module, l0 int32) int32
//go:linkname F_errsave_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errsave_finish
func F_errsave_finish(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errcode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errcode
func F_errcode(m *base.Module, l0 int32)
//go:linkname F_errcode_for_file_access github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errcode_for_file_access
func F_errcode_for_file_access(m *base.Module)
//go:linkname F_errmsg github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errmsg
func F_errmsg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errmsg_plural github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errmsg_plural
func F_errmsg_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errdetail github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errdetail
func F_errdetail(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errdetail_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errdetail_internal
func F_errdetail_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errdetail_log github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errdetail_log
func F_errdetail_log(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errdetail_log_plural github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errdetail_log_plural
func F_errdetail_log_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errhint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errhint
func F_errhint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errhint_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errhint_internal
func F_errhint_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errcontext_msg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errcontext_msg
func F_errcontext_msg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_set_errcontext_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_errcontext_domain
func F_set_errcontext_domain(m *base.Module, l0 int32)
//go:linkname F_errhidestmt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errhidestmt
func F_errhidestmt(m *base.Module)
//go:linkname F_errhidecontext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errhidecontext
func F_errhidecontext(m *base.Module)
//go:linkname F_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errposition
func F_errposition(m *base.Module, l0 int32) int32
//go:linkname F_internalerrposition github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_internalerrposition
func F_internalerrposition(m *base.Module, l0 int32)
//go:linkname F_internalerrquery github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_internalerrquery
func F_internalerrquery(m *base.Module, l0 int32) int32
//go:linkname F_err_generic_string github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_err_generic_string
func F_err_generic_string(m *base.Module, l0 int32, l1 int32)
//go:linkname F_geterrposition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_geterrposition
func F_geterrposition(m *base.Module) int32
//go:linkname F_getinternalerrposition github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getinternalerrposition
func F_getinternalerrposition(m *base.Module) int32
//go:linkname F_format_elog_string github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_format_elog_string
func F_format_elog_string(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FlushErrorState github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FlushErrorState
func F_FlushErrorState(m *base.Module)
//go:linkname F_ThrowErrorData github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ThrowErrorData
func F_ThrowErrorData(m *base.Module, l0 int32)
//go:linkname F_ReThrowError github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReThrowError
func F_ReThrowError(m *base.Module, l0 int32)
//go:linkname F_get_formatted_log_time github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_formatted_log_time
func F_get_formatted_log_time(m *base.Module) int32
//go:linkname F_get_formatted_start_time github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_formatted_start_time
func F_get_formatted_start_time(m *base.Module) int32
//go:linkname F_appendJSONKeyValueFmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_appendJSONKeyValueFmt
func F_appendJSONKeyValueFmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_load_external_function github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_load_external_function
func F_load_external_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_load_file github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_load_file
func F_load_file(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fmgr_info_cxt_security github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fmgr_info_cxt_security
func F_fmgr_info_cxt_security(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_fetch_finfo_record github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fetch_finfo_record
func F_fetch_finfo_record(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_detoast_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_detoast_datum
func F_pg_detoast_datum(m *base.Module, l0 int32) int32
//go:linkname F_DirectFunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DirectFunctionCall1Coll
func F_DirectFunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DirectFunctionCall3Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DirectFunctionCall3Coll
func F_DirectFunctionCall3Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_DirectFunctionCall5Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DirectFunctionCall5Coll
func F_DirectFunctionCall5Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_CallerFInfoFunctionCall2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CallerFInfoFunctionCall2
func F_CallerFInfoFunctionCall2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_FunctionCall2Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FunctionCall2Coll
func F_FunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_FunctionCall4Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FunctionCall4Coll
func F_FunctionCall4Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_FunctionCall8Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FunctionCall8Coll
func F_FunctionCall8Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_OidFunctionCall0Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_OidFunctionCall0Coll
func F_OidFunctionCall0Coll(m *base.Module, l0 int32) int32
//go:linkname F_OidFunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_OidFunctionCall1Coll
func F_OidFunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_OidFunctionCall2Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_OidFunctionCall2Coll
func F_OidFunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_DirectInputFunctionCallSafe github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DirectInputFunctionCallSafe
func F_DirectInputFunctionCallSafe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SendFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SendFunctionCall
func F_SendFunctionCall(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_OidInputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_OidInputFunctionCall
func F_OidInputFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_Int64GetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Int64GetDatum
func F_Int64GetDatum(m *base.Module, l0 int64) int32
//go:linkname F_pg_detoast_datum_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_detoast_datum_copy
func F_pg_detoast_datum_copy(m *base.Module, l0 int32) int32
//go:linkname F_pg_detoast_datum_packed github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_detoast_datum_packed
func F_pg_detoast_datum_packed(m *base.Module, l0 int32) int32
//go:linkname F_get_fn_expr_argtype github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_fn_expr_argtype
func F_get_fn_expr_argtype(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_fn_opclass_options github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_fn_opclass_options
func F_get_fn_opclass_options(m *base.Module, l0 int32) int32
//go:linkname F_CheckFunctionValidatorAccess github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckFunctionValidatorAccess
func F_CheckFunctionValidatorAccess(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_InitMaterializedSRF github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InitMaterializedSRF
func F_InitMaterializedSRF(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_call_result_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_call_result_type
func F_get_call_result_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_init_MultiFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_init_MultiFuncCall
func F_init_MultiFuncCall(m *base.Module, l0 int32) int32
//go:linkname F_build_function_result_tupdesc_t github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_function_result_tupdesc_t
func F_build_function_result_tupdesc_t(m *base.Module, l0 int32) int32
//go:linkname F_get_expr_result_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_expr_result_type
func F_get_expr_result_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_expr_result_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_expr_result_tupdesc
func F_get_expr_result_tupdesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_func_arg_info github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_func_arg_info
func F_get_func_arg_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_estimate_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hash_estimate_size
func F_hash_estimate_size(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hash_destroy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hash_destroy
func F_hash_destroy(m *base.Module, l0 int32)
//go:linkname F_get_hash_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_hash_value
func F_get_hash_value(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hash_search github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_search
func F_hash_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_corrupted github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_corrupted
func F_hash_corrupted(m *base.Module, l0 int32)
//go:linkname F_hash_seq_search github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_seq_search
func F_hash_seq_search(m *base.Module, l0 int32) int32
//go:linkname F_GetBackendTypeDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetBackendTypeDesc
func F_GetBackendTypeDesc(m *base.Module, l0 int32) int32
//go:linkname F_SetDatabasePath github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SetDatabasePath
func F_SetDatabasePath(m *base.Module, l0 int32)
//go:linkname F_ValidatePgVersion github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ValidatePgVersion
func F_ValidatePgVersion(m *base.Module, l0 int32)
//go:linkname F_has_rolreplication github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_has_rolreplication
func F_has_rolreplication(m *base.Module, l0 int32) int32
//go:linkname F_InitializeSessionUserId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InitializeSessionUserId
func F_InitializeSessionUserId(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_InitializeSessionUserIdStandalone github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InitializeSessionUserIdStandalone
func F_InitializeSessionUserIdStandalone(m *base.Module)
//go:linkname F_InitializeSystemUser github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InitializeSystemUser
func F_InitializeSystemUser(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetUserNameFromId github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetUserNameFromId
func F_GetUserNameFromId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_load_libraries github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_load_libraries
func F_load_libraries(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_process_startup_options github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_process_startup_options
func F_process_startup_options(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SwitchToUntrustedUser github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SwitchToUntrustedUser
func F_SwitchToUntrustedUser(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RestoreUserContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RestoreUserContext
func F_RestoreUserContext(m *base.Module, l0 int32)
//go:linkname F_local2local github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_local2local
func F_local2local(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_latin2mic github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_latin2mic
func F_latin2mic(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_mic2latin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_mic2latin
func F_mic2latin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_mic2latin_with_table github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_mic2latin_with_table
func F_mic2latin_with_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_UtfToLocal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UtfToLocal
func F_UtfToLocal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_InitializeClientEncoding github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InitializeClientEncoding
func F_InitializeClientEncoding(m *base.Module)
//go:linkname F_pg_do_encoding_conversion github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_do_encoding_conversion
func F_pg_do_encoding_conversion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_report_invalid_encoding github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_invalid_encoding
func F_report_invalid_encoding(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_client_to_server github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_client_to_server
func F_pg_client_to_server(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_any_to_server github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_any_to_server
func F_pg_any_to_server(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_server_to_client github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_server_to_client
func F_pg_server_to_client(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_server_to_any github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_server_to_any
func F_pg_server_to_any(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_mb2wchar_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_mb2wchar_with_len
func F_pg_mb2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_mblen_cstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_mblen_cstr
func F_pg_mblen_cstr(m *base.Module, l0 int32) int32
//go:linkname F_report_invalid_encoding_db github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_report_invalid_encoding_db
func F_report_invalid_encoding_db(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_mblen_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mblen_range
func F_pg_mblen_range(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mbstrlen_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mbstrlen_with_len
func F_pg_mbstrlen_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mbcliplen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mbcliplen
func F_pg_mbcliplen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_mbcharcliplen github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_mbcharcliplen
func F_pg_mbcharcliplen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_verifymbstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_verifymbstr
func F_pg_verifymbstr(m *base.Module, l0 int32, l1 int32)
//go:linkname F_report_untranslatable_char github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_untranslatable_char
func F_report_untranslatable_char(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_AbsoluteConfigLocation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AbsoluteConfigLocation
func F_AbsoluteConfigLocation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_guc_strdup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_guc_strdup
func F_guc_strdup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assignable_custom_variable_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_assignable_custom_variable_name
func F_assignable_custom_variable_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_set_config_option github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_config_option
func F_set_config_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_guc_malloc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_guc_malloc
func F_guc_malloc(m *base.Module, l0 int32) int32
//go:linkname F_get_guc_variables github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_guc_variables
func F_get_guc_variables(m *base.Module, l0 int32) int32
//go:linkname F_call_bool_check_hook github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_call_bool_check_hook
func F_call_bool_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_call_real_check_hook github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_call_real_check_hook
func F_call_real_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_call_string_check_hook github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_call_string_check_hook
func F_call_string_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ResetAllOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResetAllOptions
func F_ResetAllOptions(m *base.Module)
//go:linkname F_push_old_value github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_push_old_value
func F_push_old_value(m *base.Module, l0 int32, l1 int32)
//go:linkname F_discard_stack_value github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_discard_stack_value
func F_discard_stack_value(m *base.Module, l0 int32, l1 int32)
//go:linkname F_set_extra_field github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_extra_field
func F_set_extra_field(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_set_string_field github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_string_field
func F_set_string_field(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RestrictSearchPath github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RestrictSearchPath
func F_RestrictSearchPath(m *base.Module)
//go:linkname F_ReportChangedGUCOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReportChangedGUCOptions
func F_ReportChangedGUCOptions(m *base.Module)
//go:linkname F_parse_int github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_int
func F_parse_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_parse_and_validate_value github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parse_and_validate_value
func F_parse_and_validate_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_AlterSystemSetConfigFile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AlterSystemSetConfigFile
func F_AlterSystemSetConfigFile(m *base.Module, l0 int32)
//go:linkname F_init_custom_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_init_custom_variable
func F_init_custom_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_define_custom_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_define_custom_variable
func F_define_custom_variable(m *base.Module, l0 int32)
//go:linkname F_GetConfigOptionByName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetConfigOptionByName
func F_GetConfigOptionByName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_TransformGUCArray github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TransformGUCArray
func F_TransformGUCArray(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ProcessConfigFile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ProcessConfigFile
func F_ProcessConfigFile(m *base.Module, l0 int32)
//go:linkname F_ParseConfigFp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ParseConfigFp
func F_ParseConfigFp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_flatten_set_variable_args github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_flatten_set_variable_args
func F_flatten_set_variable_args(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SetPGVariable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SetPGVariable
func F_SetPGVariable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_rusage_show github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_rusage_show
func F_pg_rusage_show(m *base.Module, l0 int32) int32
//go:linkname F_ENRMetadataGetTupDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ENRMetadataGetTupDesc
func F_ENRMetadataGetTupDesc(m *base.Module, l0 int32) int32
//go:linkname F_check_enable_rls github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_enable_rls
func F_check_enable_rls(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_stack_depth github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_stack_depth
func F_check_stack_depth(m *base.Module)
//go:linkname F_superuser github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_superuser
func F_superuser(m *base.Module) int32
//go:linkname F_superuser_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_superuser_arg
func F_superuser_arg(m *base.Module, l0 int32) int32
//go:linkname F_RegisterTimeout github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RegisterTimeout
func F_RegisterTimeout(m *base.Module, l0 int32, l1 int32)
//go:linkname F_enable_timeout_after github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_enable_timeout_after
func F_enable_timeout_after(m *base.Module, l0 int32, l1 int32)
//go:linkname F_enable_timeouts github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_enable_timeouts
func F_enable_timeouts(m *base.Module, l0 int32, l1 int32)
//go:linkname F_disable_timeout github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_disable_timeout
func F_disable_timeout(m *base.Module, l0 int32)
//go:linkname F_create_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_internal
func F_create_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_dsa_allocate_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsa_allocate_extended
func F_dsa_allocate_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dsa_free github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dsa_free
func F_dsa_free(m *base.Module, l0 int32, l1 int32)
//go:linkname F_rebin_segment github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_rebin_segment
func F_rebin_segment(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dsa_get_address github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsa_get_address
func F_dsa_get_address(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FreePageBtreeCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreePageBtreeCleanup
func F_FreePageBtreeCleanup(m *base.Module, l0 int32) int32
//go:linkname F_FreePageManagerPutInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreePageManagerPutInternal
func F_FreePageManagerPutInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GenerationContextCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GenerationContextCreate
func F_GenerationContextCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_MemoryContextReset github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextReset
func F_MemoryContextReset(m *base.Module, l0 int32)
//go:linkname F_MemoryContextDeleteChildren github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MemoryContextDeleteChildren
func F_MemoryContextDeleteChildren(m *base.Module, l0 int32)
//go:linkname F_MemoryContextDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextDelete
func F_MemoryContextDelete(m *base.Module, l0 int32)
//go:linkname F_GetMemoryChunkContext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetMemoryChunkContext
func F_GetMemoryChunkContext(m *base.Module, l0 int32) int32
//go:linkname F_GetMemoryChunkSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetMemoryChunkSpace
func F_GetMemoryChunkSpace(m *base.Module, l0 int32) int32
//go:linkname F_MemoryContextStats github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextStats
func F_MemoryContextStats(m *base.Module, l0 int32)
//go:linkname F_MemoryContextAllocationFailure github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextAllocationFailure
func F_MemoryContextAllocationFailure(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MemoryContextSizeFailure github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextSizeFailure
func F_MemoryContextSizeFailure(m *base.Module, l0 int32)
//go:linkname F_MemoryContextAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextAlloc
func F_MemoryContextAlloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextAllocExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextAllocExtended
func F_MemoryContextAllocExtended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ProcessLogMemoryContextInterrupt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcessLogMemoryContextInterrupt
func F_ProcessLogMemoryContextInterrupt(m *base.Module)
//go:linkname F_palloc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_palloc
func F_palloc(m *base.Module, l0 int32) int32
//go:linkname F_palloc0 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_palloc0
func F_palloc0(m *base.Module, l0 int32) int32
//go:linkname F_palloc_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_palloc_extended
func F_palloc_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_palloc_aligned github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_palloc_aligned
func F_palloc_aligned(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pfree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pfree
func F_pfree(m *base.Module, l0 int32)
//go:linkname F_repalloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_repalloc
func F_repalloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_repalloc0 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_repalloc0
func F_repalloc0(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MemoryContextAllocHuge github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextAllocHuge
func F_MemoryContextAllocHuge(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_repalloc_huge github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_repalloc_huge
func F_repalloc_huge(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextStrdup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MemoryContextStrdup
func F_MemoryContextStrdup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pstrdup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pstrdup
func F_pstrdup(m *base.Module, l0 int32) int32
//go:linkname F_pnstrdup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pnstrdup
func F_pnstrdup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetPortalByName github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetPortalByName
func F_GetPortalByName(m *base.Module, l0 int32) int32
//go:linkname F_CreatePortal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreatePortal
func F_CreatePortal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_PortalDrop github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PortalDrop
func F_PortalDrop(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateNewPortal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CreateNewPortal
func F_CreateNewPortal(m *base.Module) int32
//go:linkname F_MarkPortalActive github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MarkPortalActive
func F_MarkPortalActive(m *base.Module, l0 int32)
//go:linkname F_MarkPortalFailed github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MarkPortalFailed
func F_MarkPortalFailed(m *base.Module, l0 int32)
//go:linkname F_PortalHashTableDeleteAll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PortalHashTableDeleteAll
func F_PortalHashTableDeleteAll(m *base.Module)
//go:linkname F_ResourceOwnerEnlarge github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResourceOwnerEnlarge
func F_ResourceOwnerEnlarge(m *base.Module, l0 int32)
//go:linkname F_ResourceOwnerRemember github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResourceOwnerRemember
func F_ResourceOwnerRemember(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ResourceOwnerForget github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResourceOwnerForget
func F_ResourceOwnerForget(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CreateAuxProcessResourceOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CreateAuxProcessResourceOwner
func F_CreateAuxProcessResourceOwner(m *base.Module)
//go:linkname F_ReleaseAuxProcessResources github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReleaseAuxProcessResources
func F_ReleaseAuxProcessResources(m *base.Module, l0 int32)
//go:linkname F_LogicalTapeClose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LogicalTapeClose
func F_LogicalTapeClose(m *base.Module, l0 int32)
//go:linkname F_LogicalTapeWrite github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LogicalTapeWrite
func F_LogicalTapeWrite(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ltsReadFillBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltsReadFillBuffer
func F_ltsReadFillBuffer(m *base.Module, l0 int32) int32
//go:linkname F_LogicalTapeBackspace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LogicalTapeBackspace
func F_LogicalTapeBackspace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LogicalTapeSetBlocks github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LogicalTapeSetBlocks
func F_LogicalTapeSetBlocks(m *base.Module, l0 int32) int64
//go:linkname F_qsort_interruptible github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_qsort_interruptible
func F_qsort_interruptible(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_PrepareSortSupportComparisonShim github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PrepareSortSupportComparisonShim
func F_PrepareSortSupportComparisonShim(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PrepareSortSupportFromOrderingOp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PrepareSortSupportFromOrderingOp
func F_PrepareSortSupportFromOrderingOp(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplesort_end github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_end
func F_tuplesort_end(m *base.Module, l0 int32)
//go:linkname F_tuplesort_performsort github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_performsort
func F_tuplesort_performsort(m *base.Module, l0 int32)
//go:linkname F_getlen github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getlen
func F_getlen(m *base.Module, l0 int32) int32
//go:linkname F_tuplesort_skiptuples github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_skiptuples
func F_tuplesort_skiptuples(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_tuplesort_rescan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_rescan
func F_tuplesort_rescan(m *base.Module, l0 int32)
//go:linkname F_tuplesort_begin_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_begin_heap
func F_tuplesort_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_tuplesort_begin_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_begin_datum
func F_tuplesort_begin_datum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_tuplesort_puttupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_puttupleslot
func F_tuplesort_puttupleslot(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplesort_putindextuplevalues github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplesort_putindextuplevalues
func F_tuplesort_putindextuplevalues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_tuplesort_putbrintuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_putbrintuple
func F_tuplesort_putbrintuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_putdatum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_putdatum
func F_tuplesort_putdatum(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_getdatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_getdatum
func F_tuplesort_getdatum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_tuplestore_begin_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_begin_heap
func F_tuplestore_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplestore_alloc_read_pointer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplestore_alloc_read_pointer
func F_tuplestore_alloc_read_pointer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tuplestore_select_read_pointer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplestore_select_read_pointer
func F_tuplestore_select_read_pointer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_puttuple_common github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tuplestore_puttuple_common
func F_tuplestore_puttuple_common(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_putvalues github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplestore_putvalues
func F_tuplestore_putvalues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_tuplestore_gettupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_gettupleslot
func F_tuplestore_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_tuplestore_advance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplestore_advance
func F_tuplestore_advance(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tuplestore_skiptuples github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_skiptuples
func F_tuplestore_skiptuples(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_tuplestore_rescan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_rescan
func F_tuplestore_rescan(m *base.Module, l0 int32)
//go:linkname F_GetComboCommandId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetComboCommandId
func F_GetComboCommandId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetTransactionSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetTransactionSnapshot
func F_GetTransactionSnapshot(m *base.Module) int32
//go:linkname F_InvalidateCatalogSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InvalidateCatalogSnapshot
func F_InvalidateCatalogSnapshot(m *base.Module)
//go:linkname F_GetCatalogSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetCatalogSnapshot
func F_GetCatalogSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_PushActiveSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PushActiveSnapshot
func F_PushActiveSnapshot(m *base.Module, l0 int32)
//go:linkname F_RegisterSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RegisterSnapshot
func F_RegisterSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_UnregisterSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_UnregisterSnapshot
func F_UnregisterSnapshot(m *base.Module, l0 int32)
//go:linkname F_ExportSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExportSnapshot
func F_ExportSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_parseIntFromText github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parseIntFromText
func F_parseIntFromText(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_parseXidFromText github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parseXidFromText
func F_parseXidFromText(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SetTransactionSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SetTransactionSnapshot
func F_SetTransactionSnapshot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RestoreSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RestoreSnapshot
func F_RestoreSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_RestoreTransactionSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RestoreTransactionSnapshot
func F_RestoreTransactionSnapshot(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tzload github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tzload
func F_tzload(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tzparse github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tzparse
func F_tzparse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_localtime github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_localtime
func F_pg_localtime(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_interpret_timezone_abbrev github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_interpret_timezone_abbrev
func F_pg_interpret_timezone_abbrev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_tzset github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_tzset
func F_pg_tzset(m *base.Module, l0 int32) int32
//go:linkname F_pg_tzset_offset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_tzset_offset
func F_pg_tzset_offset(m *base.Module, l0 int32) int32
//go:linkname F_pg_strftime github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_strftime
func F_pg_strftime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_jit_compile_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jit_compile_expr
func F_jit_compile_expr(m *base.Module, l0 int32) int32
//go:linkname F_binaryheap_allocate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_binaryheap_allocate
func F_binaryheap_allocate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_blockreftable_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_blockreftable_insert
func F_blockreftable_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_update_controlfile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_update_controlfile
func F_update_controlfile(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_dirent_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_dirent_type
func F_get_dirent_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_bytes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hash_bytes
func F_hash_bytes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_getaddrinfo_all github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_getaddrinfo_all
func F_pg_getaddrinfo_all(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_json_lex_number github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_json_lex_number
func F_json_lex_number(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_makeJsonLexContextCstringLen github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeJsonLexContextCstringLen
func F_makeJsonLexContextCstringLen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_freeJsonLexContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_freeJsonLexContext
func F_freeJsonLexContext(m *base.Module, l0 int32)
//go:linkname F_pg_parse_json github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_parse_json
func F_pg_parse_json(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_json_lex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_json_lex
func F_json_lex(m *base.Module, l0 int32) int32
//go:linkname F_parse_array_element github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_parse_array_element
func F_parse_array_element(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ScanKeywordLookup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ScanKeywordLookup
func F_ScanKeywordLookup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_md5_binary github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_md5_binary
func F_pg_md5_binary(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_md5_encrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_md5_encrypt
func F_pg_md5_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_json_parse_manifest_incremental_chunk github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_json_parse_manifest_incremental_chunk
func F_json_parse_manifest_incremental_chunk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pg_get_line_append github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_get_line_append
func F_pg_get_line_append(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_psprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_psprintf
func F_psprintf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetDatabasePath github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetDatabasePath
func F_GetDatabasePath(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rmtree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_rmtree
func F_rmtree(m *base.Module, l0 int32) int32
//go:linkname F_pg_saslprep github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_saslprep
func F_pg_saslprep(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_scram_SaltedPassword github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_scram_SaltedPassword
func F_scram_SaltedPassword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_scram_H github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scram_H
func F_scram_H(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_scram_ServerKey github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_scram_ServerKey
func F_scram_ServerKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeStringInfo
func F_makeStringInfo(m *base.Module) int32
//go:linkname F_initStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initStringInfo
func F_initStringInfo(m *base.Module, l0 int32)
//go:linkname F_resetStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_resetStringInfo
func F_resetStringInfo(m *base.Module, l0 int32)
//go:linkname F_appendStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_appendStringInfo
func F_appendStringInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_enlargeStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_enlargeStringInfo
func F_enlargeStringInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_appendStringInfoVA github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendStringInfoVA
func F_appendStringInfoVA(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_appendStringInfoString github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendStringInfoString
func F_appendStringInfoString(m *base.Module, l0 int32, l1 int32)
//go:linkname F_appendBinaryStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendBinaryStringInfo
func F_appendBinaryStringInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_appendStringInfoSpaces github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendStringInfoSpaces
func F_appendStringInfoSpaces(m *base.Module, l0 int32, l1 int32)
//go:linkname F_appendBinaryStringInfoNT github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_appendBinaryStringInfoNT
func F_appendBinaryStringInfoNT(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_wait_result_to_str github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_wait_result_to_str
func F_wait_result_to_str(m *base.Module, l0 int32) int32
//go:linkname F_pg_encoding_mblen github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_encoding_mblen
func F_pg_encoding_mblen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_encoding_verifymbchar github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_encoding_verifymbchar
func F_pg_encoding_verifymbchar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_cryptohash_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_cryptohash_create
func F_pg_cryptohash_create(m *base.Module, l0 int32) int32
//go:linkname F_pg_cryptohash_init github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_cryptohash_init
func F_pg_cryptohash_init(m *base.Module, l0 int32) int32
//go:linkname F_pg_cryptohash_final github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_cryptohash_final
func F_pg_cryptohash_final(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_cryptohash_free github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_cryptohash_free
func F_pg_cryptohash_free(m *base.Module, l0 int32)
//go:linkname F_pg_cryptohash_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_cryptohash_error
func F_pg_cryptohash_error(m *base.Module, l0 int32) int32
//go:linkname F_pg_hmac_create github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_hmac_create
func F_pg_hmac_create(m *base.Module, l0 int32) int32
//go:linkname F_pg_hmac_init github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_hmac_init
func F_pg_hmac_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_hmac_final github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_hmac_final
func F_pg_hmac_final(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_hmac_free github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_hmac_free
func F_pg_hmac_free(m *base.Module, l0 int32)
//go:linkname F_pg_inet_net_ntop github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_inet_net_ntop
func F_pg_inet_net_ntop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_join_path_components github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_join_path_components
func F_join_path_components(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_canonicalize_path_enc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_canonicalize_path_enc
func F_canonicalize_path_enc(m *base.Module, l0 int32)
//go:linkname F_make_relative_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_relative_path
func F_make_relative_path(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_usleep github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_usleep
func F_pg_usleep(m *base.Module, l0 int32)
//go:linkname F_pg_toupper github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_toupper
func F_pg_toupper(m *base.Module, l0 int32) int32
//go:linkname F_pg_qsort_med3 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_qsort_med3
func F_pg_qsort_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_qsort_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_qsort_arg
func F_qsort_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_dostr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dostr
func F_dostr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_fmtint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fmtint
func F_fmtint(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_dopr_outchmulti github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dopr_outchmulti
func F_dopr_outchmulti(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_leading_pad github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_leading_pad
func F_leading_pad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pg_snprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_snprintf
func F_pg_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_sprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_sprintf
func F_pg_sprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_printf github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_printf
func F_pg_printf(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_strerror_r github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_strerror_r
func F_pg_strerror_r(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgl_pclose github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgl_pclose
func F_pgl_pclose(m *base.Module, l0 int32) int32
//go:linkname F_pgl_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgl_exit
func F_pgl_exit(m *base.Module, l0 int32)
//go:linkname F_pgl_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgl_recv
func F_pgl_recv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pgl_send github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgl_send
func F_pgl_send(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_plpgsql_build_datatype github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_plpgsql_build_datatype
func F_plpgsql_build_datatype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_plpgsql_parse_word github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_parse_word
func F_plpgsql_parse_word(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_plpgsql_parse_dblword github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_parse_dblword
func F_plpgsql_parse_dblword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_plpgsql_build_recfield github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plpgsql_build_recfield
func F_plpgsql_build_recfield(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_parse_wordtype github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_parse_wordtype
func F_plpgsql_parse_wordtype(m *base.Module, l0 int32) int32
//go:linkname F_plpgsql_parse_wordrowtype github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_parse_wordrowtype
func F_plpgsql_parse_wordrowtype(m *base.Module, l0 int32) int32
//go:linkname F_plpgsql_adddatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_plpgsql_adddatum
func F_plpgsql_adddatum(m *base.Module, l0 int32)
//go:linkname F_plpgsql_build_record github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_build_record
func F_plpgsql_build_record(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_plpgsql_build_datatype_arrayof github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_plpgsql_build_datatype_arrayof
func F_plpgsql_build_datatype_arrayof(m *base.Module, l0 int32) int32
//go:linkname F_plpgsql_recognize_err_condition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_recognize_err_condition
func F_plpgsql_recognize_err_condition(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_parse_err_condition github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_plpgsql_parse_err_condition
func F_plpgsql_parse_err_condition(m *base.Module, l0 int32) int32
//go:linkname F_plpgsql_add_initdatums github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_add_initdatums
func F_plpgsql_add_initdatums(m *base.Module, l0 int32) int32
//go:linkname F_assign_simple_var github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_assign_simple_var
func F_assign_simple_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_exec_move_row github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_exec_move_row
func F_exec_move_row(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_exec_stmt_block github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_exec_stmt_block
func F_exec_stmt_block(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_exec_cast_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_cast_value
func F_exec_cast_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_plpgsql_create_econtext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plpgsql_create_econtext
func F_plpgsql_create_econtext(m *base.Module, l0 int32)
//go:linkname F_exec_eval_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_exec_eval_datum
func F_exec_eval_datum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_exec_assign_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_assign_value
func F_exec_assign_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_exec_assign_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_assign_expr
func F_exec_assign_expr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_plpgsql_exec_get_datum_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_plpgsql_exec_get_datum_type
func F_plpgsql_exec_get_datum_type(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_instantiate_empty_record_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_instantiate_empty_record_variable
func F_instantiate_empty_record_variable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_plpgsql_exec_get_datum_type_info github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_exec_get_datum_type_info
func F_plpgsql_exec_get_datum_type_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_exec_prepare_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_prepare_plan
func F_exec_prepare_plan(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_exec_eval_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exec_eval_expr
func F_exec_eval_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_exec_run_select github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exec_run_select
func F_exec_run_select(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_exec_check_assignable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_exec_check_assignable
func F_exec_check_assignable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_exec_for_query github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_exec_for_query
func F_exec_for_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_exec_stmt_execsql github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_exec_stmt_execsql
func F_exec_stmt_execsql(m *base.Module, l0 int32, l1 int32)
//go:linkname F_plpgsql_fulfill_promise github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_fulfill_promise
func F_plpgsql_fulfill_promise(m *base.Module, l0 int32, l1 int32)
//go:linkname F_exec_init_tuple_store github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_init_tuple_store
func F_exec_init_tuple_store(m *base.Module, l0 int32)
//go:linkname F_exec_eval_using_params github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_exec_eval_using_params
func F_exec_eval_using_params(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_format_preparedparamsdata github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_format_preparedparamsdata
func F_format_preparedparamsdata(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_ns_push github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_ns_push
func F_plpgsql_ns_push(m *base.Module, l0 int32, l1 int32)
//go:linkname F_plpgsql_ns_additem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_ns_additem
func F_plpgsql_ns_additem(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_check_labels github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_labels
func F_check_labels(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_check_assignable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_check_assignable
func F_check_assignable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_word_is_not_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_word_is_not_variable
func F_word_is_not_variable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_read_cursor_args github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_read_cursor_args
func F_read_cursor_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_make_execsql_stmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_execsql_stmt
func F_make_execsql_stmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_read_into_target github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_read_into_target
func F_read_into_target(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_complete_direction github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_complete_direction
func F_complete_direction(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_plpgsql_push_back_token github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_plpgsql_push_back_token
func F_plpgsql_push_back_token(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_plpgsql_append_source_text github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_append_source_text
func F_plpgsql_append_source_text(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_plpgsql_peek github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_plpgsql_peek
func F_plpgsql_peek(m *base.Module, l0 int32) int32
//go:linkname F_plpgsql_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_yyerror
func F_plpgsql_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_SN_create_env github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SN_create_env
func F_SN_create_env(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_r_en_ending_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_r_en_ending_1
func F_r_en_ending_1(m *base.Module, l0 int32) int32
//go:linkname F_r_e_ending_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_r_e_ending_1
func F_r_e_ending_1(m *base.Module, l0 int32) int32
//go:linkname F_r_remove_suffix_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_r_remove_suffix_1
func F_r_remove_suffix_1(m *base.Module, l0 int32) int32
//go:linkname F_r_remove_second_order_prefix_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_remove_second_order_prefix_1
func F_r_remove_second_order_prefix_1(m *base.Module, l0 int32) int32
//go:linkname F_r_check_vowel_harmony github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_check_vowel_harmony
func F_r_check_vowel_harmony(m *base.Module, l0 int32) int32
//go:linkname F_r_mark_ysA github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_r_mark_ysA
func F_r_mark_ysA(m *base.Module, l0 int32) int32
//go:linkname F_r_mark_ymUs_ github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_r_mark_ymUs_
func F_r_mark_ymUs_(m *base.Module, l0 int32) int32
//go:linkname F_r_mark_yUm github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_r_mark_yUm
func F_r_mark_yUm(m *base.Module, l0 int32) int32
//go:linkname F_r_mark_lAr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_mark_lAr
func F_r_mark_lAr(m *base.Module, l0 int32) int32
//go:linkname F_r_stem_suffix_chain_before_ki github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_r_stem_suffix_chain_before_ki
func F_r_stem_suffix_chain_before_ki(m *base.Module, l0 int32) int32
//go:linkname F_r_mark_possessives github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_r_mark_possessives
func F_r_mark_possessives(m *base.Module, l0 int32) int32
//go:linkname F_skip_b_utf8 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_skip_b_utf8
func F_skip_b_utf8(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_in_grouping_b_U github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_in_grouping_b_U
func F_in_grouping_b_U(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_in_grouping_b github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_in_grouping_b
func F_in_grouping_b(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_find_among github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_find_among
func F_find_among(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_slice_from_s github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_slice_from_s
func F_slice_from_s(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_slice_del github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_slice_del
func F_slice_del(m *base.Module, l0 int32) int32
//go:linkname F_BinarySearchRange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BinarySearchRange
func F_BinarySearchRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_px_crypt_des github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_px_crypt_des
func F_px_crypt_des(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__crypt_gensalt_sha github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__crypt_gensalt_sha
func F__crypt_gensalt_sha(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_px_crypt_md5 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_px_crypt_md5
func F_px_crypt_md5(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_mbuf_free github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mbuf_free
func F_mbuf_free(m *base.Module, l0 int32) int32
//go:linkname F_mbuf_append github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_mbuf_append
func F_mbuf_append(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mbuf_create github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mbuf_create
func F_mbuf_create(m *base.Module, l0 int32) int32
//go:linkname F_pullf_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pullf_create
func F_pullf_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pullf_free github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pullf_free
func F_pullf_free(m *base.Module, l0 int32)
//go:linkname F_pullf_read github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pullf_read
func F_pullf_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pullf_read_max github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pullf_read_max
func F_pullf_read_max(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pushf_write github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pushf_write
func F_pushf_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgp_armor_encode github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgp_armor_encode
func F_pgp_armor_encode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_pgp_cfb_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgp_cfb_create
func F_pgp_cfb_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_pgp_cfb_free github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgp_cfb_free
func F_pgp_cfb_free(m *base.Module, l0 int32)
//go:linkname F_cfb_process github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cfb_process
func F_cfb_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_pgp_cfb_decrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgp_cfb_decrypt
func F_pgp_cfb_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pgp_parse_pkt_hdr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgp_parse_pkt_hdr
func F_pgp_parse_pkt_hdr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pgp_expect_packet_end github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgp_expect_packet_end
func F_pgp_expect_packet_end(m *base.Module, l0 int32) int32
//go:linkname F_process_data_packets github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_process_data_packets
func F_process_data_packets(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pgp_encrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_encrypt
func F_pgp_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgp_get_keyid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_get_keyid
func F_pgp_get_keyid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgp_elgamal_decrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_elgamal_decrypt
func F_pgp_elgamal_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pgp_rsa_decrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgp_rsa_decrypt
func F_pgp_rsa_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgp_mpi_create github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_mpi_create
func F_pgp_mpi_create(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgp_mpi_read github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_mpi_read
func F_pgp_mpi_read(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_init_work github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_init_work
func F_init_work(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_decrypt_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_decrypt_internal
func F_decrypt_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_pgp_key_free github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgp_key_free
func F_pgp_key_free(m *base.Module, l0 int32)
//go:linkname F_calc_key_id github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_calc_key_id
func F_calc_key_id(m *base.Module, l0 int32) int32
//go:linkname F_pgp_set_pubkey github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_set_pubkey
func F_pgp_set_pubkey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pgp_s2k_read github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgp_s2k_read
func F_pgp_s2k_read(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgp_s2k_process github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgp_s2k_process
func F_pgp_s2k_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pgp_free github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgp_free
func F_pgp_free(m *base.Module, l0 int32) int32
//go:linkname F_px_gen_salt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_px_gen_salt
func F_px_gen_salt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_px_find_hmac github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_px_find_hmac
func F_px_find_hmac(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_px_THROW_ERROR github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_px_THROW_ERROR
func F_px_THROW_ERROR(m *base.Module, l0 int32)
//go:linkname F_px_debug github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_px_debug
func F_px_debug(m *base.Module, l0 int32, l1 int32)
//go:linkname F_px_find_combo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_px_find_combo
func F_px_find_combo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_internal_citext_pattern_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_internal_citext_pattern_cmp
func F_internal_citext_pattern_cmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generate_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_trgm
func F_generate_trgm(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_trigrams github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_trigrams
func F_make_trigrams(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_cnt_sml github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_cnt_sml
func F_cnt_sml(m *base.Module, l0 int32, l1 int32, l2 int32) float32
//go:linkname F_calc_word_similarity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_calc_word_similarity
func F_calc_word_similarity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float32
//go:linkname F_hstoreUpgrade github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstoreUpgrade
func F_hstoreUpgrade(m *base.Module, l0 int32) int32
//go:linkname F_hstoreUniquePairs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstoreUniquePairs
func F_hstoreUniquePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hstorePairs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hstorePairs
func F_hstorePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hstore_defined github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hstore_defined
func F_hstore_defined(m *base.Module, l0 int32) int32
//go:linkname F_hstore_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_delete
func F_hstore_delete(m *base.Module, l0 int32) int32
//go:linkname F_hstore_contains github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_contains
func F_hstore_contains(m *base.Module, l0 int32) int32
//go:linkname F_hstore_avals github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstore_avals
func F_hstore_avals(m *base.Module, l0 int32) int32
//go:linkname F_hstore_skeys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hstore_skeys
func F_hstore_skeys(m *base.Module, l0 int32) int32
//go:linkname F_array_iterator github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_iterator
func F_array_iterator(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_deparse_lquery github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_deparse_lquery
func F_deparse_lquery(m *base.Module, l0 int32) int32
//go:linkname F_ltree_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ltree_concat
func F_ltree_concat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_queryin github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_queryin
func F_queryin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gbt_num_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_num_compress
func F_gbt_num_compress(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gbt_num_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_num_fetch
func F_gbt_num_fetch(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gbt_num_union github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_num_union
func F_gbt_num_union(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gbt_num_same github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_num_same
func F_gbt_num_same(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gbt_num_distance github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gbt_num_distance
func F_gbt_num_distance(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_gbt_num_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_num_picksplit
func F_gbt_num_picksplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gbt_var_bin_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_var_bin_union
func F_gbt_var_bin_union(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_gbt_var_union github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gbt_var_union
func F_gbt_var_union(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_gbt_var_node_cp_len github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_var_node_cp_len
func F_gbt_var_node_cp_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gbt_var_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_var_picksplit
func F_gbt_var_picksplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_gbt_var_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_var_consistent
func F_gbt_var_consistent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_inner_int_union github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_inner_int_union
func F_inner_int_union(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_copy_intArrayType github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_copy_intArrayType
func F_copy_intArrayType(m *base.Module, l0 int32) int32
//go:linkname F__int_unique github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__int_unique
func F__int_unique(m *base.Module, l0 int32) int32
//go:linkname F_inner_int_inter github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_inner_int_inter
func F_inner_int_inter(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rt__int_size github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_rt__int_size
func F_rt__int_size(m *base.Module, l0 int32, l1 int32)
//go:linkname F_intarray_match_first github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_intarray_match_first
func F_intarray_match_first(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_intarray_add_elem github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_intarray_add_elem
func F_intarray_add_elem(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_int_to_intset github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_int_to_intset
func F_int_to_intset(m *base.Module, l0 int32) int32
//go:linkname F_hemdist_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hemdist_3
func F_hemdist_3(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_cube_union_v0 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_union_v0
func F_cube_union_v0(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cube_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_yyensure_buffer_stack
func F_cube_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_cube_yy_create_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cube_yy_create_buffer
func F_cube_yy_create_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_yy_fatal_error_6 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_yy_fatal_error_6
func F_yy_fatal_error_6(m *base.Module, l0 int32)
//go:linkname F_cube_yyrestart github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cube_yyrestart
func F_cube_yyrestart(m *base.Module, l0 int32, l1 int32)
//go:linkname F_cube_scanner_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_scanner_finish
func F_cube_scanner_finish(m *base.Module, l0 int32)
//go:linkname F_seg_yyparse github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_seg_yyparse
func F_seg_yyparse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_seg_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_seg_yyensure_buffer_stack
func F_seg_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_seg_yy_create_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_seg_yy_create_buffer
func F_seg_yy_create_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_yy_fatal_error_7 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_yy_fatal_error_7
func F_yy_fatal_error_7(m *base.Module, l0 int32)
//go:linkname F_seg_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_seg_yyerror
func F_seg_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_seg_scanner_init github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_seg_scanner_init
func F_seg_scanner_init(m *base.Module, l0 int32, l1 int32)
//go:linkname F_initBloomState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_initBloomState
func F_initBloomState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BloomInitMetapage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BloomInitMetapage
func F_BloomInitMetapage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_string2ean github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_string2ean
func F_string2ean(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ean2isn github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ean2isn
func F_ean2isn(m *base.Module, l0 int64, l1 int32, l2 int32)
//go:linkname F_find_word github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_word
func F_find_word(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstatindex_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstatindex_impl
func F_pgstatindex_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstatginindex_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstatginindex_internal
func F_pgstatginindex_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_relation
func F_pgstat_relation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_uuid_generate_random github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_uuid_generate_random
func F_uuid_generate_random(m *base.Module, l0 int32)
//go:linkname F_uuid_unparse github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_uuid_unparse
func F_uuid_unparse(m *base.Module, l0 int32, l1 int32)
//go:linkname F__emscripten_memcpy_bulkmem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_memcpy_bulkmem
func F__emscripten_memcpy_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__emscripten_memset_bulkmem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_memset_bulkmem
func F__emscripten_memset_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_abort github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_abort
func F_abort(m *base.Module)
//go:linkname F_access github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_access
func F_access(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_R github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_R
func F_R(m *base.Module, l0 float64) float64
//go:linkname F___isspace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___isspace
func F___isspace(m *base.Module, l0 int32) int32
//go:linkname F_bsearch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bsearch
func F_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_chmod github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_chmod
func F_chmod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___clock_gettime github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F___clock_gettime
func F___clock_gettime(m *base.Module, l0 int32, l1 int32)
//go:linkname F_close github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_close
func F_close(m *base.Module, l0 int32) int32
//go:linkname F___cos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___cos
func F___cos(m *base.Module, l0 float64, l1 float64) float64
//go:linkname F___sin github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___sin
func F___sin(m *base.Module, l0 float64, l1 float64, l2 int32) float64
//go:linkname F_dup2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dup2
func F_dup2(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_memmove github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_memmove
func F_memmove(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___time github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___time
func F___time(m *base.Module) int64
//go:linkname F___gettimeofday github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___gettimeofday
func F___gettimeofday(m *base.Module, l0 int32)
//go:linkname F_exp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exp
func F_exp(m *base.Module, l0 float64) float64
//go:linkname F_fclose github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fclose
func F_fclose(m *base.Module, l0 int32) int32
//go:linkname F_fflush github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fflush
func F_fflush(m *base.Module, l0 int32) int32
//go:linkname F___toread github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___toread
func F___toread(m *base.Module, l0 int32) int32
//go:linkname F_do_getc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_do_getc
func F_do_getc(m *base.Module, l0 int32) int32
//go:linkname F_fgets github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fgets
func F_fgets(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fileno github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fileno
func F_fileno(m *base.Module, l0 int32) int32
//go:linkname F_fputc github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fputc
func F_fputc(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fread github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fread
func F_fread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_freopen github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_freopen
func F_freopen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___fseeko_unlocked github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___fseeko_unlocked
func F___fseeko_unlocked(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_fsync github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fsync
func F_fsync(m *base.Module, l0 int32) int32
//go:linkname F_fwrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fwrite
func F_fwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_netlink_msg_to_ifaddr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_netlink_msg_to_ifaddr
func F_netlink_msg_to_ifaddr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fputs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fputs
func F_fputs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_iswalpha github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_iswalpha
func F_iswalpha(m *base.Module, l0 int32) int32
//go:linkname F_iswprint github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_iswprint
func F_iswprint(m *base.Module, l0 int32) int32
//go:linkname F_kill github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_kill
func F_kill(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ldexp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ldexp
func F_ldexp(m *base.Module, l0 float64, l1 int32) float64
//go:linkname F_log github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_log
func F_log(m *base.Module, l0 float64) float64
//go:linkname F___lseek github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___lseek
func F___lseek(m *base.Module, l0 int32, l1 int64, l2 int32) int64
//go:linkname F_memchr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_memchr
func F_memchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mkdir github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mkdir
func F_mkdir(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_open github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_open
func F_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pipe github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pipe
func F_pipe(m *base.Module, l0 int32) int32
//go:linkname F_pow github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pow
func F_pow(m *base.Module, l0 float64, l1 float64) float64
//go:linkname F_pwrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pwrite
func F_pwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_pwritev github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pwritev
func F_pwritev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_read github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_read
func F_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_readlink github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_readlink
func F_readlink(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rename github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_rename
func F_rename(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rmdir github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_rmdir
func F_rmdir(m *base.Module, l0 int32) int32
//go:linkname F_setitimer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_setitimer
func F_setitimer(m *base.Module, l0 int32) int32
//go:linkname F_sigemptyset github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sigemptyset
func F_sigemptyset(m *base.Module, l0 int32)
//go:linkname F_sigprocmask github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sigprocmask
func F_sigprocmask(m *base.Module, l0 int32, l1 int32)
//go:linkname F_snprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_snprintf
func F_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_sscanf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sscanf
func F_sscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strchr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strchr
func F_strchr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___strchrnul github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___strchrnul
func F___strchrnul(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strcmp
func F_strcmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strlcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strlcpy
func F_strlcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strlen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strlen
func F_strlen(m *base.Module, l0 int32) int32
//go:linkname F_strncmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strncmp
func F_strncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strncpy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strncpy
func F_strncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strstr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strstr
func F_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___shgetc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___shgetc
func F___shgetc(m *base.Module, l0 int32) int32
//go:linkname F___floatscan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___floatscan
func F___floatscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_strtox_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strtox_1
func F_strtox_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_strtod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strtod
func F_strtod(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_strtox_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strtox_2
func F_strtox_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64
//go:linkname F_strtol github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strtol
func F_strtol(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_symlink github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_symlink
func F_symlink(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___syscall_ret github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___syscall_ret
func F___syscall_ret(m *base.Module, l0 int32) int32
//go:linkname F_syslog github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_syslog
func F_syslog(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_casemap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_casemap
func F_casemap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___vfprintf_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___vfprintf_internal
func F___vfprintf_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_out github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_out
func F_out(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_write github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_write
func F_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_emscripten_builtin_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_emscripten_builtin_free
func F_emscripten_builtin_free(m *base.Module, l0 int32)
//go:linkname F_emscripten_builtin_realloc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_emscripten_builtin_realloc
func F_emscripten_builtin_realloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dispose_chunk github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dispose_chunk
func F_dispose_chunk(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sbrk github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sbrk
func F_sbrk(m *base.Module, l0 int32) int32
//go:linkname F___addtf3 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___addtf3
func F___addtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___ashlti3 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___ashlti3
func F___ashlti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32)
//go:linkname F___wasm_longjmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___wasm_longjmp
func F___wasm_longjmp(m *base.Module, l0 int32, l1 int32)
//go:linkname F___lshrti3 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___lshrti3
func F___lshrti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32)
//go:linkname F___multi3 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F___multi3
func F___multi3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F_connect github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_connect
func F_connect(m *base.Module, l0 int32) int32
//go:linkname F_freeaddrinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_freeaddrinfo
func F_freeaddrinfo(m *base.Module, l0 int32)
//go:linkname F_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_recv
func F_recv(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_recvfrom github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recvfrom
func F_recvfrom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_send github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_send
func F_send(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sendto github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sendto
func F_sendto(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_socket github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_socket
func F_socket(m *base.Module, l0 int32, l1 int32, l2 int32) int32
