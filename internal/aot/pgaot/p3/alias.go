package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	_ "unsafe"
)
//go:linkname F_brin_build_desc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_brin_build_desc
func F_brin_build_desc(m *base.Module, l0 int32) int32
//go:linkname F_brinsummarize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brinsummarize
func F_brinsummarize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_minmax_multi_get_strategy_procinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_minmax_multi_get_strategy_procinfo
func F_minmax_multi_get_strategy_procinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_brin_range_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_brin_range_deserialize
func F_brin_range_deserialize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_build_expanded_ranges github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_expanded_ranges
func F_build_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_minmax_multi_get_procinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_minmax_multi_get_procinfo
func F_minmax_multi_get_procinfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_reduce_expanded_ranges github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_reduce_expanded_ranges
func F_reduce_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_brin_range_serialize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_brin_range_serialize
func F_brin_range_serialize(m *base.Module, l0 int32) int32
//go:linkname F_brin_initialize_empty_new_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_brin_initialize_empty_new_buffer
func F_brin_initialize_empty_new_buffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_brinRevmapInitialize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_brinRevmapInitialize
func F_brinRevmapInitialize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_brinRevmapTerminate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_brinRevmapTerminate
func F_brinRevmapTerminate(m *base.Module, l0 int32)
//go:linkname F_brin_deform_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_brin_deform_tuple
func F_brin_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_attrmap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_attrmap
func F_make_attrmap(m *base.Module, l0 int32) int32
//go:linkname F_free_attrmap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_free_attrmap
func F_free_attrmap(m *base.Module, l0 int32)
//go:linkname F_build_attrmap_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_attrmap_by_name
func F_build_attrmap_by_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mask_unused_space github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mask_unused_space
func F_mask_unused_space(m *base.Module, l0 int32)
//go:linkname F_detoast_external_attr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_detoast_external_attr
func F_detoast_external_attr(m *base.Module, l0 int32) int32
//go:linkname F_detoast_attr_slice github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_detoast_attr_slice
func F_detoast_attr_slice(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_toast_raw_datum_size github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_toast_raw_datum_size
func F_toast_raw_datum_size(m *base.Module, l0 int32) int32
//go:linkname F_heap_compute_data_size github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_compute_data_size
func F_heap_compute_data_size(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_fill_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_fill_tuple
func F_heap_fill_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_heap_attisnull github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_attisnull
func F_heap_attisnull(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_nocachegetattr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_nocachegetattr
func F_nocachegetattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_getsysattr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_getsysattr
func F_heap_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_copytuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_copytuple
func F_heap_copytuple(m *base.Module, l0 int32) int32
//go:linkname F_heap_copy_tuple_as_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_copy_tuple_as_datum
func F_heap_copy_tuple_as_datum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_heap_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_form_tuple
func F_heap_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_modify_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_modify_tuple
func F_heap_modify_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_heap_modify_tuple_by_cols github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_modify_tuple_by_cols
func F_heap_modify_tuple_by_cols(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_heap_form_minimal_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_form_minimal_tuple
func F_heap_form_minimal_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_heap_copy_minimal_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_copy_minimal_tuple
func F_heap_copy_minimal_tuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_form_tuple
func F_index_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_nocache_index_getattr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_nocache_index_getattr
func F_nocache_index_getattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_index_deform_tuple_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_deform_tuple_internal
func F_index_deform_tuple_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_relation_open github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_relation_open
func F_relation_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_try_relation_open github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_try_relation_open
func F_try_relation_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_relation_openrv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_relation_openrv
func F_relation_openrv(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_relation_close github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_relation_close
func F_relation_close(m *base.Module, l0 int32, l1 int32)
//go:linkname F_add_local_int_reloption github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_add_local_int_reloption
func F_add_local_int_reloption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_untransformRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_untransformRelOptions
func F_untransformRelOptions(m *base.Module, l0 int32) int32
//go:linkname F_build_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_build_reloptions
func F_build_reloptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_heap_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_reloptions
func F_heap_reloptions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_partitioned_table_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_partitioned_table_reloptions
func F_partitioned_table_reloptions(m *base.Module, l0 int32)
//go:linkname F_view_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_view_reloptions
func F_view_reloptions(m *base.Module, l0 int32)
//go:linkname F_index_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_reloptions
func F_index_reloptions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ScanKeyInit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ScanKeyInit
func F_ScanKeyInit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_toast_compress_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_toast_compress_datum
func F_toast_compress_datum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_convert_tuples_by_position github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_convert_tuples_by_position
func F_convert_tuples_by_position(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_execute_attr_map_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_execute_attr_map_tuple
func F_execute_attr_map_tuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_execute_attr_map_slot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_execute_attr_map_slot
func F_execute_attr_map_slot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_execute_attr_map_cols github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_execute_attr_map_cols
func F_execute_attr_map_cols(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CreateTupleDescCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateTupleDescCopy
func F_CreateTupleDescCopy(m *base.Module, l0 int32) int32
//go:linkname F_TupleDescCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TupleDescCopy
func F_TupleDescCopy(m *base.Module, l0 int32, l1 int32)
//go:linkname F_FreeTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeTupleDesc
func F_FreeTupleDesc(m *base.Module, l0 int32)
//go:linkname F_DecrTupleDescRefCount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecrTupleDescRefCount
func F_DecrTupleDescRefCount(m *base.Module, l0 int32)
//go:linkname F_TupleDescInitEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TupleDescInitEntry
func F_TupleDescInitEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_TupleDescInitBuiltinEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TupleDescInitBuiltinEntry
func F_TupleDescInitBuiltinEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GinFormTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GinFormTuple
func F_GinFormTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_writeListPage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_writeListPage
func F_writeListPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ginInsertCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ginInsertCleanup
func F_ginInsertCleanup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ginEntryInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ginEntryInsert
func F_ginEntryInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_ginCompressPostingList github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ginCompressPostingList
func F_ginCompressPostingList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ginPostingListDecode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ginPostingListDecode
func F_ginPostingListDecode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ginPostingListDecodeAllSegments github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ginPostingListDecodeAllSegments
func F_ginPostingListDecodeAllSegments(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_initGinState github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initGinState
func F_initGinState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_index_getattr_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_index_getattr_1
func F_index_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GinNewBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GinNewBuffer
func F_GinNewBuffer(m *base.Module, l0 int32) int32
//go:linkname F_ginExtractEntries github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ginExtractEntries
func F_ginExtractEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_gistplacetopage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistplacetopage
func F_gistplacetopage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32
//go:linkname F_gistGetNodeBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gistGetNodeBuffer
func F_gistGetNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_gistPushItupToNodeBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistPushItupToNodeBuffer
func F_gistPushItupToNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_gistPopItupFromNodeBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistPopItupFromNodeBuffer
func F_gistPopItupFromNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rtree_internal_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_rtree_internal_consistent
func F_rtree_internal_consistent(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_size_box github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_size_box
func F_size_box(m *base.Module, l0 int32) float64
//go:linkname F_gistdentryinit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistdentryinit
func F_gistdentryinit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_gistDeCompressAtt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistDeCompressAtt
func F_gistDeCompressAtt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_gistgetadjusted github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistgetadjusted
func F_gistgetadjusted(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gistpenalty github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gistpenalty
func F_gistpenalty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float32
//go:linkname F_gistFetchTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gistFetchTuple
func F_gistFetchTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_gistcheckpage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gistcheckpage
func F_gistcheckpage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_gistGetFakeLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistGetFakeLSN
func F_gistGetFakeLSN(m *base.Module, l0 int32) int64
//go:linkname F_gistXLogUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gistXLogUpdate
func F_gistXLogUpdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64
//go:linkname F__hash_ovflblkno_to_bitno github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__hash_ovflblkno_to_bitno
func F__hash_ovflblkno_to_bitno(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__hash_getbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_getbuf
func F__hash_getbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__hash_checkpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_checkpage
func F__hash_checkpage(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_HeapCheckForSerializableConflictOut github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HeapCheckForSerializableConflictOut
func F_HeapCheckForSerializableConflictOut(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_heap_getnext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_getnext
func F_heap_getnext(m *base.Module, l0 int32) int32
//go:linkname F_heapgettup_pagemode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heapgettup_pagemode
func F_heapgettup_pagemode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_heap_fetch_next_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_fetch_next_buffer
func F_heap_fetch_next_buffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_heap_getattr_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_getattr_1
func F_heap_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_heap_fetch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_fetch
func F_heap_fetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_heap_hot_search_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_hot_search_buffer
func F_heap_hot_search_buffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_FreeBulkInsertState github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FreeBulkInsertState
func F_FreeBulkInsertState(m *base.Module, l0 int32)
//go:linkname F_heap_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_insert
func F_heap_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_heap_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_delete
func F_heap_delete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_compute_new_xmax_infomask github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_compute_new_xmax_infomask
func F_compute_new_xmax_infomask(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_simple_heap_delete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_simple_heap_delete
func F_simple_heap_delete(m *base.Module, l0 int32, l1 int32)
//go:linkname F_test_lockmode_for_conflict github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_test_lockmode_for_conflict
func F_test_lockmode_for_conflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_HeapTupleSatisfiesVacuum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HeapTupleSatisfiesVacuum
func F_HeapTupleSatisfiesVacuum(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_HeapTupleHeaderIsOnlyLocked github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_HeapTupleHeaderIsOnlyLocked
func F_HeapTupleHeaderIsOnlyLocked(m *base.Module, l0 int32) int32
//go:linkname F_HeapTupleSatisfiesVisibility github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_HeapTupleSatisfiesVisibility
func F_HeapTupleSatisfiesVisibility(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_toast_flatten_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_toast_flatten_tuple
func F_toast_flatten_tuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetVisibilityMapPins github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetVisibilityMapPins
func F_GetVisibilityMapPins(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_heap_page_prune_and_freeze github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_page_prune_and_freeze
func F_heap_page_prune_and_freeze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_visibilitymap_clear github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_visibilitymap_clear
func F_visibilitymap_clear(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_visibilitymap_pin github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_visibilitymap_pin
func F_visibilitymap_pin(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_visibilitymap_set github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_visibilitymap_set
func F_visibilitymap_set(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_visibilitymap_get_status github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_visibilitymap_get_status
func F_visibilitymap_get_status(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetIndexAmRoutineByAmId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetIndexAmRoutineByAmId
func F_GetIndexAmRoutineByAmId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_IndexAmTranslateCompareType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IndexAmTranslateCompareType
func F_IndexAmTranslateCompareType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_RelationGetIndexScan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationGetIndexScan
func F_RelationGetIndexScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_systable_beginscan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_beginscan
func F_systable_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_systable_getnext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_systable_getnext
func F_systable_getnext(m *base.Module, l0 int32) int32
//go:linkname F_systable_endscan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_systable_endscan
func F_systable_endscan(m *base.Module, l0 int32)
//go:linkname F_systable_inplace_update_begin github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_inplace_update_begin
func F_systable_inplace_update_begin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_systable_inplace_update_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_inplace_update_finish
func F_systable_inplace_update_finish(m *base.Module, l0 int32, l1 int32)
//go:linkname F_systable_inplace_update_cancel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_systable_inplace_update_cancel
func F_systable_inplace_update_cancel(m *base.Module, l0 int32)
//go:linkname F_index_open github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_open
func F_index_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_index_insert
func F_index_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_index_insert_cleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_insert_cleanup
func F_index_insert_cleanup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_index_beginscan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_beginscan
func F_index_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_index_endscan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_endscan
func F_index_endscan(m *base.Module, l0 int32)
//go:linkname F_index_restrpos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_restrpos
func F_index_restrpos(m *base.Module, l0 int32)
//go:linkname F_index_parallelrescan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_parallelrescan
func F_index_parallelrescan(m *base.Module, l0 int32)
//go:linkname F_index_getnext_slot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_getnext_slot
func F_index_getnext_slot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_index_getprocinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_getprocinfo
func F_index_getprocinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_btoidvectorcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_btoidvectorcmp
func F_btoidvectorcmp(m *base.Module, l0 int32) int32
//go:linkname F__bt_dedup_finish_pending github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_dedup_finish_pending
func F__bt_dedup_finish_pending(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_swap_posting github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_swap_posting
func F__bt_swap_posting(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_finish_split github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_finish_split
func F__bt_finish_split(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__bt_checkpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_checkpage
func F__bt_checkpage(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_getbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_getbuf
func F__bt_getbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_relbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_relbuf
func F__bt_relbuf(m *base.Module, l0 int32)
//go:linkname F__bt_unlockbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_unlockbuf
func F__bt_unlockbuf(m *base.Module, l0 int32)
//go:linkname F__bt_getmeta github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_getmeta
func F__bt_getmeta(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_relandgetbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_relandgetbuf
func F__bt_relandgetbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_getrootheight github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_getrootheight
func F__bt_getrootheight(m *base.Module, l0 int32) int32
//go:linkname F_btvacuumscan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_btvacuumscan
func F_btvacuumscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F__bt_parallel_primscan_schedule github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_parallel_primscan_schedule
func F__bt_parallel_primscan_schedule(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_compare github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_compare
func F__bt_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_first github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_first
func F__bt_first(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_next github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_next
func F__bt_next(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__bt_start_array_keys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_start_array_keys
func F__bt_start_array_keys(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_start_prim_scan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_start_prim_scan
func F__bt_start_prim_scan(m *base.Module, l0 int32) int32
//go:linkname F__bt_checkkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_checkkeys
func F__bt_checkkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F__bt_scanbehind_checkkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_scanbehind_checkkeys
func F__bt_scanbehind_checkkeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_set_startikey github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_set_startikey
func F__bt_set_startikey(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_end_vacuum_callback github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_end_vacuum_callback
func F__bt_end_vacuum_callback(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_allequalimage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_allequalimage
func F__bt_allequalimage(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_btree_xlog_split github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_btree_xlog_split
func F_btree_xlog_split(m *base.Module, l0 int32, l1 int32)
//go:linkname F_btree_xlog_updates github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_btree_xlog_updates
func F_btree_xlog_updates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__bt_restore_meta github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_restore_meta
func F__bt_restore_meta(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_restore_page github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__bt_restore_page
func F__bt_restore_page(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_spgTestLeafTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_spgTestLeafTuple
func F_spgTestLeafTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_spgGetCache github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_spgGetCache
func F_spgGetCache(m *base.Module, l0 int32) int32
//go:linkname F_spgExtractNodeLabels github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_spgExtractNodeLabels
func F_spgExtractNodeLabels(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_spgvacuumscan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_spgvacuumscan
func F_spgvacuumscan(m *base.Module, l0 int32)
//go:linkname F_try_table_open github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_try_table_open
func F_try_table_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_openrv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_table_openrv
func F_table_openrv(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_openrv_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_table_openrv_extended
func F_table_openrv_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_table_slot_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_table_slot_create
func F_table_slot_create(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_parallelscan_estimate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_table_parallelscan_estimate
func F_table_parallelscan_estimate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_parallelscan_initialize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_table_parallelscan_initialize
func F_table_parallelscan_initialize(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_table_beginscan_parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_table_beginscan_parallel
func F_table_beginscan_parallel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_tuple_get_latest_tid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_table_tuple_get_latest_tid
func F_table_tuple_get_latest_tid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_TransactionTreeSetCommitTsData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TransactionTreeSetCommitTsData
func F_TransactionTreeSetCommitTsData(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32)
//go:linkname F_TransactionIdGetCommitTsData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransactionIdGetCommitTsData
func F_TransactionIdGetCommitTsData(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SetCommitTsLimit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SetCommitTsLimit
func F_SetCommitTsLimit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GenericXLogRegisterBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GenericXLogRegisterBuffer
func F_GenericXLogRegisterBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GenericXLogFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GenericXLogFinish
func F_GenericXLogFinish(m *base.Module, l0 int32)
//go:linkname F_RecordNewMultiXact github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RecordNewMultiXact
func F_RecordNewMultiXact(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_mXactCachePut github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mXactCachePut
func F_mXactCachePut(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_MultiXactIdSetOldestMember github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MultiXactIdSetOldestMember
func F_MultiXactIdSetOldestMember(m *base.Module)
//go:linkname F_SetMultiXactIdLimit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SetMultiXactIdLimit
func F_SetMultiXactIdLimit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_find_multixact_start github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_find_multixact_start
func F_find_multixact_start(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MultiXactSetNextMXact github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MultiXactSetNextMXact
func F_MultiXactSetNextMXact(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetOldestMultiXactId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetOldestMultiXactId
func F_GetOldestMultiXactId(m *base.Module) int32
//go:linkname F_CreateParallelContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateParallelContext
func F_CreateParallelContext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_WaitForParallelWorkersToFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WaitForParallelWorkersToFinish
func F_WaitForParallelWorkersToFinish(m *base.Module, l0 int32)
//go:linkname F_WaitForParallelWorkersToAttach github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WaitForParallelWorkersToAttach
func F_WaitForParallelWorkersToAttach(m *base.Module, l0 int32)
//go:linkname F_DestroyParallelContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DestroyParallelContext
func F_DestroyParallelContext(m *base.Module, l0 int32)
//go:linkname F_RmgrNotFound github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RmgrNotFound
func F_RmgrNotFound(m *base.Module, l0 int32)
//go:linkname F_SimpleLruZeroPage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SimpleLruZeroPage
func F_SimpleLruZeroPage(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_SimpleLruWaitIO github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SimpleLruWaitIO
func F_SimpleLruWaitIO(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SimpleLruReadPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SimpleLruReadPage
func F_SimpleLruReadPage(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_SlruReportIOError github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SlruReportIOError
func F_SlruReportIOError(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_SimpleLruReadPage_ReadOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SimpleLruReadPage_ReadOnly
func F_SimpleLruReadPage_ReadOnly(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_SimpleLruWritePage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SimpleLruWritePage
func F_SimpleLruWritePage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SimpleLruDoesPhysicalPageExist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SimpleLruDoesPhysicalPageExist
func F_SimpleLruDoesPhysicalPageExist(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_SlruScanDirectory github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SlruScanDirectory
func F_SlruScanDirectory(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SlruDeleteSegment github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SlruDeleteSegment
func F_SlruDeleteSegment(m *base.Module, l0 int64)
//go:linkname F_SubTransSetParent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SubTransSetParent
func F_SubTransSetParent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_restoreTimeLineHistoryFiles github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_restoreTimeLineHistoryFiles
func F_restoreTimeLineHistoryFiles(m *base.Module, l0 int32, l1 int32)
//go:linkname F_readTimeLineHistory github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_readTimeLineHistory
func F_readTimeLineHistory(m *base.Module, l0 int32) int32
//go:linkname F_existsTimeLineHistory github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_existsTimeLineHistory
func F_existsTimeLineHistory(m *base.Module, l0 int32) int32
//go:linkname F_findNewestTimeLine github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_findNewestTimeLine
func F_findNewestTimeLine(m *base.Module, l0 int32) int32
//go:linkname F_tliSwitchPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tliSwitchPoint
func F_tliSwitchPoint(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_TransactionIdDidCommit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransactionIdDidCommit
func F_TransactionIdDidCommit(m *base.Module, l0 int32) int32
//go:linkname F_TransactionIdPrecedes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TransactionIdPrecedes
func F_TransactionIdPrecedes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TransactionIdDidAbort github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TransactionIdDidAbort
func F_TransactionIdDidAbort(m *base.Module, l0 int32) int32
//go:linkname F_TransactionIdCommitTree github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TransactionIdCommitTree
func F_TransactionIdCommitTree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_TransactionIdAsyncCommitTree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TransactionIdAsyncCommitTree
func F_TransactionIdAsyncCommitTree(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64)
//go:linkname F_MarkAsPreparingGuts github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MarkAsPreparingGuts
func F_MarkAsPreparingGuts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32)
//go:linkname F_TwoPhaseGetGXact github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TwoPhaseGetGXact
func F_TwoPhaseGetGXact(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_XlogReadTwoPhaseData github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XlogReadTwoPhaseData
func F_XlogReadTwoPhaseData(m *base.Module, l0 int64, l1 int32, l2 int32)
//go:linkname F_RemoveTwoPhaseFile github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RemoveTwoPhaseFile
func F_RemoveTwoPhaseFile(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ProcessTwoPhaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ProcessTwoPhaseBuffer
func F_ProcessTwoPhaseBuffer(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_PrepareRedoAdd github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PrepareRedoAdd
func F_PrepareRedoAdd(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32)
//go:linkname F_PrescanPreparedTransactions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PrescanPreparedTransactions
func F_PrescanPreparedTransactions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_StandbyRecoverPreparedTransactions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_StandbyRecoverPreparedTransactions
func F_StandbyRecoverPreparedTransactions(m *base.Module)
//go:linkname F_TwoPhaseTransactionGid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TwoPhaseTransactionGid
func F_TwoPhaseTransactionGid(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReadNextFullTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadNextFullTransactionId
func F_ReadNextFullTransactionId(m *base.Module) int64
//go:linkname F_AdvanceNextFullTransactionIdPastXid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AdvanceNextFullTransactionIdPastXid
func F_AdvanceNextFullTransactionIdPastXid(m *base.Module, l0 int32)
//go:linkname F_SetTransactionIdLimit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SetTransactionIdLimit
func F_SetTransactionIdLimit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetTopTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetTopTransactionId
func F_GetTopTransactionId(m *base.Module) int32
//go:linkname F_GetCurrentCommandId github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetCurrentCommandId
func F_GetCurrentCommandId(m *base.Module, l0 int32) int32
//go:linkname F_CommandCounterIncrement github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CommandCounterIncrement
func F_CommandCounterIncrement(m *base.Module)
//go:linkname F_StartTransactionCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_StartTransactionCommand
func F_StartTransactionCommand(m *base.Module)
//go:linkname F_ShowTransactionStateRec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ShowTransactionStateRec
func F_ShowTransactionStateRec(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CommitTransactionCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CommitTransactionCommand
func F_CommitTransactionCommand(m *base.Module)
//go:linkname F_CleanupTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CleanupTransaction
func F_CleanupTransaction(m *base.Module)
//go:linkname F_AbortTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AbortTransaction
func F_AbortTransaction(m *base.Module)
//go:linkname F_AbortSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AbortSubTransaction
func F_AbortSubTransaction(m *base.Module)
//go:linkname F_CleanupSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CleanupSubTransaction
func F_CleanupSubTransaction(m *base.Module)
//go:linkname F_DefineSavepoint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DefineSavepoint
func F_DefineSavepoint(m *base.Module, l0 int32)
//go:linkname F_PreventInTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreventInTransactionBlock
func F_PreventInTransactionBlock(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BeginTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BeginTransactionBlock
func F_BeginTransactionBlock(m *base.Module)
//go:linkname F_PrepareTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PrepareTransactionBlock
func F_PrepareTransactionBlock(m *base.Module, l0 int32) int32
//go:linkname F_EndTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EndTransactionBlock
func F_EndTransactionBlock(m *base.Module, l0 int32) int32
//go:linkname F_PopTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PopTransaction
func F_PopTransaction(m *base.Module)
//go:linkname F_AbortOutOfAnyTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AbortOutOfAnyTransaction
func F_AbortOutOfAnyTransaction(m *base.Module)
//go:linkname F_xact_redo_abort github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_xact_redo_abort
func F_xact_redo_abort(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_WALInsertLockRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WALInsertLockRelease
func F_WALInsertLockRelease(m *base.Module)
//go:linkname F_WALInsertLockAcquireExclusive github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WALInsertLockAcquireExclusive
func F_WALInsertLockAcquireExclusive(m *base.Module)
//go:linkname F_GetXLogBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetXLogBuffer
func F_GetXLogBuffer(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_XLogFlush github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogFlush
func F_XLogFlush(m *base.Module, l0 int64)
//go:linkname F_UpdateMinRecoveryPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_UpdateMinRecoveryPoint
func F_UpdateMinRecoveryPoint(m *base.Module, l0 int64, l1 int32)
//go:linkname F_XLogFileClose github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_XLogFileClose
func F_XLogFileClose(m *base.Module)
//go:linkname F_XLogFileInit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_XLogFileInit
func F_XLogFileInit(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_XLogFileOpen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_XLogFileOpen
func F_XLogFileOpen(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_XLogFileName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogFileName
func F_XLogFileName(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_issue_xlog_fsync github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_issue_xlog_fsync
func F_issue_xlog_fsync(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_XLogNeedsFlush github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogNeedsFlush
func F_XLogNeedsFlush(m *base.Module, l0 int64) int32
//go:linkname F_InstallXLogFileSegment github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InstallXLogFileSegment
func F_InstallXLogFileSegment(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_CheckXLogRemoved github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckXLogRemoved
func F_CheckXLogRemoved(m *base.Module, l0 int64, l1 int32)
//go:linkname F_RemoveNonParentXlogFiles github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RemoveNonParentXlogFiles
func F_RemoveNonParentXlogFiles(m *base.Module, l0 int64, l1 int32)
//go:linkname F_ReadControlFile github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReadControlFile
func F_ReadControlFile(m *base.Module)
//go:linkname F_CheckRequiredParameterValues github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckRequiredParameterValues
func F_CheckRequiredParameterValues(m *base.Module)
//go:linkname F_UpdateFullPageWrites github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UpdateFullPageWrites
func F_UpdateFullPageWrites(m *base.Module)
//go:linkname F_GetRedoRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetRedoRecPtr
func F_GetRedoRecPtr(m *base.Module) int64
//go:linkname F_GetInsertRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetInsertRecPtr
func F_GetInsertRecPtr(m *base.Module) int64
//go:linkname F_XLogShutdownWalRcv github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogShutdownWalRcv
func F_XLogShutdownWalRcv(m *base.Module)
//go:linkname F_RestoreArchivedFile github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RestoreArchivedFile
func F_RestoreArchivedFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_ExecuteRecoveryCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecuteRecoveryCommand
func F_ExecuteRecoveryCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_XLogArchiveForceDone github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogArchiveForceDone
func F_XLogArchiveForceDone(m *base.Module, l0 int32)
//go:linkname F_XLogArchiveCheckDone github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_XLogArchiveCheckDone
func F_XLogArchiveCheckDone(m *base.Module, l0 int32) int32
//go:linkname F_XLogArchiveIsBusy github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogArchiveIsBusy
func F_XLogArchiveIsBusy(m *base.Module, l0 int32) int32
//go:linkname F_XLogArchiveCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogArchiveCleanup
func F_XLogArchiveCleanup(m *base.Module, l0 int32)
//go:linkname F_build_backup_content github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_backup_content
func F_build_backup_content(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_XLogBeginInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_XLogBeginInsert
func F_XLogBeginInsert(m *base.Module)
//go:linkname F_XLogRegisterBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogRegisterBuffer
func F_XLogRegisterBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogRegisterData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XLogRegisterData
func F_XLogRegisterData(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogRegisterBufData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogRegisterBufData
func F_XLogRegisterBufData(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_log_newpage_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_log_newpage_buffer
func F_log_newpage_buffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_log_newpage_range github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_log_newpage_range
func F_log_newpage_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_XLogReaderAllocate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_XLogReaderAllocate
func F_XLogReaderAllocate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_XLogBeginRead github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogBeginRead
func F_XLogBeginRead(m *base.Module, l0 int32, l1 int64)
//go:linkname F_XLogReadRecord github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogReadRecord
func F_XLogReadRecord(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReadPageInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReadPageInternal
func F_ReadPageInternal(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_report_invalid_record github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_invalid_record
func F_report_invalid_record(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ValidXLogRecordHeader github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ValidXLogRecordHeader
func F_ValidXLogRecordHeader(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) int32
//go:linkname F_WALRead github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WALRead
func F_WALRead(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_XLogRecGetBlockTag github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XLogRecGetBlockTag
func F_XLogRecGetBlockTag(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_RestoreBlockImage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RestoreBlockImage
func F_RestoreBlockImage(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReadRecord github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReadRecord
func F_ReadRecord(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CheckForStandbyTrigger github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckForStandbyTrigger
func F_CheckForStandbyTrigger(m *base.Module) int32
//go:linkname F_rescanLatestTimeLine github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_rescanLatestTimeLine
func F_rescanLatestTimeLine(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_XLogFileRead github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_XLogFileRead
func F_XLogFileRead(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_recoveryPausesHere github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recoveryPausesHere
func F_recoveryPausesHere(m *base.Module, l0 int32)
//go:linkname F_CheckRecoveryConsistency github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckRecoveryConsistency
func F_CheckRecoveryConsistency(m *base.Module)
//go:linkname F_PromoteIsTriggered github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PromoteIsTriggered
func F_PromoteIsTriggered(m *base.Module) int32
//go:linkname F_GetXLogReplayRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetXLogReplayRecPtr
func F_GetXLogReplayRecPtr(m *base.Module, l0 int32) int64
//go:linkname F_error_multiple_recovery_targets github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_error_multiple_recovery_targets
func F_error_multiple_recovery_targets(m *base.Module)
//go:linkname F_XLogReadBufferForRedoExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_XLogReadBufferForRedoExtended
func F_XLogReadBufferForRedoExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_XLogReadBufferExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogReadBufferExtended
func F_XLogReadBufferExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_XLogInitBufferForRedo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogInitBufferForRedo
func F_XLogInitBufferForRedo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CreateFakeRelcacheEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateFakeRelcacheEntry
func F_CreateFakeRelcacheEntry(m *base.Module, l0 int32) int32
//go:linkname F_forget_invalid_pages github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_forget_invalid_pages
func F_forget_invalid_pages(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogReadDetermineTimeline github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogReadDetermineTimeline
func F_XLogReadDetermineTimeline(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)
//go:linkname F_WALReadRaiseError github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WALReadRaiseError
func F_WALReadRaiseError(m *base.Module, l0 int32)
//go:linkname F_AppendStringToManifest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AppendStringToManifest
func F_AppendStringToManifest(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sendFile github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sendFile
func F_sendFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32
//go:linkname F__tarWriteHeader github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__tarWriteHeader
func F__tarWriteHeader(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_bbsink_forward_archive_contents github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_bbsink_forward_archive_contents
func F_bbsink_forward_archive_contents(m *base.Module, l0 int32, l1 int32)
//go:linkname F_bbsink_forward_begin_manifest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_bbsink_forward_begin_manifest
func F_bbsink_forward_begin_manifest(m *base.Module, l0 int32)
//go:linkname F_bbsink_forward_manifest_contents github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bbsink_forward_manifest_contents
func F_bbsink_forward_manifest_contents(m *base.Module, l0 int32, l1 int32)
//go:linkname F_throttle github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_throttle
func F_throttle(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getRelationsInNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getRelationsInNamespace
func F_getRelationsInNamespace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_string_to_privilege github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_string_to_privilege
func F_string_to_privilege(m *base.Module, l0 int32) int64
//go:linkname F_privilege_to_string github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_privilege_to_string
func F_privilege_to_string(m *base.Module, l0 int64) int32
//go:linkname F_restrict_and_check_grant github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_restrict_and_check_grant
func F_restrict_and_check_grant(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int64
//go:linkname F_merge_acl_with_grant github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_merge_acl_with_grant
func F_merge_acl_with_grant(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int32, l7 int32) int32
//go:linkname F_recordExtensionInitPrivWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_recordExtensionInitPrivWorker
func F_recordExtensionInitPrivWorker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecGrant_common github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecGrant_common
func F_ExecGrant_common(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_aclcheck_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_aclcheck_error
func F_aclcheck_error(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_class_aclmask_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_class_aclmask_ext
func F_pg_class_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int64
//go:linkname F_object_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_object_aclcheck
func F_object_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_pg_attribute_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_attribute_aclcheck
func F_pg_attribute_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_pg_attribute_aclmask_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_attribute_aclmask_ext
func F_pg_attribute_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int64
//go:linkname F_pg_attribute_aclcheck_all_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_attribute_aclcheck_all_ext
func F_pg_attribute_aclcheck_all_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32
//go:linkname F_pg_class_aclcheck_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_class_aclcheck_ext
func F_pg_class_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pg_parameter_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_parameter_aclcheck
func F_pg_parameter_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pg_largeobject_aclcheck_snapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_largeobject_aclcheck_snapshot
func F_pg_largeobject_aclcheck_snapshot(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_object_ownercheck github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_object_ownercheck
func F_object_ownercheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_has_bypassrls_privilege github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_has_bypassrls_privilege
func F_has_bypassrls_privilege(m *base.Module, l0 int32) int32
//go:linkname F_recordDependencyOnNewAcl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recordDependencyOnNewAcl
func F_recordDependencyOnNewAcl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetNewOidWithIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetNewOidWithIndex
func F_GetNewOidWithIndex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_findDependentObjects github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_findDependentObjects
func F_findDependentObjects(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_deleteObjectsInList github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_deleteObjectsInList
func F_deleteObjectsInList(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_free_object_addresses github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_free_object_addresses
func F_free_object_addresses(m *base.Module, l0 int32)
//go:linkname F_performMultipleDeletions github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_performMultipleDeletions
func F_performMultipleDeletions(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_recordDependencyOnExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recordDependencyOnExpr
func F_recordDependencyOnExpr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_find_expr_references_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_expr_references_walker
func F_find_expr_references_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_add_exact_object_address github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_exact_object_address
func F_add_exact_object_address(m *base.Module, l0 int32, l1 int32)
//go:linkname F_record_object_address_dependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_record_object_address_dependencies
func F_record_object_address_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_SystemAttributeByName github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SystemAttributeByName
func F_SystemAttributeByName(m *base.Module, l0 int32) int32
//go:linkname F_RemoveStatistics github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RemoveStatistics
func F_RemoveStatistics(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationClearMissing github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationClearMissing
func F_RelationClearMissing(m *base.Module, l0 int32)
//go:linkname F_AddRelationNewConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AddRelationNewConstraints
func F_AddRelationNewConstraints(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_StorePartitionBound github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_StorePartitionBound
func F_StorePartitionBound(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_index_check_primary_key github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_index_check_primary_key
func F_index_check_primary_key(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_index_create github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_create
func F_index_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32) int32
//go:linkname F_index_constraint_create github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_constraint_create
func F_index_constraint_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_FormIndexDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FormIndexDatum
func F_FormIndexDatum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_BuildIndexInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BuildIndexInfo
func F_BuildIndexInfo(m *base.Module, l0 int32) int32
//go:linkname F_index_concurrently_build github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_index_concurrently_build
func F_index_concurrently_build(m *base.Module, l0 int32, l1 int32)
//go:linkname F_IndexGetRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_IndexGetRelation
func F_IndexGetRelation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BuildSpeculativeIndexInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BuildSpeculativeIndexInfo
func F_BuildSpeculativeIndexInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_validate_index github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_validate_index
func F_validate_index(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CatalogCloseIndexes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CatalogCloseIndexes
func F_CatalogCloseIndexes(m *base.Module, l0 int32)
//go:linkname F_CatalogTupleInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CatalogTupleInsert
func F_CatalogTupleInsert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CatalogTupleInsertWithInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CatalogTupleInsertWithInfo
func F_CatalogTupleInsertWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CatalogTuplesMultiInsertWithInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CatalogTuplesMultiInsertWithInfo
func F_CatalogTuplesMultiInsertWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CatalogTupleUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CatalogTupleUpdate
func F_CatalogTupleUpdate(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RangeVarGetRelidExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RangeVarGetRelidExtended
func F_RangeVarGetRelidExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_LookupExplicitNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LookupExplicitNamespace
func F_LookupExplicitNamespace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_recomputeNamespacePath github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recomputeNamespacePath
func F_recomputeNamespacePath(m *base.Module)
//go:linkname F_RangeVarGetAndCheckCreationNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RangeVarGetAndCheckCreationNamespace
func F_RangeVarGetAndCheckCreationNamespace(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_isAnyTempNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isAnyTempNamespace
func F_isAnyTempNamespace(m *base.Module, l0 int32) int32
//go:linkname F_NameListToString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_NameListToString
func F_NameListToString(m *base.Module, l0 int32) int32
//go:linkname F_OpfamilyIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_OpfamilyIsVisibleExt
func F_OpfamilyIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_StatisticsObjIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_StatisticsObjIsVisibleExt
func F_StatisticsObjIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_ts_parser_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_ts_parser_oid
func F_get_ts_parser_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_ts_dict_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_ts_dict_oid
func F_get_ts_dict_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_ts_config_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_ts_config_oid
func F_get_ts_config_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LookupNamespaceNoError github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LookupNamespaceNoError
func F_LookupNamespaceNoError(m *base.Module, l0 int32) int32
//go:linkname F_LookupCreationNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LookupCreationNamespace
func F_LookupCreationNamespace(m *base.Module, l0 int32) int32
//go:linkname F_isTempToastNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_isTempToastNamespace
func F_isTempToastNamespace(m *base.Module, l0 int32) int32
//go:linkname F_get_collation_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_collation_oid
func F_get_collation_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fetch_search_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fetch_search_path
func F_fetch_search_path(m *base.Module, l0 int32) int32
//go:linkname F_RunFunctionExecuteHook github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RunFunctionExecuteHook
func F_RunFunctionExecuteHook(m *base.Module, l0 int32)
//go:linkname F_RunObjectPostAlterHookStr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RunObjectPostAlterHookStr
func F_RunObjectPostAlterHookStr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_object_address github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_object_address
func F_get_object_address(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_getObjectDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getObjectDescription
func F_getObjectDescription(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_textarray_to_strvaluelist github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_textarray_to_strvaluelist
func F_textarray_to_strvaluelist(m *base.Module, l0 int32) int32
//go:linkname F_get_object_catcache_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_object_catcache_oid
func F_get_object_catcache_oid(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_object_attnum_acl
func F_get_object_attnum_acl(m *base.Module, l0 int32) int32
//go:linkname F_get_catalog_object_by_oid_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_catalog_object_by_oid_extended
func F_get_catalog_object_by_oid_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getPublicationSchemaInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getPublicationSchemaInfo
func F_getPublicationSchemaInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getObjectTypeDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getObjectTypeDescription
func F_getObjectTypeDescription(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getRelationIdentity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_getRelationIdentity
func F_getRelationIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_getOpFamilyIdentity github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getOpFamilyIdentity
func F_getOpFamilyIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_partition_parent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_partition_parent
func F_get_partition_parent(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_partition_ancestors github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_partition_ancestors
func F_get_partition_ancestors(m *base.Module, l0 int32) int32
//go:linkname F_index_get_partition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_get_partition
func F_index_get_partition(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_map_partition_varattnos github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_map_partition_varattnos
func F_map_partition_varattnos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_has_partition_attrs github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_has_partition_attrs
func F_has_partition_attrs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_default_partition_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_default_partition_oid
func F_get_default_partition_oid(m *base.Module, l0 int32) int32
//go:linkname F_get_proposed_default_constraint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_proposed_default_constraint
func F_get_proposed_default_constraint(m *base.Module, l0 int32) int32
//go:linkname F_RemoveAttrDefault github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RemoveAttrDefault
func F_RemoveAttrDefault(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetAttrDefaultOid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetAttrDefaultOid
func F_GetAttrDefaultOid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetAttrDefaultColumnAddress github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetAttrDefaultColumnAddress
func F_GetAttrDefaultColumnAddress(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errdetail_relkind_not_supported github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errdetail_relkind_not_supported
func F_errdetail_relkind_not_supported(m *base.Module, l0 int32)
//go:linkname F_findNotNullConstraintAttnum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_findNotNullConstraintAttnum
func F_findNotNullConstraintAttnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_findNotNullConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_findNotNullConstraint
func F_findNotNullConstraint(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ConstraintSetParentConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ConstraintSetParentConstraint
func F_ConstraintSetParentConstraint(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_relation_idx_constraint_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_relation_idx_constraint_oid
func F_get_relation_idx_constraint_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DeconstructFkConstraintRow github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DeconstructFkConstraintRow
func F_DeconstructFkConstraintRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_recordDependencyOn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recordDependencyOn
func F_recordDependencyOn(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_recordMultipleDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recordMultipleDependencies
func F_recordMultipleDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_recordDependencyOnCurrentExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recordDependencyOnCurrentExtension
func F_recordDependencyOnCurrentExtension(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getExtensionOfObject github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getExtensionOfObject
func F_getExtensionOfObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_deleteDependencyRecordsFor github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_deleteDependencyRecordsFor
func F_deleteDependencyRecordsFor(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_changeDependencyFor github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_changeDependencyFor
func F_changeDependencyFor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_changeDependenciesOf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_changeDependenciesOf
func F_changeDependenciesOf(m *base.Module, l0 int32, l1 int32)
//go:linkname F_changeDependenciesOn github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_changeDependenciesOn
func F_changeDependenciesOn(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sequenceIsOwned github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sequenceIsOwned
func F_sequenceIsOwned(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getOwnedSequences github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getOwnedSequences
func F_getOwnedSequences(m *base.Module, l0 int32) int32
//go:linkname F_get_index_constraint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_index_constraint
func F_get_index_constraint(m *base.Module, l0 int32) int32
//go:linkname F_find_inheritance_children github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_inheritance_children
func F_find_inheritance_children(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_inheritance_children_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_inheritance_children_extended
func F_find_inheritance_children_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_find_all_inheritors github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_all_inheritors
func F_find_all_inheritors(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_has_superclass github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_has_superclass
func F_has_superclass(m *base.Module, l0 int32) int32
//go:linkname F_typeInheritsFrom github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_typeInheritsFrom
func F_typeInheritsFrom(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_StoreSingleInheritance github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_StoreSingleInheritance
func F_StoreSingleInheritance(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_DeleteInheritsTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DeleteInheritsTuple
func F_DeleteInheritsTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_LargeObjectExists github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LargeObjectExists
func F_LargeObjectExists(m *base.Module, l0 int32) int32
//go:linkname F_LargeObjectExistsWithSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LargeObjectExistsWithSnapshot
func F_LargeObjectExistsWithSnapshot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_NamespaceCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_NamespaceCreate
func F_NamespaceCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ParameterAclLookup github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ParameterAclLookup
func F_ParameterAclLookup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_is_schema_publication github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_is_schema_publication
func F_is_schema_publication(m *base.Module, l0 int32) int32
//go:linkname F_GetSchemaPublicationRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetSchemaPublicationRelations
func F_GetSchemaPublicationRelations(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetPublicationSchemas github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetPublicationSchemas
func F_GetPublicationSchemas(m *base.Module, l0 int32) int32
//go:linkname F_recordDependencyOnOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_recordDependencyOnOwner
func F_recordDependencyOnOwner(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_changeDependencyOnOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_changeDependencyOnOwner
func F_changeDependencyOnOwner(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_updateAclDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_updateAclDependencies
func F_updateAclDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_deleteSharedDependencyRecordsFor github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_deleteSharedDependencyRecordsFor
func F_deleteSharedDependencyRecordsFor(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetPublicationsStr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetPublicationsStr
func F_GetPublicationsStr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AddSubscriptionRelState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AddSubscriptionRelState
func F_AddSubscriptionRelState(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32)
//go:linkname F_RemoveSubscriptionRel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RemoveSubscriptionRel
func F_RemoveSubscriptionRel(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetSubscriptionRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetSubscriptionRelations
func F_GetSubscriptionRelations(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_moveArrayTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_moveArrayTypeName
func F_moveArrayTypeName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeArrayTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeArrayTypeName
func F_makeArrayTypeName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_toast_table github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_toast_table
func F_create_toast_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_transformStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformStmt
func F_transformStmt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addNSItemForReturning github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_addNSItemForReturning
func F_addNSItemForReturning(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_resolve_aggregate_transtype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_resolve_aggregate_transtype
func F_resolve_aggregate_transtype(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_findTargetlistEntrySQL99 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_findTargetlistEntrySQL99
func F_findTargetlistEntrySQL99(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_addTargetToGroupList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_addTargetToGroupList
func F_addTargetToGroupList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_checkTargetlistEntrySQL92 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_checkTargetlistEntrySQL92
func F_checkTargetlistEntrySQL92(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_coerce_to_target_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_coerce_to_target_type
func F_coerce_to_target_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_find_coercion_pathway github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_coercion_pathway
func F_find_coercion_pathway(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_typeIsOfTypedTable github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_typeIsOfTypedTable
func F_typeIsOfTypedTable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_select_common_type_from_oids github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_select_common_type_from_oids
func F_select_common_type_from_oids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_valid_polymorphic_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_check_valid_polymorphic_signature
func F_check_valid_polymorphic_signature(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_valid_internal_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_check_valid_internal_signature
func F_check_valid_internal_signature(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_TypeCategory github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TypeCategory
func F_TypeCategory(m *base.Module, l0 int32) int32
//go:linkname F_IsPreferredType github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_IsPreferredType
func F_IsPreferredType(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_IsBinaryCoercible github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IsBinaryCoercible
func F_IsBinaryCoercible(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assign_collations_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_assign_collations_walker
func F_assign_collations_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assign_expr_collations github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_assign_expr_collations
func F_assign_expr_collations(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformExpr
func F_transformExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_func_signature_string github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_func_signature_string
func F_func_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_LookupFuncName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LookupFuncName
func F_LookupFuncName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_parsestate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_parsestate
func F_make_parsestate(m *base.Module, l0 int32) int32
//go:linkname F_free_parsestate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_free_parsestate
func F_free_parsestate(m *base.Module, l0 int32)
//go:linkname F_parser_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parser_errposition
func F_parser_errposition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LookupOperName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LookupOperName
func F_LookupOperName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_op_signature_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_op_signature_string
func F_op_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_refnameNamespaceItem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_refnameNamespaceItem
func F_refnameNamespaceItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_GetCTEForRTE github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetCTEForRTE
func F_GetCTEForRTE(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_colNameToVar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_colNameToVar
func F_colNameToVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getRTEPermissionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getRTEPermissionInfo
func F_getRTEPermissionInfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_buildNSItemFromLists github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_buildNSItemFromLists
func F_buildNSItemFromLists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_expandRTE github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expandRTE
func F_expandRTE(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_attnumAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_attnumAttName
func F_attnumAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_attnumTypeId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_attnumTypeId
func F_attnumTypeId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_attnumCollationId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_attnumCollationId
func F_attnumCollationId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FigureColnameInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FigureColnameInternal
func F_FigureColnameInternal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExpandColumnRefStar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExpandColumnRefStar
func F_ExpandColumnRefStar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExpandRowReference github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExpandRowReference
func F_ExpandRowReference(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_resolveTargetListUnknowns github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_resolveTargetListUnknowns
func F_resolveTargetListUnknowns(m *base.Module, l0 int32, l1 int32)
//go:linkname F_markTargetListOrigins github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_markTargetListOrigins
func F_markTargetListOrigins(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LookupTypeNameOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LookupTypeNameOid
func F_LookupTypeNameOid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_typenameType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_typenameType
func F_typenameType(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_typenameTypeId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_typenameTypeId
func F_typenameTypeId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetColumnDefCollation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetColumnDefCollation
func F_GetColumnDefCollation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_typeStringToTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_typeStringToTypeName
func F_typeStringToTypeName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generateClonedIndexStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generateClonedIndexStmt
func F_generateClonedIndexStmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_raw_parser github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_raw_parser
func F_raw_parser(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_scanner_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scanner_yyerror
func F_scanner_yyerror(m *base.Module, l0 int32, l1 int32)
//go:linkname F_scanner_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scanner_errposition
func F_scanner_errposition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_scanner_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_scanner_finish
func F_scanner_finish(m *base.Module, l0 int32)
//go:linkname F_downcase_truncate_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_downcase_truncate_identifier
func F_downcase_truncate_identifier(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_AlterObjectNamespace_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AlterObjectNamespace_internal
func F_AlterObjectNamespace_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_std_typanalyze github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_std_typanalyze
func F_std_typanalyze(m *base.Module, l0 int32) int32
//go:linkname F_AsyncExistsPendingNotify github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AsyncExistsPendingNotify
func F_AsyncExistsPendingNotify(m *base.Module, l0 int32) int32
//go:linkname F_AddEventToPendingNotifies github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AddEventToPendingNotifies
func F_AddEventToPendingNotifies(m *base.Module, l0 int32)
//go:linkname F_asyncQueueAdvanceTail github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_asyncQueueAdvanceTail
func F_asyncQueueAdvanceTail(m *base.Module)
//go:linkname F_mark_index_clustered github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_mark_index_clustered
func F_mark_index_clustered(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_make_new_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_new_heap
func F_make_new_heap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_finish_heap_swap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_finish_heap_swap
func F_finish_heap_swap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_CommentObject github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CommentObject
func F_CommentObject(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CopyLimitPrintoutLength github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopyLimitPrintoutLength
func F_CopyLimitPrintoutLength(m *base.Module, l0 int32) int32
//go:linkname F_CopyReadLine github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CopyReadLine
func F_CopyReadLine(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CopyReadAttributesCSV github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CopyReadAttributesCSV
func F_CopyReadAttributesCSV(m *base.Module, l0 int32) int32
//go:linkname F_CopyAttributeOutText github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CopyAttributeOutText
func F_CopyAttributeOutText(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CopySendTextLikeEndOfRow github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopySendTextLikeEndOfRow
func F_CopySendTextLikeEndOfRow(m *base.Module, l0 int32)
//go:linkname F_get_database_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_database_name
func F_get_database_name(m *base.Module, l0 int32) int32
//go:linkname F_heap_getattr_6 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_getattr_6
func F_heap_getattr_6(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_defGetString github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_defGetString
func F_defGetString(m *base.Module, l0 int32) int32
//go:linkname F_defGetQualifiedName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_defGetQualifiedName
func F_defGetQualifiedName(m *base.Module, l0 int32) int32
//go:linkname F_errorConflictingDefElem github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errorConflictingDefElem
func F_errorConflictingDefElem(m *base.Module, l0 int32, l1 int32)
//go:linkname F_EventTriggerTableRewrite github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_EventTriggerTableRewrite
func F_EventTriggerTableRewrite(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_obtain_object_name_namespace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_obtain_object_name_namespace
func F_obtain_object_name_namespace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_EventTriggerCollectSimpleCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_EventTriggerCollectSimpleCommand
func F_EventTriggerCollectSimpleCommand(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainOnePlan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExplainOnePlan
func F_ExplainOnePlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_show_buffer_usage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_show_buffer_usage
func F_show_buffer_usage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExplainPrintJIT github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExplainPrintJIT
func F_ExplainPrintJIT(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainTargetRel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExplainTargetRel
func F_ExplainTargetRel(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainIndexScanDetails github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExplainIndexScanDetails
func F_ExplainIndexScanDetails(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_show_instrumentation_count github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_show_instrumentation_count
func F_show_instrumentation_count(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_show_expression github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_show_expression
func F_show_expression(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_show_sort_group_keys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_show_sort_group_keys
func F_show_sort_group_keys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_show_incremental_sort_group_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_show_incremental_sort_group_info
func F_show_incremental_sort_group_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_show_wal_usage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_show_wal_usage
func F_show_wal_usage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExplainPropertyList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExplainPropertyList
func F_ExplainPropertyList(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainIndentText github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExplainIndentText
func F_ExplainIndentText(m *base.Module, l0 int32)
//go:linkname F_ExplainPropertyText github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExplainPropertyText
func F_ExplainPropertyText(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainProperty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExplainProperty
func F_ExplainProperty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
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
//go:linkname F_ExplainSaveGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExplainSaveGroup
func F_ExplainSaveGroup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_extension_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_extension_name
func F_get_extension_name(m *base.Module, l0 int32) int32
//go:linkname F_parse_extension_control_file github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parse_extension_control_file
func F_parse_extension_control_file(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_extension_script_filename github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_extension_script_filename
func F_get_extension_script_filename(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_required_extension github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_required_extension
func F_get_required_extension(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_heap_getattr_7 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_heap_getattr_7
func F_heap_getattr_7(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_transformGenericOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformGenericOptions
func F_transformGenericOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_WaitForOlderSnapshots github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WaitForOlderSnapshots
func F_WaitForOlderSnapshots(m *base.Module, l0 int32, l1 int32)
//go:linkname F_DefineIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_DefineIndex
func F_DefineIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)
//go:linkname F_ChooseRelationName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ChooseRelationName
func F_ChooseRelationName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_IndexSetParentIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_IndexSetParentIndex
func F_IndexSetParentIndex(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetDefaultOpClass github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetDefaultOpClass
func F_GetDefaultOpClass(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_typeDepNeeded github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_typeDepNeeded
func F_typeDepNeeded(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_init_sequence github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_init_sequence
func F_init_sequence(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_read_seq_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_seq_tuple
func F_read_seq_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SequenceChangePersistence github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SequenceChangePersistence
func F_SequenceChangePersistence(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateStatistics github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CreateStatistics
func F_CreateStatistics(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_defGetStreamingMode github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_defGetStreamingMode
func F_defGetStreamingMode(m *base.Module, l0 int32) int32
//go:linkname F_publicationListToArray github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_publicationListToArray
func F_publicationListToArray(m *base.Module, l0 int32) int32
//go:linkname F_check_publications github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_publications
func F_check_publications(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fetch_table_list github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fetch_table_list
func F_fetch_table_list(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CheckTableNotInUse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CheckTableNotInUse
func F_CheckTableNotInUse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CloneRowTriggersToPartition github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CloneRowTriggersToPartition
func F_CloneRowTriggersToPartition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CloneForeignKeyConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CloneForeignKeyConstraints
func F_CloneForeignKeyConstraints(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_set_attnotnull github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_attnotnull
func F_set_attnotnull(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetAttributeCompression github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetAttributeCompression
func F_GetAttributeCompression(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetForeignKeyCheckTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetForeignKeyCheckTriggers
func F_GetForeignKeyCheckTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_addFkRecurseReferenced github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_addFkRecurseReferenced
func F_addFkRecurseReferenced(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32)
//go:linkname F_CheckRelationTableSpaceMove github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckRelationTableSpaceMove
func F_CheckRelationTableSpaceMove(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SetRelationTableSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SetRelationTableSpace
func F_SetRelationTableSpace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RenameRelationInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RenameRelationInternal
func F_RenameRelationInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ATPrepCmd github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ATPrepCmd
func F_ATPrepCmd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_find_composite_type_dependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_find_composite_type_dependencies
func F_find_composite_type_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ATRewriteTable github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ATRewriteTable
func F_ATRewriteTable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ATExecSetTableSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ATExecSetTableSpace
func F_ATExecSetTableSpace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_check_of_type github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_check_of_type
func F_check_of_type(m *base.Module, l0 int32)
//go:linkname F_ATExecChangeOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ATExecChangeOwner
func F_ATExecChangeOwner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_AlterTableNamespaceInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AlterTableNamespaceInternal
func F_AlterTableNamespaceInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_AtEOSubXact_on_commit_actions github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AtEOSubXact_on_commit_actions
func F_AtEOSubXact_on_commit_actions(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PartConstraintImpliedByRelConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PartConstraintImpliedByRelConstraint
func F_PartConstraintImpliedByRelConstraint(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ATParseTransformCmd github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ATParseTransformCmd
func F_ATParseTransformCmd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ATExecAddIdentity github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ATExecAddIdentity
func F_ATExecAddIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_ATExecSetIdentity github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ATExecSetIdentity
func F_ATExecSetIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_dropconstraint_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dropconstraint_internal
func F_dropconstraint_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_ATExecSetNotNull github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ATExecSetNotNull
func F_ATExecSetNotNull(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_ATExecSetOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ATExecSetOptions
func F_ATExecSetOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_SetIndexStorageProperties github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SetIndexStorageProperties
func F_SetIndexStorageProperties(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_ATExecAddConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ATExecAddConstraint
func F_ATExecAddConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_ATExecAlterConstrEnforceability github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ATExecAlterConstrEnforceability
func F_ATExecAlterConstrEnforceability(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32
//go:linkname F_ATExecAlterConstrDeferrability github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ATExecAlterConstrDeferrability
func F_ATExecAlterConstrDeferrability(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_CreateInheritance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CreateInheritance
func F_CreateInheritance(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RemoveInheritance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RemoveInheritance
func F_RemoveInheritance(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_drop_parent_dependency github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_drop_parent_dependency
func F_drop_parent_dependency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_relation_mark_replica_identity github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_relation_mark_replica_identity
func F_relation_mark_replica_identity(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ATExecSetRowSecurity github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ATExecSetRowSecurity
func F_ATExecSetRowSecurity(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ATExecForceNoForceRowSecurity github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ATExecForceNoForceRowSecurity
func F_ATExecForceNoForceRowSecurity(m *base.Module, l0 int32, l1 int32)
//go:linkname F_validatePartitionedIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_validatePartitionedIndex
func F_validatePartitionedIndex(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ATExecAddColumn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ATExecAddColumn
func F_ATExecAddColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_GetParentedForeignKeyRefs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetParentedForeignKeyRefs
func F_GetParentedForeignKeyRefs(m *base.Module, l0 int32) int32
//go:linkname F_ATAddCheckNNConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ATAddCheckNNConstraint
func F_ATAddCheckNNConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_QueuePartitionConstraintValidation github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_QueuePartitionConstraintValidation
func F_QueuePartitionConstraintValidation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ATPostAlterTypeParse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ATPostAlterTypeParse
func F_ATPostAlterTypeParse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ATTypedTableRecursion github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ATTypedTableRecursion
func F_ATTypedTableRecursion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_QueueFKConstraintValidation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_QueueFKConstraintValidation
func F_QueueFKConstraintValidation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_RememberConstraintForRebuilding github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RememberConstraintForRebuilding
func F_RememberConstraintForRebuilding(m *base.Module, l0 int32, l1 int32)
//go:linkname F_DropForeignKeyConstraintTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DropForeignKeyConstraintTriggers
func F_DropForeignKeyConstraintTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_tablespace_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_tablespace_oid
func F_get_tablespace_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_tablespace_directories github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_tablespace_directories
func F_create_tablespace_directories(m *base.Module, l0 int32, l1 int32)
//go:linkname F_destroy_tablespace_directories github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_destroy_tablespace_directories
func F_destroy_tablespace_directories(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_PrepareTempTablespaces github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PrepareTempTablespaces
func F_PrepareTempTablespaces(m *base.Module)
//go:linkname F_TriggerSetParentTrigger github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TriggerSetParentTrigger
func F_TriggerSetParentTrigger(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_fastgetattr_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fastgetattr_2
func F_fastgetattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CopyTriggerDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CopyTriggerDesc
func F_CopyTriggerDesc(m *base.Module, l0 int32) int32
//go:linkname F_FreeTriggerDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeTriggerDesc
func F_FreeTriggerDesc(m *base.Module, l0 int32)
//go:linkname F_ExecCallTriggerFunc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecCallTriggerFunc
func F_ExecCallTriggerFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_AfterTriggerSaveEvent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AfterTriggerSaveEvent
func F_AfterTriggerSaveEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)
//go:linkname F_ExecARInsertTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecARInsertTriggers
func F_ExecARInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_GetTupleForTrigger github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetTupleForTrigger
func F_GetTupleForTrigger(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_afterTriggerMarkEvents github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_afterTriggerMarkEvents
func F_afterTriggerMarkEvents(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_afterTriggerInvokeEvents github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_afterTriggerInvokeEvents
func F_afterTriggerInvokeEvents(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_AfterTriggerEndSubXact github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AfterTriggerEndSubXact
func F_AfterTriggerEndSubXact(m *base.Module, l0 int32)
//go:linkname F_get_rels_with_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_rels_with_domain
func F_get_rels_with_domain(m *base.Module, l0 int32) int32
//go:linkname F_AlterDomainAddConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AlterDomainAddConstraint
func F_AlterDomainAddConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_AlterTypeNamespace_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AlterTypeNamespace_oid
func F_AlterTypeNamespace_oid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_check_role_grantor github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_check_role_grantor
func F_check_role_grantor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_vacuum_delay_point github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vacuum_delay_point
func F_vacuum_delay_point(m *base.Module, l0 int32)
//go:linkname F_vac_bulkdel_one_index github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vac_bulkdel_one_index
func F_vac_bulkdel_one_index(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_vac_cleanup_one_index github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_vac_cleanup_one_index
func F_vac_cleanup_one_index(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecInitExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitExpr
func F_ExecInitExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expr_setup_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_expr_setup_walker
func F_expr_setup_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecPushExprSetupSteps github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecPushExprSetupSteps
func F_ExecPushExprSetupSteps(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecInitExprRec github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitExprRec
func F_ExecInitExprRec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExprEvalPushStep github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExprEvalPushStep
func F_ExprEvalPushStep(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecInitExprWithParams github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitExprWithParams
func F_ExecInitExprWithParams(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecPrepareExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecPrepareExpr
func F_ExecPrepareExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecPrepareQual github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecPrepareQual
func F_ExecPrepareQual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecPrepareExprList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecPrepareExprList
func F_ExecPrepareExprList(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecCheck
func F_ExecCheck(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecBuildAggTrans github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBuildAggTrans
func F_ExecBuildAggTrans(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ExecReadyInterpretedExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecReadyInterpretedExpr
func F_ExecReadyInterpretedExpr(m *base.Module, l0 int32)
//go:linkname F_ExecAggCopyTransValue github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecAggCopyTransValue
func F_ExecAggCopyTransValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ExecOpenIndices github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecOpenIndices
func F_ExecOpenIndices(m *base.Module, l0 int32, l1 int32)
//go:linkname F_check_exclusion_or_unique_constraint github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_exclusion_or_unique_constraint
func F_check_exclusion_or_unique_constraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32
//go:linkname F_ExecCheckIndexConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecCheckIndexConstraints
func F_ExecCheckIndexConstraints(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ExecutorStart github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecutorStart
func F_ExecutorStart(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecutorRun github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecutorRun
func F_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_ExecutorFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecutorFinish
func F_ExecutorFinish(m *base.Module, l0 int32)
//go:linkname F_ExecutorEnd github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecutorEnd
func F_ExecutorEnd(m *base.Module, l0 int32)
//go:linkname F_ExecCloseResultRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecCloseResultRelations
func F_ExecCloseResultRelations(m *base.Module, l0 int32)
//go:linkname F_ExecutorRewind github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecutorRewind
func F_ExecutorRewind(m *base.Module, l0 int32)
//go:linkname F_ExecCheckOneRelPerms github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecCheckOneRelPerms
func F_ExecCheckOneRelPerms(m *base.Module, l0 int32) int32
//go:linkname F_ExecPartitionCheckEmitError github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecPartitionCheckEmitError
func F_ExecPartitionCheckEmitError(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecBuildSlotValueDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecBuildSlotValueDescription
func F_ExecBuildSlotValueDescription(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecRelGenVirtualNotNull github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecRelGenVirtualNotNull
func F_ExecRelGenVirtualNotNull(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecUpdateLockMode github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecUpdateLockMode
func F_ExecUpdateLockMode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_EvalPlanQualInit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EvalPlanQualInit
func F_EvalPlanQualInit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ExecParallelFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecParallelFinish
func F_ExecParallelFinish(m *base.Module, l0 int32)
//go:linkname F_ExecEndNode github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecEndNode
func F_ExecEndNode(m *base.Module, l0 int32)
//go:linkname F_tuples_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuples_equal
func F_tuples_equal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_should_refetch_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_should_refetch_tuple
func F_should_refetch_tuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecSimpleRelationDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecSimpleRelationDelete
func F_ExecSimpleRelationDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CheckSubscriptionRelkind github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckSubscriptionRelkind
func F_CheckSubscriptionRelkind(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecScan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecScan
func F_ExecScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecAssignScanProjectionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecAssignScanProjectionInfo
func F_ExecAssignScanProjectionInfo(m *base.Module, l0 int32)
//go:linkname F_ExecScanReScan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecScanReScan
func F_ExecScanReScan(m *base.Module, l0 int32)
//go:linkname F_ExecStoreMinimalTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecStoreMinimalTuple
func F_ExecStoreMinimalTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MakeTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MakeTupleTableSlot
func F_MakeTupleTableSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecResetTupleTable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecResetTupleTable
func F_ExecResetTupleTable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecDropSingleTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecDropSingleTupleTableSlot
func F_ExecDropSingleTupleTableSlot(m *base.Module, l0 int32)
//go:linkname F_ExecSetSlotDescriptor github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecSetSlotDescriptor
func F_ExecSetSlotDescriptor(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecStoreBufferHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecStoreBufferHeapTuple
func F_ExecStoreBufferHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecForceStoreHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecForceStoreHeapTuple
func F_ExecForceStoreHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecFetchSlotMinimalTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecFetchSlotMinimalTuple
func F_ExecFetchSlotMinimalTuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecFetchSlotHeapTupleDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecFetchSlotHeapTupleDatum
func F_ExecFetchSlotHeapTupleDatum(m *base.Module, l0 int32) int32
//go:linkname F_ExecTypeFromTLInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecTypeFromTLInternal
func F_ExecTypeFromTLInternal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecCleanTypeFromTL github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecCleanTypeFromTL
func F_ExecCleanTypeFromTL(m *base.Module, l0 int32) int32
//go:linkname F_BlessTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BlessTupleDesc
func F_BlessTupleDesc(m *base.Module, l0 int32) int32
//go:linkname F_TupleDescGetAttInMetadata github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_TupleDescGetAttInMetadata
func F_TupleDescGetAttInMetadata(m *base.Module, l0 int32) int32
//go:linkname F_HeapTupleHeaderGetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HeapTupleHeaderGetDatum
func F_HeapTupleHeaderGetDatum(m *base.Module, l0 int32) int32
//go:linkname F_begin_tup_output_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_begin_tup_output_tupdesc
func F_begin_tup_output_tupdesc(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_do_tup_output github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_do_tup_output
func F_do_tup_output(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_end_tup_output github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_end_tup_output
func F_end_tup_output(m *base.Module, l0 int32)
//go:linkname F_CreateExecutorState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateExecutorState
func F_CreateExecutorState(m *base.Module) int32
//go:linkname F_FreeExecutorState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FreeExecutorState
func F_FreeExecutorState(m *base.Module, l0 int32)
//go:linkname F_ReScanExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReScanExprContext
func F_ReScanExprContext(m *base.Module, l0 int32)
//go:linkname F_UpdateChangedParamSet github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UpdateChangedParamSet
func F_UpdateChangedParamSet(m *base.Module, l0 int32, l1 int32)
//go:linkname F_executor_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_executor_errposition
func F_executor_errposition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RegisterExprContextCallback github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RegisterExprContextCallback
func F_RegisterExprContextCallback(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_UnregisterExprContextCallback github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UnregisterExprContextCallback
func F_UnregisterExprContextCallback(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecGetTriggerOldSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetTriggerOldSlot
func F_ExecGetTriggerOldSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetRootToChildMap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetRootToChildMap
func F_ExecGetRootToChildMap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetInsertedCols github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecGetInsertedCols
func F_ExecGetInsertedCols(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetAllUpdatedCols github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetAllUpdatedCols
func F_ExecGetAllUpdatedCols(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_sql_stmt_retval github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_sql_stmt_retval
func F_check_sql_stmt_retval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_CreateCommandName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CreateCommandName
func F_CreateCommandName(m *base.Module, l0 int32) int32
//go:linkname F_InstrStartNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_InstrStartNode
func F_InstrStartNode(m *base.Module, l0 int32)
//go:linkname F_InstrStopNode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InstrStopNode
func F_InstrStopNode(m *base.Module, l0 int32, l1 float64)
//go:linkname F_InstrEndLoop github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_InstrEndLoop
func F_InstrEndLoop(m *base.Module, l0 int32)
//go:linkname F_hash_agg_entry_size github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_agg_entry_size
func F_hash_agg_entry_size(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_hash_tables github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_hash_tables
func F_build_hash_tables(m *base.Module, l0 int32)
//go:linkname F_initialize_phase github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_initialize_phase
func F_initialize_phase(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hashagg_reset_spill_state github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hashagg_reset_spill_state
func F_hashagg_reset_spill_state(m *base.Module, l0 int32)
//go:linkname F_ExecReScanBitmapAnd github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecReScanBitmapAnd
func F_ExecReScanBitmapAnd(m *base.Module, l0 int32)
//go:linkname F_ExecParallelHashEnsureBatchAccessors github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecParallelHashEnsureBatchAccessors
func F_ExecParallelHashEnsureBatchAccessors(m *base.Module, l0 int32)
//go:linkname F_ExecHashTableDestroy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecHashTableDestroy
func F_ExecHashTableDestroy(m *base.Module, l0 int32)
//go:linkname F_ExecHashTableDetachBatch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecHashTableDetachBatch
func F_ExecHashTableDetachBatch(m *base.Module, l0 int32)
//go:linkname F_ExecReScanHash github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecReScanHash
func F_ExecReScanHash(m *base.Module, l0 int32)
//go:linkname F_ExecHashTableDetach github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExecHashTableDetach
func F_ExecHashTableDetach(m *base.Module, l0 int32)
//go:linkname F_switchToPresortedPrefixMode github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_switchToPresortedPrefixMode
func F_switchToPresortedPrefixMode(m *base.Module, l0 int32)
//go:linkname F_isCurrentGroup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_isCurrentGroup
func F_isCurrentGroup(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_reorderqueue_pop github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_reorderqueue_pop
func F_reorderqueue_pop(m *base.Module, l0 int32) int32
//go:linkname F_ExecIndexEvalRuntimeKeys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecIndexEvalRuntimeKeys
func F_ExecIndexEvalRuntimeKeys(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_recompute_limits github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recompute_limits
func F_recompute_limits(m *base.Module, l0 int32)
//go:linkname F_MemoizeHash_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoizeHash_hash
func F_MemoizeHash_hash(m *base.Module, l0 int32) int32
//go:linkname F_MemoizeHash_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoizeHash_equal
func F_MemoizeHash_equal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cache_reduce_memory github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cache_reduce_memory
func F_cache_reduce_memory(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecInitGenerated github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecInitGenerated
func F_ExecInitGenerated(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_SeqNext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SeqNext
func F_SeqNext(m *base.Module, l0 int32) int32
//go:linkname F_release_partition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_release_partition
func F_release_partition(m *base.Module, l0 int32)
//go:linkname F_update_frameheadpos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_update_frameheadpos
func F_update_frameheadpos(m *base.Module, l0 int32)
//go:linkname F_WinSetMarkPosition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WinSetMarkPosition
func F_WinSetMarkPosition(m *base.Module, l0 int32, l1 int64)
//go:linkname F_update_frametailpos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_update_frametailpos
func F_update_frametailpos(m *base.Module, l0 int32)
//go:linkname F_update_grouptailpos github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_update_grouptailpos
func F_update_grouptailpos(m *base.Module, l0 int32)
//go:linkname F_WinGetPartitionLocalMemory github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WinGetPartitionLocalMemory
func F_WinGetPartitionLocalMemory(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_WinGetPartitionRowCount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WinGetPartitionRowCount
func F_WinGetPartitionRowCount(m *base.Module, l0 int32) int64
//go:linkname F_WinRowsArePeers github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_WinRowsArePeers
func F_WinRowsArePeers(m *base.Module, l0 int32, l1 int64, l2 int64) int32
//go:linkname F_WinGetFuncArgInFrame github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WinGetFuncArgInFrame
func F_WinGetFuncArgInFrame(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_WinGetFuncArgCurrent github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WinGetFuncArgCurrent
func F_WinGetFuncArgCurrent(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_connect_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SPI_connect_ext
func F_SPI_connect_ext(m *base.Module, l0 int32)
//go:linkname F_SPI_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_finish
func F_SPI_finish(m *base.Module) int32
//go:linkname F_AtEOSubXact_SPI github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AtEOSubXact_SPI
func F_AtEOSubXact_SPI(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SPI_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_execute
func F_SPI_execute(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_freetuptable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_freetuptable
func F_SPI_freetuptable(m *base.Module, l0 int32)
//go:linkname F_SPI_prepare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_prepare
func F_SPI_prepare(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__SPI_make_plan_non_temp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__SPI_make_plan_non_temp
func F__SPI_make_plan_non_temp(m *base.Module, l0 int32) int32
//go:linkname F_SPI_keepplan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SPI_keepplan
func F_SPI_keepplan(m *base.Module, l0 int32)
//go:linkname F_SPI_getvalue github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_getvalue
func F_SPI_getvalue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_palloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_palloc
func F_SPI_palloc(m *base.Module, l0 int32) int32
//go:linkname F_SPI_datumTransfer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SPI_datumTransfer
func F_SPI_datumTransfer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SPI_cursor_open github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_cursor_open
func F_SPI_cursor_open(m *base.Module, l0 int32) int32
//go:linkname F_SPI_cursor_close github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_cursor_close
func F_SPI_cursor_close(m *base.Module, l0 int32)
//go:linkname F_SPI_result_code_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SPI_result_code_string
func F_SPI_result_code_string(m *base.Module, l0 int32) int32
//go:linkname F_SPI_plan_get_cached_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_plan_get_cached_plan
func F_SPI_plan_get_cached_plan(m *base.Module, l0 int32) int32
//go:linkname F_GetForeignDataWrapperExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetForeignDataWrapperExtended
func F_GetForeignDataWrapperExtended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetForeignServerExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetForeignServerExtended
func F_GetForeignServerExtended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetFdwRoutineByRelId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetFdwRoutineByRelId
func F_GetFdwRoutineByRelId(m *base.Module, l0 int32) int32
//go:linkname F_bloom_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bloom_create
func F_bloom_create(m *base.Module, l0 int64, l1 int32, l2 int64) int32
//go:linkname F_bloom_add_element github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bloom_add_element
func F_bloom_add_element(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dshash_create github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dshash_create
func F_dshash_create(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dshash_find_or_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dshash_find_or_insert
func F_dshash_find_or_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dshash_delete_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_dshash_delete_entry
func F_dshash_delete_entry(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dshash_release_lock github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dshash_release_lock
func F_dshash_release_lock(m *base.Module, l0 int32, l1 int32)
//go:linkname F_intset_create github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_intset_create
func F_intset_create(m *base.Module) int32
//go:linkname F_intset_add_member github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_intset_add_member
func F_intset_add_member(m *base.Module, l0 int32, l1 int64)
//go:linkname F_pairingheap_allocate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pairingheap_allocate
func F_pairingheap_allocate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pairingheap_add github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pairingheap_add
func F_pairingheap_add(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pairingheap_remove_first github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pairingheap_remove_first
func F_pairingheap_remove_first(m *base.Module, l0 int32) int32
//go:linkname F_pairingheap_remove github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pairingheap_remove
func F_pairingheap_remove(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sanitize_char_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sanitize_char_1
func F_sanitize_char_1(m *base.Module, l0 int32)
//go:linkname F_parse_scram_secret github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_scram_secret
func F_parse_scram_secret(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_set_authn_id github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_set_authn_id
func F_set_authn_id(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sendAuthRequest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_sendAuthRequest
func F_sendAuthRequest(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AtEOSubXact_LargeObject github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AtEOSubXact_LargeObject
func F_AtEOSubXact_LargeObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_password_type github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_password_type
func F_get_password_type(m *base.Module, l0 int32) int32
//go:linkname F_check_usermap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_usermap
func F_check_usermap(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pq_getkeepalivesidle github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_getkeepalivesidle
func F_pq_getkeepalivesidle(m *base.Module, l0 int32) int32
//go:linkname F_pq_getkeepalivesinterval github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_getkeepalivesinterval
func F_pq_getkeepalivesinterval(m *base.Module, l0 int32) int32
//go:linkname F_pq_getkeepalivescount github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_getkeepalivescount
func F_pq_getkeepalivescount(m *base.Module, l0 int32) int32
//go:linkname F_pq_gettcpusertimeout github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_gettcpusertimeout
func F_pq_gettcpusertimeout(m *base.Module, l0 int32) int32
//go:linkname F_pq_getbyte github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_getbyte
func F_pq_getbyte(m *base.Module) int32
//go:linkname F_pq_recvbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_recvbuf
func F_pq_recvbuf(m *base.Module) int32
//go:linkname F_pq_startmsgread github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_startmsgread
func F_pq_startmsgread(m *base.Module)
//go:linkname F_pq_getmessage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_getmessage
func F_pq_getmessage(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_beginmessage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_beginmessage
func F_pq_beginmessage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pq_sendstring github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_sendstring
func F_pq_sendstring(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pq_endmessage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_endmessage
func F_pq_endmessage(m *base.Module, l0 int32)
//go:linkname F_pq_endmessage_reuse github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_endmessage_reuse
func F_pq_endmessage_reuse(m *base.Module, l0 int32)
//go:linkname F_pq_begintypsend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_begintypsend
func F_pq_begintypsend(m *base.Module, l0 int32)
//go:linkname F_pq_puttextmessage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_puttextmessage
func F_pq_puttextmessage(m *base.Module)
//go:linkname F_pq_getmsgbyte github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_getmsgbyte
func F_pq_getmsgbyte(m *base.Module, l0 int32) int32
//go:linkname F_pq_getmsgint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pq_getmsgint
func F_pq_getmsgint(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_copymsgbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_copymsgbytes
func F_pq_copymsgbytes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_getmsgint64 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pq_getmsgint64
func F_pq_getmsgint64(m *base.Module, l0 int32) int64
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
//go:linkname F_bms_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_copy
func F_bms_copy(m *base.Module, l0 int32) int32
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
//go:linkname F_list_make1_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_make1_impl
func F_list_make1_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_make2_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_list_make2_impl
func F_list_make2_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_make3_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_make3_impl
func F_list_make3_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lappend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lappend
func F_lappend(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend_int github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_lappend_int
func F_lappend_int(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend_xid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lappend_xid
func F_lappend_xid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lcons github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lcons
func F_lcons(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lcons_int github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lcons_int
func F_lcons_int(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lcons_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lcons_oid
func F_lcons_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_list_copy
func F_list_copy(m *base.Module, l0 int32) int32
//go:linkname F_list_concat_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_concat_copy
func F_list_concat_copy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_member github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_member
func F_list_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_member_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_member_ptr
func F_list_member_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_nth_cell github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_delete_nth_cell
func F_list_delete_nth_cell(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_free
func F_list_free(m *base.Module, l0 int32)
//go:linkname F_list_delete_cell github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_delete_cell
func F_list_delete_cell(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_first github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_delete_first
func F_list_delete_first(m *base.Module, l0 int32) int32
//go:linkname F_list_append_unique_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_list_append_unique_oid
func F_list_append_unique_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_free_deep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_free_deep
func F_list_free_deep(m *base.Module, l0 int32)
//go:linkname F_list_copy_tail github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_copy_tail
func F_list_copy_tail(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeVar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeVar
func F_makeVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_makeVarFromTargetEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeVarFromTargetEntry
func F_makeVarFromTargetEntry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeTargetEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeTargetEntry
func F_makeTargetEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_flatCopyTargetEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_flatCopyTargetEntry
func F_flatCopyTargetEntry(m *base.Module, l0 int32) int32
//go:linkname F_makeConst github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeConst
func F_makeConst(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_makeNullConst github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeNullConst
func F_makeNullConst(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeBoolConst github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeBoolConst
func F_makeBoolConst(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeAlias github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makeAlias
func F_makeAlias(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeRangeVar github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeRangeVar
func F_makeRangeVar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeTypeNameFromNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeTypeNameFromNameList
func F_makeTypeNameFromNameList(m *base.Module, l0 int32) int32
//go:linkname F_make_andclause github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_andclause
func F_make_andclause(m *base.Module, l0 int32) int32
//go:linkname F_make_ands_explicit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_ands_explicit
func F_make_ands_explicit(m *base.Module, l0 int32) int32
//go:linkname F_make_ands_implicit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_ands_implicit
func F_make_ands_implicit(m *base.Module, l0 int32) int32
//go:linkname F_makeIndexInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeIndexInfo
func F_makeIndexInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32
//go:linkname F_exprType github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_exprType
func F_exprType(m *base.Module, l0 int32) int32
//go:linkname F_exprTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprTypmod
func F_exprTypmod(m *base.Module, l0 int32) int32
//go:linkname F_applyRelabelType github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_applyRelabelType
func F_applyRelabelType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_exprCollation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprCollation
func F_exprCollation(m *base.Module, l0 int32) int32
//go:linkname F_expression_returns_set github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expression_returns_set
func F_expression_returns_set(m *base.Module, l0 int32) int32
//go:linkname F_exprLocation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exprLocation
func F_exprLocation(m *base.Module, l0 int32) int32
//go:linkname F_fix_opfuncids github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fix_opfuncids
func F_fix_opfuncids(m *base.Module, l0 int32)
//go:linkname F_check_functions_in_node github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_functions_in_node
func F_check_functions_in_node(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_table_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_table_walker_impl
func F_range_table_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_expression_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_expression_tree_mutator_impl
func F_expression_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_query_or_expression_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_query_or_expression_tree_mutator_impl
func F_query_or_expression_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_raw_expression_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_raw_expression_tree_walker_impl
func F_raw_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_planstate_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_planstate_tree_walker_impl
func F_planstate_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_outNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_outNode
func F_outNode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_nodeToString github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_nodeToString
func F_nodeToString(m *base.Module, l0 int32) int32
//go:linkname F_makeParamList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeParamList
func F_makeParamList(m *base.Module, l0 int32) int32
//go:linkname F__jumbleNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__jumbleNode
func F__jumbleNode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AppendJumble32 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AppendJumble32
func F_AppendJumble32(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AppendJumble8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AppendJumble8
func F_AppendJumble8(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AppendJumble github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AppendJumble
func F_AppendJumble(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_stringToNode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_stringToNode
func F_stringToNode(m *base.Module, l0 int32) int32
//go:linkname F_tbm_free github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tbm_free
func F_tbm_free(m *base.Module, l0 int32)
//go:linkname F_tbm_add_tuples github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tbm_add_tuples
func F_tbm_add_tuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_tbm_mark_page_lossy github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tbm_mark_page_lossy
func F_tbm_mark_page_lossy(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tbm_end_iterate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tbm_end_iterate
func F_tbm_end_iterate(m *base.Module, l0 int32)
//go:linkname F_makeFloat github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeFloat
func F_makeFloat(m *base.Module, l0 int32) int32
//go:linkname F_makeString github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeString
func F_makeString(m *base.Module, l0 int32) int32
//go:linkname F_set_rel_consider_parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_rel_consider_parallel
func F_set_rel_consider_parallel(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_set_rel_size github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_rel_size
func F_set_rel_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_set_rel_pathlist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_rel_pathlist
func F_set_rel_pathlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_make_rel_from_joinlist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_rel_from_joinlist
func F_make_rel_from_joinlist(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clauselist_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clauselist_selectivity
func F_clauselist_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_clauselist_selectivity_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clauselist_selectivity_ext
func F_clauselist_selectivity_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64
//go:linkname F_clause_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_clause_selectivity
func F_clause_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_cost_qual_eval_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cost_qual_eval_walker
func F_cost_qual_eval_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cost_index github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_cost_index
func F_cost_index(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32)
//go:linkname F_cost_bitmap_heap_scan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cost_bitmap_heap_scan
func F_cost_bitmap_heap_scan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64)
//go:linkname F_cost_qual_eval_node github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cost_qual_eval_node
func F_cost_qual_eval_node(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_set_pathtarget_cost_width github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_pathtarget_cost_width
func F_set_pathtarget_cost_width(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_process_equivalence github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_process_equivalence
func F_process_equivalence(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_canonicalize_ec_expression github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_canonicalize_ec_expression
func F_canonicalize_ec_expression(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ec_add_clause_to_derives_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ec_add_clause_to_derives_hash
func F_ec_add_clause_to_derives_hash(m *base.Module, l0 int32, l1 int32)
//go:linkname F_find_computable_ec_member github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_find_computable_ec_member
func F_find_computable_ec_member(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_generate_join_implied_equalities github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_join_implied_equalities
func F_generate_join_implied_equalities(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_reconsider_outer_join_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_reconsider_outer_join_clause
func F_reconsider_outer_join_clause(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rebuild_eclass_attr_needed github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_rebuild_eclass_attr_needed
func F_rebuild_eclass_attr_needed(m *base.Module, l0 int32)
//go:linkname F_match_clause_to_index github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_match_clause_to_index
func F_match_clause_to_index(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_index_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_index_paths
func F_get_index_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_build_index_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_build_index_paths
func F_build_index_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_match_index_to_operand github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_match_index_to_operand
func F_match_index_to_operand(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_pathkey_from_sortinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_pathkey_from_sortinfo
func F_make_pathkey_from_sortinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_make_pathkeys_for_sortclauses_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_pathkeys_for_sortclauses_extended
func F_make_pathkeys_for_sortclauses_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_rel_is_distinct_for github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_rel_is_distinct_for
func F_rel_is_distinct_for(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_remove_rel_from_query github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_remove_rel_from_query
func F_remove_rel_from_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_remove_rel_from_restrictinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_remove_rel_from_restrictinfo
func F_remove_rel_from_restrictinfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_remove_rel_from_joinlist github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_remove_rel_from_joinlist
func F_remove_rel_from_joinlist(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_innerrel_is_unique_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_innerrel_is_unique_ext
func F_innerrel_is_unique_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_remove_self_joins_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_remove_self_joins_recurse
func F_remove_self_joins_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_replace_nestloop_params_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_replace_nestloop_params_mutator
func F_replace_nestloop_params_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fix_indexqual_operand github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fix_indexqual_operand
func F_fix_indexqual_operand(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_add_base_rels_to_query github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_base_rels_to_query
func F_add_base_rels_to_query(m *base.Module, l0 int32, l1 int32)
//go:linkname F_build_base_rel_tlists github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_build_base_rel_tlists
func F_build_base_rel_tlists(m *base.Module, l0 int32, l1 int32)
//go:linkname F_add_vars_to_targetlist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_add_vars_to_targetlist
func F_add_vars_to_targetlist(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_add_vars_to_attr_needed github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_vars_to_attr_needed
func F_add_vars_to_attr_needed(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_deconstruct_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_deconstruct_recurse
func F_deconstruct_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_distribute_quals_to_rels github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_distribute_quals_to_rels
func F_distribute_quals_to_rels(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32)
//go:linkname F_distribute_restrictinfo_to_rels github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_distribute_restrictinfo_to_rels
func F_distribute_restrictinfo_to_rels(m *base.Module, l0 int32, l1 int32)
//go:linkname F_build_implied_join_equality github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_implied_join_equality
func F_build_implied_join_equality(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_subquery_planner github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_subquery_planner
func F_subquery_planner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 int32) int32
//go:linkname F_preprocess_expression github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_preprocess_expression
func F_preprocess_expression(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_preprocess_groupclause github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_preprocess_groupclause
func F_preprocess_groupclause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plan_create_index_workers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plan_create_index_workers
func F_plan_create_index_workers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fix_expr_common github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fix_expr_common
func F_fix_expr_common(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fix_scan_expr_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fix_scan_expr_mutator
func F_fix_scan_expr_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fix_param_node github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fix_param_node
func F_fix_param_node(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_search_indexed_tlist_for_var github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_search_indexed_tlist_for_var
func F_search_indexed_tlist_for_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_search_indexed_tlist_for_phv github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_search_indexed_tlist_for_phv
func F_search_indexed_tlist_for_phv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_convert_testexpr_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_convert_testexpr_mutator
func F_convert_testexpr_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_convert_ANY_sublink_to_join github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_convert_ANY_sublink_to_join
func F_convert_ANY_sublink_to_join(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_convert_EXISTS_sublink_to_join github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_EXISTS_sublink_to_join
func F_convert_EXISTS_sublink_to_join(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SS_replace_correlation_vars github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SS_replace_correlation_vars
func F_SS_replace_correlation_vars(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SS_process_sublinks github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SS_process_sublinks
func F_SS_process_sublinks(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pull_up_sublinks_jointree_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pull_up_sublinks_jointree_recurse
func F_pull_up_sublinks_jointree_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_canonicalize_qual github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_canonicalize_qual
func F_canonicalize_qual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_build_setop_child_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_build_setop_child_paths
func F_build_setop_child_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_generate_append_tlist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_generate_append_tlist
func F_generate_append_tlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_generate_setop_tlist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generate_setop_tlist
func F_generate_setop_tlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_add_row_identity_var github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_row_identity_var
func F_add_row_identity_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_add_row_identity_columns github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_add_row_identity_columns
func F_add_row_identity_columns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_contain_agg_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_contain_agg_clause
func F_contain_agg_clause(m *base.Module, l0 int32) int32
//go:linkname F_contain_subplans github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_contain_subplans
func F_contain_subplans(m *base.Module, l0 int32) int32
//go:linkname F_contain_volatile_functions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_contain_volatile_functions
func F_contain_volatile_functions(m *base.Module, l0 int32) int32
//go:linkname F_is_parallel_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_is_parallel_safe
func F_is_parallel_safe(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_nonnullable_rels github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_nonnullable_rels
func F_find_nonnullable_rels(m *base.Module, l0 int32) int32
//go:linkname F_is_pseudo_constant_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_is_pseudo_constant_clause
func F_is_pseudo_constant_clause(m *base.Module, l0 int32) int32
//go:linkname F_eval_const_expressions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_eval_const_expressions
func F_eval_const_expressions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_eval_const_expressions_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_eval_const_expressions_mutator
func F_eval_const_expressions_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_evaluate_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_evaluate_expr
func F_evaluate_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_convert_saop_to_hashed_saop_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_convert_saop_to_hashed_saop_walker
func F_convert_saop_to_hashed_saop_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_SAOP_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_SAOP_expr
func F_make_SAOP_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_expand_inherited_rtentry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_expand_inherited_rtentry
func F_expand_inherited_rtentry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_remove_join_clause_from_rels github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_remove_join_clause_from_rels
func F_remove_join_clause_from_rels(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_extract_or_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_extract_or_clause
func F_extract_or_clause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_compare_path_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_compare_path_costs
func F_compare_path_costs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_set_cheapest github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_cheapest
func F_set_cheapest(m *base.Module, l0 int32)
//go:linkname F_add_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_path
func F_add_path(m *base.Module, l0 int32, l1 int32)
//go:linkname F_create_append_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_append_path
func F_create_append_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 float64) int32
//go:linkname F_create_merge_append_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_merge_append_path
func F_create_merge_append_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_group_result_path github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_group_result_path
func F_create_group_result_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_gather_merge_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_gather_merge_path
func F_create_gather_merge_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_create_gather_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_gather_path
func F_create_gather_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_create_projection_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_projection_path
func F_create_projection_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_apply_projection_to_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_apply_projection_to_path
func F_apply_projection_to_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_set_projection_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_set_projection_path
func F_create_set_projection_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_incremental_sort_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_incremental_sort_path
func F_create_incremental_sort_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) int32
//go:linkname F_create_sort_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_sort_path
func F_create_sort_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32
//go:linkname F_create_upper_unique_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_create_upper_unique_path
func F_create_upper_unique_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32
//go:linkname F_create_agg_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_agg_path
func F_create_agg_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 float64) int32
//go:linkname F_create_groupingsets_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_groupingsets_path
func F_create_groupingsets_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_create_setop_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_create_setop_path
func F_create_setop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 float64, l7 float64) int32
//go:linkname F_find_placeholder_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_placeholder_info
func F_find_placeholder_info(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_placeholders_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_placeholders_recurse
func F_find_placeholders_recurse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_rel_data_width github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rel_data_width
func F_get_rel_data_width(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_restriction_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_restriction_selectivity
func F_restriction_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_join_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_join_selectivity
func F_join_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64
//go:linkname F_function_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_function_selectivity
func F_function_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64
//go:linkname F_add_function_cost github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_function_cost
func F_add_function_cost(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_predicate_implied_by github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_predicate_implied_by
func F_predicate_implied_by(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_predicate_implied_by_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_predicate_implied_by_recurse
func F_predicate_implied_by_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_predicate_classify github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_predicate_classify
func F_predicate_classify(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clause_is_strict_for github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_clause_is_strict_for
func F_clause_is_strict_for(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_operator_same_subexprs_proof github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_operator_same_subexprs_proof
func F_operator_same_subexprs_proof(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lookup_proof_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lookup_proof_cache
func F_lookup_proof_cache(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_simple_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_simple_rel
func F_build_simple_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_base_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_base_rel
func F_find_base_rel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fetch_upper_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fetch_upper_rel
func F_fetch_upper_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_baserel_parampathinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_baserel_parampathinfo
func F_get_baserel_parampathinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_joinrel_parampathinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_joinrel_parampathinfo
func F_get_joinrel_parampathinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_make_restrictinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_restrictinfo
func F_make_restrictinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_join_clause_is_movable_to github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_join_clause_is_movable_to
func F_join_clause_is_movable_to(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tlist_same_datatypes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tlist_same_datatypes
func F_tlist_same_datatypes(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_sortgroupref_tle github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_sortgroupref_tle
func F_get_sortgroupref_tle(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_sortgrouplist_exprs github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_sortgrouplist_exprs
func F_get_sortgrouplist_exprs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_sortgroupref_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_sortgroupref_clause
func F_get_sortgroupref_clause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_pathtarget_from_tlist github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_pathtarget_from_tlist
func F_make_pathtarget_from_tlist(m *base.Module, l0 int32) int32
//go:linkname F_pull_varnos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pull_varnos
func F_pull_varnos(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pull_vars_of_level github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pull_vars_of_level
func F_pull_vars_of_level(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pull_var_clause_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pull_var_clause_walker
func F_pull_var_clause_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_flatten_join_alias_vars github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_flatten_join_alias_vars
func F_flatten_join_alias_vars(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_qual_from_partbound github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_qual_from_partbound
func F_get_qual_from_partbound(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_PGSemaphoreReset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PGSemaphoreReset
func F_PGSemaphoreReset(m *base.Module, l0 int32)
//go:linkname F_RegisterDynamicBackgroundWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RegisterDynamicBackgroundWorker
func F_RegisterDynamicBackgroundWorker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AbsorbSyncRequests github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AbsorbSyncRequests
func F_AbsorbSyncRequests(m *base.Module)
//go:linkname F_CheckArchiveTimeout github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckArchiveTimeout
func F_CheckArchiveTimeout(m *base.Module)
//go:linkname F_RequestCheckpoint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RequestCheckpoint
func F_RequestCheckpoint(m *base.Module, l0 int32)
//go:linkname F_ProcessStartupProcInterrupts github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessStartupProcInterrupts
func F_ProcessStartupProcInterrupts(m *base.Module)
//go:linkname F_disable_startup_progress_timeout github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_disable_startup_progress_timeout
func F_disable_startup_progress_timeout(m *base.Module)
//go:linkname F_begin_startup_progress_phase github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_begin_startup_progress_phase
func F_begin_startup_progress_phase(m *base.Module)
//go:linkname F_ProcessWalSummarizerInterrupts github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcessWalSummarizerInterrupts
func F_ProcessWalSummarizerInterrupts(m *base.Module)
//go:linkname F_GetLatestLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetLatestLSN
func F_GetLatestLSN(m *base.Module, l0 int32) int64
//go:linkname F_WaitForWalSummarization github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WaitForWalSummarization
func F_WaitForWalSummarization(m *base.Module, l0 int64)
//go:linkname F_pg_set_regex_collation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_set_regex_collation
func F_pg_set_regex_collation(m *base.Module, l0 int32)
//go:linkname F_newnfa github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_newnfa
func F_newnfa(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_okcolors github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_okcolors
func F_okcolors(m *base.Module, l0 int32, l1 int32)
//go:linkname F_parse github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_parse
func F_parse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_optimize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_optimize
func F_optimize(m *base.Module, l0 int32) int32
//go:linkname F_makesearch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makesearch
func F_makesearch(m *base.Module, l0 int32, l1 int32)
//go:linkname F_compact github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_compact
func F_compact(m *base.Module, l0 int32, l1 int32)
//go:linkname F_freesubre github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_freesubre
func F_freesubre(m *base.Module, l0 int32, l1 int32)
//go:linkname F_newstate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_newstate
func F_newstate(m *base.Module, l0 int32) int32
//go:linkname F_rainbow github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_rainbow
func F_rainbow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_createarc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_createarc
func F_createarc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_dupnfa github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_dupnfa
func F_dupnfa(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_moveins github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_moveins
func F_moveins(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_moveouts github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_moveouts
func F_moveouts(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_cparc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cparc
func F_cparc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pg_reg_getcolor github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_reg_getcolor
func F_pg_reg_getcolor(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_newarc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_newarc
func F_newarc(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_wordchrs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_wordchrs
func F_wordchrs(m *base.Module, l0 int32)
//go:linkname F_nonword github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_nonword
func F_nonword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_cloneouts github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cloneouts
func F_cloneouts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_subcoloronechr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_subcoloronechr
func F_subcoloronechr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_allcases github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_allcases
func F_allcases(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_subcolorcvec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_subcolorcvec
func F_subcolorcvec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_bracket github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bracket
func F_bracket(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_cclasscvec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cclasscvec
func F_cclasscvec(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_charclasscomplement github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_charclasscomplement
func F_charclasscomplement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_deltraverse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_deltraverse
func F_deltraverse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_repeat_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_repeat_1
func F_repeat_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_removetraverse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_removetraverse
func F_removetraverse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_cleartraverse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cleartraverse
func F_cleartraverse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lexescape github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lexescape
func F_lexescape(m *base.Module, l0 int32) int32
//go:linkname F_pg_regerror github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_regerror
func F_pg_regerror(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pa_find_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pa_find_worker
func F_pa_find_worker(m *base.Module, l0 int32) int32
//go:linkname F_pa_send_data github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pa_send_data
func F_pa_send_data(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pa_switch_to_partial_serialize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pa_switch_to_partial_serialize
func F_pa_switch_to_partial_serialize(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetTupleTransactionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetTupleTransactionInfo
func F_GetTupleTransactionInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ReportApplyConflict github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReportApplyConflict
func F_ReportApplyConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_LogicalDecodingProcessRecord github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LogicalDecodingProcessRecord
func F_LogicalDecodingProcessRecord(m *base.Module, l0 int32, l1 int32)
//go:linkname F_logicalrep_worker_launch github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_logicalrep_worker_launch
func F_logicalrep_worker_launch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_logicalrep_worker_stop github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_logicalrep_worker_stop
func F_logicalrep_worker_stop(m *base.Module, l0 int32, l1 int32)
//go:linkname F_logicalrep_launcher_attach_dshmem github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_logicalrep_launcher_attach_dshmem
func F_logicalrep_launcher_attach_dshmem(m *base.Module)
//go:linkname F_CheckLogicalDecodingRequirements github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckLogicalDecodingRequirements
func F_CheckLogicalDecodingRequirements(m *base.Module)
//go:linkname F_CreateInitDecodingContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CreateInitDecodingContext
func F_CreateInitDecodingContext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_FreeDecodingContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FreeDecodingContext
func F_FreeDecodingContext(m *base.Module, l0 int32)
//go:linkname F_filter_by_origin_cb_wrapper github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_filter_by_origin_cb_wrapper
func F_filter_by_origin_cb_wrapper(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LogicalSlotAdvanceAndCheckSnapState github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalSlotAdvanceAndCheckSnapState
func F_LogicalSlotAdvanceAndCheckSnapState(m *base.Module, l0 int64, l1 int32) int64
//go:linkname F_pg_logical_slot_get_changes_guts github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_logical_slot_get_changes_guts
func F_pg_logical_slot_get_changes_guts(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_replorigin_create github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replorigin_create
func F_replorigin_create(m *base.Module, l0 int32) int32
//go:linkname F_replorigin_drop_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_replorigin_drop_by_name
func F_replorigin_drop_by_name(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_replorigin_advance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_replorigin_advance
func F_replorigin_advance(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32)
//go:linkname F_logicalrep_get_attrs_str github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_logicalrep_get_attrs_str
func F_logicalrep_get_attrs_str(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_logicalrep_rel_mark_updatable github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_logicalrep_rel_mark_updatable
func F_logicalrep_rel_mark_updatable(m *base.Module, l0 int32)
//go:linkname F_FindLogicalRepLocalIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FindLogicalRepLocalIndex
func F_FindLogicalRepLocalIndex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetRelationIdentityOrPK github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetRelationIdentityOrPK
func F_GetRelationIdentityOrPK(m *base.Module, l0 int32) int32
//go:linkname F_ReorderBufferCleanupSerializedTXNs github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReorderBufferCleanupSerializedTXNs
func F_ReorderBufferCleanupSerializedTXNs(m *base.Module, l0 int32)
//go:linkname F_ReorderBufferAllocChange github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferAllocChange
func F_ReorderBufferAllocChange(m *base.Module, l0 int32) int32
//go:linkname F_ReorderBufferFreeChange github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReorderBufferFreeChange
func F_ReorderBufferFreeChange(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReorderBufferQueueChange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReorderBufferQueueChange
func F_ReorderBufferQueueChange(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32)
//go:linkname F_ReorderBufferTXNByXid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReorderBufferTXNByXid
func F_ReorderBufferTXNByXid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_ReorderBufferQueueMessage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReorderBufferQueueMessage
func F_ReorderBufferQueueMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_ReorderBufferTransferSnapToParent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReorderBufferTransferSnapToParent
func F_ReorderBufferTransferSnapToParent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferProcessTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReorderBufferProcessTXN
func F_ReorderBufferProcessTXN(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32)
//go:linkname F_ReorderBufferRestoreCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReorderBufferRestoreCleanup
func F_ReorderBufferRestoreCleanup(m *base.Module, l0 int32)
//go:linkname F_ReorderBufferSetBaseSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferSetBaseSnapshot
func F_ReorderBufferSetBaseSnapshot(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_ReorderBufferXidSetCatalogChanges github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferXidSetCatalogChanges
func F_ReorderBufferXidSetCatalogChanges(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_ReorderBufferCopySnap github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReorderBufferCopySnap
func F_ReorderBufferCopySnap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_check_and_set_sync_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_check_and_set_sync_info
func F_check_and_set_sync_info(m *base.Module, l0 int32)
//go:linkname F_validate_remote_info github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_validate_remote_info
func F_validate_remote_info(m *base.Module, l0 int32)
//go:linkname F_synchronize_slots github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_synchronize_slots
func F_synchronize_slots(m *base.Module, l0 int32) int32
//go:linkname F_SnapBuildSnapDecRefcount github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SnapBuildSnapDecRefcount
func F_SnapBuildSnapDecRefcount(m *base.Module, l0 int32)
//go:linkname F_process_syncing_tables github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_process_syncing_tables
func F_process_syncing_tables(m *base.Module, l0 int64)
//go:linkname F_ReplicationSlotNameForTablesync github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotNameForTablesync
func F_ReplicationSlotNameForTablesync(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_UpdateTwoPhaseState github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UpdateTwoPhaseState
func F_UpdateTwoPhaseState(m *base.Module, l0 int32)
//go:linkname F_ReplicationOriginNameForLogicalRep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReplicationOriginNameForLogicalRep
func F_ReplicationOriginNameForLogicalRep(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_maybe_reread_subscription github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_maybe_reread_subscription
func F_maybe_reread_subscription(m *base.Module)
//go:linkname F_apply_dispatch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_apply_dispatch
func F_apply_dispatch(m *base.Module, l0 int32)
//go:linkname F_TargetPrivilegesCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TargetPrivilegesCheck
func F_TargetPrivilegesCheck(m *base.Module, l0 int32, l1 int64)
//go:linkname F_clear_subscription_skip_lsn github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clear_subscription_skip_lsn
func F_clear_subscription_skip_lsn(m *base.Module, l0 int64)
//go:linkname F_send_feedback github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_send_feedback
func F_send_feedback(m *base.Module, l0 int64, l1 int32, l2 int32)
//go:linkname F_DisableSubscriptionAndExit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DisableSubscriptionAndExit
func F_DisableSubscriptionAndExit(m *base.Module)
//go:linkname F_yy_fatal_error_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_yy_fatal_error_3
func F_yy_fatal_error_3(m *base.Module, l0 int32)
//go:linkname F_replication_scanner_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_replication_scanner_finish
func F_replication_scanner_finish(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotValidateName github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReplicationSlotValidateName
func F_ReplicationSlotValidateName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReplicationSlotValidateNameInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplicationSlotValidateNameInternal
func F_ReplicationSlotValidateNameInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ReplicationSlotCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReplicationSlotCreate
func F_ReplicationSlotCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_SearchNamedReplicationSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SearchNamedReplicationSlot
func F_SearchNamedReplicationSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReplicationSlotDropPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReplicationSlotDropPtr
func F_ReplicationSlotDropPtr(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotsComputeRequiredXmin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotsComputeRequiredXmin
func F_ReplicationSlotsComputeRequiredXmin(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotsComputeRequiredLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReplicationSlotsComputeRequiredLSN
func F_ReplicationSlotsComputeRequiredLSN(m *base.Module)
//go:linkname F_ReplicationSlotMarkDirty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReplicationSlotMarkDirty
func F_ReplicationSlotMarkDirty(m *base.Module)
//go:linkname F_ReplicationSlotSave github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotSave
func F_ReplicationSlotSave(m *base.Module)
//go:linkname F_ReplicationSlotPersist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReplicationSlotPersist
func F_ReplicationSlotPersist(m *base.Module)
//go:linkname F_ReplicationSlotsComputeLogicalRestartLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotsComputeLogicalRestartLSN
func F_ReplicationSlotsComputeLogicalRestartLSN(m *base.Module) int64
//go:linkname F_CheckSlotRequirements github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckSlotRequirements
func F_CheckSlotRequirements(m *base.Module)
//go:linkname F_CheckSlotPermissions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckSlotPermissions
func F_CheckSlotPermissions(m *base.Module)
//go:linkname F_ReportSlotInvalidation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReportSlotInvalidation
func F_ReportSlotInvalidation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64, l6 int32, l7 int32)
//go:linkname F_SyncRepUpdateSyncStandbysDefined github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SyncRepUpdateSyncStandbysDefined
func F_SyncRepUpdateSyncStandbysDefined(m *base.Module)
//go:linkname F_create_syncrep_config github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_syncrep_config
func F_create_syncrep_config(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_syncrep_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_syncrep_yyensure_buffer_stack
func F_syncrep_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_syncrep_yy_create_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_syncrep_yy_create_buffer
func F_syncrep_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_WalRcvForceReply github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WalRcvForceReply
func F_WalRcvForceReply(m *base.Module)
//go:linkname F_WalRcvStreaming github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WalRcvStreaming
func F_WalRcvStreaming(m *base.Module) int32
//go:linkname F_GetWalRcvFlushRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetWalRcvFlushRecPtr
func F_GetWalRcvFlushRecPtr(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_GetReplicationTransferLatency github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetReplicationTransferLatency
func F_GetReplicationTransferLatency(m *base.Module) int32
//go:linkname F_WalSndWakeup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WalSndWakeup
func F_WalSndWakeup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_checkRuleResultList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_checkRuleResultList
func F_checkRuleResultList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_setRuleCheckAsUser github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_setRuleCheckAsUser
func F_setRuleCheckAsUser(m *base.Module, l0 int32, l1 int32)
//go:linkname F_EnableDisableRule github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_EnableDisableRule
func F_EnableDisableRule(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AcquireRewriteLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AcquireRewriteLocks
func F_AcquireRewriteLocks(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_build_column_default github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_column_default
func F_build_column_default(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_view_query github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_view_query
func F_get_view_query(m *base.Module, l0 int32) int32
//go:linkname F_expand_generated_columns_in_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_expand_generated_columns_in_expr
func F_expand_generated_columns_in_expr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_contain_windowfuncs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_contain_windowfuncs
func F_contain_windowfuncs(m *base.Module, l0 int32) int32
//go:linkname F_OffsetVarNodes_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_OffsetVarNodes_walker
func F_OffsetVarNodes_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ChangeVarNodes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ChangeVarNodes
func F_ChangeVarNodes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_IncrementVarSublevelsUp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_IncrementVarSublevelsUp
func F_IncrementVarSublevelsUp(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_getInsertSelectQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getInsertSelectQuery
func F_getInsertSelectQuery(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_add_nulling_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_nulling_relids
func F_add_nulling_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_remove_nulling_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_remove_nulling_relids
func F_remove_nulling_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_statext_clauselist_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_statext_clauselist_selectivity
func F_statext_clauselist_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64
//go:linkname F_statext_mcv_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_statext_mcv_deserialize
func F_statext_mcv_deserialize(m *base.Module, l0 int32) int32
//go:linkname F_statext_ndistinct_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_statext_ndistinct_deserialize
func F_statext_ndistinct_deserialize(m *base.Module, l0 int32) int32
//go:linkname F_pgaio_io_reclaim github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgaio_io_reclaim
func F_pgaio_io_reclaim(m *base.Module, l0 int32)
//go:linkname F_pgaio_submit_staged github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgaio_submit_staged
func F_pgaio_submit_staged(m *base.Module)
//go:linkname F_pgaio_io_update_state github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgaio_io_update_state
func F_pgaio_io_update_state(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgaio_io_process_completion github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgaio_io_process_completion
func F_pgaio_io_process_completion(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgaio_wref_wait github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgaio_wref_wait
func F_pgaio_wref_wait(m *base.Module, l0 int32)
//go:linkname F_pgaio_io_perform_synchronously github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgaio_io_perform_synchronously
func F_pgaio_io_perform_synchronously(m *base.Module, l0 int32)
//go:linkname F_read_stream_begin_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_read_stream_begin_relation
func F_read_stream_begin_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_read_stream_next_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_read_stream_next_buffer
func F_read_stream_next_buffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_read_stream_reset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_read_stream_reset
func F_read_stream_reset(m *base.Module, l0 int32)
//go:linkname F_read_stream_end github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_stream_end
func F_read_stream_end(m *base.Module, l0 int32)
//go:linkname F_BufTableHashCode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BufTableHashCode
func F_BufTableHashCode(m *base.Module, l0 int32) int32
//go:linkname F_BufTableInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufTableInsert
func F_BufTableInsert(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_UnpinBufferNoOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnpinBufferNoOwner
func F_UnpinBufferNoOwner(m *base.Module, l0 int32)
//go:linkname F_PrefetchSharedBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PrefetchSharedBuffer
func F_PrefetchSharedBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ReservePrivateRefCountEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReservePrivateRefCountEntry
func F_ReservePrivateRefCountEntry(m *base.Module)
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
//go:linkname F_StartBufferIO github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_StartBufferIO
func F_StartBufferIO(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReleaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReleaseBuffer
func F_ReleaseBuffer(m *base.Module, l0 int32)
//go:linkname F_MarkBufferDirty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MarkBufferDirty
func F_MarkBufferDirty(m *base.Module, l0 int32)
//go:linkname F_ReleaseAndReadBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReleaseAndReadBuffer
func F_ReleaseAndReadBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SyncOneBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SyncOneBuffer
func F_SyncOneBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IssuePendingWritebacks github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_IssuePendingWritebacks
func F_IssuePendingWritebacks(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sort_checkpoint_bufferids github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sort_checkpoint_bufferids
func F_sort_checkpoint_bufferids(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BufferGetBlockNumber github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BufferGetBlockNumber
func F_BufferGetBlockNumber(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetNumberOfBlocksInFork github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetNumberOfBlocksInFork
func F_RelationGetNumberOfBlocksInFork(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BufferGetLSNAtomic github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufferGetLSNAtomic
func F_BufferGetLSNAtomic(m *base.Module, l0 int32) int64
//go:linkname F_UnlockReleaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UnlockReleaseBuffer
func F_UnlockReleaseBuffer(m *base.Module, l0 int32)
//go:linkname F_IncrBufferRefCount github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_IncrBufferRefCount
func F_IncrBufferRefCount(m *base.Module, l0 int32)
//go:linkname F_MarkBufferDirtyHint github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MarkBufferDirtyHint
func F_MarkBufferDirtyHint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ConditionalLockBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConditionalLockBuffer
func F_ConditionalLockBuffer(m *base.Module, l0 int32) int32
//go:linkname F_LockBufferForCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockBufferForCleanup
func F_LockBufferForCleanup(m *base.Module, l0 int32)
//go:linkname F_ConditionalLockBufferForCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ConditionalLockBufferForCleanup
func F_ConditionalLockBufferForCleanup(m *base.Module, l0 int32) int32
//go:linkname F_GetAccessStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetAccessStrategy
func F_GetAccessStrategy(m *base.Module, l0 int32) int32
//go:linkname F_IOContextForStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IOContextForStrategy
func F_IOContextForStrategy(m *base.Module, l0 int32) int32
//go:linkname F_InitLocalBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InitLocalBuffers
func F_InitLocalBuffers(m *base.Module)
//go:linkname F_GetLocalVictimBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetLocalVictimBuffer
func F_GetLocalVictimBuffer(m *base.Module) int32
//go:linkname F_BufFileCreateTemp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BufFileCreateTemp
func F_BufFileCreateTemp(m *base.Module, l0 int32) int32
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
//go:linkname F_BufFileSize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BufFileSize
func F_BufFileSize(m *base.Module, l0 int32) int64
//go:linkname F_copy_file github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_copy_file
func F_copy_file(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_file_exists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_file_exists
func F_pg_file_exists(m *base.Module, l0 int32) int32
//go:linkname F_fsync_fname github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fsync_fname
func F_fsync_fname(m *base.Module, l0 int32, l1 int32)
//go:linkname F_OpenTransientFilePerm github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_OpenTransientFilePerm
func F_OpenTransientFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CloseTransientFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CloseTransientFile
func F_CloseTransientFile(m *base.Module, l0 int32) int32
//go:linkname F_durable_rename github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_durable_rename
func F_durable_rename(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_OpenTransientFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_OpenTransientFile
func F_OpenTransientFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_durable_unlink github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_durable_unlink
func F_durable_unlink(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BasicOpenFile github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BasicOpenFile
func F_BasicOpenFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BasicOpenFilePerm github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BasicOpenFilePerm
func F_BasicOpenFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReserveExternalFD github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReserveExternalFD
func F_ReserveExternalFD(m *base.Module)
//go:linkname F_PathNameOpenFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PathNameOpenFile
func F_PathNameOpenFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_walkdir github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_walkdir
func F_walkdir(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_AllocateDir github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AllocateDir
func F_AllocateDir(m *base.Module, l0 int32) int32
//go:linkname F_ReadDirExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReadDirExtended
func F_ReadDirExtended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FreeDir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeDir
func F_FreeDir(m *base.Module, l0 int32)
//go:linkname F_OpenTemporaryFileInTablespace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_OpenTemporaryFileInTablespace
func F_OpenTemporaryFileInTablespace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FileAccess github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileAccess
func F_FileAccess(m *base.Module, l0 int32) int32
//go:linkname F_FileReadV github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileReadV
func F_FileReadV(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_FileWriteV github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FileWriteV
func F_FileWriteV(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_AllocateFile github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AllocateFile
func F_AllocateFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FreeFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FreeFile
func F_FreeFile(m *base.Module, l0 int32) int32
//go:linkname F_ReadDir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadDir
func F_ReadDir(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ClosePipeStream github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ClosePipeStream
func F_ClosePipeStream(m *base.Module, l0 int32) int32
//go:linkname F_AtEOSubXact_Files github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AtEOSubXact_Files
func F_AtEOSubXact_Files(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_FileSetCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FileSetCreate
func F_FileSetCreate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FileSetDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FileSetDelete
func F_FileSetDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FileSetDeleteAll github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileSetDeleteAll
func F_FileSetDeleteAll(m *base.Module, l0 int32)
//go:linkname F_ResetUnloggedRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ResetUnloggedRelations
func F_ResetUnloggedRelations(m *base.Module, l0 int32)
//go:linkname F_GetPageWithFreeSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetPageWithFreeSpace
func F_GetPageWithFreeSpace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fsm_readbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fsm_readbuf
func F_fsm_readbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RecordAndGetPageWithFreeSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RecordAndGetPageWithFreeSpace
func F_RecordAndGetPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_RecordPageWithFreeSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RecordPageWithFreeSpace
func F_RecordPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogRecordPageWithFreeSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XLogRecordPageWithFreeSpace
func F_XLogRecordPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_FreeSpaceMapVacuum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeSpaceMapVacuum
func F_FreeSpaceMapVacuum(m *base.Module, l0 int32)
//go:linkname F_FreeSpaceMapVacuumRange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreeSpaceMapVacuumRange
func F_FreeSpaceMapVacuumRange(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_fsm_search_avail github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fsm_search_avail
func F_fsm_search_avail(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GetFreeIndexPage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetFreeIndexPage
func F_GetFreeIndexPage(m *base.Module, l0 int32) int32
//go:linkname F_RecordFreeIndexPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RecordFreeIndexPage
func F_RecordFreeIndexPage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BarrierArriveAndWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BarrierArriveAndWait
func F_BarrierArriveAndWait(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsm_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsm_detach
func F_dsm_detach(m *base.Module, l0 int32)
//go:linkname F_dsm_pin_mapping github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dsm_pin_mapping
func F_dsm_pin_mapping(m *base.Module, l0 int32)
//go:linkname F_dsm_pin_segment github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsm_pin_segment
func F_dsm_pin_segment(m *base.Module, l0 int32)
//go:linkname F_dsm_unpin_segment github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dsm_unpin_segment
func F_dsm_unpin_segment(m *base.Module, l0 int32)
//go:linkname F_on_dsm_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_on_dsm_detach
func F_on_dsm_detach(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dsm_impl_op github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dsm_impl_op
func F_dsm_impl_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_GetNamedDSMSegment github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetNamedDSMSegment
func F_GetNamedDSMSegment(m *base.Module, l0 int32) int32
//go:linkname F_proc_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_proc_exit
func F_proc_exit(m *base.Module, l0 int32)
//go:linkname F_proc_exit_prepare github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_proc_exit_prepare
func F_proc_exit_prepare(m *base.Module, l0 int32)
//go:linkname F_on_proc_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_on_proc_exit
func F_on_proc_exit(m *base.Module, l0 int32)
//go:linkname F_before_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_before_shmem_exit
func F_before_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_cancel_before_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cancel_before_shmem_exit
func F_cancel_before_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_check_on_shmem_exit_lists_are_empty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_on_shmem_exit_lists_are_empty
func F_check_on_shmem_exit_lists_are_empty(m *base.Module)
//go:linkname F_CreateOrAttachShmemStructs github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateOrAttachShmemStructs
func F_CreateOrAttachShmemStructs(m *base.Module)
//go:linkname F_OwnLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_OwnLatch
func F_OwnLatch(m *base.Module, l0 int32)
//go:linkname F_WaitLatchOrSocket github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WaitLatchOrSocket
func F_WaitLatchOrSocket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SetLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_SetLatch
func F_SetLatch(m *base.Module, l0 int32)
//go:linkname F_RegisterPostmasterChildActive github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RegisterPostmasterChildActive
func F_RegisterPostmasterChildActive(m *base.Module)
//go:linkname F_ProcArrayAdd github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ProcArrayAdd
func F_ProcArrayAdd(m *base.Module, l0 int32)
//go:linkname F_ProcArrayApplyRecoveryInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcArrayApplyRecoveryInfo
func F_ProcArrayApplyRecoveryInfo(m *base.Module, l0 int32)
//go:linkname F_RecordKnownAssignedTransactionIds github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RecordKnownAssignedTransactionIds
func F_RecordKnownAssignedTransactionIds(m *base.Module, l0 int32)
//go:linkname F_KnownAssignedXidsRemoveTree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_KnownAssignedXidsRemoveTree
func F_KnownAssignedXidsRemoveTree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_TransactionIdIsInProgress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TransactionIdIsInProgress
func F_TransactionIdIsInProgress(m *base.Module, l0 int32) int32
//go:linkname F_GetOldestNonRemovableTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetOldestNonRemovableTransactionId
func F_GetOldestNonRemovableTransactionId(m *base.Module, l0 int32) int32
//go:linkname F_ComputeXidHorizons github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ComputeXidHorizons
func F_ComputeXidHorizons(m *base.Module, l0 int32)
//go:linkname F_GlobalVisHorizonKindForRel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GlobalVisHorizonKindForRel
func F_GlobalVisHorizonKindForRel(m *base.Module, l0 int32) int32
//go:linkname F_GetSnapshotData github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetSnapshotData
func F_GetSnapshotData(m *base.Module, l0 int32) int32
//go:linkname F_BackendPidGetProc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BackendPidGetProc
func F_BackendPidGetProc(m *base.Module, l0 int32) int32
//go:linkname F_GetConflictingVirtualXIDs github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetConflictingVirtualXIDs
func F_GetConflictingVirtualXIDs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GlobalVisTestIsRemovableXid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GlobalVisTestIsRemovableXid
func F_GlobalVisTestIsRemovableXid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GlobalVisCheckRemovableFullXid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GlobalVisCheckRemovableFullXid
func F_GlobalVisCheckRemovableFullXid(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_ExpireTreeKnownAssignedTransactionIds github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExpireTreeKnownAssignedTransactionIds
func F_ExpireTreeKnownAssignedTransactionIds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_KnownAssignedTransactionIdsIdleMaintenance github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_KnownAssignedTransactionIdsIdleMaintenance
func F_KnownAssignedTransactionIdsIdleMaintenance(m *base.Module)
//go:linkname F_ProcSignalInit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcSignalInit
func F_ProcSignalInit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SendProcSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SendProcSignal
func F_SendProcSignal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EmitProcSignalBarrier github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EmitProcSignalBarrier
func F_EmitProcSignalBarrier(m *base.Module) int64
//go:linkname F_WaitForProcSignalBarrier github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WaitForProcSignalBarrier
func F_WaitForProcSignalBarrier(m *base.Module, l0 int64)
//go:linkname F_ProcessProcSignalBarrier github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcessProcSignalBarrier
func F_ProcessProcSignalBarrier(m *base.Module)
//go:linkname F_shm_mq_set_receiver github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shm_mq_set_receiver
func F_shm_mq_set_receiver(m *base.Module, l0 int32, l1 int32)
//go:linkname F_shm_mq_sendv github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shm_mq_sendv
func F_shm_mq_sendv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_shm_mq_receive github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_shm_mq_receive
func F_shm_mq_receive(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_shm_toc_allocate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_shm_toc_allocate
func F_shm_toc_allocate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_shm_toc_lookup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_shm_toc_lookup
func F_shm_toc_lookup(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_shm_toc_estimate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_shm_toc_estimate
func F_shm_toc_estimate(m *base.Module, l0 int32) int32
//go:linkname F_ShmemInitStruct github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ShmemInitStruct
func F_ShmemInitStruct(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mul_size github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_mul_size
func F_mul_size(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReceiveSharedInvalidMessages github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReceiveSharedInvalidMessages
func F_ReceiveSharedInvalidMessages(m *base.Module)
//go:linkname F_ShutdownRecoveryTransactionEnvironment github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ShutdownRecoveryTransactionEnvironment
func F_ShutdownRecoveryTransactionEnvironment(m *base.Module)
//go:linkname F_StandbyReleaseAllLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_StandbyReleaseAllLocks
func F_StandbyReleaseAllLocks(m *base.Module)
//go:linkname F_ResolveRecoveryConflictWithVirtualXIDs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResolveRecoveryConflictWithVirtualXIDs
func F_ResolveRecoveryConflictWithVirtualXIDs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ResolveRecoveryConflictWithSnapshotFullXid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ResolveRecoveryConflictWithSnapshotFullXid
func F_ResolveRecoveryConflictWithSnapshotFullXid(m *base.Module, l0 int64, l1 int32, l2 int32)
//go:linkname F_LogStandbySnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogStandbySnapshot
func F_LogStandbySnapshot(m *base.Module) int64
//go:linkname F_CreateWaitEventSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateWaitEventSet
func F_CreateWaitEventSet(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AddWaitEventToSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AddWaitEventToSet
func F_AddWaitEventToSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ModifyWaitEvent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ModifyWaitEvent
func F_ModifyWaitEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_WaitEventSetWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_WaitEventSetWait
func F_WaitEventSetWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_inv_read github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inv_read
func F_inv_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ConditionVariableInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ConditionVariableInit
func F_ConditionVariableInit(m *base.Module, l0 int32)
//go:linkname F_ConditionVariablePrepareToSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ConditionVariablePrepareToSleep
func F_ConditionVariablePrepareToSleep(m *base.Module, l0 int32)
//go:linkname F_ConditionVariableCancelSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ConditionVariableCancelSleep
func F_ConditionVariableCancelSleep(m *base.Module)
//go:linkname F_ConditionVariableSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ConditionVariableSleep
func F_ConditionVariableSleep(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ConditionVariableSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConditionVariableSignal
func F_ConditionVariableSignal(m *base.Module, l0 int32)
//go:linkname F_ConditionVariableBroadcast github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ConditionVariableBroadcast
func F_ConditionVariableBroadcast(m *base.Module, l0 int32)
//go:linkname F_LockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockRelationOid
func F_LockRelationOid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ConditionalLockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConditionalLockRelationOid
func F_ConditionalLockRelationOid(m *base.Module, l0 int32, l1 int32) int32
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
//go:linkname F_LockPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockPage
func F_LockPage(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_UnlockPage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UnlockPage
func F_UnlockPage(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_UnlockTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockTuple
func F_UnlockTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XactLockTableWait github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_XactLockTableWait
func F_XactLockTableWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_WaitForLockersMultiple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_WaitForLockersMultiple
func F_WaitForLockersMultiple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockDatabaseObject github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LockDatabaseObject
func F_LockDatabaseObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockSharedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockSharedObject
func F_LockSharedObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockAcquireExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockAcquireExtended
func F_LockAcquireExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_LockRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LockRelease
func F_LockRelease(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LockReleaseAll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LockReleaseAll
func F_LockReleaseAll(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LockRefindAndRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockRefindAndRelease
func F_LockRefindAndRelease(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ReleaseLockIfHeld github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReleaseLockIfHeld
func F_ReleaseLockIfHeld(m *base.Module, l0 int32, l1 int32)
//go:linkname F_VirtualXactLockTableInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_VirtualXactLockTableInsert
func F_VirtualXactLockTableInsert(m *base.Module, l0 int32)
//go:linkname F_LWLockRegisterTranche github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LWLockRegisterTranche
func F_LWLockRegisterTranche(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LWLockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockAcquire
func F_LWLockAcquire(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LWLockRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LWLockRelease
func F_LWLockRelease(m *base.Module, l0 int32)
//go:linkname F_LWLockReleaseAll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockReleaseAll
func F_LWLockReleaseAll(m *base.Module)
//go:linkname F_ReleasePredicateLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ReleasePredicateLocks
func F_ReleasePredicateLocks(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PredicateLockRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PredicateLockRelation
func F_PredicateLockRelation(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PredicateLockPage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PredicateLockPage
func F_PredicateLockPage(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_TransferPredicateLocksToHeapRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransferPredicateLocksToHeapRelation
func F_TransferPredicateLocksToHeapRelation(m *base.Module, l0 int32)
//go:linkname F_FlagRWConflict github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FlagRWConflict
func F_FlagRWConflict(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CheckForSerializableConflictIn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckForSerializableConflictIn
func F_CheckForSerializableConflictIn(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockErrorCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockErrorCleanup
func F_LockErrorCleanup(m *base.Module)
//go:linkname F_ProcSendSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcSendSignal
func F_ProcSendSignal(m *base.Module, l0 int32)
//go:linkname F_s_lock_stuck github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_s_lock_stuck
func F_s_lock_stuck(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_perform_spin_delay github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_perform_spin_delay
func F_perform_spin_delay(m *base.Module, l0 int32)
//go:linkname F_PageInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PageInit
func F_PageInit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PageAddItemExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PageAddItemExtended
func F_PageAddItemExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_PageGetTempPageCopySpecial github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageGetTempPageCopySpecial
func F_PageGetTempPageCopySpecial(m *base.Module, l0 int32) int32
//go:linkname F_PageRestoreTempPage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PageRestoreTempPage
func F_PageRestoreTempPage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PageIndexTupleDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageIndexTupleDelete
func F_PageIndexTupleDelete(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PageIndexMultiDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PageIndexMultiDelete
func F_PageIndexMultiDelete(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PageIndexTupleOverwrite github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageIndexTupleOverwrite
func F_PageIndexTupleOverwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_checksum_page github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_checksum_page
func F_pg_checksum_page(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgr_bulk_start_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgr_bulk_start_rel
func F_smgr_bulk_start_rel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgr_bulk_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_smgr_bulk_finish
func F_smgr_bulk_finish(m *base.Module, l0 int32)
//go:linkname F_smgr_bulk_write github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_smgr_bulk_write
func F_smgr_bulk_write(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_smgr_bulk_get_buf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgr_bulk_get_buf
func F_smgr_bulk_get_buf(m *base.Module, l0 int32) int32
//go:linkname F__mdfd_getseg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F__mdfd_getseg
func F__mdfd_getseg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_mdnblocks github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mdnblocks
func F_mdnblocks(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DropRelationFiles github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DropRelationFiles
func F_DropRelationFiles(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_smgropen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgropen
func F_smgropen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgrclose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrclose
func F_smgrclose(m *base.Module, l0 int32)
//go:linkname F_smgrcreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgrcreate
func F_smgrcreate(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_smgrextend github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrextend
func F_smgrextend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_smgrzeroextend github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgrzeroextend
func F_smgrzeroextend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_smgrprefetch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_smgrprefetch
func F_smgrprefetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_smgrwritev github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_smgrwritev
func F_smgrwritev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_RegisterSyncRequest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RegisterSyncRequest
func F_RegisterSyncRequest(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ProcessStartupPacket github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcessStartupPacket
func F_ProcessStartupPacket(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CreateDestReceiver github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CreateDestReceiver
func F_CreateDestReceiver(m *base.Module, l0 int32) int32
//go:linkname F_ProcessInterrupts github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ProcessInterrupts
func F_ProcessInterrupts(m *base.Module)
//go:linkname F_ProcessClientWriteInterrupt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ProcessClientWriteInterrupt
func F_ProcessClientWriteInterrupt(m *base.Module, l0 int32)
//go:linkname F_pg_parse_query github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_parse_query
func F_pg_parse_query(m *base.Module, l0 int32) int32
//go:linkname F_pg_analyze_and_rewrite_fixedparams github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_analyze_and_rewrite_fixedparams
func F_pg_analyze_and_rewrite_fixedparams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_rewrite_query github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_rewrite_query
func F_pg_rewrite_query(m *base.Module, l0 int32) int32
//go:linkname F_pg_analyze_and_rewrite_withcb github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_analyze_and_rewrite_withcb
func F_pg_analyze_and_rewrite_withcb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_plan_query github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_plan_query
func F_pg_plan_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_plan_queries github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_plan_queries
func F_pg_plan_queries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_PostgresMain github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PostgresMain
func F_PostgresMain(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateQueryDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CreateQueryDesc
func F_CreateQueryDesc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_FreeQueryDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeQueryDesc
func F_FreeQueryDesc(m *base.Module, l0 int32)
//go:linkname F_RunFromStore github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RunFromStore
func F_RunFromStore(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int64
//go:linkname F_EnsurePortalSnapshotExists github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EnsurePortalSnapshotExists
func F_EnsurePortalSnapshotExists(m *base.Module)
//go:linkname F_CommandIsReadOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CommandIsReadOnly
func F_CommandIsReadOnly(m *base.Module, l0 int32) int32
//go:linkname F_PreventCommandIfReadOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreventCommandIfReadOnly
func F_PreventCommandIfReadOnly(m *base.Module, l0 int32)
//go:linkname F_PreventCommandIfParallelMode github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreventCommandIfParallelMode
func F_PreventCommandIfParallelMode(m *base.Module, l0 int32)
//go:linkname F_standard_ProcessUtility github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_standard_ProcessUtility
func F_standard_ProcessUtility(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_CreateCommandTag github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateCommandTag
func F_CreateCommandTag(m *base.Module, l0 int32) int32
//go:linkname F_ProcessUtilityForAlterTable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessUtilityForAlterTable
func F_ProcessUtilityForAlterTable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getNextFlagFromString github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_getNextFlagFromString
func F_getNextFlagFromString(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_SplitToVariants github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SplitToVariants
func F_SplitToVariants(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_CheckAffix github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckAffix
func F_CheckAffix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_readstoplist github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_readstoplist
func F_readstoplist(m *base.Module, l0 int32, l1 int32)
//go:linkname F_searchstoplist github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_searchstoplist
func F_searchstoplist(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tt_setup_firstcall github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tt_setup_firstcall
func F_tt_setup_firstcall(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tt_process_call github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tt_process_call
func F_tt_process_call(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_progress_parallel_incr_param github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_progress_parallel_incr_param
func F_pgstat_progress_parallel_incr_param(m *base.Module, l0 int32, l1 int64)
//go:linkname F_pgstat_beinit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_beinit
func F_pgstat_beinit(m *base.Module)
//go:linkname F_pgstat_bestart_initial github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_bestart_initial
func F_pgstat_bestart_initial(m *base.Module)
//go:linkname F_pgstat_bestart_final github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_bestart_final
func F_pgstat_bestart_final(m *base.Module)
//go:linkname F_pgstat_clear_backend_activity_snapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_clear_backend_activity_snapshot
func F_pgstat_clear_backend_activity_snapshot(m *base.Module)
//go:linkname F_pgstat_report_activity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_report_activity
func F_pgstat_report_activity(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_read_current_status github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_read_current_status
func F_pgstat_read_current_status(m *base.Module)
//go:linkname F_pgstat_report_stat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_report_stat
func F_pgstat_report_stat(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_clear_snapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_clear_snapshot
func F_pgstat_clear_snapshot(m *base.Module)
//go:linkname F_pgstat_fetch_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_fetch_entry
func F_pgstat_fetch_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pgstat_build_snapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_build_snapshot
func F_pgstat_build_snapshot(m *base.Module)
//go:linkname F_pgstat_prep_pending_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_prep_pending_entry
func F_pgstat_prep_pending_entry(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pgstat_fetch_stat_backend_by_pid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_fetch_stat_backend_by_pid
func F_pgstat_fetch_stat_backend_by_pid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_fetch_stat_bgwriter github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_fetch_stat_bgwriter
func F_pgstat_fetch_stat_bgwriter(m *base.Module) int32
//go:linkname F_pgstat_report_checkpointer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_report_checkpointer
func F_pgstat_report_checkpointer(m *base.Module)
//go:linkname F_pgstat_report_checksum_failures_in_db github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_report_checksum_failures_in_db
func F_pgstat_report_checksum_failures_in_db(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_fetch_stat_dbentry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_fetch_stat_dbentry
func F_pgstat_fetch_stat_dbentry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_assoc_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_assoc_relation
func F_pgstat_assoc_relation(m *base.Module, l0 int32)
//go:linkname F_pgstat_fetch_stat_tabentry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_fetch_stat_tabentry
func F_pgstat_fetch_stat_tabentry(m *base.Module, l0 int32) int32
//go:linkname F_find_tabstat_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_tabstat_entry
func F_find_tabstat_entry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_release_entry_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_release_entry_ref
func F_pgstat_release_entry_ref(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pgstat_unlock_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_unlock_entry
func F_pgstat_unlock_entry(m *base.Module, l0 int32)
//go:linkname F_pgstat_get_entry_ref_locked github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_get_entry_ref_locked
func F_pgstat_get_entry_ref_locked(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pgstat_drop_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_drop_entry
func F_pgstat_drop_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pgstat_drop_all_entries github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_drop_all_entries
func F_pgstat_drop_all_entries(m *base.Module)
//go:linkname F_pgstat_reset_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_reset_entry
func F_pgstat_reset_entry(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64)
//go:linkname F_pgstat_report_wal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_report_wal
func F_pgstat_report_wal(m *base.Module, l0 int32)
//go:linkname F_AtEOSubXact_PgStat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AtEOSubXact_PgStat
func F_AtEOSubXact_PgStat(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_execute_transactional_drops github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_execute_transactional_drops
func F_pgstat_execute_transactional_drops(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_create_transactional github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgstat_create_transactional
func F_pgstat_create_transactional(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_create_drop_transactional_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_drop_transactional_internal
func F_create_drop_transactional_internal(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_aclcopy github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_aclcopy
func F_aclcopy(m *base.Module, l0 int32) int32
//go:linkname F_aclupdate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_aclupdate
func F_aclupdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_check_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_acl
func F_check_acl(m *base.Module, l0 int32)
//go:linkname F_aclitemsort github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_aclitemsort
func F_aclitemsort(m *base.Module, l0 int32)
//go:linkname F_roles_is_member_of github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_roles_is_member_of
func F_roles_is_member_of(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_has_privs_of_role github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_has_privs_of_role
func F_has_privs_of_role(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_aclmembers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_aclmembers
func F_aclmembers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_convert_any_priv_string github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_any_priv_string
func F_convert_any_priv_string(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_get_role_oid_or_public github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_role_oid_or_public
func F_get_role_oid_or_public(m *base.Module, l0 int32) int32
//go:linkname F_convert_column_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_column_name
func F_convert_column_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_role_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_role_oid
func F_get_role_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_can_set_role github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_can_set_role
func F_check_can_set_role(m *base.Module, l0 int32, l1 int32)
//go:linkname F_select_best_grantor github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_select_best_grantor
func F_select_best_grantor(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_rolespec_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_rolespec_oid
func F_get_rolespec_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rolespec_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rolespec_name
func F_get_rolespec_name(m *base.Module, l0 int32) int32
//go:linkname F_indexam_property github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_indexam_property
func F_indexam_property(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_expand_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_expand_array
func F_expand_array(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DatumGetExpandedArray github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DatumGetExpandedArray
func F_DatumGetExpandedArray(m *base.Module, l0 int32) int32
//go:linkname F_DatumGetAnyArrayP github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DatumGetAnyArrayP
func F_DatumGetAnyArrayP(m *base.Module, l0 int32) int32
//go:linkname F_deconstruct_expanded_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_expanded_array
func F_deconstruct_expanded_array(m *base.Module, l0 int32)
//go:linkname F_mcelem_array_contained_selec github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mcelem_array_contained_selec
func F_mcelem_array_contained_selec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) float64
//go:linkname F_ReadDimensionInt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReadDimensionInt
func F_ReadDimensionInt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CopyArrayEls github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CopyArrayEls
func F_CopyArrayEls(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_construct_empty_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_empty_array
func F_construct_empty_array(m *base.Module, l0 int32) int32
//go:linkname F_array_iter_next github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_iter_next
func F_array_iter_next(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_array_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_recv
func F_array_recv(m *base.Module, l0 int32) int32
//go:linkname F_array_get_element github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_get_element
func F_array_get_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_array_seek github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_seek
func F_array_seek(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_construct_md_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_md_array
func F_construct_md_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_deconstruct_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_array
func F_deconstruct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_array_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_ref
func F_array_ref(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_array_set github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_array_set
func F_array_set(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_construct_array_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_array_builtin
func F_construct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_deconstruct_array_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_array_builtin
func F_deconstruct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_array_contains_nulls github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_contains_nulls
func F_array_contains_nulls(m *base.Module, l0 int32) int32
//go:linkname F_array_eq github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_array_eq
func F_array_eq(m *base.Module, l0 int32) int32
//go:linkname F_initArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_initArrayResult
func F_initArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_accumArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_accumArrayResult
func F_accumArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeArrayResult
func F_makeArrayResult(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ArrayGetNItemsSafe github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ArrayGetNItemsSafe
func F_ArrayGetNItemsSafe(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ArrayCheckBounds github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ArrayCheckBounds
func F_ArrayCheckBounds(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ArrayGetIntegerTypmods github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ArrayGetIntegerTypmods
func F_ArrayGetIntegerTypmods(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_bool_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_bool_with_len
func F_parse_bool_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_cash_mul_float8 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cash_mul_float8
func F_cash_mul_float8(m *base.Module, l0 int64, l1 float64) int64
//go:linkname F_cryptohash_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cryptohash_internal
func F_cryptohash_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_anytime_typmod_check github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_anytime_typmod_check
func F_anytime_typmod_check(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_date2j github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_date2j
func F_date2j(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_j2date github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_j2date
func F_j2date(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_j2day github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_j2day
func F_j2day(m *base.Module, l0 int32) int32
//go:linkname F_DecodeNumberField github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecodeNumberField
func F_DecodeNumberField(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_DetermineTimeZoneOffsetInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DetermineTimeZoneOffsetInternal
func F_DetermineTimeZoneOffsetInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DetermineTimeZoneOffset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DetermineTimeZoneOffset
func F_DetermineTimeZoneOffset(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DetermineTimeZoneAbbrevOffsetTS github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DetermineTimeZoneAbbrevOffsetTS
func F_DetermineTimeZoneAbbrevOffsetTS(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_DecodeTimezoneName github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DecodeTimezoneName
func F_DecodeTimezoneName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EncodeDateTime github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_EncodeDateTime
func F_EncodeDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_datumGetSize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_datumGetSize
func F_datumGetSize(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datumCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_datumCopy
func F_datumCopy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datumTransfer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_datumTransfer
func F_datumTransfer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datumIsEqual github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_datumIsEqual
func F_datumIsEqual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_db_dir_size github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_db_dir_size
func F_db_dir_size(m *base.Module, l0 int32) int64
//go:linkname F_calculate_indexes_size github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_calculate_indexes_size
func F_calculate_indexes_size(m *base.Module, l0 int32) int64
//go:linkname F_domain_state_setup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_domain_state_setup
func F_domain_state_setup(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_domain_check_input github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_domain_check_input
func F_domain_check_input(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_domain_check github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_domain_check
func F_domain_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_enum_range_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_enum_range_internal
func F_enum_range_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EOH_flatten_into github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EOH_flatten_into
func F_EOH_flatten_into(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_make_expanded_record_from_typeid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_expanded_record_from_typeid
func F_make_expanded_record_from_typeid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_expanded_record_from_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_expanded_record_from_tupdesc
func F_make_expanded_record_from_tupdesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_expanded_record_from_exprecord github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_expanded_record_from_exprecord
func F_make_expanded_record_from_exprecord(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expanded_record_fetch_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expanded_record_fetch_tupdesc
func F_expanded_record_fetch_tupdesc(m *base.Module, l0 int32) int32
//go:linkname F_expanded_record_set_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_expanded_record_set_tuple
func F_expanded_record_set_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_build_dummy_expanded_header github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_dummy_expanded_header
func F_build_dummy_expanded_header(m *base.Module, l0 int32)
//go:linkname F_deconstruct_expanded_record github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_deconstruct_expanded_record
func F_deconstruct_expanded_record(m *base.Module, l0 int32)
//go:linkname F_expanded_record_lookup_field github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expanded_record_lookup_field
func F_expanded_record_lookup_field(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_float_overflow_error github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float_overflow_error
func F_float_overflow_error(m *base.Module)
//go:linkname F_float_underflow_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_float_underflow_error
func F_float_underflow_error(m *base.Module)
//go:linkname F_float_zero_divide_error github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float_zero_divide_error
func F_float_zero_divide_error(m *base.Module)
//go:linkname F_float4in_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_float4in_internal
func F_float4in_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32
//go:linkname F_float8in_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_float8in_internal
func F_float8in_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_float8out_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float8out_internal
func F_float8out_internal(m *base.Module, l0 float64) int32
//go:linkname F_format_type_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_format_type_extended
func F_format_type_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_format_type_be github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_format_type_be
func F_format_type_be(m *base.Module, l0 int32) int32
//go:linkname F_format_type_with_typemod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_format_type_with_typemod
func F_format_type_with_typemod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_str_tolower github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_str_tolower
func F_str_tolower(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_parse_format github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_format
func F_parse_format(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_get_th github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_th
func F_get_th(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_int_to_roman github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_int_to_roman
func F_int_to_roman(m *base.Module, l0 int32) int32
//go:linkname F_convert_and_check_filename github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_convert_and_check_filename
func F_convert_and_check_filename(m *base.Module, l0 int32) int32
//go:linkname F_read_binary_file github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_binary_file
func F_read_binary_file(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32
//go:linkname F_pg_stat_file github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_stat_file
func F_pg_stat_file(m *base.Module, l0 int32) int32
//go:linkname F_path_decode github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_path_decode
func F_path_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_path_encode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_path_encode
func F_path_encode(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_box_ar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_box_ar
func F_box_ar(m *base.Module, l0 int32) float64
//go:linkname F_point_sl github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_point_sl
func F_point_sl(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_line_interpt_line github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_line_interpt_line
func F_line_interpt_line(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lseg_closept_lseg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lseg_closept_lseg
func F_lseg_closept_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_lseg_closept_point github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lseg_closept_point
func F_lseg_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_line_closept_point github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_line_closept_point
func F_line_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_line_contain_point github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_line_contain_point
func F_line_contain_point(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lseg_crossing github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lseg_crossing
func F_lseg_crossing(m *base.Module, l0 float64, l1 float64, l2 float64, l3 float64) int32
//go:linkname F_lseg_inside_poly github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lseg_inside_poly
func F_lseg_inside_poly(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_point_div_point github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_point_div_point
func F_point_div_point(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_poly_to_circle github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_poly_to_circle
func F_poly_to_circle(m *base.Module, l0 int32, l1 int32)
//go:linkname F_escape_json github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_escape_json
func F_escape_json(m *base.Module, l0 int32, l1 int32)
//go:linkname F_escape_json_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_escape_json_with_len
func F_escape_json_with_len(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_json_agg_transfn_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_json_agg_transfn_worker
func F_json_agg_transfn_worker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_json_object_agg_transfn_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_json_object_agg_transfn_worker
func F_json_object_agg_transfn_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_JsonbExtractScalar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonbExtractScalar
func F_JsonbExtractScalar(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_jsonb_put_escaped_value github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_jsonb_put_escaped_value
func F_jsonb_put_escaped_value(m *base.Module, l0 int32, l1 int32)
//go:linkname F_jsonb_agg_transfn_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_jsonb_agg_transfn_worker
func F_jsonb_agg_transfn_worker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_jsonb_object_agg_transfn_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_jsonb_object_agg_transfn_worker
func F_jsonb_object_agg_transfn_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_JsonbValueToJsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_JsonbValueToJsonb
func F_JsonbValueToJsonb(m *base.Module, l0 int32) int32
//go:linkname F_appendElement github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_appendElement
func F_appendElement(m *base.Module, l0 int32, l1 int32)
//go:linkname F_JsonbIteratorNext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonbIteratorNext
func F_JsonbIteratorNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_JsonbIteratorInit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_JsonbIteratorInit
func F_JsonbIteratorInit(m *base.Module, l0 int32) int32
//go:linkname F_equalsJsonbScalarValue github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_equalsJsonbScalarValue
func F_equalsJsonbScalarValue(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getKeyJsonValueFromContainer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_getKeyJsonValueFromContainer
func F_getKeyJsonValueFromContainer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_parse_json_or_errsave github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_parse_json_or_errsave
func F_pg_parse_json_or_errsave(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_worker
func F_get_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_jsonb_get_element github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_jsonb_get_element
func F_jsonb_get_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_push_path github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_push_path
func F_push_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_populate_record_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_populate_record_worker
func F_populate_record_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_parse_jsonb_index_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_jsonb_index_flags
func F_parse_jsonb_index_flags(m *base.Module, l0 int32) int32
//go:linkname F_iterate_jsonb_values github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_iterate_jsonb_values
func F_iterate_jsonb_values(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_populate_array_report_expected_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_populate_array_report_expected_array
func F_populate_array_report_expected_array(m *base.Module, l0 int32, l1 int32)
//go:linkname F_jspInitByBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_jspInitByBuffer
func F_jspInitByBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_jspOperationName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jspOperationName
func F_jspOperationName(m *base.Module, l0 int32) int32
//go:linkname F_jspGetNext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_jspGetNext
func F_jspGetNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetJsonTableExecContext github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetJsonTableExecContext
func F_GetJsonTableExecContext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_jsonb_path_query_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_jsonb_path_query_internal
func F_jsonb_path_query_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_executeItemOptUnwrapTarget github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_executeItemOptUnwrapTarget
func F_executeItemOptUnwrapTarget(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_executeAnyItem github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_executeAnyItem
func F_executeAnyItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_patternsel_common github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_patternsel_common
func F_patternsel_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64
//go:linkname F_match_pattern_prefix github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_match_pattern_prefix
func F_match_pattern_prefix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_string_to_const github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_string_to_const
func F_string_to_const(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_regex_selectivity_sub github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_regex_selectivity_sub
func F_regex_selectivity_sub(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_get_multirange_io_data github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_multirange_io_data
func F_get_multirange_io_data(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_multirange
func F_make_multirange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_multirange_get_range github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_multirange_get_range
func F_multirange_get_range(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overlaps_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_overlaps_multirange_internal
func F_range_overlaps_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_multirange_contains_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_multirange_contains_multirange_internal
func F_multirange_contains_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_adjacent_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_adjacent_multirange_internal
func F_range_adjacent_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_multirange_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_multirange_cmp
func F_multirange_cmp(m *base.Module, l0 int32) int32
//go:linkname F_namestrcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_namestrcmp
func F_namestrcmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_network_recv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_network_recv
func F_network_recv(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_inet_spg_consistent_bitmap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_inet_spg_consistent_bitmap
func F_inet_spg_consistent_bitmap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_result_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_result_opt_error
func F_make_result_opt_error(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_var_from_str github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_set_var_from_str
func F_set_var_from_str(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_mul_var github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_mul_var
func F_mul_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_apply_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_apply_typmod
func F_apply_typmod(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_cmp_abs_common github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cmp_abs_common
func F_cmp_abs_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_sub_var github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sub_var
func F_sub_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_add_abs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_add_abs
func F_add_abs(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_compute_bucket github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_compute_bucket
func F_compute_bucket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_numericvar_to_int64 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_numericvar_to_int64
func F_numericvar_to_int64(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_accum_sum_add github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_accum_sum_add
func F_accum_sum_add(m *base.Module, l0 int32, l1 int32)
//go:linkname F_accum_sum_final github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_accum_sum_final
func F_accum_sum_final(m *base.Module, l0 int32, l1 int32)
//go:linkname F_numericvar_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_numericvar_deserialize
func F_numericvar_deserialize(m *base.Module, l0 int32, l1 int32)
//go:linkname F_do_numeric_discard github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_do_numeric_discard
func F_do_numeric_discard(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_numeric_stddev_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_numeric_stddev_internal
func F_numeric_stddev_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_dobyteatrim github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_dobyteatrim
func F_dobyteatrim(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_db_encoding_convert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_db_encoding_convert
func F_db_encoding_convert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_strncoll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_strncoll
func F_pg_strncoll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_strxfrm github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_strxfrm
func F_pg_strxfrm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_pg_locale_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_create_pg_locale_builtin
func F_create_pg_locale_builtin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_collation_actual_version_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_collation_actual_version_builtin
func F_get_collation_actual_version_builtin(m *base.Module, l0 int32) int32
//go:linkname F_strlower_libc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strlower_libc
func F_strlower_libc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_create_pg_locale_libc github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_pg_locale_libc
func F_create_pg_locale_libc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_stat_io_build_tuples github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_stat_io_build_tuples
func F_pg_stat_io_build_tuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64)
//go:linkname F_make_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_range
func F_make_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_range_serialize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_serialize
func F_range_serialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_range_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_deserialize
func F_range_deserialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_range_get_typcache github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_get_typcache
func F_range_get_typcache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_range_cmp_bounds github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_cmp_bounds
func F_range_cmp_bounds(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_contains_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_contains_internal
func F_range_contains_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overlaps_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_range_overlaps_internal
func F_range_overlaps_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overleft_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_overleft_internal
func F_range_overleft_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_minus_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_minus_internal
func F_range_minus_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_empty_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_empty_range
func F_make_empty_range(m *base.Module, l0 int32) int32
//go:linkname F_range_union_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_union_internal
func F_range_union_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_find_simplified_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_find_simplified_clause
func F_find_simplified_clause(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_super_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_super_union
func F_range_super_union(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RE_compile_and_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RE_compile_and_cache
func F_RE_compile_and_cache(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RE_wchar_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RE_wchar_execute
func F_RE_wchar_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_textregexreplace_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_textregexreplace_extended
func F_textregexreplace_extended(m *base.Module, l0 int32) int32
//go:linkname F_setup_regexp_matches github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_setup_regexp_matches
func F_setup_regexp_matches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_regexp_match github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_regexp_match
func F_regexp_match(m *base.Module, l0 int32) int32
//go:linkname F_build_regexp_split_result github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_regexp_split_result
func F_build_regexp_split_result(m *base.Module, l0 int32) int32
//go:linkname F_stringToQualifiedNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_stringToQualifiedNameList
func F_stringToQualifiedNameList(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_format_procedure github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_format_procedure
func F_format_procedure(m *base.Module, l0 int32) int32
//go:linkname F_format_procedure_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_format_procedure_extended
func F_format_procedure_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_format_operator_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_format_operator_extended
func F_format_operator_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RI_FKey_check_ins github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RI_FKey_check_ins
func F_RI_FKey_check_ins(m *base.Module, l0 int32) int32
//go:linkname F_ri_CheckTrigger github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ri_CheckTrigger
func F_ri_CheckTrigger(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ri_FetchConstraintInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ri_FetchConstraintInfo
func F_ri_FetchConstraintInfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ri_restrict github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ri_restrict
func F_ri_restrict(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ri_set github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ri_set
func F_ri_set(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_record_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_record_cmp
func F_record_cmp(m *base.Module, l0 int32) int32
//go:linkname F_quote_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_quote_identifier
func F_quote_identifier(m *base.Module, l0 int32) int32
//go:linkname F_generate_relation_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generate_relation_name
func F_generate_relation_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generate_qualified_relation_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_generate_qualified_relation_name
func F_generate_qualified_relation_name(m *base.Module, l0 int32) int32
//go:linkname F_get_rule_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_rule_expr
func F_get_rule_expr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_query_def github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_query_def
func F_get_query_def(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_pg_get_viewdef_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_get_viewdef_worker
func F_pg_get_viewdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fastgetattr_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fastgetattr_3
func F_fastgetattr_3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_set_rtable_names github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_rtable_names
func F_set_rtable_names(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_generate_function_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_generate_function_name
func F_generate_function_name(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_pg_get_indexdef_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_get_indexdef_worker
func F_pg_get_indexdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_generate_operator_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_operator_name
func F_generate_operator_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_rule_sortgroupclause github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rule_sortgroupclause
func F_get_rule_sortgroupclause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_target_list github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_target_list
func F_get_target_list(m *base.Module, l0 int32, l1 int32)
//go:linkname F_appendContextKeyword github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_appendContextKeyword
func F_appendContextKeyword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_get_rte_alias github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_rte_alias
func F_get_rte_alias(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_values_def github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_values_def
func F_get_values_def(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_rule_orderby github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_rule_orderby
func F_get_rule_orderby(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pg_get_statisticsobj_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_get_statisticsobj_worker
func F_pg_get_statisticsobj_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_quote_qualified_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_quote_qualified_identifier
func F_quote_qualified_identifier(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_deparse_context_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_set_deparse_context_plan
func F_set_deparse_context_plan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_window_frame_options github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_window_frame_options
func F_get_window_frame_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_generate_operator_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generate_operator_clause
func F_generate_operator_clause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_range_partbound_string github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_range_partbound_string
func F_get_range_partbound_string(m *base.Module, l0 int32) int32
//go:linkname F_make_colname_unique github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_make_colname_unique
func F_make_colname_unique(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_printSubscripts github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_printSubscripts
func F_printSubscripts(m *base.Module, l0 int32, l1 int32)
//go:linkname F_resolve_special_varno github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_resolve_special_varno
func F_resolve_special_varno(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_xmltable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_xmltable
func F_get_xmltable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_json_table github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_json_table
func F_get_json_table(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_rule_expr_funccall github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_rule_expr_funccall
func F_get_rule_expr_funccall(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_from_clause_coldeflist github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_from_clause_coldeflist
func F_get_from_clause_coldeflist(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_var_eq_const github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_var_eq_const
func F_var_eq_const(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) float64
//go:linkname F_statistic_proc_security_check github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_statistic_proc_security_check
func F_statistic_proc_security_check(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generic_restriction_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generic_restriction_selectivity
func F_generic_restriction_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) float64
//go:linkname F_examine_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_examine_variable
func F_examine_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_scalarineqsel_wrapper github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scalarineqsel_wrapper
func F_scalarineqsel_wrapper(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_estimate_array_length github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_estimate_array_length
func F_estimate_array_length(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_get_stats_slot_range github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_stats_slot_range
func F_get_stats_slot_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_estimate_num_groups github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_estimate_num_groups
func F_estimate_num_groups(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32) float64
//go:linkname F_estimate_hash_bucket_stats github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_estimate_hash_bucket_stats
func F_estimate_hash_bucket_stats(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32)
//go:linkname F_index_other_operands_eval_cost github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_other_operands_eval_cost
func F_index_other_operands_eval_cost(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_gincost_pattern github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gincost_pattern
func F_gincost_pattern(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_timestamp2tm github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_timestamp2tm
func F_timestamp2tm(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_TimestampDifference github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TimestampDifference
func F_TimestampDifference(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)
//go:linkname F_timestamptz_to_str github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_timestamptz_to_str
func F_timestamptz_to_str(m *base.Module, l0 int64) int32
//go:linkname F_interval_um_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_interval_um_internal
func F_interval_um_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hemdistcache_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hemdistcache_1
func F_hemdistcache_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_findoprnd_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_findoprnd_recurse
func F_findoprnd_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_infix_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_infix_1
func F_infix_1(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_maketree github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_maketree
func F_maketree(m *base.Module, l0 int32) int32
//go:linkname F_plainnode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_plainnode
func F_plainnode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_freetree github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_freetree
func F_freetree(m *base.Module, l0 int32)
//go:linkname F_QT2QTN github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_QT2QTN
func F_QT2QTN(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_QTNodeCompare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_QTNodeCompare
func F_QTNodeCompare(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_QTN2QT github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_QTN2QT
func F_QTN2QT(m *base.Module, l0 int32) int32
//go:linkname F_calc_rank github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_calc_rank
func F_calc_rank(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32
//go:linkname F_TS_execute_ternary github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TS_execute_ternary
func F_TS_execute_ternary(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TS_phrase_output github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TS_phrase_output
func F_TS_phrase_output(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_gen_random_uuid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gen_random_uuid
func F_gen_random_uuid(m *base.Module, l0 int32) int32
//go:linkname F_anybit_typmodin github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_anybit_typmodin
func F_anybit_typmodin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cstring_to_text github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cstring_to_text
func F_cstring_to_text(m *base.Module, l0 int32) int32
//go:linkname F_cstring_to_text_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cstring_to_text_with_len
func F_cstring_to_text_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_text_to_cstring github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_text_to_cstring
func F_text_to_cstring(m *base.Module, l0 int32) int32
//go:linkname F_text_to_cstring_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_text_to_cstring_buffer
func F_text_to_cstring_buffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_byteasend github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_byteasend
func F_byteasend(m *base.Module, l0 int32) int32
//go:linkname F_text_overlay github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_text_overlay
func F_text_overlay(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SplitIdentifierString github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SplitIdentifierString
func F_SplitIdentifierString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_replace_text_regexp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_replace_text_regexp
func F_replace_text_regexp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_split_text github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_split_text
func F_split_text(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_array_to_text_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_to_text_internal
func F_array_to_text_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_varstr_levenshtein github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_varstr_levenshtein
func F_varstr_levenshtein(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_varstr_levenshtein_less_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_varstr_levenshtein_less_equal
func F_varstr_levenshtein_less_equal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_query_to_xml_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_query_to_xml_internal
func F_query_to_xml_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_map_sql_table_to_xmlschema github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_map_sql_table_to_xmlschema
func F_map_sql_table_to_xmlschema(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CatalogCacheInitializeCache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CatalogCacheInitializeCache
func F_CatalogCacheInitializeCache(m *base.Module, l0 int32)
//go:linkname F_SearchCatCache2 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SearchCatCache2
func F_SearchCatCache2(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReleaseCatCache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReleaseCatCache
func F_ReleaseCatCache(m *base.Module, l0 int32)
//go:linkname F_fastgetattr_4 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fastgetattr_4
func F_fastgetattr_4(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ReleaseCatCacheList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReleaseCatCacheList
func F_ReleaseCatCacheList(m *base.Module, l0 int32)
//go:linkname F_compute_function_hashkey github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_compute_function_hashkey
func F_compute_function_hashkey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_InvalidateSystemCachesExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InvalidateSystemCachesExtended
func F_InvalidateSystemCachesExtended(m *base.Module)
//go:linkname F_ProcessCommittedInvalidationMessages github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ProcessCommittedInvalidationMessages
func F_ProcessCommittedInvalidationMessages(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_AtEOSubXact_Inval github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AtEOSubXact_Inval
func F_AtEOSubXact_Inval(m *base.Module, l0 int32)
//go:linkname F_PrepareInvalidationState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PrepareInvalidationState
func F_PrepareInvalidationState(m *base.Module) int32
//go:linkname F_RegisterRelcacheInvalidation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RegisterRelcacheInvalidation
func F_RegisterRelcacheInvalidation(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CacheInvalidateRelcache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CacheInvalidateRelcache
func F_CacheInvalidateRelcache(m *base.Module, l0 int32)
//go:linkname F_CacheInvalidateRelcacheAll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CacheInvalidateRelcacheAll
func F_CacheInvalidateRelcacheAll(m *base.Module)
//go:linkname F_CacheInvalidateRelcacheByTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CacheInvalidateRelcacheByTuple
func F_CacheInvalidateRelcacheByTuple(m *base.Module, l0 int32)
//go:linkname F_get_opfamily_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_opfamily_member
func F_get_opfamily_member(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_opfamily_member_for_cmptype github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_opfamily_member_for_cmptype
func F_get_opfamily_member_for_cmptype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_equality_op_for_ordering_op github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_equality_op_for_ordering_op
func F_get_equality_op_for_ordering_op(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_mergejoin_opfamilies github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_mergejoin_opfamilies
func F_get_mergejoin_opfamilies(m *base.Module, l0 int32) int32
//go:linkname F_get_opfamily_proc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opfamily_proc
func F_get_opfamily_proc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_op_index_interpretation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_op_index_interpretation
func F_get_op_index_interpretation(m *base.Module, l0 int32) int32
//go:linkname F_get_negator github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_negator
func F_get_negator(m *base.Module, l0 int32) int32
//go:linkname F_get_attname github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_attname
func F_get_attname(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_attnum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_attnum
func F_get_attnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_attoptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_attoptions
func F_get_attoptions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_constraint_type github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_constraint_type
func F_get_constraint_type(m *base.Module, l0 int32) int32
//go:linkname F_get_language_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_language_name
func F_get_language_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_opclass_family github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opclass_family
func F_get_opclass_family(m *base.Module, l0 int32) int32
//go:linkname F_get_opclass_input_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_opclass_input_type
func F_get_opclass_input_type(m *base.Module, l0 int32) int32
//go:linkname F_get_opcode github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opcode
func F_get_opcode(m *base.Module, l0 int32) int32
//go:linkname F_op_input_types github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_op_input_types
func F_op_input_types(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_op_mergejoinable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_op_mergejoinable
func F_op_mergejoinable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_op_hashjoinable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_op_hashjoinable
func F_op_hashjoinable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_op_strict github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_op_strict
func F_op_strict(m *base.Module, l0 int32) int32
//go:linkname F_get_commutator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_commutator
func F_get_commutator(m *base.Module, l0 int32) int32
//go:linkname F_get_oprrest github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_oprrest
func F_get_oprrest(m *base.Module, l0 int32) int32
//go:linkname F_get_oprjoin github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_oprjoin
func F_get_oprjoin(m *base.Module, l0 int32) int32
//go:linkname F_get_func_name github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_name
func F_get_func_name(m *base.Module, l0 int32) int32
//go:linkname F_get_func_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_rettype
func F_get_func_rettype(m *base.Module, l0 int32) int32
//go:linkname F_get_func_leakproof github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_leakproof
func F_get_func_leakproof(m *base.Module, l0 int32) int32
//go:linkname F_get_func_support github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_func_support
func F_get_func_support(m *base.Module, l0 int32) int32
//go:linkname F_get_relname_relid github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_relname_relid
func F_get_relname_relid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rel_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rel_name
func F_get_rel_name(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_namespace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_rel_namespace
func F_get_rel_namespace(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_relkind github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_rel_relkind
func F_get_rel_relkind(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_relispartition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_rel_relispartition
func F_get_rel_relispartition(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_persistence github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_rel_persistence
func F_get_rel_persistence(m *base.Module, l0 int32) int32
//go:linkname F_get_typisdefined github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_typisdefined
func F_get_typisdefined(m *base.Module, l0 int32) int32
//go:linkname F_get_typlen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_typlen
func F_get_typlen(m *base.Module, l0 int32) int32
//go:linkname F_get_typlenbyval github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_typlenbyval
func F_get_typlenbyval(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_typlenbyvalalign github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typlenbyvalalign
func F_get_typlenbyvalalign(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_type_io_data github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_type_io_data
func F_get_type_io_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_get_typstorage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_typstorage
func F_get_typstorage(m *base.Module, l0 int32) int32
//go:linkname F_get_typavgwidth github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_typavgwidth
func F_get_typavgwidth(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_typtype github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typtype
func F_get_typtype(m *base.Module, l0 int32) int32
//go:linkname F_type_is_rowtype github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_type_is_rowtype
func F_type_is_rowtype(m *base.Module, l0 int32) int32
//go:linkname F_type_is_enum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_type_is_enum
func F_type_is_enum(m *base.Module, l0 int32) int32
//go:linkname F_type_is_range github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_type_is_range
func F_type_is_range(m *base.Module, l0 int32) int32
//go:linkname F_type_is_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_type_is_multirange
func F_type_is_multirange(m *base.Module, l0 int32) int32
//go:linkname F_get_type_category_preferred github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_type_category_preferred
func F_get_type_category_preferred(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_typ_typrelid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typ_typrelid
func F_get_typ_typrelid(m *base.Module, l0 int32) int32
//go:linkname F_get_element_type github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_element_type
func F_get_element_type(m *base.Module, l0 int32) int32
//go:linkname F_get_array_type github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_array_type
func F_get_array_type(m *base.Module, l0 int32) int32
//go:linkname F_get_base_element_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_base_element_type
func F_get_base_element_type(m *base.Module, l0 int32) int32
//go:linkname F_getTypeInputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getTypeInputInfo
func F_getTypeInputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_getTypeOutputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getTypeOutputInfo
func F_getTypeOutputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_typcollation github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_typcollation
func F_get_typcollation(m *base.Module, l0 int32) int32
//go:linkname F_get_attavgwidth github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_get_attavgwidth
func F_get_attavgwidth(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_attstatsslot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_attstatsslot
func F_get_attstatsslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_free_attstatsslot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_free_attstatsslot
func F_free_attstatsslot(m *base.Module, l0 int32)
//go:linkname F_get_namespace_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_namespace_name
func F_get_namespace_name(m *base.Module, l0 int32) int32
//go:linkname F_get_range_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_range_multirange
func F_get_range_multirange(m *base.Module, l0 int32) int32
//go:linkname F_get_multirange_range github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_multirange_range
func F_get_multirange_range(m *base.Module, l0 int32) int32
//go:linkname F_get_index_isvalid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_index_isvalid
func F_get_index_isvalid(m *base.Module, l0 int32) int32
//go:linkname F_get_index_isclustered github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_index_isclustered
func F_get_index_isclustered(m *base.Module, l0 int32) int32
//go:linkname F_get_publication_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_publication_name
func F_get_publication_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_subscription_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_subscription_name
func F_get_subscription_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationGetPartitionKey github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetPartitionKey
func F_RelationGetPartitionKey(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetPartitionQual github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetPartitionQual
func F_RelationGetPartitionQual(m *base.Module, l0 int32) int32
//go:linkname F_get_partition_qual_relid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_partition_qual_relid
func F_get_partition_qual_relid(m *base.Module, l0 int32) int32
//go:linkname F_CompleteCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CompleteCachedPlan
func F_CompleteCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_SaveCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SaveCachedPlan
func F_SaveCachedPlan(m *base.Module, l0 int32)
//go:linkname F_DropCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DropCachedPlan
func F_DropCachedPlan(m *base.Module, l0 int32)
//go:linkname F_GetCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetCachedPlan
func F_GetCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ReleaseCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReleaseCachedPlan
func F_ReleaseCachedPlan(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CachedPlanAllowsSimpleValidityCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CachedPlanAllowsSimpleValidityCheck
func F_CachedPlanAllowsSimpleValidityCheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RelationRebuildRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationRebuildRelation
func F_RelationRebuildRelation(m *base.Module, l0 int32)
//go:linkname F_RelationBuildDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationBuildDesc
func F_RelationBuildDesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationIncrementReferenceCount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationIncrementReferenceCount
func F_RelationIncrementReferenceCount(m *base.Module, l0 int32)
//go:linkname F_RelationInitPhysicalAddr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationInitPhysicalAddr
func F_RelationInitPhysicalAddr(m *base.Module, l0 int32)
//go:linkname F_RelationDestroyRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationDestroyRelation
func F_RelationDestroyRelation(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AtEOSubXact_RelationCache github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_AtEOSubXact_RelationCache
func F_AtEOSubXact_RelationCache(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_write_item github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_write_item
func F_write_item(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RelationGetFKeyList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationGetFKeyList
func F_RelationGetFKeyList(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexList github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationGetIndexList
func F_RelationGetIndexList(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetReplicaIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationGetReplicaIndex
func F_RelationGetReplicaIndex(m *base.Module, l0 int32) int32
//go:linkname F_errtableconstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errtableconstraint
func F_errtableconstraint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_unlink_initfile github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_unlink_initfile
func F_unlink_initfile(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationCacheInitFileRemoveInDir github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_RelationCacheInitFileRemoveInDir
func F_RelationCacheInitFileRemoveInDir(m *base.Module, l0 int32)
//go:linkname F_RelidByRelfilenumber github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelidByRelfilenumber
func F_RelidByRelfilenumber(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationMapUpdateMap github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationMapUpdateMap
func F_RelationMapUpdateMap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
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
//go:linkname F_SearchSysCacheLocked1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchSysCacheLocked1
func F_SearchSysCacheLocked1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchSysCacheCopy
func F_SearchSysCacheCopy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchSysCacheLockedCopy1 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCacheLockedCopy1
func F_SearchSysCacheLockedCopy1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheExists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCacheExists
func F_SearchSysCacheExists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_GetSysCacheOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetSysCacheOid
func F_GetSysCacheOid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SearchSysCacheCopyAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SearchSysCacheCopyAttName
func F_SearchSysCacheCopyAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SysCacheGetAttr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SysCacheGetAttr
func F_SysCacheGetAttr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SysCacheGetAttrNotNull github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SysCacheGetAttrNotNull
func F_SysCacheGetAttrNotNull(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetSysCacheHashValue github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_GetSysCacheHashValue
func F_GetSysCacheHashValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchSysCacheList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SearchSysCacheList
func F_SearchSysCacheList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_getTSCurrentConfig github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_getTSCurrentConfig
func F_getTSCurrentConfig(m *base.Module) int32
//go:linkname F_lookup_type_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lookup_type_cache
func F_lookup_type_cache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookup_rowtype_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_lookup_rowtype_tupdesc
func F_lookup_rowtype_tupdesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookup_rowtype_tupdesc_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_lookup_rowtype_tupdesc_copy
func F_lookup_rowtype_tupdesc_copy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assign_record_type_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_assign_record_type_typmod
func F_assign_record_type_typmod(m *base.Module, l0 int32)
//go:linkname F_AtEOXact_TypeCache github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AtEOXact_TypeCache
func F_AtEOXact_TypeCache(m *base.Module)
//go:linkname F_errstart_cold github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errstart_cold
func F_errstart_cold(m *base.Module, l0 int32, l1 int32)
//go:linkname F_write_stderr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_write_stderr
func F_write_stderr(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errmsg_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errmsg_internal
func F_errmsg_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errfinish github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errfinish
func F_errfinish(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_EmitErrorReport github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_EmitErrorReport
func F_EmitErrorReport(m *base.Module)
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
//go:linkname F_errcode_for_socket_access github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errcode_for_socket_access
func F_errcode_for_socket_access(m *base.Module)
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
//go:linkname F_errdetail_plural github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errdetail_plural
func F_errdetail_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errhint github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_errhint
func F_errhint(m *base.Module, l0 int32, l1 int32)
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
//go:linkname F_internalerrquery github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_internalerrquery
func F_internalerrquery(m *base.Module, l0 int32) int32
//go:linkname F_geterrposition github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_geterrposition
func F_geterrposition(m *base.Module) int32
//go:linkname F_format_elog_string github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_format_elog_string
func F_format_elog_string(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expand_dynamic_library_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_expand_dynamic_library_name
func F_expand_dynamic_library_name(m *base.Module, l0 int32) int32
//go:linkname F_internal_load_library github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_internal_load_library
func F_internal_load_library(m *base.Module, l0 int32) int32
//go:linkname F_fmgr_info github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fmgr_info
func F_fmgr_info(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fmgr_info_cxt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fmgr_info_cxt
func F_fmgr_info_cxt(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_detoast_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_detoast_datum
func F_pg_detoast_datum(m *base.Module, l0 int32) int32
//go:linkname F_DirectFunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DirectFunctionCall1Coll
func F_DirectFunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DirectFunctionCall2Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DirectFunctionCall2Coll
func F_DirectFunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_DirectFunctionCall3Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DirectFunctionCall3Coll
func F_DirectFunctionCall3Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_DirectFunctionCall4Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DirectFunctionCall4Coll
func F_DirectFunctionCall4Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_DirectFunctionCall5Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DirectFunctionCall5Coll
func F_DirectFunctionCall5Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_FunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FunctionCall1Coll
func F_FunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FunctionCall2Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_FunctionCall2Coll
func F_FunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_FunctionCall3Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FunctionCall3Coll
func F_FunctionCall3Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_FunctionCall4Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FunctionCall4Coll
func F_FunctionCall4Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_OidFunctionCall0Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_OidFunctionCall0Coll
func F_OidFunctionCall0Coll(m *base.Module, l0 int32) int32
//go:linkname F_OidFunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_OidFunctionCall1Coll
func F_OidFunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_InputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InputFunctionCall
func F_InputFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_InputFunctionCallSafe github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InputFunctionCallSafe
func F_InputFunctionCallSafe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_DirectInputFunctionCallSafe github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DirectInputFunctionCallSafe
func F_DirectInputFunctionCallSafe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_OutputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OutputFunctionCall
func F_OutputFunctionCall(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReceiveFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReceiveFunctionCall
func F_ReceiveFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_OidInputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_OidInputFunctionCall
func F_OidInputFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_OidOutputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OidOutputFunctionCall
func F_OidOutputFunctionCall(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_Int64GetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Int64GetDatum
func F_Int64GetDatum(m *base.Module, l0 int64) int32
//go:linkname F_Float8GetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_Float8GetDatum
func F_Float8GetDatum(m *base.Module, l0 float64) int32
//go:linkname F_pg_detoast_datum_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_detoast_datum_copy
func F_pg_detoast_datum_copy(m *base.Module, l0 int32) int32
//go:linkname F_pg_detoast_datum_packed github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_detoast_datum_packed
func F_pg_detoast_datum_packed(m *base.Module, l0 int32) int32
//go:linkname F_get_fn_expr_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_fn_expr_rettype
func F_get_fn_expr_rettype(m *base.Module, l0 int32) int32
//go:linkname F_get_fn_expr_argtype github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_fn_expr_argtype
func F_get_fn_expr_argtype(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_call_expr_argtype github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_call_expr_argtype
func F_get_call_expr_argtype(m *base.Module, l0 int32, l1 int32) int32
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
//go:linkname F_end_MultiFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_end_MultiFuncCall
func F_end_MultiFuncCall(m *base.Module, l0 int32)
//go:linkname F_build_function_result_tupdesc_t github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_function_result_tupdesc_t
func F_build_function_result_tupdesc_t(m *base.Module, l0 int32) int32
//go:linkname F_resolve_anyelement_from_others github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_resolve_anyelement_from_others
func F_resolve_anyelement_from_others(m *base.Module, l0 int32)
//go:linkname F_resolve_anyrange_from_others github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_resolve_anyrange_from_others
func F_resolve_anyrange_from_others(m *base.Module, l0 int32)
//go:linkname F_resolve_anymultirange_from_others github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_resolve_anymultirange_from_others
func F_resolve_anymultirange_from_others(m *base.Module, l0 int32)
//go:linkname F_get_expr_result_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_expr_result_tupdesc
func F_get_expr_result_tupdesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_func_arg_info github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_func_arg_info
func F_get_func_arg_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_func_input_arg_names github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_func_input_arg_names
func F_get_func_input_arg_names(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_function_result_tupdesc_d github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_function_result_tupdesc_d
func F_build_function_result_tupdesc_d(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_extract_variadic_args github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_extract_variadic_args
func F_extract_variadic_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_create github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_create
func F_hash_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_search github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_search
func F_hash_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_search_with_hash_value github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_search_with_hash_value
func F_hash_search_with_hash_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_hash_seq_init github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_seq_init
func F_hash_seq_init(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hash_seq_search github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_seq_search
func F_hash_seq_search(m *base.Module, l0 int32) int32
//go:linkname F_AtEOSubXact_HashTables github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AtEOSubXact_HashTables
func F_AtEOSubXact_HashTables(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SwitchToSharedLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SwitchToSharedLatch
func F_SwitchToSharedLatch(m *base.Module)
//go:linkname F_load_libraries github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_load_libraries
func F_load_libraries(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_BaseInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BaseInit
func F_BaseInit(m *base.Module)
//go:linkname F_InitPostgres github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InitPostgres
func F_InitPostgres(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_UtfToLocal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UtfToLocal
func F_UtfToLocal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_pg_do_encoding_conversion github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_do_encoding_conversion
func F_pg_do_encoding_conversion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_report_invalid_encoding github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_invalid_encoding
func F_report_invalid_encoding(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_verify_mbstr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_verify_mbstr
func F_pg_verify_mbstr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_any_to_server github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_any_to_server
func F_pg_any_to_server(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_perform_default_encoding_conversion github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_perform_default_encoding_conversion
func F_perform_default_encoding_conversion(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_unicode_to_server github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_unicode_to_server
func F_pg_unicode_to_server(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_mb2wchar_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_mb2wchar_with_len
func F_pg_mb2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_wchar2mb_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_wchar2mb_with_len
func F_pg_wchar2mb_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_report_invalid_encoding_db github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_report_invalid_encoding_db
func F_report_invalid_encoding_db(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_mblen_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mblen_range
func F_pg_mblen_range(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mblen_unbounded github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_mblen_unbounded
func F_pg_mblen_unbounded(m *base.Module, l0 int32) int32
//go:linkname F_pg_mbstrlen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_mbstrlen
func F_pg_mbstrlen(m *base.Module, l0 int32) int32
//go:linkname F_pg_mbstrlen_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mbstrlen_with_len
func F_pg_mbstrlen_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mbcliplen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mbcliplen
func F_pg_mbcliplen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_encoding_conversion_args github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_encoding_conversion_args
func F_check_encoding_conversion_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_report_untranslatable_char github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_report_untranslatable_char
func F_report_untranslatable_char(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_appendStringInfoStringQuoted github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_appendStringInfoStringQuoted
func F_appendStringInfoStringQuoted(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ProcessConfigFileInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ProcessConfigFileInternal
func F_ProcessConfigFileInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_option github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_option
func F_find_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_set_config_with_handle github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_set_config_with_handle
func F_set_config_with_handle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_assignable_custom_variable_name github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_assignable_custom_variable_name
func F_assignable_custom_variable_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SetConfigOption github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SetConfigOption
func F_SetConfigOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_guc_malloc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_guc_malloc
func F_guc_malloc(m *base.Module, l0 int32) int32
//go:linkname F_get_guc_variables github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_guc_variables
func F_get_guc_variables(m *base.Module, l0 int32) int32
//go:linkname F_convert_GUC_name_for_parameter_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_convert_GUC_name_for_parameter_acl
func F_convert_GUC_name_for_parameter_acl(m *base.Module, l0 int32) int32
//go:linkname F_AtEOXact_GUC github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AtEOXact_GUC
func F_AtEOXact_GUC(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ShowGUCOption github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ShowGUCOption
func F_ShowGUCOption(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_config_unit_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_config_unit_name
func F_get_config_unit_name(m *base.Module, l0 int32) int32
//go:linkname F_parse_int github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_int
func F_parse_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_parse_real github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_parse_real
func F_parse_real(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_config_enum_lookup_by_value github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_config_enum_lookup_by_value
func F_config_enum_lookup_by_value(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_config_enum_get_options github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_config_enum_get_options
func F_config_enum_get_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_parse_and_validate_value github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parse_and_validate_value
func F_parse_and_validate_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_set_config_option_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_set_config_option_ext
func F_set_config_option_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_DefineCustomBoolVariable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_DefineCustomBoolVariable
func F_DefineCustomBoolVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_init_custom_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_init_custom_variable
func F_init_custom_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_define_custom_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_define_custom_variable
func F_define_custom_variable(m *base.Module, l0 int32)
//go:linkname F_DefineCustomStringVariable github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DefineCustomStringVariable
func F_DefineCustomStringVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_DefineCustomEnumVariable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_DefineCustomEnumVariable
func F_DefineCustomEnumVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_MarkGUCPrefixReserved github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MarkGUCPrefixReserved
func F_MarkGUCPrefixReserved(m *base.Module, l0 int32)
//go:linkname F_do_serialize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_do_serialize
func F_do_serialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_TransformGUCArray github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TransformGUCArray
func F_TransformGUCArray(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_validate_option_array_item github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_validate_option_array_item
func F_validate_option_array_item(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GUC_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GUC_yyensure_buffer_stack
func F_GUC_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_GUC_yy_create_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GUC_yy_create_buffer
func F_GUC_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GUC_flex_fatal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GUC_flex_fatal
func F_GUC_flex_fatal(m *base.Module, l0 int32)
//go:linkname F_ParseConfigFp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ParseConfigFp
func F_ParseConfigFp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_FreeConfigVariables github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeConfigVariables
func F_FreeConfigVariables(m *base.Module, l0 int32)
//go:linkname F_ExtractSetVariableArgs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ExtractSetVariableArgs
func F_ExtractSetVariableArgs(m *base.Module, l0 int32) int32
//go:linkname F_pg_rusage_show github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_rusage_show
func F_pg_rusage_show(m *base.Module, l0 int32) int32
//go:linkname F_check_enable_rls github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_enable_rls
func F_check_enable_rls(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_stack_is_too_deep github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_stack_is_too_deep
func F_stack_is_too_deep(m *base.Module) int32
//go:linkname F_superuser github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_superuser
func F_superuser(m *base.Module) int32
//go:linkname F_superuser_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_superuser_arg
func F_superuser_arg(m *base.Module, l0 int32) int32
//go:linkname F_schedule_alarm github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_schedule_alarm
func F_schedule_alarm(m *base.Module, l0 int64)
//go:linkname F_RegisterTimeout github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RegisterTimeout
func F_RegisterTimeout(m *base.Module, l0 int32, l1 int32)
//go:linkname F_enable_timeout_after github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_enable_timeout_after
func F_enable_timeout_after(m *base.Module, l0 int32, l1 int32)
//go:linkname F_disable_timeout github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_disable_timeout
func F_disable_timeout(m *base.Module, l0 int32)
//go:linkname F_AllocSetContextCreateInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AllocSetContextCreateInternal
func F_AllocSetContextCreateInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_create_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_internal
func F_create_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_dsa_create_in_place_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dsa_create_in_place_ext
func F_dsa_create_in_place_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_dsa_allocate_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsa_allocate_extended
func F_dsa_allocate_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_for_freed_segments_locked github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_check_for_freed_segments_locked
func F_check_for_freed_segments_locked(m *base.Module, l0 int32)
//go:linkname F_get_segment_by_index github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_segment_by_index
func F_get_segment_by_index(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rebin_segment github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_rebin_segment
func F_rebin_segment(m *base.Module, l0 int32, l1 int32)
//go:linkname F_unlink_span github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_unlink_span
func F_unlink_span(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dsa_get_address github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsa_get_address
func F_dsa_get_address(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FreePageManagerGet github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreePageManagerGet
func F_FreePageManagerGet(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FreePageManagerPut github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreePageManagerPut
func F_FreePageManagerPut(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_FreePageBtreeRemovePage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreePageBtreeRemovePage
func F_FreePageBtreeRemovePage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_FreePageBtreeConsolidate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreePageBtreeConsolidate
func F_FreePageBtreeConsolidate(m *base.Module, l0 int32, l1 int32)
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
//go:linkname F_MemoryContextMemConsumed github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextMemConsumed
func F_MemoryContextMemConsumed(m *base.Module, l0 int32, l1 int32)
//go:linkname F_MemoryContextStats github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextStats
func F_MemoryContextStats(m *base.Module, l0 int32)
//go:linkname F_MemoryContextAllocationFailure github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextAllocationFailure
func F_MemoryContextAllocationFailure(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MemoryContextSizeFailure github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextSizeFailure
func F_MemoryContextSizeFailure(m *base.Module, l0 int32)
//go:linkname F_MemoryContextAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextAlloc
func F_MemoryContextAlloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextAllocZero github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoryContextAllocZero
func F_MemoryContextAllocZero(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextAllocExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_MemoryContextAllocExtended
func F_MemoryContextAllocExtended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_palloc_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_palloc_extended
func F_palloc_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_palloc_aligned github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_palloc_aligned
func F_palloc_aligned(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pfree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pfree
func F_pfree(m *base.Module, l0 int32)
//go:linkname F_repalloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_repalloc
func F_repalloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_repalloc_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_repalloc_extended
func F_repalloc_extended(m *base.Module, l0 int32, l1 int32) int32
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
//go:linkname F_MarkPortalFailed github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MarkPortalFailed
func F_MarkPortalFailed(m *base.Module, l0 int32)
//go:linkname F_ResourceOwnerRemember github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResourceOwnerRemember
func F_ResourceOwnerRemember(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ResourceOwnerForget github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResourceOwnerForget
func F_ResourceOwnerForget(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ResourceOwnerReleaseAll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ResourceOwnerReleaseAll
func F_ResourceOwnerReleaseAll(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ResourceOwnerDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ResourceOwnerDelete
func F_ResourceOwnerDelete(m *base.Module, l0 int32)
//go:linkname F_CreateAuxProcessResourceOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CreateAuxProcessResourceOwner
func F_CreateAuxProcessResourceOwner(m *base.Module)
//go:linkname F_LogicalTapeCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalTapeCreate
func F_LogicalTapeCreate(m *base.Module, l0 int32) int32
//go:linkname F_LogicalTapeWrite github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_LogicalTapeWrite
func F_LogicalTapeWrite(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LogicalTapeRead github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalTapeRead
func F_LogicalTapeRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ltsReadFillBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltsReadFillBuffer
func F_ltsReadFillBuffer(m *base.Module, l0 int32) int32
//go:linkname F_qsort_interruptible_med3 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_qsort_interruptible_med3
func F_qsort_interruptible_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_tuplesort_end github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_end
func F_tuplesort_end(m *base.Module, l0 int32)
//go:linkname F_tuplesort_reset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_reset
func F_tuplesort_reset(m *base.Module, l0 int32)
//go:linkname F_tuplesort_sort_memtuples github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_sort_memtuples
func F_tuplesort_sort_memtuples(m *base.Module, l0 int32)
//go:linkname F_tuplesort_performsort github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_performsort
func F_tuplesort_performsort(m *base.Module, l0 int32)
//go:linkname F_tuplesort_merge_order github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplesort_merge_order
func F_tuplesort_merge_order(m *base.Module, l0 int64) int32
//go:linkname F_tuplesort_rescan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_rescan
func F_tuplesort_rescan(m *base.Module, l0 int32)
//go:linkname F_tuplesort_readtup_alloc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplesort_readtup_alloc
func F_tuplesort_readtup_alloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tuplesort_estimate_shared github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_estimate_shared
func F_tuplesort_estimate_shared(m *base.Module, l0 int32) int32
//go:linkname F_tuplesort_initialize_shared github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_initialize_shared
func F_tuplesort_initialize_shared(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_begin_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_begin_heap
func F_tuplesort_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_comparetup_heap_tiebreak github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_comparetup_heap_tiebreak
func F_comparetup_heap_tiebreak(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplesort_begin_index_brin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_begin_index_brin
func F_tuplesort_begin_index_brin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tuplesort_puttupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_puttupleslot
func F_tuplesort_puttupleslot(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplesort_putindextuplevalues github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplesort_putindextuplevalues
func F_tuplesort_putindextuplevalues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_tuplesort_putbrintuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_putbrintuple
func F_tuplesort_putbrintuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_gettupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_gettupleslot
func F_tuplesort_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_tuplestore_begin_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_begin_heap
func F_tuplestore_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplestore_end github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplestore_end
func F_tuplestore_end(m *base.Module, l0 int32)
//go:linkname F_tuplestore_select_read_pointer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplestore_select_read_pointer
func F_tuplestore_select_read_pointer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_puttuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplestore_puttuple
func F_tuplestore_puttuple(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_putvalues github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplestore_putvalues
func F_tuplestore_putvalues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_tuplestore_gettupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_gettupleslot
func F_tuplestore_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_tuplestore_skiptuples github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_skiptuples
func F_tuplestore_skiptuples(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_tuplestore_rescan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_rescan
func F_tuplestore_rescan(m *base.Module, l0 int32)
//go:linkname F_tuplestore_copy_read_pointer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplestore_copy_read_pointer
func F_tuplestore_copy_read_pointer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplestore_get_stats github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tuplestore_get_stats
func F_tuplestore_get_stats(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetTransactionSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetTransactionSnapshot
func F_GetTransactionSnapshot(m *base.Module) int32
//go:linkname F_GetLatestSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetLatestSnapshot
func F_GetLatestSnapshot(m *base.Module) int32
//go:linkname F_PushActiveSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PushActiveSnapshot
func F_PushActiveSnapshot(m *base.Module, l0 int32)
//go:linkname F_PushActiveSnapshotWithLevel github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_PushActiveSnapshotWithLevel
func F_PushActiveSnapshotWithLevel(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PushCopiedSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PushCopiedSnapshot
func F_PushCopiedSnapshot(m *base.Module, l0 int32)
//go:linkname F_UpdateActiveSnapshotCommandId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UpdateActiveSnapshotCommandId
func F_UpdateActiveSnapshotCommandId(m *base.Module)
//go:linkname F_PopActiveSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PopActiveSnapshot
func F_PopActiveSnapshot(m *base.Module)
//go:linkname F_RegisterSnapshotOnOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RegisterSnapshotOnOwner
func F_RegisterSnapshotOnOwner(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_UnregisterSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_UnregisterSnapshot
func F_UnregisterSnapshot(m *base.Module, l0 int32)
//go:linkname F_UnregisterSnapshotNoOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_UnregisterSnapshotNoOwner
func F_UnregisterSnapshotNoOwner(m *base.Module, l0 int32)
//go:linkname F_EstimateSnapshotSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_EstimateSnapshotSpace
func F_EstimateSnapshotSpace(m *base.Module, l0 int32) int32
//go:linkname F_XidInMVCCSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XidInMVCCSnapshot
func F_XidInMVCCSnapshot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tzload github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tzload
func F_tzload(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getrule github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_getrule
func F_getrule(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_localtime github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_localtime
func F_pg_localtime(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_tz_acceptable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_tz_acceptable
func F_pg_tz_acceptable(m *base.Module, l0 int32) int32
//go:linkname F_pg_tzset github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_tzset
func F_pg_tzset(m *base.Module, l0 int32) int32
//go:linkname F_pg_strftime github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_strftime
func F_pg_strftime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_jit_compile_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jit_compile_expr
func F_jit_compile_expr(m *base.Module, l0 int32) int32
//go:linkname F_binaryheap_allocate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_binaryheap_allocate
func F_binaryheap_allocate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_binaryheap_add_unordered github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_binaryheap_add_unordered
func F_binaryheap_add_unordered(m *base.Module, l0 int32, l1 int32)
//go:linkname F_binaryheap_build github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_binaryheap_build
func F_binaryheap_build(m *base.Module, l0 int32)
//go:linkname F_binaryheap_remove_first github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_binaryheap_remove_first
func F_binaryheap_remove_first(m *base.Module, l0 int32) int32
//go:linkname F_binaryheap_replace_first github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_binaryheap_replace_first
func F_binaryheap_replace_first(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_checksum_init github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_checksum_init
func F_pg_checksum_init(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_controlfile github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_controlfile
func F_get_controlfile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_dirent_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_dirent_type
func F_get_dirent_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_bytes github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hash_bytes
func F_hash_bytes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_getaddrinfo_all github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_getaddrinfo_all
func F_pg_getaddrinfo_all(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_getnameinfo_all github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_getnameinfo_all
func F_pg_getnameinfo_all(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_json_lex github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_json_lex
func F_json_lex(m *base.Module, l0 int32) int32
//go:linkname F_inc_lex_level github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_inc_lex_level
func F_inc_lex_level(m *base.Module, l0 int32)
//go:linkname F_dec_lex_level github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dec_lex_level
func F_dec_lex_level(m *base.Module, l0 int32)
//go:linkname F_json_errdetail github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_json_errdetail
func F_json_errdetail(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_md5_encrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_md5_encrypt
func F_pg_md5_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pglz_compress github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pglz_compress
func F_pglz_compress(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_prng_seed github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_prng_seed
func F_pg_prng_seed(m *base.Module, l0 int32, l1 int64)
//go:linkname F_psprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_psprintf
func F_psprintf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pvsnprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pvsnprintf
func F_pvsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_forkname_chars github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_forkname_chars
func F_forkname_chars(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetRelationPath github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetRelationPath
func F_GetRelationPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_pg_saslprep github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_saslprep
func F_pg_saslprep(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_initStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initStringInfo
func F_initStringInfo(m *base.Module, l0 int32)
//go:linkname F_resetStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_resetStringInfo
func F_resetStringInfo(m *base.Module, l0 int32)
//go:linkname F_appendStringInfoString github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendStringInfoString
func F_appendStringInfoString(m *base.Module, l0 int32, l1 int32)
//go:linkname F_appendBinaryStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendBinaryStringInfo
func F_appendBinaryStringInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_appendStringInfoChar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_appendStringInfoChar
func F_appendStringInfoChar(m *base.Module, l0 int32, l1 int32)
//go:linkname F_appendStringInfoSpaces github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_appendStringInfoSpaces
func F_appendStringInfoSpaces(m *base.Module, l0 int32, l1 int32)
//go:linkname F_convert_case github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_convert_case
func F_convert_case(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_pg_u_isdigit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_u_isdigit
func F_pg_u_isdigit(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_u_isalnum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_u_isalnum
func F_pg_u_isalnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_u_isspace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_u_isspace
func F_pg_u_isspace(m *base.Module, l0 int32) int32
//go:linkname F_wait_result_to_str github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_wait_result_to_str
func F_wait_result_to_str(m *base.Module, l0 int32) int32
//go:linkname F_pg_ascii_verifystr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pg_ascii_verifystr
func F_pg_ascii_verifystr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_encoding_verifymbchar github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_encoding_verifymbchar
func F_pg_encoding_verifymbchar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_cryptohash_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_cryptohash_create
func F_pg_cryptohash_create(m *base.Module, l0 int32) int32
//go:linkname F_pg_cryptohash_init github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_cryptohash_init
func F_pg_cryptohash_init(m *base.Module, l0 int32) int32
//go:linkname F_pg_cryptohash_update github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_cryptohash_update
func F_pg_cryptohash_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32
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
//go:linkname F_pg_get_encoding_from_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_get_encoding_from_locale
func F_pg_get_encoding_from_locale(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_absolute_path github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_absolute_path
func F_make_absolute_path(m *base.Module, l0 int32) int32
//go:linkname F_get_share_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_share_path
func F_get_share_path(m *base.Module, l0 int32)
//go:linkname F_pg_strong_random github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_strong_random
func F_pg_strong_random(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_usleep github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_usleep
func F_pg_usleep(m *base.Module, l0 int32)
//go:linkname F_pg_strcasecmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_strcasecmp
func F_pg_strcasecmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_strncasecmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_strncasecmp
func F_pg_strncasecmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pqsignal_be github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pqsignal_be
func F_pqsignal_be(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_qsort github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_qsort
func F_pg_qsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_qsort_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_qsort_arg
func F_qsort_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_dopr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dopr
func F_dopr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_snprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_snprintf
func F_pg_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_sprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_sprintf
func F_pg_sprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_fprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_fprintf
func F_pg_fprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgl_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgl_exit
func F_pgl_exit(m *base.Module, l0 int32)
//go:linkname F_pgmem_kill github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgmem_kill
func F_pgmem_kill(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgmem_shmat github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgmem_shmat
func F_pgmem_shmat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgmem_shmdt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgmem_shmdt
func F_pgmem_shmdt(m *base.Module, l0 int32) int32
//go:linkname F_pgmem_shmctl github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgmem_shmctl
func F_pgmem_shmctl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgmem_send github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgmem_send
func F_pgmem_send(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgmem_sigprocmask github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgmem_sigprocmask
func F_pgmem_sigprocmask(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgmem_dlsym github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgmem_dlsym
func F_pgmem_dlsym(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_compile github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plpgsql_compile
func F_plpgsql_compile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_build_datatype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_datatype
func F_build_datatype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_plpgsql_build_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_plpgsql_build_variable
func F_plpgsql_build_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_add_dummy_return github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_add_dummy_return
func F_add_dummy_return(m *base.Module, l0 int32)
//go:linkname F_resolve_column_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_resolve_column_ref
func F_resolve_column_ref(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_plpgsql_adddatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_plpgsql_adddatum
func F_plpgsql_adddatum(m *base.Module, l0 int32)
//go:linkname F_plpgsql_exec_function github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_plpgsql_exec_function
func F_plpgsql_exec_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_assign_simple_var github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_assign_simple_var
func F_assign_simple_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_exec_move_row github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_exec_move_row
func F_exec_move_row(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_revalidate_rectypeid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_revalidate_rectypeid
func F_revalidate_rectypeid(m *base.Module, l0 int32)
//go:linkname F_assign_record_var github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_assign_record_var
func F_assign_record_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_exec_move_row_from_fields github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_exec_move_row_from_fields
func F_exec_move_row_from_fields(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_plpgsql_exec_trigger github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_exec_trigger
func F_plpgsql_exec_trigger(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_exec_event_trigger github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plpgsql_exec_event_trigger
func F_plpgsql_exec_event_trigger(m *base.Module, l0 int32, l1 int32)
//go:linkname F_instantiate_empty_record_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_instantiate_empty_record_variable
func F_instantiate_empty_record_variable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_plpgsql_exec_get_datum_type_info github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_exec_get_datum_type_info
func F_plpgsql_exec_get_datum_type_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_exec_eval_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exec_eval_expr
func F_exec_eval_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_exec_save_simple_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_exec_save_simple_expr
func F_exec_save_simple_expr(m *base.Module, l0 int32, l1 int32)
//go:linkname F_plpgsql_ns_push github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_ns_push
func F_plpgsql_ns_push(m *base.Module, l0 int32, l1 int32)
//go:linkname F_plpgsql_ns_additem github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_ns_additem
func F_plpgsql_ns_additem(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_plpgsql_mark_local_assignment_targets github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_mark_local_assignment_targets
func F_plpgsql_mark_local_assignment_targets(m *base.Module, l0 int32)
//go:linkname F_mark_stmt github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mark_stmt
func F_mark_stmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F_plpgsql_free_function_memory github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plpgsql_free_function_memory
func F_plpgsql_free_function_memory(m *base.Module, l0 int32)
//go:linkname F_plpgsql_dumptree github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_dumptree
func F_plpgsql_dumptree(m *base.Module, l0 int32)
//go:linkname F_dump_block github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_dump_block
func F_dump_block(m *base.Module, l0 int32)
//go:linkname F_plpgsql_yyparse github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_plpgsql_yyparse
func F_plpgsql_yyparse(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_read_sql_construct github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_read_sql_construct
func F_read_sql_construct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32
//go:linkname F_check_assignable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_check_assignable
func F_check_assignable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_word_is_not_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_word_is_not_variable
func F_word_is_not_variable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_cword_is_not_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cword_is_not_variable
func F_cword_is_not_variable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_read_into_scalar_list github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_read_into_scalar_list
func F_read_into_scalar_list(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_plpgsql_yylex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_plpgsql_yylex
func F_plpgsql_yylex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_internal_yylex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_internal_yylex
func F_internal_yylex(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_push_back_token github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_plpgsql_push_back_token
func F_plpgsql_push_back_token(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_plpgsql_append_source_text github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_plpgsql_append_source_text
func F_plpgsql_append_source_text(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_plpgsql_peek2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_plpgsql_peek2
func F_plpgsql_peek2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_plpgsql_scanner_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_plpgsql_scanner_errposition
func F_plpgsql_scanner_errposition(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_plpgsql_yyerror
func F_plpgsql_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_plpgsql_scanner_init github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plpgsql_scanner_init
func F_plpgsql_scanner_init(m *base.Module, l0 int32) int32
//go:linkname F_SN_create_env github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SN_create_env
func F_SN_create_env(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SN_close_env github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SN_close_env
func F_SN_close_env(m *base.Module, l0 int32, l1 int32)
//go:linkname F_r_e_ending_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_r_e_ending_2
func F_r_e_ending_2(m *base.Module, l0 int32) int32
//go:linkname F_r_undouble_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_r_undouble_3
func F_r_undouble_3(m *base.Module, l0 int32) int32
//go:linkname F_r_remove_suffix_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_r_remove_suffix_2
func F_r_remove_suffix_2(m *base.Module, l0 int32) int32
//go:linkname F_find_among github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_find_among
func F_find_among(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_among_b github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_among_b
func F_find_among_b(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_replace_s github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_replace_s
func F_replace_s(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_slice_del github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_slice_del
func F_slice_del(m *base.Module, l0 int32) int32
//go:linkname F_insert_s github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_insert_s
func F_insert_s(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_BinarySearchRange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BinarySearchRange
func F_BinarySearchRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__crypt_blowfish_rn github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__crypt_blowfish_rn
func F__crypt_blowfish_rn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_mbuf_append github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_mbuf_append
func F_mbuf_append(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mbuf_create_from_data github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mbuf_create_from_data
func F_mbuf_create_from_data(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pullf_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pullf_create
func F_pullf_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pullf_free github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pullf_free
func F_pullf_free(m *base.Module, l0 int32)
//go:linkname F_pullf_read github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pullf_read
func F_pullf_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pullf_read_fixed github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pullf_read_fixed
func F_pullf_read_fixed(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pullf_create_mbuf_reader github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pullf_create_mbuf_reader
func F_pullf_create_mbuf_reader(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pushf_create github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pushf_create
func F_pushf_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pushf_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pushf_free
func F_pushf_free(m *base.Module, l0 int32)
//go:linkname F_pushf_write github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pushf_write
func F_pushf_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pushf_flush github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pushf_flush
func F_pushf_flush(m *base.Module, l0 int32) int32
//go:linkname F_pushf_create_mbuf_writer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pushf_create_mbuf_writer
func F_pushf_create_mbuf_writer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_px_find_digest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_px_find_digest
func F_px_find_digest(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_px_find_cipher github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_px_find_cipher
func F_px_find_cipher(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgp_cfb_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgp_cfb_create
func F_pgp_cfb_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_pgp_cfb_free github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgp_cfb_free
func F_pgp_cfb_free(m *base.Module, l0 int32)
//go:linkname F_pgp_cfb_encrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgp_cfb_encrypt
func F_pgp_cfb_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pgp_compress_filter github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgp_compress_filter
func F_pgp_compress_filter(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgp_parse_pkt_hdr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgp_parse_pkt_hdr
func F_pgp_parse_pkt_hdr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pgp_create_pkt_reader github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgp_create_pkt_reader
func F_pgp_create_pkt_reader(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pgp_skip_packet github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgp_skip_packet
func F_pgp_skip_packet(m *base.Module, l0 int32) int32
//go:linkname F_pgp_expect_packet_end github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_pgp_expect_packet_end
func F_pgp_expect_packet_end(m *base.Module, l0 int32) int32
//go:linkname F_pgp_decrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgp_decrypt
func F_pgp_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgp_create_pkt_writer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgp_create_pkt_writer
func F_pgp_create_pkt_writer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_bytes_to_mpi github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bytes_to_mpi
func F_bytes_to_mpi(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgp_rsa_encrypt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgp_rsa_encrypt
func F_pgp_rsa_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgp_mpi_free github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgp_mpi_free
func F_pgp_mpi_free(m *base.Module, l0 int32) int32
//go:linkname F_pgp_mpi_write github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgp_mpi_write
func F_pgp_mpi_write(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_init_work github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_init_work
func F_init_work(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_create_secmsg github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_secmsg
func F_create_secmsg(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgp_key_free github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgp_key_free
func F_pgp_key_free(m *base.Module, l0 int32)
//go:linkname F__pgp_read_public_key github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__pgp_read_public_key
func F__pgp_read_public_key(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgp_s2k_read github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgp_s2k_read
func F_pgp_s2k_read(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgp_s2k_process github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgp_s2k_process
func F_pgp_s2k_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pgp_free github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgp_free
func F_pgp_free(m *base.Module, l0 int32) int32
//go:linkname F_px_THROW_ERROR github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_px_THROW_ERROR
func F_px_THROW_ERROR(m *base.Module, l0 int32)
//go:linkname F_px_debug github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_px_debug
func F_px_debug(m *base.Module, l0 int32, l1 int32)
//go:linkname F_citextcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_citextcmp
func F_citextcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_gin_extract_query_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gin_extract_query_trgm
func F_gin_extract_query_trgm(m *base.Module, l0 int32) int32
//go:linkname F_index_strategy_get_limit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_strategy_get_limit
func F_index_strategy_get_limit(m *base.Module, l0 int32) float64
//go:linkname F_generate_trgm github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_trgm
func F_generate_trgm(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_calc_word_similarity github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_calc_word_similarity
func F_calc_word_similarity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float32
//go:linkname F_hstoreCheckKeyLen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hstoreCheckKeyLen
func F_hstoreCheckKeyLen(m *base.Module, l0 int32) int32
//go:linkname F_hstoreCheckValLen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hstoreCheckValLen
func F_hstoreCheckValLen(m *base.Module, l0 int32) int32
//go:linkname F_hstorePairs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_hstorePairs
func F_hstorePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hstore_fetchval github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hstore_fetchval
func F_hstore_fetchval(m *base.Module, l0 int32) int32
//go:linkname F_hstoreArrayToPairs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_hstoreArrayToPairs
func F_hstoreArrayToPairs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_ltree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parse_ltree
func F_parse_ltree(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ltree_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ltree_concat
func F_ltree_concat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lca_inner github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lca_inner
func F_lca_inner(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makepol_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_makepol_2
func F_makepol_2(m *base.Module, l0 int32) int32
//go:linkname F_findoprnd_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_findoprnd_1
func F_findoprnd_1(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ltree_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_ltree_execute
func F_ltree_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gbt_num_consistent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_num_consistent
func F_gbt_num_consistent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_gbt_num_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gbt_num_picksplit
func F_gbt_num_picksplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gbt_var_bin_union github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_var_bin_union
func F_gbt_var_bin_union(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_gbt_var_node_cp_len github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gbt_var_node_cp_len
func F_gbt_var_node_cp_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gbt_var_penalty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gbt_var_penalty
func F_gbt_var_penalty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_gbt_var_picksplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gbt_var_picksplit
func F_gbt_var_picksplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_gin_btree_extract_query github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gin_btree_extract_query
func F_gin_btree_extract_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_validateConnectbyTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_validateConnectbyTupleDesc
func F_validateConnectbyTupleDesc(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_build_tuplestore_recursively github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_build_tuplestore_recursively
func F_build_tuplestore_recursively(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32)
//go:linkname F_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_execute
func F_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_infix_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_infix_3
func F_infix_3(m *base.Module, l0 int32, l1 int32)
//go:linkname F__int_unique github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__int_unique
func F__int_unique(m *base.Module, l0 int32) int32
//go:linkname F_new_intArrayType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_new_intArrayType
func F_new_intArrayType(m *base.Module, l0 int32) int32
//go:linkname F_isort github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isort
func F_isort(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_intarray_concat_arrays github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_intarray_concat_arrays
func F_intarray_concat_arrays(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DoubleMetaphone github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DoubleMetaphone
func F_DoubleMetaphone(m *base.Module, l0 int32, l1 int32)
//go:linkname F_write_box github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_write_box
func F_write_box(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_write_point_as_box github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_write_point_as_box
func F_write_point_as_box(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_cube_yylex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cube_yylex
func F_cube_yylex(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cube_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cube_yyensure_buffer_stack
func F_cube_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_cube_yy_create_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cube_yy_create_buffer
func F_cube_yy_create_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_yy_fatal_error_6 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_yy_fatal_error_6
func F_yy_fatal_error_6(m *base.Module, l0 int32)
//go:linkname F_cube_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_cube_yyerror
func F_cube_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_seg_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_seg_yyensure_buffer_stack
func F_seg_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_seg_yy_create_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_seg_yy_create_buffer
func F_seg_yy_create_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_yy_fatal_error_7 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_yy_fatal_error_7
func F_yy_fatal_error_7(m *base.Module, l0 int32)
//go:linkname F_seg_yyrestart github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_seg_yyrestart
func F_seg_yyrestart(m *base.Module, l0 int32, l1 int32)
//go:linkname F_seg_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_seg_yyerror
func F_seg_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_BloomFillMetapage github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_BloomFillMetapage
func F_BloomFillMetapage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstatindex_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstatindex_impl
func F_pgstatindex_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_relpages_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_relpages_impl
func F_pg_relpages_impl(m *base.Module, l0 int32) int64
//go:linkname F_pgstatginindex_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstatginindex_internal
func F_pgstatginindex_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_relation
func F_pgstat_relation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_uuid_generate_time github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_uuid_generate_time
func F_uuid_generate_time(m *base.Module, l0 int32)
//go:linkname F_uuid_unparse github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_uuid_unparse
func F_uuid_unparse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_amcheck_lock_relation_and_check github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_amcheck_lock_relation_and_check
func F_amcheck_lock_relation_and_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_PageGetItemIdCareful_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PageGetItemIdCareful_2
func F_PageGetItemIdCareful_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GetStrictOldestNonRemovableTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetStrictOldestNonRemovableTransactionId
func F_GetStrictOldestNonRemovableTransactionId(m *base.Module, l0 int32) int32
//go:linkname F_bt_index_block_validate github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bt_index_block_validate
func F_bt_index_block_validate(m *base.Module, l0 int32, l1 int64)
//go:linkname F_bt_page_print_tuples github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bt_page_print_tuples
func F_bt_page_print_tuples(m *base.Module, l0 int32) int32
//go:linkname F_verify_hash_page github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_verify_hash_page
func F_verify_hash_page(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_page_from_raw github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_page_from_raw
func F_get_page_from_raw(m *base.Module, l0 int32) int32
//go:linkname F_apw_detach_shmem github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_apw_detach_shmem
func F_apw_detach_shmem(m *base.Module, l0 int32, l1 int32)
//go:linkname F_apw_dump_now github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_apw_dump_now
func F_apw_dump_now(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgss_store github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgss_store
func F_pgss_store(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 float64, l6 int64, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32)
//go:linkname F_entry_reset github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_entry_reset
func F_entry_reset(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int64
//go:linkname F_qtext_load_file github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_qtext_load_file
func F_qtext_load_file(m *base.Module, l0 int32) int32
//go:linkname F_InitBitVector github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InitBitVector
func F_InitBitVector(m *base.Module, l0 int32) int32
//go:linkname F_CheckElement_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CheckElement_1
func F_CheckElement_1(m *base.Module, l0 int32)
//go:linkname F_CheckDim_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_CheckDim_1
func F_CheckDim_1(m *base.Module, l0 int32)
//go:linkname F_Float4ToHalf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_Float4ToHalf
func F_Float4ToHalf(m *base.Module, l0 float32) int32
//go:linkname F_HnswInitLockTranche github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HnswInitLockTranche
func F_HnswInitLockTranche(m *base.Module)
//go:linkname F_HnswUpdateNeighborsOnDisk github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HnswUpdateNeighborsOnDisk
func F_HnswUpdateNeighborsOnDisk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_HnswInsertAppendPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HnswInsertAppendPage
func F_HnswInsertAppendPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_tidhash_insert_hash_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_tidhash_insert_hash_internal
func F_tidhash_insert_hash_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_offsethash_insert_hash_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_offsethash_insert_hash_internal
func F_offsethash_insert_hash_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_HnswOptionalProcInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HnswOptionalProcInfo
func F_HnswOptionalProcInfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_HnswInitSupport github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_HnswInitSupport
func F_HnswInitSupport(m *base.Module, l0 int32, l1 int32)
//go:linkname F_HnswNewBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_HnswNewBuffer
func F_HnswNewBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_HnswInitElement github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HnswInitElement
func F_HnswInitElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 int32, l5 int32) int32
//go:linkname F_HnswGetMetaPageInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HnswGetMetaPageInfo
func F_HnswGetMetaPageInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_HnswGetEntryPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HnswGetEntryPoint
func F_HnswGetEntryPoint(m *base.Module, l0 int32) int32
//go:linkname F_HnswFindElementNeighbors github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HnswFindElementNeighbors
func F_HnswFindElementNeighbors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_HnswGetTypeInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HnswGetTypeInfo
func F_HnswGetTypeInfo(m *base.Module, l0 int32) int32
//go:linkname F_InitBuildState_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InitBuildState_2
func F_InitBuildState_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_VectorArraySet github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_VectorArraySet
func F_VectorArraySet(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_IvfflatKmeans github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_IvfflatKmeans
func F_IvfflatKmeans(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_VectorArrayInit github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_VectorArrayInit
func F_VectorArrayInit(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_VectorArrayFree github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_VectorArrayFree
func F_VectorArrayFree(m *base.Module, l0 int32)
//go:linkname F_IvfflatCheckNorm github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_IvfflatCheckNorm
func F_IvfflatCheckNorm(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IvfflatNormVectors github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_IvfflatNormVectors
func F_IvfflatNormVectors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_IvfflatCheckMemoryUsage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_IvfflatCheckMemoryUsage
func F_IvfflatCheckMemoryUsage(m *base.Module, l0 int32)
//go:linkname F_IvfflatInitRegisterPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_IvfflatInitRegisterPage
func F_IvfflatInitRegisterPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_IvfflatCommitBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_IvfflatCommitBuffer
func F_IvfflatCommitBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_IvfflatGetMetaPageInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_IvfflatGetMetaPageInfo
func F_IvfflatGetMetaPageInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_IvfflatUpdateList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IvfflatUpdateList
func F_IvfflatUpdateList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_IvfflatGetTypeInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_IvfflatGetTypeInfo
func F_IvfflatGetTypeInfo(m *base.Module, l0 int32) int32
//go:linkname F_CheckElement_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckElement_3
func F_CheckElement_3(m *base.Module, l0 float32)
//go:linkname F_CheckDim_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckDim_3
func F_CheckDim_3(m *base.Module, l0 int32)
//go:linkname F_R github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_R
func F_R(m *base.Module, l0 float64) float64
//go:linkname F___isspace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___isspace
func F___isspace(m *base.Module, l0 int32) int32
//go:linkname F_bsearch github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_bsearch
func F_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F___clock_gettime github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F___clock_gettime
func F___clock_gettime(m *base.Module, l0 int32, l1 int32)
//go:linkname F_close github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_close
func F_close(m *base.Module, l0 int32) int32
//go:linkname F___cos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___cos
func F___cos(m *base.Module, l0 float64, l1 float64) float64
//go:linkname F___rem_pio2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___rem_pio2
func F___rem_pio2(m *base.Module, l0 float64, l1 int32) int32
//go:linkname F___sin github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___sin
func F___sin(m *base.Module, l0 float64, l1 float64, l2 int32) float64
//go:linkname F___memcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___memcpy
func F___memcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___memset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___memset
func F___memset(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_time github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_time
func F_time(m *base.Module) int64
//go:linkname F_gettimeofday github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_gettimeofday
func F_gettimeofday(m *base.Module, l0 int32)
//go:linkname F___math_xflow github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___math_xflow
func F___math_xflow(m *base.Module, l0 int32, l1 float64) float64
//go:linkname F_fclose github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fclose
func F_fclose(m *base.Module, l0 int32) int32
//go:linkname F_fflush github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fflush
func F_fflush(m *base.Module, l0 int32) int32
//go:linkname F_do_getc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_do_getc
func F_do_getc(m *base.Module, l0 int32) int32
//go:linkname F_fgets github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fgets
func F_fgets(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fopen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fopen
func F_fopen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fread github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fread
func F_fread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_fscanf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fscanf
func F_fscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___fstatat github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___fstatat
func F___fstatat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_fsync github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fsync
func F_fsync(m *base.Module, l0 int32) int32
//go:linkname F_ftruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ftruncate
func F_ftruncate(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F___fwritex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___fwritex
func F___fwritex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getcwd github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_getcwd
func F_getcwd(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___getopt_msg github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___getopt_msg
func F___getopt_msg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_isalnum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isalnum
func F_isalnum(m *base.Module, l0 int32) int32
//go:linkname F_isdigit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isdigit
func F_isdigit(m *base.Module, l0 int32) int32
//go:linkname F_iswalnum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_iswalnum
func F_iswalnum(m *base.Module, l0 int32) int32
//go:linkname F_iswalpha github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_iswalpha
func F_iswalpha(m *base.Module, l0 int32) int32
//go:linkname F_iswspace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_iswspace
func F_iswspace(m *base.Module, l0 int32) int32
//go:linkname F_log github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_log
func F_log(m *base.Module, l0 float64) float64
//go:linkname F_mbstowcs github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_mbstowcs
func F_mbstowcs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_memchr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_memchr
func F_memchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_memcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_memcmp
func F_memcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_do_tzset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_do_tzset
func F_do_tzset(m *base.Module)
//go:linkname F___gmtime_r github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___gmtime_r
func F___gmtime_r(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___get_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___get_locale
func F___get_locale(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_open github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_open
func F_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pow github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pow
func F_pow(m *base.Module, l0 float64, l1 float64) float64
//go:linkname F_pread github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pread
func F_pread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_raise github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_raise
func F_raise(m *base.Module, l0 int32)
//go:linkname F_read github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_read
func F_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_readdir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_readdir
func F_readdir(m *base.Module, l0 int32) int32
//go:linkname F_readlink github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_readlink
func F_readlink(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rename github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_rename
func F_rename(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rmdir github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_rmdir
func F_rmdir(m *base.Module, l0 int32) int32
//go:linkname F_scalbn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_scalbn
func F_scalbn(m *base.Module, l0 float64, l1 int32) float64
//go:linkname F_sigemptyset github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sigemptyset
func F_sigemptyset(m *base.Module, l0 int32)
//go:linkname F_snprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_snprintf
func F_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_sscanf github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_sscanf
func F_sscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_stat github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_stat
func F_stat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strchr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strchr
func F_strchr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___strchrnul github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___strchrnul
func F___strchrnul(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcspn github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strcspn
func F_strcspn(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_is_leap github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_is_leap
func F_is_leap(m *base.Module, l0 int32) int32
//go:linkname F_strlcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strlcpy
func F_strlcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strlen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strlen
func F_strlen(m *base.Module, l0 int32) int32
//go:linkname F_strspn github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_strspn
func F_strspn(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strstr github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strstr
func F_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___shgetc github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___shgetc
func F___shgetc(m *base.Module, l0 int32) int32
//go:linkname F_fmodl github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_fmodl
func F_fmodl(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F_scanexp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_scanexp
func F_scanexp(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_strtof github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strtof
func F_strtof(m *base.Module, l0 int32, l1 int32) float32
//go:linkname F_strtod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strtod
func F_strtod(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_strtox_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strtox_2
func F_strtox_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64
//go:linkname F_strtol github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_strtol
func F_strtol(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___tan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___tan
func F___tan(m *base.Module, l0 float64, l1 float64, l2 int32) float64
//go:linkname F_toupper github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_toupper
func F_toupper(m *base.Module, l0 int32) int32
//go:linkname F_casemap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_casemap
func F_casemap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_unlink github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_unlink
func F_unlink(m *base.Module, l0 int32) int32
//go:linkname F_vfprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_vfprintf
func F_vfprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_vsnprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vsnprintf
func F_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_wcslen github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_wcslen
func F_wcslen(m *base.Module, l0 int32) int32
//go:linkname F_wcrtomb github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_wcrtomb
func F_wcrtomb(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_write github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_write
func F_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_emscripten_builtin_malloc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_emscripten_builtin_malloc
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32
//go:linkname F_emscripten_builtin_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_emscripten_builtin_free
func F_emscripten_builtin_free(m *base.Module, l0 int32)
//go:linkname F___addtf3 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___addtf3
func F___addtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___ashlti3 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___ashlti3
func F___ashlti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32)
//go:linkname F___divtf3 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___divtf3
func F___divtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___wasm_longjmp github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F___wasm_longjmp
func F___wasm_longjmp(m *base.Module, l0 int32, l1 int32)
//go:linkname F___lshrti3 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___lshrti3
func F___lshrti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32)
//go:linkname F___udivmodti4 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___udivmodti4
func F___udivmodti4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F_bind github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_bind
func F_bind(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_connect github.com/shibukawa/pgmem/internal/aot/pgaot/p2.F_connect
func F_connect(m *base.Module, l0 int32) int32
//go:linkname F_socket github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_socket
func F_socket(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13822 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13822
func Fn13822(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13825 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13825
func Fn13825(m *base.Module, l0 float32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname Fn13826 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13826
func Fn13826(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname Fn13830 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13830
func Fn13830(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13835 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13835
func Fn13835(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname Fn13842 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13842
func Fn13842(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13849 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13849
func Fn13849(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname Fn13850 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13850
func Fn13850(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname Fn13851 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13851
func Fn13851(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13854 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13854
func Fn13854(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname Fn13855 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13855
func Fn13855(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13858 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13858
func Fn13858(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13866 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13866
func Fn13866(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32)
//go:linkname Fn13868 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13868
func Fn13868(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13869 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13869
func Fn13869(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13880 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13880
func Fn13880(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13885 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13885
func Fn13885(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13886 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13886
func Fn13886(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13889 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13889
func Fn13889(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13891 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13891
func Fn13891(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13893 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13893
func Fn13893(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname Fn13901 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13901
func Fn13901(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname Fn13908 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13908
func Fn13908(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname Fn13913 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13913
func Fn13913(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13917 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13917
func Fn13917(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13918 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13918
func Fn13918(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname Fn13919 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13919
func Fn13919(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname Fn13924 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13924
func Fn13924(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname Fn13925 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13925
func Fn13925(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13929 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13929
func Fn13929(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn13930 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13930
func Fn13930(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname Fn13931 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13931
func Fn13931(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13938 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13938
func Fn13938(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13939 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13939
func Fn13939(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname Fn13945 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13945
func Fn13945(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13948 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13948
func Fn13948(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13952 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13952
func Fn13952(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13953 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13953
func Fn13953(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13954 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13954
func Fn13954(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13957 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13957
func Fn13957(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13958 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13958
func Fn13958(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13965 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13965
func Fn13965(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname Fn13966 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13966
func Fn13966(m *base.Module, l0 int64) int32
//go:linkname Fn13968 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.Fn13968
func Fn13968(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13971 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn13971
func Fn13971(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname Fn13972 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13972
func Fn13972(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname Fn13980 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn13980
func Fn13980(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13985 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn13985
func Fn13985(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname Fn13990 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn13990
func Fn13990(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn14009 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn14009
func Fn14009(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname Fn14010 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.Fn14010
func Fn14010(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn14015 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn14015
func Fn14015(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname Fn14018 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.Fn14018
func Fn14018(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname Fn14019 github.com/shibukawa/pgmem/internal/aot/pgaot/p2.Fn14019
func Fn14019(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn14020 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn14020
func Fn14020(m *base.Module, l0 int32, l1 int32) int32
//go:linkname Fn14026 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.Fn14026
func Fn14026(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
